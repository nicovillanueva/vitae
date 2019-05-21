FROM golang:1.12 as builder
WORKDIR /src/
ADD . /src/
RUN go get && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o vitae ./cli

FROM alpine as runtime
RUN apk --no-cache add ca-certificates
COPY --from=0 /src/vitae /bin/
COPY ./cv-data.yaml /opt/
COPY ./latex/nsv-cv.pdf /opt/
ENV CV_DATA_PATH /opt/cv-data.yaml
ENV CV_PDF_PATH /opt/nsv-cv.pdf
ENTRYPOINT ["vitae"]
CMD ["start"]
