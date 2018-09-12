package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jcasmer/GoBootcamp/memory_db/db"
)

var (
	//ErrNotFound es retornado cuando no existe el indice o el item en la bd
	ErrRequired          = errors.New("is required")
	ErrOwnerRequired     = errors.New("Owner is required")
	ErrIdArticleRequired = errors.New("IdArticle is required")
	ErrQuantity          = errors.New("Quantity must be greater than 0")
)

// Articles struct (Model)
type Articles struct {
	Id       string `json:"id"`
	Quantity int    `json:"quantity"`
}

// Car struct
type Carts struct {
	Id      string     `json:"id"`
	Owner   string     `json:"owner"`
	Article []Articles `json:"articles"`
}

type CartsArticles struct {
	id       string              `json:"id"`
	articles map[string]Articles `json:"articles"`
}

type Service struct {
	dataBase db.DbInter
}

var arcticles []Articles
var cartsA []Carts
var dbName = "db.json"

func NewService(db db.DbInter) *Service {

	s := Service{dataBase: db}
	return &s
}

func (c Carts) ValidateCart() error {
	//cart validations
	if c.Owner == "" {
		return ErrOwnerRequired
	}
	_, err := strconv.Atoi(c.Owner)
	if err == nil {
		return errors.New("Owner must be string")
	}
	return nil
}

func (a Articles) ValidateArticle() error {
	//article validations
	fmt.Println(a)
	if a.Id == "" {
		return ErrIdArticleRequired
	}
	_, err := strconv.Atoi(a.Id)
	if err != nil {
		return errors.New("id article must be int")
	}
	if a.Quantity <= 0 {
		return ErrQuantity
	}
	// _, err = strconv.Atoi(a.Quantity)
	// if err != nil {
	// 	return errors.New("quantity must be int")
	// }
	return nil
}

// Get all carts
func (s *Service) getCarts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(cartsA)

}

// Create single cart
func (s *Service) createCart(w http.ResponseWriter, r *http.Request) {
	// create a sinle cart without articles
	w.Header().Set("Content-Type", "application/json")

	var cart Carts
	err := json.NewDecoder(r.Body).Decode(&cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = cart.ValidateCart()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cart.Id = strconv.Itoa(rand.Intn(100000000))
	// cartsA = append(cartsA, cart)

	value, _ := json.Marshal(cart)

	resul := s.dataBase.CreateWithIndex(cart.Id, string(value))
	if resul != nil {
		http.Error(w, resul.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(cart)
	// http. .StatusCreated

}

func (s *Service) getCart(w http.ResponseWriter, r *http.Request) {
	// retrive a specific cart with its articles
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, "id is required", http.StatusBadRequest)
		return
	}
	var cart Carts
	car, erro := s.dataBase.Retrieve(params["id"])
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusNotFound)
		return
	}

	err := json.Unmarshal([]byte(car), &cart)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(cart)

}

func (s *Service) addArticles(w http.ResponseWriter, r *http.Request) {
	// add article to specific cart
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, "id cart is required", http.StatusBadRequest)
		return
	}

	var article Articles
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = article.ValidateArticle()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// value, _ := json.Marshal(cart)

	var cart Carts
	car, erro := s.dataBase.Retrieve(params["id"])
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusNotFound)
		return
	}

	_ = json.Unmarshal([]byte(car), &cart)
	// validate the article to add not exists
	for _, item := range cart.Article {
		if item.Id == article.Id {
			http.Error(w, "The article to add already exist", http.StatusNotFound)
			return
		}
	}

	cart.Article = append(cart.Article, article)

	value, _ := json.Marshal(cart)

	resul := s.dataBase.Update(cart.Id, string(value))
	if resul != nil {
		http.Error(w, resul.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	// json.NewEncoder(w).Encode(cart)
	// http. .StatusCreated

}
