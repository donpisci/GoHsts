package GoHsts

import (
	"net/http/httptest"
	"testing"
	"net/http"
)

//TestHstsHandler tests that the HSTS header has been added
//to the response with the default values
func TestHstsHandler(t *testing.T) {
	w := httptest.NewRecorder()
	
	r, err := http.NewRequest("GET", "/", nil)
	
	if err != nil {
		t.Fatalf("Error constructing test HTTP request [%s]", err)
	}
	
	handler := NewHstsHandler()
	
	handler.HstsHandlerFunc(nil).ServeHTTP(w, r)
	
	if header := w.HeaderMap.Get(hstsHeaderName); header == "" {
		t.Error("HSTS header not present")
	} else {
		t.Log(header)
	}
	
}

//TestHstsHandler_SetMaxAge tests that the max age value on the 
//HSTS header is set correctly to a custom value
func TestHstsHandler_SetMaxAge(t *testing.T) {
	//Arrange
	expected := "max-age=1; includeSubDomains"
	
	w := httptest.NewRecorder()
	
	r, err := http.NewRequest("GET", "/", nil)
	
	if err != nil {
		t.Fatalf("Error constructing test HTTP request [%s]", err)
	}
	
	handler := NewHstsHandler()
	handler.MaxAge = 1
	
	//Act
	handler.HstsHandlerFunc(nil).ServeHTTP(w, r)
	
	header := w.HeaderMap.Get(hstsHeaderName)
	
	//Assert
	if header == "" {
		t.Error("HSTS header not present")
	}
	
	if header != expected {
		t.Error("MaxAge is not set correctly")
	}
	t.Log(header)
}

//TestHstsHandler_ExcludeSubdomains tests that the includeSubdomains
//part of the header value os not included
func TestHstsHandler_ExcludeSubdomains(t *testing.T) {
	//Arrange
	expected := "max-age=31536000"
	
	w := httptest.NewRecorder()
	
	r, err := http.NewRequest("GET", "/", nil)
	
	if err != nil {
		t.Fatalf("Error constructing test HTTP request [%s]", err)
	}
	
	handler := NewHstsHandler()
	handler.IncludeSubdomains = false
	
	//Act
	handler.HstsHandlerFunc(nil).ServeHTTP(w, r)
	
	//Assert
	header := w.HeaderMap.Get(hstsHeaderName)
	
	if header == "" {
		t.Error("HSTS header not present")
	}
	
	if header != expected {
		t.Error("HSTS header is not correctly formed")
	}
	t.Log(header)
}

//TestHstsHandler_SetMaxAgeAndExcludeSubdomains tests that both the MaxAge
//and IncludeSubdomains properties can be set 
func TestHstsHandler_SetMaxAgeAndExcludeSubdomains(t *testing.T) {
	//Arrange
	expected := "max-age=1"
	
	w := httptest.NewRecorder()
	
	r, err := http.NewRequest("GET", "/", nil)
	
	if err != nil {
		t.Fatalf("Error constructing test HTTP request [%s]", err)
	}
	
	handler := NewHstsHandler()
	handler.MaxAge = 1
	handler.IncludeSubdomains = false
	
	//Act
	handler.HstsHandlerFunc(nil).ServeHTTP(w, r)
	
	//Assert
	header := w.HeaderMap.Get(hstsHeaderName)
	
	if header == "" {
		t.Error("HSTS header not present")
	}
	
	if header != expected {
		t.Error("HSTS header is not correctly formed")
	}
	t.Log(header)
}

func TestHstsHandler_NewHstsHandler(t *testing.T) {
	//Arrange
	expectedMaxAge := 31536000
	expectedIncludeSubdomain := true
	
	//Act
	h := NewHstsHandler()
	
	//Assert
	if h.MaxAge != expectedMaxAge {
		t.Error("MaxAge value is incorrect")
	}
	
	if h.IncludeSubdomains != expectedIncludeSubdomain {
		t.Error("IncludeSubdomain value is incorrect")
	}
}