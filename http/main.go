package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	handler "github.com/ermiasgashu/Construction-Machinary-Rental/http/handlers"
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
// func admin(w http.ResponseWriter, r *http.Request) {
// 	templ.ExecuteTemplate(w, "admin.layout", nil)
// }
// func userr(w http.ResponseWriter, r *http.Request) {
// 	templ.ExecuteTemplate(w, "user.layout", nil)
// }

// func loginAs(w http.ResponseWriter, r *http.Request) {
// 	templ.ExecuteTemplate(w, "loginAsCompany.layout", nil)
// }

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

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("../ui/assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	mux.HandleFunc("/", index)
	mux.HandleFunc("/users/signup", userHandler.AddUser)

	// mux.HandleFunc("/admin", admin)
	// mux.HandleFunc("/user", userr)
	// http.HandleFunc("/users/signup", index)
	http.ListenAndServe(":8080", mux)
}
