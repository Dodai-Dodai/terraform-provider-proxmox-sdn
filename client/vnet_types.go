package client

type SDNVnet struct {
	Vnet      string   `json:"vnet"`                // required field
	Zone      string   `json:"zone"`                // required field
	Type      string   `json:"type,omitempty"`      // optional but always set 'vnet'
	Alias     *string  `json:"alias,omitempty"`     // optional
	Tag       *int64   `json:"tag,omitempty"`       // optional
	Vlanaware *IntBool `json:"vlanaware,omitempty"` // optional
}
