package protocol

import (
	"errors"
	"fmt"

	"github.com/bytedance/sonic"
	"github.com/gucooing/BaPs/pkg/mx"
	"github.com/gucooing/BaPs/protocol/cmd"
	"github.com/gucooing/BaPs/protocol/proto"
)

type NetworkProtocolResponse struct {
	Packet   string `json:"packet"`   // 包数据
	Protocol string `json:"protocol"` // 协议名称
}

// UnmarshalRequest 解码req数据包
func UnmarshalRequest(b []byte) (proto.Message, *proto.BasePacket, error) {
	base := new(proto.BasePacket)
	err := sonic.Unmarshal(b, base)
	if err != nil {
		return nil, nil, err
	}
	packet := cmd.Get().GetRequestPacketByCmdId(base.GetProtocol())
	if packet == nil {
		return nil, nil, errors.New(fmt.Sprintf("request unknown cmd id: %v", base.GetProtocol()))
	}
	err = sonic.Unmarshal(b, packet)
	if err != nil {
		return nil, nil, err
	}

	return packet, base, nil
}

// MarshalResponse 编码rsp数据包
func MarshalResponse(m proto.Message) (*NetworkProtocolResponse, error) {
	if m == nil {
		return nil, errors.New("message nil")
	}
	jsonData, err := sonic.MarshalString(m)
	if err != nil {
		return nil, err
	}
	v := &NetworkProtocolResponse{
		Packet:   jsonData,
		Protocol: mx.Protocol(m.GetProtocol()).String(),
	}

	return v, nil
}
