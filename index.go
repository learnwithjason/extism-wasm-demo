package main

import (
	"time"

	pdk "github.com/extism/go-pdk"
)

type JsonTime struct {
	Time   time.Time `json:"time"`
	Format string    `json:"format"`
}

//export format
func format() int32 {
	var t JsonTime
	err := pdk.InputJSON(&t)
	if err != nil {
		pdk.SetError(err)
		return 1
	}

	pdk.OutputString(t.Time.Format(t.Format))
	return 0
}

func main() {}
