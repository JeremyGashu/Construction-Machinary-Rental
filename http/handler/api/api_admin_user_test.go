package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"

	"github.com/ermiasgashu/Construction-Machinary-Rental/admin/repository"
	"github.com/ermiasgashu/Construction-Machinary-Rental/admin/service"
	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

func TestAdminUserDelete(t *testing.T) {
	// tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))

	userRepo := repository.NewMockUserRepository(nil)
	userService := service.NewUserServiceImpl(userRepo)

	// handler.NewA
	userAPIHandler := NewAdminUserHandler(userService)

	router := httprouter.New()
	router.DELETE("/v1/admin/user/:username", userAPIHandler.DeleteUser)
	ts := httptest.NewTLSServer(router)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	req, err := http.NewRequest("DELETE", url+"/v1/admin/user/Mokeuname", nil)
	if err != nil {
		fmt.Println(err)
	}

	resp, err := tc.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestAdminUserNew(t *testing.T) {
	userRepo := repository.NewMockUserRepository(nil)
	userService := service.NewUserServiceImpl(userRepo)

	// handler.NewA
	userAPIHandler := NewAdminUserHandler(userService)

	router := httprouter.New()
	router.POST("/v1/admin/user", userAPIHandler.PostUser)
	ts := httptest.NewTLSServer(router)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL
	user := entity.User{}

	user.Account = entity.UserMock.Account
	user.DeliveryAddress = entity.UserMock.DeliveryAddress
	user.Email = entity.UserMock.Email
	user.FirstName = entity.UserMock.FirstName
	user.ImagePath = entity.UserMock.ImagePath
	user.LastName = entity.UserMock.LastName
	user.Password = entity.UserMock.Password
	user.Phone = entity.UserMock.Phone
	user.Username = entity.UserMock.Username

	requestByte, _ := json.Marshal(user)
	requestBody := bytes.NewReader(requestByte)
	resp, err := tc.Post(sURL+"/v1/admin/user", "application/JSON", requestBody)

	if err != nil {
		t.Fatal(err)
	}

	//body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

// TestAdminCompanys handle ()
func TestAdminUsers(t *testing.T) {

	// tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))
	Userrepo := repository.NewMockUserRepository(nil)
	userserv := service.NewUserServiceImpl(Userrepo)
	Userhandler := NewAdminUserHandler(userserv)
	router := httprouter.New()
	router.GET("/v1/admin/user", Userhandler.GetUsers)
	ts := httptest.NewTLSServer(router)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/admin/user")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

// TestAdminCompanys handle ()
func TestAdminUser(t *testing.T) {

	// tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))
	Userrepo := repository.NewMockUserRepository(nil)
	userserv := service.NewUserServiceImpl(Userrepo)
	Userhandler := NewAdminUserHandler(userserv)
	router := httprouter.New()
	router.GET("/v1/admin/user/:username", Userhandler.GetSingleUser)
	ts := httptest.NewTLSServer(router)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/admin/user/Mokeuname")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

// TestAdminCompanysNew hanlde
func TestAdminUserUpdate(t *testing.T) {
	// tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))
	Userrepo := repository.NewMockUserRepository(nil)
	userserv := service.NewUserServiceImpl(Userrepo)
	Userhandler := NewAdminUserHandler(userserv)
	router := httprouter.New()
	router.PUT("/v1/admin/user/:username", Userhandler.PutUser)
	ts := httptest.NewTLSServer(router)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL
	User := entity.User{}
	User.Username = entity.UserMock.Username

	User.Account = entity.UserMock.Account
	User.DeliveryAddress = entity.UserMock.DeliveryAddress
	User.Email = entity.UserMock.Email
	User.FirstName = entity.UserMock.FirstName
	User.ImagePath = entity.UserMock.ImagePath

	User.LastName = entity.UserMock.LastName
	User.Password = entity.UserMock.Password

	User.Phone = entity.UserMock.Phone
	User.Username = entity.UserMock.Username
	requestByte, _ := json.Marshal(User)
	requestBody := bytes.NewReader(requestByte)
	req, _ := http.NewRequest("PUT", sURL+"/v1/admin/user/Mokeuname", requestBody)
	resp, err := tc.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	//body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

}
