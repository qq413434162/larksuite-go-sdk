package larksuite

import (
	"context"

	lark "github.com/larksuite/oapi-sdk-go/v3"
	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

type Client struct {
	client *lark.Client
}

func NewClient(appId, appSecret string) *Client {
	client := lark.NewClient(
		appId,
		appSecret,
	)
	return &Client{
		client,
	}
}

// SendMessageByChatId 发送消息
func (s *Client) SendMessageByChatId(
	_ context.Context,
	req *MessageReq,
) (
	messageId string,
	err error,
) {
	result, err := s.client.Im.Message.Create(context.Background(), larkim.NewCreateMessageReqBuilder().
		ReceiveIdType(larkim.ReceiveIdTypeChatId).
		Body(larkim.NewCreateMessageReqBodyBuilder().
			MsgType(req.MsgType).
			ReceiveId(req.ReceiveId).
			Content(req.Content).
			Build()).
		Build())
	if err != nil {
		return
	}
	if !result.Success() {
		return
	}
	if result == nil || result.Data == nil || result.Data.MessageId == nil {
		return
	}
	messageId = *result.Data.MessageId
	return
}

// UploadFile 上传文件
func (s *Client) UploadFile(
	_ context.Context,
	req *FileReq,
) (
	fileKey string,
	err error,
) {
	result, err := s.client.Im.File.Create(context.Background(),
		larkim.NewCreateFileReqBuilder().
			Body(larkim.NewCreateFileReqBodyBuilder().
				FileType(req.FileType).
				FileName(req.FileName).
				File(req.File).
				Build()).
			Build())
	if err != nil {
		return
	}
	if result == nil || result.Data == nil || result.Data.FileKey == nil {
		return
	}
	fileKey = *result.Data.FileKey
	return
}
