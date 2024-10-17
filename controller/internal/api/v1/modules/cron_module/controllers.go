package cron_module

import (
	"fmt"
	res "go-http/internal/api/response"
	"go-http/pkg/validate"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	GetCronsUrl   = "/cron"
	GetCronUrl    = "/cron/{id}"
	CreateCronUrl = "/cron"
	DeleteCronUrl = "/cron/{id}"
	UpdateCronUrl = "/cron/{id}"
)

func Router(api *chi.Mux) {
	api.Get(GetCronsUrl, GetCrons)
	api.Get(GetCronUrl, GetCron)
	api.Post(CreateCronUrl, CreateCron)
	api.Delete(DeleteCronUrl, DeleteCron)
	api.Put(UpdateCronUrl, UpdateCron)
}

// GET URL/cron
func GetCrons(w http.ResponseWriter, r *http.Request) {
	data, err := GetAllCronsSQL()
	if err != nil {
		fmt.Printf("Error occured while calling db: %v\n", err) // replace with logger
		res.Response(w, 400, nil, res.FailedDbConnMessage)
		return
	}

	res.Response(w, 200, data, "", res.PaginationLinks{})
}

// GET URL/cron/[id:string]
func GetCron(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	data, err := GetCronSQL(id)
	if err != nil {
		fmt.Printf("Error occured while calling db: %v\n", err) // replace with logger
		res.Response(w, 400, nil, res.FailedDbConnMessage)
		return
	}

	res.Response(w, 200, data, "")
}

// POST URL/cron
func CreateCron(w http.ResponseWriter, r *http.Request) {

	// validate the sent body
	payload := new(CreateCronBody)
	if err := validate.RequestBody(r, payload); err != nil {
		res.Response(w, 400, err, res.FailedPayloadValidationMessage)
		return
	}

	data, err := CreateCronSQL(*payload)
	if err != nil {
		fmt.Printf("Error occured while calling db: %v\n", err) // replace with logger
		res.Response(w, 400, nil, res.FailedDbConnMessage)
		return
	}

	res.Response(w, 200, data, "")
}

// DELETE URL/cron/[id:string]
func DeleteCron(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")

	data, err := DeleteCronSQL(id)
	if err != nil {
		fmt.Printf("Error occured while calling db: %v\n", err) // replace with logger
		res.Response(w, 400, nil, res.FailedDbConnMessage)
		return
	}

	res.Response(w, 200, data, "")
}

// PUT URL/cron/[id:string]
func UpdateCron(w http.ResponseWriter, r *http.Request) {

	// validate the sent body
	payload := new(UpdateCronBody)
	if err := validate.RequestBody(r, payload); err != nil {
		res.Response(w, 400, err, res.FailedPayloadValidationMessage)
		return
	}

	id := chi.URLParam(r, "id")

	data, err := UpdateCronSQL(id, *payload)
	if err != nil {
		fmt.Printf("Error occured while calling db: %v\n", err) // replace with logger
		res.Response(w, 400, nil, res.FailedDbConnMessage)
		return
	}

	res.Response(w, 200, data, "")
}
