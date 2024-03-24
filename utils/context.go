package utils

type contextKey struct {
	name string
}

var DatabaseKey = &contextKey{name: "Database"}
var SessionManagerKey = &contextKey{name: "SessionManager"}
var TSDBKey = &contextKey{name: "TSDB"}
var LocationKey = &contextKey{name: "Location"}
