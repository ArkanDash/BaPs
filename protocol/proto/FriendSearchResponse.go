// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type FriendSearchResponse struct{
	message ProtoMessage
	ResponsePacket

    SearchResult []*FriendDB
}

func (x *FriendSearchResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *FriendSearchResponse) ProtoReflect() Message {
	return x
}

func (x *FriendSearchResponse) GetProtocol() int32 {
	return mx.Protocol_Friend_Search
}

func (x *FriendSearchResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

