package muggle

import (
	"magic"
	"testing"
	"v2ray.com/core/common"
)

func TestStartInstanceConflict(t *testing.T) {
	cfgStr := `
	{
	   "policy": {
	       "levels": {
	           "1": {
	               "handshake": 4,
	               "connIdle": 300,
	               "uplinkOnly": 0,
	               "downlinkOnly": 0,
	               "statsUserUplink": false,
	               "statsUserDownlink": false,
	               "bufferSize": 10240
	           }
	       }
	   },
	   "dns": {
	       "servers": [
	           "localhost",
	           "223.5.5.5",
	           "1.1.1.1",
	           "8.8.8.8"
	       ]
	   },
	   "inbounds": [
	       {
	           "settings": {
	               "ip": "127.0.0.1",
	               "udp": true,
	               "allowTransparent": true
	           },
	           "port": 9999,
	           "protocol": "socks",
	           "listen": "127.0.0.1",
	           "sniffing": {
	               "enabled": true,
	               "destOverride": [
	                   "http",
	                   "tls"
	               ]
	           },
	           "streamSettings": {
	               "auth": "noauth",
	               "udp": true,
	               "ip": "127.0.0.1"
	           }
	       }
	   ],
	   "outbounds": [
	       {
	           "settings": {
	               "vnext": [
	                   {
	                       "address": "peozeo.xyz",
	                       "port": 443,
	                       "udp": true,
	                       "users": [
	                           {
	                               "level": 1,
	                               "alterId": 16,
	                               "security": "aes-128-gcm",
	                               "id": "66d88922-c880-4b5a-906e-ff50fc8c209f"
	                           }
	                       ]
	                   }
	               ]
	           },
	           "streamSettings": {
	               "network": "ws",
	               "security": "tls",
	               "sockopt": {
	                   "tcpFastOpen": true
	               },
	               "wsSettings": {
	                   "connectionReuse": true,
	                   "path": "/llc"
	               },
	               "tlsSettings": {
	                   "allowInsecure": true
	               }
	           },
	           "mux": {
	               "enabled": true,
	               "concurrency": 8
	           },
	           "protocol": "vmess",
	           "tag": "proxy"
	       },
	       {
	           "settings": {
	               "domainStrategy": "UseIP"
	           },
	           "streamSettings": {
	               "sockopt": {
	                   "tcpFastOpen": true
	               }
	           },
	           "protocol": "freedom",
	           "tag": "direct"
	       },
	       {
	           "protocol": "blackhole",
	           "tag": "block",
	           "settings": {}
	       }
	   ],
	   "routing": {
	       "strategy": "rules",
	       "settings": {
	           "domainStrategy": "IPIfNonMatch",
	           "rules": [
	               {
	                   "type": "field",
	                   "outboundTag": "direct",
	                   "ip": [
	                       "0.0.0.0/8",
	                       "10.0.0.0/8",
	                       "100.64.0.0/10",
	                       "127.0.0.0/8",
	                       "169.254.0.0/16",
	                       "172.16.0.0/12",
	                       "192.0.0.0/24",
	                       "192.0.2.0/24",
	                       "192.168.0.0/16",
	                       "198.18.0.0/15",
	                       "198.51.100.0/24",
	                       "203.0.113.0/24",
	                       "::1/128",
	                       "fc00::/7",
	                       "fe80::/10"
	                   ]
	               }
	           ]
	       }
	   }
	}
		`

	pbBytes, err := magic.GetPbConfigBytes([]byte(cfgStr))

	var sig = make(chan int)

	go func(pbBytes []byte) {
		err = StartInstance(pbBytes)

		common.Must(err)
	}(pbBytes)

	go func(pbBytes []byte) {
		err = StartInstance(pbBytes)

		common.Must(err)
	}(pbBytes)

	<-sig

}

func TestStartInstance(t *testing.T) {
	cfgStr := `
	{
	   "policy": {
	       "levels": {
	           "1": {
	               "handshake": 4,
	               "connIdle": 300,
	               "uplinkOnly": 0,
	               "downlinkOnly": 0,
	               "statsUserUplink": false,
	               "statsUserDownlink": false,
	               "bufferSize": 10240
	           }
	       }
	   },
	   "dns": {
	       "servers": [
	           "localhost",
	           "223.5.5.5",
	           "1.1.1.1",
	           "8.8.8.8"
	       ]
	   },
	   "inbounds": [
	       {
	           "settings": {
	               "ip": "127.0.0.1",
	               "udp": true,
	               "allowTransparent": true
	           },
	           "port": 9999,
	           "protocol": "socks",
	           "listen": "127.0.0.1",
	           "sniffing": {
	               "enabled": true,
	               "destOverride": [
	                   "http",
	                   "tls"
	               ]
	           },
	           "streamSettings": {
	               "auth": "noauth",
	               "udp": true,
	               "ip": "127.0.0.1"
	           }
	       }
	   ],
	   "outbounds": [
	       {
	           "settings": {
	               "vnext": [
	                   {
	                       "address": "peozeo.xyz",
	                       "port": 443,
	                       "udp": true,
	                       "users": [
	                           {
	                               "level": 1,
	                               "alterId": 16,
	                               "security": "aes-128-gcm",
	                               "id": "66d88922-c880-4b5a-906e-ff50fc8c209f"
	                           }
	                       ]
	                   }
	               ]
	           },
	           "streamSettings": {
	               "network": "ws",
	               "security": "tls",
	               "sockopt": {
	                   "tcpFastOpen": true
	               },
	               "wsSettings": {
	                   "connectionReuse": true,
	                   "path": "/llc"
	               },
	               "tlsSettings": {
	                   "allowInsecure": true
	               }
	           },
	           "mux": {
	               "enabled": true,
	               "concurrency": 8
	           },
	           "protocol": "vmess",
	           "tag": "proxy"
	       },
	       {
	           "settings": {
	               "domainStrategy": "UseIP"
	           },
	           "streamSettings": {
	               "sockopt": {
	                   "tcpFastOpen": true
	               }
	           },
	           "protocol": "freedom",
	           "tag": "direct"
	       },
	       {
	           "protocol": "blackhole",
	           "tag": "block",
	           "settings": {}
	       }
	   ],
	   "routing": {
	       "strategy": "rules",
	       "settings": {
	           "domainStrategy": "IPIfNonMatch",
	           "rules": [
	               {
	                   "type": "field",
	                   "outboundTag": "direct",
	                   "ip": [
	                       "0.0.0.0/8",
	                       "10.0.0.0/8",
	                       "100.64.0.0/10",
	                       "127.0.0.0/8",
	                       "169.254.0.0/16",
	                       "172.16.0.0/12",
	                       "192.0.0.0/24",
	                       "192.0.2.0/24",
	                       "192.168.0.0/16",
	                       "198.18.0.0/15",
	                       "198.51.100.0/24",
	                       "203.0.113.0/24",
	                       "::1/128",
	                       "fc00::/7",
	                       "fe80::/10"
	                   ]
	               }
	           ]
	       }
	   }
	}
		`

	pbBytes, err := magic.GetPbConfigBytes([]byte(cfgStr))

	var sig = make(chan int)

	err = StartInstance(pbBytes)

	common.Must(err)

	<-sig

}
