package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	compRepo "github.com/ermiasgashu/Construction-Machinary-Rental/company/repository"
	compService "github.com/ermiasgashu/Construction-Machinary-Rental/company/service"
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

// func login(w http.ResponseWriter, r *http.Request) {
// 	templ.ExecuteTemplate(w, "signup.layout", nil)
// }
func admin(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "admin.layout", nil)
}

func users(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "user.layout", nil)
}

func loginAs(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "loginAsCompany.layout", nil)
}

func companies(w http.ResponseWriter, r *http.Request) {
	templ.ExecuteTemplate(w, "company.layout", nil)

}

// func company(w http.ResponseWriter, r *http.Request) {
// 	templ.ExecuteTemplate(w, "company.layout", nil)
// }
func main() {

	// session, err := store.Get(r, "session-name")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	dbconn, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	if err := dbconn.Ping(); err != nil {
		panic(err)
	}

	userRepo := usrRepo.NewPsqlUserRepository(dbconn)
	userService := usrService.NewUserServiceImpl(userRepo)
	userHandler := handler.NewUserHandler(userService, templ)

	companyRepo := compRepo.NewCompanyRepo(dbconn)
	compService := compService.NewCompanyService(companyRepo)
	companyHandler := handler.NewCompanyHandler(compService, templ)

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("../ui/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", index)

	mux.HandleFunc("/users", userHandler.UserSignup) //TODO needs authentication to access this route
	mux.HandleFunc("/users/signup", userHandler.AddUser)

	mux.HandleFunc("/admin", admin)

	mux.HandleFunc("/companies/signup", companyHandler.CompanySigup)
	mux.HandleFunc("/companies", companyHandler.CompanyIndex) //TODO needs authentication to access this route

	http.ListenAndServe(":8080", mux)
}
