package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"strings"

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

//ParseDate -
func ParseDate(date string) string {
	dat := strings.Split(date, "T")
	return dat[0]
}
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
	// b := UserServ.Pay("eyuelyemane", 2000)
	// fmt.Println(b)
	adminUsersHandler := handlers.NewAdminUserHandler(templ, UserServ)

	authHandler := handlers.NewCompanyAuthHandler(CompanyServ)

	allAuthHandler := handlers.NewAuthHander(CompanyServ, UserServ, AdminServ)

	apiAdminUsersHandler := api.NewAdminUserHandler(UserServ)

	materialRepo := comprep.NewMaterialRepository(dbconn)
	ser := compser.NewMaterialService(materialRepo)

	hand := api.NewCompanyMaterialHandler(ser)
	handlol := handlers.NewCompanyMaterialHandler(templ, ser, CompanyServ)
	userMaterialHandler := handlers.NewUserMaterialHandler(ser, templ, UserServ, CompanyServ)
	userProfileHandler := handlers.NewUserProfileHandler(ser, templ, UserServ)

	// materialHandle := handlers.NewCompanyMaterialHandler(templ, ser)
	// serv := api.NewCompanyMaterialHandler(materialSer)
	// ap := api.NewCompanyUseCaseHander(*CompanyServ)
	// CommentRepo := repository.NewCommentRepositoryImpl(dbconn)
	// CommentServ := service.NewCommentServiceImpl(CommentRepo)
	// adminCommentsHandler := handlers.NewAdminCommentHandler(templ, CommentServ)

	userSignupHandler := handlers.NewUserSignupHandler(UserServ, templ)
	cpnySignupHandler := handlers.NewCompanySignUpHandler(CompanyServ, templ)
	cpnyProfileHandler := handlers.NewCompanyProfileHandler(ser, templ, CompanyServ)

	//THIS WILL BE CLASSIFIED AS CLIENT AND SERVER FOR LATER US
	// fs := http.FileServer(http.Dir("../ui/assets"))
	router.ServeFiles("/assets/*filepath", http.Dir("../ui/assets"))
	// http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	//router.GET("/", index)
	router.POST("/login", allAuthHandler.Login)
	router.GET("/logout", allAuthHandler.Logout)

	router.GET("/", handlol.IndexMaterialSearch)
	// router.GET("/", handlol.IndexMaterialSearch)
	router.POST("/", handlol.IndexMaterialSearch)

	router.GET("/v1/search/material/:name", hand.SearchMaterial)

	//Signing in as a company is a must to access this page
	router.GET("/user", middleware.UserLoginRequired(userMaterialHandler.UserIndex)) //Loggin ing is must to access this page
	router.GET("/user/materials/:id", middleware.UserLoginRequired(userMaterialHandler.Material))
	router.GET("/user/rent/:material_id", middleware.UserLoginRequired(userMaterialHandler.UserRentMaterial))
	router.POST("/user/rent", middleware.UserLoginRequired(userMaterialHandler.UserRentMaterial))
	router.GET("/user/search", middleware.UserLoginRequired(handlol.MaterialSearch))
	router.POST("/user/search", middleware.UserLoginRequired(handlol.MaterialSearch))
	router.GET("/user/profile", middleware.UserLoginRequired(userProfileHandler.ProfileIndex))
	router.GET("/user/material/rented", middleware.UserLoginRequired(userMaterialHandler.GetRentedMaterials))
	router.POST("/user/profile", middleware.UserLoginRequired(userProfileHandler.UpdateProfile))
	router.GET("/user/ondiscount", middleware.UserLoginRequired(userMaterialHandler.MaterialsOnDiscount))
	router.GET("/user/register", login)
	router.POST("/user/register", userSignupHandler.SignupHandler)

	router.GET("/company", middleware.CompaniesLoginReequired(handlol.CompanyIndex))
	router.GET("/company/materials", middleware.CompaniesLoginReequired(handlol.CompanyMaterials)) // company login is essential USE MIDDLE WARE
	router.GET("/company/material/new", middleware.CompaniesLoginReequired(handlol.CompanyMaterialsNew))
	router.GET("/company/material/update", middleware.CompaniesLoginReequired(handlol.CompanyMaterialsUpdate))
	router.GET("/company/material/delete", middleware.CompaniesLoginReequired(handlol.CompanyMaterialsDelete))
	router.POST("/company/material/update", middleware.CompaniesLoginReequired(handlol.CompanyMaterialsUpdate))
	router.GET("/company/rented/delete/cid/:company_id/mid/:material_id/uid/:user_id", middleware.CompaniesLoginReequired(handlol.DeleteRentedMaterial))
	router.POST("/company/material/new", middleware.CompaniesLoginReequired(handlol.CompanyMaterialsNew))
	router.GET("/company/materials/rented", middleware.CompaniesLoginReequired(handlol.GetRentedMaterials))
	router.GET("/company/profile", middleware.CompaniesLoginReequired(cpnyProfileHandler.ProfileIndex))
	router.POST("/company/profile", middleware.CompaniesLoginReequired(cpnyProfileHandler.UpdateProfile))

	router.POST("/companies/register", cpnySignupHandler.SignupHandler)
	router.GET("/company/register", loginAs)

	router.GET("/admin", middleware.AdminLoginRequired(admin))
	router.GET("/admin/admins", adminAdminsHandler.AdminAdmins)
	router.POST("/admin/admins/new", middleware.AdminLoginRequired(adminAdminsHandler.AdminAdminsNew))
	router.GET("/admin/admins/new", middleware.AdminLoginRequired(adminAdminsHandler.AdminAdminsNew))
	router.GET("/admin/admins/delete", middleware.AdminLoginRequired(adminAdminsHandler.AdminAdminsDelete))
	//handle company
	router.GET("/admin/company", middleware.AdminLoginRequired(adminCompanysHandler.AdminCompanys))
	router.POST("/admin/company/new", middleware.AdminLoginRequired(adminCompanysHandler.AdminCompanysNew))
	router.GET("/admin/company/new", middleware.AdminLoginRequired(adminCompanysHandler.AdminCompanysNew))
	router.GET("/admin/requests", middleware.AdminLoginRequired(adminCompanysHandler.Unactivated))
	router.GET("/admin/company/approve", middleware.AdminLoginRequired(adminCompanysHandler.Approve))

	router.GET("/admin/company/delete", middleware.AdminLoginRequired(adminCompanysHandler.AdminCompanysDelete))
	//handle user
	router.GET("/admin/user", middleware.AdminLoginRequired(adminUsersHandler.AdminUsers))
	router.POST("/admin/user/new", middleware.AdminLoginRequired(adminUsersHandler.AdminUsersNew))
	router.GET("/admin/user/new", middleware.AdminLoginRequired(adminUsersHandler.AdminUsersNew))
	router.GET("/admin/users/delete", middleware.AdminLoginRequired(adminUsersHandler.AdminUsersDelete))
	//handle user
	// router.GET("/admin/comment", adminCommentsHandler.AdminComments)
	// router.GET("/admin/comment/new", adminCommentsHandler.AdminCommentsNew)
	// router.POST("/admin/comment/new", adminCommentsHandler.AdminCommentsNew)
	// router.GET("/admin/comment/delete", adminCommentsHandler.AdminCommentsDelete)

	// http.HandleFunc("/company/material", materialHandle.CompanyMaterials)
	// http.HandleFunc("/company/material/new", materialHandle.CompanyMaterialsNew)
	// http.HandleFunc("/company/material/update", materialHandle.CompanyMaterialsUpdate)
	// http.HandleFunc("/company/material/delete", materialHandle.CompanyMaterialsDelete)

	router.GET("/v1/companies/materials", hand.Materials)                          //to show all materials for the user
	router.GET("/v1/companies/owner/:company_id/materials", hand.MaterialsByOwner) //lists all materials of a single company
	router.GET("/v1/companies/materials/:material_id", hand.Material)
	router.PUT("/v1/companies/materials/:id", hand.UpdateMaterial)
	router.DELETE("/v1/companies/materials/delete/:material_id", hand.DeleteMaterial)
	router.POST("/v1/companies/materials", hand.StoreMaterial)
	router.GET("/v1/materials/discount", hand.MaterialsOnDiscount)

	router.POST("/v1/companies/login", authHandler.Login)

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
	router.GET("/v1/admin/admins/:username", apiAdminAdminsHandler.GetSingleAdmin) //TESTED
	router.GET("/v1/admin/admins", apiAdminAdminsHandler.GetAdmins)                //TESTED
	router.PUT("/v1/admin/admins/:username", apiAdminAdminsHandler.PutAdmin)       //TESTED
	router.POST("/v1/admin/admins", apiAdminAdminsHandler.PostAdmin)               //TESTED
	router.DELETE("/v1/admin/admins/:username", apiAdminAdminsHandler.DeleteAdmin) //TESTED

	http.ListenAndServe(":8080", router)
}
