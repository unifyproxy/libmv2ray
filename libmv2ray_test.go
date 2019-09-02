package libmv2ray

import "testing"

func TestStartV2ray(t *testing.T) {
	var a = V2Manager{}
	a.config = "{\n  \"log\": {\n    \"access\": \"/var/log/v2ray/access.log\",\n    \"error\": \"/var/log/v2ray/error.log\",\n    \"loglevel\": \"warning\"\n  },\n  \"inbound\": {\n    \"port\": 12550,\n    \"protocol\": \"vmess\",\n    \"settings\": {\n      \"clients\": [\n        {\n          \"id\": \"da9a735d-0dfe-e46d-3059-7e74f1ed6334\",\n          \"level\": 1,\n          \"alterId\": 1000\n        }\n      ]\n    }\n  },\n  \"outbound\": {\n    \"protocol\": \"freedom\",\n    \"settings\": {}\n  },\n  \"inboundDetour\": [],\n  \"outboundDetour\": [\n    {\n      \"protocol\": \"blackhole\",\n      \"settings\": {},\n      \"tag\": \"blocked\"\n    }\n  ],\n  \"routing\": {\n    \"strategy\": \"rules\",\n    \"settings\": {\n      \"rules\": [\n        {\n          \"type\": \"field\",\n          \"ip\": [\n            \"0.0.0.0/8\",\n            \"10.0.0.0/8\",\n            \"100.64.0.0/10\",\n            \"127.0.0.0/8\",\n            \"169.254.0.0/16\",\n            \"172.16.0.0/12\",\n            \"192.0.0.0/24\",\n            \"192.0.2.0/24\",\n            \"192.168.0.0/16\",\n            \"198.18.0.0/15\",\n            \"198.51.100.0/24\",\n            \"203.0.113.0/24\",\n            \"::1/128\",\n            \"fc00::/7\",\n            \"fe80::/10\"\n          ],\n          \"outboundTag\": \"blocked\"\n        }\n      ]\n    }\n  }\n}"
	a.StartV2ray()
}