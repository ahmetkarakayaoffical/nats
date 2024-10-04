package openuem_nats

import "time"

type DeployResult struct {
	Action string    `json:"action,omitempty"`
	When   time.Time `json:"when,omitempty"`
	Info   string    `json:"info,omitempty"`
}
