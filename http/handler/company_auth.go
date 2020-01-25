package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"

	"github.com/julienschmidt/httprouter"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
)

//CompanyInfo -
type CompanyInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//CompanyAuthHandler -
type CompanyAuthHandler struct {
	company admin.CompanyService
}

//NewCompanyAuthHandler -
func NewCompanyAuthHandler(cmp admin.CompanyService) *CompanyAuthHandler {
	return &CompanyAuthHandler{company: cmp}
}

//Login -
func (cah *CompanyAuthHandler) Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) { //for the api
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	comp := &CompanyInfo{}

	err := json.Unmarshal(body, comp)
	if err != nil {

		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	authenticated := cah.company.AuthCompany(comp.Email, comp.Password)

	if !authenticated {

		// fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		// w.Write([]byte("Invalid Info"))
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	tokenStirng, err := GenerateJWT(*comp)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		// w.Write([]byte("Token Generation Failed"))
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"Token:\"" + tokenStirng + "}")) //set the token as cookie
	// http.Error(w, http.StatusText(http.StatusAccepted), http.StatusAccepted)

	// fmt.Println(tokenStirng)
	// fmt.Println(r.Header["Token"][0]) to access the token generated
}

//TestJWT - test route
func (cah *CompanyAuthHandler) TestJWT(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("Welcome You are authorized..."))
	// http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	// return
}

//GenerateJWT -Generate some random token
func GenerateJWT(data CompanyInfo) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["logged-company"] = data //TODO Change it
	tokenString, err := token.SignedString([]byte("secret-key"))
	if err != nil {
		return tokenString, err
	}
	return tokenString, nil
}
