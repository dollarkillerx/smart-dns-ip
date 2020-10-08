# smart-dns-ip
smart dns ip server (基于ip2region的IP库Golang实现了Restful API &amp;&amp; GRPC)

config: 
```yaml
RestfulAddr: "0.0.0.0:8086"
GRPCAddr: "0.0.0.0:8087"
DBPath: "./ip2region.db" # https://github.com/lionsoul2014/ip2region/blob/master/data/ip2region.db
```

Restful API:
``` 

```

GRPC API:
```proto
message IPSearchRequest {
  string ip = 1;
}

message IPSearchResponse {
  int64 cityId = 1;
  string country = 2;
  string region = 3;
  string city = 4;
  string isp = 5;
}

service IPSearch {
  rpc IPSearch(IPSearchRequest) returns (IPSearchResponse);
}
```