# XrayR

[![](https://img.shields.io/badge/TgChat-@XrayR讨论-blue.svg)](https://t.me/XrayR_project)
[![](https://img.shields.io/badge/Channel-@XrayR通知-blue.svg)](https://t.me/XrayR_channel)
![](https://img.shields.io/github/stars/XrayR-project/XrayR)
![](https://img.shields.io/github/forks/XrayR-project/XrayR)
![](https://github.com/XrayR-project/XrayR/actions/workflows/release.yml/badge.svg)
![](https://github.com/XrayR-project/XrayR/actions/workflows/docker.yml/badge.svg)
[![Github All Releases](https://img.shields.io/github/downloads/XrayR-project/XrayR/total.svg)]()

[English](https://github.com/XrayR-project/XrayR/blob/master/README-en.md) | [Iranian](https://github.com/XrayR-project/XrayR/blob/master/README_Fa.md) | [Vietnamese](https://github.com/XrayR-project/XrayR/blob/master/README-vi.md)

A Xray backend framework that can easily support many panels.

一个基于 Xray 的后端框架，支持 V2ray、Trojan、Shadowsocks 协议，极易扩展，支持多面板对接。

如果您喜欢本项目，可以右上角点个 star+watch，持续关注本项目的进展。

使用教程：[详细使用教程](https://xrayr-project.github.io/XrayR-doc/)

## 免责声明

本项目只是本人个人学习开发并维护，本人不保证任何可用性，也不对使用本软件造成的任何后果负责。

## 特点

* 永久开源且免费。
* 基于 Xray-core v26.x（内核版本 v1.260327.0），紧跟上游最新特性。
* 支持 V2ray、Trojan、Shadowsocks 多种协议。
* 支持 VLESS、XTLS、REALITY 等新特性。
* 支持 VLESS 后量子加密（PQE，mlkem768x25519plus）。
* 支持 XHTTP/SplitHTTP 扩展传输（Mode、XPaddingObfs、NoGRPCHeader、Xmux 等）。
* 支持 TUN 入站，便于节点本机透明代理。
* 支持基于进程名的路由/审计规则。
* 支持单实例对接多面板、多节点，无需重复启动。
* 支持本地 + 全局（Redis）在线 IP 限制。
* 支持节点端口级别、用户级别限速，以及超速自动限速。
* 支持 Prometheus Metrics 端点暴露。
* 配置简单明了，修改配置自动热重启实例。
* 方便编译和升级，可快速跟进 Xray-core 新特性。

## 功能介绍

| 功能 | V2ray | Trojan | Shadowsocks | SS-Plugin |
|------|:-----:|:------:|:-----------:|:---------:|
| 获取节点信息 | √ | √ | √ | √ |
| 获取用户信息 | √ | √ | √ | √ |
| 用户流量统计上报 | √ | √ | √ | √ |
| 服务器信息上报 | √ | √ | √ | √ |
| 自动申请 TLS 证书 | √ | √ | √ | √ |
| 自动续签 TLS 证书 | √ | √ | √ | √ |
| 在线人数统计 | √ | √ | √ | √ |
| 在线用户限制（本地 + 全局 Redis） | √ | √ | √ | √ |
| 审计规则 | √ | √ | √ | √ |
| 节点端口限速 | √ | √ | √ | √ |
| 按用户限速 | √ | √ | √ | √ |
| 超速自动限速 | √ | √ | √ | √ |
| 自定义 DNS | √ | √ | √ | √ |
| REALITY | √ | √ | — | — |
| VLESS PQE（后量子加密） | √ | — | — | — |
| XHTTP 扩展传输 | √ | √ | √ | √ |
| TUN 入站 | √ | √ | √ | √ |
| 进程名路由/审计 | √ | √ | √ | √ |
| Prometheus Metrics | √ | √ | √ | √ |

## 支持前端

| 前端 | V2ray | Trojan | Shadowsocks |
|------|:-----:|:------:|:-----------:|
| sspanel-uim | √ | √ | √ (单端口多用户和 V2ray-Plugin) |
| v2board / NewV2board | √ | √ | √ |
| [PMPanel](https://github.com/ByteInternetHK/PMPanel) | √ | √ | √ |
| [ProxyPanel](https://github.com/ProxyPanel/ProxyPanel) | √ | √ | √ |
| [WHMCS (V2RaySocks)](https://v2raysocks.doxtex.com/) | √ | √ | √ |
| [GoV2Panel](https://github.com/pingProMax/gov2panel) | √ | √ | √ |
| [BunPanel](https://github.com/pennyMorant/bunpanel-release) | √ | √ | √ |

## 软件安装

### 一键安装

```
wget -N https://raw.githubusercontent.com/XrayR-project/XrayR-release/master/install.sh && bash install.sh
```

### 使用 Docker 部署

[Docker 部署教程](https://xrayr-project.github.io/XrayR-doc/xrayr-xia-zai-he-an-zhuang/install/docker)

### 手动安装

[手动安装教程](https://xrayr-project.github.io/XrayR-doc/xrayr-xia-zai-he-an-zhuang/install/manual)

## 配置文件及详细使用教程

[详细使用教程](https://xrayr-project.github.io/XrayR-doc/)

### 新增配置项（v0.10.0+）

以下配置项位于 `Nodes` 的 `ControllerConfig` 下：

```yaml
      # VLESS 后量子加密（Xray-core v26.6.22+，mlkem768x25519plus）
      # 仅对 VLESS 入站生效。启用后必须提供完整的 PQE 字符串（含填充材料）。
      PQEConfig:
        Enable: false
        Decryption: # 格式: mlkem768x25519plus.<mode>.<from>-<to>s.<base64 padding>
        Encryption: # 必须与 Decryption 兼容
      # TUN 入站（Xray-core v26.1.23+），用于节点本机透明代理
      TUNConfig:
        Enable: false
        Name: tun0
        Tag: # 可选，默认 tun_<nodeTag>
        Settings: {} # 转发给 xray-core TUN 入站的原始 JSON
      # 基于进程名的路由/审计规则（Xray-core v26.1.23+）
      ProcessRouteConfig:
        Enable: false
        RejectNames: # 拒绝的进程名列表，如 ["BitTorrent", "Thunder"]
        RejectPaths: # 拒绝的进程可执行文件路径列表
      # Prometheus Metrics 端点（Xray-core app/metrics）
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

## Telgram

[XrayR 后端讨论](https://t.me/XrayR_project)

[XrayR 通知](https://t.me/XrayR_channel)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/XrayR-project/XrayR.svg)](https://starchart.cc/XrayR-project/XrayR)
