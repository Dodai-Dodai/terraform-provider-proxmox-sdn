私は今、ProxmoxのSDN Zoneを作成するTerraformのカスタムプロバイダーを作成しています。このプロバイダーは、ProxmoxのAPIを使用してSDN Zoneを作成、更新、削除するために、`pvesh`コマンドを使用します。このプロバイダーは、TerraformのリソースとしてSDN Zoneを表現し、Terraformの構成ファイルでSDN Zoneを定義できるようにします。
このとき、nodesやpeerなど、リストであるフィールドをどのように扱うかがわかりません。
参考に、現状のProxmox上のSDN Zone一覧を示します。
```sh
pvesh get cluster/sdn/zones --output-format json | jq
[
  {
    "bridge": "vmbr0",
    "digest": "0fc7f123b70be7ad2f6230994ff91a72a081a781",
    "tag": 100,
    "type": "qinq",
    "zone": "qinqzone"
  },
  {
    "dhcp": "dnsmasq",
    "digest": "0fc7f123b70be7ad2f6230994ff91a72a081a781",
    "ipam": "pve",
    "type": "simple",
    "zone": "simple"
  },
  {
    "bridge": "vmbr0",
    "digest": "0fc7f123b70be7ad2f6230994ff91a72a081a781",
    "type": "vlan",
    "zone": "test1"
  },
  {
    "bridge": "vmbr0",
    "digest": "0fc7f123b70be7ad2f6230994ff91a72a081a781",
    "type": "vlan",
    "zone": "vlan"
  },
  {
    "digest": "0fc7f123b70be7ad2f6230994ff91a72a081a781",
    "peers": "150.89.233.20,150.89.233.21",
    "type": "vxlan",
    "zone": "vx"
  },
  {
    "digest": "0fc7f123b70be7ad2f6230994ff91a72a081a781",
    "ipam": "pve",
    "mtu": 1450,
    "peers": "150.89.233.20,150.89.233.21",
    "type": "vxlan",
    "zone": "vxlan"
  },
  {
    "dhcp": "dnsmasq",
    "digest": "0fc7f123b70be7ad2f6230994ff91a72a081a781",
    "ipam": "pve",
    "mtu": 1400,
    "nodes": "pvewata01",
    "type": "simple",
    "zone": "vzone1"
  }
]
```
このとき、datasourceとして取得すると、このようになっています
```json
{
  "version": 4,
  "terraform_version": "1.5.7",
  "serial": 6,
  "lineage": "6c3a815c-d1c8-68c9-9cb7-cc86ded1e48e",
  "outputs": {
    "zones": {
      "value": {
        "zones": [
          {
            "dns": null,
            "dnszone": null,
            "evpn": null,
            "ipam": null,
            "mtu": null,
            "nodes": null,
            "qinq": {
              "bridge": "vmbr0",
              "tag": 100,
              "vlanprotocol": null
            },
            "reversedns": null,
            "type": "qinq",
            "vlan": null,
            "vxlan": null,
            "zone": "qinqzone"
          },
          {
            "dns": null,
            "dnszone": null,
            "evpn": null,
            "ipam": "pve",
            "mtu": null,
            "nodes": null,
            "qinq": null,
            "reversedns": null,
            "type": "simple",
            "vlan": null,
            "vxlan": null,
            "zone": "simple"
          },
          {
            "dns": null,
            "dnszone": null,
            "evpn": null,
            "ipam": null,
            "mtu": null,
            "nodes": null,
            "qinq": null,
            "reversedns": null,
            "type": "vlan",
            "vlan": {
              "bridge": "vmbr0"
            },
            "vxlan": null,
            "zone": "test1"
          },
          {
            "dns": null,
            "dnszone": null,
            "evpn": null,
            "ipam": null,
            "mtu": null,
            "nodes": null,
            "qinq": null,
            "reversedns": null,
            "type": "vlan",
            "vlan": {
              "bridge": "vmbr0"
            },
            "vxlan": null,
            "zone": "vlan"
          },
          {
            "dns": null,
            "dnszone": null,
            "evpn": null,
            "ipam": null,
            "mtu": null,
            "nodes": null,
            "qinq": null,
            "reversedns": null,
            "type": "vxlan",
            "vlan": null,
            "vxlan": {
              "peer": null
            },
            "zone": "vx"
          },
          {
            "dns": null,
            "dnszone": null,
            "evpn": null,
            "ipam": "pve",
            "mtu": 1450,
            "nodes": null,
            "qinq": null,
            "reversedns": null,
            "type": "vxlan",
            "vlan": null,
            "vxlan": {
              "peer": null
            },
            "zone": "vxlan"
          },
          {
            "dns": null,
            "dnszone": null,
            "evpn": null,
            "ipam": "pve",
            "mtu": 1400,
            "nodes": [
              "pvewata01"
            ],
            "qinq": null,
            "reversedns": null,
            "type": "simple",
            "vlan": null,
            "vxlan": null,
            "zone": "vzone1"
          }
        ]
      },
      "type": [
        "object",
        {
          "zones": [
            "list",
            [
              "object",
              {
                "dns": "string",
                "dnszone": "string",
                "evpn": [
                  "object",
                  {
                    "advertisesubnets": "bool",
                    "controller": "string",
                    "disablearpndsuppression": "bool",
                    "exitnodes": [
                      "set",
                      "string"
                    ],
                    "exitnodeslocalrouting": "bool",
                    "mac": "string",
                    "primaryexitnode": "string",
                    "rtimport": "string",
                    "vrf_vxlan": "number"
                  }
                ],
                "ipam": "string",
                "mtu": "number",
                "nodes": [
                  "set",
                  "string"
                ],
                "qinq": [
                  "object",
                  {
                    "bridge": "string",
                    "tag": "number",
                    "vlanprotocol": "string"
                  }
                ],
                "reversedns": "string",
                "type": "string",
                "vlan": [
                  "object",
                  {
                    "bridge": "string"
                  }
                ],
                "vxlan": [
                  "object",
                  {
                    "peer": [
                      "set",
                      "string"
                    ]
                  }
                ],
                "zone": "string"
              }
            ]
          ]
        }
      ]
    }
  },
  "resources": [
    {
      "mode": "data",
      "type": "proxmox-sdn_zone",
      "name": "zones",
      "provider": "provider[\"registry.terraform.io/dodai-dodai/proxmox-sdn\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "zones": [
              {
                "dns": null,
                "dnszone": null,
                "evpn": null,
                "ipam": null,
                "mtu": null,
                "nodes": null,
                "qinq": {
                  "bridge": "vmbr0",
                  "tag": 100,
                  "vlanprotocol": null
                },
                "reversedns": null,
                "type": "qinq",
                "vlan": null,
                "vxlan": null,
                "zone": "qinqzone"
              },
              {
                "dns": null,
                "dnszone": null,
                "evpn": null,
                "ipam": "pve",
                "mtu": null,
                "nodes": null,
                "qinq": null,
                "reversedns": null,
                "type": "simple",
                "vlan": null,
                "vxlan": null,
                "zone": "simple"
              },
              {
                "dns": null,
                "dnszone": null,
                "evpn": null,
                "ipam": null,
                "mtu": null,
                "nodes": null,
                "qinq": null,
                "reversedns": null,
                "type": "vlan",
                "vlan": {
                  "bridge": "vmbr0"
                },
                "vxlan": null,
                "zone": "test1"
              },
              {
                "dns": null,
                "dnszone": null,
                "evpn": null,
                "ipam": null,
                "mtu": null,
                "nodes": null,
                "qinq": null,
                "reversedns": null,
                "type": "vlan",
                "vlan": {
                  "bridge": "vmbr0"
                },
                "vxlan": null,
                "zone": "vlan"
              },
              {
                "dns": null,
                "dnszone": null,
                "evpn": null,
                "ipam": null,
                "mtu": null,
                "nodes": null,
                "qinq": null,
                "reversedns": null,
                "type": "vxlan",
                "vlan": null,
                "vxlan": {
                  "peer": null
                },
                "zone": "vx"
              },
              {
                "dns": null,
                "dnszone": null,
                "evpn": null,
                "ipam": "pve",
                "mtu": 1450,
                "nodes": null,
                "qinq": null,
                "reversedns": null,
                "type": "vxlan",
                "vlan": null,
                "vxlan": {
                  "peer": null
                },
                "zone": "vxlan"
              },
              {
                "dns": null,
                "dnszone": null,
                "evpn": null,
                "ipam": "pve",
                "mtu": 1400,
                "nodes": [
                  "pvewata01"
                ],
                "qinq": null,
                "reversedns": null,
                "type": "simple",
                "vlan": null,
                "vxlan": null,
                "zone": "vzone1"
              }
            ]
          },
          "sensitive_attributes": []
        }
      ]
    },
    {
      "mode": "managed",
      "type": "proxmox-sdn_zone",
      "name": "test1",
      "provider": "provider[\"registry.terraform.io/dodai-dodai/proxmox-sdn\"]",
      "instances": [
        {
          "schema_version": 0,
          "attributes": {
            "dns": null,
            "dnszone": null,
            "evpn": null,
            "ipam": null,
            "mtu": null,
            "nodes": null,
            "qinq": null,
            "reversedns": null,
            "type": "vlan",
            "vlan": {
              "bridge": "vmbr0"
            },
            "vxlan": null,
            "zone": "test1"
          },
          "sensitive_attributes": []
        }
      ]
    }
  ],
  "check_results": null
}
```
以下にソースコードを示します。

