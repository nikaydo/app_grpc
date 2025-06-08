package router

import (
	"main/internal/database"
	h "main/internal/handlers"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	auth "github.com/nikaydo/grpc-contract/gen/auth"
)

type Router struct {
	Handlers h.Handlers
	Auth     auth.AuthClient
}

func RouterInit(db database.Database, g auth.AuthClient) Router {
	return Router{Handlers: h.Handlers{Db: db, Auth: g}}
}

func (rt *Router) Router() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/user", func(r chi.Router) {
		r.Use(rt.Handlers.CheckJWT)
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "./web/user.html")
		})
		r.Get("/token", rt.Handlers.Token)
		r.Delete("/token", rt.Handlers.Token)
		r.Get("/tokens", rt.Handlers.GetTokens)
	})

	r.Post("/signup", rt.Handlers.SignUp)
	r.Post("/signin", rt.Handlers.SignIn)
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/index.html")
	})
	r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./web/register.html")
	})
	return r
}
