package api

import (
	"encoding/json"
	"regexp"

	"github.com/xtls/xray-core/infra/conf"
)

const (
	UserNotModified = "users not modified"
	NodeNotModified = "node not modified"
	RuleNotModified = "rules not modified"
)

// Config API config
type Config struct {
	APIHost             string  `mapstructure:"ApiHost"`
	NodeID              int     `mapstructure:"NodeID"`
	Key                 string  `mapstructure:"ApiKey"`
	NodeType            string  `mapstructure:"NodeType"`
	EnableVless         bool    `mapstructure:"EnableVless"`
	VlessFlow           string  `mapstructure:"VlessFlow"`
	Timeout             int     `mapstructure:"Timeout"`
	SpeedLimit          float64 `mapstructure:"SpeedLimit"`
	DeviceLimit         int     `mapstructure:"DeviceLimit"`
	RuleListPath        string  `mapstructure:"RuleListPath"`
	DisableCustomConfig bool    `mapstructure:"DisableCustomConfig"`
}

// NodeStatus Node status
type NodeStatus struct {
	CPU    float64
	Mem    float64
	Disk   float64
	Uptime uint64
}

type NodeInfo struct {
	AcceptProxyProtocol bool
	Authority           string
	NodeType            string // Must be V2ray, Trojan, and Shadowsocks
	NodeID              int
	Port                uint32
	SpeedLimit          uint64 // Bps
	AlterID             uint16
	TransportProtocol   string
	FakeType            string
	Host                string
	Path                string
	EnableTLS           bool
	EnableSniffing      bool
	RouteOnly           bool
	EnableVless         bool
	VlessFlow           string
	CypherMethod        string
	ServerKey           string
	ServiceName         string
	Method              string
	Header              json.RawMessage
	HttpHeaders         map[string]*conf.StringList
	Headers             map[string]string
	NameServerConfig    []*conf.NameServerConfig
	EnableREALITY       bool
	REALITYConfig       *REALITYConfig
	Show                bool
	EnableTFO           bool
	Dest                string
	ProxyProtocolVer    uint64
	ServerNames         []string
	PrivateKey          string
	MinClientVer        string
	MaxClientVer        string
	MaxTimeDiff         uint64
	ShortIds            []string
	Xver                uint64
	Flow                string
	Security            string
	Key                 string
	RejectUnknownSni    bool
	// XHTTP (SplitHTTP) extended options for v26.x
	XHTTPConfig *XHTTPConfig
	// VLESS Post-Quantum Encryption (PQE) options for v26.6+
	EnablePQE bool
	PQEConfig *PQEConfig
}

// XHTTPConfig holds extended XHTTP/SplitHTTP transport options introduced in Xray-core v26.x.
// These fields map 1:1 to conf.SplitHTTPConfig and allow panels to leverage the new
// CDN-bypass and multiplexing capabilities.
type XHTTPConfig struct {
	Mode                 string          `json:"mode"`                 // auto, packet-up, stream-up, stream-one
	Path                 string          `json:"path"`
	Host                 string          `json:"host"`
	XPaddingObfsMode     bool            `json:"xPaddingObfsMode"`
	NoGRPCHeader         bool            `json:"noGRPCHeader"`
	ScMaxEachPostBytes   Int32Range      `json:"scMaxEachPostBytes"`
	ScMinPostsIntervalMs Int32Range      `json:"scMinPostsIntervalMs"`
	Xmux                 XmuxConfig      `json:"xmux"`
	Extra                json.RawMessage `json:"extra"`
}

// Int32Range mirrors conf.Int32Range for JSON marshalling.
type Int32Range struct {
	From int32 `json:"from"`
	To   int32 `json:"to"`
}

// XmuxConfig maps to conf.XmuxConfig.
type XmuxConfig struct {
	MaxConcurrency   Int32Range `json:"maxConcurrency"`
	MaxConnections   Int32Range `json:"maxConnections"`
	CMaxReuseTimes   Int32Range `json:"cMaxReuseTimes"`
	HMaxRequestTimes Int32Range `json:"hMaxRequestTimes"`
	HMaxReusableSecs Int32Range `json:"hMaxReusableSecs"`
	HKeepAlivePeriod int64      `json:"hKeepAlivePeriod"`
}

// PQEConfig holds VLESS Post-Quantum Encryption parameters (Xray-core v26.6.22+).
// When EnablePQE is true and the underlying xray-core supports PQE, the controller
// will enable post-quantum encryption for VLESS users.
type PQEConfig struct {
	// EncAlgo specifies the post-quantum encryption algorithm, e.g. "mlkem768x25519".
	EncAlgo string `json:"encAlgo"`
	// EncPassword is the optional password for the PQE layer.
	EncPassword string `json:"encPassword"`
}

type UserInfo struct {
	UID         int
	Email       string
	UUID        string
	Passwd      string
	Port        uint32
	AlterID     uint16
	Method      string
	SpeedLimit  uint64 // Bps
	DeviceLimit int
}

type OnlineUser struct {
	UID int
	IP  string
}

type UserTraffic struct {
	UID      int
	Email    string
	Upload   int64
	Download int64
}

type ClientInfo struct {
	APIHost  string
	NodeID   int
	Key      string
	NodeType string
}

type DetectRule struct {
	ID      int
	Pattern *regexp.Regexp
}

type DetectResult struct {
	UID    int
	RuleID int
}

type REALITYConfig struct {
	Dest             string
	ProxyProtocolVer uint64
	ServerNames      []string
	PrivateKey       string
	MinClientVer     string
	MaxClientVer     string
	MaxTimeDiff      uint64
	ShortIds         []string
}
