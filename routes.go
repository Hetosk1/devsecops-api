package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is running"))
	})

	r.Post("/api/v1/auth/register", RegisterUser)
	r.Post("/api/v1/auth/login", LoginUser)

	r.Get("/api/v1/user/me", GetMe)
	r.Get("/api/v1/users", GetAllUsers)

	r.Route("/api/v1/products", func(r chi.Router) {
		r.Get("/", GetAllProducts)
		r.Post("/", CreateProduct)
		r.Get("/{id}", GetProductByID)
		r.Put("/{id}", UpdateProduct)
		r.Delete("/{id}", DeleteProduct)
	})

	r.Route("/api/v1/categories", func(r chi.Router) {
		r.Get("/", GetAllCategories)
		r.Post("/", CreateCategory)
	})

	r.Route("/api/v1/cart", func(r chi.Router) {
		r.Get("/", GetCart)
		r.Post("/", AddToCart)
		r.Put("/{itemId}", UpdateCartItem)
		r.Delete("/{itemId}", RemoveCartItem)
	})

	r.Route("/api/v1/orders", func(r chi.Router) {
		r.Post("/", PlaceOrder)
		r.Get("/my", GetMyOrders)
		r.Get("/", GetAllOrders)
		r.Put("/{id}/status", UpdateOrderStatus)
	})

	r.Route("/api/v1/products/{id}/reviews", func(r chi.Router) {
		r.Post("/", AddReview)
		r.Get("/", GetProductReviews)
	})

	return r
}
