package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"EffectIndex",
		"GET",
		"/effects",
		EffectIndex,
	},
	Route{
		"EffectCreate",
		"POST",
		"/effects",
		EffectCreate,
	},
	Route{
		"EffectShow",
		"GET",
		"/effects/{effectId}",
		EffectShow,
	},
}
