package openuem_nats

import "time"

type OpenUEMUpdateRequest struct {
	DownloadFrom string    `json:"download,omitempty"`
	DownloadHash string    `json:"download_hash,omitempty"`
	UpdateAt     time.Time `json:"updated_at,omitempty"`
	UpdateNow    bool      `json:"update_now,omitempty"`
}

type OpenUEMRollbackRequest struct {
	RollbackAt  time.Time `json:"updated_at,omitempty"`
	RollbackNow bool      `json:"update_now,omitempty"`
}
