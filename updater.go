package openuem_nats

import "time"

type OpenUEMUpdateRequest struct {
	LatestVersion string    `json:"latest_version,omitempty"`
	DownloadFrom  string    `json:"download,omitempty"`
	DownloadHash  string    `json:"download_hash,omitempty"`
	UpdateAt      time.Time `json:"updated_at,omitempty"`
	UpdateNow     bool      `json:"update_now,omitempty"`
}

const UPDATE_ERROR = "admin.update.agents.task_status_error"
const UPDATE_PENDING = "admin.update.agents.task_status_pending"
const UPDATE_SUCCESS = "admin.update.agents.task_status_success"

func TaskUpdatePossibleStatus() []string {
	return []string{UPDATE_ERROR, UPDATE_PENDING, UPDATE_SUCCESS}
}
