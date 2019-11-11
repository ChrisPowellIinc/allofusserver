package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"github.com/stripe/stripe-go/paymentmethod"

	"github.com/stripe/stripe-go/customer"

	"github.com/stripe/stripe-go/setupintent"

	// "github.com/aws/aws-sdk-go/aws/session"
	"github.com/globalsign/mgo/bson"

	"github.com/ChrisPowellIinc/allofusserver/models"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/checkout/session"
	"golang.org/x/crypto/bcrypt"

	"github.com/go-chi/render"
)

// Me return the logged in user
func (handler *Handler) Me(w http.ResponseWriter, r *http.Request) {
	// get logged in user:
	userEmail, err := models.GetLoggedInUserID(r.Context())
	if err != nil {
		log.Printf("error getting logged in user: %v\n", err)
		models.HandleResponse(w, r, "Could not authenticate user", http.StatusBadRequest, nil)
		return
	}
	db := handler.config.DB
	user := &models.User{}
	err = db.C("user").Find(bson.M{"email": userEmail}).One(user) //bson.M{"$eq": userEmail}
	if err != nil {
		models.HandleResponse(w, r, "Could not get user", http.StatusInternalServerError, nil)
		return
	}
	models.HandleResponse(w, r, "User", http.StatusOK, map[string]interface{}{"user": user})
}

// Login : Logs in an existing user
func (handler *Handler) Login(w http.ResponseWriter, r *http.Request) {
	res := models.Response{}
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		res.Message = "Username or password incorrect"
		res.Status = http.StatusUnauthorized
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, res)
		return
	}
	// Check if the user with that username exists
	db := handler.config.DB
	userPassword := user.PasswordString
	// Get first matched record
	err = db.C("user").Find(bson.M{"username": user.Username}).One(&user)
	if err != nil {
		log.Println(err)
		res.Message = "Username or password incorrect"
		res.Status = http.StatusUnauthorized
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, res)
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userPassword))
	if err != nil {
		log.Println("Passwords do not match")
		res.Message = "Username or password incorrect"
		res.Status = http.StatusUnauthorized
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, res)
		return
	}

	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"user_email": user.Email})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(handler.config.Constants.JWTSecret))

	// _, token, err := jwt.TokenAuth.Encode(jwtauth.Claims{"user_email": user.Email})

	if err != nil {
		log.Println(err)
		res.Message = "Username or password incorrect"
		res.Status = http.StatusInternalServerError
		render.JSON(w, r, res)
		return
	}

	res = models.Response{
		Message: "Login Successful",
		Status:  http.StatusOK,
		Data: map[string]interface{}{
			"token":      tokenString,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"phone":      user.Phone,
			"email":      user.Email,
			"username":   user.Username,
			"image":      user.Image,
		},
	}
	render.Status(r, http.StatusOK)
	render.JSON(w, r, res)
}

// Register : Registers a user
func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var res models.Response
	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
		res.Message = "Username or password incorrect"
		res.Status = http.StatusUnauthorized
		render.Status(r, http.StatusUnauthorized)
		render.JSON(w, r, res)
		return
	}
	errs := user.Validate(handler.config)
	if len(errs) > 0 {
		log.Println(errs)
		res.Message = "There are some problems in your forms"
		res.Status = http.StatusBadRequest
		res.Data = errs
		render.Status(r, http.StatusBadRequest)
		render.JSON(w, r, res)
		return
	}

	user.Password, err = bcrypt.GenerateFromPassword([]byte(user.PasswordString), bcrypt.DefaultCost)
	if err != nil {
		res.Message = "Sorry a problem occured, please try again"
		res.Status = http.StatusInternalServerError
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res)
		return
	}

	db := handler.config.DB
	err = db.C("user").Insert(&user)
	if err != nil {
		res.Message = "Sorry a problem occured, please try again"
		res.Status = http.StatusInternalServerError
		render.Status(r, http.StatusInternalServerError)
		render.JSON(w, r, res)
		return
	}

	res.Message = "User created successfully"
	res.Status = http.StatusCreated
	render.Status(r, http.StatusCreated)
	render.JSON(w, r, res)
}

