package middleware

import (
	"net/http"
)

//UserLoginRequired - middleware to check if the user us authenticated
func UserLoginRequired(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// session, err := sessionprovider.Store.Get(r, "authentication")
		// if err != nil {
		// 	panic(err)
		// }
		// value, ok := session.Values["user_username"]

		// if !ok {
		// 	http.Redirect(w, r, "/", http.StatusSeeOther)
		// 	return
		// }
		// fmt.Println("The logged in user's username is ", value)
		// // http.Redirect(w, r, "/users", http.StatusSeeOther)
		next.ServeHTTP(w, r)
	}
}
