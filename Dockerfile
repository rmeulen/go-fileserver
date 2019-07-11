FROM golang:1 as build
WORKDIR /go/src/
COPY main.go .
RUN CGO_ENABLED=0 go build -o /go/bin/fileserver .

FROM scratch
COPY --from=build /go/bin/fileserver .
ENV FILE_ROOT "/root"
ENV PORT 8080
EXPOSE ${PORT}
VOLUME ${FILE_ROOT}
ENTRYPOINT [ "./fileserver"]