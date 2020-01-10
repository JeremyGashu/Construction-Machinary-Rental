package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/ermiasgashu/Construction-Machinary-Rental/middleware"

	// compRepo "github.com/ermiasgashu/Construction-Machinary-Rental/company/repository"
	// compService "github.com/ermiasgashu/Construction-Machinary-Rental/company/service"
	handler "github.com/ermiasgashu/Construction-Machinary-Rental/http/handler"
	usrRepo "github.com/ermiasgashu/Construction-Machinary-Rental/user/repository"
	usrService "github.com/ermiasgashu/Construction-Machinary-Rental/user/service"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "1234"
	dbname   = "constructiondb"
)

var templ = template.Must(template.ParseGlob("../ui/templates/*"))

func index(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "index.layout", nil)
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

	userRepo := usrRepo.NewPsqlUserRepository(dbconn)
	userService := usrService.NewUserServiceImpl(userRepo)
	userHandler := handler.NewUserHandler(userService, templ)

	// companyRepo := compRepo.NewCompanyRepo(dbconn)
	// compService := compService.NewCompanyService(companyRepo)
	// companyHandler := handler.NewCompanyHandler(compService, templ)

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("../ui/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", index)

	mux.HandleFunc("/users", middleware.UserLoginRequired(userHandler.UserSignup)) //TODO needs authentication to access this route
	mux.HandleFunc("/users/login", userHandler.UserLogin)
	mux.HandleFunc("/users/signup", userHandler.AddUser)
	mux.HandleFunc("/users/logout", userHandler.LogOut)

	// mux.HandleFunc("/admin", admin) //TODO admin handlers

	// mux.HandleFunc("/companies/signup", companyHandler.CompanySigup)
	// //using local host and the default saved username and password
	// mux.HandleFunc("/companies", companyHandler.CompanyIndex) //TODO needs authentication to access this route

	http.ListenAndServe(":8080", mux)
}