**client/client.go**
```go
package client

import (
	"bytes"
	"fmt"

	"golang.org/x/crypto/ssh"
)

type SSHProxmoxClient struct {
	client *ssh.Client
}

func NewSSHProxmoxClient(user, password, address string) (*SSHProxmoxClient, error) {
	sshConfig := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", address, sshConfig)
	if err != nil {
		return nil, err
	}

	return &SSHProxmoxClient{client: client}, nil
}

func (c *SSHProxmoxClient) RunCommand(cmd string) (string, error) {
	fmt.Println("Running command:", cmd)
	session, err := c.client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	var stdoutBuf bytes.Buffer
	var stderrBuf bytes.Buffer
	session.Stdout = &stdoutBuf
	session.Stderr = &stderrBuf

	err = session.Run(cmd)
	if err != nil {
		return "", fmt.Errorf("failed to run command: %w\nStderr: %s", err, stderrBuf.String())
	}

	return stdoutBuf.String(), nil
}

func (c *SSHProxmoxClient) Close() {
	c.client.Close()
}
```

**client/types.go**
```go
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
```

**client/zone.go**
```go
package client

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

// CreateSDNZone creates a new SDN zone in Proxmox.
func (c *SSHProxmoxClient) CreateSDNZone(zone SDNZone) error {
	command := fmt.Sprintf(
		"pvesh create /cluster/sdn/zones  --type %s --zone %s",
		zone.Type, zone.Zone, // required fields,
	)

	// common optional fields
	if zone.MTU != nil && *zone.MTU != 0 {
		command += fmt.Sprintf(" --mtu %d", *zone.MTU)
	}

	if zone.Nodes != nil && len(zone.Nodes) > 0 {
		nodes := strings.Join(zone.Nodes, ",")
		command += fmt.Sprintf(" --nodes %s", nodes)
	}

	if zone.IPAM != nil && *zone.IPAM != "" {
		command += fmt.Sprintf(" --ipam %s", *zone.IPAM)
	}

	if zone.DNS != nil && *zone.DNS != "" {
		command += fmt.Sprintf(" --dns %s", *zone.DNS)
	}

	if zone.ReverseDNS != nil && *zone.ReverseDNS != "" {
		command += fmt.Sprintf(" --reversedns %s", *zone.ReverseDNS)
	}

	if zone.DNSZone != nil && *zone.DNSZone != "" {
		command += fmt.Sprintf(" --dnszone %s", *zone.DNSZone)
	}

	switch zone.Type {
	// case "simple":
	// 	if zone.Simple != nil {
	// 		if zone.Simple.AutoDHCP != nil {
	// 			command += fmt.Sprintf(" --auto-dhcp %t", *zone.Simple.AutoDHCP)
	// 		}
	// 	}
	case "vlan":
		if zone.VLAN != nil {
			// required field so no need to check for nil
			command += fmt.Sprintf(" --bridge %s", zone.VLAN.Bridge)
		}
	case "qinq":
		if zone.QinQ != nil {
			// required field so no need to check for nil
			command += fmt.Sprintf(" --bridge %s", zone.QinQ.Bridge)
			command += fmt.Sprintf(" --tag %d", zone.QinQ.Tag)
			if zone.QinQ.VLANProtocol != nil && *zone.QinQ.VLANProtocol != "" {
				command += fmt.Sprintf(" --vlan-protocol %s", *zone.QinQ.VLANProtocol)
			}
		}
	case "vxlan":
		if zone.VXLAN != nil && len(zone.VXLAN.Peer) > 0 {
			peers := strings.Join(zone.VXLAN.Peer, ",")
			command += fmt.Sprintf(" --peer %s", peers)
		} else {
			return fmt.Errorf("peer is required for VXLAN zone")
		}

	case "evpn":
		if zone.EVPN != nil {
			// required field so no need to check for nil
			command += fmt.Sprintf(" --controller %s", zone.EVPN.Controller)
			command += fmt.Sprintf(" --vrf-vxlan %d", zone.EVPN.VRFVXLAN)
			if zone.EVPN.MAC != nil && *zone.EVPN.MAC != "" {
				command += fmt.Sprintf(" --mac %s", *zone.EVPN.MAC)
			}
			if zone.EVPN.ExitNodes != nil && len(zone.EVPN.ExitNodes) > 0 {
				exitnodes := strings.Join(zone.EVPN.ExitNodes, ",")
				command += fmt.Sprintf(" --exitnodes %s", exitnodes)
			}
			if zone.EVPN.PrimaryExitNode != nil && *zone.EVPN.PrimaryExitNode != "" {
				command += fmt.Sprintf(" --primary-exitnode %s", *zone.EVPN.PrimaryExitNode)
			}
			if zone.EVPN.ExitNodesLocalRouting != nil {
				val := "false"
				if *zone.EVPN.ExitNodesLocalRouting {
					val = "true"
				}
				command += fmt.Sprintf(" --exitnodes-local-routing %t", val)
			}
			if zone.EVPN.AdvertiseSubnets != nil {
				val := "false"
				if *zone.EVPN.AdvertiseSubnets {
					val = "true"
				}
				command += fmt.Sprintf(" --advertise-subnets %t", val)
			}
			if zone.EVPN.DisableARPNdSuppression != nil {
				val := "false"
				if *zone.EVPN.DisableARPNdSuppression {
					val = "true"
				}
				command += fmt.Sprintf(" --disable-arp-nd-suppression %t", val)
			}
			if zone.EVPN.RouteTargetImport != nil && *zone.EVPN.RouteTargetImport != "" {
				command += fmt.Sprintf(" --rt-import %s", *zone.EVPN.RouteTargetImport)
			}
		}
	}

	_, err := c.RunCommand(command)
	if err != nil {
		return fmt.Errorf("failed to create SDN zone: %v", err)
	}
	return nil
}

// GetSDNZones retrieves the list of SDN zones from Proxmox.
func (c *SSHProxmoxClient) GetSDNZones() ([]SDNZone, error) {
	command := "pvesh get /cluster/sdn/zones --output-format json"
	output, err := c.RunCommand(command)
	if err != nil {
		return nil, fmt.Errorf("failed to get SDN zones: %v", err)
	}

	var data []map[string]interface{}
	if err := json.Unmarshal([]byte(output), &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
	}

	var zones []SDNZone
	for _, d := range data {
		zone := SDNZone{
			Zone: d["zone"].(string),
			Type: d["type"].(string),
		}
		if mtu, ok := d["mtu"].(float64); ok {
			mtuInt := int64(mtu)
			zone.MTU = &mtuInt
		}
		if nodes, ok := d["nodes"].(string); ok {
			zone.Nodes = strings.Split(nodes, ",")
		}
		if ipam, ok := d["ipam"].(string); ok {
			zone.IPAM = &ipam
		}
		if dns, ok := d["dns"].(string); ok {
			zone.DNS = &dns
		}
		if reverseDNS, ok := d["reversedns"].(string); ok {
			zone.ReverseDNS = &reverseDNS
		}
		if dnsZone, ok := d["dnszone"].(string); ok {
			zone.DNSZone = &dnsZone
		}

		switch zone.Type {
		// case "simple":
		// 	config := &SimpleConfig{}
		// 	if dhcp, ok := (d["dhcp"]).(bool); ok {
		// 		config.AutoDHCP = &dhcp
		// 	}
		// 	zone.Simple = config
		case "vlan":
			config := &VLANConfig{}
			if bridge, ok := (d["bridge"]).(string); ok {
				config.Bridge = bridge
			}
			zone.VLAN = config
		case "qinq":
			config := &QinQConfig{}
			if bridge, ok := (d["bridge"]).(string); ok {
				config.Bridge = bridge
			}
			if tag, ok := (d["tag"]).(float64); ok {
				config.Tag = int64(tag)
			}
			if vlanProtocol, ok := (d["vlan_protocol"]).(string); ok {
				config.VLANProtocol = &vlanProtocol
			}
			zone.QinQ = config
		case "vxlan":
			config := &VXLANConfig{}
			if peer, ok := (d["peer"]).(string); ok {
				config.Peer = strings.Split(peer, ",")
			}
			zone.VXLAN = config
		case "evpn":
			config := &EVPNConfig{}
			if controller, ok := (d["controller"]).(string); ok {
				config.Controller = controller
			}
			if vrfVXLAN, ok := (d["vrf_vxlan"]).(float64); ok {
				config.VRFVXLAN = int64(vrfVXLAN)
			}
			if mac, ok := (d["mac"]).(string); ok {
				config.MAC = &mac
			}
			if exitNodes, ok := (d["exitnodes"]).(string); ok {
				config.ExitNodes = strings.Split(exitNodes, ",")
			}
			if primaryExitNode, ok := (d["primary_exitnode"]).(string); ok {
				config.PrimaryExitNode = &primaryExitNode
			}
			if exitNodesLocalRouting, ok := (d["exitnodes_local_routing"]).(bool); ok {
				config.ExitNodesLocalRouting = &exitNodesLocalRouting
			}
			if advertiseSubnets, ok := (d["advertise_subnets"]).(bool); ok {
				config.AdvertiseSubnets = &advertiseSubnets
			}
			if disableARPNdSuppression, ok := (d["disable_arp_nd_suppression"]).(bool); ok {
				config.DisableARPNdSuppression = &disableARPNdSuppression
			}
			if routeTargetImport, ok := (d["rt_import"]).(string); ok {
				config.RouteTargetImport = &routeTargetImport
			}
			zone.EVPN = config
		}
		zones = append(zones, zone)
	}

	return zones, nil
}

func (c *SSHProxmoxClient) GetSDNZone(zoneID string) (*SDNZone, error) {
	cmd := fmt.Sprintf("pvesh get /cluster/sdn/zones/%s --output-format json", zoneID)
	output, err := c.RunCommand(cmd)
	if err != nil {
		return nil, err
	}

	file, err := os.Create(fmt.Sprintf("%s.json", zoneID))
	if err != nil {
		return nil, fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	var zone SDNZone
	if err := json.Unmarshal([]byte(output), &zone); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	var data map[string]interface{}
	if err := json.Unmarshal([]byte(output), &data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	zone = SDNZone{
		Zone: data["zone"].(string),
		Type: data["type"].(string),
	}

	if mtu, ok := data["mtu"].(float64); ok {
		mtuInt := int64(mtu)
		zone.MTU = &mtuInt
	}

	if nodes, ok := data["nodes"].(string); ok {
		zone.Nodes = strings.Split(nodes, ",")
	}

	if ipam, ok := data["ipam"].(string); ok {
		zone.IPAM = &ipam
	}

	if dns, ok := data["dns"].(string); ok {
		zone.DNS = &dns
	}

	if reverseDNS, ok := data["reversedns"].(string); ok {
		zone.ReverseDNS = &reverseDNS
	}

	if dnsZone, ok := data["dnszone"].(string); ok {
		zone.DNSZone = &dnsZone
	}

	switch zone.Type {
	// case "simple":
	// 	config := &SimpleConfig{}
	// 	if dhcp, ok := (data["dhcp"]).(bool); ok {
	// 		config.AutoDHCP = &dhcp
	// 	}
	// 	zone.Simple = config
	case "vlan":
		config := &VLANConfig{}
		if bridge, ok := (data["bridge"]).(string); ok {
			config.Bridge = bridge
		}
		zone.VLAN = config
	case "qinq":
		config := &QinQConfig{}
		if bridge, ok := (data["bridge"]).(string); ok {
			config.Bridge = bridge
		}
		if tag, ok := (data["tag"]).(float64); ok {
			config.Tag = int64(tag)
		}
		if vlanProtocol, ok := (data["vlan_protocol"]).(string); ok {
			config.VLANProtocol = &vlanProtocol
		}
		zone.QinQ = config
	case "vxlan":
		if peers, ok := (data["peer"]).([]interface{}); ok {
			for _, peer := range peers {
				if peerStr, ok := peer.(string); ok {
					zone.VXLAN.Peer = append(zone.VXLAN.Peer, peerStr)
				}
			}
		}
	case "evpn":
		config := &EVPNConfig{}
		if controller, ok := (data["controller"]).(string); ok {
			config.Controller = controller
		}
		if vrfVXLAN, ok := (data["vrf_vxlan"]).(float64); ok {
			config.VRFVXLAN = int64(vrfVXLAN)
		}
		if mac, ok := (data["mac"]).(string); ok {
			config.MAC = &mac
		}
		if exitNodes, ok := (data["exitnodes"]).(string); ok {
			config.ExitNodes = strings.Split(exitNodes, ",")
		}
		if primaryExitNode, ok := (data["primary_exitnode"]).(string); ok {
			config.PrimaryExitNode = &primaryExitNode
		}
		if exitNodesLocalRouting, ok := (data["exitnodes_local_routing"]).(bool); ok {
			config.ExitNodesLocalRouting = &exitNodesLocalRouting
		}
		if advertiseSubnets, ok := (data["advertise_subnets"]).(bool); ok {
			config.AdvertiseSubnets = &advertiseSubnets
		}
		if disableARPNdSuppression, ok := (data["disable_arp_nd_suppression"]).(bool); ok {
			config.DisableARPNdSuppression = &disableARPNdSuppression
		}
		if routeTargetImport, ok := (data["rt_import"]).(string); ok {
			config.RouteTargetImport = &routeTargetImport
		}
		zone.EVPN = config
	}

	return &zone, nil
}

// UpdateSDNZone updates an existing SDN zone in Proxmox.
func (c *SSHProxmoxClient) UpdateSDNZone(zone SDNZone) error {
	command := fmt.Sprintf("pvesh set /cluster/sdn/zones/%s", zone.Zone)

	if zone.MTU != nil && *zone.MTU != 0 {
		command += fmt.Sprintf(" --mtu %d", *zone.MTU)
	} else if zone.MTU == nil /* フィールドがあり、nilが指定されたら、そのフィールドを削除する*/ {
		command += " --delete mtu"
	}

	if len(zone.Nodes) > 0 && zone.Nodes != nil {
		nodes := strings.Join(zone.Nodes, ",")
		command += fmt.Sprintf(" --nodes %s", nodes)
	} else if zone.Nodes == nil && len(zone.Nodes) == 0 {
		command += " --delete nodes"
	}

	if zone.IPAM != nil && *zone.IPAM != "" {
		command += fmt.Sprintf(" --ipam %s", *zone.IPAM)
	} else if zone.IPAM == nil && *zone.IPAM == "" {
		command += " --delete ipam"
	}

	if zone.DNS != nil && *zone.DNS != "" {
		command += fmt.Sprintf(" --dns %s", *zone.DNS)
	} else if zone.DNS == nil && *zone.DNS == "" {
		command += " --delete dns"
	}

	if zone.ReverseDNS != nil && *zone.ReverseDNS != "" {
		command += fmt.Sprintf(" --reversedns %s", *zone.ReverseDNS)
	} else if zone.ReverseDNS == nil && *zone.ReverseDNS == "" {
		command += " --delete reversedns"
	}

	if zone.DNSZone != nil && *zone.DNSZone != "" {
		command += fmt.Sprintf(" --dnszone %s", *zone.DNSZone)
	} else if zone.DNSZone == nil && *zone.DNSZone == "" {
		command += " --delete dnszone"
	}

	switch zone.Type {
	// case "simple":
	// 	if zone.Simple != nil {
	// 		if zone.Simple.AutoDHCP != nil {
	// 			command += fmt.Sprintf(" --auto-dhcp %t", *zone.Simple.AutoDHCP)
	// 		}
	// 	}
	case "vlan":
		if zone.VLAN != nil {
			command += fmt.Sprintf(" --bridge %s", zone.VLAN.Bridge)
		}

	case "qinq":
		if zone.QinQ != nil {
			command += fmt.Sprintf(" --bridge %s", zone.QinQ.Bridge)
			command += fmt.Sprintf(" --tag %d", zone.QinQ.Tag)
			if zone.QinQ.VLANProtocol != nil && *zone.QinQ.VLANProtocol != "" {
				command += fmt.Sprintf(" --vlan-protocol %s", *zone.QinQ.VLANProtocol)
			} else if zone.QinQ.VLANProtocol == nil && *zone.QinQ.VLANProtocol == "" {
				command += " --delete vlan-protocol"
			}
		}
	case "vxlan":
		if zone.VXLAN != nil && len(zone.VXLAN.Peer) > 0 {
			peers := strings.Join(zone.VXLAN.Peer, ",")
			command += fmt.Sprintf(" --peer %s", peers)
		} else {
			command += " --delete peer"
		}
	case "evpn":
		if zone.EVPN != nil {
			command += fmt.Sprintf(" --controller %s", zone.EVPN.Controller)
			command += fmt.Sprintf(" --vrf-vxlan %d", zone.EVPN.VRFVXLAN)

			if zone.EVPN.MAC != nil && *zone.EVPN.MAC != "" {
				command += fmt.Sprintf(" --mac %s", *zone.EVPN.MAC)
			} else if zone.EVPN.MAC == nil && *zone.EVPN.MAC == "" {
				command += " --delete mac"
			}

			if len(zone.EVPN.ExitNodes) > 0 && zone.EVPN.ExitNodes != nil {
				exitNodes := strings.Join(zone.EVPN.ExitNodes, ",")
				command += fmt.Sprintf(" --exitnodes %s", exitNodes)
			} else if zone.EVPN.ExitNodes == nil && len(zone.EVPN.ExitNodes) == 0 {
				command += " --delete exitnodes"
			}

			if zone.EVPN.PrimaryExitNode != nil && *zone.EVPN.PrimaryExitNode != "" {
				command += fmt.Sprintf(" --primary-exitnode %s", *zone.EVPN.PrimaryExitNode)
			} else if zone.EVPN.PrimaryExitNode == nil && *zone.EVPN.PrimaryExitNode == "" {
				command += " --delete primary-exitnode"
			}

			if zone.EVPN.ExitNodesLocalRouting != nil {
				val := "false"
				if *zone.EVPN.ExitNodesLocalRouting {
					val = "true"
				}
				command += fmt.Sprintf(" --exitnodes-local-routing %s", val)
			} else if zone.EVPN.ExitNodesLocalRouting == nil {
				command += " --delete exitnodes-local-routing"
			}

			if zone.EVPN.AdvertiseSubnets != nil {
				val := "false"
				if *zone.EVPN.AdvertiseSubnets {
					val = "true"
				}
				command += fmt.Sprintf(" --advertise-subnets %s", val)
			} else if zone.EVPN.AdvertiseSubnets == nil {
				command += " --delete advertise-subnets"
			}

			if zone.EVPN.DisableARPNdSuppression != nil {
				val := "false"
				if *zone.EVPN.DisableARPNdSuppression {
					val = "true"
				}
				command += fmt.Sprintf(" --disable-arp-nd-suppression %s", val)
			} else if zone.EVPN.DisableARPNdSuppression == nil {
				command += " --delete disable-arp-nd-suppression"
			}

			if zone.EVPN.RouteTargetImport != nil && *zone.EVPN.RouteTargetImport != "" {
				command += fmt.Sprintf(" --rt-import %s", *zone.EVPN.RouteTargetImport)
			} else if zone.EVPN.RouteTargetImport == nil && *zone.EVPN.RouteTargetImport == "" {
				command += " --delete rt-import"
			}
		}
	}

	_, err := c.RunCommand(command)
	if err != nil {
		return fmt.Errorf("failed to update SDN zone: %v", err)
	}
	return nil
}

// DeleteSDNZone deletes an existing SDN zone in Proxmox.
func (c *SSHProxmoxClient) DeleteSDNZone(zoneID string) error {
	command := fmt.Sprintf("pvesh delete /cluster/sdn/zones/%s", zoneID)
	_, err := c.RunCommand(command)
	if err != nil {
		return fmt.Errorf("failed to delete SDN zone: %v", err)
	}
	return nil
}
```

