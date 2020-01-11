package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin/repository"
	"github.com/ermiasgashu/Construction-Machinary-Rental/admin/service"
	comprep "github.com/ermiasgashu/Construction-Machinary-Rental/company/repository"
	compser "github.com/ermiasgashu/Construction-Machinary-Rental/company/service"
	handlers "github.com/ermiasgashu/Construction-Machinary-Rental/http/handler"
	"github.com/ermiasgashu/Construction-Machinary-Rental/http/handler/api"
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
func userr(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "user.layout", nil)
}
func company(w http.ResponseWriter, r *http.Request) {

	templ.ExecuteTemplate(w, "company.layout", nil)
}

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "1234"
	dbname   = "constructiondb"
)

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbconn, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	} //this i
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

	materialRepo := comprep.NewMaterialRepository(dbconn)
	ser := compser.NewMaterialService(materialRepo)
	hand := api.NewCompanyMaterialHandler(ser)
	// serv := api.NewCompanyMaterialHandler(materialSer)
	// ap := api.NewCompanyUseCaseHander(*CompanyServ)
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
	http.HandleFunc("/user", userr)
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

	// http.HandleFunc("/v1/companies/login", ap.Login)
	// http.HandleFunc("/v1/companies/secret", middleware.IsAuthorized(ap.Secret))
	router := httprouter.New()

	router.GET("/v1/companies/materials", hand.Materials)
	router.DELETE("/v1/companies/materials/delete/:material_id", hand.DeleteMaterial)

	http.ListenAndServe(":8080", router)
}
