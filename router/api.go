package router

import "github.com/gorilla/mux"

func API() *mux.Router {
	m := mux.NewRouter()
	m.Path("/posts").Methods("GET").Name(Posts)
	m.Path("/posts/{ID:.+}").Methods("GET").Name(Post)
	return m
}
