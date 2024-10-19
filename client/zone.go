package client

import (
	"encoding/json"
	"fmt"
)

// SDNZone represents the configuration for an SDN Zone in Proxmox.
type SDNZone struct {
	Zone                     string  `json:"zone"`
	Type                     string  `json:"type"`
	DHCP                     *string `json:"dhcp,omitempty"`
	DNS                      *string `json:"dns,omitempty"`
	DNSZone                  *string `json:"dnszone,omitempty"`
	Digest                   *string `json:"digest,omitempty"`
	IPAM                     *string `json:"ipam,omitempty"`
	MTU                      *int    `json:"mtu,omitempty"`
	Nodes                    *string `json:"nodes,omitempty"`
	Pending                  *bool   `json:"pending,omitempty"`
	ReverseDNS               *string `json:"reversedns,omitempty"`
	State                    *string `json:"state,omitempty"`
	AdvertiseSubnets         *bool   `json:"advertise_subnets,omitempty"`
	Bridge                   *string `json:"bridge,omitempty"`
	BridgeDisableMACLearning *bool   `json:"bridge_disable_mac_learning,omitempty"`
	Controller               *string `json:"controller,omitempty"`
	DisableARPDiscovery      *bool   `json:"disable_arp_discovery,omitempty"`
	DPID                     *int    `json:"dp_id,omitempty"`
	ExitNodes                *string `json:"exitnodes,omitempty"`
	ExitNodesLocalRouting    *bool   `json:"exitnodes_local_routing,omitempty"`
	MAC                      *string `json:"mac,omitempty"`
	Peers                    *string `json:"peers,omitempty"`
	RouteTargetImport        *string `json:"rt_import,omitempty"`
	//	Tag                      *int   `json:"tag,omitempty"`
	VLANProtocol *string `json:"vlan_protocol,omitempty"`
	VRFVXLAN     *int    `json:"vrf_vxlan,omitempty"`
	VXLANPort    *int    `json:"vxlan_port,omitempty"`
}

