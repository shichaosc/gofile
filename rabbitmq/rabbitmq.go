package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	_ "github.com/streadway/amqp"
	"log"
)

type Producer struct {
	Exchange string //交换机
	Qname    string //路由name
	logger   *log.Logger
	//Connection    *Connection
	done          chan bool              // 如果主动close，会接受数据
	notifyClose   chan *amqp.Error       // 如果异常关闭，会接受数据
	notifyConfirm chan amqp.Confirmation // 消息发送成功确认，会接受到数据
	isConnected   bool
}

type Connection struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s:%s", msg, err)
		panic(fmt.Sprintf("%s:%s", msg, err))
	}
}

func ConnMq() *Connection {
	fmt.Println("创建rabbitmq连接***********************************")
	connection, err := amqp.Dial("amqp://books:book123@139.9.92.116:5672/")
	//conn.connection = new(amqp.Connection)
	//conn.connection = connection
	if err != nil {
		log.Fatalf("Can't connect the rabbitmq server: %s", err)
	}
	ch, err := connection.Channel()
	//ch, err := conn.connection.Channel()
	if err != nil {
		log.Fatalf("Can't connect the Channel : %s", err)
	}
	conn := Connection{
		connection: connection,
		channel:    ch,
	}
	return &conn
}

var conn = ConnMq()

func (p *Producer) Publish(msg string) {

	//p.Connection.channel.NotifyClose(p.notifyClose)
	//p.Connection.channel.NotifyPublish(p.notifyConfirm)
	err := conn.channel.Publish(
		p.Exchange, //交换机
		p.Qname,    //routing_key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	if err != nil {
		log.Fatalf("publish msg fail: err=%s,msg=%s", err, msg)
	}
}

func Receive(queueName string, processUrl func(url string)) {
	if conn.channel == nil {
		conn = ConnMq()
	}

	msgs, err := conn.channel.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnErr(err, "")

	forever := make(chan bool)

	go func() {
		//fmt.Println(*msgs)
		for d := range msgs {
			fmt.Printf(" -----")
			s := string(d.Body)
			processUrl(s)
		}
	}()

	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C\n")
	<-forever
}
