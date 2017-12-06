package core

type NodeSummary struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Status            string `json:"status"`
	Datacenter        string `json:"datacenter"`
	NodeClass         string `json:"node_class"`
	Version           string `json:"version"`
	Drain             bool 	 `json:"drain"`
	StatusDescription string `json:"status_description"`
}

type Node struct {
	Id         string            `json:"id"`
	Name       string            `json:"name"`
	Attributes map[string]string `json:"attributes"`
	Meta map[string]string `json:"meta"`
	Resources  *NodeResources    `json:"resources"`
	Datacenter        string `json:"datacenter"`
	NodeClass         string `json:"node_class"`
	Drain             bool 	 `json:"drain"`
	Status            string `json:"status"`
	StatusDescription string `json:"status_description"`
	HTTPAddr          string `json:"http_addr"`
	TLSEnabled        bool `json:"tls_enabled"`
	Links             map[string]string `json:"links"`
}

type NodeResources struct {
	CPU    *int `json:"cpu"`
	Memory *int `json:"memory_mb"`
	Disk   *int `json:"disk_mb"`
	IOPS   *int `json:"iops"`
}
