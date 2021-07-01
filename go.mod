module github.com/mikejoh/coinbase-go

go 1.15

require (
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.1.3
	k8s.io/release v0.9.0 // indirect
	sigs.k8s.io/release-utils v0.2.1
)

replace github.com/mikejoh/coinbase-go/cmd/cb => /home/mikael/Repos/personal/golang-playground/coinbase-go/cmd/cb