**provider/provider.go**
```go
package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/Dodai-Dodai/terraform-provider-proxmox-sdn/client"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
)

var (
	_ provider.Provider = &ProxmoxSDNProvider{}
)

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ProxmoxSDNProvider{
			version: version,
		}
	}
}

type ProxmoxSDNProvider struct {
	version string
}

func (p *ProxmoxSDNProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "proxmox-sdn"
	resp.Version = p.version
}

func (p *ProxmoxSDNProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Description: "The Proxmox host",
				Required:    true,
			},
			"username": schema.StringAttribute{
				Description: "The Proxmox username",
				Required:    true,
			},
			"password": schema.StringAttribute{
				Description: "The Proxmox password",
				Required:    true,
				Sensitive:   true,
			},
		},
	}
}

type ProxmoxSDNProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func (p *ProxmoxSDNProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring ProxmoxSDNProvider Client")
	var config ProxmoxSDNProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "proxmoSDNHost", config.Host)
	ctx = tflog.SetField(ctx, "proxmoSDNUsername", config.Username)
	ctx = tflog.SetField(ctx, "proxmoSDNPassword", config.Password)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "proxmoSDNPassword")

	tflog.Debug(ctx, "Creating ProxmoxSDNProvider Client")

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown Proxmox host",
			"The Provider cannot create the client without a host",
		)
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown Proxmox username",
			"The Provider cannot create the client without a username",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown Proxmox password",
			"The Provider cannot create the client without a password",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	host := os.Getenv("PROXMOXSDN_HOST")
	username := os.Getenv("PROXMOXSDN_USERNAME")
	password := os.Getenv("PROXMOXSDN_PASSWORD")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Proxmox host",
			"The Provider cannot create the client without a host"+
				"Set the PROXMOXSDN_HOST environment variable or provide the host attribute"+
				"if either is already set, ensure that the value is not empty",
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing Proxmox username",
			"The Provider cannot create the client without a username"+
				"Set the PROXMOXSDN_USERNAME environment variable or provide the username attribute"+
				"if either is already set, ensure that the value is not empty",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Proxmox password",
			"The Provider cannot create the client without a password"+
				"Set the PROXMOXSDN_PASSWORD environment variable or provide the password attribute"+
				"if either is already set, ensure that the value is not empty",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	client, err := client.NewSSHProxmoxClient(username, password, host)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to create client",
			fmt.Sprintf("An error occurred while creating the client: %s", err.Error()),
		)
		return
	}

	resp.DataSourceData = client
	resp.ResourceData = client

	tflog.Info(ctx, "Configured proxmox-sdn client", map[string]any{"success": true})
}

func (p *ProxmoxSDNProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewProxmoxSDNZoneDataSource,
	}
}

func (p *ProxmoxSDNProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewProxmoxSDNZoneResource,
	}
}
```

