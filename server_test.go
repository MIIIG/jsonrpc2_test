package main

import "testing"

func TestHandler(t *testing.T) {
	expectedHandlerResult := "somthng"
	if actualHandlerResult := Handler(mock1, mock2); actualHandlerResult != expectedHandlerResult {
		t.Errorf("expected %s; got: %s", expectedHandlerResult, &actualHandlerResult)
	}
}
