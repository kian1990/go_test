package main

import (
	"context"
	"log"
	"time"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
  if err != nil {
    log.Panicf("%s: %s", msg, err)
  }
}

func main() {
	// 连接 RabbitMQ 服务
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "无法连接到 RabbitMQ")
	defer conn.Close()

	// 创建通道
	ch, err := conn.Channel()
	failOnError(err, "无法创建通道")
	defer ch.Close()

	// 声明队列
	q, err := ch.QueueDeclare(
		"hello", // 队列名称
		false,    // 是否持久化
		false,   // 是否自动删除
		false,   // 是否独占
		false,   // 是否等待
		nil,     // 额外参数
	)
	failOnError(err, "无法声明队列")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 发送消息
	body := "Hello, RabbitMQ!"
	err = ch.PublishWithContext(ctx,
		"",     // 默认交换机
		q.Name, // 队列名称
		false,  // 是否持久化
		false,  // 是否强制路由
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "无法发送消息")
	log.Printf("发送消息: %s", body)
}
