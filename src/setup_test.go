package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	. "gopkg.in/check.v1"
)

type MySuite struct{}

var _ = Suite(&MySuite{})
var r = GetRouter()

func PerformRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func (s *MySuite) SetUpSuite(c *C) {
	w := PerformRequest(r, http.MethodDelete, "/users/all", nil)
	c.Check(w.Code, Equals, http.StatusOK)

	user := []byte(`{"username": "Alex", "password": "123", "isAdmin": "true"}`)
	w = PerformRequest(r, http.MethodPost, "/users", user)
	c.Check(w.Code, Equals, http.StatusOK)

	user = []byte(`{"username": "UserToDelete", "password": "4711", "isAdmin": "false"}`)
	w = PerformRequest(r, http.MethodPost, "/users", user)
	c.Check(w.Code, Equals, http.StatusOK)

}
