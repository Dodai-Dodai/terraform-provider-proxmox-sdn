package client

import (
	"encoding/json"
	"fmt"
)

// Vnetを作成する関数
func (c *SSHProxmoxClient) CreateVnet(vnet SDNVnet) error {
	command := fmt.Sprintf(
		"pvesh create /cluster/sdn/vnets --vnet %s --zone %s",
		vnet.Vnet, vnet.Zone, // required fields
	)

	if vnet.Type != "" {
		// always set 'vnet'
	}

	if vnet.Tag != nil && *vnet.Tag != 0 {
		command += fmt.Sprintf(" --tag %d", *vnet.Tag)
	}

	if vnet.Alias != nil && *vnet.Alias != "" {
		command += fmt.Sprintf(" --alias %s", *vnet.Alias)
	}

	if vnet.Vlanaware != nil {
		val := false
		if *vnet.Vlanaware {
			val = true
		}
		command += fmt.Sprintf(" --vlanaware %t", val)
	}

	_, err := c.RunCommand(command)
	if err != nil {
		return fmt.Errorf("failed to create vnet: %w", err)
	}

	// 変更を反映
	if _, err = c.RunCommand("pvesh set /cluster/sdn"); err != nil {
		return fmt.Errorf("failed to apply changes to SDN: %w", err)
	}
	return nil
}

// Vnetを取得する関数
func (c *SSHProxmoxClient) GetVnets() ([]SDNVnet, error) {
	// pvesh get /cluster/sdn/vnets --output-format json
	cmd := "pvesh get /cluster/sdn/vnets --output-format json"
	output, err := c.RunCommand(cmd)
	if err != nil {
		return nil, err
	}

	var vnets []SDNVnet
	if err := json.Unmarshal([]byte(output), &vnets); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}
	// debug
	// fmt.Println(vnets)
	return vnets, nil
}

func (c *SSHProxmoxClient) GetVnet(vnetID string) (*SDNVnet, error) {
	cmd := fmt.Sprintf("pvesh get /cluster/sdn/vnets/%s --output-format json", vnetID)
	output, err := c.RunCommand(cmd)
	if err != nil {
		return nil, err
	}

	var vnet SDNVnet
	if err := json.Unmarshal([]byte(output), &vnet); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}
	// debug
	// fmt.Println(vnet)
	return &vnet, nil
}

// Vnetを更新する関数
func (c *SSHProxmoxClient) UpdateVnet(vnet SDNVnet) error {
	command := fmt.Sprintf(
		"pvesh set /cluster/sdn/vnets/%s", vnet.Vnet, // required field
	)

	if vnet.Zone != "" {
		command += fmt.Sprintf(" --zone %s", vnet.Zone) // required field
	}

	if vnet.Type != "" {
		command += fmt.Sprintf(" --type %s", vnet.Type)
	}

	if vnet.Alias != nil && *vnet.Alias != "" {
		command += fmt.Sprintf(" --alias %s", *vnet.Alias)
	}

	if vnet.Tag != nil && *vnet.Tag != 0 {
		command += fmt.Sprintf(" --tag %d", *vnet.Tag)
	}

	if vnet.Vlanaware != nil {
		val := false
		if *vnet.Vlanaware {
			val = true
		}
		command += fmt.Sprintf(" --vlanaware %t", val)
	}

	_, err := c.RunCommand(command)
	if err != nil {
		return fmt.Errorf("failed to update vnet: %w", err)
	}

	// 変更を反映
	if _, err = c.RunCommand("pvesh set /cluster/sdn"); err != nil {
		return fmt.Errorf("failed to apply changes to SDN: %w", err)
	}
	return nil
}

// Vnetを削除する関数
func (c *SSHProxmoxClient) DeleteVnet(vnetID string) error {
	command := fmt.Sprintf(
		"pvesh delete /cluster/sdn/vnets/%s", vnetID,
	)

	_, err := c.RunCommand(command)
	if err != nil {
		return fmt.Errorf("failed to delete vnet: %w", err)
	}

	// 変更を反映
	if _, err = c.RunCommand("pvesh set /cluster/sdn"); err != nil {
		return fmt.Errorf("failed to apply changes to SDN: %w", err)
	}
	return nil
}
