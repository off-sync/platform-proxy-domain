# Off-Sync.com Platform Proxy Domain

[![CircleCI](https://circleci.com/gh/off-sync/platform-proxy-domain.svg?style=svg)](https://circleci.com/gh/off-sync/platform-proxy-domain)

The domain project contains the models for the Off-Sync.com Platform Proxy.

This project distinguishes the following aggregate roots:
* Services;
* Frontends.

## Services

A [Service](https://godoc.org/github.com/off-sync/platform-proxy-domain/services#Service) defines a functionality provided by one or more backend servers.

## Frontends

A [Frontend](https://godoc.org/github.com/off-sync/platform-proxy-domain/frontends#Frontend) defines how a Service is exposed through the Platform Proxy.

A Frontend can contain an optional [Certificate](https://godoc.org/github.com/off-sync/platform-proxy-domain/frontends#Certificate).
