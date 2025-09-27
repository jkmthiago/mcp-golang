package mcp

import (
	"mcp/tools"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func NewMcpServer()*server.MCPServer {
	mcps := server.NewMCPServer(
		"Fibralink LIA MCP API Server",
		"0.1.0",
		server.WithToolCapabilities(true),
		server.WithLogging(),
	)

	mcps.AddTool(
		mcp.NewTool("calculator",
			mcp.WithDescription("Do a basic aritmetic operation by passing 2 numbers and the operation to be done and returns the result of the operation"),
			mcp.WithNumber("x", mcp.Required(), mcp.Description("The first number for the operation")),
			mcp.WithNumber("y", mcp.Required(), mcp.Description("The second number for the operation")),
			mcp.WithString("operation", mcp.Enum("add", "subtract", "multiply", "divide"), mcp.Required(), mcp.Description("")),
		),
		tools.HandleCalculator,
	)

	return mcps
}
