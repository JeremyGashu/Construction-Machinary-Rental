package api

import (
	"encoding/json"

	"fmt"
	"net/http"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
	"github.com/julienschmidt/httprouter"
)

// AdminUserHandler handles User related http requests
type AdminUserHandler struct {
	UserService admin.UserService
}

// NewAdminUserHandler returns new AdminUserHandler object
func NewAdminUserHandler(cmntService admin.UserService) *AdminUserHandler {
	return &AdminUserHandler{UserService: cmntService}
}

// GetUsers handles GET /v1/admin/Users request
func (ach *AdminUserHandler) GetUsers(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	Users, errs := ach.UserService.Users()

	if errs != nil {
		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(Users, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return

}

// GetSingleUser handles GET /v1/admin/Users/:id request
func (ach *AdminUserHandler) GetSingleUser(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id := ps.ByName("username")

	User, errs := ach.UserService.User(id)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(User, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// PostUser handles POST /v1/admin/Users request
func (ach *AdminUserHandler) PostUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	User := &entity.User{}

	err := json.Unmarshal(body, User)

	if err != nil {

		fmt.Println(err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	errs := ach.UserService.StoreUser(*User)

	if errs != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/v1/admin/Users/%s", User.Username)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

// PutUser handles PUT /v1/admin/Users/:id request
func (ach *AdminUserHandler) PutUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id := ps.ByName("username")

	User, errs := ach.UserService.User(id)

	if errs != nil {

		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &User)

	errs = ach.UserService.UpdateUser(User)

	if errs != nil {

		fmt.Println(errs)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(User, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}

// DeleteUser handles DELETE /v1/admin/Users/:id request
func (ach *AdminUserHandler) DeleteUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id := ps.ByName("username")

	errs := ach.UserService.DeleteUser(id)

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
