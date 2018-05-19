 package model

// import (
//     “net”
//     “net/http”
//     “time”
// )

// var (
//     HttpClient *http.Client
// )

// // init HTTPClient  默认开启长链接(http 1.1之后) 开启http keepalive功能，也即是否重用连接，
// func init() {
//     HttpClient = createHTTPClient()
// }

// const (
//     MaxIdleConns        int = 100 // 连接池对所有host的最大链接数量，host也即dest-ip
//     MaxIdleConnsPerHost int = 100 // 连接池对每个host的最大链接数量
//     IdleConnTimeout     int = 90  // 空闲timeout设置，也即socket在该时间内没有交互则自动关闭连接（注意：该timeout起点是从每次空闲开始计时，若有交互则重置为0）,该参数通常设置为分钟级别
//     RequestTimeout      int = 30  // 请求以及连接的超时时间
// )

// // createHTTPClient for connection re-use
// func createHTTPClient() *http.Client {
//     client := &http.Client{
//         Transport: &http.Transport{
//             Proxy: http.ProxyFromEnvironment,
//             DialContext: (&net.Dialer{
//                 Timeout:   time.Duration(RequestTimeout) * time.Second,
//                 KeepAlive: time.Duration(RequestTimeout) * time.Second,
//             }).DialContext,

//             MaxIdleConns:        MaxIdleConns,
//             MaxIdleConnsPerHost: MaxIdleConnsPerHost,
//             IdleConnTimeout:     time.Duration(IdleConnTimeout) * time.Second,
//         },

//         Timeout: time.Duration(RequestTimeout) * time.Second,
//     }

//     return client
// }