package websocket

import (
	"time"
)

const bitstampWsUrl = "wss://ws.bitstamp.net"
const wsTimeout = 60 * time.Second
const skipVerifyTLS = false

type wsClientConfig struct {
	domain        string
	timeout       time.Duration
	skipVerifyTLS bool
}

func defaultWsClientConfig() *wsClientConfig {
	return &wsClientConfig{
		domain:        bitstampWsUrl,
		timeout:       wsTimeout,
		skipVerifyTLS: skipVerifyTLS,
	}
}

type WsOption func(*wsClientConfig)

func WsUrl(domain string) WsOption {
	return func(config *wsClientConfig) {
		config.domain = domain
	}
}

func Timeout(timeout time.Duration) WsOption {
	return func(config *wsClientConfig) {
		config.timeout = timeout
	}
}

func SkipVerifyTLS(skipVerifyTLS bool) WsOption {
	return func(config *wsClientConfig) {
		config.skipVerifyTLS = skipVerifyTLS
	}
}
