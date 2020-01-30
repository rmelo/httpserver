package entity

import "time"

//Time is a struct representing time
type Time struct {
	Sleep int    `json:"sleep"`
	Time  string `json:"time"`
}

//NewTime creates an instance of Time.
func NewTime() Time {
	return Time{Time: time.Now().String()}
}
