package middleware

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("some-tempo-key")

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

func CompanyLoginRequired(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// session, err := sessionprovider.Store.Get(r, "authentication")
		// user, err := r.Cookie("company_id")
		// if err != nil {
		// 	w.WriteHeader(http.StatusUnauthorized)
		// 	w.Write([]byte("Status Unauthorized, Please login first"))
		// 	return
		// }
		next.ServeHTTP(w, r)
	}
}

func IsAuthorized(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] != nil {

			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}

				return mySigningKey, nil
			})
			// if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 	fmt.Println(claims)
			// }

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			fmt.Println(token)

			if token.Valid {
				next(w, r)
			}
		} else {

			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
