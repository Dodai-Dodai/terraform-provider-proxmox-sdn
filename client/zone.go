package client

import (
	"encoding/json"
	"fmt"
	"strings"
)

// CreateSDNZone creates a new SDN zone in Proxmox.
func (c *SSHProxmoxClient) CreateSDNZone(zone SDNZone) error {
	command := fmt.Sprintf(
		"pvesh create /cluster/sdn/zones  --type %s --zone %s",
		zone.Type, zone.Zone, // required fields,
	)

	if zone.MTU != nil && *zone.MTU != 0 {
		command += fmt.Sprintf(" --mtu %d", *zone.MTU)
	}

	if len(zone.Nodes) > 0 {
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

	case "vlan":
		command += fmt.Sprintf(" --bridge %s", *zone.Bridge)

	case "qinq":
		command += fmt.Sprintf(" --bridge %s", *zone.Bridge)
		command += fmt.Sprintf(" --tag %d", *zone.Tag)
		if zone.VLANProtocol != nil && *zone.VLANProtocol != "" {
			command += fmt.Sprintf(" --vlan-protocol %s", *zone.VLANProtocol)
		}

	case "vxlan":
		command += fmt.Sprintf(" --peers %s", strings.Join(zone.Peers, ","))

	case "evpn":
		command += fmt.Sprintf(" --controller %s", *zone.Controller)
		command += fmt.Sprintf(" --vrf-vxlan %d", *zone.VRFVXLAN)

		if zone.MAC != nil && *zone.MAC != "" {
			command += fmt.Sprintf(" --mac %s", *zone.MAC)
		}

		if len(zone.ExitNodes) > 0 {
			exitNodes := strings.Join(zone.ExitNodes, ",")
			command += fmt.Sprintf(" --exitnodes %s", exitNodes)
		}

		if zone.PrimaryExitNode != nil && *zone.PrimaryExitNode != "" {
			command += fmt.Sprintf(" --primary-exitnode %s", *zone.PrimaryExitNode)
		}

		if zone.ExitNodesLocalRouting != nil {
			val := "false"
			if *zone.ExitNodesLocalRouting {
				val = "true"
			}
			command += fmt.Sprintf(" --exitnodes-local-routing %s", val)
		}

		if zone.AdvertiseSubnets != nil {
			val := "false"
			if *zone.AdvertiseSubnets {
				val = "true"
			}
			command += fmt.Sprintf(" --advertise-subnets %s", val)
		}

		if zone.DisableARPNdSuppression != nil {
			val := "false"
			if *zone.DisableARPNdSuppression {
				val = "true"
			}
			command += fmt.Sprintf(" --disable-arp-nd-suppression %s", val)
		}

		if zone.RouteTargetImport != nil && *zone.RouteTargetImport != "" {
			command += fmt.Sprintf(" --rt-import %s", *zone.RouteTargetImport)
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
	// pvesh get /cluster/sdn/zones --output-format json
	cmd := "pvesh get /cluster/sdn/zones --output-format json"
	output, err := c.RunCommand(cmd)
	if err != nil {
		return nil, err
	}
	//[{"digest":"6eeeb0e5517fe5ba4a3a86a6203dfe445c0929b5","nodes":"pvewata01,pvewata02","peers":"150.89.233.20,150.89.233.21","type":"vxlan","zone":"test2"}]
	// 上記のようなJSON文字列をパースする
	var zones []SDNZone
	if err := json.Unmarshal([]byte(output), &zones); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}
	// debug
	fmt.Println(zones)
	return zones, nil
}

func (c *SSHProxmoxClient) GetSDNZone(zoneID string) (*SDNZone, error) {
	cmd := fmt.Sprintf("pvesh get /cluster/sdn/zones/%s --output-format json", zoneID)
	output, err := c.RunCommand(cmd)
	if err != nil {
		return nil, err
	}

	var zone SDNZone
	if err := json.Unmarshal([]byte(output), &zone); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}
	// debug
	fmt.Println(zone)
	return &zone, nil
}

// // UpdateSDNZone updates an existing SDN zone in Proxmox.
// func (c *SSHProxmoxClient) UpdateSDNZone(zone SDNZone) error {
// 	command := fmt.Sprintf("pvesh set /cluster/sdn/zones/%s", zone.Zone)

// 	if zone.MTU != nil && *zone.MTU != 0 {
// 		command += fmt.Sprintf(" --mtu %d", *zone.MTU)
// 	} else if zone.MTU == nil /* フィールドがあり、nilが指定されたら、そのフィールドを削除する*/ {
// 		command += " --delete mtu"
// 	}

// 	if len(zone.Nodes) > 0 && zone.Nodes != nil {
// 		nodes := strings.Join(zone.Nodes, ",")
// 		command += fmt.Sprintf(" --nodes %s", nodes)
// 	} else if zone.Nodes == nil && len(zone.Nodes) == 0 {
// 		command += " --delete nodes"
// 	}

// 	if zone.IPAM != nil && *zone.IPAM != "" {
// 		command += fmt.Sprintf(" --ipam %s", *zone.IPAM)
// 	} else if zone.IPAM == nil && *zone.IPAM == "" {
// 		command += " --delete ipam"
// 	}

// 	if zone.DNS != nil && *zone.DNS != "" {
// 		command += fmt.Sprintf(" --dns %s", *zone.DNS)
// 	} else if zone.DNS == nil && *zone.DNS == "" {
// 		command += " --delete dns"
// 	}

// 	if zone.ReverseDNS != nil && *zone.ReverseDNS != "" {
// 		command += fmt.Sprintf(" --reversedns %s", *zone.ReverseDNS)
// 	} else if zone.ReverseDNS == nil && *zone.ReverseDNS == "" {
// 		command += " --delete reversedns"
// 	}

