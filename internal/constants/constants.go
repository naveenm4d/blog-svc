package constants

type Status string

const (
	StatusUndefined Status = "Undefined"
	StatusDrafted   Status = "Drafted"
	StatusPublished Status = "Published"
	StatusArchived  Status = "Archived"
)

const (
	ResponseSuccess bool = true
	ResponseFailed  bool = false
)
