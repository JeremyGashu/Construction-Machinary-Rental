package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
)

//AuthHandler -
type AuthHandler struct {
	company admin.CompanyService
	user    admin.UserService
	admin   admin.AdminService
}

//Info -
type Info struct {
	Type      string
	Value     string
	ID        int
	Activated bool
}

//NewAuthHander -
func NewAuthHander(c admin.CompanyService, u admin.UserService, adm admin.AdminService) *AuthHandler {
	return &AuthHandler{company: c, user: u, admin: adm}
}

//Login -
func (ah *AuthHandler) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	loggedAs := r.FormValue("type")
	username := r.FormValue("username")
	password := r.FormValue("pass1")
	// fmt.Println(loggedAs, username, password)
	if loggedAs == "user" {
		authenticated := ah.user.AuthUser(username, password)
		if authenticated {

			userIn := Info{Type: "user", Value: username}
			token, _ := GenerateToken(userIn)
			// fmt.Println(token)

			// w.Header().Add("authorization", token)
			coo := http.Cookie{
				Name:    "auth-information",
				Value:   token,
				Expires: time.Now().Add(time.Hour * 2),
			}
			http.SetCookie(w, &coo)
			// w.Header().Add("authorization", token)
			// fmt.Println("Cookie set and redirecting...")
			http.Redirect(w, r, "/user", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}

		// fmt.Println(c) //TODO add user login session
	} else if loggedAs == "provider" {
		authenticated := ah.company.AuthCompany(username, password)
		if authenticated {
			comp, err := ah.company.CompanyByEmail(username)
			if err != nil {
				fmt.Println(err)
			}
			userIn := Info{Type: "provider", Value: username, ID: comp.CompanyID, Activated: comp.Activated}
			// fmt.Println(userIn)

			token, _ := GenerateToken(userIn) //token generated

			coo := http.Cookie{
				Name:    "auth-information",
				Value:   token,
				Expires: time.Now().Add(time.Hour * 2),
			} //token saved as cookie in jwt foormat which is more secured
			http.SetCookie(w, &coo)

			// w.Header().Add("authorization", token)
			http.Redirect(w, r, "/company", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	} else if loggedAs == "admin" {
		authenticated := ah.admin.AuthAdmin(username, password)
		// fmt.Println(authenticated)
		if authenticated {

			userIn := Info{Type: "admin", Value: username}

			token, _ := GenerateToken(userIn) //token generated

			coo := http.Cookie{
				Name:    "auth-information",
				Value:   token,
				Expires: time.Now().Add(time.Hour * 2),
			} //token saved as cookie in jwt foormat which is more secured
			http.SetCookie(w, &coo)

			// w.Header().Add("authorization", token)
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

//Logout -
func (ah *AuthHandler) Logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	c := http.Cookie{
		Name:    "auth-information",
		Value:   "",
		Expires: time.Now().Add(-1 * time.Millisecond),
	}
	http.SetCookie(w, &c)
	// fmt.Println("Cookie set and redirecting...")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//GenerateToken - generated client side token to make secure connection
func GenerateToken(data Info) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["auth-information"] = data
	tokenString, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
