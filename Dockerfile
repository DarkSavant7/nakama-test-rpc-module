FROM heroiclabs/nakama-pluginbuilder:3.21.1 AS go-builder

ENV GO111MODULE on
ENV CGO_ENABLED 1

WORKDIR /test-rpc-module

COPY go.mod .
COPY main.go .
COPY database/ database/
COPY rpc/ rpc/
COPY structures/ structures/
COPY vendor/ vendor/

RUN go build --trimpath --mod=vendor --buildmode=plugin -o ./test-rpc-module.so

FROM registry.heroiclabs.com/heroiclabs/nakama:3.21.1

COPY --from=go-builder /test-rpc-module/test-rpc-module.so /nakama/data/modules/
COPY local.yml /nakama/data/