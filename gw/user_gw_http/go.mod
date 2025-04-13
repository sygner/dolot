module dolott_user_gw_http

go 1.24.0

require (
	github.com/gofiber/fiber/v2 v2.52.6
	google.golang.org/grpc v1.71.0
	google.golang.org/protobuf v1.36.6
	safir/libs/appconfigs v0.0.0-00010101000000-000000000000
	safir/libs/appstates v0.0.0-00010101000000-000000000000
	safir/libs/idgen v0.0.0-00010101000000-000000000000
)

require (
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/klauspost/compress v1.18.0 // indirect
	github.com/mattn/go-colorable v0.1.14 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/rivo/uniseg v0.4.7 // indirect
	github.com/rs/xid v1.6.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.59.0 // indirect
	golang.org/x/net v0.37.0 // indirect
	golang.org/x/sys v0.31.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
)

replace safir/libs/appconfigs => ../../libs/appconfigs

replace safir/libs/appstates => ../../libs/appstates

replace safir/libs/idgen => ../../libs/idgen
