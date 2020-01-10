package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
	"github.com/julienschmidt/httprouter"
)

// AdminCompanyHandler handles Company related http requests
type AdminCompanyHandler struct {
	CompanyService admin.CompanyService
}

// NewAdminCompanyHandler returns new AdminCompanyHandler object
func NewAdminCompanyHandler(cmntService admin.CompanyService) *AdminCompanyHandler {
	return &AdminCompanyHandler{CompanyService: cmntService}
}

// GetCompanys handles GET /v1/admin/Companys request
func (ach *AdminCompanyHandler) GetCompanys(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	Companys, errs := ach.CompanyService.Companies()

	if errs != nil {
		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(Companys, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return

}

// GetSingleCompany handles GET /v1/admin/Companys/:id request
func (ach *AdminCompanyHandler) GetSingleCompany(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")

		fmt.Println(err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	Company, errs := ach.CompanyService.Company(id)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(Company, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// PostCompany handles POST /v1/admin/Companys request
func (ach *AdminCompanyHandler) PostCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	Company := &entity.Company{}

	err := json.Unmarshal(body, Company)

	if err != nil {

		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	errs := ach.CompanyService.StoreCompany(*Company)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/admin/Companys/%d", Company.CompanyID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// PutCompany handles PUT /v1/admin/Companys/:id request
func (ach *AdminCompanyHandler) PutCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	Company, errs := ach.CompanyService.Company(id)

	if errs != nil {

		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &Company)

	errs = ach.CompanyService.UpdateCompany(Company)

	if errs != nil {

		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(Company, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// DeleteCompany handles DELETE /v1/admin/Companys/:id request
func (ach *AdminCompanyHandler) DeleteCompany(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {

		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	errs := ach.CompanyService.DeleteCompany(id)

	if errs != nil {

		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
