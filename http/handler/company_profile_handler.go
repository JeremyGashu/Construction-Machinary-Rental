package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/company"
	"github.com/julienschmidt/httprouter"
)

//CompanyProfileHandler -
type CompanyProfileHandler struct {
	comp            admin.CompanyService
	materialService company.MaterialService
	tmpl            *template.Template
}

//NewCompanyProfileHandler -
func NewCompanyProfileHandler(ms company.MaterialService, t *template.Template, usr admin.CompanyService) *CompanyProfileHandler {
	return &CompanyProfileHandler{materialService: ms, tmpl: t, comp: usr}
}

//ProfileIndex -
func (cph *CompanyProfileHandler) ProfileIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := GetUserFromJWT(r)
	if err != nil {
		fmt.Println(err)
	}
	usr, err := cph.comp.Company(int(user))
	fmt.Println(usr.Rating)
	if err != nil {
		fmt.Println(err)
	}
	cph.tmpl.ExecuteTemplate(w, "company.edit.layout", usr)
}

//UpdateProfile -
func (cph *CompanyProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := GetUserFromJWT(r)
	if err != nil {
		fmt.Println(err)
	}
	usr, err := cph.comp.Company(int(user))
	someUser := usr
	if err != nil {
		fmt.Println(err)
	}
	phone := r.FormValue("phone")
	name := r.FormValue("name")
	address := r.FormValue("address")
	description := r.FormValue("description")
	mf, fh, err := r.FormFile("catimg")
	if mf != nil {
		usr.ImagePath = fh.Filename

		if err != nil {
			panic(err)
		}

		defer mf.Close()

		writeFile(&mf, usr.ImagePath)

		fmt.Println(usr.ImagePath)
	}
	usr.Name = name
	usr.Description = description
	usr.PhoneNo = phone
	usr.Address = address

	err = cph.comp.UpdateCompany(usr)
	if err != nil {
		cph.tmpl.ExecuteTemplate(w, "company.edit.layout", someUser)
	}
	http.Redirect(w, r, "/company", http.StatusSeeOther)
}
