package model

import "time"

type Request struct {
	ExecutionNanoSecond int64
	IP                  string
	Size                int64
	Method              string
	Path                string
	Route               string
	Time                time.Time
}
