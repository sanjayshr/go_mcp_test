package main

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestAdd(t *testing.T) {
	testCases := []struct {
		name       string
		args       AddParams
		wantResult string
		wantErr    bool
		wantErrMsg string
	}{
		{
			name:       "simple addition",
			args:       AddParams{A: "10", B: "20"},
			wantResult: "30",
			wantErr:    false,
		},
		{
			name:       "addition with negative number",
			args:       AddParams{A: "-10", B: "20"},
			wantResult: "10",
			wantErr:    false,
		},
		{
			name:       "addition with zero",
			args:       AddParams{A: "10", B: "0"},
			wantResult: "10",
			wantErr:    false,
		},
		{
			name:       "invalid first number",
			args:       AddParams{A: "abc", B: "20"},
			wantErr:    true,
			wantErrMsg: "invalid number: abc",
		},
		{
			name:       "invalid second number",
			args:       AddParams{A: "10", B: "xyz"},
			wantErr:    true,
			wantErrMsg: "invalid number: xyz",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			params := &mcp.CallToolParamsFor[AddParams]{
				Arguments: tc.args,
			}
			// The Add function doesn't use the ServerSession, so we can pass nil.
			result, err := Add(context.Background(), nil, params)
			if err != nil {
				t.Fatalf("Add() returned an unexpected error: %v", err)
			}

			if tc.wantErr {
				if !result.IsError {
					t.Fatalf("Add() did not return an error, but one was expected")
				}
				if len(result.Content) != 1 {
					t.Fatalf("expected 1 content item for error, got %d", len(result.Content))
				}
				textContent, ok := result.Content[0].(*mcp.TextContent)
				if !ok {
					t.Fatalf("expected TextContent for error, got %T", result.Content[0])
				}
				if textContent.Text != tc.wantErrMsg {
					t.Errorf("Add() returned error message %q, want %q", textContent.Text, tc.wantErrMsg)
				}
			} else {
				if result.IsError {
					t.Fatalf("Add() returned an error, but none was expected: %v", result.Content)
				}
				if len(result.Content) != 1 {
					t.Fatalf("expected 1 content item, got %d", len(result.Content))
				}
				textContent, ok := result.Content[0].(*mcp.TextContent)
				if !ok {
					t.Fatalf("expected TextContent, got %T", result.Content[0])
				}
				if textContent.Text != tc.wantResult {
					t.Errorf("Add() returned %q, want %q", textContent.Text, tc.wantResult)
				}
			}
		})
	}
}
