// go-swagger examples.
//
// The purpose of this application is to provide some
// use cases describing how to generate docs for your API
//
//     Schemes: http
//     Host: localhost:8080
//     BasePath: /
//     Version: 0.0.1
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
// swagger:meta
package main

import (
    "crypto/rand"
	"encoding/json"
    "encoding/base64"
    "io/ioutil"
	"fmt"
	"log"
    "os"
	"net/http"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
    "golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/gin-contrib/cors"
)

// A Character model.
//
// This is used for operations that want an character as body of the request
// swagger:parameters character
type Character struct {
	//in: body
	//required: true
	UserID		string 		`json:"user-id" form:"user-id"`
	Name    	string 		`json:"name" form:"name"`
	Title 		string 		`json:"title" form:"title"`
	Attributes 	[]string 	`json:"attributes" form:"attributes"`
	Level   	int 		`json:"level" form:"level"`
	Experience	int 		`json:"experience" form:"experience"`
}
//
// An Message model.
//
// This is used for operations that returns message
// swagger:response message
type Message struct {
	//in: body
	content string `json:"content"`
}

type Credentials struct {
	Web struct {
		Cid     string `json:"client_id"`
		Csecret string `json:"client_secret"`
	} `json:"web"`
}

type User struct {
    Sub string `json:"sub"`
    Name string `json:"name"`
    GivenName string `json:"given_name"`
    FamilyName string `json:"family_name"`
    Profile string `json:"profile"`
    Picture string `json:"picture"`
    Email string `json:"email"`
    EmailVerified string `json:"email_verified"`
    Gender string `json:"gender"`
}

type ClientState struct {
	State string `json:"state"`
	UserID string `json:"user-id"`
	ValidFrom string `json:"valid-from"`
}

var cred Credentials
var conf *oauth2.Config
var state string

func randToken() string {
	b := make([]byte, 32)
	rand.Read(b)
	return base64.StdEncoding.EncodeToString(b)
}

func init() {
	file, err := ioutil.ReadFile("client_secret_10258018262-dbsduq6hcpl9nvrjkf8uksv7ja9jqc6a.apps.googleusercontent.com.json")
    if err != nil {
        log.Printf("File error: %v\n", err)
        os.Exit(1)
	}
	json.Unmarshal(file, &cred)

    conf = &oauth2.Config{
        ClientID:     cred.Web.Cid,
        ClientSecret: cred.Web.Csecret,
        RedirectURL:  "http://127.0.0.1:8080/auth",
        Scopes: []string{
            "https://www.googleapis.com/auth/userinfo.email",
        },
        Endpoint: google.Endpoint,
    }
}

func getLoginURL(state string) string {
    return conf.AuthCodeURL(state)
}

func MiddleDB(mongo *mgo.Session) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("mongo", mongo)
		currentState := c.Query("state")
		var clientState ClientState
		s := mongo.DB("game").C("sessions")
		if (currentState == "") {
			clientState.State = randToken()
			err := s.Insert(&clientState)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Authorization error: could not create", "currentState" : currentState})
				c.Abort()
			}
			c.Set("state", clientState)
			c.Next()
		}
		log.Println("looking for "+currentState)
		err := s.Find(bson.M{"state": currentState}).One(&clientState)

		if err != nil {
			switch err {
			default:
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Authorization error: not found internal"})
				c.Abort()
			case mgo.ErrNotFound:
				c.JSON(http.StatusNotFound, gin.H{"error": "Authorization error: not found", "currentState" : currentState})
				c.Abort()
			}
		}
		log.Println("Found " + clientState.UserID + " asd")
		c.Set("state", clientState)
		c.Next()
	}
}

func AuthorizeRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		s, ok := c.Keys["state"].(ClientState)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not get state"})
		}
		v := s.UserID
		if v == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
		}
		c.Next()
	}
}


func authHandler(c *gin.Context) {
    retrievedState := c.Keys["state"].(ClientState)
	if retrievedState.State != c.Query("state") {
        c.AbortWithError(http.StatusUnauthorized, fmt.Errorf("Invalid session state: %s", retrievedState))
        return
    }

	tok, err := conf.Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
        return
	}

	client := conf.Client(oauth2.NoContext, tok)
	email, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
    if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
        return
	}
    defer email.Body.Close()
    data, _ := ioutil.ReadAll(email.Body)
	log.Println("Email body: ", string(data))
	var user User
	json.Unmarshal(data, &user)
	retrievedState.UserID = user.Email
	//s, ok := c.Keys["mongo"].(*mgo.Session)
	s := c.Keys["mongo"].(*mgo.Session)
	s = s.Copy()
	defer s.Close()

	
	sess := s.DB("game").C("sessions")
	log.Println("Try to upsert ", retrievedState.UserID)
	log.Println("By state ", retrievedState.State)
	err = sess.Update(bson.M{"state": retrievedState.State}, &retrievedState)
	if err != nil {
		if mgo.IsDup(err) {
			log.Println("dup error")
		}
		log.Println("Failed insert user: ", err)
	}
	
	c.Redirect(302, "http://127.0.0.1:8888/?code="+retrievedState.State+"&user="+user.Email)
    //c.Status(http.StatusOK)
}

