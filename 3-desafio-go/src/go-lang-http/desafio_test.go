package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Greeting(t *testing.T) {
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
			t.Fatal(err)
	}

	res := httptest.NewRecorder()
	handler(res, req)

	exp := "<b>Code.education Rocks!</b>"
	act := res.Body.String()
	if exp != act {
			t.Fatalf("Expected %s, received %s", exp, act)
	}	
}
