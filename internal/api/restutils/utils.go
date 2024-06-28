package restutils

import (
	"errors"
	"fmt"
	"github.com/go-chi/render"
	"net/http"
)

func ToSuccessPayloadResponse(w http.ResponseWriter, r *http.Request, resp interface{}) {
	render.Status(r, http.StatusOK)
	render.Respond(w, r, resp)
}

func ToErrorResponse(w http.ResponseWriter, r *http.Request, err error, httpStatusCode int) {
	render.Status(r, httpStatusCode)
	errorResponse := &ErrResponse{
		StatusText: "Invalid request.",
		ErrorText:  err.Error(),
	}
	render.Respond(w, r, errorResponse)
}

func UnmarshalJSONRequest(req *http.Request, payload interface{}) error {
	if payload == nil {
		return errors.New("input payload is empty")
	}
	if err := render.Decode(req, payload); err != nil {
		return fmt.Errorf("error decoding request body: %w", err)
	}
	return nil
}
