package controller

import (
	"encoding/json"

	"github.com/XrayR-project/XrayR/common/limiter"
	"github.com/XrayR-project/XrayR/common/mylego"
)

type Config struct {
	ListenIP                  string                           `mapstructure:"ListenIP"`
	SendIP                    string                           `mapstructure:"SendIP"`
	UpdatePeriodic            int                              `mapstructure:"UpdatePeriodic"`
	CertConfig                *mylego.CertConfig               `mapstructure:"CertConfig"`
	EnableDNS                 bool                             `mapstructure:"EnableDNS"`
	DNSType                   string                           `mapstructure:"DNSType"`
	DisableUploadTraffic      bool                             `mapstructure:"DisableUploadTraffic"`
	DisableGetRule            bool                             `mapstructure:"DisableGetRule"`
	EnableProxyProtocol       bool                             `mapstructure:"EnableProxyProtocol"`
	EnableFallback            bool                             `mapstructure:"EnableFallback"`
	DisableIVCheck            bool                             `mapstructure:"DisableIVCheck"`
	DisableSniffing           bool                             `mapstructure:"DisableSniffing"`
	AutoSpeedLimitConfig      *AutoSpeedLimitConfig            `mapstructure:"AutoSpeedLimitConfig"`
	GlobalDeviceLimitConfig   *limiter.GlobalDeviceLimitConfig `mapstructure:"GlobalDeviceLimitConfig"`
	FallBackConfigs           []*FallBackConfig                `mapstructure:"FallBackConfigs"`
	DisableLocalREALITYConfig bool                             `mapstructure:"DisableLocalREALITYConfig"`
	EnableREALITY             bool                             `mapstructure:"EnableREALITY"`
	REALITYConfigs            *REALITYConfig                   `mapstructure:"REALITYConfigs"`
	// VLESS Post-Quantum Encryption (Xray-core v26.6.22+, mlkem768x25519plus)
	// When non-empty, overrides the VLESS inbound Decryption string and each user's
	// Encryption string with the PQE value, enabling post-quantum encryption.
	PQEConfig *PQEConfig `mapstructure:"PQEConfig"`
	// TUN inbound configuration (Xray-core v26.1.23+). When non-nil, an additional
	// TUN inbound is registered for transparent proxying on the node.
	TUNConfig *TUNConfig `mapstructure:"TUNConfig"`
	// ProcessRouteConfig enables process-name based routing/audit rules
	// (Xray-core v26.1.23+). Matched process names are rejected via the rule manager.
	ProcessRouteConfig *ProcessRouteConfig `mapstructure:"ProcessRouteConfig"`
	// MetricsConfig enables the Prometheus metrics endpoint exposed by xray-core's
	// app/metrics feature. When non-nil, a metrics server is started.
	MetricsConfig *MetricsConfig `mapstructure:"MetricsConfig"`
}

// PQEConfig configures VLESS Post-Quantum Encryption.
// The Decryption string format is: mlkem768x25519plus.<mode>.<rtt>.<seconds>
// where mode in {native, xorpub, random}, rtt in {1rtt, 0rtt}, seconds is an int.
type PQEConfig struct {
	// Enable turns on PQE for VLESS inbound and users.
	Enable bool `mapstructure:"Enable"`
	// Decryption is the PQE decryption string for the inbound, e.g.
	// "mlkem768x25519plus.native.1rtt.60". Leave empty to use a safe default.
	Decryption string `mapstructure:"Decryption"`
	// Encryption is the PQE encryption string applied to each VLESS user.
	// It must be compatible with Decryption. Leave empty to use a safe default.
	Encryption string `mapstructure:"Encryption"`
}

// TUNConfig configures a TUN inbound. Xray-core's TUN inbound requires a JSON
// settings blob; XrayR forwards SettingsRaw verbatim to xray-core.
type TUNConfig struct {
	Enable     bool            `mapstructure:"Enable"`
	Name       string          `mapstructure:"Name"`
	Tag        string          `mapstructure:"Tag"`
	SettingsRaw json.RawMessage `mapstructure:"Settings"`
}

// ProcessRouteConfig configures process-name based routing/audit.
type ProcessRouteConfig struct {
	Enable       bool     `mapstructure:"Enable"`
	RejectNames  []string `mapstructure:"RejectNames"`
	RejectPaths  []string `mapstructure:"RejectPaths"`
}

// MetricsConfig configures the Prometheus metrics endpoint.
type MetricsConfig struct {
	Enable bool   `mapstructure:"Enable"`
	Listen string `mapstructure:"Listen"` // e.g. "127.0.0.1:9090"
	Path   string `mapstructure:"Path"`   // e.g. "/metrics"
}

type AutoSpeedLimitConfig struct {
	Limit         int `mapstructure:"Limit"` // mbps
	WarnTimes     int `mapstructure:"WarnTimes"`
	LimitSpeed    int `mapstructure:"LimitSpeed"`    // mbps
	LimitDuration int `mapstructure:"LimitDuration"` // minute
}

type FallBackConfig struct {
	SNI              string `mapstructure:"SNI"`
	Alpn             string `mapstructure:"Alpn"`
	Path             string `mapstructure:"Path"`
	Dest             string `mapstructure:"Dest"`
	ProxyProtocolVer uint64 `mapstructure:"ProxyProtocolVer"`
}

type REALITYConfig struct {
	Show             bool     `mapstructure:"Show"`
	Dest             string   `mapstructure:"Dest"`
	ProxyProtocolVer uint64   `mapstructure:"ProxyProtocolVer"`
	ServerNames      []string `mapstructure:"ServerNames"`
	PrivateKey       string   `mapstructure:"PrivateKey"`
	MinClientVer     string   `mapstructure:"MinClientVer"`
	MaxClientVer     string   `mapstructure:"MaxClientVer"`
	MaxTimeDiff      uint64   `mapstructure:"MaxTimeDiff"`
	ShortIds         []string `mapstructure:"ShortIds"`
}
