// Code generated by protoc-gen-go.
// source: hadoop_ha
// DO NOT EDIT!

/*
Package hadoop_ha is a generated protocol buffer package.

It is generated from these files:
        active_node_info.proto

It has these top-level messages:
        ActiveNodeInfo
*/
package hadoop_ha

import proto "github.com/golang/protobuf/proto"
import "fmt"
import "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ActiveNodeInfo struct {
	NameserviceId    *string `protobuf:"bytes,1,req,name=nameserviceId" json:"nameserviceId,omitempty"`
	NamenodeId       *string `protobuf:"bytes,2,req,name=namenodeId" json:"namenodeId,omitempty"`
	Hostname         *string `protobuf:"bytes,3,req,name=hostname" json:"hostname,omitempty"`
	Port             *int32  `protobuf:"varint,4,req,name=port" json:"port,omitempty"`
	ZkfcPort         *int32  `protobuf:"varint,5,req,name=zkfcPort" json:"zkfcPort,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *ActiveNodeInfo) Reset() { *m = ActiveNodeInfo{} }
func (m *ActiveNodeInfo) String() string { return proto.CompactTextString(m) }
func (*ActiveNodeInfo) ProtoMessage() {}
func (*ActiveNodeInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ActiveNodeInfo) GetNameserviceId() string {
	if m != nil && m.NameserviceId != nil {
		return *m.NameserviceId
	}
	return ""
}

func (m *ActiveNodeInfo) GetNamenodeId() string {
	if m != nil && m.NamenodeId != nil {
		return *m.NamenodeId
	}
	return ""
}

func (m *ActiveNodeInfo) GetHostname() string {
	if m != nil && m.Hostname != nil {
		return *m.Hostname
	}
	return ""
}

func (m *ActiveNodeInfo) GetPort() int32 {
	if m != nil && m.Port != nil {
		return *m.Port
	}
	return 0
}

func (m *ActiveNodeInfo) GetZkfcPort() int32 {
	if m != nil && m.ZkfcPort != nil {
		return *m.ZkfcPort
	}
	return 0
}

func init() {
	proto.RegisterType((*ActiveNodeInfo)(nil), "hadoop_ha.ActiveNodeInfo")
}

var fileDescriptor0 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x1c, 0xca, 0xb1, 0x0a, 0xc2, 0x30,
	0x10, 0xc6, 0x71, 0xac, 0x2d, 0xe8, 0xa1, 0x45, 0x02, 0x42, 0x46, 0x71, 0x72, 0xf2, 0x1d, 0x1c,
	0xbb, 0x88, 0x6f, 0x20, 0x25, 0xb9, 0xd2, 0x20, 0xcd, 0x1d, 0x97, 0xa3, 0x83, 0x4f, 0xaf, 0x97,
	0xf1, 0xfb, 0x7d, 0x7f, 0x00, 0xc5, 0xa2, 0x77, 0x16, 0x52, 0x72, 0xfd, 0x3c, 0x0a, 0xbf, 0x59,
	0x30, 0xa6, 0xa0, 0x24, 0xd7, 0x05, 0xfa, 0x47, 0xd0, 0xb4, 0xe2, 0x93, 0x22, 0x0e, 0x79, 0x22,
	0x77, 0x86, 0x63, 0x1e, 0x17, 0x2c, 0x28, 0x6b, 0x0a, 0x38, 0x44, 0xbf, 0xb9, 0x34, 0xb7, 0xbd,
	0x73, 0x00, 0xc6, 0xd9, 0xb2, 0xe8, 0x9b, 0x6a, 0x27, 0xd8, 0xcd, 0x54, 0xd4, 0xdc, 0x6f, 0xab,
	0x1c, 0xa0, 0x65, 0x12, 0xf5, 0xed, 0x7f, 0x75, 0xf6, 0x7f, 0x3f, 0x53, 0x78, 0x99, 0x74, 0x26,
	0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcb, 0x31, 0x77, 0xc8, 0x8b, 0x00, 0x00, 0x00,
}
