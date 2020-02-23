FROM golang as build-stage
COPY go.mod /
COPY go.sum /
COPY main.go /
RUN cd / && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o publisher

FROM alpine
COPY --from=build-stage /publisher /
EXPOSE 8080
CMD ["/publisher"]
