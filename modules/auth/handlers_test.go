package auth

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/ChrisPowellIinc/allofusserver/internal/jwt"

	"github.com/ChrisPowellIinc/allofusserver/migrations"
	"github.com/ChrisPowellIinc/allofusserver/models"
	"github.com/ChrisPowellIinc/allofusserver/tests"
	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	log.Println("Start test seeding")
	db := tests.Con.DB
	user := models.User{
		FirstName:      "Spankie",
		LastName:       "Dee",
		Phone:          "08103169310",
		Email:          "edem@gmail.com",
		Username:       "spankie",
		PasswordString: "password",
		Image:          "me.jpg",
	}
	// Create
	var err error
	user.Password, err = bcrypt.GenerateFromPassword([]byte(user.PasswordString), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Could not generate password")
	}
	db.C("user").Insert(&user)
	log.Println("End test seeding")
}

func TestMain(m *testing.M) {
	// tests.Con.DB.DropTableIfExists(&models.User{})
	migrations.MakeMigrations(tests.Con)
	Seed()
	jwt.Register([]byte(tests.Con.JWTSecret))
	log.Println("Starting tests")
	exitCode := m.Run()
	err := tests.Con.DB.DropDatabase()
	if err != nil {
		log.Println("Error dropping the database; ", err)
	}
	os.Exit(exitCode)
}

func TestLoginHandler(t *testing.T) {
	log.Println("== Login Tests ==")
	handler := New(tests.Con)
	data := tests.LoginFixtures
	for _, d := range data {
		// log.Println(d.Name)
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest("POST", "/", bytes.NewBuffer([]byte(d.Data)))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		loginHandler := http.HandlerFunc(handler.Login)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		loginHandler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != d.Status {
			t.Errorf("handler returned wrong status code: got %v want %v", status, d.Status)
		}
		// Check the response body if it is what we expect.
		responseBody, err := ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Fatal("Could not read the body of the response")
		}
		data := make(map[string]interface{})
		json.Unmarshal(responseBody, &data)
		if data["message"] != d.Message["message"] {
			t.Fatalf("handler returned unexpected body[message]: got %v want %v", data["message"], d.Message["message"])
		}

		if data["status"] != float64(d.Message["status"].(int)) {
			t.Fatalf("handler returned unexpected body[status]: got %v want %v", data["status"], d.Message["status"])
		}

		if reflect.TypeOf(data["data"]) != reflect.TypeOf(d.Message["data"]) {
			t.Fatalf("incompatible types: want: %T; got %T", d.Message["data"], data["data"])
		}

		if data["data"] != nil {
			token, ok := data["data"].(map[string]interface{})["token"].(string)
			if !ok {
				t.Fatalf("token is not a string: %v; type: %T", data["data"].(map[string]interface{})["token"], data["data"].(map[string]interface{})["token"])
			}
			if len(token) <= 10 {
				t.Fatalf("handler returned unexpected token: got %v want token string longer than %v", token, len(token))
			}
		}

	}
}

func TestRegisterHandler(t *testing.T) {
	log.Println("== Register Tests ==")
	handler := New(tests.Con)
	data := tests.RegisterFixtures
	for _, d := range data {
		log.Println(d.Name)
		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		req, err := http.NewRequest("POST", "/register", bytes.NewBuffer([]byte(d.Data)))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
		// req.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(handler.Register)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != d.Status {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
		// Check the response body if it is what we expect.
		responseBody, err := ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Fatal("Could not read the body of the response")
		}
		data := make(map[string]interface{})
		json.Unmarshal(responseBody, &data)
		if data["message"] != d.Message["message"] {
			t.Fatalf("handler returned unexpected body[message]: got %v want %v", data["message"], d.Message["message"])
		}

		if data["status"] != float64(d.Message["status"].(int)) {
			t.Fatalf("handler returned unexpected body[status]: got %v want %v", data["status"], d.Message["status"])
		}

		if reflect.TypeOf(data["data"]) != reflect.TypeOf(d.Message["data"]) {
			t.Fatalf("incompatible types: got %T want: %T", data["data"], d.Message["data"])
		}

		if data["data"] != nil {
			respData, ok := data["data"].(map[string]interface{})
			if !ok {
				t.Fatalf("response data is not what was expected: got %T; want: %T", data["data"], d.Message["data"])
			}
			// Loop through the map to get the ones that match what is expected
			for i, v := range d.Message["data"].(map[string]interface{}) {
				if respData[i] != v {
					t.Fatalf("Unexpected %v Error: got %v want %v", i, respData[i], v)
				}
			}
		}
	}
}
