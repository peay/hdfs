// Code generated by protoc-gen-go.
// source: ClientDatanodeProtocol.proto
// DO NOT EDIT!

package hadoop_hdfs

import proto "github.com/golang/protobuf/proto"
import math "math"
import hadoop_common "github.com/peay/hdfs/protocol/hadoop_common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// *
// block - block for which visible length is requested
type GetReplicaVisibleLengthRequestProto struct {
	Block            *ExtendedBlockProto `protobuf:"bytes,1,req,name=block" json:"block,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *GetReplicaVisibleLengthRequestProto) Reset()         { *m = GetReplicaVisibleLengthRequestProto{} }
func (m *GetReplicaVisibleLengthRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetReplicaVisibleLengthRequestProto) ProtoMessage()    {}

func (m *GetReplicaVisibleLengthRequestProto) GetBlock() *ExtendedBlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

// *
// length - visible length of the block
type GetReplicaVisibleLengthResponseProto struct {
	Length           *uint64 `protobuf:"varint,1,req,name=length" json:"length,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetReplicaVisibleLengthResponseProto) Reset()         { *m = GetReplicaVisibleLengthResponseProto{} }
func (m *GetReplicaVisibleLengthResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetReplicaVisibleLengthResponseProto) ProtoMessage()    {}

func (m *GetReplicaVisibleLengthResponseProto) GetLength() uint64 {
	if m != nil && m.Length != nil {
		return *m.Length
	}
	return 0
}

// *
// void request
type RefreshNamenodesRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RefreshNamenodesRequestProto) Reset()         { *m = RefreshNamenodesRequestProto{} }
func (m *RefreshNamenodesRequestProto) String() string { return proto.CompactTextString(m) }
func (*RefreshNamenodesRequestProto) ProtoMessage()    {}

// *
// void response
type RefreshNamenodesResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *RefreshNamenodesResponseProto) Reset()         { *m = RefreshNamenodesResponseProto{} }
func (m *RefreshNamenodesResponseProto) String() string { return proto.CompactTextString(m) }
func (*RefreshNamenodesResponseProto) ProtoMessage()    {}

// *
// blockPool - block pool to be deleted
// force - if false, delete the block pool only if it is empty.
//         if true, delete the block pool even if it has blocks.
type DeleteBlockPoolRequestProto struct {
	BlockPool        *string `protobuf:"bytes,1,req,name=blockPool" json:"blockPool,omitempty"`
	Force            *bool   `protobuf:"varint,2,req,name=force" json:"force,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *DeleteBlockPoolRequestProto) Reset()         { *m = DeleteBlockPoolRequestProto{} }
func (m *DeleteBlockPoolRequestProto) String() string { return proto.CompactTextString(m) }
func (*DeleteBlockPoolRequestProto) ProtoMessage()    {}

func (m *DeleteBlockPoolRequestProto) GetBlockPool() string {
	if m != nil && m.BlockPool != nil {
		return *m.BlockPool
	}
	return ""
}

func (m *DeleteBlockPoolRequestProto) GetForce() bool {
	if m != nil && m.Force != nil {
		return *m.Force
	}
	return false
}

// *
// void response
type DeleteBlockPoolResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *DeleteBlockPoolResponseProto) Reset()         { *m = DeleteBlockPoolResponseProto{} }
func (m *DeleteBlockPoolResponseProto) String() string { return proto.CompactTextString(m) }
func (*DeleteBlockPoolResponseProto) ProtoMessage()    {}

// *
// Gets the file information where block and its metadata is stored
// block - block for which path information is being requested
// token - block token
//
// This message is deprecated in favor of file descriptor passing.
type GetBlockLocalPathInfoRequestProto struct {
	Block            *ExtendedBlockProto       `protobuf:"bytes,1,req,name=block" json:"block,omitempty"`
	Token            *hadoop_common.TokenProto `protobuf:"bytes,2,req,name=token" json:"token,omitempty"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *GetBlockLocalPathInfoRequestProto) Reset()         { *m = GetBlockLocalPathInfoRequestProto{} }
func (m *GetBlockLocalPathInfoRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetBlockLocalPathInfoRequestProto) ProtoMessage()    {}

func (m *GetBlockLocalPathInfoRequestProto) GetBlock() *ExtendedBlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *GetBlockLocalPathInfoRequestProto) GetToken() *hadoop_common.TokenProto {
	if m != nil {
		return m.Token
	}
	return nil
}

