# XrayR

[![](https://img.shields.io/badge/TgChat-@XrayR讨论-blue.svg)](https://t.me/XrayR_project)
[![](https://img.shields.io/badge/Channel-@XrayR通知-blue.svg)](https://t.me/XrayR_channel)
![](https://img.shields.io/github/stars/XrayR-project/XrayR)
![](https://img.shields.io/github/forks/XrayR-project/XrayR)
![](https://github.com/XrayR-project/XrayR/actions/workflows/release.yml/badge.svg)
![](https://github.com/XrayR-project/XrayR/actions/workflows/docker.yml/badge.svg)
[![Github All Releases](https://img.shields.io/github/downloads/XrayR-project/XrayR/total.svg)]()

[Chinese](https://github.com/XrayR-project/XrayR/blob/master/README.md) | [English](https://github.com/XrayR-project/XrayR/blob/master/README-en.md) | [Vietnamese](https://github.com/XrayR-project/XrayR/blob/master/README-vi.md)

A Xray backend framework that can easily support many panels.

یک فریمورک بک‌اند مبتنی بر Xray که از پروتکل‌های V2ray، Trojan و Shadowsocks پشتیبانی می‌کند، به راحتی قابل گسترش است و از اتصال چند پنل پشتیبانی می‌کند.

اگر این پروژه را دوست دارید، می‌توانید در گوشه بالا سمت راست روی star+watch کلیک کنید تا روند پیشرفت این پروژه را دنبال کنید.

آموزش: [آموزش با جزئیات](https://xrayr-project.github.io/XrayR-doc/)

## سلب مسئولیت

این پروژه فقط مطالعه، توسعه و نگهداری شخصی من است. من هیچ‌گونه قابلیت استفاده‌ای را تضمین نمی‌کنم و مسئولیتی در قبال عواقب ناشی از استفاده از این نرم‌افزار ندارم.

## ویژگی‌ها

* متن‌باز دائمی و رایگان.
* مبتنی بر Xray-core v26.x (نسخه هسته v1.260327.0)، با پیگیری نزدیک به upstream.
* پشتیبانی از پروتکل‌های V2Ray، Trojan، Shadowsocks.
* پشتیبانی از ویژگی‌های جدید مانند VLESS، XTLS، REALITY.
* پشتیبانی از رمزنگاری پساکوانتومی VLESS (PQE، mlkem768x25519plus).
* پشتیبانی از انتقال توسعه‌یافته XHTTP/SplitHTTP (Mode، XPaddingObfs، NoGRPCHeader، Xmux و غیره).
* پشتیبانی از TUN inbound برای پروکسی شفاف روی میزبان نود.
* پشتیبانی از قوانین مسیریابی/حسابرسی مبتنی بر نام فرآیند.
* پشتیبانی از اتصال یک نمونه به چند پنل و چند نود بدون نیاز به شروع مکرر.
* پشتیبانی از محدودیت IP آنلاین محلی + سراسری (Redis).
* پشتیبانی از محدودیت سرعت در سطح پورت نود و سطح کاربر، به همراه محدودیت خودکار سرعت بیش از حد.
* پشتیبانی از در معرض قرار دادن نقطه پایانی Prometheus metrics.
* پیکربندی ساده و واضح؛ تغییر پیکربندی به طور خودکار نمونه را راه‌اندازی مجدد می‌کند.
* کامپایل و ارتقاء آسان، می‌تواند به سرعت ویژگی‌های جدید Xray-core را دنبال کند.

## امکانات

| امکانات | V2ray | Trojan | Shadowsocks | SS-Plugin |
|--------|:-----:|:------:|:-----------:|:---------:|
| دریافت اطلاعات نود | √ | √ | √ | √ |
| دریافت اطلاعات کاربر | √ | √ | √ | √ |
| گزارش آمار ترافیک کاربر | √ | √ | √ | √ |
| گزارش اطلاعات سرور | √ | √ | √ | √ |
| درخواست خودکار گواهی TLS | √ | √ | √ | √ |
| تمدید خودکار گواهی TLS | √ | √ | √ | √ |
| تعداد کاربران آنلاین | √ | √ | √ | √ |
| محدودیت کاربر آنلاین (محلی + Redis سراسری) | √ | √ | √ | √ |
| قوانین حسابرسی | √ | √ | √ | √ |
| محدودیت سرعت پورت نود | √ | √ | √ | √ |
| محدودیت سرعت بر اساس کاربر | √ | √ | √ | √ |
| محدودیت خودکار سرعت بیش از حد | √ | √ | √ | √ |
| DNS سفارشی | √ | √ | √ | √ |
| REALITY | √ | √ | — | — |
| VLESS PQE (رمزنگاری پساکوانتومی) | √ | — | — | — |
| انتقال توسعه‌یافته XHTTP | √ | √ | √ | √ |
| TUN inbound | √ | √ | √ | √ |
| مسیریابی/حسابرسی بر اساس نام فرآیند | √ | √ | √ | √ |
| Prometheus Metrics | √ | √ | √ | √ |

## پشتیبانی از پنل‌ها

| پنل | V2ray | Trojan | Shadowsocks |
|-----|:-----:|:------:|:-----------:|
| sspanel-uim | √ | √ | √ (چند کاربره تک‌پورت و V2Ray-Plugin) |
| v2board / NewV2board | √ | √ | √ |
| [PMPanel](https://github.com/ByteInternetHK/PMPanel) | √ | √ | √ |
| [ProxyPanel](https://github.com/ProxyPanel/ProxyPanel) | √ | √ | √ |
| [WHMCS (V2RaySocks)](https://v2raysocks.doxtex.com/) | √ | √ | √ |
| [GoV2Panel](https://github.com/pingProMax/gov2panel) | √ | √ | √ |
| [BunPanel](https://github.com/pennyMorant/bunpanel-release) | √ | √ | √ |

## نصب نرم‌افزار

### نصب یک‌کلیکی

```
wget -N https://raw.githubusercontent.com/XrayR-project/XrayR-release/master/install.sh && bash install.sh
```

### استقرار با Docker

[آموزش استقرار Docker](https://xrayr-project.github.io/XrayR-doc/xrayr-xia-zai-he-an-zhuang/install/docker)

### نصب دستی

[آموزش نصب دستی](https://xrayr-project.github.io/XrayR-doc/xrayr-xia-zai-he-an-zhuang/install/manual)

## فایل پیکربندی و آموزش جامع

[آموزش جامع](https://xrayr-project.github.io/XrayR-doc/)

### گزینه‌های پیکربندی جدید (v0.10.0+)

گزینه‌های زیر زیر `ControllerConfig` هر ورودی `Nodes` قرار می‌گیرند:

```yaml
      # رمزنگاری پساکوانتومی VLESS (Xray-core v26.6.22+، mlkem768x25519plus)
      # فقط روی VLESS inbound اعمال می‌شود. رشته PQE کامل (با padding) لازم است.
      PQEConfig:
        Enable: false
        Decryption: # قالب: mlkem768x25519plus.<mode>.<from>-<to>s.<base64 padding>
        Encryption: # باید با Decryption سازگار باشد
      # TUN inbound (Xray-core v26.1.23+) برای پروکسی شفاف روی میزبان نود
      TUNConfig:
        Enable: false
        Name: tun0
        Tag: # اختیاری، پیش‌فرض tun_<nodeTag>
        Settings: {} # JSON خام به TUN inbound xray-core ارسال می‌شود
      # قوانین مسیریابی/حسابرسی مبتنی بر نام فرآیند (Xray-core v26.1.23+)
      ProcessRouteConfig:
        Enable: false
        RejectNames: # فهرست نام فرآیندهای مورد رد، مثلاً ["BitTorrent", "Thunder"]
        RejectPaths: # فهرست مسیرهای فایل اجرایی مورد رد
      # نقطه پایانی Prometheus metrics (Xray-core app/metrics)
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

[بحث بک‌اند XrayR](https://t.me/XrayR_project)

[کانال اعلان XrayR](https://t.me/XrayR_channel)

## Stargazers over time

[![Stargazers over time](https://starchart.cc/XrayR-project/XrayR.svg)](https://starchart.cc/XrayR-project/XrayR)
