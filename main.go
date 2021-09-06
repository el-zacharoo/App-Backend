package main

import (
	"fmt"
	"net/http"

	"github.com/el-zacharoo/go-101/handler"
	"github.com/el-zacharoo/go-101/store"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func main() {
	s := store.Store{}
	s.Connect()

	// chi
	r := chi.NewRouter()
	r.Use(middleware.Logger,
		cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: false,
			MaxAge:           300,
		},
		),
	)

	p := &handler.Person{
		Store: s,
	}
	r.Route("/people", func(r chi.Router) {
		r.Post("/", p.Create)
		r.Get("/{id}", p.Get)
		r.Put("/{id}", p.Update)
		r.Delete("/{id}", p.Delete)
	})

	o := &handler.Org{
		Store: s,
	}
	r.Route("/organisations", func(r chi.Router) {
		r.Post("/", o.Create)
		r.Get("/{id}", o.Get)
		r.Put("/{id}", o.Update)
		r.Delete("/{id}", o.Delete)
	})

	prd := &handler.Product{
		Store: s,
	}
	r.Route("/product", func(r chi.Router) {
		r.Post("/", prd.Create)
		r.Get("/{id}", prd.Get)
		r.Put("/{id}", prd.Update)
		r.Delete("/{id}", prd.Delete)
	})

	// start server
	if err := http.ListenAndServe(":8080", r); err != nil {
		fmt.Print(err)
	}
}

// func CORS(r *mux.Router) mux.MiddlewareFunc {
// 	return func(next http.Handler) http.Handler {
// 		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
// 			if req.Method == http.MethodOptions {
// 				w.Header().Set("Access-Control-Allow-Origin", "*")
// 				w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
// 				w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, X-CSRF-Token")
// 			}
// 			// allMethods, err := getAllMethodsForRoute(r, req)
// 			// if err == nil {
// 			// 	for _, v := range allMethods {
// 			// 		if v == http.MethodOptions {
// 			// 			w.Header().Set("Access-Control-Allow-Methzods", strings.Join(allMethods, ","))
// 			// 		}
// 			// 	}
// 			// }

// 			next.ServeHTTP(w, req)
// 		})
// 	}
// }
