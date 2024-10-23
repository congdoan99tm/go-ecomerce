# FROM golang:alpine


# WORKDIR /build/

# COPY . .

# RUN go build -o nameDemo .

# WORKDIR /dist

# RUN cp /build/nameDemo .

# EXPOSE 8009

# CMD [ "/dist/nameDemo" ]

#FROM golang:alpine AS builder
#
#WORKDIR /build/
#
#COPY ../demo-docker/demodocker .
#
#RUN go build -o nameDemo1 .
#
#FROM scratch
#
#COPY --from=builder /build/nameDemo1 /
#
#ENTRYPOINT [ "/nameDemo1" ]

FROM golang:alpine AS builder

WORKDIR /build/

COPY . .

RUN go mod download

RUN go build -o crm.shopdev.com ./cmd/server

FROM scratch

COPY ./config /config

COPY --from=builder /build/crm.shopdev.com /

ENTRYPOINT ["/crm.shopdev/com", "config/local.yaml"]