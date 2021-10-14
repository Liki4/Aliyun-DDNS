# Aliyun-DDNS
Aliyun Dynamic DNS service server.

## Build
```shell
go build -trimpath -ldflags "-w -s" -o build/linux-amd64
```

## Config
copy `ddns.sample.toml` as `ddns.toml` and modify the toml file
```toml
[site]
name = "aliddns"
port = 8000
```

## Api
```
# post /updateDomainRecord
{
    "accessKeyId": "",
    "accessSecret": "",
    "domain": "example.com",
    "record": "www",
    "ip": "1.1.1.1"
}
```