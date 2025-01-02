package client

import (
	"encoding/json"
	"fmt"
	"strings"
)

// controllerを作成する関数
func (c *SSHProxmoxClient) CreateSDNController(controller SDNController) error {
	command := fmt.Sprintf(
		"pvesh create /cluster/sdn/controllers --type %s --controller %s",
		controller.Type, controller.Controller, // required fields,
	)

	switch controller.Type {
	case "evpn":
		peers := strings.Join(controller.Peers, ",")
		command += fmt.Sprintf(" --peers %s", peers)

		if controller.ASN != nil {
			command += fmt.Sprintf(" --asn %d", *controller.ASN)
		}
	case "bgp":
		command += fmt.Sprintf(" --node %s", *controller.Node)

		if controller.ASN != nil {
			command += fmt.Sprintf(" --asn %d", *controller.ASN)
		}

		peers := strings.Join(controller.Peers, ",")
		command += fmt.Sprintf(" --peers %s", peers)

		if controller.Ebgp != nil {
			val := "false"
			if *controller.Ebgp {
				val = "true"
			}
			command += fmt.Sprintf(" --ebgp %s", val)
		}

		if controller.Loopback != nil && *controller.Loopback != "" {
			command += fmt.Sprintf(" --loopback %s", *controller.Loopback)
		}

		if controller.EbgpMultihop != nil && *controller.EbgpMultihop != 0 {
			command += fmt.Sprintf(" --ebgp-multihop %d", *controller.EbgpMultihop)
		}

		if controller.BgpMultipathAsPathRelax != nil {
			val := "false"
			if *controller.BgpMultipathAsPathRelax {
				val = "true"
			}
			command += fmt.Sprintf(" --bgp-multipath-aspath-relax %s", val)
		}
	case "isis":
		command += fmt.Sprintf(" --node %s", *controller.Node)

		command += fmt.Sprintf(" --isis-domain %s", *controller.ISISDomain)

		command += fmt.Sprintf(" --isis-net %s", *controller.ISISNet)

		command += fmt.Sprintf(" --isis-ifaces %s", *controller.ISISIfaces)

		if controller.Loopback != nil && *controller.Loopback != "" {
			command += fmt.Sprintf(" --loopback %s", *controller.Loopback)
		}
	default:
		return fmt.Errorf("unsupported controller type: %s", controller.Type)
	}

	_, err := c.RunCommand(command)
	if err != nil {
		return fmt.Errorf("failed to create SDN controller: %v", err)
	}

	// 変更を反映
	if _, err = c.RunCommand("pvesh set /cluster/sdn"); err != nil {
		return fmt.Errorf("failed to apply changes to SDN: %v", err)
	}

	return nil

}

// controllerを取得する関数
func (c *SSHProxmoxClient) GetSDNController(controllerID string) (*SDNController, error) {
	command := fmt.Sprintf(
		"pvesh get /cluster/sdn/controllers/%s --output-format json", controllerID,
	)

	output, err := c.RunCommand(command)
	if err != nil {
		return nil, fmt.Errorf("failed to get SDN controller: %v", err)
	}

	var controller SDNController
	if err := json.Unmarshal([]byte(output), &controller); err != nil {
		return nil, fmt.Errorf("failed to unmarshal SDN controller: %v", err)
	}

	return &controller, nil
}

// controllerの一覧を取得する関数
func (c *SSHProxmoxClient) GetSDNControllers() ([]SDNController, error) {
	command := "pvesh get /cluster/sdn/controllers --output-format json"
	output, err := c.RunCommand(command)
	if err != nil {
		return nil, fmt.Errorf("failed to get SDN controllers: %v", err)
	}

	var controllers []SDNController
	if err := json.Unmarshal([]byte(output), &controllers); err != nil {
		return nil, fmt.Errorf("failed to unmarshal SDN controllers: %v", err)
	}
	// debug
	fmt.Println(controllers)
	return controllers, nil
}

// controllerを更新する関数
func (c *SSHProxmoxClient) UpdateSDNController(controller SDNController) error {
	command := fmt.Sprintf(
		"pvesh set /cluster/sdn/controllers/%s",
		controller.Controller, // required field
	)

	switch controller.Type {
	case "evpn":
		if controller.Peers != nil {
			peers := strings.Join(controller.Peers, ",")
			command += fmt.Sprintf(" --peers %s", peers)
		}

		if controller.ASN != nil {
			command += fmt.Sprintf(" --asn %d", *controller.ASN)
		}
	case "bgp":
		if controller.Node != nil {
			command += fmt.Sprintf(" --node %s", *controller.Node)
		}

		if controller.ASN != nil {
			command += fmt.Sprintf(" --asn %d", *controller.ASN)
		}

		if controller.Peers != nil {
			peers := strings.Join(controller.Peers, ",")
			command += fmt.Sprintf(" --peers %s", peers)
		}

		if controller.Ebgp != nil {
			val := "false"
			if *controller.Ebgp {
				val = "true"
			}
			command += fmt.Sprintf(" --ebgp %s", val)
		}

		if controller.Loopback != nil && *controller.Loopback != "" {
			command += fmt.Sprintf(" --loopback %s", *controller.Loopback)
		}

		if controller.EbgpMultihop != nil && *controller.EbgpMultihop != 0 {
			command += fmt.Sprintf(" --ebgp-multihop %d", *controller.EbgpMultihop)
		}

		if controller.BgpMultipathAsPathRelax != nil {
			val := "false"
			if *controller.BgpMultipathAsPathRelax {
				val = "true"
			}
			command += fmt.Sprintf(" --bgp-multipath-aspath-relax %s", val)
		}
	case "isis":
		if controller.Node != nil {
			command += fmt.Sprintf(" --node %s", *controller.Node)
		}

		if controller.ISISDomain != nil {
			command += fmt.Sprintf(" --isis-domain %s", *controller.ISISDomain)
		}

		if controller.ISISNet != nil {
			command += fmt.Sprintf(" --isis-net %s", *controller.ISISNet)
		}

		if controller.ISISIfaces != nil {
			command += fmt.Sprintf(" --isis-ifaces %s", *controller.ISISIfaces)
		}

		if controller.Loopback != nil && *controller.Loopback != "" {
			command += fmt.Sprintf(" --loopback %s", *controller.Loopback)
		}
	default:
		return fmt.Errorf("unsupported controller type: %s", controller.Type)
	}

	_, err := c.RunCommand(command)
	if err != nil {
		return fmt.Errorf("failed to update SDN controller: %v", err)
	}

	// 変更を反映
	if _, err = c.RunCommand("pvesh set /cluster/sdn"); err != nil {
		return fmt.Errorf("failed to apply changes to SDN: %v", err)
	}

	return nil
}

// controllerを削除する関数
func (c *SSHProxmoxClient) DeleteSDNController(controllerID string) error {
	command := fmt.Sprintf(
		"pvesh delete /cluster/sdn/controllers/%s", controllerID,
	)

	_, err := c.RunCommand(command)
	if err != nil {
		return fmt.Errorf("failed to delete SDN controller: %v", err)
	}

	// 変更を反映
	if _, err = c.RunCommand("pvesh set /cluster/sdn"); err != nil {
		return fmt.Errorf("failed to apply changes to SDN: %v", err)
	}

	return nil
}
