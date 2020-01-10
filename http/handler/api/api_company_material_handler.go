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

//AllMaterials - GET /
func (ch *CompanyMaterialHandler) AllMaterials(w http.ResponseWriter, r *http.Request) {

}

func (ch *CompanyMaterialHandler) DeleteMaterial(w http.ResponseWriter, r *http.Request) {

}
