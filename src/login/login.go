package login

import (
	"context"
	"db"
	"encoding/json"
	"exception"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"time"
)

const SecretKey = "SERVER_PRIVATE_KXY"

/**
TODO read database
*/
func Handler(w http.ResponseWriter, r *http.Request) {
	var userLoginData UserCredentials
	err := json.NewDecoder(r.Body).Decode(&userLoginData)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "request data has error.")
		return
	}

	filter := bson.D{{"userName", userLoginData.UserName}}

	var result User
	err = db.GetInstance().Collection("users").FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "request data has error.")
		return
	}

	// TODO 驗證密碼
	if userLoginData.PxssCode != "" {
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("Error logging in")
		fmt.Fprint(w, "Invalid credentials")
	}

	// create token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		exception.Fatal(err)
	}
	response := Token{tokenString}
	JsonResponse(response, w)
}

/**
jwt 驗證.
*/
func ValidateTokenMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (i interface{}, err error) {
			return []byte(SecretKey), nil
		})

	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}

}

/**

 */
func JsonResponse(response interface{}, w http.ResponseWriter) {

	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