// *
// block - block for which file path information is being returned
// localPath - file path where the block data is stored
// localMetaPath - file path where the block meta data is stored
//
// This message is deprecated in favor of file descriptor passing.
type GetBlockLocalPathInfoResponseProto struct {
	Block            *ExtendedBlockProto `protobuf:"bytes,1,req,name=block" json:"block,omitempty"`
	LocalPath        *string             `protobuf:"bytes,2,req,name=localPath" json:"localPath,omitempty"`
	LocalMetaPath    *string             `protobuf:"bytes,3,req,name=localMetaPath" json:"localMetaPath,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *GetBlockLocalPathInfoResponseProto) Reset()         { *m = GetBlockLocalPathInfoResponseProto{} }
func (m *GetBlockLocalPathInfoResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetBlockLocalPathInfoResponseProto) ProtoMessage()    {}

func (m *GetBlockLocalPathInfoResponseProto) GetBlock() *ExtendedBlockProto {
	if m != nil {
		return m.Block
	}
	return nil
}

func (m *GetBlockLocalPathInfoResponseProto) GetLocalPath() string {
	if m != nil && m.LocalPath != nil {
		return *m.LocalPath
	}
	return ""
}

func (m *GetBlockLocalPathInfoResponseProto) GetLocalMetaPath() string {
	if m != nil && m.LocalMetaPath != nil {
		return *m.LocalMetaPath
	}
	return ""
}

// *
// forUpgrade - if true, clients are advised to wait for restart and quick
//              upgrade restart is instrumented. Otherwise, datanode does
//              the regular shutdown.
type ShutdownDatanodeRequestProto struct {
	ForUpgrade       *bool  `protobuf:"varint,1,req,name=forUpgrade" json:"forUpgrade,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *ShutdownDatanodeRequestProto) Reset()         { *m = ShutdownDatanodeRequestProto{} }
func (m *ShutdownDatanodeRequestProto) String() string { return proto.CompactTextString(m) }
func (*ShutdownDatanodeRequestProto) ProtoMessage()    {}

func (m *ShutdownDatanodeRequestProto) GetForUpgrade() bool {
	if m != nil && m.ForUpgrade != nil {
		return *m.ForUpgrade
	}
	return false
}

type ShutdownDatanodeResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *ShutdownDatanodeResponseProto) Reset()         { *m = ShutdownDatanodeResponseProto{} }
func (m *ShutdownDatanodeResponseProto) String() string { return proto.CompactTextString(m) }
func (*ShutdownDatanodeResponseProto) ProtoMessage()    {}

// * Tell datanode to evict active clients that are writing
type EvictWritersRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *EvictWritersRequestProto) Reset()         { *m = EvictWritersRequestProto{} }
func (m *EvictWritersRequestProto) String() string { return proto.CompactTextString(m) }
func (*EvictWritersRequestProto) ProtoMessage()    {}

type EvictWritersResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *EvictWritersResponseProto) Reset()         { *m = EvictWritersResponseProto{} }
func (m *EvictWritersResponseProto) String() string { return proto.CompactTextString(m) }
func (*EvictWritersResponseProto) ProtoMessage()    {}

// *
// Ping datanode for liveness and quick info
type GetDatanodeInfoRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetDatanodeInfoRequestProto) Reset()         { *m = GetDatanodeInfoRequestProto{} }
func (m *GetDatanodeInfoRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetDatanodeInfoRequestProto) ProtoMessage()    {}

type GetDatanodeInfoResponseProto struct {
	LocalInfo        *DatanodeLocalInfoProto `protobuf:"bytes,1,req,name=localInfo" json:"localInfo,omitempty"`
	XXX_unrecognized []byte                  `json:"-"`
}

func (m *GetDatanodeInfoResponseProto) Reset()         { *m = GetDatanodeInfoResponseProto{} }
func (m *GetDatanodeInfoResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetDatanodeInfoResponseProto) ProtoMessage()    {}

func (m *GetDatanodeInfoResponseProto) GetLocalInfo() *DatanodeLocalInfoProto {
	if m != nil {
		return m.LocalInfo
	}
	return nil
}

type TriggerBlockReportRequestProto struct {
	Incremental      *bool  `protobuf:"varint,1,req,name=incremental" json:"incremental,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *TriggerBlockReportRequestProto) Reset()         { *m = TriggerBlockReportRequestProto{} }
func (m *TriggerBlockReportRequestProto) String() string { return proto.CompactTextString(m) }
func (*TriggerBlockReportRequestProto) ProtoMessage()    {}

func (m *TriggerBlockReportRequestProto) GetIncremental() bool {
	if m != nil && m.Incremental != nil {
		return *m.Incremental
	}
	return false
}

type TriggerBlockReportResponseProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *TriggerBlockReportResponseProto) Reset()         { *m = TriggerBlockReportResponseProto{} }
func (m *TriggerBlockReportResponseProto) String() string { return proto.CompactTextString(m) }
func (*TriggerBlockReportResponseProto) ProtoMessage()    {}

type GetBalancerBandwidthRequestProto struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *GetBalancerBandwidthRequestProto) Reset()         { *m = GetBalancerBandwidthRequestProto{} }
func (m *GetBalancerBandwidthRequestProto) String() string { return proto.CompactTextString(m) }
func (*GetBalancerBandwidthRequestProto) ProtoMessage()    {}

// *
// bandwidth - balancer bandwidth value of the datanode.
type GetBalancerBandwidthResponseProto struct {
	Bandwidth        *uint64 `protobuf:"varint,1,req,name=bandwidth" json:"bandwidth,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *GetBalancerBandwidthResponseProto) Reset()         { *m = GetBalancerBandwidthResponseProto{} }
func (m *GetBalancerBandwidthResponseProto) String() string { return proto.CompactTextString(m) }
func (*GetBalancerBandwidthResponseProto) ProtoMessage()    {}

func (m *GetBalancerBandwidthResponseProto) GetBandwidth() uint64 {
	if m != nil && m.Bandwidth != nil {
		return *m.Bandwidth
	}
	return 0
}

func init() {
}
