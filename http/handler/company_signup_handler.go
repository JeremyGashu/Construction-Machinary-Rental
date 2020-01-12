package handlers

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

//UserSignupHandler -
type CompanySignupHandler struct {
	cpService admin.CompanyService
	tmpl      *template.Template
}

//NewUserSignupHandler -
func NewCompanySignUpHandler(usr admin.CompanyService, tpl *template.Template) *CompanySignupHandler {
	return &CompanySignupHandler{cpService: usr, tmpl: tpl}
}

//SignupHandler -
func (ush *CompanySignupHandler) SignupHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == http.MethodPost {
		ctg := entity.Company{}
		ctg.Name = r.FormValue("name")
		ctg.Address = r.FormValue("address")
		ctg.Password = r.FormValue("password")
		ctg.Description = r.FormValue("description")
		ctg.Email = r.FormValue("email")
		ctg.PhoneNo = r.FormValue("phone")

		err := ush.cpService.StoreCompany(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/company", http.StatusSeeOther)

	} else {

		ush.tmpl.ExecuteTemplate(w, "loginAsCompany.layout", nil)

	}
}
