# Prometheus ksyun sd

腾讯云KEC的服务发现.

类似[ec2_sd_config](https://prometheus.io/docs/prometheus/latest/configuration/configuration/#ec2_sd_config), prometheus对没有原生支持的sd通过`file_sd`实现, 本repo采用此思路.

## txyun api
腾讯云api https://github.com/TencentCloud/tencentcloud-sdk-go
 
## 用法 go build -o tencent-sd

```
Usage of ./tencent-sd:
  -config string
        config file path (default "sd-config.yaml")
  -output string
        output file path (default "sd.yaml")
```

配置文件

```
ak: ""
sk: ""
region: ""
port: 9100
interval: 1m
```



## 过滤器

支持的过滤器见https://cloud.tencent.com/document/product/213/15728



