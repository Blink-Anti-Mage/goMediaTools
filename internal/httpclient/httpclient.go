package httpclient

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"
)

var (
	defaultHttpClient *http.Client
)

func init() {
	defaultHttpClient = newClient(true)
}

const (
	MaxIdleConnNum     int = 1000
	MaxIdleConnPerHost int = 1000
	IdleConnTimeout    int = 90
	Timeout            int = 15
)

func newClient(verifyFlag bool) *http.Client {
	if verifyFlag != true {
		return &http.Client{
			Transport: &http.Transport{
				Proxy:           http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				DialContext: (&net.Dialer{
					Timeout:   time.Duration(Timeout) * time.Second,
					KeepAlive: time.Duration(Timeout) * time.Second,
				}).DialContext,

				MaxIdleConns:        MaxIdleConnNum,
				MaxIdleConnsPerHost: MaxIdleConnPerHost,
				IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
			},

			Timeout: time.Duration(Timeout) * time.Second,
		}
	} else {
		return &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   time.Duration(Timeout) * time.Second,
					KeepAlive: time.Duration(Timeout) * time.Second,
				}).DialContext,

				MaxIdleConns:        MaxIdleConnNum,
				MaxIdleConnsPerHost: MaxIdleConnPerHost,
				IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
			},

			Timeout: time.Duration(Timeout) * time.Second,
		}
	}
}

func GetClient(verifyFlag bool, timeout time.Duration) *http.Client {
	return createNotMultiplexHTTPClient(verifyFlag, timeout)
}

func createNotMultiplexHTTPClient(verifyFlag bool, timeout time.Duration) *http.Client {
	if timeout <= 0 {
		timeout = time.Duration(Timeout) * time.Second
	}

	if !verifyFlag {
		return &http.Client{
			Transport: &http.Transport{
				Proxy:           http.ProxyFromEnvironment,
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				DialContext: (&net.Dialer{
					Timeout:   timeout,
					KeepAlive: time.Duration(Timeout) * time.Second,
				}).DialContext,
				IdleConnTimeout: time.Duration(IdleConnTimeout) * time.Second,
			},

			Timeout: timeout,
		}
	} else {
		return &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   timeout,
					KeepAlive: time.Duration(Timeout) * time.Second,
				}).DialContext,

				IdleConnTimeout: time.Duration(IdleConnTimeout) * time.Second,
			},

			Timeout: timeout,
		}
	}
}

func GetShortByProxyClient(timeout time.Duration, proxyURL string) *http.Client {
	p, err := url.Parse(proxyURL)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	httpProxy := &http.Client{
		Transport: &http.Transport{
			Proxy:             http.ProxyURL(p),
			DisableKeepAlives: true,
		},
		Timeout: timeout,
	}

	return httpProxy
}
