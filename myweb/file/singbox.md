---
date: 2023-12-24
authors:
  - kian
categories:
  - 软件
---

# **SingBox**
**SingBox 配置文件 config.json**
<!-- more -->

## 新版配置
```json
{
  "log": {
    "disabled": true,
    "level": "debug",
    "timestamp": true
  },
  "experimental": {
    "clash_api": {
      "external_controller": "127.0.0.1:9090",
      "external_ui": "ui",
      "secret": "",
      "default_mode": "rule"
    },
    "cache_file": {
      "enabled": true,
      "store_fakeip": false
    }
  },
  "dns": {
    "servers": [
      {
        "tag": "proxyDns",
        "address": "tls://8.8.8.8",
        "detour": "proxy"
      },
      {
        "tag": "localDns",
        "address": "https://223.5.5.5/dns-query",
        "detour": "direct"
      },
      {
        "tag": "block",
        "address": "rcode://success"
      }
    ],
    "rules": [
      {
        "domain": [
          "ghproxy.com",
          "cdn.jsdelivr.net",
          "testingcf.jsdelivr.net"
        ],
        "server": "localDns"
      },
      {
        "rule_set": "geosite-category-ads-all",
        "server": "block"
      },
      {
        "outbound": "any",
        "server": "localDns",
        "disable_cache": true
      },
      {
        "rule_set": "geosite-cn",
        "server": "localDns"
      },
      {
        "clash_mode": "direct",
        "server": "localDns"
      },
      {
        "clash_mode": "global",
        "server": "proxyDns"
      },
      {
        "rule_set": "geosite-geolocation-!cn",
        "server": "proxyDns"
      }
    ]
  },
  "inbounds": [
    {
      "type": "tun",
      "inet4_address": "172.19.0.1/30",
      "inet6_address": "fdfe:dcba:9876::1/126",
      "mtu": 9000,
      "auto_route": true,
      "strict_route": true,
      "sniff": true,
      "endpoint_independent_nat": false,
      "stack": "system",
      "platform": {
        "http_proxy": {
          "enabled": true,
          "server": "127.0.0.1",
          "server_port": 2080
        }
      }
    },
    {
      "type": "mixed",
      "listen": "127.0.0.1",
      "listen_port": 2080,
      "sniff": true,
      "users": []
    }
  ],
  "outbounds": [
    {
      "tag": "proxy",
      "type": "selector",
      "outbounds": [
        "auto",
        "direct",
        "🇺🇸 US1",
        "🇺🇸 US2",
        "🇺🇸 US3",
        "🇺🇸 US4"
      ]
    },
    {
      "tag": "🤖 OpenAI",
      "type": "selector",
      "outbounds": [
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ],
      "default": "🇺🇸 美国节点"
    },
    {
      "tag": "🌌 Google",
      "type": "selector",
      "outbounds": [
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "📟 Telegram",
      "type": "selector",
      "outbounds": [
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "🐦 Twitter",
      "type": "selector",
      "outbounds": [
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "👤 Facebook",
      "type": "selector",
      "outbounds": [
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "🛍️ Amazon",
      "type": "selector",
      "outbounds": [
        "direct",
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "🍎 Apple",
      "type": "selector",
      "outbounds": [
        "direct",
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "🧩 Microsoft",
      "type": "selector",
      "outbounds": [
        "direct",
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "🎮 Game",
      "type": "selector",
      "outbounds": [
        "direct",
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "📺 Bilibili",
      "type": "selector",
      "outbounds": [
        "direct",
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点"
      ]
    },
    {
      "tag": "🎬 MediaVideo",
      "type": "selector",
      "outbounds": [
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "🌏 !cn",
      "type": "selector",
      "outbounds": [
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点",
        "direct"
      ]
    },
    {
      "tag": "🇭🇰 香港节点",
      "type": "selector",
      "outbounds": [
        "proxy"
      ]
    },
    {
      "tag": "🇹🇼 台湾节点",
      "type": "selector",
      "outbounds": [
        "proxy"
      ]
    },
    {
      "tag": "🇸🇬 狮城节点",
      "type": "selector",
      "outbounds": [
        "proxy"
      ]
    },
    {
      "tag": "🇯🇵 日本节点",
      "type": "selector",
      "outbounds": [
        "proxy"
      ]
    },
    {
      "tag": "🇺🇸 美国节点",
      "type": "selector",
      "outbounds": [
        "🇺🇸 US1",
        "🇺🇸 US2",
        "🇺🇸 US3",
        "🇺🇸 US4",
        "proxy"
      ]
    },
    {
      "tag": "✈️ 其他节点",
      "type": "selector",
      "outbounds": [
        "proxy"
      ]
    },
    {
      "tag": "🌏 cn",
      "type": "selector",
      "outbounds": [
        "direct",
        "proxy"
      ]
    },
    {
      "tag": "🛑 AdBlock",
      "type": "selector",
      "outbounds": [
        "block",
        "direct"
      ]
    },
    {
      "tag": "auto",
      "type": "urltest",
      "outbounds": [
        "🇺🇸 US1",
        "🇺🇸 US2",
        "🇺🇸 US3",
        "🇺🇸 US4"
      ],
      "url": "http://www.gstatic.com/generate_204",
      "interval": "5m",
      "tolerance": 50
    },
    {
      "type": "direct",
      "tag": "direct"
    },
    {
      "type": "dns",
      "tag": "dns-out"
    },
    {
      "type": "block",
      "tag": "block"
    },
    {
      "tag": "🇺🇸 US1",
      "type": "shadowsocks",
      "server": "ipv4",
      "server_port": 10000,
      "method": "2022-blake3-chacha20-poly1305",
      "password": "xxxxxxxx"
    },
    {
      "tag": "🇺🇸 US2",
      "type": "shadowsocks",
      "server": "[ipv6]",
      "server_port": 10000,
      "method": "2022-blake3-chacha20-poly1305",
      "password": "xxxxxxxx"
    },
    {
      "tag": "🇺🇸 US3",
      "type": "shadowsocks",
      "server": "[ipv6]",
      "server_port": 10000,
      "method": "2022-blake3-chacha20-poly1305",
      "password": "xxxxxxxx"
    },
    {
      "tag": "🇺🇸 US4",
      "type": "shadowsocks",
      "server": "[ipv6]",
      "server_port": 10000,
      "method": "2022-blake3-chacha20-poly1305",
      "password": "xxxxxxxx"
    }
  ],
  "route": {
    "auto_detect_interface": true,
    "final": "proxy",
    "rules": [
      {
        "protocol": "dns",
        "outbound": "dns-out"
      },
      {
        "network": "udp",
        "port": 443,
        "outbound": "block"
      },
      {
        "rule_set": "geosite-category-ads-all",
        "outbound": "🛑 AdBlock"
      },
      {
        "clash_mode": "direct",
        "outbound": "direct"
      },
      {
        "clash_mode": "global",
        "outbound": "proxy"
      },
      {
        "domain": [
          "clash.razord.top",
          "yacd.metacubex.one",
          "yacd.haishan.me",
          "d.metacubex.one"
        ],
        "outbound": "direct"
      },
      {
        "rule_set": "geosite-openai",
        "outbound": "🤖 OpenAI"
      },
      {
        "rule_set": "geosite-youtube",
        "outbound": "🌌 Google"
      },
      {
        "rule_set": "geoip-google",
        "outbound": "🌌 Google"
      },
      {
        "rule_set": [
          "geosite-google",
          "geosite-github"
        ],
        "outbound": "🌌 Google"
      },
      {
        "rule_set": "geoip-telegram",
        "outbound": "📟 Telegram"
      },
      {
        "rule_set": "geosite-telegram",
        "outbound": "📟 Telegram"
      },
      {
        "rule_set": "geoip-twitter",
        "outbound": "🐦 Twitter"
      },
      {
        "rule_set": "geosite-twitter",
        "outbound": "🐦 Twitter"
      },
      {
        "rule_set": "geoip-facebook",
        "outbound": "👤 Facebook"
      },
      {
        "rule_set": [
          "geosite-facebook",
          "geosite-instagram"
        ],
        "outbound": "👤 Facebook"
      },
      {
        "rule_set": "geosite-amazon",
        "outbound": "🛍️ Amazon"
      },
      {
        "rule_set": "geosite-apple",
        "outbound": "🍎 Apple"
      },
      {
        "rule_set": "geosite-microsoft",
        "outbound": "🧩 Microsoft"
      },
      {
        "rule_set": "geosite-category-games",
        "outbound": "🎮 Game"
      },
      {
        "rule_set": "geosite-bilibili",
        "outbound": "📺 Bilibili"
      },
      {
        "rule_set": "geoip-netflix",
        "outbound": "🎬 MediaVideo"
      },
      {
        "rule_set": [
          "geosite-tiktok",
          "geosite-netflix",
          "geosite-hbo",
          "geosite-disney",
          "geosite-primevideo"
        ],
        "outbound": "🎬 MediaVideo"
      },
      {
        "rule_set": "geosite-geolocation-!cn",
        "outbound": "🌏 !cn"
      },
      {
        "ip_is_private": true,
        "outbound": "🌏 cn"
      },
      {
        "rule_set": "geoip-cn",
        "outbound": "🌏 cn"
      },
      {
        "rule_set": "geosite-cn",
        "outbound": "🌏 cn"
      }
    ],
    "rule_set": [
      {
        "tag": "geoip-google",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geoip/google.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geoip-telegram",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geoip/telegram.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geoip-twitter",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geoip/twitter.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geoip-facebook",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geoip/facebook.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geoip-netflix",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geoip/netflix.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geoip-cn",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geoip/cn.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-openai",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/openai.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-youtube",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/youtube.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-google",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/google.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-github",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/github.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-telegram",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/telegram.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-twitter",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/twitter.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-facebook",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/facebook.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-instagram",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/instagram.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-amazon",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/amazon.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-apple",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/apple.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-microsoft",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/microsoft.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-category-games",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/category-games.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-bilibili",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/bilibili.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-tiktok",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/tiktok.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-netflix",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/netflix.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-hbo",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/hbo.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-disney",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/disney.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-primevideo",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/primevideo.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-cn",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/cn.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-geolocation-!cn",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/geolocation-!cn.srs",
        "download_detour": "direct"
      },
      {
        "tag": "geosite-category-ads-all",
        "type": "remote",
        "format": "binary",
        "url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@sing/geo/geosite/category-ads-all.srs",
        "download_detour": "direct"
      }
    ]
  }
}
```
## 旧版配置

