package controllers

import (
	"pruebaTecnica/src/api/middlewares"
)

func (s *Server) initializeRoutes() {

	t := s.Router.StrictSlash(true).PathPrefix("/api").Subrouter()

	// Person Routes
	t.HandleFunc("/person", middlewares.SetMiddlewareJSON(s.GetPersons)).Methods("GET")
	t.HandleFunc("/person", middlewares.SetMiddlewareJSON(s.CreatePerson)).Methods("POST")
	t.HandleFunc("/person", middlewares.SetMiddlewareJSON(s.UpdatePerson)).Methods("PUT")
	t.HandleFunc("/person/byId/{id}", middlewares.SetMiddlewareJSON(s.DeleteById)).Methods("DELETE")
}
