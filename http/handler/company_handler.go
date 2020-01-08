package handler

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/ermiasgashu/Construction-Machinary-Rental/company"
)

//CompanyHandler -
type CompanyHandler struct {
	compService company.Service
	tmpl        *template.Template
	store       *sessions.CookieStore
}

//NewCompanyHandler -
func NewCompanyHandler(serv company.Service, tem *template.Template, str *sessions.CookieStore) *CompanyHandler {
	return &CompanyHandler{compService: serv, tmpl: tem}
}

//CompanyIndex -
func (ch *CompanyHandler) CompanyIndex(w http.ResponseWriter, r *http.Request) {
	ch.tmpl.ExecuteTemplate(w, "company.layout", nil)
}

//CompanySigup -
func (ch *CompanyHandler) CompanySigup(w http.ResponseWriter, r *http.Request) {
	ch.tmpl.ExecuteTemplate(w, "loginAsCompany.layout", nil)
}