package client

import (
	"errors"
	// mcpClient "github.com/mark3labs/mcp-go/client" // 临时注释以解决Go版本兼容性问题
	// "github.com/mark3labs/mcp-go/mcp" // 临时注释以解决Go版本兼容性问题
)

func NewClient(baseUrl, name, version, serverName string) (interface{}, error) {
	// 临时注释所有MCP相关代码
	// client, err := mcpClient.NewSSEMCPClient(baseUrl)
	// if err != nil {
	// 	return nil, err
	// }

	// ctx := context.Background()

	// // 启动client
	// if err := client.Start(ctx); err != nil {
	// 	return nil, err
	// }

	// // 初始化
	// initRequest := mcp.InitializeRequest{}
	// initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	// initRequest.Params.ClientInfo = mcp.Implementation{
	// 	Name:    name,
	// 	Version: version,
	// }

	// result, err := client.Initialize(ctx, initRequest)
	// if err != nil {
	// 	return nil, err
	// }
	// if result.ServerInfo.Name != serverName {
	// 	return nil, errors.New("server name mismatch")
	// }
	// return client, nil

	// 临时返回nil
	return nil, errors.New("MCP功能暂时禁用")
}
