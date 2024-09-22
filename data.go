package openuem_nats

type Notification struct {
	From             string `json:"from,omitempty"`
	To               string `json:"to,omitempty"`
	Subject          string `json:"subject,omitempty"`
	MessageTitle     string `json:"message_title,omitempty"`
	MessageGreeting  string `json:"message_greeting,omitempty"`
	MessageText      string `json:"message_text,omitempty"`
	MessageAction    string `json:"message_action,omitempty"`
	MessageActionURL string `json:"message_action_url,omitempty"`
}
