package apiserver

import (
	"database/sql"
	"github.com/AlmasOrazgaliev/assignment1/controller"
	"github.com/AlmasOrazgaliev/assignment1/model"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
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
	s.router.HandleFunc("/registration/", s.handleRegistration())
	s.router.HandleFunc("/create/", s.handleUserCreate())
	s.router.HandleFunc("/seller/", s.handleSellerCreate())
	s.router.HandleFunc("/save_user/", s.handleSaveUser())
	s.router.HandleFunc("/publish/", s.handlePublishItem())
	s.router.HandleFunc("/save_item/", s.handleSaveItem())
	s.router.HandleFunc("/by_price/", s.handleSearchByPrice())
	s.router.HandleFunc("/search/", s.handleSearchByName())
	s.router.HandleFunc("/by_rating/", s.handleSearchByRating())
	s.router.HandleFunc("/admin_mode")
}

func (s *server) handleHome() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			res, _ := s.controller.FindUser(user.Email, user.Password)
			if res == nil {
				http.Redirect(w, r, "/registration/", http.StatusFound)
			}
		}
		items := s.controller.AllItems()
		html.ExecuteTemplate(w, "main", items)
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
		if r.FormValue("is_seller") == "" {
			user.IsSeller = false
		} else {
			user.IsSeller = true
		}
		err = s.controller.CreateUser(&user)
		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/home/", http.StatusFound)
	}
}

func (s *server) handlePublishItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html, err := template.ParseFiles("templates/publish.html")
		if err != nil {
			panic(err)
		}
		err = html.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	}
}

func (s *server) handleSaveItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		price, err := strconv.Atoi(r.FormValue("price")) ////////////////////////////////////////////////////////////////
		item := model.Item{
			Name:        r.FormValue("name"),
			Price:       price,
			Description: r.FormValue("description"),
		}
		err = s.controller.CreateItem(&item)
		http.Redirect(w, r, "/home/", http.StatusFound)
	}
}

func (s *server) handleSearchByPrice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html, err := template.ParseFiles("templates/home.html", "templates/main.html")
		if err != nil {
			panic(err)
		}
		err = r.ParseForm()
		if err != nil {
			panic(err)
		}
		min, err := strconv.Atoi(r.FormValue("min"))
		if err != nil {
			panic(err)
		}
		max, err := strconv.Atoi(r.FormValue("max"))
		items := s.controller.SearchByPrice(min, max)
		html.ExecuteTemplate(w, "main", items)
	}
}

func (s *server) handleSearchByRating() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		html, err := template.ParseFiles("templates/home.html", "templates/main.html")
		if err != nil {
			panic(err)
		}
		err = r.ParseForm()
		if err != nil {
			panic(err)
		}
		min, err := strconv.ParseFloat(r.FormValue("min"), 64)
		if err != nil {
			panic(err)
		}
		max, err := strconv.ParseFloat(r.FormValue("max"), 64)
		items := s.controller.SearchByRating(min, max)
		html.ExecuteTemplate(w, "main", items)
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

func (s *server) handleAdminMode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
