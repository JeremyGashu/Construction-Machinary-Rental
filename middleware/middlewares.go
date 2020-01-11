package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
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

//CompanyLoginRequired -
func CompanyLoginRequired(next httprouter.Handle) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return []byte("secret-key"), nil
			})

			// fmt.Println(token)
			if err != nil {
				fmt.Println(err)
			}
			if token.Valid {
				// fmt.Println(token.Claims.(jwt.MapClaims)["logged-company"].(map[string]interface{})["email"])
				//GETTING THE CLAIM AFTER LONG TRY
				next(w, r, ps)
			}
			return
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.Write([]byte("Token is not sent"))
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
	}
}

// //CompanyLoginRequired -
// func CompanyLoginRequired(next http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
// 			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("There was an error")
// 			}
// 			return []byte("secret-key"), nil
// 		})

// 		// fmt.Println(token)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		if token.Valid {
// 			fmt.Println("Congrats...")
// 			next.ServeHTTP(w, r)
// 		}
// 		return
// 	}
// }
