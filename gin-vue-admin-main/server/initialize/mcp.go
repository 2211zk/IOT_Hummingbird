package initialize

// "github.com/mark3labs/mcp-go/server" // 临时注释以解决Go版本兼容性问题

func McpRun() interface{} { // 临时修改返回类型
	// config := global.GVA_CONFIG.MCP // 临时注释

	// s := server.NewMCPServer( // 临时注释
	// 	config.Name,
	// 	config.Version,
	// )

	// // global.GVA_MCP_SERVER = s // 临时注释以解决Go版本兼容性问题

	// mcpTool.RegisterAllTools(s) // 临时注释

	// return server.NewSSEServer(s, // 临时注释
	// 	server.WithSSEEndpoint(config.SSEPath),
	// 	server.WithMessageEndpoint(config.MessagePath),
	// 	server.WithBaseURL(config.UrlPrefix))

	// 临时返回nil
	return nil
}
