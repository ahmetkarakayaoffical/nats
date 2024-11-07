package openuem_nats

const REGISTER_CERTIFICATE_SENT = "users.certificate_sent"
const REGISTER_COMPLETE = "users.completed"
const REGISTER_IN_REVIEW = "users.review_request"

func RegisterStatus() []string {
	return []string{REGISTER_CERTIFICATE_SENT, REGISTER_COMPLETE, REGISTER_IN_REVIEW}
}
