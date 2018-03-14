package main

// A simple example publisher app. Writes a hello message to a grpc pubsub server.
// You can run the server provided in this examples section.

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/infobloxopen/atlas-pubsub/examples/hello"
	pubsubgrpc "github.com/infobloxopen/atlas-pubsub/grpc"
	"google.golang.org/grpc"
)

var topic = flag.String("topic", hello.DefaultTopicName, "the topic to publish to")

func main() {
	flag.Parse()
	conn, err := grpc.Dial("pubsub.atlas.svc.cluster.local:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial to grpc server: %v", err)
	}
<<<<<<< HEAD
	for i := 0; ; i++ {
		p := pubsubgrpc.NewPublisher(*topic, conn)
		g := hello.Greetings[i%len(hello.Greetings)]

		msg := fmt.Sprintf("%s %s", g.Message, time.Now())
		md := map[string]string{"language": g.Language}
		log.Printf("printing %q %v", msg, md)
		p.Publish(context.Background(), []byte(msg), md)
=======
	p := pubsubgrpc.NewPublisher(*topic, conn)
	for {
		msg := fmt.Sprintf("Hello! %s", time.Now())
		log.Printf("printing %q", msg)
		err := p.Publish(context.Background(), []byte(msg))
		if err != nil {
			log.Printf("Failed to send a message")
			log.Println(err.Error())
		}
>>>>>>> Create a Makefile, dockerfiles for the server and the example publisher and subscriber. Created 3 kubernetes yaml files for deployment
		time.Sleep(time.Second)
	}
}
