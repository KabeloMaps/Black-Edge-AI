package queue

import (
	"fmt"
)

func Subscribe(channel string, handler func([]byte)) {
	sub := RDB.Subscribe(Ctx, channel)
	ch := sub.Channel()

	go func() {
		for msg := range ch {
			fmt.Println("📥 Event:", channel)
			handler([]byte(msg.Payload))
		}
	}()
}