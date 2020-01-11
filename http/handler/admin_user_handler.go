package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

// AdminUserHandler handles User handler User requests
type AdminUserHandler struct {
	tmpl    *template.Template
	UserSrv admin.UserService
}

// NewAdminUserHandler initializes and returns new UserCateogryHandler
func NewAdminUserHandler(T *template.Template, CS admin.UserService) *AdminUserHandler {
	return &AdminUserHandler{tmpl: T, UserSrv: CS}
}

// AdminUsers handle requests on route /admin/users
func (ach *AdminUserHandler) AdminUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	Users, err := ach.UserSrv.Users()
	if err != nil {
		panic(err)
	}
	ach.tmpl.ExecuteTemplate(w, "admin.user.layout", Users)
}

// AdminUsersNew hanlde requests on route /admin/Users/new
func (ach *AdminUserHandler) AdminUsersNew(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if r.Method == http.MethodPost {
		ctg := entity.User{}
		ctg.Username = r.FormValue("username")
		ctg.FirstName = r.FormValue("firstname")
		ctg.LastName = r.FormValue("lastname")
		ctg.Email = r.FormValue("email")
		ctg.Password = r.FormValue("password")
		ctg.DeliveryAddress = r.FormValue("daddress")
		acc, errs := strconv.Atoi(r.FormValue("account"))
		if errs != nil {
			panic(errs)
		}
		ctg.Account = acc
		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		ctg.ImagePath = fh.Filename

		writeFile(&mf, fh.Filename)

		err = ach.UserSrv.StoreUser(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/user", http.StatusSeeOther)

	} else {

		ach.tmpl.ExecuteTemplate(w, "admin.user.new.layout", nil)

	}
}

// AdminUsersUpdate handle requests on /User/categories/update
func (ach *AdminUserHandler) AdminUsersUpdate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("username")

		cat, err := ach.UserSrv.User(idRaw)

		if err != nil {
			panic(err)
		}

		ach.tmpl.ExecuteTemplate(w, "admin.user.update.layout", cat)

	} else if r.Method == http.MethodPost {

		ctg := entity.User{}
		ctg.FirstName = r.FormValue("firstname")
		ctg.LastName = r.FormValue("lastname")
		ctg.Email = r.FormValue("email")
		ctg.Password = r.FormValue("password")
		ctg.DeliveryAddress = r.FormValue("daddress")
		acc, errs := strconv.Atoi(r.FormValue("account"))
		if errs != nil {
			panic(errs)
		}
		ctg.Account = acc

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

		err = ach.UserSrv.UpdateUser(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/user", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/user", http.StatusSeeOther)
	}

}

// AdminUsersDelete handle requests on route /User/categories/delete
func (ach *AdminUserHandler) AdminUsersDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("username")

		err := ach.UserSrv.DeleteUser(idRaw)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/admin/user", http.StatusSeeOther)
}