// GetStripeSessionID gets session if from stripe
func (handler *Handler) GetStripeSessionID(w http.ResponseWriter, r *http.Request) {
	// get logged in user:
	userEmail, err := models.GetLoggedInUserID(r.Context())
	if err != nil {
		log.Printf("error getting logged in user: %v\n", err)
		models.HandleResponse(w, r, "Could not authenticate user", http.StatusBadRequest, nil)
		return
	}
	db := handler.config.DB
	user := &models.User{}
	err = db.C("user").Find(bson.M{"email": userEmail}).One(user) //bson.M{"$eq": userEmail}
	if err != nil {
		models.HandleResponse(w, r, "Could not get user", http.StatusInternalServerError, nil)
		return
	}
	// Set your secret key: remember to change this to your live secret key in production
	// See your keys here: https://dashboard.stripe.com/account/apikeys
	stripe.Key = os.Getenv("SK_STRIPE")
	params := &stripe.CheckoutSessionParams{
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		CustomerEmail: stripe.String(userEmail),
		// Customer:      stripe.String(user.CustomerID),
		Mode:       stripe.String(string(stripe.CheckoutSessionModeSetup)),
		SuccessURL: stripe.String(fmt.Sprintf("%s?session_id={CHECKOUT_SESSION_ID}", os.Getenv("SUCCESS_URL"))),
		CancelURL:  stripe.String(os.Getenv("CANCEL_URL")),
	}
	// stripe.
	session, err := session.New(params)
	if err != nil {
		log.Printf("Error getting session: %v\n", err)
		models.HandleResponse(w, r, "Internal server error", http.StatusInternalServerError, nil)
		return
	}
	// store the session id in the users model
	user.SessionID = session.ID
	err = db.C("user").Update(bson.M{"email": user.Email}, user)
	if err != nil {
		models.HandleResponse(w, r, "Could not update user", http.StatusInternalServerError, nil)
		return
	}

	models.HandleResponse(w, r, "Session retrieved successfuly", http.StatusOK, map[string]interface{}{"data": session})
	return
}

func (handler *Handler) StripeWebhookHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{}
	err := json.NewDecoder(r.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		models.HandleResponse(w, r, "Error with request", http.StatusBadRequest, nil)
		return
	}
	b, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("\nWEBHOOK REQUEST: %v\n", string(b))
	// log.Printf("\nTYPE: %+v\n", response["type"])
	// log.Printf("\nDATA: %+v\n", response["data"].(map[string]interface{})["object"].(map[string]interface{})["setup_intent"].(string))
	if response["type"].(string) == "checkout.session.completed" {
		setupIntent, ok := response["data"].(map[string]interface{})["object"].(map[string]interface{})["setup_intent"].(string)
		userEmail, ok := response["data"].(map[string]interface{})["object"].(map[string]interface{})["customer_email"].(string)
		if !ok {
			log.Printf("Cannot fetch the setup intent\n")
		}
		db := handler.config.DB
		user := &models.User{}
		err = db.C("user").Find(bson.M{"email": userEmail}).One(user)
		if err != nil {
			models.HandleResponse(w, r, "Could not get user", http.StatusInternalServerError, nil)
			return
		}
		// Set your secret key: remember to change this to your live secret key in production
		// See your keys here: https://dashboard.stripe.com/account/apikeys
		stripe.Key = os.Getenv("SK_STRIPE")
		s, _ := setupintent.Get(setupIntent, nil)
		user.SetupIntentID = setupIntent // s.ID
		user.PaymentMethodID = s.PaymentMethod.ID
		b, _ := json.MarshalIndent(s, "", "  ")
		log.Printf("SetupIntent: %v", string(b))
		var cus *stripe.Customer
		var cusErr error
		if user.CustomerID == "" {
			// create the customer and add the payment method
			cus, cusErr = customer.New(&stripe.CustomerParams{
				Email:         stripe.String(userEmail),
				Name:          stripe.String(user.FirstName),
				Phone:         stripe.String(user.Phone),
				PaymentMethod: stripe.String(user.PaymentMethodID),
				InvoiceSettings: &stripe.CustomerInvoiceSettingsParams{
					DefaultPaymentMethod: stripe.String(user.PaymentMethodID),
				},
			})
			user.CustomerID = cus.ID // user will be updated after the if statement
		} else {
			// update the customer payment method:
			params := &stripe.PaymentMethodAttachParams{
				Customer: stripe.String(user.CustomerID),
			}
			_, cusErr = paymentmethod.Attach("pm_card_visa", params)
		}
		if cusErr != nil {
			log.Printf("\n%v\n", err)
		} else {
			cusJSON, _ := json.MarshalIndent(cus, "", "  ")
			log.Printf("customer: %v", string(cusJSON))
			err = db.C("user").Update(bson.M{"email": user.Email}, user)
			if err != nil {
				models.HandleResponse(w, r, "Could not update user", http.StatusInternalServerError, nil)
				return
			}
		}
		// cusIter := customer.List(&stripe.CustomerListParams{}) // New(params)
		// for cusIter.Next() {
		// 	js, _ := json.MarshalIndent(cusIter.Current(), "", "  ")
		// 	log.Println(string(js))
		// }

	}
	models.HandleResponse(w, r, "Success", http.StatusOK, nil)
	return
}

