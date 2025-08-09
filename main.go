package main

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

type AddParams struct {
	A string `json:"a" jsonschema:"the first number to add"`
	B string `json:"b" jsonschema:"the second number to add"`
}

func Add(ctx context.Context, cc *mcp.ServerSession, params *mcp.CallToolParamsFor[AddParams]) (*mcp.CallToolResultFor[any], error) {
	a, err := strconv.Atoi(params.Arguments.A)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("invalid number: %s", params.Arguments.A)}},
		}, nil
	}
	b, err := strconv.Atoi(params.Arguments.B)
	if err != nil {
		return &mcp.CallToolResultFor[any]{
			IsError: true,
			Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("invalid number: %s", params.Arguments.B)}},
		}, nil
	}

	sum := a + b
	return &mcp.CallToolResultFor[any]{
		Content: []mcp.Content{&mcp.TextContent{Text: fmt.Sprintf("%d", sum)}},
	}, nil
}

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "adder", Version: "v1.0.0"}, nil)

	mcp.AddTool(server, &mcp.Tool{Name: "add", Description: "adds two numbers"}, Add)

	if err := server.Run(context.Background(), mcp.NewStdioTransport()); err != nil {
		log.Fatal(err)
	}
}
