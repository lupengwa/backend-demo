package api

import (
	"demo-backend/internal/api/restutils"
	"net/http"
)

type Handler interface {
	GetRestUriToHandlerConfig() map[restutils.RestApiUriKey]http.HandlerFunc
}
