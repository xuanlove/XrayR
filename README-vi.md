# XrayR

[![](https://img.shields.io/badge/TgChat-@XrayR讨论-blue.svg)](https://t.me/XrayR_project)
[![](https://img.shields.io/badge/Channel-@XrayR通知-blue.svg)](https://t.me/XrayR_channel)
![](https://img.shields.io/github/stars/XrayR-project/XrayR)
![](https://img.shields.io/github/forks/XrayR-project/XrayR)
![](https://github.com/XrayR-project/XrayR/actions/workflows/release.yml/badge.svg)
![](https://github.com/XrayR-project/XrayR/actions/workflows/docker.yml/badge.svg)
[![Github All Releases](https://img.shields.io/github/downloads/XrayR-project/XrayR/total.svg)]()

[Chinese](https://github.com/XrayR-project/XrayR/blob/master/README.md) | [English](https://github.com/XrayR-project/XrayR/blob/master/README-en.md) | [Iranian](https://github.com/XrayR-project/XrayR/blob/master/README_Fa.md)

A Xray backend framework that can easily support many panels.

Khung back-end dựa trên Xray, hỗ trợ các giao thức V2ray, Trojan, Shadowsocks, dễ dàng mở rộng và hỗ trợ kết nối nhiều panel.

Nếu bạn thích dự án này, bạn có thể nhấp vào star+watch ở góc trên bên phải để tiếp tục theo dõi tiến trình của dự án này.

Hướng dẫn sử dụng: [Hướng dẫn chi tiết](https://xrayr-project.github.io/XrayR-doc/)

## Tuyên bố miễn trừ

Dự án này chỉ là học tập, phát triển và bảo trì cá nhân của tôi. Tôi không đảm bảo bất kỳ khả năng sẵn sàng nào và không chịu trách nhiệm cho bất kỳ hậu quả nào do việc sử dụng phần mềm này.

## Điểm nổi bật

* Mã nguồn mở vĩnh viễn và miễn phí.
* Dựa trên Xray-core v26.x (phiên bản nhân v1.260327.0), theo sát upstream.
* Hỗ trợ nhiều giao thức V2Ray, Trojan, Shadowsocks.
* Hỗ trợ các tính năng mới như VLESS, XTLS, REALITY.
* Hỗ trợ mã hóa hậu lượng tử VLESS (PQE, mlkem768x25519plus).
* Hỗ trợ truyền tải mở rộng XHTTP/SplitHTTP (Mode, XPaddingObfs, NoGRPCHeader, Xmux, v.v.).
* Hỗ trợ TUN inbound để proxy trong suốt trên máy chủ node.
* Hỗ trợ quy tắc định tuyến/kiểm toán dựa trên tên tiến trình.
* Hỗ trợ một instance kết nối nhiều panel, nhiều node, không cần khởi động lại nhiều lần.
* Hỗ trợ giới hạn IP trực tuyến cục bộ + toàn cục (Redis).
* Hỗ trợ giới hạn tốc độ cấp cổng node và cấp người dùng, cùng với tự động giới hạn tốc độ khi vượt tốc.
* Hỗ trợ Prometheus metrics endpoint.
* Cấu hình đơn giản và rõ ràng; sửa đổi cấu hình sẽ tự động khởi động lại instance.
* Dễ dàng biên dịch và nâng cấp, có thể nhanh chóng theo sát các tính năng mới của Xray-core.

## Chức năng

| Chức năng | V2ray | Trojan | Shadowsocks | SS-Plugin |
|-----------|:-----:|:------:|:-----------:|:---------:|
| Nhận thông tin node | √ | √ | √ | √ |
| Nhận thông tin người dùng | √ | √ | √ | √ |
| Báo cáo thống kê lưu lượng người dùng | √ | √ | √ | √ |
| Báo cáo thông tin máy chủ | √ | √ | √ | √ |
| Tự động đăng ký chứng chỉ TLS | √ | √ | √ | √ |
| Tự động gia hạn chứng chỉ TLS | √ | √ | √ | √ |
| Số người trực tuyến | √ | √ | √ | √ |
| Giới hạn người dùng trực tuyến (cục bộ + Redis toàn cục) | √ | √ | √ | √ |
| Quy tắc kiểm toán | √ | √ | √ | √ |
| Giới hạn tốc độ cổng node | √ | √ | √ | √ |
| Giới hạn tốc độ theo người dùng | √ | √ | √ | √ |
| Tự động giới hạn tốc độ khi vượt tốc | √ | √ | √ | √ |
| DNS tùy chỉnh | √ | √ | √ | √ |
| REALITY | √ | √ | — | — |
| VLESS PQE (Mã hóa hậu lượng tử) | √ | — | — | — |
| Truyền tải mở rộng XHTTP | √ | √ | √ | √ |
| TUN inbound | √ | √ | √ | √ |
| Định tuyến/kiểm toán theo tên tiến trình | √ | √ | √ | √ |
| Prometheus Metrics | √ | √ | √ | √ |

## Hỗ trợ Panel

| Panel | V2ray | Trojan | Shadowsocks |
|-------|:-----:|:------:|:-----------:|
| sspanel-uim | √ | √ | √ (Nhiều người dùng một cổng và V2Ray-Plugin) |
| v2board / NewV2board | √ | √ | √ |
| [PMPanel](https://github.com/ByteInternetHK/PMPanel) | √ | √ | √ |
| [ProxyPanel](https://github.com/ProxyPanel/ProxyPanel) | √ | √ | √ |
| [WHMCS (V2RaySocks)](https://v2raysocks.doxtex.com/) | √ | √ | √ |
| [GoV2Panel](https://github.com/pingProMax/gov2panel) | √ | √ | √ |
| [BunPanel](https://github.com/pennyMorant/bunpanel-release) | √ | √ | √ |

## Cài đặt phần mềm

### Cài đặt một chạm

```
wget -N https://raw.githubusercontent.com/XrayR-project/XrayR-release/master/install.sh && bash install.sh
```

### Triển khai bằng Docker

[Hướng dẫn triển khai Docker](https://xrayr-project.github.io/XrayR-doc/xrayr-xia-zai-he-an-zhuang/install/docker)

### Cài đặt thủ công

[Hướng dẫn cài đặt thủ công](https://xrayr-project.github.io/XrayR-doc/xrayr-xia-zai-he-an-zhuang/install/manual)

## Tệp cấu hình và hướng dẫn sử dụng chi tiết

[Hướng dẫn chi tiết](https://xrayr-project.github.io/XrayR-doc/)

### Tùy chọn cấu hình mới (v0.10.0+)

Các tùy chọn sau được đặt dưới `ControllerConfig` của mỗi mục `Nodes`:

```yaml
      # Mã hóa hậu lượng tử VLESS (Xray-core v26.6.22+, mlkem768x25519plus)
      # Chỉ có hiệu lực trên VLESS inbound. Yêu cầu chuỗi PQE đầy đủ (có padding).
      PQEConfig:
        Enable: false
        Decryption: # Định dạng: mlkem768x25519plus.<mode>.<from>-<to>s.<base64 padding>
        Encryption: # Phải tương thích với Decryption
      # TUN inbound (Xray-core v26.1.23+) để proxy trong suốt trên máy chủ node
      TUNConfig:
        Enable: false
        Name: tun0
        Tag: # Tùy chọn, mặc định tun_<nodeTag>
        Settings: {} # JSON thô chuyển tiếp cho TUN inbound của xray-core
      # Quy tắc định tuyến/kiểm toán theo tên tiến trình (Xray-core v26.1.23+)
      ProcessRouteConfig:
        Enable: false
        RejectNames: # Danh sách tên tiến trình cần từ chối, ví dụ ["BitTorrent", "Thunder"]
        RejectPaths: # Danh sách đường dẫn tệp thực thi cần từ chối
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

[Thảo luận back-end XrayR](https://t.me/XrayR_project)

[Thông báo XrayR](https://t.me/XrayR_channel)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/XrayR-project/XrayR.svg)](https://starchart.cc/XrayR-project/XrayR)
