FROM golang:1.20-alpine as builder

ENV HOME /go/src/orb

WORKDIR "$HOME"
COPY src/ .

RUN go mod download
RUN go mod verify
RUN go build -o orb

FROM golang:1.20-alpine

ENV HOME /go/src/orb
RUN mkdir -p "$HOME"
WORKDIR "$HOME"

COPY src/conf/ conf/
COPY --from=builder "$HOME"/orb $HOME

EXPOSE 8080
CMD ["./orb"]