**resource/types.go**
```go
package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type zonesModel struct {
	Zone       types.String `tfsdk:"zone"`       //required
	Type       types.String `tfsdk:"type"`       //required
	MTU        types.Int64  `tfsdk:"mtu"`        //optional
	Nodes      types.Set    `tfsdk:"nodes"`      //optional
	IPAM       types.String `tfsdk:"ipam"`       //optional
	DNS        types.String `tfsdk:"dns"`        //optional
	ReverseDNS types.String `tfsdk:"reversedns"` //optional
	DNSZone    types.String `tfsdk:"dnszone"`    //optional
	VLAN       types.Object `tfsdk:"vlan"`       //optional
	QinQ       types.Object `tfsdk:"qinq"`       //optional
	VXLAN      types.Object `tfsdk:"vxlan"`      //optional
	EVPN       types.Object `tfsdk:"evpn"`       //optional
}

type VLANConfigModel struct {
	Bridge types.String `tfsdk:"bridge"` //required
}

type QinQConfigModel struct {
	Bridge       types.String `tfsdk:"bridge"`       //required
	Tag          types.Int64  `tfsdk:"tag"`          //required
	VLANProtocol types.String `tfsdk:"vlanprotocol"` //optional
}

type VXLANConfigModel struct {
	Peer types.Set `tfsdk:"peer"` //required
}

type EVPNConfigModel struct {
	Controller              types.String `tfsdk:"controller"`              //required
	VRFVXLAN                types.Int64  `tfsdk:"vrf_vxlan"`               //required
	MAC                     types.String `tfsdk:"mac"`                     //optional
	ExitNodes               types.Set    `tfsdk:"exitnodes"`               //optional
	PrimaryExitNode         types.String `tfsdk:"primaryexitnode"`         //optional
	ExitNodesLocalRouting   types.Bool   `tfsdk:"exitnodeslocalrouting"`   //optional
	AdvertiseSubnets        types.Bool   `tfsdk:"advertisesubnets"`        //optional
	DisableARPNdSuppression types.Bool   `tfsdk:"disablearpndsuppression"` //optional
	RouteTargetImport       types.String `tfsdk:"rtimport"`                //optional
}
```

