package tools

import (
	"context"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
)

func HandleCalculator(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	
	op, err := req.RequireString("operation")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	
	x, err := req.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	
	y, err := req.RequireFloat("y")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	
	log.Println("Handle Calculations")
	
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
