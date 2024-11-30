package client

type DhcpRange struct {
	StartAddress string `json:"start-address"`
	EndAddress   string `json:"end-address"`
}

type SDNSubnets struct {
	Subnet        string      `json:"subnet"`                    // required field
	Cidr          string      `json:"cidr,omitempty"`            // optional
	Type          string      `json:"type"`                      // required but always set 'subnet'
	Vnet          string      `json:"vnet"`                      // required field
	DhcpDnsServer *string     `json:"dhcp-dns-server,omitempty"` // optional
	DhcpRange     []DhcpRange `json:"dhcp-range,omitempty"`      // optional
	DnsZonePrefix *string     `json:"dnszoneprefix,omitempty"`   // optional
	Gateway       *string     `json:"gateway,omitempty"`         // optional
	Snat          *IntBool    `json:"snat,omitempty"`            // optional
	Zone          *string     `json:"zone,omitempty"`            // optional
}
