package main

import (
	"database/sql"
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

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	templ.ExecuteTemplate(w, "index.layout", nil)
}
func login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	templ.ExecuteTemplate(w, "login.layout", nil)
}
func loginAs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	templ.ExecuteTemplate(w, "loginAsCompany.layout", nil)
}
func admin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	templ.ExecuteTemplate(w, "admin.layout", nil)
}
func userr(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	templ.ExecuteTemplate(w, "user.layout", nil)
}
func company(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

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
	apiAdminAdminsHandler := api.NewAdminAdminsHandler(AdminServ)

	//company
	CompanyRepo := repository.NewCompanyRepositoryImpl(dbconn)
	CompanyServ := service.NewCompanyServiceImpl(CompanyRepo)
	adminCompanysHandler := handlers.NewAdminCompanyHandler(templ, CompanyServ)
	apiAdminCompanysHandler := api.NewAdminCompanyHandler(CompanyServ)
	router := httprouter.New()
	//User
	UserRepo := repository.NewUserRepositoryImpl(dbconn)
	UserServ := service.NewUserServiceImpl(UserRepo)
	adminUsersHandler := handlers.NewAdminUserHandler(templ, UserServ)
	apiAdminUsersHandler := api.NewAdminUserHandler(UserServ)

	materialRepo := comprep.NewMaterialRepository(dbconn)
	ser := compser.NewMaterialService(materialRepo)
	hand := api.NewCompanyMaterialHandler(ser)

	materialHandle := handlers.NewCompanyMaterialHandler(templ, ser)
	// serv := api.NewCompanyMaterialHandler(materialSer)
	// ap := api.NewCompanyUseCaseHander(*CompanyServ)
	CommentRepo := repository.NewCommentRepositoryImpl(dbconn)
	CommentServ := service.NewCommentServiceImpl(CommentRepo)
	adminCommentsHandler := handlers.NewAdminCommentHandler(templ, CommentServ)

	// fs := http.FileServer(http.Dir("../ui/assets"))
	router.ServeFiles("/assets/*filepath", http.Dir("../ui/assets"))
	// http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	router.GET("/", index)
	router.GET("/login", login)
	router.GET("/signinCompany", loginAs)
	router.GET("/admin", admin)
	router.GET("/user", userr)
	router.GET("/company", company)

	//handle admin
	// router := httprouter.New()

	router.GET("/admin/admins", adminAdminsHandler.AdminAdmins)
	router.POST("/admin/admins/new", adminAdminsHandler.AdminAdminsNew)
	router.POST("/admin/admins/update", adminAdminsHandler.AdminAdminsUpdate)
	router.GET("/admin/admins/update", adminAdminsHandler.AdminAdminsUpdate)
	router.GET("/admin/admins/delete", adminAdminsHandler.AdminAdminsDelete)
	//handle company
	router.GET("/admin/company", adminCompanysHandler.AdminCompanys)
	router.POST("/admin/company/new", adminCompanysHandler.AdminCompanysNew)
	router.POST("/admin/company/update", adminCompanysHandler.AdminCompanysNew)
	router.GET("/admin/company/update", adminCompanysHandler.AdminCompanysUpdate)
	router.GET("/admin/company/delete", adminCompanysHandler.AdminCompanysDelete)
	//handle user
	router.GET("/admin/user", adminUsersHandler.AdminUsers)
	router.POST("/admin/user/new", adminUsersHandler.AdminUsersNew)
	router.GET("/admin/user/new", adminUsersHandler.AdminUsersNew)
	router.PUT("/admin/user/update", adminUsersHandler.AdminUsersUpdate)
	router.GET("/admin/users/delete", adminUsersHandler.AdminUsersDelete)
	//handle user
	http.HandleFunc("/admin/comment", adminCommentsHandler.AdminComments)
	http.HandleFunc("/admin/comment/new", adminCommentsHandler.AdminCommentsNew)
	http.HandleFunc("/admin/comment/update", adminCommentsHandler.AdminCommentsUpdate)
	http.HandleFunc("/admin/comment/delete", adminCommentsHandler.AdminCommentsDelete)

	http.HandleFunc("/company/material", materialHandle.CompanyMaterials)
	http.HandleFunc("/company/material/new", materialHandle.CompanyMaterialsNew)
	http.HandleFunc("/company/material/update", materialHandle.CompanyMaterialsUpdate)
	http.HandleFunc("/company/material/delete", materialHandle.CompanyMaterialsDelete)
	// http.HandleFunc("/v1/companies/login", ap.Login)
	// http.HandleFunc("/v1/companies/secret", middleware.IsAuthorized(ap.Secret))

	router.GET("/v1/companies/materials", hand.Materials)
	router.GET("/v1/companies/materials/:material_id", hand.Material)
	router.DELETE("/v1/companies/materials/delete/:material_id", hand.DeleteMaterial)
	router.POST("/v1/companies/materials/", hand.StoreMaterial)
	//handle company api
	router.GET("/v1/admin/company/:id", apiAdminCompanysHandler.GetSingleCompany)
	router.GET("/v1/admin/company", apiAdminCompanysHandler.GetCompanys)
	router.PUT("/v1/admin/company/:id", apiAdminCompanysHandler.PutCompany)
	router.POST("/v1/admin/company", apiAdminCompanysHandler.PostCompany)
	router.DELETE("/v1/admin/company/:id", apiAdminCompanysHandler.DeleteCompany)
	//handle user api
	router.GET("/v1/admin/user/:username", apiAdminUsersHandler.GetSingleUser)
	router.GET("/v1/admin/user", apiAdminUsersHandler.GetUsers)
	router.PUT("/v1/admin/user/:username", apiAdminUsersHandler.PutUser)
	router.POST("/v1/admin/user", apiAdminUsersHandler.PostUser)
	router.DELETE("/v1/admin/user/:username", apiAdminUsersHandler.DeleteUser)
	//handle Admin api
	router.GET("/v1/admin/admins/:username", apiAdminAdminsHandler.GetSingleAdmin)
	router.GET("/v1/admin/admins", apiAdminAdminsHandler.GetAdmins)
	router.PUT("/v1/admin/admins/:username", apiAdminAdminsHandler.PutAdmin)
	router.POST("/v1/admin/admins", apiAdminAdminsHandler.PostAdmin)
	router.DELETE("/v1/admin/admins/:username", apiAdminAdminsHandler.DeleteAdmin)
	http.ListenAndServe(":8080", nil)
}
