package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	jwtreq "github.com/dgrijalva/jwt-go/request"
	"crypto/rsa"
	"encoding/json"
)

const (
	privKeyPath = "app.rsa"
	pubKeyPath = "app.rsa.pub"
)

//var VerifyKey, SignKey []byte
var (
	VerifyKey *rsa.PublicKey
	SignKey   *rsa.PrivateKey
)


//struct User for parsing login credentials
type User struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// UserClaims is a set of JWT claims that contain UserProfile.
type UserClaims struct {
	Profile User           `json:"profile"`
	jwt.StandardClaims
}

var claims UserClaims

func initKeys() {
	var err error

	signBytes, err := ioutil.ReadFile(privKeyPath)

	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err !=nil{
		fmt.Println("pvt key not read")
		return
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err !=nil{
		fmt.Println("pub key not read")
		return
	}
}

func main() {

	initKeys()

	router := mux.NewRouter()
	router.HandleFunc("/login" , LoginHandler)
	router.HandleFunc("/auth" , ValidateTokenMiddleware)

	server := &http.Server{
		Addr: ":8000",
		Handler: router,
	}
	log.Println("Listening...")
	server.ListenAndServe()
}


func LoginHandler(w http.ResponseWriter , r *http.Request) {
	var user User

	err:=json.NewDecoder(r.Body).Decode(&user)
	if err!=nil {
		fmt.Println("Error reading username and/or Password")
	}

	fmt.Println(user)

	// validate user credentials
	if user.UserName != "alex" && user.Password != "12345" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintln(w, "Wrong credintials")
		return
	}

	//set claims
	claims = UserClaims{
		User{UserName: user.UserName , Password: "secret"},
		jwt.StandardClaims{
			Issuer: "administrator007",
			Subject: "Testing",
		},
	}

	fmt.Println(claims)

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	ss, err := token.SignedString(SignKey)
	if err != nil {
		fmt.Println("Error while Signing Token!")
		return
	}

	fmt.Println(ss)
	fmt.Fprintln(w,ss)

}

//AUTH TOKEN VALIDATION
func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request) {

	//validate token
	token, err := jwtreq.ParseFromRequestWithClaims(r, jwtreq.AuthorizationHeaderExtractor,&claims,func(token *jwt.Token) (interface{}, error){
		return VerifyKey, nil
	})

	if err == nil {

		if token.Valid{
			fmt.Println("Token is valid")
			ProtectedHandler(w,r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorised access to this resource")
	}

}


func ProtectedHandler(w http.ResponseWriter, r *http.Request){

	fmt.Fprintln(w,"Access to DB Granted !")

}