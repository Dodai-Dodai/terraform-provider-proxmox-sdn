package client

import (
	"encoding/json"
	"strings"
)

type SDNZone struct {
	Zone       string   `json:"zone"`
	Type       string   `json:"type"`
	MTU        *int64   `json:"mtu,omitempty"`
	Nodes      []string `json:"nodes,omitempty"`
	IPAM       *string  `json:"ipam,omitempty"`
	DNS        *string  `json:"dns,omitempty"`
	ReverseDNS *string  `json:"reversedns,omitempty"`
	DNSZone    *string  `json:"dnszone,omitempty"`
	// Simple     *SimpleConfig `json:"simple,omitempty"`
	VLAN  *VLANConfig  `json:"vlan,omitempty"`
	QinQ  *QinQConfig  `json:"qinq,omitempty"`
	VXLAN *VXLANConfig `json:"vxlan,omitempty"`
	EVPN  *EVPNConfig  `json:"evpn,omitempty"`
}

// type SimpleConfig struct {
// 	AutoDHCP *bool `json:"auto_dhcp"`
// }

func (s *SDNZone) UnmarshalJSON(data []byte) error {
	type Alias SDNZone
	aux := &struct {
		Nodes string `json:"nodes,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.Nodes != "" {
		s.Nodes = strings.Split(aux.Nodes, ",")
	} else {
		s.Nodes = []string{}
	}
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
	Peer []string `json:"peers"`
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

func (v *VXLANConfig) UnmarshalJSON(data []byte) error {
	var tmp struct {
		Peers string `json:"peers"`
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	if tmp.Peers != "" {
		v.Peer = strings.Split(tmp.Peers, ",")
	} else {
		v.Peer = []string{}
	}
	return nil
}

func (e *EVPNConfig) UnmarshalJSON(data []byte) error {
	type Alias EVPNConfig
	aux := &struct {
		ExitNodes string `json:"exitnodes,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	if aux.ExitNodes != "" {
		e.ExitNodes = strings.Split(aux.ExitNodes, ",")
	} else {
		e.ExitNodes = []string{}
	}
	return nil
}