```json
{
  "log": {
    "disabled": true,
    "level": "debug",
    "timestamp": true
  },
  "dns": {
    "servers": [
      {
        "tag": "proxyDns",
        "address": "tls://8.8.8.8",
        "detour": "proxy"
      },
      {
        "tag": "localDns",
        "address": "https://223.5.5.5/dns-query",
        "detour": "direct"
      },
      {
        "tag": "block",
        "address": "rcode://success"
      }
    ],
    "rules": [
      {
        "domain": [
          "ghproxy.com",
          "cdn.jsdelivr.net",
          "testingcf.jsdelivr.net"
        ],
        "server": "localDns"
      },
      {
        "geosite": "category-ads-all",
        "server": "block"
      },
      {
        "outbound": "any",
        "server": "localDns",
        "disable_cache": true
      },
      {
        "geosite": "cn",
        "server": "localDns"
      },
      {
        "clash_mode": "direct",
        "server": "localDns"
      },
      {
        "clash_mode": "global",
        "server": "proxyDns"
      },
      {
        "geosite": "geolocation-!cn",
        "server": "proxyDns"
      }
    ]
  },
  "inbounds": [
    {
      "type": "tun",
      "inet4_address": "172.19.0.1/30",
      "inet6_address": "fdfe:dcba:9876::1/126",
      "mtu": 9000,
      "auto_route": true,
      "strict_route": true,
      "sniff": true,
      "endpoint_independent_nat": false,
      "stack": "system",
      "platform": {
        "http_proxy": {
          "enabled": true,
          "server": "127.0.0.1",
          "server_port": 2080
        }
      }
    },
    {
      "type": "mixed",
      "listen": "127.0.0.1",
      "listen_port": 2080,
      "sniff": true,
      "users": []
    }
  ],
  "outbounds": [
    {
      "tag": "proxy",
      "type": "selector",
      "outbounds": [
        "auto",
        "direct",
        "🇺🇸 US1",
        "🇺🇸 US2",
        "🇺🇸 US3",
        "🇺🇸 US4"
      ]
    },
    {
      "tag": "🤖 OpenAI",
      "type": "selector",
      "outbounds": [
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ],
      "default": "🇺🇸 美国节点"
    },
    {
      "tag": "🌌 Google",
      "type": "selector",
      "outbounds": [
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "📟 Telegram",
      "type": "selector",
      "outbounds": [
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "🐦 Twitter",
      "type": "selector",
      "outbounds": [
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "👤 Facebook",
      "type": "selector",
      "outbounds": [
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "🛍️ Amazon",
      "type": "selector",
      "outbounds": [
        "direct",
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "🍎 Apple",
      "type": "selector",
      "outbounds": [
        "direct",
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "🧩 Microsoft",
      "type": "selector",
      "outbounds": [
        "direct",
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "🎮 Game",
      "type": "selector",
      "outbounds": [
        "direct",
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "📺 Bilibili",
      "type": "selector",
      "outbounds": [
        "direct",
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点"
      ]
    },
    {
      "tag": "🎬 MediaVideo",
      "type": "selector",
      "outbounds": [
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点"
      ]
    },
    {
      "tag": "🌏 !cn",
      "type": "selector",
      "outbounds": [
        "🇭🇰 香港节点",
        "🇹🇼 台湾节点",
        "🇸🇬 狮城节点",
        "🇯🇵 日本节点",
        "🇺🇸 美国节点",
        "✈️ 其他节点",
        "direct"
      ]
    },
    {
      "tag": "🇭🇰 香港节点",
      "type": "selector",
      "outbounds": [
        "proxy"
      ]
    },
    {
      "tag": "🇹🇼 台湾节点",
      "type": "selector",
      "outbounds": [
        "proxy"
      ]
    },
    {
      "tag": "🇸🇬 狮城节点",
      "type": "selector",
      "outbounds": [
        "proxy"
      ]
    },
    {
      "tag": "🇯🇵 日本节点",
      "type": "selector",
      "outbounds": [
        "proxy"
      ]
    },
    {
      "tag": "🇺🇸 美国节点",
      "type": "selector",
      "outbounds": [
        "🇺🇸 US1",
        "🇺🇸 US2",
        "🇺🇸 US3",
        "🇺🇸 US4",
        "proxy"
      ]
    },
    {
      "tag": "✈️ 其他节点",
      "type": "selector",
      "outbounds": [
        "proxy"
      ]
    },
    {
      "tag": "🌏 cn",
      "type": "selector",
      "outbounds": [
        "direct",
        "proxy"
      ]
    },
    {
      "tag": "🛑 AdBlock",
      "type": "selector",
      "outbounds": [
        "block",
        "direct"
      ]
    },
    {
      "tag": "auto",
      "type": "urltest",
      "outbounds": [
        "🇺🇸 US1",
        "🇺🇸 US2",
        "🇺🇸 US3",
        "🇺🇸 US4"
      ],
      "url": "http://www.gstatic.com/generate_204",
      "interval": "5m",
      "tolerance": 50
    },
    {
      "type": "direct",
      "tag": "direct"
    },
    {
      "type": "dns",
      "tag": "dns-out"
    },
    {
      "type": "block",
      "tag": "block"
    },
    {
      "tag": "🇺🇸 US1",
      "type": "shadowsocks",
      "server": "ipv4",
      "server_port": 10000,
      "method": "2022-blake3-chacha20-poly1305",
      "password": "xxxxxxxx"
    },
    {
      "tag": "🇺🇸 US2",
      "type": "shadowsocks",
      "server": "[ipv6]",
      "server_port": 10000,
      "method": "2022-blake3-chacha20-poly1305",
      "password": "xxxxxxxx"
    },
    {
      "tag": "🇺🇸 US3",
      "type": "shadowsocks",
      "server": "[ipv6]",
      "server_port": 10000,
      "method": "2022-blake3-chacha20-poly1305",
      "password": "xxxxxxxx"
    },
    {
      "tag": "🇺🇸 US4",
      "type": "shadowsocks",
      "server": "[ipv6]",
      "server_port": 10000,
      "method": "2022-blake3-chacha20-poly1305",
      "password": "xxxxxxxx"
    }
  ],
  "route": {
    "auto_detect_interface": true,
    "final": "proxy",
    "geoip": {
      "download_url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@release/geoip.db",
      "download_detour": "direct"
    },
    "geosite": {
      "download_url": "https://testingcf.jsdelivr.net/gh/MetaCubeX/meta-rules-dat@release/geosite.db",
      "download_detour": "direct"
    },
    "rules": [
      {
        "protocol": "dns",
        "outbound": "dns-out"
      },
      {
        "network": "udp",
        "port": 443,
        "outbound": "block"
      },
      {
        "geosite": "category-ads-all",
        "outbound": "🛑 AdBlock"
      },
      {
        "clash_mode": "direct",
        "outbound": "direct"
      },
      {
        "clash_mode": "global",
        "outbound": "proxy"
      },
      {
        "domain": [
          "clash.razord.top",
          "yacd.metacubex.one",
          "yacd.haishan.me",
          "d.metacubex.one"
        ],
        "outbound": "direct"
      },
      {
        "geosite": "openai",
        "outbound": "🤖 OpenAI"
      },
      {
        "geosite": "youtube",
        "outbound": "🌌 Google"
      },
      {
        "geoip": "google",
        "outbound": "🌌 Google"
      },
      {
        "geosite": [
          "google",
          "github"
        ],
        "outbound": "🌌 Google"
      },
      {
        "geoip": "telegram",
        "outbound": "📟 Telegram"
      },
      {
        "geosite": "telegram",
        "outbound": "📟 Telegram"
      },
      {
        "geoip": "twitter",
        "outbound": "🐦 Twitter"
      },
      {
        "geosite": "twitter",
        "outbound": "🐦 Twitter"
      },
      {
        "geoip": "facebook",
        "outbound": "👤 Facebook"
      },
      {
        "geosite": [
          "facebook",
          "instagram"
        ],
        "outbound": "👤 Facebook"
      },
      {
        "geosite": "amazon",
        "outbound": "🛍️ Amazon"
      },
      {
        "geosite": "apple",
        "outbound": "🍎 Apple"
      },
      {
        "geosite": "apple-cn",
        "outbound": "🍎 Apple"
      },
      {
        "geosite": "microsoft",
        "outbound": "🧩 Microsoft"
      },
      {
        "geosite": "category-games",
        "outbound": "🎮 Game"
      },
      {
        "geosite": "bilibili",
        "outbound": "📺 Bilibili"
      },
      {
        "geoip": "netflix",
        "outbound": "🎬 MediaVideo"
      },
      {
        "geosite": [
          "tiktok",
          "netflix",
          "hbo",
          "disney",
          "primevideo"
        ],
        "outbound": "🎬 MediaVideo"
      },
      {
        "geosite": "geolocation-!cn",
        "outbound": "🌏 !cn"
      },
      {
        "geoip": [
          "private",
          "cn"
        ],
        "outbound": "🌏 cn"
      },
      {
        "geosite": "cn",
        "outbound": "🌏 cn"
      }
    ]
  },
  "experimental": {
    "clash_api": {
      "external_controller": "127.0.0.1:9090",
      "external_ui": "ui",
      "secret": "",
      "default_mode": "rule",
      "store_selected": true
    }
  }
}
```
