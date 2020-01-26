package middleware

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
)

var mySigningKey = []byte("some-tempo-key")

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

//LoginRequired -
func UserLoginRequired(next httprouter.Handle) httprouter.Handle { //middleware to check is user is logged in or not
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		coo, err := r.Cookie("auth-information")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)

		} else {
			if coo != nil {
				token, err := jwt.Parse(coo.Value, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret-key"), nil
				})
				if err != nil {
					fmt.Println(err)
				}
				if token.Valid {
					typ := token.Claims.(jwt.MapClaims)["auth-information"].(map[string]interface{})["Type"]
					if typ == "user" {
						next(w, r, ps)
					} else {
						http.Redirect(w, r, "/", http.StatusSeeOther)
					}

					// // fmt.Println("Valid Token Thanks")
					// fmt.Println(token.Claims.(jwt.MapClaims)["auth-information"].(map[string]interface{})["Value"]) // GETTING THE CLAIM AFTER LONG TRY

				} else {
					http.Redirect(w, r, "/", http.StatusSeeOther)
				}
			}
		}

	}
}

//CompaniesLoginReequired -
func CompaniesLoginReequired(next httprouter.Handle) httprouter.Handle { //middleware to check is user is logged in or not
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		coo, err := r.Cookie("auth-information")
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)

		} else {
			if coo != nil {
				token, err := jwt.Parse(coo.Value, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret-key"), nil
				})
				if err != nil {
					fmt.Println(err)
				}
				if token.Valid {
					typ := token.Claims.(jwt.MapClaims)["auth-information"].(map[string]interface{})["Type"]
					if typ == "provider" {
						next(w, r, ps)
					} else {
						http.Redirect(w, r, "/", http.StatusSeeOther)
					}

					// // fmt.Println("Valid Token Thanks")
					// fmt.Println(token.Claims.(jwt.MapClaims)["auth-information"].(map[string]interface{})["Value"]) // GETTING THE CLAIM AFTER LONG TRY

				} else {
					http.Redirect(w, r, "/", http.StatusSeeOther)
				}
			}
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
