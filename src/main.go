package main

import (
	"log"
	"mcp/mcp"

	"github.com/mark3labs/mcp-go/server"
	"github.com/mark3labs/mcp-go/util"
)

// var logger = log.New(os.Stdout, "MCP - ", log.LstdFlags)

func main() {
	mcps := mcp.NewMcpServer()
	logger := util.DefaultLogger()
	log.Println("Starting MCP Test Server")
	httpServer := server.NewStreamableHTTPServer(mcps, server.WithEndpointPath("/api/v1/mcp"), server.WithLogger(logger))
	if err := httpServer.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}