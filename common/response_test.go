package common

import (
	"errors"
	"testing"
)

// When a request fails before a response is received (e.g. timeout / connection
// reset), resp.RawResponse is nil. MCPToolResponse previously dereferenced it in
// the error branch and panicked the tool handler.
func TestMCPToolResponseNilRawResponseDoesNotPanic(t *testing.T) {
	resp := &APIResponse{RawResponse: nil}
	result, err := MCPToolResponse(resp, []int{200}, errors.New("dial tcp: i/o timeout"))
	if err == nil {
		t.Fatal("expected the transport error to be propagated")
	}
	if result == nil {
		t.Fatal("expected a tool result describing the error, got nil")
	}
}
