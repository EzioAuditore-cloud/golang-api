# **基于GIN实现的IM系统** 

## **基于kafka的消息队列**

消息发送方作为生产者，如果接收方不在线，消息状态被标记为0，消费后写入MySql，已读消息则先被标记为1再消费
![image](https://github.com/EzioAuditore-cloud/golang-api/assets/62204263/19a90209-6060-4e81-9767-daab7fa93442)


## **配置**

IP、端口配置: ./config/globleConf.yaml

MySQL、Redis 配置: ./database/config/DB.yaml

log 配置: ./middleWare/logger/config/log.yaml

## **启动:**

go run main.go
