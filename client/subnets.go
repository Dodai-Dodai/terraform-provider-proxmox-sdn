package client

import (
	"encoding/json"
	"fmt"
	"log"
)

// Subnetを作成する関数
func (c *SSHProxmoxClient) CreateSubnet(subnet SDNSubnets) error {
	command := fmt.Sprintf("pvesh create /cluster/sdn/vnets/%s/subnets --subnet %s --type subnet", subnet.Vnet, subnet.Subnet)

	if subnet.DhcpDnsServer != nil && *subnet.DhcpDnsServer != "" {
		command += fmt.Sprintf(" --dhcp-dns-server %s", *subnet.DhcpDnsServer)
	}

	if len(subnet.DhcpRange) > 0 {
		for _, dr := range subnet.DhcpRange {
			command += fmt.Sprintf(" --dhcp-range start-address=%s,end-address=%s", dr.StartAddress, dr.EndAddress)
		}
	}

	if subnet.DnsZonePrefix != nil && *subnet.DnsZonePrefix != "" {
		command += fmt.Sprintf(" --dnszoneprefix %s", *subnet.DnsZonePrefix)
	}

	if subnet.Gateway != nil && *subnet.Gateway != "" {
		command += fmt.Sprintf(" --gateway %s", *subnet.Gateway)
	}

	if subnet.Snat != nil {
		val := false
		if *subnet.Snat {
			val = true
		}
		command += fmt.Sprintf(" --snat %t", val)
	}

	log.Println(command)

	_, err := c.RunCommand(command)
	if err != nil {
		return fmt.Errorf("failed to create subnet: %w", err)
	}

	// 変更を反映
	if _, err = c.RunCommand("pvesh set /cluster/sdn"); err != nil {
		return fmt.Errorf("failed to apply changes to SDN: %w", err)
	}
	return nil
}

// Subnetsを取得する関数
func (c *SSHProxmoxClient) GetSubnets(vnet string) ([]SDNSubnets, error) {
	// pvesh get /cluster/sdn/subnets --output-format json
	cmd := fmt.Sprintf("pvesh get /cluster/sdn/vnets/%s/subnets --output-format json", vnet)
	output, err := c.RunCommand(cmd)
	if err != nil {
		return nil, err
	}

	var subnets []SDNSubnets
	if err := json.Unmarshal([]byte(output), &subnets); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return subnets, nil
}

func (c *SSHProxmoxClient) GetSubnet(vnet string, zone string, subnetID string) (*SDNSubnets, error) {
	// subnetIDのスラッシュを-に変換 subnetIDはIPアドレス/プレフィックス長の形式
	subnetID = subnetID[:len(subnetID)-3] + "-" + subnetID[len(subnetID)-2:]

	cmd := fmt.Sprintf("pvesh get /cluster/sdn/vnets/%s/subnets/%s-%s --output-format json", vnet, zone, subnetID)
	output, err := c.RunCommand(cmd)
	if err != nil {
		return nil, err
	}

	var subnet SDNSubnets
	if err := json.Unmarshal([]byte(output), &subnet); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return &subnet, nil
}

func (c *SSHProxmoxClient) UpdateSubnet(subnet SDNSubnets) error {
	command := fmt.Sprintf("hog")

	if subnet.DhcpDnsServer != nil && *subnet.DhcpDnsServer != "" {
		command += fmt.Sprintf(" --dhcp-dns-server %s", *subnet.DhcpDnsServer)
	}

	if len(subnet.DhcpRange) > 0 {
		dhcpRangeJson, err := json.Marshal(subnet.DhcpRange)
		if err != nil {
			return fmt.Errorf("failed to marshal dhcp-range: %w", err)
		}
		command += fmt.Sprintf(" --dhcp-range %s", string(dhcpRangeJson))
	}

	if subnet.DnsZonePrefix != nil && *subnet.DnsZonePrefix != "" {
		command += fmt.Sprintf(" --dnszoneprefix %s", *subnet.DnsZonePrefix)
	}

	if subnet.Gateway != nil && *subnet.Gateway != "" {
		command += fmt.Sprintf(" --gateway %s", *subnet.Gateway)
	}

	if subnet.Snat != nil {
		val := false
		if *subnet.Snat {
			val = true
		}
		command += fmt.Sprintf(" --snat %t", val)
	}

	_, err := c.RunCommand(command)
	if err != nil {
		return fmt.Errorf("failed to update subnet: %w", err)
	}

	// 変更を反映
	if _, err = c.RunCommand("pvesh set /cluster/sdn"); err != nil {
		return fmt.Errorf("failed to apply changes to SDN: %w", err)
	}
	return nil
}

func (c *SSHProxmoxClient) DeleteSubnet(vnet string, zone string, subnetID string) error {
	// subnetIDのスラッシュを-に変換 subnetIDはIPアドレス/プレフィックス長の形式
	subnetID = subnetID[:len(subnetID)-3] + "-" + subnetID[len(subnetID)-2:]
	cmd := fmt.Sprintf("pvesh delete /cluster/sdn/vnets/%s/subnets/%s-%s", vnet, zone, subnetID)
	_, err := c.RunCommand(cmd)
	if err != nil {
		return fmt.Errorf("failed to delete subnet: %w", err)
	}

	// 変更を反映
	if _, err = c.RunCommand("pvesh set /cluster/sdn"); err != nil {
		return fmt.Errorf("failed to apply changes to SDN: %w", err)
	}
	return nil
}
