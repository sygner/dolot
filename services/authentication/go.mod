module dolott_authentication

go 1.23.3

require (
	github.com/lib/pq v1.10.9
	github.com/matthewhartstonge/argon2 v1.0.1
	github.com/redis/go-redis/v9 v9.6.1
	google.golang.org/grpc v1.66.2
	safir/libs/appconfigs v0.0.0-00010101000000-000000000000
	safir/libs/appstates v0.0.0-00010101000000-000000000000
	safir/libs/idgen v0.0.0-00010101000000-000000000000
)

replace safir/libs/appconfigs => ../../libs/appconfigs

replace safir/libs/appstates => ../../libs/appstates

replace safir/libs/idgen => ../../libs/idgen

require google.golang.org/protobuf v1.34.2

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/rs/xid v1.6.0 // indirect
	golang.org/x/crypto v0.27.0 // indirect
	golang.org/x/net v0.29.0 // indirect
	golang.org/x/sys v0.25.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240903143218-8af14fe29dc1 // indirect
)
