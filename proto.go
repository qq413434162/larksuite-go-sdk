package larksuite

import (
	"os"
)

type MessageReq struct {
	ReceiveId string `json:"receive_id"`
	MsgType   string `json:"msg_type"`
	Content   string `json:"file"`
}

type FileReq struct {
	FileType string   `json:"file_type"`
	FileName string   `json:"file_name"`
	File     *os.File `json:"file"`
}

type msgTypeEnum struct {
	Text        string `json:"text" desc:"纯文本"`
	Interactive string `json:"interactive" desc:"消息卡片"`
	File        string `json:"file" desc:"文件"`
}

var MsgTypeEnum = msgTypeEnum{
	"text",
	"interactive",
	"file",
}
