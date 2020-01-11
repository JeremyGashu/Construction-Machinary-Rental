package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/ermiasgashu/Construction-Machinary-Rental/company"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

//CompanyMaterialHandler -
type CompanyMaterialHandler struct {
	materials company.MaterialService
}

//NewCompanyMaterialHandler -
func NewCompanyMaterialHandler(mat company.MaterialService) *CompanyMaterialHandler {
	return &CompanyMaterialHandler{materials: mat}
}

//Materials - GET /
func (ch *CompanyMaterialHandler) Materials(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	materials, err := ch.materials.Materials()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(materials)
}

//DeleteMaterial -
func (ch *CompanyMaterialHandler) DeleteMaterial(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("material_id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	err = ch.materials.DeleteMaterial(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	http.Redirect(w, r, "/v1/companies/materials", http.StatusSeeOther)
}

//StoreMaterial -
func (ch *CompanyMaterialHandler) StoreMaterial(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	material := &entity.Material{}

	err := json.Unmarshal(body, material)

	if err != nil {

		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	errs := ch.materials.AddMaterial(*material)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/admin/Users/%s", material.Name)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

//Material -
func (ch *CompanyMaterialHandler) Material(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("material_id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	material, err := ch.materials.Material(ID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
	json.NewEncoder(w).Encode(material)
}

//UpdateMaterial -
func (ch *CompanyMaterialHandler) UpdateMaterial(w http.ResponseWriter, r *http.Request) {

}
