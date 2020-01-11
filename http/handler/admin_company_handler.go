package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
	"github.com/julienschmidt/httprouter"
)

// AdminCompanyHandler handles Admin handler admin requests
type AdminCompanyHandler struct {
	tmpl       *template.Template
	CompanySrv admin.CompanyService
}

// NewAdminCompanyHandler initializes and returns new AdminCompanyHandler
func NewAdminCompanyHandler(T *template.Template, CS admin.CompanyService) *AdminCompanyHandler {
	return &AdminCompanyHandler{tmpl: T, CompanySrv: CS}
}

// AdminCompanys handle requests on route /admin/Companys
func (ach *AdminCompanyHandler) AdminCompanys(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Companys, err := ach.CompanySrv.Companies()
	if err != nil {
		panic(err)
	}
	ach.tmpl.ExecuteTemplate(w, "admin.company.layout", Companys)
}

// AdminCompanysNew hanlde requests on route /admin/Companys/new
func (ach *AdminCompanyHandler) AdminCompanysNew(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if r.Method == http.MethodPost {

		ctg := entity.Company{}
		ctg.Name = r.FormValue("name")
		ctg.Description = r.FormValue("description")
		ctg.Email = r.FormValue("email")
		ctg.PhoneNo = r.FormValue("phone")
		ctg.Password = r.FormValue("password")
		ctg.Address = r.FormValue("address")
		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		ctg.ImagePath = fh.Filename

		writeFile(&mf, fh.Filename)

		err = ach.CompanySrv.StoreCompany(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/company", http.StatusSeeOther)

	} else {

		ach.tmpl.ExecuteTemplate(w, "admin.company.new.layout", nil)

	}
}

// AdminCompanysUpdate handle requests on /admin/Companys/update
func (ach *AdminCompanyHandler) AdminCompanysUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		cat, err := ach.CompanySrv.Company(id)
		if err != nil {
			panic(err)
		}

		ach.tmpl.ExecuteTemplate(w, "admin.company.update.layout", cat)

	} else if r.Method == http.MethodPost {

		ctg := entity.Company{}
		ctg.Name = r.FormValue("name")
		ctg.Description = r.FormValue("description")
		ctg.Email = r.FormValue("email")
		ctg.Password = r.FormValue("password")
		ctg.Address = r.FormValue("address")
		mf, fh, err := r.FormFile("catimg")
		if mf != nil {
			ctg.ImagePath = fh.Filename

			if err != nil {
				panic(err)
			}

			defer mf.Close()

			writeFile(&mf, ctg.ImagePath)

			fmt.Println(ctg.ImagePath)
		} else {
			ctg.ImagePath = r.FormValue("catimg")
		}

		err = ach.CompanySrv.UpdateCompany(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/company", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/company", http.StatusSeeOther)
	}

}

// AdminCompanysDelete handle requests on route /admin/categories/delete
func (ach *AdminCompanyHandler) AdminCompanysDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		err = ach.CompanySrv.DeleteCompany(id)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/admin/company", http.StatusSeeOther)
}
