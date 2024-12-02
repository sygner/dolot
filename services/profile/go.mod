module dolott_profile

go 1.23.3

require (
	github.com/lib/pq v1.10.9
	google.golang.org/grpc v1.68.0
	google.golang.org/protobuf v1.35.2
	safir/libs/appconfigs v0.0.0-00010101000000-000000000000
	safir/libs/appstates v0.0.0-00010101000000-000000000000
	safir/libs/idgen v0.0.0-00010101000000-000000000000
)

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/rs/xid v1.6.0 // indirect
	golang.org/x/net v0.31.0 // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241118233622-e639e219e697 // indirect
)

replace safir/libs/appconfigs => ../../libs/appconfigs

replace safir/libs/appstates => ../../libs/appstates

replace safir/libs/idgen => ../../libs/idgen
