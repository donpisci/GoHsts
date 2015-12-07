[![GoDoc](https://godoc.org/github.com/donpisci/GoHsts?status.svg)](https://godoc.org/github.com/donpisci/GoHsts)

# GoHsts Middleware

GoHsts is a simple middleware implementation that appends the HTTP Secure Transport Policy (HSTS) header to every response.
Once the header has been added and a supporting browser has received the response, the browser will prevent any 
communications from being sent over HTTP to the specified domain and instead route any subsequent requests over HTTPS.

More information about HSTS can be found [here](http://tools.ietf.org/html/rfc6797)

This middleware handler uses implements the Go http.Handler interface

## Usage

```
go get github.com/donpisci/GoHsts
```

To instatiate the hstsHandler, use the provided constructor
```
hstsMiddleware := NewHstsHandler()
```
Using this constructor will set the structs fields to default values

There are two fields on the hstsHandler struct; MaxAge (int) and IncludeSubDomains (bool).
The MaxAge field determines how long the header is valid for in seconds and by default is set to 31536000 (1 year)
The IncludeSubDomains field does what it says on the tin. By default, this is set to true and is the recommended setting.

Once the values have been set, the HSTS middleware can be used like standard middleware:
```
http.ListenAndServe("/", hstsMiddleware.HstsHandlerFunc(otherMiddleware))
```
