package apiserver

import (
	"database/sql"
	"github.com/AlmasOrazgaliev/assignment1/controller"
	"github.com/AlmasOrazgaliev/assignment1/model"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"net/http"
)

type server struct {
	router     *mux.Router
	controller *controller.Controller
}

func newServer(db *sql.DB) *server {
	s := &server{
		router:     mux.NewRouter(),
		controller: controller.New(db),
	}
	s.configureRouter()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/home/", s.handleHome())
}

func (s *server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "home")
		html, err := template.ParseFiles("templates/home.html", "templates/main.html")
		if err != nil {
			panic(err)
		}
		err = r.ParseForm()
		if err != nil {
			panic(err)
		}
		user := model.User{
			Email:    r.FormValue("email"),
			Password: r.FormValue("password"),
		}
		if len(user.Password) != 0 && len(user.Password) != 0 {
			res, _ := s.controller.FindByEmail(user.Email)
			if res != nil {
				http.Redirect(w, r, "/registration/", http.StatusFound)
			}
		}
		items := s.controller.AllItems()
		html.ExecuteTemplate(w, "templates/main", items)
	}
}

func (s *server) handleRegistration() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html, err := template.ParseFiles("templates/reg.html")
		if err != nil {
			panic(err)
		}
		err = r.ParseForm()
		if err != nil {
			panic(err)
		}
		err = html.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	}
}

func (s *server) handleUserCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html, err := template.ParseFiles("templates/create.html")
		if err != nil {
			panic(err)
		}
		err = r.ParseForm()
		if err != nil {
			panic(err)
		}
		err = html.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	}
}

func (s *server) handleSellerCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html, err := template.ParseFiles("templates/seller.html")
		if err != nil {
			panic(err)
		}
		err = r.ParseForm()
		if err != nil {
			panic(err)
		}
		err = html.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	}
}

func (s *server) handleSaveUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		user := model.User{}
		user.Email = r.FormValue("email")
		user.Password = r.FormValue("password")
		user.IsSeller = r.FormValue("is_seller")
		s.controller.CreateUser(user)
		http.Redirect(w, r, "/home/", http.StatusFound)
	}
}
