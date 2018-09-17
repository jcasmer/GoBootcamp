package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jcasmer/GoBootcamp/memory_db/db"
	// "github.com/jcasmer/GoBootcamp/rest_api/handler"
)

// var dbName = "db.json"

func TestgetCart(t *testing.T) {

	t.Run("Open", func(tt *testing.T) {
		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		w := httptest.NewRecorder()

		// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
		// pass 'nil' as the third parameter.
		r, err := http.NewRequest("GET", "/carts/1", nil)
		if err != nil {
			t.Fatal(err)
		}
		dataBase, _ := db.OpenDB("db.json")
		service := NewService(dataBase)
		handler := http.HandlerFunc(service.getCart)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(w, r)

		// Check the status code is what we expect.
		fmt.Println(w.Code)
		if status := w.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		// Check the response body is what we expect.
		// expected := `{"Id": string}`
		// if rr.Body.String() != expected {
		// 	t.Errorf("handler returned unexpected body: got %v want %v",
		// 		rr.Body.String(), expected)
		// }
	})
}
