package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"testing"

	. "../model"
	. "../utils"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func TestGet(t *testing.T) { TestingT(t) }

func (s *MySuite) TestGetHelloMessage(c *C) {
	w := PerformRequest(r, http.MethodGet, "/", nil)
	c.Check(w.Code, Equals, http.StatusOK)
	var msg Message
	err := json.Unmarshal([]byte(w.Body.String()), &msg)
	c.Assert(err, Equals, nil)
	c.Check(msg.Message, Equals, "Hello! I am the go app running for you :-)")
}

func (s *MySuite) TestGetByUsername(c *C) {
	w := PerformRequest(r, http.MethodGet, "/users?attribute=username&value=Alex", nil)
	c.Check(w.Code, Equals, http.StatusOK)
	var results []interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &results)
	c.Assert(err, Equals, nil)
	c.Check(len(results), Equals, 1)
	c.Assert(mustMatchRegex(w.Body.String(), "username.*:.*Alex"), Equals, true)
	c.Assert(mustMatchRegex(w.Body.String(), "password.*:.*123"), Equals, true)
	c.Assert(mustMatchRegex(w.Body.String(), "isadmin.*:.*true"), Equals, true)
}

func (s *MySuite) TestGetByPassword(c *C) {
	w := PerformRequest(r, http.MethodGet, "/users?attribute=password&value=123", nil)
	c.Check(w.Code, Equals, http.StatusOK)
	var results []interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &results)
	c.Assert(err, Equals, nil)
	c.Check(len(results), Equals, 1)
	c.Assert(mustMatchRegex(w.Body.String(), "username.*:.*Alex"), Equals, true)
	c.Assert(mustMatchRegex(w.Body.String(), "password.*:.*123"), Equals, true)
	c.Assert(mustMatchRegex(w.Body.String(), "isadmin.*:.*true"), Equals, true)
}

func (s *MySuite) TestGetAll(c *C) {
	w := PerformRequest(r, http.MethodGet, "/users/all", nil)
	c.Check(w.Code, Equals, http.StatusOK)
	var receivedUsers []interface{}
	err := json.Unmarshal([]byte(w.Body.String()), &receivedUsers)
	c.Assert(err, Equals, nil)
	expectedNumberOfEntries := 2
	c.Check(len(receivedUsers), Equals, expectedNumberOfEntries)
}

func mustMatchRegex(input string, regex string) bool {
	re := regexp.MustCompile(regex)
	match := re.MatchString(input)
	if !match {
		LogIt(ERROR, "mustMatchRegex", fmt.Sprintf("regex '%s' didn't match in '%s'", regex, input))
	}
	return match
}
