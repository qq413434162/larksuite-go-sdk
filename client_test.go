package larksuite

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"os"
	"testing"

	larkim "github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
)

const (
	AppId     = "x"
	AppSecret = "x"
	// NotifierChatId 会话号
	NotifierChatId = "x"
)

func TestLarkSuite_SendMessage_Text(t *testing.T) {
	Convey("Test LarkSuite_SendMessage", t, func() {
		u := NewClient(AppId, AppSecret)

		req := &MessageReq{
			ReceiveId: NotifierChatId,
			MsgType:   MsgTypeEnum.Text,
			Content: `
{
	"text": "test1"
}`,
		}
		messageId, err := u.SendMessageByChatId(
			context.TODO(),
			req,
		)
		So(
			err,
			ShouldBeNil,
		)
		So(
			len(messageId),
			ShouldBeGreaterThan,
			0,
		)
	})
}

func TestLarkSuite_SendMessage_File(t *testing.T) {
	Convey("Test LarkSuite_SendMessage_File", t, func() {
		u := NewClient(AppId, AppSecret)

		req := &MessageReq{
			ReceiveId: NotifierChatId,
			MsgType:   MsgTypeEnum.File,
			Content:   `{"file_key":"file_v2_b30fb3b7-61a6-43dd-ac6b-22619e3451cg"}`,
		}
		messageId, err := u.SendMessageByChatId(
			context.TODO(),
			req,
		)
		So(
			err,
			ShouldBeNil,
		)
		So(
			len(messageId),
			ShouldBeGreaterThan,
			0,
		)
	})
}

func TestLarkSuite_UploadFile(t *testing.T) {
	Convey("Test LarkSuite_UploadFile", t, func() {
		u := NewClient(AppId, AppSecret)
		ctx := context.TODO()
		f, err := os.Open("/Users/x/Downloads/Untitled.xlsx")
		So(
			err,
			ShouldBeNil,
		)

		req := &FileReq{
			FileType: larkim.FileTypeXls,
			FileName: "test.xls",
			File:     f,
		}
		fileKey, err := u.UploadFile(
			ctx,
			req,
		)
		So(
			err,
			ShouldBeNil,
		)
		So(
			len(fileKey),
			ShouldBeGreaterThan,
			0,
		)
	})
}
