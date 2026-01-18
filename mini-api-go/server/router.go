package server

import "net/http"

func (app *App) Get(path string, handler func(c *Context)) {
	app.mux.HandleFunc("GET "+path, func(rw http.ResponseWriter, r *http.Request) {
		c := &Context{RWriter: rw, Request: r, Ctx: r.Context()}
		handler(c)
	})
	app.handlerCount++
}

func (app *App) Post(path string, handler func(c *Context)) {
	app.mux.HandleFunc("POST "+path, func(rw http.ResponseWriter, r *http.Request) {
		c := &Context{RWriter: rw, Request: r, Ctx: r.Context()}
		handler(c)
	})
	app.handlerCount++
}

func (app *App) Put(path string, handler func(c *Context)) {
	app.mux.HandleFunc("PUT "+path, func(rw http.ResponseWriter, r *http.Request) {
		c := &Context{RWriter: rw, Request: r, Ctx: r.Context()}
		handler(c)
	})
	app.handlerCount++
}

func (app *App) Delete(path string, handler func(c *Context)) {
	app.mux.HandleFunc("DELETE "+path, func(rw http.ResponseWriter, r *http.Request) {
		c := &Context{RWriter: rw, Request: r, Ctx: r.Context()}
		handler(c)
	})
	app.handlerCount++
}