func loginHandler(c *gin.Context) {
    state = randToken()
    session := sessions.Default(c)
    session.Set("state", state)
    session.Save()
    c.Writer.Write([]byte("<html><title>Golang Google</title> <body> <a href='" + getLoginURL(state) + "'><button>Login with Google!</button> </a> </body></html>"))
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:8888")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session)

	//c := session.DB("game").C("sessions")
	//store := sessions.NewMongoStore(c, 120, true, []byte("secret"))
	
	router := gin.Default()
   
	router.Use(CORSMiddleware())
	//config := cors.DefaultConfig()
	//config.AllowMethods = []string{"PUT", "PATCH", "POST", "DELETE", "GET"}
	router.Use(cors.Default())
	router.Use(MiddleDB(session))
	//router.Use(sessions.Sessions("rpggamesession", store))
	authorized := router.Group("/game")
	authorized.Use(AuthorizeRequest())
	{
		router.GET("/characters", allCharacters)
		router.POST("/characters", addCharacter)
		router.GET("/characters/:name", characterByName)
		router.PUT("/characters/:name", updateCharacter)
		router.DELETE("/characters/:name", deleteCharacter)
	}
	router.GET("/state", getState)
	router.GET("/getlogin", getLogin)
    router.GET("/login", loginHandler)
	router.GET("/auth", authHandler)
	router.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}
		session.Set("count", count)
		session.Save()
		c.JSON(200, gin.H{"count": count})
	})
	router.Run("127.0.0.1:8080")
}

func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("game").C("characters")

	index := mgo.Index{
		Key:        []string{"name", "user-id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func allCharacters(c *gin.Context) {
	s, ok := c.Keys["mongo"].(*mgo.Session)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "mongo is not ok"})
	}
	s = s.Copy()
	defer s.Close()

	character := s.DB("game").C("characters")

	var characters []Character
	err := character.Find(bson.M{"userid" : c.Keys["state"].(ClientState).UserID}).All(&characters)
	//err := character.Find(bson.M{}).All(&characters)
	log.Println("looking for "+c.Keys["state"].(ClientState).UserID+" characters")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		//ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
		//log.Println("Failed get all characters: ", err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "body": characters})
}

func getLogin(c *gin.Context) {
	state := c.Keys["state"].(ClientState)
	c.JSON(http.StatusOK, gin.H{"url": getLoginURL(state.State), "state": state.State, "user": state.UserID})
}

func getState(c *gin.Context) {
	retrievedState := c.Keys["state"].(ClientState)
	c.JSON(http.StatusOK, gin.H{"state": retrievedState})
}

// addCharacter swagger:route POST /characters characters character
//
// Handler to create a character.
//
// Responses:
//        200: message
//        400: message
//        401: message
func addCharacter(c *gin.Context) {
	s, ok := c.Keys["mongo"].(*mgo.Session)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "mongo is not ok"})
	}
	s = s.Copy()
	defer s.Close()

	var char Character
	if err := c.ShouldBindJSON(&char); err == nil {
		log.Println(char)
	} else {
		log.Println("still shit")
	}
	char.UserID = c.Keys["state"].(ClientState).UserID

	characters := s.DB("game").C("characters")
	log.Println(char)
	err := characters.Insert(char)
	if err != nil {
		if mgo.IsDup(err) {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Character with this Name already exists"})
			c.Abort()
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		log.Println("Failed insert character: ", err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"state": "ok"})
}

func characterByName(c *gin.Context){
	session := c.Keys["mongo"].(*mgo.Session)
	session = session.Copy()
	defer session.Close()

	//name := pat.Param(r, "name")

	chars := session.DB("game").C("characters")

	var character Character
	err := chars.Find(bson.M{"name": c.Param("name"), "userid": c.Keys["state"].(ClientState).UserID}).One(&character)

	if err != nil {
		switch err {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			log.Println("Failed find character: ", err)
			return
		case mgo.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"message": "Success", "body": character})
}

func updateCharacter(c *gin.Context) {
	session := c.Keys["mongo"].(*mgo.Session)
	session = session.Copy()
	defer session.Close()

	//name := pat.Param(r, "name")

	var character Character
	if err := c.ShouldBindJSON(&character); err == nil {
		log.Println(character)
	} else {
		log.Println("error")
	}

	chars := session.DB("game").C("characters")
	character.UserID = c.Keys["state"].(ClientState).UserID
	//err, id = c.Upsert(bson.M{"name": name}, &character)
	err := chars.Update(bson.M{"name": c.Param("name"), "userid": c.Keys["state"].(ClientState).UserID}, &character)
	if err != nil {
		switch err {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			log.Println("Failed update character: ", err)
			return
		case mgo.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
			return
		}
	}

	c.JSON(http.StatusNoContent, gin.H{})
}

func deleteCharacter(c *gin.Context) {
	session := c.Keys["mongo"].(*mgo.Session)
	session = session.Copy()
	defer session.Close()

	chars := session.DB("game").C("characters")

	log.Println(c.Param("name")+" " + c.Keys["state"].(ClientState).UserID)

	err := chars.Remove(bson.M{"name": c.Param("name"), "userid": c.Keys["state"].(ClientState).UserID})
	if err != nil {
		switch err {
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
			log.Println("Failed delete character: ", err)
			return
		case mgo.ErrNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
			return
		}
	}

	c.JSON(http.StatusNoContent, gin.H{})
}