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
	"github.com/ermiasgashu/Construction-Machinary-Rental/middleware"
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
	apiAdminAdminsHandler := api.NewAdminAdminsHandler(AdminServ)

	CompanyRepo := repository.NewCompanyRepositoryImpl(dbconn)
	CompanyServ := service.NewCompanyServiceImpl(CompanyRepo)
	adminCompanysHandler := handlers.NewAdminCompanyHandler(templ, CompanyServ)
	apiAdminCompanysHandler := api.NewAdminCompanyHandler(CompanyServ)

	//company
	router := httprouter.New()
	//User
	UserRepo := repository.NewUserRepositoryImpl(dbconn)
	UserServ := service.NewUserServiceImpl(UserRepo)
	adminUsersHandler := handlers.NewAdminUserHandler(templ, UserServ)

	authHandler := handlers.NewCompanyAuthHandler(CompanyServ)

	apiAdminUsersHandler := api.NewAdminUserHandler(UserServ)

	materialRepo := comprep.NewMaterialRepository(dbconn)
	ser := compser.NewMaterialService(materialRepo)
	hand := api.NewCompanyMaterialHandler(ser)

	// materialHandle := handlers.NewCompanyMaterialHandler(templ, ser)
	// serv := api.NewCompanyMaterialHandler(materialSer)
	// ap := api.NewCompanyUseCaseHander(*CompanyServ)
	CommentRepo := repository.NewCommentRepositoryImpl(dbconn)
	CommentServ := service.NewCommentServiceImpl(CommentRepo)
	adminCommentsHandler := handlers.NewAdminCommentHandler(templ, CommentServ)

	userSignupHandler := handlers.NewUserSignupHandler(UserServ, templ)
	cpnySignupHandler := handlers.NewCompanySignUpHandler(CompanyServ, templ)

	//THIS WILL BE CLASSIFIED AS CLIENT AND SERVER FOR LATER US
	// fs := http.FileServer(http.Dir("../ui/assets"))
	router.ServeFiles("/assets/*filepath", http.Dir("../ui/assets"))
	// http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	router.GET("/", index)
	router.GET("/login", login)
	router.GET("/signinCompany", loginAs)
	router.GET("/admin", admin)     //Signing in as a company is a must to access this page
	router.GET("/user", userr)      //Loggin ing is must to access this page
	router.GET("/company", company) // company login is essential USE MIDDLE WARE

	router.POST("/user/register", userSignupHandler.SignupHandler)
	router.POST("/companies/register", cpnySignupHandler.SignupHandler)

	router.GET("/admin/admins", adminAdminsHandler.AdminAdmins)
	router.POST("/admin/admins/new", adminAdminsHandler.AdminAdminsNew)
	router.GET("/admin/admins/new", adminAdminsHandler.AdminAdminsNew)
	router.GET("/admin/admins/delete", adminAdminsHandler.AdminAdminsDelete)
	//handle company
	router.GET("/admin/company", adminCompanysHandler.AdminCompanys)
	router.POST("/admin/company/new", adminCompanysHandler.AdminCompanysNew)
	router.GET("/admin/company/new", adminCompanysHandler.AdminCompanysNew)
	router.GET("/admin/requests", adminCompanysHandler.Unactivated)
	router.GET("/admin/company/approve", adminCompanysHandler.Approve)

	router.GET("/admin/company/delete", adminCompanysHandler.AdminCompanysDelete)
	//handle user
	router.GET("/admin/user", adminUsersHandler.AdminUsers)
	router.POST("/admin/user/new", adminUsersHandler.AdminUsersNew)
	router.GET("/admin/user/new", adminUsersHandler.AdminUsersNew)

	router.GET("/admin/users/delete", adminUsersHandler.AdminUsersDelete)
	//handle user
	router.GET("/admin/comment", adminCommentsHandler.AdminComments)
	router.GET("/admin/comment/new", adminCommentsHandler.AdminCommentsNew)
	router.POST("/admin/comment/new", adminCommentsHandler.AdminCommentsNew)
	router.GET("/admin/comment/delete", adminCommentsHandler.AdminCommentsDelete)

	// http.HandleFunc("/company/material", materialHandle.CompanyMaterials)
	// http.HandleFunc("/company/material/new", materialHandle.CompanyMaterialsNew)
	// http.HandleFunc("/company/material/update", materialHandle.CompanyMaterialsUpdate)
	// http.HandleFunc("/company/material/delete", materialHandle.CompanyMaterialsDelete)

	router.GET("/v1/companies/materials", hand.Materials)
	router.GET("/v1/companies/materials/:material_id", middleware.CompanyLoginRequired(hand.Material))
	router.PUT("/v1/companies/materials/:id", middleware.CompanyLoginRequired(hand.UpdateMaterial))
	router.DELETE("/v1/companies/materials/delete/:material_id", middleware.CompanyLoginRequired(hand.DeleteMaterial))
	router.POST("/v1/companies/materials", middleware.CompanyLoginRequired(hand.StoreMaterial))
	router.POST("/v1/companies/login", authHandler.Login)
	// router.GET("/v1/companies/test", middleware.CompanyLoginRequired(authHandler.TestJWT))

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

	http.ListenAndServe(":8080", router)
}
