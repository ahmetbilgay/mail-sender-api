package models

type MailModel struct {
	MsgSender   string `json:"msg_sender"`
	MsgTo       string `json:"msg_to"`
	Subject     string `json:"subject"`
	BodyMessage string `json:"body_message"`
	Password    string `json:"password"`
}
