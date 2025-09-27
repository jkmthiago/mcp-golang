package tools

import (
	"context"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
)

func HandleCalculator(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	log.Println("Handle Calculations")
	
	op, err := req.RequireString("operation")
	if err != nil {
		return nil, err
	}

	x, err := req.RequireFloat("x")
	if err != nil {
		return nil, err
	}

	y, err := req.RequireFloat("y")
	if err != nil {
		return nil, err
	}

	switch op {
	case "add":
		return mcp.NewToolResultText(fmt.Sprintf("%.7f", x + y)), nil
	case "subtract":
		return mcp.NewToolResultText(fmt.Sprintf("%.7f", x - y)), nil
	case "multiply":
		return mcp.NewToolResultText(fmt.Sprintf("%.7f", x * y)), nil
	case "divide":
		if y == 0 {
			return mcp.NewToolResultError("cannot divide by zero"), nil
		}
		return mcp.NewToolResultText(fmt.Sprintf("%.7f", x / y)), nil
	}

	return mcp.NewToolResultError("invalid operation"), nil
}
