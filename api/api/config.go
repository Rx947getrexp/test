package api

/*
{
  "inbounds": [
    {
      "port": 1080,
      "protocol": "socks",
      "settings": {
        "auth": "noauth"
      }
    }
  ],
  "outbounds": [
    {
      "protocol": "freedom",
      "settings": {},
      "tag": "direct"
    },
    {
      "protocol": "blackhole",
      "settings": {},
      "tag": "blocked"
    }
  ],
  "routing": {
    "rules": [
      {
        "type": "field",
        "domain": ["*.ru"],
        "outboundTag": "direct"
      },
      {
        "type": "field",
        "domain": ["*"],
        "outboundTag": "default" // 确保其他所有流量都通过默认的代理规则
      }
    ]
  }
}


{
  "inbounds": [
    {
      "port": 1080,
      "protocol": "socks",
      "settings": {
        "auth": "noauth"
      }
    }
  ],
  "outbounds": [
    {
      "protocol": "freedom",
      "settings": {}
    },
    {
      "protocol": "blackhole",
      "settings": {},
      "tag": "blocked"
    }
  ],
  "routing": {
    "rules": [
      {
        "type": "ip",
        "ip": [
          "192.0.2.1",
          "192.0.2.0/24",
          "0.0.0.0/0"
        ],
        "outboundTag": "direct"
      },
      {
        "type": "ip",
        "ip": ["0.0.0.0/0"],
        "outboundTag": "freedom"
      }
    ]
  }
}



*/
