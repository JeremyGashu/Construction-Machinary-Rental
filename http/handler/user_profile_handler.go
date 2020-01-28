package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/company"
)

//UserProfileHandler - User material handler
type UserProfileHandler struct {
	user            admin.UserService
	materialService company.MaterialService
	tmpl            *template.Template
}

//NewUserProfileHandler -
func NewUserProfileHandler(ms company.MaterialService, t *template.Template, usr admin.UserService) *UserProfileHandler {
	return &UserProfileHandler{materialService: ms, tmpl: t, user: usr}
}

//ProfileIndex -
func (uph *UserProfileHandler) ProfileIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := GetLogedUserFromJWT(r)
	if err != nil {
		fmt.Println(err)
	}
	usr, err := uph.user.User(user)
	if err != nil {
		fmt.Println(err)
	}
	uph.tmpl.ExecuteTemplate(w, "user.edit.layout", usr)
}

//UpdateProfile -
func (uph *UserProfileHandler) UpdateProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := GetLogedUserFromJWT(r)
	if err != nil {
		fmt.Println(err)
	}
	usr, err := uph.user.User(user)
	someUser := usr
	if err != nil {
		fmt.Println(err)
	}
	fname := r.FormValue("firstname")
	lname := r.FormValue("lastname")
	email := r.FormValue("email")
	phone := r.FormValue("phone")
	delAd := r.FormValue("delivery")
	mf, fh, err := r.FormFile("catimg")
	if mf != nil {
		usr.ImagePath = fh.Filename

		if err != nil {
			panic(err)
		}

		defer mf.Close()

		writeFile(&mf, usr.ImagePath)

		fmt.Println(usr.ImagePath)
	} else {
		usr.ImagePath = r.FormValue("catimg")
	}
	usr.FirstName = fname
	usr.LastName = lname
	usr.Email = email
	usr.Phone = phone
	usr.DeliveryAddress = delAd
	err = uph.user.UpdateUser(usr)
	if err != nil {
		uph.tmpl.ExecuteTemplate(w, "user.edit.layout", someUser)
	}
	http.Redirect(w, r, "/user", http.StatusSeeOther)
}
