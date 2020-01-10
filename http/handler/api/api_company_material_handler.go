package api

import (
	"net/http"

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
func (ch *CompanyMaterialHandler) Materials(w http.ResponseWriter, r *http.Request) {

}

//DeleteMaterial -
func (ch *CompanyMaterialHandler) DeleteMaterial(w http.ResponseWriter, r *http.Request) {

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