// 	if zone.DNSZone != nil && *zone.DNSZone != "" {
// 		command += fmt.Sprintf(" --dnszone %s", *zone.DNSZone)
// 	} else if zone.DNSZone == nil && *zone.DNSZone == "" {
// 		command += " --delete dnszone"
// 	}

// 	switch zone.Type {
// 	// case "simple":
// 	// 	if zone.Simple != nil {
// 	// 		if zone.Simple.AutoDHCP != nil {
// 	// 			command += fmt.Sprintf(" --auto-dhcp %t", *zone.Simple.AutoDHCP)
// 	// 		}
// 	// 	}
// 	case "vlan":
// 		if zone.VLAN != nil {
// 			command += fmt.Sprintf(" --bridge %s", zone.VLAN.Bridge)
// 		}

// 	case "qinq":
// 		if zone.QinQ != nil {
// 			command += fmt.Sprintf(" --bridge %s", zone.QinQ.Bridge)
// 			command += fmt.Sprintf(" --tag %d", zone.QinQ.Tag)
// 			if zone.QinQ.VLANProtocol != nil && *zone.QinQ.VLANProtocol != "" {
// 				command += fmt.Sprintf(" --vlan-protocol %s", *zone.QinQ.VLANProtocol)
// 			} else if zone.QinQ.VLANProtocol == nil && *zone.QinQ.VLANProtocol == "" {
// 				command += " --delete vlan-protocol"
// 			}
// 		}
// 	case "vxlan":
// 		if zone.VXLAN != nil && len(zone.VXLAN.Peer) > 0 {
// 			peers := strings.Join(zone.VXLAN.Peer, ",")
// 			command += fmt.Sprintf(" --peers %s", peers)
// 		} else {
// 			command += " --delete peers"
// 		}
// 	case "evpn":
// 		if zone.EVPN != nil {
// 			command += fmt.Sprintf(" --controller %s", zone.EVPN.Controller)
// 			command += fmt.Sprintf(" --vrf-vxlan %d", zone.EVPN.VRFVXLAN)

// 			if zone.EVPN.MAC != nil && *zone.EVPN.MAC != "" {
// 				command += fmt.Sprintf(" --mac %s", *zone.EVPN.MAC)
// 			} else if zone.EVPN.MAC == nil && *zone.EVPN.MAC == "" {
// 				command += " --delete mac"
// 			}

// 			if len(zone.EVPN.ExitNodes) > 0 && zone.EVPN.ExitNodes != nil {
// 				exitNodes := strings.Join(zone.EVPN.ExitNodes, ",")
// 				command += fmt.Sprintf(" --exitnodes %s", exitNodes)
// 			} else if zone.EVPN.ExitNodes == nil && len(zone.EVPN.ExitNodes) == 0 {
// 				command += " --delete exitnodes"
// 			}

// 			if zone.EVPN.PrimaryExitNode != nil && *zone.EVPN.PrimaryExitNode != "" {
// 				command += fmt.Sprintf(" --primary-exitnode %s", *zone.EVPN.PrimaryExitNode)
// 			} else if zone.EVPN.PrimaryExitNode == nil && *zone.EVPN.PrimaryExitNode == "" {
// 				command += " --delete primary-exitnode"
// 			}

// 			if zone.EVPN.ExitNodesLocalRouting != nil {
// 				val := "false"
// 				if *zone.EVPN.ExitNodesLocalRouting {
// 					val = "true"
// 				}
// 				command += fmt.Sprintf(" --exitnodes-local-routing %s", val)
// 			} else if zone.EVPN.ExitNodesLocalRouting == nil {
// 				command += " --delete exitnodes-local-routing"
// 			}

// 			if zone.EVPN.AdvertiseSubnets != nil {
// 				val := "false"
// 				if *zone.EVPN.AdvertiseSubnets {
// 					val = "true"
// 				}
// 				command += fmt.Sprintf(" --advertise-subnets %s", val)
// 			} else if zone.EVPN.AdvertiseSubnets == nil {
// 				command += " --delete advertise-subnets"
// 			}

// 			if zone.EVPN.DisableARPNdSuppression != nil {
// 				val := "false"
// 				if *zone.EVPN.DisableARPNdSuppression {
// 					val = "true"
// 				}
// 				command += fmt.Sprintf(" --disable-arp-nd-suppression %s", val)
// 			} else if zone.EVPN.DisableARPNdSuppression == nil {
// 				command += " --delete disable-arp-nd-suppression"
// 			}

// 			if zone.EVPN.RouteTargetImport != nil && *zone.EVPN.RouteTargetImport != "" {
// 				command += fmt.Sprintf(" --rt-import %s", *zone.EVPN.RouteTargetImport)
// 			} else if zone.EVPN.RouteTargetImport == nil && *zone.EVPN.RouteTargetImport == "" {
// 				command += " --delete rt-import"
// 			}
// 		}
// 	}

// 	_, err := c.RunCommand(command)
// 	if err != nil {
// 		return fmt.Errorf("failed to update SDN zone: %v", err)
// 	}
// 	return nil
// }

// DeleteSDNZone deletes an existing SDN zone in Proxmox.
func (c *SSHProxmoxClient) DeleteSDNZone(zoneID string) error {
	command := fmt.Sprintf("pvesh delete /cluster/sdn/zones/%s", zoneID)
	_, err := c.RunCommand(command)
	if err != nil {
		return fmt.Errorf("failed to delete SDN zone: %v", err)
	}
	return nil
}
