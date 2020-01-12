package handlers

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

//UserSignupHandler -
type UserSignupHandler struct {
	usrService admin.UserService
	tmpl       *template.Template
}

//NewUserSignupHandler -
func NewUserSignupHandler(usr admin.UserService, tpl *template.Template) *UserSignupHandler {
	return &UserSignupHandler{usrService: usr, tmpl: tpl}
}

//SignupHandler -
func (ush *UserSignupHandler) SignupHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == http.MethodPost {
		ctg := entity.User{}
		ctg.Username = r.FormValue("username")
		ctg.FirstName = r.FormValue("firstname")
		ctg.LastName = r.FormValue("lastname")
		ctg.Email = r.FormValue("email")
		ctg.Phone = r.FormValue("phone")
		ctg.Password = r.FormValue("password")
		ctg.DeliveryAddress = r.FormValue("address")

		err := ush.usrService.StoreUser(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/user", http.StatusSeeOther)

	} else {

		ush.tmpl.ExecuteTemplate(w, "login.layout", nil)

	}
}
