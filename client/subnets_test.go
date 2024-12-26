package client

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateSubnet(t *testing.T) {
	subnetID := "192.168.50.0/24"
	vnetName := "vnetnam"
	proxmoxHost := os.Getenv("PROXMOX_HOST")
	proxmoxUser := os.Getenv("PROXMOX_USERNAME")
	proxmoxPassword := os.Getenv("PROXMOX_PASSWORD")

	if proxmoxHost == "" || proxmoxUser == "" || proxmoxPassword == "" {
		t.Fatal("必要な環境変数(PROXMOX_HOST, PROXMOX_USERNAME, PROXMOX_PASSWORD)が設定されていません")
	}

	client, err := NewSSHProxmoxClient(proxmoxUser, proxmoxPassword, proxmoxHost)
	if err != nil {
		t.Fatalf("Proxmoxクライアントの作成に失敗しました: %v", err)
	}

	subnet := SDNSubnets{
		Subnet:  subnetID,
		Type:    "subnet",
		Vnet:    vnetName,
		Gateway: StringPointer("192.168.50.254"),
		// 必要に応じて他のフィールドを設定します
		DhcpRange: []DhcpRange{
			{
				StartAddress: "192.168.50.100",
				EndAddress:   "192.168.50.110",
			},
		},
	}

	err = client.CreateSubnet(subnet)
	if err != nil {
		t.Fatalf("CreateSubnetに失敗しました: %v", err)
	}
}

func TestGetSubnets(t *testing.T) {
	subnetID := "192.168.50.0/24"
	vnetName := "vnetnam"
	zone := "zonenam"
	proxmoxHost := os.Getenv("PROXMOX_HOST")
	proxmoxUser := os.Getenv("PROXMOX_USERNAME")
	proxmoxPassword := os.Getenv("PROXMOX_PASSWORD")

	if proxmoxHost == "" || proxmoxUser == "" || proxmoxPassword == "" {
		t.Fatal("必要な環境変数(PROXMOX_HOST, PROXMOX_USERNAME, PROXMOX_PASSWORD)が設定されていません")
	}

	client, err := NewSSHProxmoxClient(proxmoxUser, proxmoxPassword, proxmoxHost)
	if err != nil {
		t.Fatalf("Proxmoxクライアントの作成に失敗しました: %v", err)
	}

	subnet, err := client.GetSubnet(vnetName, zone, subnetID)
	if err != nil {
		t.Fatalf("GetSubnetsに失敗しました: %v", err)
	}

	// if subnet.Subnet != subnetID {
	// 	t.Errorf("期待されたサブネットID %s が取得されましたが、実際には %s でした", subnetID, subnet.Subnet)
	// }

	// 成功か失敗かにかかわらず、取得されたサブネット情報を表示します
	fmt.Println(subnet)
}

func TestDeleteSubnet(t *testing.T) {
	subnetID := "192.168.50.0/24"
	vnetName := "vnetnam"
	zone := "zonenam"
	proxmoxHost := os.Getenv("PROXMOX_HOST")
	proxmoxUser := os.Getenv("PROXMOX_USERNAME")
	proxmoxPassword := os.Getenv("PROXMOX_PASSWORD")

	if proxmoxHost == "" || proxmoxUser == "" || proxmoxPassword == "" {
		t.Fatal("必要な環境変数(PROXMOX_HOST, PROXMOX_USERNAME, PROXMOX_PASSWORD)が設定されていません")
	}

	client, err := NewSSHProxmoxClient(proxmoxUser, proxmoxPassword, proxmoxHost)
	if err != nil {
		t.Fatalf("Proxmoxクライアントの作成に失敗しました: %v", err)
	}

	err = client.DeleteSubnet(vnetName, zone, subnetID)
	if err != nil {
		t.Fatalf("DeleteSubnetに失敗しました: %v", err)
	}

	_, err = client.GetSubnet(vnetName, zone, subnetID)
	if err == nil {
		t.Errorf("削除されたサブネットの取得でエラーが発生するはずですが、実際には発生しませんでした")
	}
}

// ヘルパー関数
func StringPointer(s string) *string {
	return &s
}

func BoolPointer(b bool) *bool {
	return &b
}
