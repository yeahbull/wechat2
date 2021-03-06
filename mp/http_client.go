// @description wechat2 是腾讯微信公众平台 api 的 golang 语言封装
// @link        https://github.com/philsong/wechat2 for the canonical source repository
// @license     https://github.com/philsong/wechat2/blob/master/LICENSE
// @authors     chanxuehong(chanxuehong@gmail.com)

package mp

import (
	"net"
	"net/http"
	"time"
)

// net/http:
//
// type Client struct {
//     // Transport specifies the mechanism by which individual
//     // HTTP requests are made.
//     // If nil, DefaultTransport is used.
//     Transport RoundTripper
//
//     ...
// }
//
// var DefaultClient = &Client{}
//
// var DefaultTransport RoundTripper = &Transport{
//     Proxy: ProxyFromEnvironment,
//     Dial: (&net.Dialer{
//         Timeout:   30 * time.Second,
//         KeepAlive: 30 * time.Second,
//     }).Dial,
//     TLSHandshakeTimeout: 10 * time.Second,
// }

// 一般请求的 http.Client
var TextHttpClient = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	},
	Timeout: 15 * time.Second,
}

// 多媒体上传下载请求的 http.Client
var MediaHttpClient = &http.Client{
	Transport: &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	},
	Timeout: 300 * time.Second, // 因为目前微信支持最大的文件是 10MB, 请求超时时间保守设置为 300 秒
}
