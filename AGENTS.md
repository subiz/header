# AGENTS.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## What this repo is

`github.com/subiz/header` is the shared **schema + contract library** for Subiz's microservices. It holds the protobuf definitions (gRPC service interfaces, message types, enums) that every service imports, plus hand-written Go helpers for the domain types those protos describe. Changing a `.proto` here changes the API surface of the whole platform, so treat edits as cross-service contract changes.

## Build & code generation

**Never hand-edit generated files.** The following are outputs and are overwritten on every build:
- `*.pb.go`, `*_grpc.pb.go` (from `protoc`)
- `locale.generated.go`, `locale.generated.proto` (from `lang.js`)

Regeneration runs inside a Docker image (protoc + protoc-gen-go + Node are pinned there, not on the host):

```sh
./build.sh                       # runs the protobuild container against the working dir
# rebuild the container after changing Dockerfile / lang.js / protoc-go.sh:
docker build --progress=plain -t live360vn/protobuild:3.4 . && ./build.sh
```

Inside the container, `protoc-go.sh` first runs `node ./lang.js` (generates the locale/currency data), then compiles every `.proto` (skipping `vendor`/`proto`/`node_modules`). Output lands under `github.com/subiz/header/...` and is copied back to the repo root.

To edit the API: change the `.proto`, run `./build.sh`, then commit both the `.proto` and the regenerated `*.pb.go`.

### proto → package layout
Each `.proto` sets `option go_package`. Root-level protos (`header.proto`, `service.proto`, `request.proto`, `type.proto`, `common.proto`) generate into package `header` at the repo root. Others generate into sibling packages: `account/`, `common/`, `payment/`, `fabikon/`, `googlekon/`, `shopee/`, `tiktokon/`, `zalokon/`. `service.proto` + `request.proto` define the gRPC services; `header.proto` is the large shared message catalog.

## Tests

Standard Go tests; most live in `all_test.go` (package `header`).

```sh
go test ./...                    # everything
go test .                        # root package only
go test -run TestNormPhone .     # single test
```

Note: `go vet` reports pre-existing warnings (proto structs contain a `sync.Mutex`, so copying a `Block`/message by value trips the lock-copy check; there is also known unreachable code). These are baseline noise — don't treat them as regressions unless your change adds new ones.

## Hand-written code (package `header`, repo root)

These are the files you actually edit for logic (as opposed to generated types):

- `header.go` — the core domain toolkit and by far the largest hand-written file: user **attribute** read/update logic (`UpdateAttribute`, `MakePrimaryValue`, `AttributeValue`), phone/email normalization (`NormPhone`, `VNPhone`, `EmailAddress`, `GetVietnamPhoneISP`), the **Block** rich-content model (`DeltaToBlock`, `BlockToText`, `BlockToHTML`, `CompileBlock`), order totals, and error conversion (`ToErr`/`FromErr`). Its `init()` loads permission scope expansion.
- `grpc.go` — gRPC plumbing shared by all services: `DialGrpc`, sharded routing (`WithShardRedirect`, `NewServerShardInterceptor`, `GetAccShard`, `GetHostShard`), context conversion between the proto `common.Context` and Go `context.Context` (`ToGrpcCtx`/`FromGrpcCtx`), panic recovery and error-stack interceptors, Prometheus + OTel wiring.
- `attrdef.go`, `userutil.go`, `string.go`, `business_hour.go`, `json.go`, `klock.go`, `gocql.go` — narrower helpers (attribute definitions, user utilities, string ops, business-hours math, JSON assignment, clock, gocql type marshaling).

### Sharding model
Requests route to a shard derived from account/host id. Client interceptors (`WithShardRedirect`) transparently re-dial the correct backend when a call lands on the wrong shard; server interceptors (`NewServerShardInterceptor`, `NewServerShardInterceptor2`) enforce shard ownership. When touching gRPC dialing or routing, preserve this redirect-on-wrong-shard behavior.

## The JS files

`lang.js`, `perm.js`, `block.js` are not app code — they support generation and mirror Go logic:
- `lang.js` — the generator for `locale.generated.*` (currency codes, locale data via `@subiz/langmap`). Runs automatically in `protoc-go.sh`.
- `perm.js` / `perm.json` — permission scope definitions and expansion (`accessFeature`, `expandScope`). `perm.json` is the source of truth; `header.go`'s `init()`/`expandScope` consumes the same data on the Go side, so keep the two in sync.
- `block.js` — sample/reference fixtures for the Block content model.
