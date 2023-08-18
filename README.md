# 飞书 Go SDK

Lark Suite Go SDK 是一个方便易用的软件包，基于[飞书官方sdk](https://github.com/larksuite/oapi-sdk-go)深度封装，用于处理即时通讯和文件上传等功能。该SDK简化了与Lark Suite API的工作流程，使开发人员能够快速将通讯和文件上传功能集成到他们的应用程序中。

## 特性

- 与Lark Suite API的无缝集成，用于发送即时通讯、上传文件等功能。
- 提供创建消息、发送消息、上传文件等易于使用的函数。
- 使用`github.com/larksuite/oapi-sdk-go/v3`软件包处理API调用。
- 支持自定义选项，以更好地控制集成。

## 安装

要安装飞书 Go SDK，只需在您的项目目录中运行以下命令：

```sh
go get github.com/qq413434162/larksuite-go-sdk
```
## 使用
### 并在您的Go代码中导入SDK软件包：
```go
import "github.com/qq413434162/larksuite-go-sdk"
```
### 创建一个新的飞书客户端实例：
```go
client := larksuite.NewClient("your-app-id", "your-app-secret")
```
### 发送即时通讯消息
```go
messageReq := larksuite.MessageReq{
// 设置消息内容和接收者等信息
}

messageID, err := client.SendMessageByChatId(context.Background(), &messageReq)
if err != nil {
// 处理错误
}

```
### 上传文件
```go
fileReq := larksuite.FileReq{
// 设置文件类型、文件名、文件内容等信息
}

fileKey, err := client.UploadFile(context.Background(), &fileReq)
if err != nil {
// 处理错误
}
```
## 更多
有关更多示例和详细文档，请参阅飞书的[API文档](https://open.feishu.cn/document/server-docs/api-call-guide/server-api-list)和[SDK的GitHub](https://github.com/larksuite/oapi-sdk-go/)。


## 贡献
欢迎贡献、提交问题和功能请求！请随时在[GitHub](https://github.com/qq413434162/larksuite-go-sdk)上打开拉取请求或提交问题。

## 许可证
此SDK采用MIT license许可证发布。有关详情，请参阅飞书官方SDK的LICENSE。

**免责声明：** 此SDK未经飞书官方维护。请自行承担风险。