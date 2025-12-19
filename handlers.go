package main

import (
	"encoding/json"
	"net/http"
)

// ----------- AUTH -----------

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User registered successfully",
	})
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "User logged in successfully",
		"token":   "dummy-jwt-token",
	})
}

// ----------- USERS -----------

func GetMe(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"id":    "1",
		"name":  "Het",
		"email": "het@example.com",
	})
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]map[string]string{
		{"id": "1", "name": "Het"},
		{"id": "2", "name": "User2"},
	})
}

// ----------- PRODUCTS -----------

func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]map[string]interface{}{
		{"id": 1, "name": "iPhone", "price": 79999},
	})
}

func GetProductByID(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"id":    1,
		"name":  "iPhone",
		"price": 79999,
	})
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Product created",
	})
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Product updated",
	})
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Product deleted",
	})
}

// ----------- CATEGORIES -----------

func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]map[string]string{
		{"id": "1", "name": "Electronics"},
	})
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Category created",
	})
}

// ----------- CART -----------

func GetCart(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]interface{}{
		"items": []string{},
	})
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Item added to cart",
	})
}

func UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Cart item updated",
	})
}

func RemoveCartItem(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Cart item removed",
	})
}

// ----------- ORDERS -----------

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Order placed successfully",
	})
}

func GetMyOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]map[string]interface{}{
		{"id": 1, "status": "PLACED"},
	})
}

func GetAllOrders(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]map[string]interface{}{
		{"id": 1, "status": "SHIPPED"},
	})
}

func UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Order status updated",
	})
}

// ----------- REVIEWS -----------

func AddReview(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Review added",
	})
}

func GetProductReviews(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode([]map[string]interface{}{
		{"rating": 5, "comment": "Excellent product"},
	})
}

