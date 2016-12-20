package main

import "net/http"
import (
	"fmt"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	// Mock Group is for later usage. Since we have only saleschannel, this is the only mock group
	Route{
		"MockGroup",
		"GET",
		"/",
		Index,
	},

}

func Index(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Welcome!\n")
}

