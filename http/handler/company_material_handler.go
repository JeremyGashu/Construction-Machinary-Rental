package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/ermiasgashu/Construction-Machinary-Rental/company"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

// CompanyMaterialHandler handles Company handler Company requests
type CompanyMaterialHandler struct {
	tmpl        *template.Template
	MaterialSrv company.MaterialService
}

// NewCompanyMaterialHandler initializes and returns new CompanyMaterialHandler
func NewCompanyMaterialHandler(T *template.Template, CS company.MaterialService) *CompanyMaterialHandler {
	return &CompanyMaterialHandler{tmpl: T, MaterialSrv: CS}
}

// CompanyMaterials handle requests on route /Company/Materials
func (ach *CompanyMaterialHandler) CompanyMaterials(w http.ResponseWriter, r *http.Request) {
	Materials, err := ach.MaterialSrv.Materials()
	if err != nil {
		panic(err)
		fmt.Println(err)
	}
	ach.tmpl.ExecuteTemplate(w, "company.material.layout", Materials)
}

// CompanyMaterialsNew hanlde requests on route /Company/Materials/new
func (ach *CompanyMaterialHandler) CompanyMaterialsNew(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		ctg := entity.Material{}
		ctg.Name = r.FormValue("name")
		// own, _ := strconv.Atoi(r.FormValue("owner"))
		// ctg.Owner = own
		ppd, _ := strconv.ParseFloat(r.FormValue("priceperday"), 10)
		ctg.PricePerDay = ppd
		ondiscount, _ := strconv.ParseBool(r.FormValue("ondiscount"))
		ctg.OnDiscount = ondiscount
		discout, _ := strconv.ParseFloat(r.FormValue("discount"), 10)
		ctg.Discount = float32(discout)
		onsal, _ := strconv.ParseBool(r.FormValue("onsale"))
		ctg.OnSale = onsal
		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		ctg.ImagePath = fh.Filename

		writeFile(&mf, fh.Filename)

		err = ach.MaterialSrv.AddMaterial(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/company/material", http.StatusSeeOther)

	} else {

		ach.tmpl.ExecuteTemplate(w, "company.material.new.layout", nil)

	}
}

// CompanyMaterialsUpdate handle requests on /Company/Materials/update
func (ach *CompanyMaterialHandler) CompanyMaterialsUpdate(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		cat, err := ach.MaterialSrv.Material(id)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}

		ach.tmpl.ExecuteTemplate(w, "company.material.update.layout", cat)

	} else if r.Method == http.MethodPost {

		ctg := entity.Material{}
		ctg.Name = r.FormValue("name")
		own, _ := strconv.Atoi(r.FormValue("owner"))
		ctg.Owner = own
		ppd, _ := strconv.ParseFloat(r.FormValue("priceperday"), 10)
		ctg.PricePerDay = ppd
		ondiscount, _ := strconv.ParseBool(r.FormValue("ondiscount"))
		ctg.OnDiscount = ondiscount
		discout, _ := strconv.ParseFloat(r.FormValue("discount"), 10)
		ctg.Discount = float32(discout)
		onsal, _ := strconv.ParseBool(r.FormValue("onsale"))
		ctg.OnSale = onsal
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

		err = ach.MaterialSrv.UpdateMaterial(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/material", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/material", http.StatusSeeOther)
	}

}

// CompanyMaterialsDelete handle requests on route /Company/categories/delete
func (ach *CompanyMaterialHandler) CompanyMaterialsDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		idRaw := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idRaw)

		if err != nil {
			panic(err)
		}

		err = ach.MaterialSrv.DeleteMaterial(id)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/company/material", http.StatusSeeOther)
}
