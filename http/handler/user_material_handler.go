package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/company"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
	"github.com/julienschmidt/httprouter"
)

//UserMaterialHandler - User material handler
type UserMaterialHandler struct {
	comp            admin.CompanyService
	user            admin.UserService
	materialService company.MaterialService
	tmpl            *template.Template
}

//NewUserMaterialHandler -
func NewUserMaterialHandler(ms company.MaterialService, t *template.Template, usr admin.UserService, c admin.CompanyService) *UserMaterialHandler {
	return &UserMaterialHandler{materialService: ms, tmpl: t, user: usr, comp: c}
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
	// fmt.Println(materials)
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

	LogedUser string
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

	//ADDED FOR TEST

	// info.LogedUser = user
	//TEST
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
		user, err := GetLogedUserFromJWT(r)
		if err != nil {
			fmt.Println(err)
		}
		var material entity.Material

		err = json.Unmarshal(body, &material)
		own, _ := mh.materialService.GetOwner(material.Owner) //getting owner from the owner id

		// fmt.Println(own.CompanyID)
		info := Information{LogedUser: user, CompanyName: own.Name, CompanyAddress: own.Address, CompanyDescription: own.Description, CompanyEmail: own.Email, CompanyImagePath: own.ImagePath, CompanyPhone: own.PhoneNo, ComppanyRating: own.Rating, MaterialID: material.ID, MaterialImagePath: material.ImagePath, MaterialName: material.Name, OnDiscount: material.OnDiscount, OnSale: material.OnSale, PricePerDay: material.PricePerDay, CompanyID: own.CompanyID}
		// fmt.Println(info)
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

		info.Username = r.FormValue("logedUser")
		info.CompanyID, _ = strconv.Atoi(r.FormValue("companyID"))
		info.DueDate = r.FormValue("returnDate")
		valid := VaidDate(fmt.Sprintf("%d-%d-%d", year, month, day), info.DueDate)

		parsedDate := strings.Split(r.FormValue("returnDate"), "-")
		info.MaterialID, _ = strconv.Atoi(r.FormValue("materialID"))
		if !valid {
			http.Redirect(w, r, "/user/rent/"+r.FormValue("materialID"), http.StatusSeeOther)
			return
		}
		price, _ := strconv.ParseFloat(r.FormValue("priceperday"), 10)
		dday, err := strconv.Atoi(parsedDate[2])
		comp, err := mh.materialService.GetOwner(info.CompanyID)

		if err != nil {
			fmt.Println(err)
		}

		if err != nil {
			fmt.Println(err)
		}
		dmonth, err := strconv.Atoi(parsedDate[1])
		if err != nil {
			fmt.Println(err)
		}
		dyear, err := strconv.Atoi(parsedDate[0])
		if err != nil {
			fmt.Println(err)
		}
		amount := float32((dyear-year)*365+(dmonth-month)*30+(dday-day)) * float32(price)

		comp.Account = comp.Account + amount
		// fmt.Println(comp)

		// fmt.Println("Loogged:" + r.FormValue("logedUser"))

		if amount > 0 {

			b := mh.user.Pay(info.Username, float64(amount), comp.CompanyID)

			if b {
				info.TransactionMade = float64(amount)
			}
		} else {
			info.TransactionMade = float64(0)

		}

		err = mh.comp.UpdateCompany(comp)
		if err != nil {
			fmt.Println(err)
		}

		err = mh.materialService.RentMaterial(info)

		// fmt.Println(info)
		if err != nil {
			fmt.Println(err)
		}
		http.Redirect(w, r, "/user", http.StatusSeeOther)
	}
}

//GetRentedMaterials -
func (mh *UserMaterialHandler) GetRentedMaterials(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := GetLogedUserFromJWT(r)
	if err != nil {
		fmt.Println(err)
	}
	usr, err := mh.user.GetRentedMaterials(user)
	if err != nil {
		fmt.Println(err)
	}
	mh.tmpl.ExecuteTemplate(w, "user.rented.info.layout", usr)
}

//GetLogedUserFromJWT - get user
func GetLogedUserFromJWT(r *http.Request) (string, error) {
	coo, err := r.Cookie("auth-information")

	token, err := jwt.Parse(coo.Value, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("There was an error")
		}
		return []byte("secret-key"), nil
	})
	if err != nil {
		return "", err
	}
	user := token.Claims.(jwt.MapClaims)["auth-information"].(map[string]interface{})["Value"].(string)
	return user, nil
}
func VaidDate(now string, due string) bool {
	NOW := strings.Split(now, "-")
	nowDay, _ := strconv.Atoi(NOW[2])
	nowMonth, _ := strconv.Atoi(NOW[1])
	nowYear, _ := strconv.Atoi(NOW[0])
	DUE := strings.Split(due, "-")
	dueD, _ := strconv.Atoi(DUE[2])
	dueM, _ := strconv.Atoi(DUE[1])
	dueY, _ := strconv.Atoi(DUE[0])
	if dueD < nowDay {
		if nowMonth < dueM && nowYear <= dueY {
			return true
		}
		return false
	} else if dueD > nowDay {
		if nowMonth <= dueM && nowYear <= dueY {
			return true
		}
	}
	return false
	// if(nowYear <= dueY && nowMonth <= dueM && nowDay < )
}
