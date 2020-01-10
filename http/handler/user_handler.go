package handlers

// import (
// 	"html/template"
// 	"net/http"

// 	"github.com/project/WebProject/register"
// )

// //AdminUserHandler ..
// type AdminUserHandler struct {
// 	tmpl     *template.Template
// 	userserv register.AuthenticationService
// }

// //NewAdminUserHandler ..
// func NewAdminUserHandler(t *template.Template, us register.AuthenticationService) *AdminUserHandler {
// 	return &AdminUserHandler{tmpl: t, userserv: us}
// }

// //AdminUsers ..
// func (auh *AdminUserHandler) AdminUsers(w http.ResponseWriter, r http.Request) {
// 	usrs, err := auh.userserv.Users()
// 	if err != nil {
// 		panic(err)
// 	}
// 	auh.tmpl.ExecuteTemplate(w, "admin", usrs)
// }
