package openuem_nats

type Notification struct {
	From                  string `json:"from,omitempty"`
	To                    string `json:"to,omitempty"`
	Subject               string `json:"subject,omitempty"`
	MessageTitle          string `json:"message_title,omitempty"`
	MessageGreeting       string `json:"message_greeting,omitempty"`
	MessageText           string `json:"message_text,omitempty"`
	MessageAction         string `json:"message_action,omitempty"`
	MessageActionURL      string `json:"message_action_url,omitempty"`
	MessageAttachFileName string `json:"message_attach_filename,omitempty"`
	MessageAttachFile     string `json:"message_attach_file,omitempty"`
}

type CertificationRequest struct {
	Username     string `json:"username,omitempty"`
	Organization string `json:"organization,omitempty"`
	Country      string `json:"country,omitempty"`
	Province     string `json:"province,omitempty"`
	Locality     string `json:"locality,omitempty"`
	Address      string `json:"address,omitempty"`
	PostalCode   string `json:"postal_code,omitempty"`
	YearsValid   int    `json:"years_valid,omitempty"`
	MonthsValid  int    `json:"months_valid,omitempty"`
	DaysValid    int    `json:"days_valid,omitempty"`
	OCSP         string `json:"ocsp,omitempty"`
}
