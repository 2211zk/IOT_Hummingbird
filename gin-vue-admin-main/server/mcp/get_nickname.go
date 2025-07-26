package mcpTool

import (
	"context"
	"errors"
	// "github.com/mark3labs/mcp-go/mcp" // 临时注释以解决Go版本兼容性问题
)

func init() {
	// RegisterTool(&GetNickname{}) // 临时注释
}

type GetNickname struct{}

// 根据用户username获取nickname
func (t *GetNickname) New() interface{} { // 临时修改返回类型
	// return mcp.NewTool("getNickname", // 临时注释
	// 	mcp.WithDescription("根据用户username获取nickname"),
	// 	mcp.WithString("username",
	// 		mcp.Required(),
	// 		mcp.Description("用户的username"),
	// 	))
	return nil
}

// Handle 处理获取昵称的请求
func (t *GetNickname) Handle(ctx context.Context, request interface{}) (interface{}, error) { // 临时修改参数和返回类型
	// 临时注释所有MCP相关代码
	return nil, errors.New("MCP功能暂时禁用")
}
