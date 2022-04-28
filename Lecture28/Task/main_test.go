package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestEnds(t *testing.T) {
	handle := HandlerNewsMarshal("localhost:9000/hello")

	mockReq := httptest.NewRequest(http.MethodGet, "localhost:9000/hello", nil)
	responseRecoder := httptest.NewRecorder()
	handle(responseRecoder, mockReq)

	actual := responseRecoder.Body.String()
	expect := string(HNmock())

	if reflect.DeepEqual(actual, expect) == false {
		t.Errorf("Not expected")

	}

}
