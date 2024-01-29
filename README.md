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