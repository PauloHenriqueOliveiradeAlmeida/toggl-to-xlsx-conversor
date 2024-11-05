package togglResponse

import "time"

type TimeEntriesResponse struct {
	Id              int64         `json:"id"`
	WorkspaceId     int           `json:"workspace_id"`
	ProjectId       int           `json:"project_id"`
	TaskId          string        `json:"task_id"`
	Billable        bool          `json:"billable"`
	Start           time.Time     `json:"start"`
	Stop            time.Time     `json:"stop"`
	Duration        time.Duration `json:"duration"`
	Description     string        `json:"description"`
	Tags            []string      `json:"tags"`
	TagIds          []string      `json:"tag_ids"`
	Duronly         bool          `json:"duronly"`
	At              time.Time     `json:"at"`
	ServerDeletedAt time.Time     `json:"server_deleted_at"`
	UserID          int           `json:"user_id"`
	UID             int           `json:"uid"`
	Wid             int           `json:"wid"`
	Pid             int           `json:"pid"`
}
