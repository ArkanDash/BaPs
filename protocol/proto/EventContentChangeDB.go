// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"time"
)

type EventContentChangeDB struct {
	AccountId             int64
	EventContentId        int64
	UseAmount             int64
	ChangeCount           int64
	AccumulateChangeCount int64
	LastUpdateDate        time.Time
	ChangeFlag            bool
}
