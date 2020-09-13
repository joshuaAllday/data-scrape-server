package api

import (
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

type Handler struct {
	HandleFunc  func(http.ResponseWriter, *http.Request)
	HandlerName string
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func GetHandlerName(h func(http.ResponseWriter, *http.Request)) string {
	handlerName := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
	pos := strings.LastIndex(handlerName, ".")
	if pos != -1 && len(handlerName) > pos {
		handlerName = handlerName[pos+1:]
	}
	return handlerName
}

func (api *API) ApiHandler(h func(http.ResponseWriter, *http.Request)) http.Handler {
	handler := Handler{
		HandleFunc:  h,
		HandlerName: GetHandlerName(h),
	}

	return handler
}
