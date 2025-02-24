package main

import (
	"database/sql"
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
	_ "github.com/go-sql-driver/mysql"
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

	// 获取消息
	msgs, err := ch.Consume(
		q.Name, // 队列名称
		"",     // 消费者标签
		true,   // 是否自动确认
		false,  // 是否独占
		false,  // 是否等待
		false,  // 是否阻塞
		nil,    // 额外参数
	)
	failOnError(err, "无法接收消息")

	// 处理消息
	// var forever chan struct{}

	// go func() {
	//   for d := range msgs {
	//     log.Printf("收到消息: %s", d.Body)
	//   }
	// }()

	// log.Printf("等待消息, 按 CTRL+C 退出")
	// <-forever

	// 连接 MySQL
	dsn := "root:root@tcp(127.0.0.1:3306)/test"
	db, err := sql.Open("mysql", dsn)
	failOnError(err, "无法连接到数据库")
	defer db.Close()

	log.Printf("等待消息, 按 CTRL+C 退出")

	// 处理消息并插入到数据库
	for msg := range msgs {
		// 打印收到的消息
		log.Printf("收到消息: %s", msg.Body)

		// 插入到数据库
		_, err := db.Exec("INSERT INTO rabbitmq (message) VALUES (?)", string(msg.Body))
		failOnError(err, "无法插入消息到数据库")
		log.Println("消息插入成功")
	}
}