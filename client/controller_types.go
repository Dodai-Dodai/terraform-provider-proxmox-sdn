package client

import (
	"encoding/json"
	"strings"
)

type SDNController struct {
	Controller              string   `json:"controller"`
	Type                    string   `json:"type"`
	ASN                     *int64   `json:"asn,omitempty"`
	BgpMultipathAsPathRelax *IntBool `json:"bgp-multipath-aspath-relax,omitempty"`
	Ebgp                    *IntBool `json:"ebgp,omitempty"`
	EbgpMultihop            *int64   `json:"ebgp-multihop,omitempty"`
	ISISDomain              *string  `json:"isis-domain,omitempty"`
	ISISIfaces              *string  `json:"isis-ifaces,omitempty"`
	ISISNet                 *string  `json:"isis-net,omitempty"`
	Loopback                *string  `json:"loopback,omitempty"`
	Node                    *string  `json:"node,omitempty"`
	Peers                   []string `json:"peers,omitempty"`
}

func (s *SDNController) UnmarshalJSON(data []byte) error {
	// まずは同じフィールド構成だが peers は string にしてある中間 struct を定義
	type Alias SDNController
	type rawSDNController struct {
		Peers string `json:"peers"`
		*Alias
	}

	tmp := &rawSDNController{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, tmp); err != nil {
		return err
	}

	// peers をカンマ区切りで配列化
	if tmp.Peers != "" {
		s.Peers = strings.Split(tmp.Peers, ",")
	} else {
		s.Peers = nil
	}

	return nil
}
