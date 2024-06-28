package demo

import (
	"demo-backend/internal/api/restutils"
	"fmt"
	"log"
	"net/http"
)

const (
	UserApiBasePath = "/movie"
)

type ApiHandler struct {
	demoService Service
}

func NewApiHandler(demoService Service) *ApiHandler {
	if demoService == nil {
		log.Panic("demo service can't be nil")
	}
	return &ApiHandler{demoService: demoService}
}

func (handler *ApiHandler) GetRestUriToHandlerConfig() map[restutils.RestApiUriKey]http.HandlerFunc {
	return map[restutils.RestApiUriKey]http.HandlerFunc{
		restutils.RestApiUriKey{
			HttpMethod: http.MethodPost,
			Path:       UserApiBasePath,
		}: handler.SearchMovie,
	}
}

func (handler *ApiHandler) SearchMovie(w http.ResponseWriter, r *http.Request) {
	var movieRequest MovieRequestDto

	if err := restutils.UnmarshalJSONRequest(r, &movieRequest); err != nil {
		restutils.ToErrorResponse(w, r, err, http.StatusBadRequest)
		return
	}

	fmt.Println("Request received", movieRequest.Name)

	movieResponse, err := handler.demoService.SearchMovie(movieRequest.Name)
	if err != nil {
		return
	}
	restutils.ToSuccessPayloadResponse(w, r, movieResponse)
}
