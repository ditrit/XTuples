package exec_module

import (
	"fmt"
	res "go-http/internal/api/response"
	"go-http/pkg/validate"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const (
	ExecUrl = "/exec"
)

func Router(api *chi.Mux) {
	api.Post(ExecUrl, Exec)
}

// POST URL/exec
func Exec(w http.ResponseWriter, r *http.Request) {

	// validate the sent body
	payload := new(ExecBody)
	if err := validate.RequestBody(r, payload); err != nil {
		res.Response(w, 400, err, res.FailedPayloadValidationMessage)
		return
	}

	// replace the function to return the data from the db
	data, err := placeholderFunc()
	if err != nil {
		fmt.Printf("Error occured while calling db: %v\n", err) // replace with logger
		res.Response(w, 400, nil, res.FailedDbConnMessage)
		return
	}

	res.Response(w, 200, data, "")
}

// Placeholder function to be replaced in the controllers after codegen.
// Usually replaced with the function which returns the data from the db.
func placeholderFunc() (struct{}, error) {
	var placeholder struct{}
	return placeholder, nil
}
