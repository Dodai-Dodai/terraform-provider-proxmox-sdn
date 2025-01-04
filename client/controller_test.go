package client

import (
	"os"
	"testing"
)

func TestCreateEVPNController(t *testing.T) {
	controllerName := "evpncontroller"
	asn := int64(65001)
	peers := []string{}
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

	controller := SDNController{
		Controller: controllerName,
		Type:       "evpn",
		ASN:        &asn,
		Peers:      peers,
	}

	err = client.CreateSDNController(controller)
	if err != nil {
		t.Fatalf("CreateSDNController failed: %v", err)
	}
}

func TestGetEVPNController(t *testing.T) {
	controllerName := "evpncontroller"
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

	controller, err := client.GetSDNController(controllerName)
	if err != nil {
		t.Fatalf("GetSDNController failed: %v", err)
	}

	t.Logf("Controller Name: %s", controller.Controller)
	t.Logf("Peers: %v", controller.Peers)
	t.Logf("ASN: %d", controller.ASN)

	if controller.Controller != controllerName {
		t.Errorf("Expected controller name %s, got %s", controllerName, controller.Controller)
	}
}

func TestDeleteEVPNController(t *testing.T) {
	controllerName := "evpncontroller"
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

	err = client.DeleteSDNController(controllerName)
	if err != nil {
		t.Fatalf("DeleteSDNController failed: %v", err)
	}

	_, err = client.GetSDNController(controllerName)
	if err == nil {
		t.Errorf("削除されたコントローラの取得でエラーが発生するはずですが、実際には発生しませんでした")
	}
}

func TestCreateBGPController(t *testing.T) {
	controllerName := "bgpcontroller"
	asn := int64(65001)
	peers := []string{}
	node := "pvewata01"
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

	controller := SDNController{
		Controller: controllerName,
		Type:       "bgp",
		ASN:        &asn,
		Peers:      peers,
		Node:       &node,
	}

	err = client.CreateSDNController(controller)
	if err != nil {
		t.Fatalf("CreateSDNController failed: %v", err)
	}

}

func TestGetBGPController(t *testing.T) {
	controllerName := "bgpcontroller"
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

	controller, err := client.GetSDNController(controllerName)
	if err != nil {
		t.Fatalf("GetSDNController failed: %v", err)
	}

	t.Logf("controller: %v", controller)
	t.Logf("peers: %v", controller.Peers)
	t.Logf("asn: %v", controller.ASN)
	t.Logf("node: %v", controller.Node)

	if controller.Controller != controllerName {
		t.Errorf("Expected controller name %s, got %s", controllerName, controller.Controller)
	}
}

func TestDeleteBGPController(t *testing.T) {
	controllerName := "bgpcontroller"
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

	err = client.DeleteSDNController(controllerName)
	if err != nil {
		t.Fatalf("DeleteSDNController failed: %v", err)
	}

	_, err = client.GetSDNController(controllerName)
	if err == nil {
		t.Errorf("削除されたコントローラの取得でエラーが発生するはずですが、実際には発生しませんでした")
	}
}

func TestCreateISISController(t *testing.T) {
	controllerName := "isiscontroller"
	node := "pvewata01"
	isisDomain := "49.0001"
	isisIfaces := "eno1"
	isisNet := "49.0001.0001.0001.0001.00"
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

	controller := SDNController{
		Controller: controllerName,
		Type:       "isis",
		Node:       &node,
		ISISDomain: &isisDomain,
		ISISIfaces: &isisIfaces,
		ISISNet:    &isisNet,
	}

	err = client.CreateSDNController(controller)
	if err != nil {
		t.Fatalf("CreateSDNController failed: %v", err)
	}

}

func TestGetISISController(t *testing.T) {
	controllerName := "isiscontroller"
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

	controller, err := client.GetSDNController(controllerName)
	if err != nil {
		t.Fatalf("GetSDNController failed: %v", err)
	}

	t.Logf("controller: %v", controller)
	t.Logf("isisDomain: %v", controller.ISISDomain)
	t.Logf("isisIfaces: %v", controller.ISISIfaces)
	t.Logf("isisNet: %v", controller.ISISNet)

	if controller.Controller != controllerName {
		t.Errorf("Expected controller name %s, got %s", controllerName, controller.Controller)
	}
}

func TestDeleteISISController(t *testing.T) {
	controllerName := "isiscontroller"
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

	err = client.DeleteSDNController(controllerName)
	if err != nil {
		t.Fatalf("DeleteSDNController failed: %v", err)
	}

	_, err = client.GetSDNController(controllerName)
	if err == nil {
		t.Errorf("削除されたコントローラの取得でエラーが発生するはずですが、実際には発生しませんでした")
	}
}
