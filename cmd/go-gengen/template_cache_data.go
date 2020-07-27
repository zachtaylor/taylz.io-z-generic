package main

// CacheTemplateData is the type consumed by CacheTemplate
type CacheTemplateData struct {
	// Package name
	Package string
	// Key type name
	Key string
	// Val type name
	Val string
	// Stdlib imports from std lib
	Stdlib []string
	// Remote imports
	Remote []string
}
