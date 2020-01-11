package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/ermiasgashu/Construction-Machinary-Rental/company"
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
func (ch *CompanyMaterialHandler) StoreMaterial(w http.ResponseWriter, r *http.Request) {

}

//Material -
func (ch *CompanyMaterialHandler) Material(w http.ResponseWriter, r *http.Request) {

}

//UpdateMaterial -
func (ch *CompanyMaterialHandler) UpdateMaterial(w http.ResponseWriter, r *http.Request) {

}
