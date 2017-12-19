package main

import (
	//"net/http"
	"reflect"
	"testing"

	//"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2"
)

func Test_randToken(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randToken(); got != tt.want {
				t.Errorf("randToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLoginURL(t *testing.T) {
	type args struct {
		state string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "pregenerated state",
			args: args{
				state:"M2DEnQA4qDhdOHzeDl2ppcMO8IPS83z1Ogswz8bOzwg=",
			},
			want: "https://accounts.google.com/o/oauth2/auth?client_id=10258018262-dbsduq6hcpl9nvrjkf8uksv7ja9jqc6a.apps.googleusercontent.com&redirect_uri=http%3A%2F%2F127.0.0.1%3A8080%2Fauth&response_type=code&scope=https%3A%2F%2Fwww.googleapis.com%2Fauth%2Fuserinfo.email&state=M2DEnQA4qDhdOHzeDl2ppcMO8IPS83z1Ogswz8bOzwg%3D",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLoginURL(tt.args.state); got != tt.want {
				t.Errorf("getLoginURL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMiddleDB(t *testing.T) {
	type args struct {
		mongo *mgo.Session
	}
	tests := []struct {
		name string
		args args
		want gin.HandlerFunc
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MiddleDB(tt.args.mongo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MiddleDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthorizeRequest(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AuthorizeRequest(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthorizeRequest() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_authHandler(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			authHandler(tt.args.c)
		})
	}
}

func Test_loginHandler(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			loginHandler(tt.args.c)
		})
	}
}

func TestCORSMiddleware(t *testing.T) {
	tests := []struct {
		name string
		want gin.HandlerFunc
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CORSMiddleware(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CORSMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ensureIndex(t *testing.T) {
	type args struct {
		s *mgo.Session
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ensureIndex(tt.args.s)
		})
	}
}

func Test_allCharacters(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allCharacters(tt.args.c)
		})
	}
}

func Test_getLogin(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getLogin(tt.args.c)
		})
	}
}

func Test_getState(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getState(tt.args.c)
		})
	}
}

func Test_addCharacter(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			addCharacter(tt.args.c)
		})
	}
}

func Test_characterByName(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			characterByName(tt.args.c)
		})
	}
}

func Test_updateCharacter(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateCharacter(tt.args.c)
		})
	}
}

func Test_deleteCharacter(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deleteCharacter(tt.args.c)
		})
	}
}
