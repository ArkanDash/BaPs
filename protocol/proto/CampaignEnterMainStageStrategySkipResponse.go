// Code generated by gucooing DO NOT EDIT.

package proto

import (
	"encoding/json"

	"github.com/gucooing/BaPs/pkg/mx"
)

type CampaignEnterMainStageStrategySkipResponse struct{
	message ProtoMessage
	ResponsePacket

    ParcelResultDB *ParcelResultDB
}

func (x *CampaignEnterMainStageStrategySkipResponse) String() string {
	j, _ := json.Marshal(x)
	return string(j)
}

func (x *CampaignEnterMainStageStrategySkipResponse) ProtoReflect() Message {
	return x
}

func (x *CampaignEnterMainStageStrategySkipResponse) GetProtocol() int32 {
	return mx.Protocol_Campaign_EnterMainStageStrategySkip
}

func (x *CampaignEnterMainStageStrategySkipResponse) SetSessionKey(base *BasePacket) {
	if x == nil {
		return
	}
	x.BasePacket = base
}

