package model

import "time"

type Response struct {
	Status    uint
	RequestId uint
	Size      int64
	Time      time.Time
}
