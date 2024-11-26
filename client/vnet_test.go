package client

import (
	"os"
	"testing"
)

func TestCreateVnet(t *testing.T) {
	vnetName := "testvnet"
	zoneName := "vxlan"
	tag := int64(100)
	proxmoxHost := os.Getenv("PROXMOX_HOST")
	proxmoxUser := os.Getenv("PROXMOX_USERNAME")
	proxmoxPassword := os.Getenv("PROXMOX_PASSWORD")

	if proxmoxHost == "" || proxmoxUser == "" || proxmoxPassword == "" {
		t.Fatal("必要な環境変数(PROXMOX_HOST, PROXMOX_USER, PROXMOX_PASSWORD)が設定されていません")
	}

	client, err := NewSSHProxmoxClient(proxmoxUser, proxmoxPassword, proxmoxHost)
	if err != nil {
		t.Fatalf("Proxmoxクライアントの作成に失敗しました: %v", err)
	}

	vnet := SDNVnet{
		Vnet: vnetName,
		Zone: zoneName,
		Tag:  &tag,
	}

	err = client.CreateVnet(vnet)
	if err != nil {
		t.Fatalf("CreateVnet failed: %v", err)
	}

}

func TestGetVnet(t *testing.T) {
	vnetName := "testvnet"
	proxmoxHost := os.Getenv("PROXMOX_HOST")
	proxmoxUser := os.Getenv("PROXMOX_USERNAME")
	proxmoxPassword := os.Getenv("PROXMOX_PASSWORD")

	if proxmoxHost == "" || proxmoxUser == "" || proxmoxPassword == "" {
		t.Fatal("必要な環境変数(PROXMOX_HOST, PROXMOX_USER, PROXMOX_PASSWORD)が設定されていません")
	}

	sshClient, err := NewSSHProxmoxClient(proxmoxUser, proxmoxPassword, proxmoxHost)
	if err != nil {
		t.Fatalf("Proxmoxクライアントの作成に失敗しました: %v", err)
	}

	vnet, err := sshClient.GetVnet(vnetName)
	if err != nil {
		t.Fatalf("GetVnet failed: %v", err)
	}

	if vnet.Vnet != vnetName {
		t.Errorf("Expected vnet name %s, got %s", vnetName, vnet.Vnet)
	}

	t.Logf("vnet: %+v", vnet)
	t.Logf("vnet.Vnet: %s", vnet.Vnet)
	t.Logf("vnet.Zone: %s", vnet.Zone)
	t.Logf("vnet.Tag: %d", *vnet.Tag)
	//t.Logf("vnet.Vlanaware: %t", *vnet.Vlanaware)

}

func TestDeleteVnet(t *testing.T) {
	vnetName := "testvnet"

	proxmoxHost := os.Getenv("PROXMOX_HOST")
	proxmoxUser := os.Getenv("PROXMOX_USERNAME")
	proxmoxPassword := os.Getenv("PROXMOX_PASSWORD")

	if proxmoxHost == "" || proxmoxUser == "" || proxmoxPassword == "" {
		t.Fatal("必要な環境変数(PROXMOX_HOST, PROXMOX_USER, PROXMOX_PASSWORD)が設定されていません")
	}

	sshClient, err := NewSSHProxmoxClient(proxmoxUser, proxmoxPassword, proxmoxHost)
	if err != nil {
		t.Fatalf("Proxmoxクライアントの作成に失敗しました: %v", err)
	}

	err = sshClient.DeleteVnet(vnetName)
	if err != nil {
		t.Fatalf("DeleteVnet failed: %v", err)
	}

	_, err = sshClient.GetVnet(vnetName)
	if err == nil {
		t.Errorf("Expected error when getting deleted vnet, but got none")
	}
}
