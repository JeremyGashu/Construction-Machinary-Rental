package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"

	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
	"github.com/ermiasgashu/Construction-Machinary-Rental/user"
)

//UserHandler -
type UserHandler struct {
	userService user.Service
	tmpl        *template.Template
	store       *sessions.CookieStore
}

//NewUserHandler -
func NewUserHandler(us user.Service, tmplate *template.Template, str *sessions.CookieStore) *UserHandler {
	return &UserHandler{userService: us, tmpl: tmplate, store: str}
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
		// password2 := r.FormValue("password2") //TODO the real auth is to be made
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

//UserLogin -
func (uh *UserHandler) UserLogin(w http.ResponseWriter, r *http.Request) {
	session, err := uh.store.Get(r, "authentication")
	if err != nil {
		fmt.Println(err)
	}
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")
		usr, err := uh.userService.User(username)
		// fmt.Println("What the user entered: ", password)
		// fmt.Println("What it actually is: ", usr.Password)
		// fmt.Println("What the user entered: ", username)
		// fmt.Println("What it actually is: ", usr.Username)
		// fmt.Println(usr.Username == username)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		} //here we checked if the username exists
		if usr.Password == password {
			session.Values["user_username"] = username
			session.Save(r, w)
			uh.tmpl.ExecuteTemplate(w, "user.layout", nil)

		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
		//TODO do real authentication in the name and in the field

	} else {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

//LogOut -
func (uh *UserHandler) LogOut(w http.ResponseWriter, r *http.Request) {
	session, err := uh.store.Get(r, "authentication")
	if err != nil {
		panic(err)
	}
	session.Options.MaxAge = -1
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
