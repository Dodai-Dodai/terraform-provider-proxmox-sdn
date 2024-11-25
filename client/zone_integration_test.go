// zone_integration_test.go
package client

import (
	"fmt"
	"os"
	"testing"
)

func TestCreateSDNZone_Integration(t *testing.T) {
	// 環境変数から接続情報を取得
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
	zone := SDNZone{
		Type:   "vlan",
		Zone:   "testzone",
		Bridge: ptrString("br0"),
		MTU:    ptrInt64(1500),
		Nodes:  []string{"node1", "node2"},
	}

	// テスト対象の関数を実行
	err = client.CreateSDNZone(zone)
	if err != nil {
		t.Fatalf("SDNゾーンの作成に失敗しました: %v", err)
	}

	// // 作成されたゾーンが存在するか確認
	// createdZone, err := client.GetSDNZone(zone.Zone)
	// if err != nil {
	// 	t.Fatalf("作成されたSDNゾーンの取得に失敗しました: %v", err)
	// }

	// if createdZone.Type != zone.Type || createdZone.Zone != zone.Zone {
	// 	t.Errorf("作成されたSDNゾーンの内容が一致しません")
	// }
}

func TestDeleteSDNZone_Integration(t *testing.T) {
	// 環境変数から接続情報を取得
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

	zone := SDNZone{
		Zone: "testzone",
	}

	// テスト対象の関数を実行
	err = client.DeleteSDNZone(zone.Zone)
	if err != nil {
		t.Fatalf("SDNゾーンの削除に失敗しました: %v", err)
	}

}

func TestGetSDNZones_Integration(t *testing.T) {
	// 環境変数から接続情報を取得
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

	// テスト対象の関数を実行
	zones, err := client.GetSDNZones()
	if err != nil {
		t.Fatalf("SDNゾーンの取得に失敗しました: %v", err)
	}

	if len(zones) == 0 {
		t.Fatal("取得されたSDNゾーンが0件です")
	}

	// 取得されたゾーンの内容を出力 (nil ポインタを安全に処理)
	for _, zone := range zones {
		bridge := "<nil>"
		if zone.Bridge != nil {
			bridge = *zone.Bridge
		}
		mtu := "<nil>"
		if zone.MTU != nil {
			mtu = fmt.Sprintf("%d", *zone.MTU)
		}
		nodes := "<nil>"
		if zone.Nodes != nil {
			nodes = fmt.Sprintf("%v", zone.Nodes)
		}

		fmt.Printf("Zone: %s, Bridge: %s, MTU: %s, nodes%s\n", zone.Zone, bridge, mtu, nodes)
	}
}

func TestGetSDNZone_Integration(t *testing.T) {
	// 環境変数から接続情報を取得
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

	zoneName := "testzone"

	// テスト対象の関数を実行
	zone, err := client.GetSDNZone(zoneName)
	if err != nil {
		t.Fatalf("SDNゾーンの取得に失敗しました: %v", err)
	}

	if zone.Zone != zoneName {
		t.Fatalf("取得されたゾーン名が一致しません: %s != %s", zone.Zone, zoneName)
	}

	bridge := "<nil>"
	if zone.Bridge != nil {
		bridge = *zone.Bridge
	}
	mtu := "<nil>"
	if zone.MTU != nil {
		mtu = fmt.Sprintf("%d", *zone.MTU)
	}
	nodes := "<nil>"
	if zone.Nodes != nil {
		nodes = fmt.Sprintf("%v", zone.Nodes)
	}

	fmt.Printf("Zone: %s, Bridge: %s, MTU: %s, nodes%s\n", zone.Zone, bridge, mtu, nodes)
}

func ptrString(s string) *string {
	return &s
}

func ptrInt64(i int64) *int64 {
	return &i
}
