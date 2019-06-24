package messaging

import (
	"bytes"
	"encoding/json"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
)

type User struct {
	UserID string `json:"subject"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Name   string `json:"name"`
}

var (
	accessToken = os.Getenv("INTERCOM_ACCESS_TOKEN")
	email       = os.Getenv("INTERCOM_EMAIL")
	adminID     = os.Getenv("INTERCOM_ADMIN_ID")
	userID      = os.Getenv("INTERCOM_USER_ID")
)

var _ = Describe("Intercom Create User with all required data", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	user := User{UserID: "777", Email: email, Phone: "1234567890", Name: "Testcase User"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(user)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createuser", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(recorder, request)

	Describe("Create user for intercom", func() {
		Context("create user", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Intercom Create User without required data", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	user := User{Phone: "1234567890", Name: "Testcase User"}
	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(user)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/createuser", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(CreateUser)
	handler.ServeHTTP(recorder, request)

	Describe("Create user for intercom", func() {
		Context("create user", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Intercom messaging, send through InApp with all required data", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	var message Message
	message.From = json.Number(adminID)
	message.To = email
	message.Body = "Test case body for inapp"

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(message)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/inappmessage", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendInAppMessage)
	handler.ServeHTTP(recorder, request)

	Describe("Send message within InApp", func() {
		Context("InApp Message", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Intercom messaging, send through InApp without required data", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	var message Message
	message.From = json.Number(adminID)
	message.To = email

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(message)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/inappmessage", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendInAppMessage)
	handler.ServeHTTP(recorder, request)

	Describe("Send message within InApp", func() {
		Context("InApp Message", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Intercom messaging, send through Email with all required data", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	var message Message
	message.UserID = userID
	message.To = email
	message.Subject = "Test case subject"
	message.Body = "Test case body for email"

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(message)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/emailmessage", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendEmailMessage)
	handler.ServeHTTP(recorder, request)

	Describe("Send message within Email", func() {
		Context("Email Message", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Intercom messaging, send through Email without required data", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	var message Message
	message.Subject = "Test case subject"
	message.Body = "Test case body for email"

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(message)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/emailmessage", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendEmailMessage)
	handler.ServeHTTP(recorder, request)

	Describe("Send message within Email", func() {
		Context("Email Message", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})

var _ = Describe("Intercom messaging, send through User with all required data", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	var message Message
	message.Email = email
	message.Body = "Test case body for email"

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(message)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/usermessage", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendUserMessage)
	handler.ServeHTTP(recorder, request)

	Describe("Send message within User", func() {
		Context("User Message", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusOK))
			})
		})
	})
})

var _ = Describe("Intercom messaging, send through User without required data", func() {

	os.Setenv("ACCESS_TOKEN", accessToken)

	var message Message
	message.Body = "Test case body for email"

	requestBody := new(bytes.Buffer)
	errr := json.NewEncoder(requestBody).Encode(message)
	if errr != nil {
		log.Fatal(errr)
	}

	request, err := http.NewRequest("POST", "/usermessage", requestBody)
	if err != nil {
		log.Fatal(err)
	}
	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(SendUserMessage)
	handler.ServeHTTP(recorder, request)

	Describe("Send message within User", func() {
		Context("User Message", func() {
			It("Should result http.StatusOK", func() {
				Expect(recorder.Code).To(Equal(http.StatusBadRequest))
			})
		})
	})
})
