package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
	"github.com/julienschmidt/httprouter"
)

// AdminAdminHandler handles Admin related http requests
type AdminAdminHandler struct {
	AdminService admin.AdminService
}

// NewAdminAdminsHandler returns new AdminAdminHandler object
func NewAdminAdminsHandler(cmntService admin.AdminService) *AdminAdminHandler {
	return &AdminAdminHandler{AdminService: cmntService}
}

// GetAdmins handles GET /v1/admin/Admins request
func (ach *AdminAdminHandler) GetAdmins(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	Admins, errs := ach.AdminService.Admins()

	if errs != nil {
		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(Admins, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return

}

// GetSingleAdmin handles GET /v1/admin/Admins/:id request
func (ach *AdminAdminHandler) GetSingleAdmin(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id := ps.ByName("username")

	Admin, errs := ach.AdminService.Admin(id)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(Admin, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// PostAdmin handles POST /v1/admin/Admins request
func (ach *AdminAdminHandler) PostAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	Admin := &entity.Admin{}

	err := json.Unmarshal(body, Admin)

	if err != nil {

		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	errs := ach.AdminService.StoreAdmin(*Admin)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/admin/Admins/%s", Admin.Username)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// PutAdmin handles PUT /v1/admin/Admins/:id request
func (ach *AdminAdminHandler) PutAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id := ps.ByName("username")

	Admin, errs := ach.AdminService.Admin(id)

	if errs != nil {

		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &Admin)

	errs = ach.AdminService.UpdateAdmin(Admin)

	if errs != nil {

		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(Admin, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// DeleteAdmin handles DELETE /v1/admin/Admins/:id request
func (ach *AdminAdminHandler) DeleteAdmin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id := ps.ByName("username")

	errs := ach.AdminService.DeleteAdmin(id)

	if errs != nil {

		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
