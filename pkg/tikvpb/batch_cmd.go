package tikvpb

import (
	"encoding/binary"

	"github.com/pingcap/kvproto/pkg/coprocessor"
	"github.com/pingcap/kvproto/pkg/kvrpcpb"
	"github.com/pingcap/kvproto/pkg/sharedbytes"
)

func TryDecodeBatchCommandsResp(data []byte) (any, bool) {
	size := len(data)
	if size == 0 {
		return nil, false
	}
	wire, idx := binary.Uvarint(data)
	if idx <= 0 {
		return nil, false
	}
	wireType := int32(wire & 0x7)
	if wireType != 2 {
		return nil, false
	}
	fieldNum := int32(wire >> 3)
	switch fieldNum {
	case 22:
		msglen, off := binary.Uvarint(data[idx:])
		begin := idx + off
		end := begin + int(msglen)
		if begin <= 0 || size < end {
			return nil, false
		}
		resp := coprocessor.Response{
			ExecDetailsV2: &kvrpcpb.ExecDetailsV2{
				TimeDetailV2: &kvrpcpb.TimeDetailV2{},
				ScanDetailV2: &kvrpcpb.ScanDetailV2{},
			},
			BatchResponses: make([]sharedbytes.SharedBytes, 0, 4),
		}
		if resp.Unmarshal(data[begin:end]) == nil {
			return &resp, true
		}
	}
	return nil, false
}
