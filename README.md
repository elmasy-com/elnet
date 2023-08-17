# Elnet

[![Go Reference](https://pkg.go.dev/badge/github.com/elmasy-com/elnet.svg)](https://pkg.go.dev/github.com/elmasy-com/elnet)
[![Go Report Card](https://goreportcard.com/badge/github.com/elmasy-com/elnet)](https://goreportcard.com/report/github.com/elmasy-com/elnet)

Get:
```bash
go get github.com/elmasy-com/elnet@latest
```

Or get the latest commit (if Go module proxy is not updated)"
```bash
go get "github.com/elmasy-com/elnet@$(curl -s 'https://api.github.com/repos/elmasy-com/elnet/commits' | jq -r '.[0].sha')"
```

## dns

[![Go Reference](https://pkg.go.dev/badge/github.com/elmasy-com/elnet/dns.svg)](https://pkg.go.dev/github.com/elmasy-com/elnet/dns)

DNS queries and helper function for domain names.

Based on [miekg's dns module](https://github.com/miekg/dns).