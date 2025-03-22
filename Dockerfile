FROM --platform=$BUILDPLATFORM golang:1.23 AS build
ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

WORKDIR /go/src/app
COPY . .

RUN go mod download
RUN go vet -v
RUN go test -v

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /go/bin/app

FROM gcr.io/distroless/static-debian11
COPY --from=build /go/src/app /
COPY --from=build /go/bin/app /
CMD ["/app"]