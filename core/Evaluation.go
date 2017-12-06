package core

type Evaluation struct {
	Id           string `json:"id"`
	JobKind      string `json:"job_kind"`
	JobId        string `json:"job_id"`
	Priority     int    `json:"priority"`
	Status       string `json:"status"`
}

type NodeAllocationState struct {
	NodeId        string   `json:"node_id"`
	AllocationId  string   `json:"allocation_id"`
	JobKind       string   `json:"job_kind"`
	JobId         string   `json:"job_id"`
	DesiredStatus string   `json:"desired_status"`
	ActualStatus  string   `json:"actual_status"`
	Trace         []*Trace `json:"trace"`
}

type Trace struct {
	Type      string `json:"type"`
	Timestamp int64  `json:"timestamp"`
	Message   string `json:"message"`
}
