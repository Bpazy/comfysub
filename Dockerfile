FROM golang:latest AS development
RUN git clone --progress --verbose --depth=1 https://github.com/Bpazy/comfysub /comfysub
WORKDIR /comfysub
RUN go env && CGO_ENABLED=0 go build ./cmd/comfysub

FROM alpine:latest AS production
ENV PORT ""
COPY --from=development /comfysub/comfysub /comfysub/comfysub
WORKDIR /comfysub
ENTRYPOINT ./comfysub -port $PORT
