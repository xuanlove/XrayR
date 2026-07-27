# XrayR

[![](https://img.shields.io/badge/TgChat-@XrayR讨论-blue.svg)](https://t.me/XrayR_project)
[![](https://img.shields.io/badge/Channel-@XrayR通知-blue.svg)](https://t.me/XrayR_channel)
![](https://img.shields.io/github/stars/XrayR-project/XrayR)
![](https://img.shields.io/github/forks/XrayR-project/XrayR)
![](https://github.com/XrayR-project/XrayR/actions/workflows/release.yml/badge.svg)
![](https://github.com/XrayR-project/XrayR/actions/workflows/docker.yml/badge.svg)
[![Github All Releases](https://img.shields.io/github/downloads/XrayR-project/XrayR/total.svg)]()

[Chinese](https://github.com/XrayR-project/XrayR/blob/master/README.md) | [Iranian](https://github.com/XrayR-project/XrayR/blob/master/README_Fa.md) | [Vietnamese](https://github.com/XrayR-project/XrayR/blob/master/README-vi.md)

A Xray backend framework that can easily support many panels.

A back-end framework based on Xray that supports V2ray, Trojan and Shadowsocks protocols, easy to extend and supports multi-panel docking.

If you like this project, you can click star+watch in the upper right corner to keep following the progress of this project.

Tutorial: [Detailed tutorial](https://xrayr-project.github.io/XrayR-doc/)

## Disclaimer

This project is just my personal learning, development and maintenance. I do not guarantee any availability, nor am I responsible for any consequences caused by the use of this software.

## Features

* Permanently open source and free.
* Based on Xray-core v26.x (kernel version v1.260327.0), closely following upstream.
* Supports V2Ray, Trojan, Shadowsocks protocols.
* Supports VLESS, XTLS, REALITY and other new features.
* Supports VLESS Post-Quantum Encryption (PQE, mlkem768x25519plus).
* Supports XHTTP/SplitHTTP extended transport (Mode, XPaddingObfs, NoGRPCHeader, Xmux, etc.).
* Supports TUN inbound for transparent proxying on the node host.
* Supports process-name based routing/audit rules.
* Supports single instance docking multiple panels and multiple nodes without repeated startup.
* Supports local + global (Redis) online IP restriction.
* Supports node port level and user level speed limit, plus automatic over-speed limiting.
* Supports Prometheus metrics endpoint exposure.
* Simple and clear configuration; modifying config triggers automatic hot restart.
* Easy to compile and upgrade, can quickly follow Xray-core new features.

## Function

| Function | V2ray | Trojan | Shadowsocks | SS-Plugin |
|----------|:-----:|:------:|:-----------:|:---------:|
| Get node information | √ | √ | √ | √ |
| Get user information | √ | √ | √ | √ |
| User traffic statistics report | √ | √ | √ | √ |
| Server information report | √ | √ | √ | √ |
| Automatically apply for TLS certificate | √ | √ | √ | √ |
| Automatic renewal of TLS certificate | √ | √ | √ | √ |
| Online user count | √ | √ | √ | √ |
| Online user limit (local + global Redis) | √ | √ | √ | √ |
| Audit rules | √ | √ | √ | √ |
| Node port speed limit | √ | √ | √ | √ |
| Per-user speed limit | √ | √ | √ | √ |
| Automatic over-speed limiting | √ | √ | √ | √ |
| Custom DNS | √ | √ | √ | √ |
| REALITY | √ | √ | — | — |
| VLESS PQE (Post-Quantum Encryption) | √ | — | — | — |
| XHTTP extended transport | √ | √ | √ | √ |
| TUN inbound | √ | √ | √ | √ |
| Process-name routing/audit | √ | √ | √ | √ |
| Prometheus Metrics | √ | √ | √ | √ |

## Support for panels

| Panel | V2ray | Trojan | Shadowsocks |
|-------|:-----:|:------:|:-----------:|
| sspanel-uim | √ | √ | √ (Single-port multi-user and V2Ray-Plugin) |
| v2board / NewV2board | √ | √ | √ |
| [PMPanel](https://github.com/ByteInternetHK/PMPanel) | √ | √ | √ |
| [ProxyPanel](https://github.com/ProxyPanel/ProxyPanel) | √ | √ | √ |
| [WHMCS (V2RaySocks)](https://v2raysocks.doxtex.com/) | √ | √ | √ |
| [GoV2Panel](https://github.com/pingProMax/gov2panel) | √ | √ | √ |
| [BunPanel](https://github.com/pennyMorant/bunpanel-release) | √ | √ | √ |

## Software Installation

### 1-Click installation

```
wget -N https://raw.githubusercontent.com/XrayR-project/XrayR-release/master/install.sh && bash install.sh
```

### Docker

[Docker deployment tutorial](https://xrayr-project.github.io/XrayR-doc/xrayr-xia-zai-he-an-zhuang/install/docker)

### Manual installation

[Manual installation tutorial](https://xrayr-project.github.io/XrayR-doc/xrayr-xia-zai-he-an-zhuang/install/manual)

## Configuration file and detailed tutorial

[Detailed tutorial](https://xrayr-project.github.io/XrayR-doc/)

### New config options (v0.10.0+)

The following options are placed under `ControllerConfig` of each `Nodes` entry:

```yaml
      # VLESS Post-Quantum Encryption (Xray-core v26.6.22+, mlkem768x25519plus)
      # Only takes effect on VLESS inbound. A full PQE string (with padding) is required.
      PQEConfig:
        Enable: false
        Decryption: # Format: mlkem768x25519plus.<mode>.<from>-<to>s.<base64 padding>
        Encryption: # Must be compatible with Decryption
      # TUN inbound (Xray-core v26.1.23+) for transparent proxying on the node host
      TUNConfig:
        Enable: false
        Name: tun0
        Tag: # Optional, defaults to tun_<nodeTag>
        Settings: {} # Raw JSON forwarded to xray-core's TUN inbound
      # Process-name based routing/audit rules (Xray-core v26.1.23+)
      ProcessRouteConfig:
        Enable: false
        RejectNames: # List of process names to reject, e.g. ["BitTorrent", "Thunder"]
        RejectPaths: # List of executable paths to reject
      # Prometheus metrics endpoint (Xray-core app/metrics)
      MetricsConfig:
        Enable: false
        Listen: 127.0.0.1:9090
        Path: /metrics
```

## Thanks

* [Project X](https://github.com/XTLS/)
* [V2Fly](https://github.com/v2fly)
* [VNet-V2ray](https://github.com/ProxyPanel/VNet-V2ray)
* [Air-Universe](https://github.com/crossfw/Air-Universe)

## Licence

[Mozilla Public License Version 2.0](https://github.com/XrayR-project/XrayR/blob/master/LICENSE)

## Telegram

[XrayR back-end discussion](https://t.me/XrayR_project)

[XrayR notification](https://t.me/XrayR_channel)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/XrayR-project/XrayR.svg)](https://starchart.cc/XrayR-project/XrayR)
