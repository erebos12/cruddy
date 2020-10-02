package test

import (
	"encoding/json"
	"net/http"
	"testing"

	. "../model"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func TestDelete(t *testing.T) { TestingT(t) }

func (s *MySuite) Test_When_Delete_Existing_User_thenExpect_OKWithSuccessMsg(c *C) {
	w := PerformRequest(r, http.MethodDelete, "/users?attribute=username&value=UserToDelete", nil)
	c.Check(w.Code, Equals, http.StatusOK)
	var msg Message
	err := json.Unmarshal([]byte(w.Body.String()), &msg)
	c.Assert(err, Equals, nil)
	c.Check(msg.Message, Equals, "Successfully deleted entity with attribute=username, value=UserToDelete")
	w = PerformRequest(r, http.MethodGet, "/users/all", nil)
	c.Check(w.Code, Equals, http.StatusOK)
	var receivedUsers []interface{}
	err = json.Unmarshal([]byte(w.Body.String()), &receivedUsers)
	c.Assert(err, Equals, nil)
	expectedRemainingEntries := 1
	c.Check(len(receivedUsers), Equals, expectedRemainingEntries)
}

func (s *MySuite) Test_When_Delete_NonExisting_User_thenExpect_NotFoundWithErrorMsg(c *C) {
	w := PerformRequest(r, http.MethodDelete, "/users?attribute=username&value=Frida", nil)
	c.Check(w.Code, Equals, http.StatusNotFound)
	var msg Message
	err := json.Unmarshal([]byte(w.Body.String()), &msg)
	c.Assert(err, Equals, nil)
	c.Check(msg.Message, Equals, "Cannot delete entity with attribute=username, value=Frida")
}

func (s *MySuite) Test_When_Delete_WithInavlidAttribute_thenExpect_NotFoundWithErrorMsg(c *C) {
	w := PerformRequest(r, http.MethodDelete, "/users?attribute=UNKNOWN&value=any", nil)
	c.Check(w.Code, Equals, http.StatusNotFound)
	var msg Message
	err := json.Unmarshal([]byte(w.Body.String()), &msg)
	c.Assert(err, Equals, nil)
	c.Check(msg.Message, Equals, "Cannot delete entity with attribute=UNKNOWN, value=any")
}
