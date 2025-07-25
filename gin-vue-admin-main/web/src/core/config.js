/**
 * 网站配置文件
 */
const greenText = (text) => `\x1b[32m${text}\x1b[0m`

const config = {
  appName: '万联智控',
  appLogo: 'logo_login.png',
  showViteLogo: true,
  logs: []
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    console.log(
      greenText(
        `> 欢迎使用万联智控智能设备管理平台`
      )
    )
    console.log(greenText(`> 当前版本:v1.0.0`))
    console.log(greenText(`> 技术支持：万联智控团队`))
    console.log(
      greenText(`> 项目地址：https://github.com/wanlian-iot`))
    console.log(greenText(`> 插件市场:https://plugin.wanlian-iot.com`))
    console.log(
      greenText(`> 技术支持社区:https://support.wanlian-iot.com`)
    )
    console.log(
      greenText(
        `> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
      )
    )
    console.log(
      greenText(`> 默认前端文件运行地址:http://127.0.0.1:${env.VITE_CLI_PORT}`)
    )
    console.log(
      greenText(
        `--------------------------------------版权声明--------------------------------------`
      )
    )
    console.log(greenText(`** 版权所有方：万联智控团队 **`))
    console.log(greenText(`** 版权持有公司：万联智控科技有限公司 **`))
    console.log(
      greenText(
        `** 商用授权请联系：contact@wanlian-iot.com **`
      )
    )
    console.log('\n')
  }
}

export default config
