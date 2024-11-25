package client

import (
	"encoding/json"
	"strings"
)

type IntBool bool

func (b *IntBool) UnmarshalJSON(data []byte) error {
	var i int
	if err := json.Unmarshal(data, &i); err != nil {
		return err
	}
	*b = IntBool(i != 0)
	return nil
}

func (b IntBool) MarshalJSON() ([]byte, error) {
	if b {
		return json.Marshal(1)
	}
	return json.Marshal(0)
}

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
	ExitNodesLocalRouting   *IntBool `json:"exitnodes-local-routing,omitempty"`    // for EVPN
	AdvertiseSubnets        *IntBool `json:"advertise-subnets,omitempty"`          // for EVPN
	DisableARPNdSuppression *IntBool `json:"disable-arp-nd-suppression,omitempty"` // for EVPN
	RouteTargetImport       *string  `json:"rt-import,omitempty"`                  // for EVPN
}

// カスタムアンマーシャラー
func (n *SDNZone) UnmarshalJSON(data []byte) error {
	type Alias SDNZone
	aux := &struct {
		Nodes                   string   `json:"nodes,omitempty"`
		Peers                   string   `json:"peers,omitempty"`
		ExitNodes               string   `json:"exitnodes,omitempty"`
		ExitNodesLocalRouting   *IntBool `json:"exitnodes-local-routing,omitempty"`
		AdvertiseSubnets        *IntBool `json:"advertise-subnets,omitempty"`
		DisableARPNdSuppression *IntBool `json:"disable-arp-nd-suppression,omitempty"`
		*Alias
	}{
		Alias: (*Alias)(n),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	// Nodes の処理
	if aux.Nodes != "" {
		n.Nodes = parseCommaSeparatedList(aux.Nodes)
	} else {
		n.Nodes = nil
	}

	// Peers の処理
	if aux.Peers != "" {
		n.Peers = parseCommaSeparatedList(aux.Peers)
	} else {
		n.Peers = nil
	}

	// ExitNodes の処理
	if aux.ExitNodes != "" {
		n.ExitNodes = parseCommaSeparatedList(aux.ExitNodes)
	} else {
		n.ExitNodes = nil
	}

	// ブール値のフィールドを設定
	n.ExitNodesLocalRouting = aux.ExitNodesLocalRouting
	n.AdvertiseSubnets = aux.AdvertiseSubnets
	n.DisableARPNdSuppression = aux.DisableARPNdSuppression

	return nil
}

// 補助関数
func parseCommaSeparatedList(s string) []string {
	rawItems := strings.Split(s, ",")
	items := make([]string, 0, len(rawItems))
	for _, item := range rawItems {
		trimmed := strings.TrimSpace(item)
		if trimmed != "" {
			items = append(items, trimmed)
		}
	}
	return items
}
