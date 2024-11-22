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
			command += fmt.Sprintf(" --peers %s", peers)
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
			if peer, ok := (d["peers"]).(string); ok {
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

	var zone SDNZone
	if err := json.Unmarshal([]byte(output), &zone); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
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
			command += fmt.Sprintf(" --peers %s", peers)
		} else {
			command += " --delete peers"
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
