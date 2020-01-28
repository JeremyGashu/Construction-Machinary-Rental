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

func TestAdminAdmins(t *testing.T) {
	// tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))

	adminRepo := repository.NewMockAdminRepo(nil)
	adminService := service.NewAdminServiceImpl(adminRepo)

	// handler.NewA
	adminAPIHandler := NewAdminAdminsHandler(adminService)

	router := httprouter.New()
	router.GET("/v1/admin/admins", adminAPIHandler.GetAdmins)
	ts := httptest.NewTLSServer(router)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/admin/admins")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

}

func TestAdminAdmin(t *testing.T) {
	// tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))

	adminRepo := repository.NewMockAdminRepo(nil)
	adminService := service.NewAdminServiceImpl(adminRepo)

	// handler.NewA
	adminAPIHandler := NewAdminAdminsHandler(adminService)

	router := httprouter.New()
	router.GET("/v1/admin/admins/:username", adminAPIHandler.GetSingleAdmin)
	ts := httptest.NewTLSServer(router)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/admin/admins/1")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

}

func TestAdminAdminDelete(t *testing.T) {

	adminRepo := repository.NewMockAdminRepo(nil)
	adminService := service.NewAdminServiceImpl(adminRepo)
	adminAPIHandler := NewAdminAdminsHandler(adminService)
	router := httprouter.New()
	router.DELETE("/v1/admin/admins/:username", adminAPIHandler.DeleteAdmin)
	ts := httptest.NewTLSServer(router)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	req, err := http.NewRequest("DELETE", url+"/v1/admin/admins/Mock uName", nil)
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
func TestAdminNew(t *testing.T) {

	adminRepo := repository.NewMockAdminRepo(nil)
	adminService := service.NewAdminServiceImpl(adminRepo)

	// handler.NewA
	adminAPIHandler := NewAdminAdminsHandler(adminService)

	router := httprouter.New()
	router.POST("/v1/admin/admins", adminAPIHandler.PostAdmin)
	ts := httptest.NewTLSServer(router)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL
	admin := entity.Admin{}
	admin.Email = entity.AdminMock.Email
	admin.FirstName = entity.AdminMock.FirstName
	admin.ImagePath = entity.AdminMock.ImagePath
	admin.LastName = entity.AdminMock.LastName
	admin.Password = entity.AdminMock.Password
	admin.Username = entity.AdminMock.Username
	requestByte, _ := json.Marshal(admin)
	requestBody := bytes.NewReader(requestByte)
	resp, err := tc.Post(sURL+"/v1/admin/admins", "application/JSON", requestBody)

	if err != nil {
		t.Fatal(err)
	}

	//body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

}

func TestAdminPut(t *testing.T) {
	adminRepo := repository.NewMockAdminRepo(nil)
	adminService := service.NewAdminServiceImpl(adminRepo)

	// handler.NewA
	adminAPIHandler := NewAdminAdminsHandler(adminService)

	router := httprouter.New()
	router.PUT("/v1/admin/admins/:username", adminAPIHandler.PutAdmin)
	ts := httptest.NewTLSServer(router)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL
	admin := entity.Admin{}
	admin.Email = entity.AdminMock.Email
	admin.FirstName = entity.AdminMock.FirstName
	admin.ImagePath = entity.AdminMock.ImagePath
	admin.LastName = entity.AdminMock.LastName
	admin.Password = entity.AdminMock.Password
	admin.Username = entity.AdminMock.Username
	requestByte, _ := json.Marshal(admin)
	requestBody := bytes.NewReader(requestByte)
	req, err := http.NewRequest("PUT", sURL+"/v1/admin/admins/Mock uName", requestBody)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := tc.Do(req)

	if err != nil {
		t.Fatal(err)
	}

	//body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

}
