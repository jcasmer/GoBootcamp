package handler

import (
	"encoding/json"
	"errors"
	"log"
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
	Router   *mux.Router
}

var arcticles []Articles
var cartsA []Carts
var dbName = "db.json"

func (s *Service) NewService() {

	s.dataBase, _ = db.OpenDB("db.json")
	s.Router = mux.NewRouter()
	s.initializeRoutes()
	// return &s
}

func (s *Service) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, s.Router))
	defer s.dataBase.Close("db.json")
}

func (s *Service) initializeRoutes() {
	s.Router.HandleFunc("/carts", s.CreateCart).Methods("POST")
	s.Router.HandleFunc("/carts/{id}", s.GetCart).Methods("GET")
	s.Router.HandleFunc("/carts/{id}", s.deleteCart).Methods("DELETE")
	s.Router.HandleFunc("/carts/{id}/items", s.AddArticles).Methods("POST")
	s.Router.HandleFunc("/carts/{id}/items", s.deleteAllArticles).Methods("DELETE")
	s.Router.HandleFunc("/carts/{id}/items/{idItem}", s.changeArticles).Methods("PUT")
	s.Router.HandleFunc("/carts/{id}/items/{idItem}", s.deleteArticle).Methods("DELETE")
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

func ResponseHttp(w http.ResponseWriter, status int, ob interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(ob)

}

// Get all carts
func (s *Service) getCarts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(cartsA)

}

// Create single cart
func (s *Service) CreateCart(w http.ResponseWriter, r *http.Request) {
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

	ResponseHttp(w, http.StatusCreated, cart)
}

func (s *Service) GetCart(w http.ResponseWriter, r *http.Request) {
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

	ResponseHttp(w, http.StatusOK, cart)

}

func (s *Service) AddArticles(w http.ResponseWriter, r *http.Request) {
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
	found := false
	// validate the article to add not exists
	for index, item := range cart.Article {
		if item.Id == article.Id {
			cart.Article[index].Quantity++
			found = true
		}
	}
	if !found {

		cart.Article = append(cart.Article, article)
	}

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

func (s *Service) changeArticles(w http.ResponseWriter, r *http.Request) {
	// change the quantity of a specific item in a cart
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, "id cart is required", http.StatusBadRequest)
		return
	}
	if params["idItem"] == "" {
		http.Error(w, "id item is required", http.StatusBadRequest)
		return
	}

	var article Articles
	article.Id = params["idItem"]
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
	for index, item := range cart.Article {
		if item.Id == article.Id {
			cart.Article[index].Quantity = article.Quantity

		}
	}
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

func (s *Service) deleteArticle(w http.ResponseWriter, r *http.Request) {
	// delete an item of a specific item in a cart
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, "id cart is required", http.StatusBadRequest)
		return
	}
	if params["idItem"] == "" {
		http.Error(w, "id item is required", http.StatusBadRequest)
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
	for index, item := range cart.Article {
		if item.Id == params["idItem"] {
			cart.Article = append(cart.Article[:int(index)], cart.Article[int(index)+1:]...)
			// delete(cart.Article, index)

		}
	}
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

func (s *Service) deleteAllArticles(w http.ResponseWriter, r *http.Request) {
	// delete all items of a specific  a cart
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, "id cart is required", http.StatusBadRequest)
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

	// clear articles
	cart.Article = nil

	value, _ := json.Marshal(cart)

	resul := s.dataBase.Update(cart.Id, string(value))
	if resul != nil {
		http.Error(w, resul.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	// json.NewEncoder(w).Encode(cart)
	// http. .StatusCreated|

}

func (s *Service) deleteCart(w http.ResponseWriter, r *http.Request) {
	// delete a specific  a cart
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	if params["id"] == "" {
		http.Error(w, "id cart is required", http.StatusBadRequest)
		return
	}
	// value, _ := json.Marshal(cart)

	resul := s.dataBase.Delete(params["id"])
	if resul != nil {
		http.Error(w, resul.Error(), http.StatusNotFound)
		return
	}
	_ = s.dataBase.Close("db.json")

	w.WriteHeader(http.StatusNoContent)
	// json.NewEncoder(w).Encode(cart)
	// http. .StatusCreated|

}
