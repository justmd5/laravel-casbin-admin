package http

import "net/http"

var http2pushHeaderKey = http.CanonicalHeaderKey("http2-push")
var TrailerHeaderKey = http.CanonicalHeaderKey("trailer")