// MyCards return the logged in user
func (handler *Handler) MyCards(w http.ResponseWriter, r *http.Request) {
	// Set your secret key: remember to change this to your live secret key in production
	// See your keys here: https://dashboard.stripe.com/account/apikeys
	stripe.Key = os.Getenv("SK_STRIPE")
	// get logged in user:
	userEmail, err := models.GetLoggedInUserID(r.Context())
	if err != nil {
		log.Printf("error getting logged in user: %v\n", err)
		models.HandleResponse(w, r, "Could not authenticate user", http.StatusBadRequest, nil)
		return
	}
	db := handler.config.DB
	user := &models.User{}
	err = db.C("user").Find(bson.M{"email": userEmail}).One(user) //bson.M{"$eq": userEmail}
	if err != nil {
		models.HandleResponse(w, r, "Could not get user", http.StatusInternalServerError, nil)
		return
	}
	pmsList := paymentmethod.List(&stripe.PaymentMethodListParams{
		Customer: stripe.String(user.CustomerID),
		Type:     stripe.String("card"),
	})
	// cs := card.List(&stripe.CardListParams{
	// 	Customer: stripe.String(user.CustomerID),
	// })
	pms := []*stripe.PaymentMethod{}
	for pmsList.Next() {
		pms = append(pms, pmsList.PaymentMethod())
	}
	// pm, _ := paymentmethod.Get(user.PaymentMethodID, &stripe.PaymentMethodParams{})
	models.HandleResponse(w, r, "User", http.StatusOK, map[string]interface{}{"pms": pms})
}

// DeleteCard deletes the logged in user's card
func (handler *Handler) DeleteCard(w http.ResponseWriter, r *http.Request) {
	pmID := chi.URLParam(r, "pmID")
	// Set your secret key: remember to change this to your live secret key in production
	// See your keys here: https://dashboard.stripe.com/account/apikeys
	stripe.Key = os.Getenv("SK_STRIPE")
	// get logged in user:
	// userEmail, err := models.GetLoggedInUserID(r.Context())
	// if err != nil {
	// 	log.Printf("error getting logged in user: %v\n", err)
	// 	models.HandleResponse(w, r, "Could not authenticate user", http.StatusBadRequest, nil)
	// 	return
	// }
	// db := handler.config.DB
	// user := &models.User{}
	// err = db.C("user").Find(bson.M{"email": userEmail}).One(user) //bson.M{"$eq": userEmail}
	// if err != nil {
	// 	models.HandleResponse(w, r, "Could not get user", http.StatusInternalServerError, nil)
	// 	return
	// }
	pm, err := paymentmethod.Detach(pmID, nil)
	if err != nil {
		log.Printf("Unable to detach paymentmethod\n")
		models.HandleResponse(w, r, "Unable to delete card", http.StatusInternalServerError, nil)
		return
	}
	models.HandleResponse(w, r, "Card deleted successfully", http.StatusOK, map[string]interface{}{"pm": pm})
}
