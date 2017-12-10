package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"goji.io"
	"goji.io/pat"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "{message: %q}", message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}

type Character struct {
	Name    	string 		`json:"name"`
	Title 		string 		`json:"title"`
	Attributes 	[]string 	`json:"attributes"`
	Level   	int 		`json:"level"`
	Experience	int 		`json:"experience"`
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session)

	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/characters"), allCharacters(session))
	mux.HandleFunc(pat.Post("/characters"), addCharacter(session))
	mux.HandleFunc(pat.Get("/characters/:name"), characterByName(session))
	mux.HandleFunc(pat.Put("/characters/:name"), updateCharacter(session))
	mux.HandleFunc(pat.Delete("/characters/:name"), deleteCharacter(session))
	http.ListenAndServe("localhost:8080", mux)
}

func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("game").C("characters")

	index := mgo.Index{
		Key:        []string{"name"},
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

func allCharacters(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		c := session.DB("game").C("characters")

		var characters []Character
		err := c.Find(bson.M{}).All(&characters)
		if err != nil {
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed get all characters: ", err)
			return
		}

		respBody, err := json.MarshalIndent(characters, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		//log.Println("Pirma")
		log.Println("Automated building")
		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

func addCharacter(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		var character Character
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&character)
		if err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			return
		}

		c := session.DB("game").C("characters")

		err = c.Insert(character)
		if err != nil {
			if mgo.IsDup(err) {
				ErrorWithJSON(w, "Character with this Name already exists", http.StatusBadRequest)
				return
			}

			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed insert character: ", err)
			return
		}

		//fmt.Printf("%+v\n", character)
		//fmt.Printf("%+v\n", r.Body)
		
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Location", r.URL.Path+"/"+character.Name)
		w.WriteHeader(http.StatusCreated)
	}
}

func characterByName(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		name := pat.Param(r, "name")

		c := session.DB("game").C("characters")

		var character Character
		err := c.Find(bson.M{"name": name}).One(&character)

		if err != nil {
			switch err {
			default:
				ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
				log.Println("Failed find character: ", err)
				return
			case mgo.ErrNotFound:
				ErrorWithJSON(w, "Character not found", http.StatusNotFound)
				return
			}
		}

		/*if character.Name == "" {
			ErrorWithJSON(w, "Character not found", http.StatusNotFound)
			return
		}*/

		respBody, err := json.MarshalIndent(character, "", "  ")
		if err != nil {
			log.Fatal(err)
		}

		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

func updateCharacter(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		name := pat.Param(r, "name")

		var character Character
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&character)
		if err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			return
		}

		c := session.DB("game").C("characters")

		//err, id = c.Upsert(bson.M{"name": name}, &character)
		err = c.Update(bson.M{"name": name}, &character)
		if err != nil {
			switch err {
			default:
				ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
				log.Println("Failed update character: ", err)
				return
			case mgo.ErrNotFound:
				ErrorWithJSON(w, "Character not found", http.StatusNotFound)
				return
			}
		}

		w.WriteHeader(http.StatusNoContent)
	}
}

func deleteCharacter(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()

		name := pat.Param(r, "name")

		c := session.DB("game").C("characters")

		err := c.Remove(bson.M{"name": name})
		if err != nil {
			switch err {
			default:
				ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
				log.Println("Failed delete character: ", err)
				return
			case mgo.ErrNotFound:
				ErrorWithJSON(w, "Character not found", http.StatusNotFound)
				return
			}
		}

		w.WriteHeader(http.StatusNoContent)
	}
}