package main

import (
	"bytes"
	"github.com/smakaroni/epulse-cats/pkg/app"
	"github.com/smakaroni/epulse-cats/pkg/config"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var a app.App

func TestMain(m *testing.M) {
	c, err := config.LoadTestConfig()
	if err != nil {
		log.Fatalln("failed to load config:", err)
	}
	a.Init(&c)

	code := m.Run()

	os.Exit(code)

}

func TestGetAll(t *testing.T) {
	req, _ := http.NewRequest("GET", "/cats/list", nil)
	res := executeReq(req)

	if res.Code != http.StatusOK {
		t.Errorf("expected response code %d but got %d", http.StatusOK, res.Code)
	}
}

func TestCreateCat(t *testing.T) {
	var catJson = []byte(`{"name": "testing", "description": "test creating cat"}`)
	req, _ := http.NewRequest("POST", "/cats/create", bytes.NewBuffer(catJson))

	res := executeReq(req)

	if res.Code != http.StatusCreated {
		t.Errorf("expected resonse code %d but got %d", http.StatusCreated, res.Code)
	}
}

func executeReq(req *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	a.Router.ServeHTTP(recorder, req)

	return recorder
}
