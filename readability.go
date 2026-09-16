// Package readability is a Go package that find the main readable
// content from a HTML page. It works by removing clutter like buttons,
// ads, background images, script, etc.
//
// This package is based from Readability.js by Mozilla, and written line
// by line to make sure it looks and works as similar as possible. This
// way, hopefully all web page that can be parsed by Readability.js
// are parse-able by go-readability as well.
package readability

import (
	"io"
	nurl "net/url"

	"golang.org/x/net/html"
)

// FromReader parses an `io.Reader` and returns the readable content. It's the wrapper
// or `Parser.Parse()` and useful if you only want to use the default parser.
func FromReader(input io.Reader, pageURL *nurl.URL) (Article, error) {
	parser := NewParser()
	return parser.Parse(input, pageURL)
}

// FromDocument parses an document and returns the readable content. It's the wrapper
// or `Parser.ParseDocument()` and useful if you only want to use the default parser.
func FromDocument(doc *html.Node, pageURL *nurl.URL) (Article, error) {
	parser := NewParser()
	return parser.ParseDocument(doc, pageURL)
}

// CheckDocument checks whether the document is readable without parsing the whole thing.
// It's the wrapper for `Parser.CheckDocument()` and useful if you only use the default
// parser.
func CheckDocument(doc *html.Node) bool {
	parser := NewParser()
	return parser.CheckDocument(doc)
}
