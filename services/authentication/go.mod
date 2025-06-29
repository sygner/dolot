module dolott_authentication

go 1.24.2

require (
	github.com/lib/pq v1.10.9
	github.com/matthewhartstonge/argon2 v1.2.0
	github.com/redis/go-redis/v9 v9.7.1
	google.golang.org/grpc v1.71.0
)

replace neo/libs/appconfigs => ../../libs/appconfigs

replace neo/libs/appstates => ../../libs/appstates

replace neo/libs/idgen => ../../libs/idgen

require (
	google.golang.org/protobuf v1.36.5
	neo/libs/appconfigs v0.0.0-00010101000000-000000000000
	neo/libs/appstates v0.0.0-00010101000000-000000000000
	neo/libs/idgen v0.0.0-00010101000000-000000000000
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/rs/xid v1.5.0 // indirect
	golang.org/x/crypto v0.36.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250303144028-a0af3efb3deb // indirect
)
