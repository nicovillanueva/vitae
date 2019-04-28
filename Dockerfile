FROM golang:1.12 as builder
WORKDIR /src/
ADD . /src/
RUN go get && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api-vitae ./cli/start-server

FROM alpine as runtime
RUN apk --no-cache add ca-certificates
COPY --from=0 /src/api-vitae /bin/
COPY ./cv-data.yaml /opt/
COPY ./CV-NicolasVillanueva.pdf /opt/
ENTRYPOINT ["api-vitae"]
CMD ["-c", "/opt/cv-data.yaml", "-p", "/opt/CV-NicolasVillanueva.pdf"]