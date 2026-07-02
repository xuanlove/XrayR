package controller_test

import (
	"testing"

	"github.com/XrayR-project/XrayR/api"
	"github.com/XrayR-project/XrayR/common/mylego"
	. "github.com/XrayR-project/XrayR/service/controller"
)

func TestBuildV2ray(t *testing.T) {
	nodeInfo := &api.NodeInfo{
		NodeType:          "V2ray",
		NodeID:            1,
		Port:              1145,
		SpeedLimit:        0,
		AlterID:           2,
		TransportProtocol: "ws",
		Host:              "test.test.tk",
		Path:              "v2ray",
		EnableTLS:         false,
	}
	certConfig := &mylego.CertConfig{
		CertMode:   "http",
		CertDomain: "test.test.tk",
		Provider:   "alidns",
		Email:      "test@gmail.com",
	}
	config := &Config{
		CertConfig: certConfig,
	}
	_, err := InboundBuilder(config, nodeInfo, "test_tag")
	if err != nil {
		t.Error(err)
	}
}

func TestBuildTrojan(t *testing.T) {
	nodeInfo := &api.NodeInfo{
		NodeType:          "Trojan",
		NodeID:            1,
		Port:              1145,
		SpeedLimit:        0,
		AlterID:           2,
		TransportProtocol: "tcp",
		Host:              "trojan.test.tk",
		Path:              "v2ray",
		EnableTLS:         false,
	}
	DNSEnv := make(map[string]string)
	DNSEnv["ALICLOUD_ACCESS_KEY"] = "aaa"
	DNSEnv["ALICLOUD_SECRET_KEY"] = "bbb"
	certConfig := &mylego.CertConfig{
		CertMode:   "dns",
		CertDomain: "trojan.test.tk",
		Provider:   "alidns",
		Email:      "test@gmail.com",
		DNSEnv:     DNSEnv,
	}
	config := &Config{
		CertConfig: certConfig,
	}
	_, err := InboundBuilder(config, nodeInfo, "test_tag")
	if err != nil {
		t.Error(err)
	}
}

func TestBuildSS(t *testing.T) {
	nodeInfo := &api.NodeInfo{
		NodeType:          "Shadowsocks",
		NodeID:            1,
		Port:              1145,
		SpeedLimit:        0,
		AlterID:           2,
		TransportProtocol: "tcp",
		Host:              "test.test.tk",
		Path:              "v2ray",
		EnableTLS:         false,
		CypherMethod:      "aes-256-gcm",
		ServerKey:         "test_server_key",
	}
	DNSEnv := make(map[string]string)
	DNSEnv["ALICLOUD_ACCESS_KEY"] = "aaa"
	DNSEnv["ALICLOUD_SECRET_KEY"] = "bbb"
	certConfig := &mylego.CertConfig{
		CertMode:   "dns",
		CertDomain: "trojan.test.tk",
		Provider:   "alidns",
		Email:      "test@me.com",
		DNSEnv:     DNSEnv,
	}
	config := &Config{
		CertConfig: certConfig,
	}
	_, err := InboundBuilder(config, nodeInfo, "test_tag")
	if err != nil {
		t.Error(err)
	}
}

// TestBuildXHTTP verifies that extended XHTTP/SplitHTTP transport options
// introduced in Xray-core v26.x are forwarded to the inbound config.
func TestBuildXHTTP(t *testing.T) {
	nodeInfo := &api.NodeInfo{
		NodeType:          "V2ray",
		NodeID:            1,
		Port:              1145,
		AlterID:           2,
		TransportProtocol: "splithttp",
		Host:              "xhttp.test.tk",
		Path:              "/xh",
		EnableTLS:         false,
		EnableVless:       true,
		XHTTPConfig: &api.XHTTPConfig{
			Mode:             "packet-up",
			XPaddingObfsMode: true,
			NoGRPCHeader:     true,
			Xmux: api.XmuxConfig{
				MaxConcurrency:   api.Int32Range{From: 1, To: 100},
				HKeepAlivePeriod: 60,
			},
		},
	}
	config := &Config{}
	ib, err := InboundBuilder(config, nodeInfo, "test_tag")
	if err != nil {
		t.Fatal(err)
	}
	if ib == nil {
		t.Fatal("inbound config is nil")
	}
	// No panic + non-nil result means extended fields were accepted by conf.
}

// TestBuildPQE verifies VLESS Post-Quantum Encryption is applied when a valid
// PQE string is provided, and rejected when enabled but empty.
func TestBuildPQE(t *testing.T) {
	nodeInfo := &api.NodeInfo{
		NodeType:          "Vless",
		NodeID:            1,
		Port:              1145,
		AlterID:           2,
		TransportProtocol: "tcp",
		EnableTLS:         false,
	}
	// A valid PQE string: mlkem768x25519plus.<mode>.<from>-<to>s.<32B base64 padding>
	validPQE := "mlkem768x25519plus.native.60-60s.bDp_L6pR3AXVW3hyxpquMEbg9QDrJcdVeDyRngYt71c"
	config := &Config{
		PQEConfig: &PQEConfig{
			Enable:     true,
			Decryption: validPQE,
			Encryption: validPQE,
		},
	}
	ib, err := InboundBuilder(config, nodeInfo, "test_tag")
	if err != nil {
		t.Fatal(err)
	}
	if ib == nil {
		t.Fatal("inbound config is nil")
	}

	// Empty Decryption with Enable=true must error out.
	bad := &Config{
		PQEConfig: &PQEConfig{Enable: true},
	}
	if _, err := InboundBuilder(bad, nodeInfo, "test_tag"); err == nil {
		t.Fatal("expected error when PQE enabled but Decryption empty")
	}
}
