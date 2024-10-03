package openuem_nats

import "time"

type DeployResult struct {
	Action      string    `json:"action,omitempty"`
	Installed   time.Time `json:"installed,omitempty"`
	Updated     time.Time `json:"updated,omitempty"`
	Uninstalled time.Time `json:"uninstalled,omitempty"`
	Info        string    `json:"info,omitempty"`
}
