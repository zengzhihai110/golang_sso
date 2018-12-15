package zconst

import "github.com/satori/go.uuid"

var InitTrackId string

func init() {
	InitTrackId = uuid.NewV4().String()
}