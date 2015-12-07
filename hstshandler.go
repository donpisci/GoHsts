/*
GoHsts provides a simple middleware inplementation that adds the 
Strict-Transport-Security header to each response.

It also contains a number of customisable settings to allow for different options.

GoHsts is implements the http.Handler interface so that it will work with net/http 
*/

package GoHsts

import (
	"fmt"
	"net/http"
)

//HstsHandler provides HSTS middleware for HTTP requests
type HstsHandler struct {
	MaxAge            int
	IncludeSubdomains bool
}

//NewHstsHandler is the constructor for the HSTSHandler struct
//and sets default values for each of its fields
func NewHstsHandler() HstsHandler {
	return HstsHandler{MaxAge: 31536000, IncludeSubdomains: true}
}

const hstsHeaderName = "Strict-Transport-Security"

//HstsHandlerFunc adds a HSTS header to the HTTP response
func (h *HstsHandler) HstsHandlerFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Is there a header already there
		if header := w.Header().Get(hstsHeaderName); header != "" {
			w.Header().Set(hstsHeaderName, createHeaderValue(h.MaxAge, h.IncludeSubdomains))
		} else {
			w.Header().Add(hstsHeaderName, createHeaderValue(h.MaxAge, h.IncludeSubdomains))
		}

		complete := make(chan bool)
		go func() {
			if next != nil {
				next.ServeHTTP(w, r)
			}
			complete <- true
		}()
		<-complete
	})
}

//createHeaderValue sets the value to use in the HSTS header
func createHeaderValue(age int, includeSubDomains bool) string {
	content := fmt.Sprintf("max-age=%d", age)

	if includeSubDomains {
		content = fmt.Sprintf("%s; includeSubDomains", content)
	}
	return content
}
