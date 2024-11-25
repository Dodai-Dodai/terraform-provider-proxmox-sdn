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

	Bridge *string `json:"bridge,omitempty"` // for VLAN, QinQ

	Tag          *int64  `json:"tag,omitempty"`           // for QinQ
	VLANProtocol *string `json:"vlan-protocol,omitempty"` // for QinQ

	Peers []string `json:"peers,omitempty"` // for VXLAN

	Controller              *string  `json:"controller,omitempty"`                 // for EVPN
	VRFVXLAN                *int64   `json:"vrf-vxlan,omitempty"`                  // for EVPN
	MAC                     *string  `json:"mac,omitempty"`                        // for EVPN
	ExitNodes               []string `json:"exitnodes,omitempty"`                  // for EVPN
	PrimaryExitNode         *string  `json:"exitnodes-primary,omitempty"`          // for EVPN
	ExitNodesLocalRouting   *bool    `json:"exitnodes-local-routing,omitempty"`    // for EVPN
	AdvertiseSubnets        *bool    `json:"advertise-subnets,omitempty"`          // for EVPN
	DisableARPNdSuppression *bool    `json:"disable-arp-nd-suppression,omitempty"` // for EVPN
	RouteTargetImport       *string  `json:"rt-import,omitempty"`                  // for EVPN
}

// []stringでJSONのマーシャライズをしたいが、Proxmoxは"hoge1, hoge2"のような形式で受け取るため、カスタムマーシャラーを作成
// カスタムマーシャラー
func (n *SDNZone) MarshalJSON() ([]byte, error) {
	type Alias SDNZone
	return json.Marshal(&struct {
		Nodes     string `json:"nodes,omitempty"`
		Peers     string `json:"peers,omitempty"`
		ExitNodes string `json:"exitnodes,omitempty"`
		*Alias
	}{
		Nodes: func() string {
			if len(n.Nodes) > 0 {
				return strings.Join(n.Nodes, ", ")
			}
			return ""
		}(),
		Peers: func() string {
			if len(n.Peers) > 0 {
				return strings.Join(n.Peers, ", ")
			}
			return ""
		}(),
		ExitNodes: func() string {
			if len(n.ExitNodes) > 0 {
				return strings.Join(n.ExitNodes, ", ")
			}
			return ""
		}(),
		Alias: (*Alias)(n),
	})
}

// カスタムアンマーシャラー
func (n *SDNZone) UnmarshalJSON(data []byte) error {
	type Alias SDNZone
	aux := &struct {
		Nodes     string `json:"nodes,omitempty"`
		Peers     string `json:"peers,omitempty"`
		ExitNodes string `json:"exitnodes,omitempty"`
		*Alias
	}{
		Alias: &Alias{},
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// Assign unmarshalled data to n
	*n = SDNZone(*aux.Alias)

	// Split and trim Nodes
	if aux.Nodes != "" {
		rawNodes := strings.Split(aux.Nodes, ",")
		n.Nodes = make([]string, len(rawNodes))
		for i, s := range rawNodes {
			n.Nodes[i] = strings.TrimSpace(s)
		}
	} else {
		n.Nodes = nil
	}

	// Split and trim Peers
	if aux.Peers != "" {
		rawPeers := strings.Split(aux.Peers, ",")
		n.Peers = make([]string, len(rawPeers))
		for i, s := range rawPeers {
			n.Peers[i] = strings.TrimSpace(s)
		}
	} else {
		n.Peers = nil
	}

	// Split and trim ExitNodes
	if aux.ExitNodes != "" {
		rawExitNodes := strings.Split(aux.ExitNodes, ",")
		n.ExitNodes = make([]string, len(rawExitNodes))
		for i, s := range rawExitNodes {
			n.ExitNodes[i] = strings.TrimSpace(s)
		}
	} else {
		n.ExitNodes = nil
	}

	return nil
}
