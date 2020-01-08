package middleware

import (
	"fmt"
	"net/http"

	"github.com/ermiasgashu/Construction-Machinary-Rental/sessionprovider"
)

//UserLoginRequired - middleware to check if the user us authenticated
func UserLoginRequired(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, err := sessionprovider.Store.Get(r, "authentication")
		if err != nil {
			panic(err)
		}
		value, ok := session.Values["user_username"]

		if !ok {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		fmt.Println("The logged in user's username is ", value)
		// http.Redirect(w, r, "/users", http.StatusSeeOther)
		handler.ServeHTTP(w, r)
	}
}
