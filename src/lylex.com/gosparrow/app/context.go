package app

import ()

// Context represents the app context of the application instance
type Context struct {
	Name      string
	Version   string
	StartedAt string
}

// Assimilate used to extended Context content.
func (ctx *Context) Assimilate() *Context {
	return ctx
}

// gContext is a global Context object
var gContext *Context

// GetContext returns the global Context object
func GetContext() *Context {
	return gContext
}

// SetContext returns the global Context object
func SetContext(c *Context) {
	gContext = c
}
