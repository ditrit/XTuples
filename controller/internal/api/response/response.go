// Package response is used for returning a predictable json object
// to the client. It is used in the controllers of the modules.
// The package also contains predefined error messages.
package response

import (
	"encoding/json"
	"net/http"
)

const (
	ParamIsNotIntMessage           = "Only integers as URL params allowed!"
	FailedDbConnMessage            = "Error while calling the database!"
	FailedPayloadValidationMessage = "Payload validation failed!"
)

// Baseline Response for the controllers
type Res struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// Baseline Response + Pagination links
type ResWithPagination struct {
	Res   `json:",inline"`
	Links PaginationLinks `json:"links,omitempty"`
}

type PaginationLinks struct {
	Prev string `json:"prev,omitempty"`
	Next string `json:"next,omitempty"`
}

// Response is used in the controllers to return an predictable
// json object. There is a 4th optional parameter, which is
// used to send pagination links to the client, if needed.
//
// All of the correct status codes can be found here:
// https://pkg.go.dev/net/http?utm_source=gopls#StatusOK
func Response(w http.ResponseWriter, status int, data interface{}, message string, links ...PaginationLinks) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	res := Res{Status: status, Message: message, Data: data}

	if len(links) == 1 {
		return json.NewEncoder(w).Encode(&ResWithPagination{
			Res:   res,
			Links: links[0]})
	}

	return json.NewEncoder(w).Encode(&res)
}