// CreateSDNZone creates a new SDN zone in Proxmox.
func (c *SSHProxmoxClient) CreateSDNZone(zone SDNZone) error {
	command := fmt.Sprintf(
		"pvesh create /cluster/sdn/zones  --type %s --zone %s",
		zone.Type, zone.Zone, // required fields,
	)

	if zone.AdvertiseSubnets != nil && *zone.AdvertiseSubnets {
		command += fmt.Sprintf(" --advertise-subnets %t", *zone.AdvertiseSubnets)
	}
	if zone.Bridge != nil && *zone.Bridge != "" {
		command += fmt.Sprintf(" --bridge %s", *zone.Bridge)
	}
	if zone.BridgeDisableMACLearning != nil && *zone.BridgeDisableMACLearning {
		command += fmt.Sprintf(" --bridge-disable-mac-learning %t", *zone.BridgeDisableMACLearning)
	}
	if zone.Controller != nil && *zone.Controller != "" {
		command += fmt.Sprintf(" --controller %s", *zone.Controller)
	}
	if zone.DHCP != nil && *zone.DHCP != "" {
		command += fmt.Sprintf(" --dhcp %s", *zone.DHCP)
	}
	if zone.DisableARPDiscovery != nil && *zone.DisableARPDiscovery {
		command += fmt.Sprintf(" --disable-arp-discovery %t", *zone.DisableARPDiscovery)
	}
	if zone.DNS != nil && *zone.DNS != "" {
		command += fmt.Sprintf(" --dns %s", *zone.DNS)
	}
	if zone.DNSZone != nil && *zone.DNSZone != "" {
		command += fmt.Sprintf(" --dnszone %s", *zone.DNSZone)
	}
	if zone.DPID != nil && *zone.DPID != 0 {
		command += fmt.Sprintf(" --dp-id %d", zone.DPID)
	}
	if zone.ExitNodes != nil && *zone.ExitNodes != "" {
		command += fmt.Sprintf(" --exitnodes %s", *zone.ExitNodes)
	}
	if zone.ExitNodesLocalRouting != nil && *zone.ExitNodesLocalRouting {
		command += fmt.Sprintf(" --exitnodes-local-routing %t", *zone.ExitNodesLocalRouting)
	}
	if zone.IPAM != nil && *zone.IPAM != "" {
		command += fmt.Sprintf(" --ipam %s", *zone.IPAM)
	}
	if zone.MAC != nil && *zone.MAC != "" {
		command += fmt.Sprintf(" --mac %s", *zone.MAC)
	}
	if zone.MTU != nil && *zone.MTU != 0 {
		command += fmt.Sprintf(" --mtu %d", *zone.MTU)
	}
	if zone.Nodes != nil && *zone.Nodes != "" {
		command += fmt.Sprintf(" --nodes %s", *zone.Nodes)
	}
	if zone.Peers != nil && *zone.Peers != "" {
		command += fmt.Sprintf(" --peers %s", *zone.Peers)
	}
	if zone.ReverseDNS != nil && *zone.ReverseDNS != "" {
		command += fmt.Sprintf(" --reversedns %s", *zone.ReverseDNS)
	}
	if zone.RouteTargetImport != nil && *zone.RouteTargetImport != "" {
		command += fmt.Sprintf(" --rt-import %s", *zone.RouteTargetImport)
	}
	// if zone.Tag != nil {
	// 	command += fmt.Sprintf(" --tag %d", *zone.Tag)
	// }
	if zone.VLANProtocol != nil && *zone.VLANProtocol != "" {
		command += fmt.Sprintf(" --vlan-protocol %s", *zone.VLANProtocol)
	}
	if zone.VRFVXLAN != nil && *zone.VRFVXLAN != 0 {
		command += fmt.Sprintf(" --vrf-vxlan %d", zone.VRFVXLAN)
	}
	if zone.VXLANPort != nil && *zone.VXLANPort != 0 {
		command += fmt.Sprintf(" --vxlan-port %d", zone.VXLANPort)
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

	var zones []SDNZone
	if err := json.Unmarshal([]byte(output), &zones); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %v", err)
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
	command := fmt.Sprintf(
		"pvesh set /cluster/sdn/zones/%s",
		zone.Zone,
	)
	if zone.AdvertiseSubnets != nil && *zone.AdvertiseSubnets {
		command += fmt.Sprintf(" --advertise-subnets %t", *zone.AdvertiseSubnets)
	}
	if zone.Bridge != nil && *zone.Bridge != "" {
		command += fmt.Sprintf(" --bridge %s", *zone.Bridge)
	}
	if zone.BridgeDisableMACLearning != nil && *zone.BridgeDisableMACLearning {
		command += fmt.Sprintf(" --bridge-disable-mac-learning %t", *zone.BridgeDisableMACLearning)
	}
	if zone.Controller != nil && *zone.Controller != "" {
		command += fmt.Sprintf(" --controller %s", *zone.Controller)
	}
	if zone.DHCP != nil && *zone.DHCP != "" {
		command += fmt.Sprintf(" --dhcp %s", *zone.DHCP)
	}
	if zone.DisableARPDiscovery != nil && *zone.DisableARPDiscovery {
		command += fmt.Sprintf(" --disable-arp-discovery %t", *zone.DisableARPDiscovery)
	}
	if zone.DNS != nil && *zone.DNS != "" {
		command += fmt.Sprintf(" --dns %s", *zone.DNS)
	}
	if zone.DNSZone != nil && *zone.DNSZone != "" {
		command += fmt.Sprintf(" --dnszone %s", *zone.DNSZone)
	}
	if zone.DPID != nil && *zone.DPID != 0 {
		command += fmt.Sprintf(" --dp-id %d", zone.DPID)
	}
	if zone.ExitNodes != nil && *zone.ExitNodes != "" {
		command += fmt.Sprintf(" --exitnodes %s", *zone.ExitNodes)
	}
	if zone.ExitNodesLocalRouting != nil && *zone.ExitNodesLocalRouting {
		command += fmt.Sprintf(" --exitnodes-local-routing %t", *zone.ExitNodesLocalRouting)
	}
	if zone.IPAM != nil && *zone.IPAM != "" {
		command += fmt.Sprintf(" --ipam %s", *zone.IPAM)
	}
	if zone.MAC != nil && *zone.MAC != "" {
		command += fmt.Sprintf(" --mac %s", *zone.MAC)
	}
	if zone.MTU != nil && *zone.MTU != 0 {
		command += fmt.Sprintf(" --mtu %d", zone.MTU)
	}
	if zone.Nodes != nil && *zone.Nodes != "" {
		command += fmt.Sprintf(" --nodes %s", *zone.Nodes)
	}
	if zone.Peers != nil && *zone.Peers != "" {
		command += fmt.Sprintf(" --peers %s", *zone.Peers)
	}
	if zone.ReverseDNS != nil && *zone.ReverseDNS != "" {
		command += fmt.Sprintf(" --reversedns %s", *zone.ReverseDNS)
	}
	if zone.RouteTargetImport != nil && *zone.RouteTargetImport != "" {
		command += fmt.Sprintf(" --rt-import %s", *zone.RouteTargetImport)
	}
	// if zone.Tag != nil {
	// 	command += fmt.Sprintf(" --tag %d", *zone.Tag)
	// }
	if zone.VLANProtocol != nil && *zone.VLANProtocol != "" {
		command += fmt.Sprintf(" --vlan-protocol %s", *zone.VLANProtocol)
	}
	if zone.VRFVXLAN != nil && *zone.VRFVXLAN != 0 {
		command += fmt.Sprintf(" --vrf-vxlan %d", zone.VRFVXLAN)
	}
	if zone.VXLANPort != nil && *zone.VXLANPort != 0 {
		command += fmt.Sprintf(" --vxlan-port %d", zone.VXLANPort)
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
