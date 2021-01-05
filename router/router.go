package router

import (
	"fmt"
	"net/http"
)

type Router struct {
	handlers map[string]func(http.ResponseWriter, *http.Request)
}

func NewRouter() *Router {
	router := new(Router)                                                       // allocate Router struct
	router.handlers = make(map[string]func(http.ResponseWriter, *http.Request)) // allocate handlers map
	return router
}

func (s *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	f, ok := s.handlers[key(r.Method, r.URL.Path)]
	if !ok {
		ruhroh(w)
		return
	}
	f(w, r)
}

func (s *Router) GET(path string, f http.HandlerFunc) {
	s.handlers[key("GET", path)] = f
}

func (s *Router) POST(path string, f http.HandlerFunc) {
	s.handlers[key("POST", path)] = f
}

func ruhroh(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"error":"erroar"}`))
}

func key(method, path string) string {
	return fmt.Sprintf("%s:%s", method, path)
}
