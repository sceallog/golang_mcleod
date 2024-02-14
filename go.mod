module vl-general-info

go 1.20

replace github.com/mono-revo/vl-backend-basement => ./.basement

require (
	github.com/confluentinc/confluent-kafka-go v1.9.2
	github.com/confluentinc/confluent-kafka-go/v2 v2.3.0
	github.com/mono-revo/vl-backend-basement v0.0.0-20240117123111-f13b25184819
	github.com/redis/go-redis/v9 v9.4.0
	go.uber.org/fx v1.20.1
	go.uber.org/zap v1.26.0
	google.golang.org/protobuf v1.32.0
	gorm.io/gorm v1.25.5
)

require (
	github.com/aws/aws-sdk-go v1.49.21 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/elastic/elastic-transport-go/v8 v8.3.0 // indirect
	github.com/elastic/go-elasticsearch/v8 v8.11.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/iancoleman/orderedmap v0.3.0 // indirect
	github.com/jhump/protoreflect v1.14.1 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/moby/patternmatcher v0.6.0 // indirect
	github.com/moby/sys/sequential v0.5.0 // indirect
	github.com/sethvargo/go-retry v0.2.4 // indirect
	github.com/stretchr/testify v1.8.3 // indirect
	google.golang.org/genproto v0.0.0-20231030173426-d783a09b4405 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231106174013-bbf56f31fb17 // indirect
	gorm.io/driver/mysql v1.5.2 // indirect
)

require (
	github.com/Masterminds/log-go v1.0.0 // indirect
	github.com/antlr/antlr4/runtime/Go/antlr/v4 v4.0.0-20230321174746-8dcc6526cfb1 // indirect
	github.com/caarlos0/env/v9 v9.0.0
	github.com/go-sql-driver/mysql v1.7.1
	github.com/google/uuid v1.5.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/pressly/goose/v3 v3.16.0
	github.com/thmeitz/ksqldb-go v0.1.0
	go.uber.org/dig v1.17.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20240112132812-db7319d0e0e3
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/sync v0.6.0
	golang.org/x/sys v0.14.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/grpc v1.60.1
)
