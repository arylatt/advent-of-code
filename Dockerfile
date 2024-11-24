FROM golang:1.23 AS build

COPY . /go/src/github.com/arylatt/advent-of-code

WORKDIR /go/src/github.com/arylatt/advent-of-code/cmd/aoc-announce

RUN go build -o ../../out/aoc-announce .

FROM gcr.io/distroless/base-debian12:nonroot

COPY --from=build /go/src/github.com/arylatt/advent-of-code/out/aoc-announce /usr/local/bin/aoc-announce

ENTRYPOINT ["/usr/local/bin/aoc-announce"]

LABEL org.opencontainers.image.source=https://github.com/arylatt/advent-of-code
