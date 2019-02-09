package user

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/ChrisPowellIinc/allofusserver/internal/jwt"
	"github.com/ChrisPowellIinc/allofusserver/tests"
	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

var token string

func TestMain(m *testing.M) {
	// run main test here...
	token = tests.GetAuth()
	if token == "" {
		log.Println("No token")
		return
	}
	log.Println(token)
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestUploadProfilePic(t *testing.T) {
	router := chi.NewRouter()
	log.Println("Testing Upload Profile Picture")
	router.Use(jwtauth.Verifier(jwt.TokenAuth))
	router.Use(jwt.AuthHandler)
	handler := New(tests.Con)

	router.HandleFunc("/upload", http.HandlerFunc(handler.UploadProfilePic))
	ts := httptest.NewServer(router)
	defer ts.Close()

	url := ts.URL + "/upload"
	filename := "cool-background.png"
	formFileName := "profile_picture"

	rr, err := tests.PostFile(filename, url, formFileName, token)
	if err != nil {
		t.Fatalf("Could not Post file: %v", err)
	}
	if status := rr.StatusCode; status != http.StatusOK {
		t.Errorf("Invalid status code; got %v want %v", status, http.StatusOK)
	}
	// t.Errorf("Testing testing %v", rr.StatusCode)
}
