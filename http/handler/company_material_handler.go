package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/company"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
	"github.com/julienschmidt/httprouter"
)

// CompanyMaterialHandler handles Company handler Company requests
type CompanyMaterialHandler struct {
	comp        admin.CompanyService
	tmpl        *template.Template
	MaterialSrv company.MaterialService
}

// NewCompanyMaterialHandler initializes and returns new CompanyMaterialHandler
func NewCompanyMaterialHandler(T *template.Template, CS company.MaterialService, c admin.CompanyService) *CompanyMaterialHandler {
	return &CompanyMaterialHandler{tmpl: T, MaterialSrv: CS, comp: c}
}

//CompanyIndex - Home page for logged companies..
func (ach *CompanyMaterialHandler) CompanyIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	usr, err := GetUserFromJWT(r)
	if err != nil {
		fmt.Println(err)
	}
	user, err := ach.comp.Company(int(usr))
	// fmt.Println(user)
	if err != nil {
		fmt.Println(err)
	}

	ach.tmpl.ExecuteTemplate(w, "company.layout", user)
}

// CompanyMaterials handle requests on route /Company/Materials
func (ach *CompanyMaterialHandler) CompanyMaterials(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := GetUserFromJWT(r)
	id := int(user)
	if err != nil {
		fmt.Println(err)
	}
	url := fmt.Sprintf("http://localhost:8080/v1/companies/owner/%d/materials", id)
	client := http.DefaultClient
	req, err := http.NewRequest("GET", url, nil)

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
	// fmt.Println(url)
	if err != nil {
		panic(err)

	}
	ach.tmpl.ExecuteTemplate(w, "company.material.layout", materials)
}

// CompanyMaterialsNew hanlde requests on route /Company/Materials/new
func (ach *CompanyMaterialHandler) CompanyMaterialsNew(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if r.Method == http.MethodPost {
		user, err := GetUserFromJWT(r)
		
		if err != nil {
			fmt.Println(err)
		}
		id := int(user)
		ctg := entity.Material{}
		ctg.Name = r.FormValue("name")
		// own, _ := strconv.Atoi(r.FormValue("owner"))
		// ctg.Owner = own
		ppd, err := strconv.ParseFloat(r.FormValue("priceperday"), 10)
		if err != nil {

			fmt.Println(1, err)
		}
		ctg.PricePerDay = ppd
		// ondiscount, _ := strconv.ParseBool(r.FormValue("ondiscount"))
		ctg.OnDiscount = false
		// discout := 0
		ctg.Discount = 0
		onsal, _ := strconv.ParseBool(r.FormValue("sellStatus"))
		ctg.OnSale = onsal
		ctg.Owner = id
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

		http.Redirect(w, r, "/company", http.StatusSeeOther)

	} else {

		ach.tmpl.ExecuteTemplate(w, "company.material.new.layout", nil)

	}
}

// CompanyMaterialsUpdate handle requests on /Company/Materials/update
func (ach *CompanyMaterialHandler) CompanyMaterialsUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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
		rawID := r.FormValue("ID")
		matid, err := strconv.Atoi(rawID)
		if err != nil {
			fmt.Println(err)
		}

		user, err := GetUserFromJWT(r)

		if err != nil {
			fmt.Println(err)
		}
		id := int(user)

		ctg := entity.Material{}
		ctg.Name = r.FormValue("name")
		ctg.Owner = id
		ppd, err := strconv.ParseFloat(r.FormValue("priceperday"), 10)
		if err != nil {
			fmt.Println(err)
		}
		onsal, err := strconv.ParseBool(r.FormValue("sellStatus"))
		if err != nil {
			fmt.Println(err)
		}
		ctg.OnSale = onsal
		ctg.PricePerDay = ppd
		ctg.ID = matid
		ond, err := strconv.ParseBool(r.FormValue("discountStatus"))
		if err != nil {
			fmt.Println(err)
		}
		ctg.OnSale = ond

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

		fmt.Println(ctg)
		err = ach.MaterialSrv.UpdateMaterial(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/company", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/company", http.StatusSeeOther)
	}

}

// CompanyMaterialsDelete handle requests on route /Company/categories/delete
func (ach *CompanyMaterialHandler) CompanyMaterialsDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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

	http.Redirect(w, r, "/company", http.StatusSeeOther)
}

//MaterialSearch handle requests on route /Company/categories/delete
func (ach *CompanyMaterialHandler) MaterialSearch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctg := entity.Material{}
	a := ctg.Name

	if r.Method == http.MethodPost {
		a = r.FormValue("search")
		material, err := ach.MaterialSrv.MaterialSearch(a)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		ach.tmpl.ExecuteTemplate(w, "user.search.layout", material)

	} else {

		ach.tmpl.ExecuteTemplate(w, "user", nil)

	}
}

//IndexMaterialSearch handle requests on route /Company/categories/delete
func (ach *CompanyMaterialHandler) IndexMaterialSearch(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	ctg := entity.Material{}
	a := ctg.Name

	if r.Method == http.MethodPost {
		a = r.FormValue("search")
		material, err := ach.MaterialSrv.MaterialSearch(a)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		ach.tmpl.ExecuteTemplate(w, "index.layout", material)

	} else {
		ach.tmpl.ExecuteTemplate(w, "index.layout", nil)

	}
}

//GetRentedMaterials -
func (ach *CompanyMaterialHandler) GetRentedMaterials(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := GetUserFromJWT(r)
	if err != nil {
		fmt.Println(err)
	}
	usr, err := ach.comp.GetRentedMaterials(int(user))
	fmt.Println(usr)
	if err != nil {
		fmt.Println(err)
	}
	ach.tmpl.ExecuteTemplate(w, "user.rented.info.layout", usr)
}

//GetUserFromJWT - GETS USER FROM JWTT
func GetUserFromJWT(r *http.Request) (float64, error) {
	coo, err := r.Cookie("auth-information")

	token, err := jwt.Parse(coo.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret-key"), nil
	})
	if err != nil {
		return 0, err
	}
	user := token.Claims.(jwt.MapClaims)["auth-information"].(map[string]interface{})["ID"].(float64)
	return user, nil
}