**resource/zone_resource.go**
```go
package provider

import (
	"context"
	"fmt"

	"github.com/Dodai-Dodai/terraform-provider-proxmox-sdn/client"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

//type proxmoxSDNZoneResourceModel = zonesModel

var (
	_ resource.Resource              = &proxmoxSDNZoneResource{}
	_ resource.ResourceWithConfigure = &proxmoxSDNZoneResource{}
)

func NewProxmoxSDNZoneResource() resource.Resource {
	return &proxmoxSDNZoneResource{}
}

type proxmoxSDNZoneResource struct {
	client *client.SSHProxmoxClient
}

func (r *proxmoxSDNZoneResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.SSHProxmoxClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got %T", req.ProviderData),
		)
		return
	}
	r.client = client
}

func (r *proxmoxSDNZoneResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zone"
}

func (r *proxmoxSDNZoneResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zone": schema.StringAttribute{
				Description: "The name of the zone",
				Required:    true,
				// 変更されたらリソースを再作成する
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"type": schema.StringAttribute{
				Description: "The type of the zone",
				Required:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"mtu": schema.Int64Attribute{
				Description: "The MTU Num of the Zone",
				Optional:    true,
				Computed:    true,
			},
			"nodes": schema.SetAttribute{
				Description: "Set of nodes in the zone",
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
			},
			"ipam": schema.StringAttribute{
				Description: "The IPAM of the zone",
				Optional:    true,
				Computed:    true,
			},
			"dns": schema.StringAttribute{
				Description: "The DNS of the zone",
				Optional:    true,
				Computed:    true,
			},
			"reversedns": schema.StringAttribute{
				Description: "The reverse dns of the zone",
				Optional:    true,
				Computed:    true,
			},
			"dnszone": schema.StringAttribute{
				Description: "The DNS zone of the zone",
				Optional:    true,
				Computed:    true,
			},

			// VLAN Config
			"vlan": schema.SingleNestedAttribute{
				Description: "The VLAN configuration of the zone",
				Optional:    true,
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"bridge": schema.StringAttribute{
						Description: "The bridge of the VLAN",
						Required:    true,
					},
				},
			},

			// QinQ Config
			"qinq": schema.SingleNestedAttribute{
				Description: "The QinQ configuration of the zone",
				Optional:    true,
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"bridge": schema.StringAttribute{
						Description: "The bridge of the QinQ",
						Required:    true,
					},
					"tag": schema.Int64Attribute{
						Description: "The tag of the QinQ",
						Required:    true,
					},
					"vlanprotocol": schema.StringAttribute{
						Description: "The VLAN protocol of the QinQ",
						Optional:    true,
						Computed:    true,
					},
				},
			},

			// VXLAN Config
			"vxlan": schema.SingleNestedAttribute{
				Description: "The VXLAN configuration of the zone",
				Optional:    true,
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"peer": schema.SetAttribute{
						Description: "The peer of the VXLAN",
						ElementType: types.StringType,
						Required:    true,
					},
				},
			},

			// EVPN Config
			"evpn": schema.SingleNestedAttribute{
				Description: "The EVPN configuration of the zone",
				Optional:    true,
				Computed:    true,
				Attributes: map[string]schema.Attribute{
					"controller": schema.StringAttribute{
						Description: "The controller of the EVPN",
						Required:    true,
					},
					"vrf_vxlan": schema.Int64Attribute{
						Description: "The VRF VXLAN of the EVPN",
						Required:    true,
					},
					"mac": schema.StringAttribute{
						Description: "The MAC of the EVPN",
						Optional:    true,
						Computed:    true,
					},
					"exitnodes": schema.SetAttribute{
						Description: "The exit nodes of the EVPN",
						ElementType: types.StringType,
						Optional:    true,
						Computed:    true,
					},
					"primaryexitnode": schema.StringAttribute{
						Description: "The primary exit node of the EVPN",
						Optional:    true,
						Computed:    true,
					},
					"exitnodeslocalrouting": schema.BoolAttribute{
						Description: "The exit nodes local routing of the EVPN",
						Optional:    true,
						Computed:    true,
					},
					"advertisesubnets": schema.BoolAttribute{
						Description: "The advertise subnets of the EVPN",
						Optional:    true,
						Computed:    true,
					},
					"disablearpndsuppression": schema.BoolAttribute{
						Description: "The disable arp nd suppression of the EVPN",
						Optional:    true,
						Computed:    true,
					},
					"rtimport": schema.StringAttribute{
						Description: "The route import of the EVPN",
						Optional:    true,
						Computed:    true,
					},
				},
			},
		},
	}
}

func (r *proxmoxSDNZoneResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan zonesModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// planの構造体をclientの構造体に変換
	zone, diagsConvert := convertZonesModeltoClientSDNZone(ctx, plan)
	resp.Diagnostics.Append(diagsConvert...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.CreateSDNZone(*zone)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to create SDN zone",
			err.Error(),
		)
		return
	}

	createdZone, err := r.client.GetSDNZone(zone.Zone)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to get created SDN zone",
			err.Error(),
		)
		return
	}

	state, diagsState := convertSDNZoneToZonesModel(ctx, *createdZone)
	resp.Diagnostics.Append(diagsState...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, state)
	resp.Diagnostics.Append(diags...)
}

func convertZonesModeltoClientSDNZone(ctx context.Context, model zonesModel) (*client.SDNZone, diag.Diagnostics) {
	var diags diag.Diagnostics

	zone := &client.SDNZone{
		Zone: model.Zone.ValueString(),
		Type: model.Type.ValueString(),
	}

	if !model.MTU.IsNull() && !model.MTU.IsUnknown() {
		mtu := model.MTU.ValueInt64()
		zone.MTU = &mtu
	}

	if !model.Nodes.IsNull() && !model.Nodes.IsUnknown() {
		var nodes []string
		diagNodes := model.Nodes.ElementsAs(ctx, &nodes, false)
		diags.Append(diagNodes...)
		if diagNodes.HasError() {
			return nil, diags
		}
		zone.Nodes = nodes
	}

	if !model.IPAM.IsNull() && !model.IPAM.IsUnknown() {
		ipam := model.IPAM.ValueString()
		zone.IPAM = &ipam
	}

	if !model.DNS.IsNull() && !model.DNS.IsUnknown() {
		dns := model.DNS.ValueString()
		zone.DNS = &dns
	}

	if !model.ReverseDNS.IsNull() && !model.ReverseDNS.IsUnknown() {
		reverseDNS := model.ReverseDNS.ValueString()
		zone.ReverseDNS = &reverseDNS
	}

	if !model.DNSZone.IsNull() && !model.DNSZone.IsUnknown() {
		dnsZone := model.DNSZone.ValueString()
		zone.DNSZone = &dnsZone
	}

	if !model.VLAN.IsNull() && !model.VLAN.IsUnknown() {
		var vlanConfig VLANConfigModel
		diags := model.VLAN.As(ctx, &vlanConfig, basetypes.ObjectAsOptions{})
		if diags.HasError() {
			return nil, diags
		}
		clientVLANConfig := &client.VLANConfig{
			Bridge: vlanConfig.Bridge.ValueString(),
		}
		zone.VLAN = clientVLANConfig
	}

	if !model.QinQ.IsNull() && !model.QinQ.IsUnknown() {
		var qinqConfig QinQConfigModel
		diags := model.QinQ.As(ctx, &qinqConfig, basetypes.ObjectAsOptions{})
		if diags.HasError() {
			return nil, diags
		}
		clientQinQConfig := &client.QinQConfig{
			Bridge: qinqConfig.Bridge.ValueString(),
			Tag:    qinqConfig.Tag.ValueInt64(),
		}
		if !qinqConfig.VLANProtocol.IsNull() {
			vlanProtocol := qinqConfig.VLANProtocol.ValueString()
			clientQinQConfig.VLANProtocol = &vlanProtocol
		}
		zone.QinQ = clientQinQConfig
	}

	if !model.VXLAN.IsNull() && !model.VXLAN.IsUnknown() {
		var vxlanConfig VXLANConfigModel
		diags := model.VXLAN.As(ctx, &vxlanConfig, basetypes.ObjectAsOptions{})
		if diags.HasError() {
			return nil, diags
		}
		var peers []string
		diagsPeers := vxlanConfig.Peer.ElementsAs(ctx, &peers, false)
		diags.Append(diagsPeers...)
		if diagsPeers.HasError() {
			return nil, diags
		}
		clientVXLANConfig := &client.VXLANConfig{
			Peer: peers,
		}
		zone.VXLAN = clientVXLANConfig
	}

	if !model.EVPN.IsNull() && !model.EVPN.IsUnknown() {
		var evpnConfig EVPNConfigModel
		diags := model.EVPN.As(ctx, &evpnConfig, basetypes.ObjectAsOptions{})
		if diags.HasError() {
			return nil, diags
		}
		clientEVPNConfig := &client.EVPNConfig{
			Controller: evpnConfig.Controller.ValueString(),
			VRFVXLAN:   evpnConfig.VRFVXLAN.ValueInt64(),
		}
		if !evpnConfig.MAC.IsNull() {
			mac := evpnConfig.MAC.ValueString()
			clientEVPNConfig.MAC = &mac
		}
		if !evpnConfig.PrimaryExitNode.IsNull() {
			primaryExitNode := evpnConfig.PrimaryExitNode.ValueString()
			clientEVPNConfig.PrimaryExitNode = &primaryExitNode
		}
		if !evpnConfig.ExitNodesLocalRouting.IsNull() {
			exitNodesLocalRouting := evpnConfig.ExitNodesLocalRouting.ValueBool()
			clientEVPNConfig.ExitNodesLocalRouting = &exitNodesLocalRouting
		}
		if !evpnConfig.AdvertiseSubnets.IsNull() {
			advertiseSubnets := evpnConfig.AdvertiseSubnets.ValueBool()
			clientEVPNConfig.AdvertiseSubnets = &advertiseSubnets
		}
		if !evpnConfig.DisableARPNdSuppression.IsNull() {
			disableARPNdSuppression := evpnConfig.DisableARPNdSuppression.ValueBool()
			clientEVPNConfig.DisableARPNdSuppression = &disableARPNdSuppression
		}
		if !evpnConfig.RouteTargetImport.IsNull() {
			routeTargetImport := evpnConfig.RouteTargetImport.ValueString()
			clientEVPNConfig.RouteTargetImport = &routeTargetImport
		}
		var exitNodes []string
		diagsExitNodes := evpnConfig.ExitNodes.ElementsAs(ctx, &exitNodes, false)
		diags.Append(diagsExitNodes...)
		if diagsExitNodes.HasError() {
			return nil, diags
		}
		clientEVPNConfig.ExitNodes = exitNodes
		zone.EVPN = clientEVPNConfig
	}
	return zone, diags
}

// Read refreshes the Terraform state with the latest data.
func (r *proxmoxSDNZoneResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state zonesModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zone, err := r.client.GetSDNZone(state.Zone.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to get SDN zone",
			err.Error(),
		)
		return
	}

	updatedState, diagsState := convertSDNZoneToZonesModel(ctx, *zone)
	resp.Diagnostics.Append(diagsState...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, updatedState)
	resp.Diagnostics.Append(diags...)
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *proxmoxSDNZoneResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan zonesModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	zone, diagsConvert := convertZonesModeltoClientSDNZone(ctx, plan)
	resp.Diagnostics.Append(diagsConvert...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.UpdateSDNZone(*zone)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to update SDN zone",
			err.Error(),
		)
		return
	}

	updatedZone, err := r.client.GetSDNZone(zone.Zone)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to get updated SDN zone",
			err.Error(),
		)
		return
	}

	updatedSteta, diagsState := convertSDNZoneToZonesModel(ctx, *updatedZone)
	resp.Diagnostics.Append(diagsState...)
	if resp.Diagnostics.HasError() {
		return
	}

	diags = resp.State.Set(ctx, updatedSteta)
	resp.Diagnostics.Append(diags...)
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *proxmoxSDNZoneResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state zonesModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	err := r.client.DeleteSDNZone(state.Zone.ValueString())
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to delete SDN zone",
			err.Error(),
		)
		return
	}

}
```

