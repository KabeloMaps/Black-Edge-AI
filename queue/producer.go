package queue

import "encoding/json"

func Publish(channel string, payload interface{}) {
	data, _ := json.Marshal(payload)
	RDB.Publish(Ctx, channel, data)
}