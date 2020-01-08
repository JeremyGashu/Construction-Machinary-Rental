package middleware

import (
	"fmt"
	"net/http"

	"github.com/ermiasgashu/Construction-Machinary-Rental/sessionprovider"
)

//UserLoginRequired - middleware to check if the user us authenticated
func UserLoginRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//TODO Check is the user session is created
		session, err := sessionprovider.Store.Get(r, "authentication")
		if err != nil {
			panic(err)
		}
		value, ok := session.Values["user_username"]
		if !ok {
			fmt.Println("The user is not logged in...")

		} else {
			fmt.Println("The logged in user's username is ", value)
		}

		handler.ServeHTTP(w, r)
	}
}
