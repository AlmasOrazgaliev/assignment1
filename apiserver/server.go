package apiserver

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/AlmasOrazgaliev/assignment1/controller"
	"github.com/AlmasOrazgaliev/assignment1/model"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
	"strconv"
)

type server struct {
	router       *mux.Router
	controller   *controller.Controller
	sessionStore sessions.Store
}

func newServer(db *sql.DB, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		controller:   controller.New(db),
		sessionStore: sessionStore,
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/items/", s.handleItems())
	s.router.HandleFunc("/signin", s.handleSignIn())
	s.router.HandleFunc("/signup", s.handleSignUp())
	s.router.HandleFunc("/createItem", s.handleCreateItem()) //
	s.router.HandleFunc("/by_price/", s.handleSearchByPrice())
	s.router.HandleFunc("/search/", s.handleSearchByName())
	s.router.HandleFunc("/by_rating/", s.handleSearchByRating())
	s.router.HandleFunc("/admin_mode/", s.handleAdminMode())
	s.router.HandleFunc("/items/{id:[0-9]+}", s.handleItemsId())
	s.router.HandleFunc("/items/{id:[0-9]+}/getRating/", s.handleGetRating()) //
	//s.router.HandleFunc("/admin_mode/")
}

func (s *server) handleItems() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := s.controller.AllItems()
		response(w, http.StatusOK, items)
	}
}

func (s *server) handleSignIn() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := model.User{}
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			errResponse(w, http.StatusBadRequest, err)
		}
		res, err := s.controller.FindUser(&u)
		if res == nil {
			errResponse(w, http.StatusNotFound, errors.New("incorrect email or password"))
		}
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
		}
		response(w, http.StatusFound, nil)
	}
}

func (s *server) handleSignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		u := model.User{}
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			errResponse(w, http.StatusBadRequest, err)
		}
		err = s.controller.CreateUser(&u)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
		} else {
			response(w, http.StatusCreated, nil)
		}
	}
}

func (s *server) handleCreateItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		item := model.Item{}
		err := json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			errResponse(w, http.StatusBadRequest, err)
		}
		err = s.controller.CreateItem(&item)
		if err != nil {
			errResponse(w, http.StatusInternalServerError, err)
		} else {
			response(w, http.StatusCreated, nil)
		}

	}
}

func (s *server) handleSearchByPrice() http.HandlerFunc {
	type minMax struct {
		min int `json:"min"`
		max int `json:"max"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		mm := minMax{}
		err := json.NewDecoder(r.Body).Decode(&mm)
		if err != nil {
			errResponse(w, http.StatusBadRequest, err)
		}
		items := s.controller.SearchByPrice(mm.min, mm.max)
		if items != nil {
			response(w, http.StatusFound, items)
		} else {
			errResponse(w, http.StatusNotFound, errors.New("no such items"))
		}
	}
}

func (s *server) handleSearchByRating() http.HandlerFunc {
	type minMax struct {
		min int `json:"min"`
		max int `json:"max"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		mm := minMax{}
		err := json.NewDecoder(r.Body).Decode(&mm)
		if err != nil {
			errResponse(w, http.StatusBadRequest, err)
		}
		items := s.controller.SearchByRating(mm.min, mm.max)
		if items != nil {
			response(w, http.StatusFound, items)
		} else {
			errResponse(w, http.StatusNotFound, errors.New("no such items"))
		}
	}
}

func (s *server) handleSearchByName() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html, err := template.ParseFiles("templates/home.html", "templates/main.html")
		if err != nil {
			panic(err)
		}
		err = r.ParseForm()
		if err != nil {
			panic(err)
		}
		err = r.ParseForm()
		if err != nil {
			panic(err)
		}
		name := r.FormValue("search")
		items := s.controller.SearchByName(name)
		html.ExecuteTemplate(w, "main", items)
	}
}

func (s *server) handleItemsId() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		html, err := template.ParseFiles("templates/item.html", "templates/home.html")
		if err != nil {
			panic(err)
		}
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err)
		}
		html.ExecuteTemplate(w, "item", s.controller.GetById(id))
	}
}
func (s *server) handleGetRating() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			panic(err)
		}
		item := s.controller.GetById(id)
		rating, _ := strconv.Atoi(r.FormValue("rating"))
		item.Rating += float64(rating)
		s.controller.UpdateItem(&item)
		http.Redirect(w, r, "/home/", http.StatusFound)
	}
}

func (s *server) handleAdminMode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		items := s.controller.NotModeratedItems()
		html, err := template.ParseFiles("templates/admin.html")
		if err != nil {
			panic(err)
		}
		html.ExecuteTemplate(w, "admin.html", items)
	}
}

func errResponse(w http.ResponseWriter, code int, err error) {
	response(w, code, map[string]string{"error": err.Error()})
}

func response(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}

}
