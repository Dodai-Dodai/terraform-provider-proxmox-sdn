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

// 	if
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
