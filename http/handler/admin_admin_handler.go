package handlers

import (
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

// AdminAdminHandler handles Admin handler admin requests
type AdminAdminHandler struct {
	tmpl     *template.Template
	AdminSrv admin.AdminService
}

// NewAdminAdminHandler initializes and returns new AdminCateogryHandler
func NewAdminAdminHandler(T *template.Template, CS admin.AdminService) *AdminAdminHandler {
	return &AdminAdminHandler{tmpl: T, AdminSrv: CS}
}

// AdminAdmins handle requests on route /admin/categories
func (ach *AdminAdminHandler) AdminAdmins(w http.ResponseWriter, r *http.Request) {
	admins, err := ach.AdminSrv.Admins()
	if err != nil {
		panic(err)
	}
	ach.tmpl.ExecuteTemplate(w, "admin.admins.layout", admins)
}

// AdminAdminsNew hanlde requests on route /admin/admins/new
func (ach *AdminAdminHandler) AdminAdminsNew(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		ctg := entity.Admin{}
		ctg.Username = r.FormValue("username")
		ctg.FirstName = r.FormValue("firstname")
		ctg.LastName = r.FormValue("lastname")
		ctg.Email = r.FormValue("email")
		ctg.Password = r.FormValue("password")

		mf, fh, err := r.FormFile("catimg")
		if err != nil {
			panic(err)
		}
		defer mf.Close()

		ctg.ImagePath = fh.Filename

		writeFile(&mf, fh.Filename)

		err = ach.AdminSrv.StoreAdmin(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin/admins", http.StatusSeeOther)

	} else {

		ach.tmpl.ExecuteTemplate(w, "admin.admins.new.layout", nil)

	}
}

// AdminAdminsUpdate handle requests on /admin/categories/update
func (ach *AdminAdminHandler) AdminAdminsUpdate(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("username")

		cat, err := ach.AdminSrv.Admin(idRaw)

		if err != nil {
			panic(err)
		}

		ach.tmpl.ExecuteTemplate(w, "admin.admins.update.layout", cat)

	} else if r.Method == http.MethodPost {

		ctg := entity.Admin{}
		ctg.FirstName = r.FormValue("firstname")
		ctg.LastName = r.FormValue("lastname")
		ctg.Email = r.FormValue("email")
		ctg.Password = r.FormValue("password")
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

		err = ach.AdminSrv.UpdateAdmin(ctg)

		if err != nil {
			panic(err)
		}

		http.Redirect(w, r, "/admin", http.StatusSeeOther)

	} else {
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
	}

}

// AdminAdminsDelete handle requests on route /admin/categories/delete
func (ach *AdminAdminHandler) AdminAdminsDelete(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		idRaw := r.URL.Query().Get("username")

		err := ach.AdminSrv.DeleteAdmin(idRaw)

		if err != nil {
			panic(err)
		}

	}

	http.Redirect(w, r, "/admin/admins", http.StatusSeeOther)
}

func writeFile(mf *multipart.File, fname string) {

	wd, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	path := filepath.Join(wd, "/../ui", "assets", "img", fname)
	image, err := os.Create(path)

	if err != nil {
		panic(err)
	}
	defer image.Close()
	io.Copy(image, *mf)
}
