package client

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SDNZoneを作成する関数
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
			command += fmt.Sprintf(" --exitnodes-primary %s", *zone.PrimaryExitNode)
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

	// 変更を反映
	if _, err = c.RunCommand("pvesh set /cluster/sdn"); err != nil {
		return fmt.Errorf("failed to apply changes to SDN: %v", err)
	}
	return nil
}

// SDNZoneを取得する関数
func (c *SSHProxmoxClient) GetSDNZones() ([]SDNZone, error) {
	// pvesh get /cluster/sdn/zones --output-format json
	cmd := "pvesh get /cluster/sdn/zones --output-format json"
	output, err := c.RunCommand(cmd)
	if err != nil {
		return nil, err
	}

	var zones []SDNZone
	if err := json.Unmarshal([]byte(output), &zones); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}
	// debug
	// fmt.Println(zones)
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
	// fmt.Println(zone)
	return &zone, nil
}

// SDNZoneを更新する関数
func (c *SSHProxmoxClient) UpdateSDNZone(zone SDNZone) error {
	command := fmt.Sprintf(
		"pvesh set /cluster/sdn/zones/%s",
		zone.Zone,
	)

	// if zone.Type != "" {
	// 	command += fmt.Sprintf(" --type %s", zone.Type)
	// }

	if zone.MTU != nil {
		if *zone.MTU != 0 {
			// 値がある場合は設定
			command += fmt.Sprintf(" --mtu %d", *zone.MTU)
		} else {
			// 値がない場合は削除
			command += " --delete mtu"
		}
	}

	if len(zone.Nodes) > 0 {
		nodes := strings.Join(zone.Nodes, ",")
		command += fmt.Sprintf(" --nodes %s", nodes)
	} else if zone.Nodes != nil {
		// 空のスライスが渡された場合は削除
		command += " --delete nodes"
	}

	if zone.IPAM != nil {
		if *zone.IPAM != "" {
			command += fmt.Sprintf(" --ipam %s", *zone.IPAM)
		} else {
			command += " --delete ipam"
		}
	}

	if zone.DNS != nil {
		if *zone.DNS != "" {
			command += fmt.Sprintf(" --dns %s", *zone.DNS)
		} else {
			command += " --delete dns"
		}
	}

	if zone.ReverseDNS != nil {
		if *zone.ReverseDNS != "" {
			command += fmt.Sprintf(" --reversedns %s", *zone.ReverseDNS)
		} else {
			command += " --delete reversedns"
		}
	}

	if zone.DNSZone != nil {
		if *zone.DNSZone != "" {
			command += fmt.Sprintf(" --dnszone %s", *zone.DNSZone)
		} else {
			command += " --delete dnszone"
		}
	}

	// zone.Type に応じたフィールドを処理
	switch zone.Type {
	case "vlan":
		if zone.Bridge != nil {
			if *zone.Bridge != "" {
				command += fmt.Sprintf(" --bridge %s", *zone.Bridge)
			} else {
				command += " --delete bridge"
			}
		}

	case "qinq":
		if zone.Bridge != nil {
			if *zone.Bridge != "" {
				command += fmt.Sprintf(" --bridge %s", *zone.Bridge)
			} else {
				command += " --delete bridge"
			}
		}

		if zone.Tag != nil {
			if *zone.Tag != 0 {
				command += fmt.Sprintf(" --tag %d", *zone.Tag)
			} else {
				command += " --delete tag"
			}
		}

		if zone.VLANProtocol != nil {
			if *zone.VLANProtocol != "" {
				command += fmt.Sprintf(" --vlan-protocol %s", *zone.VLANProtocol)
			} else {
				command += " --delete vlan-protocol"
			}
		}

	case "vxlan":
		if len(zone.Peers) > 0 {
			peers := strings.Join(zone.Peers, ",")
			command += fmt.Sprintf(" --peers %s", peers)
		} else if zone.Peers != nil {
			command += " --delete peers"
		}

	case "evpn":
		if zone.Controller != nil {
			if *zone.Controller != "" {
				command += fmt.Sprintf(" --controller %s", *zone.Controller)
			} else {
				command += " --delete controller"
			}
		}

		if zone.VRFVXLAN != nil {
			if *zone.VRFVXLAN != 0 {
				command += fmt.Sprintf(" --vrf-vxlan %d", *zone.VRFVXLAN)
			} else {
				command += " --delete vrf-vxlan"
			}
		}

		if zone.MAC != nil {
			if *zone.MAC != "" {
				command += fmt.Sprintf(" --mac %s", *zone.MAC)
			} else {
				command += " --delete mac"
			}
		}

		if len(zone.ExitNodes) > 0 {
			exitNodes := strings.Join(zone.ExitNodes, ",")
			command += fmt.Sprintf(" --exitnodes %s", exitNodes)
		} else if zone.ExitNodes != nil {
			command += " --delete exitnodes"
		}

		if zone.PrimaryExitNode != nil {
			if *zone.PrimaryExitNode != "" {
				command += fmt.Sprintf(" --primary-exitnode %s", *zone.PrimaryExitNode)
			} else {
				command += " --delete primary-exitnode"
			}
		}

		if zone.ExitNodesLocalRouting != nil {
			val := "0"
			if *zone.ExitNodesLocalRouting {
				val = "1"
			}
			command += fmt.Sprintf(" --exitnodes-local-routing %s", val)
		}

		if zone.AdvertiseSubnets != nil {
			val := "0"
			if *zone.AdvertiseSubnets {
				val = "1"
			}
			command += fmt.Sprintf(" --advertise-subnets %s", val)
		}

		if zone.DisableARPNdSuppression != nil {
			val := "0"
			if *zone.DisableARPNdSuppression {
				val = "1"
			}
			command += fmt.Sprintf(" --disable-arp-nd-suppression %s", val)
		}

		if zone.RouteTargetImport != nil {
			if *zone.RouteTargetImport != "" {
				command += fmt.Sprintf(" --rt-import %s", *zone.RouteTargetImport)
			} else {
				command += " --delete rt-import"
			}
		}
	}

	// コマンドを実行
	_, err := c.RunCommand(command)
	if err != nil {
		return fmt.Errorf("failed to update SDN zone: %v", err)
	}
	// 変更を反映
	if _, err = c.RunCommand("pvesh set /cluster/sdn"); err != nil {
		return fmt.Errorf("failed to apply changes to SDN: %v", err)
	}

	return nil
}

// SDNZoneを削除する関数
func (c *SSHProxmoxClient) DeleteSDNZone(zoneID string) error {
	command := fmt.Sprintf("pvesh delete /cluster/sdn/zones/%s", zoneID)
	_, err := c.RunCommand(command)
	if err != nil {
		return fmt.Errorf("failed to delete SDN zone: %v", err)
	}

	// 変更を反映
	if _, err = c.RunCommand("pvesh set /cluster/sdn"); err != nil {
		return fmt.Errorf("failed to apply changes to SDN: %v", err)
	}
	return nil
}
