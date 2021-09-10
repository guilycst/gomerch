package controllers

import (
	"fmt"
	"gomerch/models"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", models.FindAll())
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", models.FindAll())
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		product, err := getProductFormData(w, r)
		if err != nil {
			panic(err.Error())
		}

		models.AddNewProduct(*product)
		http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	parsedProductId, err := getIdFromQuery(w, r)
	if err != nil {
		panic(err.Error())
	}
	models.DeleteProduct(parsedProductId)
	http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	parsedProductId, err := getIdFromQuery(w, r)
	if err != nil {
		panic(err.Error())
	}
	temp.ExecuteTemplate(w, "Edit", models.FindProduct(parsedProductId))
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		product, err := getProductFormData(w, r)
		if err != nil {
			panic(err.Error())
		}

		models.UpdateProduct(*product)
		http.Redirect(w, r, "/home", http.StatusPermanentRedirect)
	}
}

func getIdFromQuery(w http.ResponseWriter, r *http.Request) (uint, error) {
	productId := r.URL.Query().Get("id")
	parsedProductId, err := strconv.ParseUint(productId, 10, 32)
	if err != nil {
		log.Println("Failed to parse id:", err)
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return 0, err
	}
	return uint(parsedProductId), err
}

func getProductFormData(w http.ResponseWriter, r *http.Request) (*models.Product, error) {
	id := r.FormValue("id")
	name := r.FormValue("name")
	description := r.FormValue("description")
	price := r.FormValue("price")
	amount := r.FormValue("amount")

	hasId := len(id) > 0
	var parsedId uint
	if hasId {
		formId, err := strconv.ParseUint(id, 10, 32)
		if err != nil {
			log.Println(fmt.Sprintf("Failed to parse id %d", formId), err)
			http.Error(w, "Invalid id", http.StatusBadRequest)
			return nil, err
		}
		parsedId = uint(formId)
	}

	parsedPrice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Failed to parse price:", err)
		http.Error(w, "Invalid price", http.StatusBadRequest)
		return nil, err
	}

	parsedAmount, err := strconv.Atoi(amount)
	if err != nil {
		log.Println("Failed to parse amount:", err)
		http.Error(w, "Invalid amount", http.StatusBadRequest)
		return nil, err
	}

	product := models.Product{
		Model:       gorm.Model{ID: parsedId},
		Name:        name,
		Description: description,
		Price:       parsedPrice,
		Amount:      parsedAmount,
	}

	return &product, nil
}
