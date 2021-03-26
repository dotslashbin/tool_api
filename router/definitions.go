package router

import (
	"encoding/json"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"index",
		"GET",
		"/",
		testRoute,
	},
}

func testRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	test := []string{"this", "is", "testing"}

	json.NewEncoder(w).Encode(test)
}
