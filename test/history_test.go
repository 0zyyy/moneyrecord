package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHellow(t *testing.T) {
	assert.Equal(t, 123, 123, "IT'S MUST BE EQUAL")
}

func Teapot(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusTeapot)
}

func TestCreate(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080/api/v1/yyy", nil)
	res := httptest.NewRecorder()

	Teapot(res, req)

	if res.Code != http.StatusTeapot {
		t.Errorf("got status %d but wanted %d", res.Code, http.StatusTeapot)
	}
	fmt.Printf("status %d,%d", res.Code, http.StatusTeapot)

}
