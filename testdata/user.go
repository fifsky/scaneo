package testdata

import "time"

//go:generate scangen -s $GOFILE
type Users struct {
	Id        int64
	Username  string
	Password  string
	CreatedAt time.Time
}
