package tests

import "net/http"

type AuthRequest struct {
	Name    string
	Data    string
	Status  int
	Message map[string]interface{}
}

var LoginFixtures = []AuthRequest{
	{
		Name:   "Correct details",
		Data:   `{"username": "spankie", "password": "password"}`,
		Status: http.StatusOK,
		Message: map[string]interface{}{
			"message": "Login Successful",
			"status":  200,
			"data": map[string]interface{}{
				"token": "null",
			},
		},
	},
	{
		Name:   "Empty Password",
		Data:   `{"username": "spankie", "password": ""}`,
		Status: http.StatusUnauthorized,
		Message: map[string]interface{}{
			"message": "Username or password incorrect",
			"status":  http.StatusUnauthorized,
			"data":    nil,
		},
	},
	{
		Name:   "Empty username and password",
		Data:   `{"username": "", "password": ""}`,
		Status: http.StatusUnauthorized,
		Message: map[string]interface{}{
			"message": "Username or password incorrect",
			"status":  http.StatusUnauthorized,
			"data":    nil,
		},
	},
	{
		Name:   "No username, empty password",
		Data:   `{"password": ""}`,
		Status: http.StatusUnauthorized,
		Message: map[string]interface{}{
			"message": "Username or password incorrect",
			"status":  http.StatusUnauthorized,
			"data":    nil,
		},
	},
	{
		Name:   "Empty username, no password",
		Data:   `{"username": ""}`,
		Status: http.StatusUnauthorized,
		Message: map[string]interface{}{
			"message": "Username or password incorrect",
			"status":  http.StatusUnauthorized,
			"data":    nil,
		},
	},
}

var (
	newUsername = "chris1"
	newEmail    = "chris1@gmail.com"
)

var RegisterFixtures = []AuthRequest{
	{
		Name:   "Correct Details",
		Data:   `{"username": "chris", "password": "password", "first_name": "Chris", "last_name": "Powell", "email": "chris@gmail.com", "phone": "08103169311"}`,
		Status: http.StatusCreated,
		Message: map[string]interface{}{
			"message": "User created successfully",
			"status":  http.StatusCreated,
			"data":    nil,
		},
	},
	{
		Name:   "Missing First Name",
		Data:   `{"username": "` + newUsername + `", "password": "password", "last_name": "Powell", "email": "` + newEmail + `", "phone": "08103169310"}`,
		Status: http.StatusBadRequest,
		Message: map[string]interface{}{
			"message": "There are some problems in your forms",
			"status":  http.StatusBadRequest,
			"data": map[string]interface{}{
				"first_name": "You must enter a valid first name",
			},
		},
	},
	{
		Name:   "Missing Username",
		Data:   `{"password": "password", "first_name": "Chris", "last_name": "Powell", "email": "` + newEmail + `", "phone": "08103169310"}`,
		Status: http.StatusBadRequest,
		Message: map[string]interface{}{
			"message": "There are some problems in your forms",
			"status":  http.StatusBadRequest,
			"data": map[string]interface{}{
				"username": "You must enter a valid username",
			},
		},
	},
	{
		Name:   "Duplicate Username",
		Data:   `{"username": "spankie", "password": "password", "first_name": "Chris", "last_name": "Powell", "email": "` + newEmail + `", "phone": "08103169310"}`,
		Status: http.StatusBadRequest,
		Message: map[string]interface{}{
			"message": "There are some problems in your forms",
			"status":  http.StatusBadRequest,
			"data": map[string]interface{}{
				"username": "Username is taken",
			},
		},
	},
	{
		Name:   "Missing Password",
		Data:   `{"username": "` + newUsername + `", "first_name": "Chris", "last_name": "Powell", "email": "` + newEmail + `", "phone": "08103169310"}`,
		Status: http.StatusBadRequest,
		Message: map[string]interface{}{
			"message": "There are some problems in your forms",
			"status":  http.StatusBadRequest,
			"data": map[string]interface{}{
				"password": "You must enter a valid password: minimum of 6 characters",
			},
		},
	},
	{
		Name:   "Short password",
		Data:   `{"username": "` + newUsername + `", "password": "pa", "first_name": "Chris", "last_name": "Powell", "email": "` + newEmail + `", "phone": "08103169310"}`,
		Status: http.StatusBadRequest,
		Message: map[string]interface{}{
			"message": "There are some problems in your forms",
			"status":  http.StatusBadRequest,
			"data": map[string]interface{}{
				"password": "You must enter a valid password: minimum of 6 characters",
			},
		},
	},
	{
		Name:   "Missing Email",
		Data:   `{"username": "` + newUsername + `", "password": "password", "first_name": "Chris", "last_name": "Powell", "phone": "08103169310"}`,
		Status: http.StatusBadRequest,
		Message: map[string]interface{}{
			"message": "There are some problems in your forms",
			"status":  http.StatusBadRequest,
			"data": map[string]interface{}{
				"email": "You must enter a valid email address",
			},
		},
	},
	{
		Name:   "Duplicate Email",
		Data:   `{"username": "` + newUsername + `", "password": "password", "first_name": "Chris", "last_name": "Powell", "email": "edem@gmail.com", "phone": "08103169310"}`,
		Status: http.StatusBadRequest,
		Message: map[string]interface{}{
			"message": "There are some problems in your forms",
			"status":  http.StatusBadRequest,
			"data": map[string]interface{}{
				"email": "Email already exists",
			},
		},
	},
	{
		Name:   "Invalid Email",
		Data:   `{"username": "` + newUsername + `", "password": "password", "first_name": "Chris", "last_name": "Powell", "email": "gmail.com", "phone": "08103169310"}`,
		Status: http.StatusBadRequest,
		Message: map[string]interface{}{
			"message": "There are some problems in your forms",
			"status":  http.StatusBadRequest,
			"data": map[string]interface{}{
				"email": "You must enter a valid email address",
			},
		},
	},
	{
		Name:   "Invalid Phone Number",
		Data:   `{"username": "` + newUsername + `", "password": "password", "first_name": "Chris", "last_name": "Powell", "email": "gmail.com", "phone": "0810"}`,
		Status: http.StatusBadRequest,
		Message: map[string]interface{}{
			"message": "There are some problems in your forms",
			"status":  http.StatusBadRequest,
			"data": map[string]interface{}{
				"phone": "You must enter a valid phone number",
			},
		},
	},
	{
		Name:   "No Phone Number",
		Data:   `{"username": "` + newUsername + `", "password": "password", "first_name": "Chris", "last_name": "Powell", "email": "gmail.com"}`,
		Status: http.StatusBadRequest,
		Message: map[string]interface{}{
			"message": "There are some problems in your forms",
			"status":  http.StatusBadRequest,
			"data": map[string]interface{}{
				"phone": "You must enter a valid phone number",
			},
		},
	},
	{
		Name:   "Used Phone Number",
		Data:   `{"username": "` + newUsername + `", "password": "password", "first_name": "Chris", "last_name": "Powell", "email": "gmail.com", "phone": "08103169311"}`,
		Status: http.StatusBadRequest,
		Message: map[string]interface{}{
			"message": "There are some problems in your forms",
			"status":  http.StatusBadRequest,
			"data": map[string]interface{}{
				"phone": "Phone number is taken",
			},
		},
	},
}
