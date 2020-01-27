package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
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
func (mh *UserMaterialHandler) UserIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

//Information - template info provider
type Information struct {
	CompanyName        string
	CompanyEmail       string
	CompanyAddress     string
	CompanyPhone       string
	CompanyDescription string
	ComppanyRating     int
	CompanyImagePath   string
	CompanyID          int

	MaterialID        int
	MaterialName      string
	PricePerDay       float64
	OnDiscount        bool
	OnSale            bool
	MaterialImagePath string
}

//Material -
func (mh *UserMaterialHandler) Material(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	client := http.DefaultClient
	rawID := ps.ByName("id")
	id, err := strconv.Atoi(rawID)
	if err != nil {
		fmt.Println(err)
	}
	url := fmt.Sprintf("http://localhost:8080/v1/companies/materials/%d", id)
	req, err := http.NewRequest("GET", url, nil)
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
	var material entity.Material

	err = json.Unmarshal(body, &material)

	own, _ := mh.materialService.GetOwner(material.Owner) //getting owner from the owner id

	info := Information{CompanyName: own.Name, CompanyAddress: own.Address, CompanyDescription: own.Description, CompanyEmail: own.Email, CompanyImagePath: own.ImagePath, CompanyPhone: own.PhoneNo, ComppanyRating: own.Rating, MaterialID: material.ID, MaterialImagePath: material.ImagePath, MaterialName: material.Name, OnDiscount: material.OnDiscount, OnSale: material.OnSale, PricePerDay: material.PricePerDay}

	if err != nil {
		fmt.Println(err)
	}
	mh.tmpl.ExecuteTemplate(w, "user.single.material.layout", info)
}

//UserRentMaterial -
func (mh *UserMaterialHandler) UserRentMaterial(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Method == http.MethodGet {
		client := http.DefaultClient
		rawID := ps.ByName("material_id")
		id, err := strconv.Atoi(rawID)
		if err != nil {
			fmt.Println(err)
		}
		url := fmt.Sprintf("http://localhost:8080/v1/companies/materials/%d", id)
		req, err := http.NewRequest("GET", url, nil)
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
		var material entity.Material

		err = json.Unmarshal(body, &material)
		own, _ := mh.materialService.GetOwner(material.Owner) //getting owner from the owner id
		fmt.Println(own.CompanyID)
		info := Information{CompanyName: own.Name, CompanyAddress: own.Address, CompanyDescription: own.Description, CompanyEmail: own.Email, CompanyImagePath: own.ImagePath, CompanyPhone: own.PhoneNo, ComppanyRating: own.Rating, MaterialID: material.ID, MaterialImagePath: material.ImagePath, MaterialName: material.Name, OnDiscount: material.OnDiscount, OnSale: material.OnSale, PricePerDay: material.PricePerDay, CompanyID: own.CompanyID}

		if err != nil {
			fmt.Println(err)
		}
		mh.tmpl.ExecuteTemplate(w, "user.rent.layout", info)
	} else if r.Method == http.MethodPost {
		var info entity.RentInformation

		rentdate := time.Now()
		day := rentdate.Day()
		month := int(rentdate.Month())
		year := rentdate.Year()

		info.RentDate = fmt.Sprintf("%d-%d-%d", year, month, day)
		info.CompanyID, _ = strconv.Atoi(r.FormValue("companyID"))
		info.DueDate = r.FormValue("returnDate")
		info.MaterialID, _ = strconv.Atoi(r.FormValue("materialID"))

		coo, err := r.Cookie("auth-information")
		if err != nil {
			http.Redirect(w, r, "/user", http.StatusSeeOther)

		} else {
			if coo != nil {
				token, err := jwt.Parse(coo.Value, func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("secret-key"), nil
				})
				if err != nil {
					fmt.Println(err)
				}
				user := token.Claims.(jwt.MapClaims)["auth-information"].(map[string]interface{})["Value"].(string)
				info.Username = user
				// // fmt.Println("Valid Token Thanks")
				// fmt.Println(token.Claims.(jwt.MapClaims)["auth-information"].(map[string]interface{})["Value"]) // GETTING THE CLAIM AFTER LONG TRY
			}
		}
		info.TransactionMade = 1200
		err = mh.materialService.RentMaterial(info)
		// fmt.Println(info)
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/user", http.StatusSeeOther)
	}
}
