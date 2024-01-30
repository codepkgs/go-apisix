# 说明

## APISIX 版本

> 基于 apisix api v3 版本。

## 安装

```bash
go get -u github.com/codepkgs/go-apisix
```

# 使用

* 创建apisix client

    ```go
    client, err := apisix.NewClient("http://127.0.0.1:9180", "your-admin-token")
    if err != nil {
        fmt.Println(err)
    }
    ```

* 分页查询

    > 如果没有指定分页，则默认获取所有该类型的资源。

    ```go
    ssls, err := client.GetSSLs(apisix.WithPageNumber(1), apisix.WithPageSize(10))
    if err != nil {
        fmt.Println(err)
    }
    
    for _, ssl := range ssls {
        fmt.Printf("%#v\n", ssl)
    }
    ```

* 根据资源名或标签Key过滤资源
    
    ```go
    ups, err := client.GetUpstreams(apisix.WithName("apisix"), apisix.WithLabelKey("env"))
    if err != nil {
        fmt.Println(err)
    }
    
    for _, up := range ups {
        fmt.Printf("%#v\n", up)
    }
    ```

# 路由管理

* 查询所有路由（如果路由比较多的话建议使用分页查询）

  ```go
  routes, err := client.GetRoutes()
  if err != nil {
      fmt.Println(err)
  } else {
      fmt.Println(len(routes))
  
      for _, route := range routes {
          fmt.Printf("%#v\n", route)
      }
  }
  ```

* 分页查询所有路由

  ```go
  pageNumber := 1
  pageSize := 10
  
  for {
      routes, err := client.GetRoutes(apisix.WithPageNumber(pageNumber), apisix.WithPageSize(pageSize))
      if err != nil {
          fmt.Println(err)
          break
      }
  
      if len(routes) != 0 {
          for _, route := range routes {
              fmt.Printf("%#v\n", route)
          }
          pageNumber++
      } else {
          break
      }
  }
  ```

* 创建路由

  ```go
  r, err := client.CreateRoute(
      "测试路由",
      apisix.CreateOrModifyRouteWithDesc("测试路由"),
      apisix.CreateOrModifyRouteWithLabels(map[string]string{"env": "test", "app": "apisix_dashboard"}),
      apisix.CreateOrModifyRouteWithHosts([]string{"d.pgops.com", "dash.pgops.com"}),
      apisix.CreateOrModifyRouteWithUri("/*"),
      apisix.CreateOrModifyRouteWithUpstream(&apisix.Upstream{
          Name:   "apisix-dashboard-upstream",
          Desc:   "apisix dashboard upstream",
          Labels: map[string]string{"env": "test"},
          Type:   apisix.RoundRobin,
          Scheme: apisix.HTTP,
          Nodes: client.ConvertUpstreamNodeStructToMap(
              []*apisix.UpstreamNode{
                  {Host: "172.16.158.29", Port: 9100, Weight: 1},
                  {Host: "172.16.158.30", Port: 9100, Weight: 1},
              },
          ),
      }),
      apisix.CreateOrModifyRouteWithStatus(apisix.RouteEnable),
      apisix.CreateOrModifyRouteWithPlugins(map[string]any{"redirect": map[string]any{"http_to_https": true}}),
  )
  
  if err != nil {
      fmt.Println(err)
  } else {
      fmt.Printf("%#v", r)
  }
  ```

* 删除路由

  ```go
  resp, err := client.DeleteRoute("498406734138704757")
  fmt.Println(err)
  fmt.Printf("%#v\n", resp)
  ```
  
* 修改路由

  > 明确出现的属性会被更新，未出现的属性则不会发生变化。

  ```go
  r, err := client.ModifyRoute("00000000000000000449",
      apisix.CreateOrModifyRouteWithName("测试修改路由"),
      apisix.CreateOrModifyRouteWithDesc("测试修改路由"),
      apisix.CreateOrModifyRouteWithUpstream(&apisix.Upstream{
          Name:   "apisix-dashboard-upstream",
          Desc:   "apisix dashboard upstream",
          Labels: map[string]string{"env": "test"},
          Type:   apisix.RoundRobin,
          Scheme: apisix.HTTP,
          Nodes: client.ConvertUpstreamNodeStructToMap(
              []*apisix.UpstreamNode{
                  {Host: "172.16.158.29", Port: 9000, Weight: 1},
              },
          ),
      }),
      apisix.CreateOrModifyRouteWithStatus(apisix.RouteEnable),
      apisix.CreateOrModifyRouteWithPlugins(nil),
  )
  
  if err != nil {
      fmt.Println(err)
  } else {
      fmt.Printf("%#v", r)
  }
  ```

# SSL证书管理

* 查询所有证书

  ```go
  ssls, err := client.GetSSLs()
  if err != nil {
      fmt.Println(err)
  } else {
      for _, ssl := range ssls {
          fmt.Printf("%#v\n", ssl)
      }
  }
  ```

* 查询指定ID的证书

  ```go
  ssl, _ := client.GetSSL("00000000000000000413")
  fmt.Printf("%#v", ssl)
  ```

* 删除证书

  ```go
  resp, err := client.DeleteSSL("497242557483320181")
  fmt.Println(resp, err)
  ```

* 创建证书

  ```go
  f, _ := os.OpenFile("/Users/codepkgs/Desktop/*.pgops.com.key", os.O_RDONLY, 0644)
  key, _ := io.ReadAll(f)
  defer f.Close()
  
  f1, _ := os.OpenFile("/Users/codepkgs/Desktop/*.pgops.com.crt", os.O_RDONLY, 0644)
  cert, _ := io.ReadAll(f1)
  defer f1.Close()
  
  sr, err := client.CreateSSL(
      key, cert, []string{"*.pgops.com", "pgops.com"},
      apisix.CreateSSLWithStatus(apisix.SSLEnable),
      apisix.CreateSSLWithSSLProtocols([]apisix.SSLProtocol{apisix.TLSv11, apisix.TLSv12, apisix.TLSv13}),
  )
  fmt.Println(err)
  fmt.Printf("%#v", sr)
  ```
  
# Upstream管理

* 查询所有的Upstream

  ```go
  ups, err := client.GetUpstreams()
  fmt.Println(err)
  for _, up := range ups {
      fmt.Printf("%#v\n", up)
  }
  ```

* 查询指定ID的Upstream

  ```go
  up, err := client.GetUpstream("00000000000000000306")
  fmt.Println(err)
  fmt.Printf("%#v\n", up)
  ```
  
* 删除指定ID的Upstream
  
  ```go
  up, err := client.DeleteUpstream("00000000000000000306")
  fmt.Println(err)
  fmt.Printf("%#v\n", up)
  ```
  
* 创建Upstream

  ```go
  upi, err := client.CreateUpstream("test2",
      apisix.CreateOrModifyUpstreamWithLoadBalancerType(apisix.RoundRobin),
      apisix.CreateOrModifyUpstreamWithNodes([]apisix.UpstreamNode{
          {Host: "172.16.158.29", Port: 9100, Weight: 1},
          {Host: "172.16.158.29", Port: 9101, Weight: 2},
      }),
      apisix.CreateOrModifyUpstreamWithDesc("测试apisix"),
      apisix.CreateOrModifyUpstreamWithKeepalivePool(apisix.KeepalivePool{
          Size:        64,
          Requests:    1000,
          IdleTimeout: 60,
      }),
  )
  fmt.Println(err)
  fmt.Printf("%#v", upi)
  ```
  
* 修改Upstream

  > 修改Upstream和创建Upstream类似。

