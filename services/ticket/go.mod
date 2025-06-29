module dolott_ticket

go 1.24.2

require (
	github.com/lib/pq v1.10.9
	google.golang.org/grpc v1.71.0
	neo/libs/appconfigs v0.0.0-00010101000000-000000000000
	neo/libs/appstates v0.0.0-00010101000000-000000000000
	neo/libs/idgen v0.0.0-00010101000000-000000000000
)

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/rs/xid v1.5.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/text v0.23.0 // indirect
)

require (
	golang.org/x/sys v0.31.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250311190419-81fb87f6b8bf // indirect
	google.golang.org/protobuf v1.36.5
)

replace neo/libs/appconfigs => ../../libs/appconfigs

replace neo/libs/appstates => ../../libs/appstates

replace neo/libs/idgen => ../../libs/idgen
