package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/ermiasgashu/Construction-Machinary-Rental/company"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
	"github.com/julienschmidt/httprouter"
)

//UserMaterialHandler - User material handler
type UserMaterialHandler struct {
	materialService company.MaterialService
	tmpl            *template.Template
}

//NewUserMaterialHandler -
func NewUserMaterialHandler(ms company.MaterialService, t *template.Template) *UserMaterialHandler {
	return &UserMaterialHandler{materialService: ms, tmpl: t}
}

//Materials -
func (mh *UserMaterialHandler) Materials(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", "http://localhost:8080/v1/companies/materials", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var materials []entity.Material
	err = json.Unmarshal(body, &materials)
	if err != nil {
		fmt.Println(err)
	}
	mh.tmpl.ExecuteTemplate(w, "user.layout", materials)
}
