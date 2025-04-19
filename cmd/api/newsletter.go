package main

import (
	"canvas/internal/data"
	"canvas/validator"
	"net/http"
)

func (app *application) NewsletterSignup(w http.ResponseWriter, r *http.Request){

	var input struct{
		Email string `json:"email"`
	}

	err:= app.readJSON(w,r,&input)

	if err!=nil{
		app.badRequestResponse(w,r,err)
		return
	}

	// validate the email:
	v:= validator.New()
	data.ValidateEmail(v,input.Email)
	if !v.Valid(){
		app.failedValidationResponse(w,r,v.Errors)

		return
	}

	app.models.Newsletters.Insert(input.Email)

	data:= envelope{
		"newsletter":input.Email,
	}
	

	err = app.writeJSON(w,http.StatusCreated,data,nil)

	if err!=nil{
		app.serverErrorResponse(w,r,err)
		return
	}


	

		


}