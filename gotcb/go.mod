module github.com/gotcb/tencent

go 1.12

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d // indirect
	github.com/garyburd/redigo v1.6.0
	github.com/ghodss/yaml v1.0.1-0.20190212211648-25d852aebe32 // indirect
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.2
	github.com/gorilla/websocket v1.4.1
	github.com/lib/pq v1.8.0
	github.com/micro/go-micro v1.13.2
	github.com/minio/minio-go v6.0.14+incompatible
	github.com/prometheus/common v0.6.0
	github.com/shijuvar/gokit v1.5.0
	github.com/tealeg/xlsx v1.0.5
	go.etcd.io/etcd v3.3.18+incompatible
	go.uber.org/zap v1.11.0
	golang.org/x/net v0.0.0-20191011234655-491137f69257
	google.golang.org/appengine v1.6.1 // indirect
	google.golang.org/grpc v1.24.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
)

replace github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.4

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
