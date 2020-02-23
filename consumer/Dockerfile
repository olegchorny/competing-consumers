FROM golang as build-stage
COPY go.mod /
COPY go.sum /
COPY main.go /
RUN cd / && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o consumer

FROM alpine
COPY --from=build-stage /consumer /
EXPOSE 8080
CMD ["/consumer"]
