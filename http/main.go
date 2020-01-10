package main

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin/repository"
	"github.com/ermiasgashu/Construction-Machinary-Rental/admin/service"
	handlers "github.com/ermiasgashu/Construction-Machinary-Rental/http/handler"
	_ "github.com/lib/pq"
)

var templ = template.Must(template.ParseGlob("../ui/templates/*"))

func index(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "index.layout", nil)
}
func login(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "login.layout", nil)
}
func loginAs(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "loginAsCompany.layout", nil)
}
func admin(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "admin.layout", nil)
}
func user(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "user.layout", nil)
}
func company(w http.ResponseWriter, r *http.Request) {

	templ.ExecuteTemplate(w, "company.layout", nil)
}
func main() {
	dbconn, err := sql.Open("postgres", "postgres://postgres:ebsa@localhost/constructiondb?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}
	//admin
	AdminRepo := repository.NewAdminRepositoryImpl(dbconn)
	AdminServ := service.NewAdminServiceImpl(AdminRepo)
	adminAdminsHandler := handlers.NewAdminAdminHandler(templ, AdminServ)
	//company
	CompanyRepo := repository.NewCompanyRepositoryImpl(dbconn)
	CompanyServ := service.NewCompanyServiceImpl(CompanyRepo)
	adminCompanysHandler := handlers.NewAdminCompanyHandler(templ, CompanyServ)
	//User
	UserRepo := repository.NewUserRepositoryImpl(dbconn)
	UserServ := service.NewUserServiceImpl(UserRepo)
	adminUsersHandler := handlers.NewAdminUserHandler(templ, UserServ)

	//Comment
	CommentRepo := repository.NewCommentRepositoryImpl(dbconn)
	CommentServ := service.NewCommentServiceImpl(CommentRepo)
	adminCommentsHandler := handlers.NewAdminCommentHandler(templ, CommentServ)

	fs := http.FileServer(http.Dir("../ui/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/", index)
	http.HandleFunc("/login", login)
	http.HandleFunc("/signinCompany", loginAs)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/user", user)
	http.HandleFunc("/company", company)

	//handle admin
	http.HandleFunc("/admin/admins", adminAdminsHandler.AdminAdmins)
	http.HandleFunc("/admin/admins/new", adminAdminsHandler.AdminAdminsNew)
	http.HandleFunc("/admin/admins/update", adminAdminsHandler.AdminAdminsUpdate)
	http.HandleFunc("/admin/admins/delete", adminAdminsHandler.AdminAdminsDelete)
	//handle company
	http.HandleFunc("/admin/company", adminCompanysHandler.AdminCompanys)
	http.HandleFunc("/admin/company/new", adminCompanysHandler.AdminCompanysNew)
	http.HandleFunc("/admin/company/update", adminCompanysHandler.AdminCompanysUpdate)
	http.HandleFunc("/admin/company/delete", adminCompanysHandler.AdminCompanysDelete)
	//handle user
	http.HandleFunc("/admin/user", adminUsersHandler.AdminUsers)
	http.HandleFunc("/admin/user/new", adminUsersHandler.AdminUsersNew)
	http.HandleFunc("/admin/user/update", adminUsersHandler.AdminUsersUpdate)
	http.HandleFunc("/admin/user/delete", adminUsersHandler.AdminUsersDelete)
	//handle user
	http.HandleFunc("/admin/comment", adminCommentsHandler.AdminComments)
	http.HandleFunc("/admin/comment/new", adminCommentsHandler.AdminCommentsNew)
	http.HandleFunc("/admin/comment/update", adminCommentsHandler.AdminCommentsUpdate)
	http.HandleFunc("/admin/comment/delete", adminCommentsHandler.AdminCommentsDelete)

	http.ListenAndServe(":8080", nil)
}
