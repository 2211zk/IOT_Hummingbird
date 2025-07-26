package system

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/mcp/client"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/gin-gonic/gin"
	// "github.com/mark3labs/mcp-go/mcp" // 临时注释以解决Go版本兼容性问题
)

// Create
// @Tags      mcp
// @Summary   自动McpTool
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.AutoMcpTool  true  "创建自动代码"
// @Success   200   {string}  string                 "{"success":true,"data":{},"msg":"创建成功"}"
// @Router    /autoCode/mcp [post]
func (a *AutoCodeTemplateApi) MCP(c *gin.Context) {
	var info request.AutoMcpTool
	err := c.ShouldBindJSON(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	toolFilePath, err := autoCodeTemplateService.CreateMcp(c.Request.Context(), info)
	if err != nil {
		response.FailWithMessage("创建失败", c)
		global.GVA_LOG.Error(err.Error())
		return
	}
	response.OkWithMessage("创建成功,MCP Tool路径:"+toolFilePath, c)
}

// Create
// @Tags      mcp
// @Summary   自动McpTool
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      request.AutoMcpTool  true  "创建自动代码"
// @Success   200   {string}  string                 "{"success":true,"data":{},"msg":"创建成功"}"
// @Router    /autoCode/mcpList [post]
func (a *AutoCodeTemplateApi) MCPList(c *gin.Context) {

	baseUrl := fmt.Sprintf("http://127.0.0.1:%d%s", global.GVA_CONFIG.System.Addr, global.GVA_CONFIG.MCP.SSEPath)

	testClient, err := client.NewClient(baseUrl, "testClient", "v1.0.0", global.GVA_CONFIG.MCP.Name)
	if err != nil {
		response.FailWithMessage("创建MCP客户端失败: "+err.Error(), c)
		return
	}
	defer func() {
		if testClient != nil {
			if closer, ok := testClient.(interface{ Close() }); ok {
				closer.Close()
			}
		}
	}()

	// 临时禁用MCP功能
	response.FailWithMessage("MCP功能暂时禁用", c)
	return

	// 以下代码暂时注释掉，因为MCP功能被禁用
	/*
		if err != nil {
			response.FailWithMessage("创建失败", c)
			global.GVA_LOG.Error(err.Error())
			return
		}

		mcpServerConfig := map[string]interface{}{
			"mcpServers": map[string]interface{}{
				global.GVA_CONFIG.MCP.Name: map[string]string{
					"url": baseUrl,
				},
			},
		}
		response.OkWithData(gin.H{
			"mcpServerConfig": mcpServerConfig,
			"list":            list,
		}, c)
	*/
}

// Create
// @Tags      mcp
// @Summary   测试McpTool
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      object  true  "调用MCP Tool的参数"
// @Success   200   {object}  response.Response  "{"success":true,"data":{},"msg":"测试成功"}"
// @Router    /autoCode/mcpTest [post]
func (a *AutoCodeTemplateApi) MCPTest(c *gin.Context) {
	// 定义接口请求结构
	var testRequest struct {
		Name      string                 `json:"name" binding:"required"`      // 工具名称
		Arguments map[string]interface{} `json:"arguments" binding:"required"` // 工具参数
	}

	// 绑定JSON请求体
	if err := c.ShouldBindJSON(&testRequest); err != nil {
		response.FailWithMessage("参数解析失败:"+err.Error(), c)
		return
	}

	// 创建MCP客户端
	baseUrl := fmt.Sprintf("http://127.0.0.1:%d%s", global.GVA_CONFIG.System.Addr, global.GVA_CONFIG.MCP.SSEPath)
	testClient, err := client.NewClient(baseUrl, "testClient", "v1.0.0", global.GVA_CONFIG.MCP.Name)
	if err != nil {
		response.FailWithMessage("创建MCP客户端失败:"+err.Error(), c)
		return
	}
	defer func() {
		if testClient != nil {
			if closer, ok := testClient.(interface{ Close() }); ok {
				closer.Close()
			}
		}
	}()

	// 临时禁用MCP功能
	response.FailWithMessage("MCP功能暂时禁用", c)
	return
}
