// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kv/kvserver/loqrecovery/loqrecoverypb/recovery.proto

package loqrecoverypb

import (
	fmt "fmt"
	github_com_cockroachdb_cockroach_pkg_roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
	roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// ReplicaInfo contains info about state of range replica for the purpose of range
// recovery. This information should be enough for recovery algorithm to pick a
// survivor replica in when not replicas are available.
// Information includes range descriptor as well as parts of raft state.
type ReplicaInfo struct {
	NodeID                    github_com_cockroachdb_cockroach_pkg_roachpb.NodeID  `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.NodeID" json:"node_id,omitempty"`
	StoreID                   github_com_cockroachdb_cockroach_pkg_roachpb.StoreID `protobuf:"varint,2,opt,name=store_id,json=storeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.StoreID" json:"store_id,omitempty"`
	Desc                      roachpb.RangeDescriptor                              `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc"`
	RaftAppliedIndex          uint64                                               `protobuf:"varint,4,opt,name=raft_applied_index,json=raftAppliedIndex,proto3" json:"raft_applied_index,omitempty"`
	RaftCommittedIndex        uint64                                               `protobuf:"varint,5,opt,name=raft_committed_index,json=raftCommittedIndex,proto3" json:"raft_committed_index,omitempty"`
	HasUncommittedDescriptors bool                                                 `protobuf:"varint,6,opt,name=has_uncommitted_descriptors,json=hasUncommittedDescriptors,proto3" json:"has_uncommitted_descriptors,omitempty"`
}

func (m *ReplicaInfo) Reset()         { *m = ReplicaInfo{} }
func (m *ReplicaInfo) String() string { return proto.CompactTextString(m) }
func (*ReplicaInfo) ProtoMessage()    {}
func (*ReplicaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a2cc96948b3cab0, []int{0}
}
func (m *ReplicaInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReplicaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ReplicaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplicaInfo.Merge(m, src)
}
func (m *ReplicaInfo) XXX_Size() int {
	return m.Size()
}
func (m *ReplicaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplicaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ReplicaInfo proto.InternalMessageInfo

// Collection of replica information gathered from a collect-info run on a single node.
type NodeReplicaInfo struct {
	Replicas []ReplicaInfo `protobuf:"bytes,1,rep,name=replicas,proto3" json:"replicas"`
}

func (m *NodeReplicaInfo) Reset()         { *m = NodeReplicaInfo{} }
func (m *NodeReplicaInfo) String() string { return proto.CompactTextString(m) }
func (*NodeReplicaInfo) ProtoMessage()    {}
func (*NodeReplicaInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a2cc96948b3cab0, []int{1}
}
func (m *NodeReplicaInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *NodeReplicaInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *NodeReplicaInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeReplicaInfo.Merge(m, src)
}
func (m *NodeReplicaInfo) XXX_Size() int {
	return m.Size()
}
func (m *NodeReplicaInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeReplicaInfo.DiscardUnknown(m)
}

var xxx_messageInfo_NodeReplicaInfo proto.InternalMessageInfo

// ReplicaUpdate contains information that needs to be updated on replica on the node
// to make it a designated survivor so that replica could act as a source of truth when
// doing loss of quorum recovery.
type ReplicaUpdate struct {
	RangeID       github_com_cockroachdb_cockroach_pkg_roachpb.RangeID   `protobuf:"varint,1,opt,name=range_id,json=rangeId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.RangeID" json:"range_id,omitempty" yaml:"RangeID"`
	StartKey      RecoveryKey                                            `protobuf:"bytes,2,opt,name=start_key,json=startKey,proto3,casttype=RecoveryKey" json:"start_key,omitempty" yaml:"StartKey"`
	OldReplicaID  github_com_cockroachdb_cockroach_pkg_roachpb.ReplicaID `protobuf:"varint,3,opt,name=old_replica_id,json=oldReplicaId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.ReplicaID" json:"old_replica_id,omitempty" yaml:"OldReplicaID"`
	NewReplica    *roachpb.ReplicaDescriptor                             `protobuf:"bytes,4,opt,name=new_replica,json=newReplica,proto3" json:"new_replica,omitempty" yaml:"NewReplica"`
	NextReplicaID github_com_cockroachdb_cockroach_pkg_roachpb.ReplicaID `protobuf:"varint,5,opt,name=next_replica_id,json=nextReplicaId,proto3,casttype=github.com/cockroachdb/cockroach/pkg/roachpb.ReplicaID" json:"next_replica_id,omitempty" yaml:"NextReplicaID"`
}

func (m *ReplicaUpdate) Reset()      { *m = ReplicaUpdate{} }
func (*ReplicaUpdate) ProtoMessage() {}
func (*ReplicaUpdate) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a2cc96948b3cab0, []int{2}
}
func (m *ReplicaUpdate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReplicaUpdate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ReplicaUpdate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplicaUpdate.Merge(m, src)
}
func (m *ReplicaUpdate) XXX_Size() int {
	return m.Size()
}
func (m *ReplicaUpdate) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplicaUpdate.DiscardUnknown(m)
}