**resource/zone_datasource.go**
```go
package provider

import (
	"context"
	"fmt"

	"github.com/Dodai-Dodai/terraform-provider-proxmox-sdn/client"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ datasource.DataSource              = &proxmoxSDNZoneDataSource{}
	_ datasource.DataSourceWithConfigure = &proxmoxSDNZoneDataSource{}
)

func NewProxmoxSDNZoneDataSource() datasource.DataSource {
	return &proxmoxSDNZoneDataSource{}
}

type proxmoxSDNZoneDataSource struct {
	client *client.SSHProxmoxClient
}

func (d *proxmoxSDNZoneDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_zone"
}

func (d *proxmoxSDNZoneDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"zones": schema.ListNestedAttribute{
				Description: "List of SDN zones",
				Computed:    true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"zone": schema.StringAttribute{
							Description: "The name of the zone",
							Required:    true,
						},
						"type": schema.StringAttribute{
							Description: "The type of the zone",
							Required:    true,
						},
						"mtu": schema.Int64Attribute{
							Description: "The MTU Num of the Zone",
							Optional:    true,
							Computed:    true,
						},
						"nodes": schema.SetAttribute{
							Description: "Set of nodes in the zone",
							ElementType: types.StringType,
							Optional:    true,
							Computed:    true,
						},
						"ipam": schema.StringAttribute{
							Description: "The IPAM of the zone",
							Optional:    true,
							Computed:    true,
						},
						"dns": schema.StringAttribute{
							Description: "The DNS of the zone",
							Optional:    true,
							Computed:    true,
						},
						"reversedns": schema.StringAttribute{
							Description: "The reverse dns of the zone",
							Optional:    true,
							Computed:    true,
						},
						"dnszone": schema.StringAttribute{
							Description: "The DnsZone of the zone",
							Optional:    true,
							Computed:    true,
						},

						// VLAN Config
						"vlan": schema.SingleNestedAttribute{
							Description: "The VLAN configuration of the zone",
							Optional:    true,
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"bridge": schema.StringAttribute{
									Description: "The bridge of VLAN zone",
									Required:    true,
								},
							},
						},

						// QinQ Config
						"qinq": schema.SingleNestedAttribute{
							Description: "The QinQ configuration of the zone",
							Optional:    true,
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"bridge": schema.StringAttribute{
									Description: "The bridge of QinQ zone",
									Required:    true,
								},
								"tag": schema.Int64Attribute{
									Description: "The tag num of QinQ zone",
									Required:    true,
								},
								"vlanprotocol": schema.StringAttribute{
									Description: "VLAN Protocol of the zone",
									Optional:    true,
									Computed:    true,
								},
							},
						},

						// VXLAN Config
						"vxlan": schema.SingleNestedAttribute{
							Description: "The VXLAN configuration of the zone",
							Optional:    true,
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"peer": schema.SetAttribute{
									Description: "Set of Peers in the vxlan zone",
									ElementType: types.StringType,
									Required:    true,
								},
							},
						},

						// EVPN Config
						"evpn": schema.SingleNestedAttribute{
							Description: "The EVPN configuration of the zone",
							Optional:    true,
							Computed:    true,
							Attributes: map[string]schema.Attribute{
								"controller": schema.StringAttribute{
									Description: "The controller of EVPN zone",
									Required:    true,
								},
								"vrf_vxlan": schema.Int64Attribute{
									Description: "The VRFVXLAN of EVPN zone",
									Required:    true,
								},
								"mac": schema.StringAttribute{
									Description: "The MAC of EVPN zone",
									Optional:    true,
									Computed:    true,
								},
								"exitnodes": schema.SetAttribute{
									Description: "Set of ExitNodes in the EVPN zone",
									ElementType: types.StringType,
									Optional:    true,
									Computed:    true,
								},
								"primaryexitnode": schema.StringAttribute{
									Description: "The PrimaryExitNode of EVPN zone",
									Optional:    true,
									Computed:    true,
								},
								"exitnodeslocalrouting": schema.BoolAttribute{
									Description: "The ExitNodesLocalRouting of EVPN zone",
									Optional:    true,
									Computed:    true,
								},
								"advertisesubnets": schema.BoolAttribute{
									Description: "The AdvertiseSubnets of EVPN zone",
									Optional:    true,
									Computed:    true,
								},
								"disablearpndsuppression": schema.BoolAttribute{
									Description: "The DisableARPNdSuppression of EVPN zone",
									Optional:    true,
									Computed:    true,
								},
								"rtimport": schema.StringAttribute{
									Description: "The RouteTargetImport of EVPN zone",
									Optional:    true,
									Computed:    true,
								},
							},
						},
					},
				},
			},
		},
	}
}

func (d *proxmoxSDNZoneDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.SSHProxmoxClient)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.SSHProxmoxClient, got %T", req.ProviderData),
		)
		return
	}
	d.client = client
}

func convertSDNZoneToZonesModel(ctx context.Context, zone client.SDNZone) (zonesModel, diag.Diagnostics) {
	var diags diag.Diagnostics

	nodeset, diagNodes := types.SetValueFrom(ctx, types.StringType, zone.Nodes)
	if diagNodes.HasError() {
		diags.Append(diagNodes...)
		return zonesModel{}, diags
	}
	diags.Append(diagNodes...)

	zoneModel := zonesModel{
		Zone:       types.StringValue(zone.Zone),
		Type:       types.StringValue(zone.Type),
		MTU:        types.Int64PointerValue(zone.MTU),
		Nodes:      nodeset,
		IPAM:       types.StringPointerValue(zone.IPAM),
		DNS:        types.StringPointerValue(zone.DNS),
		ReverseDNS: types.StringPointerValue(zone.ReverseDNS),
		DNSZone:    types.StringPointerValue(zone.DNSZone),
	}

	if zone.VLAN != nil {
		vlanModel := VLANConfigModel{
			Bridge: types.StringValue(zone.VLAN.Bridge),
		}

		vlanObject, diag := types.ObjectValueFrom(ctx, map[string]attr.Type{
			"bridge": types.StringType,
		}, vlanModel)
		if diag.HasError() {
			diags.Append(diag...)
			return zoneModel, diags
		}

		zoneModel.VLAN = vlanObject
	} else {
		// VLANがnilの場合、nullのObjectを作成
		zoneModel.VLAN = types.ObjectNull(map[string]attr.Type{
			"bridge": types.StringType,
		})
	}

	if zone.QinQ != nil {
		qinqModel := QinQConfigModel{
			Bridge: types.StringValue(zone.QinQ.Bridge),
			Tag:    types.Int64Value(zone.QinQ.Tag),
		}
		if zone.QinQ.VLANProtocol != nil {
			qinqModel.VLANProtocol = types.StringPointerValue(zone.QinQ.VLANProtocol)
		}

		qinqObject, diag := types.ObjectValueFrom(ctx, map[string]attr.Type{
			"bridge":       types.StringType,
			"tag":          types.Int64Type,
			"vlanprotocol": types.StringType,
		}, qinqModel)
		if diag.HasError() {
			diags.Append(diag...)
			return zoneModel, diags
		}

		zoneModel.QinQ = qinqObject
	} else {
		// QinQがnilの場合、nullのObjectを作成
		zoneModel.QinQ = types.ObjectNull(map[string]attr.Type{
			"bridge":       types.StringType,
			"tag":          types.Int64Type,
			"vlanprotocol": types.StringType,
		})
	}

	if zone.VXLAN != nil {
		peerSet, diagPeers := types.SetValueFrom(ctx, types.StringType, zone.VXLAN.Peer)
		if diagPeers.HasError() {
			diags.Append(diagPeers...)
			return zoneModel, diags
		}

		vxlanModel := VXLANConfigModel{
			Peer: peerSet,
		}

		vxlanObject, diag := types.ObjectValueFrom(ctx, map[string]attr.Type{
			"peer": types.SetType{ElemType: types.StringType},
		}, vxlanModel)
		if diag.HasError() {
			diags.Append(diag...)
			return zoneModel, diags
		}

		zoneModel.VXLAN = vxlanObject
	} else {
		// VXLANがnilの場合、nullのObjectを作成
		zoneModel.VXLAN = types.ObjectNull(map[string]attr.Type{
			"peer": types.SetType{ElemType: types.StringType},
		})
	}

	if zone.EVPN != nil {
		var exitNodesSet types.Set
		if zone.EVPN.ExitNodes != nil {
			var diagExitNodes diag.Diagnostics
			exitNodesSet, diagExitNodes = types.SetValueFrom(ctx, types.StringType, zone.EVPN.ExitNodes)
			if diagExitNodes.HasError() {
				diags.Append(diagExitNodes...)
				return zoneModel, diags
			}
		} else {
			// ExitNodesがnilの場合、nullのSetを作成
			exitNodesSet = types.SetNull(types.StringType)
		}

		evpnModel := EVPNConfigModel{
			Controller: types.StringValue(zone.EVPN.Controller),
			VRFVXLAN:   types.Int64Value(zone.EVPN.VRFVXLAN),
			ExitNodes:  exitNodesSet,
		}
		if zone.EVPN.MAC != nil {
			evpnModel.MAC = types.StringPointerValue(zone.EVPN.MAC)
		}
		if zone.EVPN.PrimaryExitNode != nil {
			evpnModel.PrimaryExitNode = types.StringPointerValue(zone.EVPN.PrimaryExitNode)
		}
		if zone.EVPN.ExitNodesLocalRouting != nil {
			evpnModel.ExitNodesLocalRouting = types.BoolPointerValue(zone.EVPN.ExitNodesLocalRouting)
		}
		if zone.EVPN.AdvertiseSubnets != nil {
			evpnModel.AdvertiseSubnets = types.BoolPointerValue(zone.EVPN.AdvertiseSubnets)
		}
		if zone.EVPN.DisableARPNdSuppression != nil {
			evpnModel.DisableARPNdSuppression = types.BoolPointerValue(zone.EVPN.DisableARPNdSuppression)
		}
		if zone.EVPN.RouteTargetImport != nil {
			evpnModel.RouteTargetImport = types.StringPointerValue(zone.EVPN.RouteTargetImport)
		}

		evpnObject, diag := types.ObjectValueFrom(ctx, map[string]attr.Type{
			"controller":              types.StringType,
			"vrf_vxlan":               types.Int64Type,
			"mac":                     types.StringType,
			"exitnodes":               types.SetType{ElemType: types.StringType},
			"primaryexitnode":         types.StringType,
			"exitnodeslocalrouting":   types.BoolType,
			"advertisesubnets":        types.BoolType,
			"disablearpndsuppression": types.BoolType,
			"rtimport":                types.StringType,
		}, evpnModel)
		if diag.HasError() {
			diags.Append(diag...)
			return zoneModel, diags
		}

		zoneModel.EVPN = evpnObject
	} else {
		// EVPNがnilの場合、nullのObjectを作成
		zoneModel.EVPN = types.ObjectNull(map[string]attr.Type{
			"controller":              types.StringType,
			"vrf_vxlan":               types.Int64Type,
			"mac":                     types.StringType,
			"exitnodes":               types.SetType{ElemType: types.StringType},
			"primaryexitnode":         types.StringType,
			"exitnodeslocalrouting":   types.BoolType,
			"advertisesubnets":        types.BoolType,
			"disablearpndsuppression": types.BoolType,
			"rtimport":                types.StringType,
		})
	}

	return zoneModel, diags

}

func (d *proxmoxSDNZoneDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state struct {
		Zones []zonesModel `tfsdk:"zones"`
	}

	zones, err := d.client.GetSDNZones()
	if err != nil {
		resp.Diagnostics.AddError("Error getting zones", err.Error())
		return
	}

	var diags diag.Diagnostics

	for _, zone := range zones {
		zoneModel, diagZone := convertSDNZoneToZonesModel(ctx, zone)
		if diagZone.HasError() {
			resp.Diagnostics.Append(diagZone...)
			continue
		}
		diags.Append(diagZone...)
		state.Zones = append(state.Zones, zoneModel)
	}

	if diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	diags.Append(resp.State.Set(ctx, state)...)
	resp.Diagnostics.Append(diags...)
}
```
