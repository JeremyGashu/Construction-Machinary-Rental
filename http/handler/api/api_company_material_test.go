package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	repository "github.com/ermiasgashu/Construction-Machinary-Rental/company/repository"
	service "github.com/ermiasgashu/Construction-Machinary-Rental/company/service"
	"github.com/julienschmidt/httprouter"

	"github.com/ermiasgashu/Construction-Machinary-Rental/entity"
)

// TestAdminCompanys handle ()
func TestCompanyMaterials(t *testing.T) {

	// tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))
	companyrepo := repository.NewMockMaterialRepository(nil)
	companyserv := service.NewMaterialService(companyrepo)
	Companyhandler := NewCompanyMaterialHandler(companyserv)
	router := httprouter.New()
	router.GET("/v1/companies/materials", Companyhandler.Materials)
	ts := httptest.NewTLSServer(router)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/companies/materials")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

// TestAdminCompanys handle ()
func TestCompanyMaterial(t *testing.T) {

	// tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))
	companyrepo := repository.NewMockMaterialRepository(nil)
	companyserv := service.NewMaterialService(companyrepo)
	Companyhandler := NewCompanyMaterialHandler(companyserv)
	router := httprouter.New()
	router.GET("/v1/companies/materials/:material_id", Companyhandler.Material)
	ts := httptest.NewTLSServer(router)
	defer ts.Close()

	tc := ts.Client()
	url := ts.URL

	resp, err := tc.Get(url + "/v1/companies/materials/1")
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

// TestAdminCompanysNew hanlde
func TestCompanyMaterialNew(t *testing.T) {
	// tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))
	companyrepo := repository.NewMockMaterialRepository(nil)
	companyserv := service.NewMaterialService(companyrepo)
	Companyhandler := NewCompanyMaterialHandler(companyserv)
	router := httprouter.New()
	router.POST("/v1/companies/materials", Companyhandler.StoreMaterial)
	ts := httptest.NewTLSServer(router)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL
	material := entity.Material{}
	material.ID = entity.MaterialMock.ID
	material.Discount = entity.MaterialMock.Discount
	material.ImagePath = entity.MaterialMock.ImagePath
	material.Name = entity.MaterialMock.Name
	material.OnDiscount = entity.MaterialMock.OnDiscount
	material.OnSale = entity.MaterialMock.OnSale
	material.Owner = entity.MaterialMock.Owner
	material.PricePerDay = entity.MaterialMock.PricePerDay
	requestByte, _ := json.Marshal(material)
	requestBody := bytes.NewReader(requestByte)
	resp, err := tc.Post(sURL+"/v1/companies/materials", "application/JSON", requestBody)

	if err != nil {
		t.Fatal(err)
	}

	//body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}

}

// TestAdminCompanysUpdate handle
func TestCompanyMaterialUpdate(t *testing.T) {

	// tmpl := template.Must(template.ParseGlob("../../ui/templates/*"))
	companyrepo := repository.NewMockMaterialRepository(nil)
	companyserv := service.NewMaterialService(companyrepo)
	Companyhandler := NewCompanyMaterialHandler(companyserv)
	router := httprouter.New()
	router.PUT("/v1/companies/materials/:id", Companyhandler.UpdateMaterial)
	ts := httptest.NewTLSServer(router)
	defer ts.Close()

	tc := ts.Client()
	sURL := ts.URL
	material := entity.Material{}
	material.ID = entity.MaterialMock.ID
	material.Discount = entity.MaterialMock.Discount
	material.ImagePath = entity.MaterialMock.ImagePath
	material.Name = entity.MaterialMock.Name
	material.OnDiscount = entity.MaterialMock.OnDiscount
	material.OnSale = entity.MaterialMock.OnSale
	material.Owner = entity.MaterialMock.Owner
	material.PricePerDay = entity.MaterialMock.PricePerDay
	requestByte, _ := json.Marshal(material)
	requestBody := bytes.NewReader(requestByte)
	req, _ := http.NewRequest("PUT", sURL+"/v1/companies/materials/1", requestBody)
	resp, err := tc.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	//body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}

// TestAdminCompanysDelete handle
func TestCompanyMaterialDelete(t *testing.T) {
	companyrepo := repository.NewMockMaterialRepository(nil)
	companyserv := service.NewMaterialService(companyrepo)
	Companyhandler := NewCompanyMaterialHandler(companyserv)
	router := httprouter.New()
	router.DELETE("/v1/companies/materials/delete/:material_id", Companyhandler.DeleteMaterial)
	ts := httptest.NewTLSServer(router)
	defer ts.Close()
	tc := ts.Client()
	sURL := ts.URL
	req, err := http.NewRequest("DELETE", sURL+"/v1/companies/materials/delete/1", nil)
	if err != nil {
		fmt.Println(err)
	}
	resp, err := tc.Do(req)
	if err != nil {
		t.Fatal(err)
	}

	//body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("want %d, got %d", http.StatusOK, resp.StatusCode)
	}
}