var xxx_messageInfo_ReplicaUpdate proto.InternalMessageInfo

// ReplicaUpdatePlan Collection of updates for all recoverable replicas in the cluster.
type ReplicaUpdatePlan struct {
	Updates []ReplicaUpdate `protobuf:"bytes,1,rep,name=updates,proto3" json:"updates"`
}

func (m *ReplicaUpdatePlan) Reset()         { *m = ReplicaUpdatePlan{} }
func (m *ReplicaUpdatePlan) String() string { return proto.CompactTextString(m) }
func (*ReplicaUpdatePlan) ProtoMessage()    {}
func (*ReplicaUpdatePlan) Descriptor() ([]byte, []int) {
	return fileDescriptor_5a2cc96948b3cab0, []int{3}
}
func (m *ReplicaUpdatePlan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ReplicaUpdatePlan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *ReplicaUpdatePlan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ReplicaUpdatePlan.Merge(m, src)
}
func (m *ReplicaUpdatePlan) XXX_Size() int {
	return m.Size()
}
func (m *ReplicaUpdatePlan) XXX_DiscardUnknown() {
	xxx_messageInfo_ReplicaUpdatePlan.DiscardUnknown(m)
}

var xxx_messageInfo_ReplicaUpdatePlan proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ReplicaInfo)(nil), "cockroach.kv.kvserver.loqrecovery.loqrecoverypb.ReplicaInfo")
	proto.RegisterType((*NodeReplicaInfo)(nil), "cockroach.kv.kvserver.loqrecovery.loqrecoverypb.NodeReplicaInfo")
	proto.RegisterType((*ReplicaUpdate)(nil), "cockroach.kv.kvserver.loqrecovery.loqrecoverypb.ReplicaUpdate")
	proto.RegisterType((*ReplicaUpdatePlan)(nil), "cockroach.kv.kvserver.loqrecovery.loqrecoverypb.ReplicaUpdatePlan")
}

func init() {
	proto.RegisterFile("kv/kvserver/loqrecovery/loqrecoverypb/recovery.proto", fileDescriptor_5a2cc96948b3cab0)
}

