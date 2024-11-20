package client

import (
	"encoding/json"
	"fmt"
)

type SDNZone struct {
	Zone       string     `json:"zone"`
	Type       string     `json:"type"`
	MTU        *int64     `json:"mtu,omitempty"`
	Nodes      NodesField `json:"nodes,omitempty"`
	IPAM       *string    `json:"ipam,omitempty"`
	DNS        *string    `json:"dns,omitempty"`
	ReverseDNS *string    `json:"reversedns,omitempty"`
	DNSZone    *string    `json:"dnszone,omitempty"`
	// Simple     *SimpleConfig `json:"simple,omitempty"`
	VLAN  *VLANConfig  `json:"vlan,omitempty"`
	QinQ  *QinQConfig  `json:"qinq,omitempty"`
	VXLAN *VXLANConfig `json:"vxlan,omitempty"`
	EVPN  *EVPNConfig  `json:"evpn,omitempty"`
}

// type SimpleConfig struct {
// 	AutoDHCP *bool `json:"auto_dhcp"`
// }

type NodesField []string

func (n *NodesField) UnmarshalJSON(data []byte) error {
	var single string
	if err := json.Unmarshal(data, &single); err == nil {
		*n = []string{single}
		return nil
	}

	var multiple []string
	if err := json.Unmarshal(data, &multiple); err != nil {
		return fmt.Errorf("failed to unmarshal nodes field: %w", err)
	}
	*n = multiple
	return nil
}

type VLANConfig struct {
	Bridge string `json:"bridge"`
}

type QinQConfig struct {
	Bridge       string  `json:"bridge"`
	Tag          int64   `json:"tag"`
	VLANProtocol *string `json:"vlan_protocol,omitempty"`
}

type VXLANConfig struct {
	Peer []string `json:"peer"`
}

type EVPNConfig struct {
	Controller              string   `json:"controller"`
	VRFVXLAN                int64    `json:"vrf_vxlan"`
	MAC                     *string  `json:"mac,omitempty"`
	ExitNodes               []string `json:"exitnodes,omitempty"`
	PrimaryExitNode         *string  `json:"primary_exitnode,omitempty"`
	ExitNodesLocalRouting   *bool    `json:"exitnodes_local_routing,omitempty"`
	AdvertiseSubnets        *bool    `json:"advertise_subnets,omitempty"`
	DisableARPNdSuppression *bool    `json:"disable_arp_nd_suppression,omitempty"`
	RouteTargetImport       *string  `json:"rt_import,omitempty"`
}
