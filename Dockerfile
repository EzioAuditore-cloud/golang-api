FROM golang:1.21.9
WORKDIR /src
COPY project .
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN go build -o chat /src/test/main.go

FROM ubuntu
COPY /project/config/globleConf.yaml .
COPY /project/database/config/DB.yaml .
COPY --from=0 /src/chat .
EXPOSE 8080
CMD [ "./chat" ]

