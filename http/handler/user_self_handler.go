package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
	"github.com/ermiasgashu/Construction-Machinary-Rental/user"
)

//UserHandler -
type UserHandler struct {
	userService user.Service
	tmpl        *template.Template
}

//NewUserHandler -
func NewUserHandler(us user.Service, tmplate *template.Template) *UserHandler {
	return &UserHandler{userService: us, tmpl: tmplate}
}

//AddUser - User Sign up function
func (uh *UserHandler) AddUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		newUser := entity.User{}
		username := r.FormValue("username")
		firstname := r.FormValue("fname")
		lastname := r.FormValue("lname")
		email := r.FormValue("email")
		phone := r.FormValue("phone")
		address := r.FormValue("address")
		password1 := r.FormValue("pass1")
		// password2 := r.FormValue("password2")
		//TODO check password1 and password2 similarity
		newUser.Username = username
		newUser.FirstName = firstname
		newUser.LastName = lastname
		newUser.Email = email
		newUser.DeliveryAddress = address
		newUser.Phone = phone
		newUser.Password = password1
		//TODO hash password with some algo...
		err := uh.userService.AddUser(newUser)
		if err != nil {
			fmt.Println(err)
			uh.tmpl.ExecuteTemplate(w, "signup.layout", nil)
		}
		//TODO make different page for each

		http.Redirect(w, r, "/", http.StatusSeeOther)
		// uh.tmpl.ExecuteTemplate(w, "index.layout", "TRUE")
	} else {
		fmt.Println(uh.tmpl.ExecuteTemplate(w, "signup.layout", nil))
	}

}

//UserSignup -
func (uh *UserHandler) UserSignup(w http.ResponseWriter, r *http.Request) {
	uh.tmpl.ExecuteTemplate(w, "user.layout", nil)
}
