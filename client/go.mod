module github.com/ark-network/ark-cli

go 1.22.2

replace github.com/ark-network/ark/common => ../common

replace github.com/ark-network/ark => ../server

replace github.com/ark-network/ark-sdk => ../pkg/client-sdk

require (
	github.com/ark-network/ark v0.0.0-00010101000000-000000000000
	github.com/ark-network/ark-sdk v0.0.0-00010101000000-000000000000
	github.com/ark-network/ark/common v0.0.0
	github.com/btcsuite/btcd v0.24.2
	github.com/btcsuite/btcd/btcec/v2 v2.3.4
	github.com/btcsuite/btcd/chaincfg/chainhash v1.1.0
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.3.0
	github.com/urfave/cli/v2 v2.27.2
	golang.org/x/crypto v0.23.0
	golang.org/x/term v0.20.0
)

require (
	github.com/btcsuite/btcd/btcutil/psbt v1.1.9 // indirect
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f // indirect
	github.com/decred/dcrd/crypto/blake256 v1.0.1 // indirect
	github.com/vulpemventures/fastsha256 v0.0.0-20160815193821-637e65642941 // indirect
)

require (
	github.com/btcsuite/btcd/btcutil v1.1.5 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.4 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.20.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/vulpemventures/go-elements v0.5.3
	github.com/xrash/smetrics v0.0.0-20240312152122-5f08fbb34913 // indirect
	golang.org/x/net v0.25.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240528184218-531527333157 // indirect
	google.golang.org/grpc v1.65.0 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)
