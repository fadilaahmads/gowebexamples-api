package services

import (
	"net/http"
	"web-api/model"
)

func Forms(r *http.Request) *model.ContactDetails {
	details := model.ContactDetails{
		Email:   r.FormValue("email"),
		Subject: r.FormValue("subject"),
		Message: r.FormValue("message"),
	}
	return &details
}
