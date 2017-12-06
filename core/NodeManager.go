package core

import (
	nomad "github.com/hashicorp/nomad/api"
)

type NodeManager struct {
	platform *Platform
}

func NewNodeManager(platform *Platform) *NodeManager {
	return &NodeManager{
		platform: platform,
	}
}

func (m *NodeManager) List() ([]*NodeSummary, error) {
	nodes, _, err := m.platform.nomadClient.Nodes().List(&nomad.QueryOptions{})

	if err != nil {
		return nil, err
	}

	res := make([]*NodeSummary, len(nodes))
	for i, v := range nodes {
		res[i] = &NodeSummary{
			Id: v.ID,
			Name: v.Name,
			Status: v.Status,
			Version: v.Version,
			Datacenter: v.Datacenter,
			Drain: v.Drain,
			NodeClass: v.NodeClass,
			StatusDescription: v.StatusDescription,
		}
	}

	return res, nil
}

func (m *NodeManager) Get(nodeId string) (*Node, error) {
	node, _, err := m.platform.nomadClient.Nodes().Info(nodeId, &nomad.QueryOptions{})

	if err != nil {
		return nil, err
	}

	return &Node{
		Id: node.ID,
		Name: node.Name,
		Attributes: node.Attributes,
		StatusDescription: node.StatusDescription,
		NodeClass: node.NodeClass,
		Drain: node.Drain,
		Datacenter: node.Datacenter,
		HTTPAddr: node.HTTPAddr,
		Links: node.Links,
		Meta: node.Meta,
		Status: node.Status,
		TLSEnabled: node.TLSEnabled,
		Resources: &NodeResources{
			CPU: node.Resources.CPU,
			Disk: node.Resources.DiskMB,
			IOPS: node.Resources.IOPS,
			Memory: node.Resources.MemoryMB,
		},
	}, nil
}