var fileDescriptor_5a2cc96948b3cab0 = []byte{
	// 675 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0x31, 0x4f, 0xdb, 0x40,
	0x14, 0xc7, 0xe3, 0x26, 0x24, 0xe9, 0x85, 0x40, 0xb9, 0xd2, 0x2a, 0xa5, 0x92, 0x1d, 0x59, 0x1d,
	0x32, 0x54, 0x76, 0x05, 0xa8, 0x03, 0x42, 0x54, 0x04, 0x16, 0x0b, 0x89, 0x56, 0x46, 0x2c, 0xa8,
	0x22, 0xba, 0xf8, 0x8e, 0xc4, 0x8a, 0xe3, 0x33, 0xe7, 0x23, 0x90, 0x4f, 0xd0, 0xa1, 0x4b, 0xc7,
	0x8e, 0xdd, 0xfa, 0x2d, 0x3a, 0x33, 0x32, 0x32, 0x59, 0xad, 0xf9, 0x06, 0x8c, 0x9d, 0xaa, 0x3b,
	0xdb, 0xb1, 0x91, 0xba, 0xa4, 0xdd, 0xce, 0xef, 0xbd, 0xff, 0xef, 0xde, 0x7b, 0xf7, 0x97, 0xc1,
	0xe6, 0x68, 0x62, 0x8e, 0x26, 0x21, 0x61, 0x13, 0xc2, 0x4c, 0x8f, 0x9e, 0x33, 0xe2, 0xd0, 0x09,
	0x61, 0xd3, 0xe2, 0x39, 0xe8, 0x9b, 0xd9, 0xd1, 0x08, 0x18, 0xe5, 0x14, 0x9a, 0x0e, 0x75, 0x46,
	0x8c, 0x22, 0x67, 0x68, 0x8c, 0x26, 0x46, 0xa6, 0x37, 0x0a, 0x1a, 0xe3, 0x81, 0x7e, 0xed, 0xb9,
	0x2c, 0x0e, 0xfa, 0xe6, 0x98, 0x70, 0x84, 0x11, 0x47, 0x09, 0x68, 0x6d, 0x75, 0x40, 0x07, 0x54,
	0x1e, 0x4d, 0x71, 0x4a, 0xa2, 0xfa, 0xf7, 0x32, 0x68, 0xd8, 0x24, 0xf0, 0x5c, 0x07, 0x59, 0xfe,
	0x19, 0x85, 0x27, 0xa0, 0xe6, 0x53, 0x4c, 0x7a, 0x2e, 0x6e, 0x29, 0x6d, 0xa5, 0xb3, 0xd0, 0xdd,
	0x8d, 0x23, 0xad, 0x7a, 0x48, 0x31, 0xb1, 0xf6, 0x7f, 0x47, 0xda, 0xc6, 0xc0, 0xe5, 0xc3, 0x8b,
	0xbe, 0xe1, 0xd0, 0x71, 0xde, 0x18, 0xee, 0xe7, 0x67, 0x33, 0x18, 0x0d, 0xcc, 0xb4, 0x03, 0x23,
	0x91, 0xd9, 0x55, 0x41, 0xb4, 0x30, 0x3c, 0x05, 0xf5, 0x90, 0x53, 0x26, 0xe1, 0x8f, 0x24, 0x7c,
	0x2f, 0x8e, 0xb4, 0xda, 0x91, 0x88, 0x49, 0xfa, 0xe6, 0x5c, 0xf4, 0x54, 0x67, 0xd7, 0x24, 0xd4,
	0xc2, 0x70, 0x1b, 0x54, 0x30, 0x09, 0x9d, 0x56, 0xb9, 0xad, 0x74, 0x1a, 0xeb, 0xba, 0x91, 0x6f,
	0x2e, 0x93, 0xd8, 0xc8, 0x1f, 0x90, 0x7d, 0x12, 0x3a, 0xcc, 0x0d, 0x38, 0x65, 0xdd, 0xca, 0x75,
	0xa4, 0x95, 0x6c, 0xa9, 0x82, 0xaf, 0x01, 0x64, 0xe8, 0x8c, 0xf7, 0x50, 0x10, 0x78, 0x2e, 0xc1,
	0x3d, 0xd7, 0xc7, 0xe4, 0xaa, 0x55, 0x69, 0x2b, 0x9d, 0x8a, 0xfd, 0x44, 0x64, 0x76, 0x93, 0x84,
	0x25, 0xe2, 0xf0, 0x0d, 0x58, 0x95, 0xd5, 0x0e, 0x1d, 0x8f, 0x5d, 0xce, 0x67, 0xf5, 0x0b, 0xb2,
	0x5e, 0x92, 0xf6, 0xb2, 0x54, 0xa2, 0xd8, 0x01, 0x2f, 0x87, 0x28, 0xec, 0x5d, 0xf8, 0xb9, 0x04,
	0xcf, 0x3a, 0x09, 0x5b, 0xd5, 0xb6, 0xd2, 0xa9, 0xdb, 0x2f, 0x86, 0x28, 0x3c, 0xce, 0x2b, 0xf2,
	0x56, 0x43, 0xfd, 0x1c, 0x2c, 0x8b, 0x7d, 0x16, 0x1f, 0xeb, 0x14, 0xd4, 0x59, 0xf2, 0x19, 0xb6,
	0x94, 0x76, 0xb9, 0xd3, 0x58, 0xdf, 0x36, 0xe6, 0xb4, 0x8b, 0x51, 0xe0, 0xa5, 0xeb, 0x98, 0x31,
	0xf5, 0x1f, 0x15, 0xd0, 0x4c, 0xf3, 0xc7, 0x01, 0x46, 0x9c, 0xc0, 0x00, 0xd4, 0x99, 0xd8, 0x61,
	0xe6, 0x8f, 0x72, 0xf7, 0x58, 0x3c, 0xa1, 0xdc, 0xab, 0xb5, 0x7f, 0x1f, 0x69, 0x4b, 0x53, 0x34,
	0xf6, 0xb6, 0xf4, 0x34, 0xa0, 0xcf, 0xfd, 0xa8, 0xa9, 0xd0, 0xae, 0xc9, 0x6b, 0x2c, 0x0c, 0xdf,
	0x81, 0xc7, 0x21, 0x47, 0x8c, 0xf7, 0x46, 0x64, 0x2a, 0x5d, 0xb3, 0xd8, 0xd5, 0xef, 0x23, 0x6d,
	0x39, 0xb9, 0xe7, 0x48, 0xa4, 0x0e, 0xc8, 0x54, 0x5c, 0xd4, 0xb0, 0xd3, 0xa9, 0x0e, 0xc8, 0xd4,
	0xae, 0x87, 0x69, 0x06, 0x7e, 0x52, 0xc0, 0x12, 0xf5, 0x70, 0x2f, 0x9d, 0x4a, 0x74, 0x5e, 0x96,
	0xe6, 0x43, 0x71, 0xa4, 0x2d, 0xbe, 0xf7, 0x70, 0xb6, 0x01, 0xd1, 0xfe, 0xd3, 0x04, 0x5b, 0x8c,
	0x0a, 0xf4, 0xdb, 0xf9, 0x66, 0xc8, 0xa4, 0xf6, 0x22, 0xcd, 0x41, 0x18, 0x7e, 0x04, 0x0d, 0x9f,
	0x5c, 0x66, 0x8d, 0x48, 0x6b, 0x35, 0xd6, 0x5f, 0xfd, 0xcd, 0xa6, 0x49, 0x45, 0xc1, 0xa8, 0xcf,
	0xee, 0x23, 0x6d, 0x25, 0xe9, 0xed, 0x90, 0x5c, 0xa6, 0x79, 0xdd, 0x06, 0xfe, 0xec, 0x03, 0x7e,
	0x56, 0xc0, 0xb2, 0x4f, 0xae, 0x78, 0x71, 0xd0, 0x05, 0x39, 0xa8, 0x13, 0x47, 0x5a, 0xf3, 0x90,
	0x5c, 0xf1, 0xe2, 0xa4, 0xab, 0x19, 0xad, 0x10, 0xfe, 0x9f, 0x51, 0x9b, 0x7e, 0x81, 0x84, 0xb7,
	0x2a, 0x5f, 0xbf, 0x69, 0x25, 0x3d, 0x04, 0x2b, 0x0f, 0xfc, 0xf3, 0xc1, 0x43, 0x3e, 0x3c, 0x05,
	0xb5, 0x0b, 0xf9, 0x95, 0x99, 0x76, 0xe7, 0x5f, 0x4d, 0x9b, 0x40, 0x53, 0xdb, 0x66, 0xd0, 0xae,
	0x79, 0xfd, 0x4b, 0x2d, 0x5d, 0xc7, 0xaa, 0x72, 0x13, 0xab, 0xca, 0x6d, 0xac, 0x2a, 0x3f, 0x63,
	0x55, 0xf9, 0x72, 0xa7, 0x96, 0x6e, 0xee, 0xd4, 0xd2, 0xed, 0x9d, 0x5a, 0x3a, 0x69, 0x3e, 0xa0,
	0xf5, 0xab, 0xf2, 0x57, 0xb8, 0xf1, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd3, 0xa8, 0xda, 0x73, 0xa1,
	0x05, 0x00, 0x00,
}

