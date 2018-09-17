package handler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/jcasmer/GoBootcamp/memory_db/db"
	// "github.com/jcasmer/GoBootcamp/rest_api/handler"
)

// var dbName = "db.json"

// func TestGetCartFail(t *testing.T) {
// 	// get test fail
// 	t.Run("Open", func(tt *testing.T) {
// 		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
// 		w := httptest.NewRecorder()

// 		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
// 		// pass 'nil' as the third parameter.
// 		r, err := http.NewRequest("GET", "http://localhost:8002/carts/1", nil)

// 		if err != nil {
// 			t.Fatal(err)
// 			return
// 		}
// 		dataBase, _ := db.OpenDB("db.json")
// 		service := NewService(dataBase)
// 		// handler := http.HandlerFunc(service.getCart(w, r))
// 		service.GetCart(w, r)
// 		res := w.Result()
// 		_, err = ioutil.ReadAll(res.Body)
// 		if err != nil {
// 			t.Fatalf("could not read response: %v", err)
// 			return
// 		}

// 		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
// 		// directly and pass in our Request and ResponseRecorder.
// 		// handler.ServeHTTP(w, r)

// 		// Check the status code is what we expect.
// 		// fmt.Println(res.StatusCode)
// 		if status := res.StatusCode; status != http.StatusOK {
// 			t.Errorf("handler returned wrong status code: got %v want %v",
// 				status, http.StatusOK)
// 			return
// 		}
// 	})
// }

func TestGetCart(t *testing.T) {
	// get test fail
	t.Run("Open", func(tt *testing.T) {
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.

		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		r, err := http.NewRequest("GET", "carts/19727887", nil)
		r.Header.Set("Content-Type", "application/json")

		if err != nil {
			t.Fatal(err)
			return
		}
		// q := url.Values{}
		// q.Add("id", "19727887")
		// r.URL.Path = q.Encode()
		fmt.Println(r.URL.String())

		w := httptest.NewRecorder()
		w.Header().Set("Content-Type", "application/json")

		dataBase, _ := db.OpenDB("db.json")
		service := NewService(dataBase)
		service.GetCart(w, r)
		// handler := http.HandlerFunc(service.GetCart)
		// service.GetCart(w, r)
		res := w.Result()

		// fmt.Println(res)

		// _, err = ioutil.ReadAll(res.Body)
		// if err != nil {
		// 	t.Fatalf("could not read response: %v", err)
		// 	return
		// }

		// handler.ServeHTTP(w, r)

		// Check the status code is what we expect
		if status := res.StatusCode; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
			return
		}
		b, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("could not read response: %v", err)
		}
		d, err := strconv.Atoi(string(bytes.TrimSpace(b)))
		fmt.Println(w.Code, d)
	})
}
