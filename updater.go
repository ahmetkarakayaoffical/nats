package openuem_nats

import "time"

type OpenUEMUpdateRequest struct {
	DownloadFrom string    `json:"download,omitempty"`
	UpdateAt     time.Time `json:"updated_at,omitempty"`
	UpdateNow    bool      `json:"update_now,omitempty"`
}