func (m *ReplicaInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplicaInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReplicaInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HasUncommittedDescriptors {
		i--
		if m.HasUncommittedDescriptors {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x30
	}
	if m.RaftCommittedIndex != 0 {
		i = encodeVarintRecovery(dAtA, i, uint64(m.RaftCommittedIndex))
		i--
		dAtA[i] = 0x28
	}
	if m.RaftAppliedIndex != 0 {
		i = encodeVarintRecovery(dAtA, i, uint64(m.RaftAppliedIndex))
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.Desc.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintRecovery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.StoreID != 0 {
		i = encodeVarintRecovery(dAtA, i, uint64(m.StoreID))
		i--
		dAtA[i] = 0x10
	}
	if m.NodeID != 0 {
		i = encodeVarintRecovery(dAtA, i, uint64(m.NodeID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *NodeReplicaInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *NodeReplicaInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *NodeReplicaInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Replicas) > 0 {
		for iNdEx := len(m.Replicas) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Replicas[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRecovery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *ReplicaUpdate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplicaUpdate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReplicaUpdate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.NextReplicaID != 0 {
		i = encodeVarintRecovery(dAtA, i, uint64(m.NextReplicaID))
		i--
		dAtA[i] = 0x28
	}
	if m.NewReplica != nil {
		{
			size, err := m.NewReplica.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintRecovery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.OldReplicaID != 0 {
		i = encodeVarintRecovery(dAtA, i, uint64(m.OldReplicaID))
		i--
		dAtA[i] = 0x18
	}
	if len(m.StartKey) > 0 {
		i -= len(m.StartKey)
		copy(dAtA[i:], m.StartKey)
		i = encodeVarintRecovery(dAtA, i, uint64(len(m.StartKey)))
		i--
		dAtA[i] = 0x12
	}
	if m.RangeID != 0 {
		i = encodeVarintRecovery(dAtA, i, uint64(m.RangeID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ReplicaUpdatePlan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ReplicaUpdatePlan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ReplicaUpdatePlan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Updates) > 0 {
		for iNdEx := len(m.Updates) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Updates[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintRecovery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintRecovery(dAtA []byte, offset int, v uint64) int {
	offset -= sovRecovery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ReplicaInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NodeID != 0 {
		n += 1 + sovRecovery(uint64(m.NodeID))
	}
	if m.StoreID != 0 {
		n += 1 + sovRecovery(uint64(m.StoreID))
	}
	l = m.Desc.Size()
	n += 1 + l + sovRecovery(uint64(l))
	if m.RaftAppliedIndex != 0 {
		n += 1 + sovRecovery(uint64(m.RaftAppliedIndex))
	}
	if m.RaftCommittedIndex != 0 {
		n += 1 + sovRecovery(uint64(m.RaftCommittedIndex))
	}
	if m.HasUncommittedDescriptors {
		n += 2
	}
	return n
}

func (m *NodeReplicaInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Replicas) > 0 {
		for _, e := range m.Replicas {
			l = e.Size()
			n += 1 + l + sovRecovery(uint64(l))
		}
	}
	return n
}

func (m *ReplicaUpdate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.RangeID != 0 {
		n += 1 + sovRecovery(uint64(m.RangeID))
	}
	l = len(m.StartKey)
	if l > 0 {
		n += 1 + l + sovRecovery(uint64(l))
	}
	if m.OldReplicaID != 0 {
		n += 1 + sovRecovery(uint64(m.OldReplicaID))
	}
	if m.NewReplica != nil {
		l = m.NewReplica.Size()
		n += 1 + l + sovRecovery(uint64(l))
	}
	if m.NextReplicaID != 0 {
		n += 1 + sovRecovery(uint64(m.NextReplicaID))
	}
	return n
}

func (m *ReplicaUpdatePlan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Updates) > 0 {
		for _, e := range m.Updates {
			l = e.Size()
			n += 1 + l + sovRecovery(uint64(l))
		}
	}
	return n
}

func sovRecovery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRecovery(x uint64) (n int) {
	return sovRecovery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ReplicaInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecovery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReplicaInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplicaInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeID", wireType)
			}
			m.NodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NodeID |= github_com_cockroachdb_cockroach_pkg_roachpb.NodeID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StoreID", wireType)
			}
			m.StoreID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StoreID |= github_com_cockroachdb_cockroach_pkg_roachpb.StoreID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Desc", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRecovery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Desc.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftAppliedIndex", wireType)
			}
			m.RaftAppliedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftAppliedIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RaftCommittedIndex", wireType)
			}
			m.RaftCommittedIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RaftCommittedIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasUncommittedDescriptors", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.HasUncommittedDescriptors = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipRecovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecovery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *NodeReplicaInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecovery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: NodeReplicaInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: NodeReplicaInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replicas", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRecovery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Replicas = append(m.Replicas, ReplicaInfo{})
			if err := m.Replicas[len(m.Replicas)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecovery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReplicaUpdate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecovery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReplicaUpdate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplicaUpdate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RangeID", wireType)
			}
			m.RangeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RangeID |= github_com_cockroachdb_cockroach_pkg_roachpb.RangeID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthRecovery
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthRecovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StartKey = append(m.StartKey[:0], dAtA[iNdEx:postIndex]...)
			if m.StartKey == nil {
				m.StartKey = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field OldReplicaID", wireType)
			}
			m.OldReplicaID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.OldReplicaID |= github_com_cockroachdb_cockroach_pkg_roachpb.ReplicaID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NewReplica", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRecovery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.NewReplica == nil {
				m.NewReplica = &roachpb.ReplicaDescriptor{}
			}
			if err := m.NewReplica.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextReplicaID", wireType)
			}
			m.NextReplicaID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextReplicaID |= github_com_cockroachdb_cockroach_pkg_roachpb.ReplicaID(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipRecovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecovery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ReplicaUpdatePlan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRecovery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ReplicaUpdatePlan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ReplicaUpdatePlan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Updates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRecovery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthRecovery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthRecovery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Updates = append(m.Updates, ReplicaUpdate{})
			if err := m.Updates[len(m.Updates)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRecovery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRecovery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipRecovery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRecovery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRecovery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowRecovery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthRecovery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRecovery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRecovery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRecovery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRecovery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRecovery = fmt.Errorf("proto: unexpected end of group")
)
