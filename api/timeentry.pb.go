// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: timeentry.proto

/*
	Package api is a generated protocol buffer package.

	It is generated from these files:
		timeentry.proto
		user.proto

	It has these top-level messages:
		TimeEntry
		LogoutInfo
		AuthInfo
		User
*/
package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type TimeEntry_Status int32

const (
	TimeEntry_ClockedOut TimeEntry_Status = 0
	TimeEntry_ClockedIn  TimeEntry_Status = 1
)

var TimeEntry_Status_name = map[int32]string{
	0: "ClockedOut",
	1: "ClockedIn",
}
var TimeEntry_Status_value = map[string]int32{
	"ClockedOut": 0,
	"ClockedIn":  1,
}

func (x TimeEntry_Status) String() string {
	return proto.EnumName(TimeEntry_Status_name, int32(x))
}
func (TimeEntry_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptorTimeentry, []int{0, 0} }

type TimeEntry struct {
	ID        int32            `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	UserID    string           `protobuf:"bytes,2,opt,name=UserID,proto3" json:"UserID,omitempty"`
	Timestamp int64            `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	State     TimeEntry_Status `protobuf:"varint,4,opt,name=State,proto3,enum=api.TimeEntry_Status" json:"State,omitempty"`
}

func (m *TimeEntry) Reset()                    { *m = TimeEntry{} }
func (m *TimeEntry) String() string            { return proto.CompactTextString(m) }
func (*TimeEntry) ProtoMessage()               {}
func (*TimeEntry) Descriptor() ([]byte, []int) { return fileDescriptorTimeentry, []int{0} }

func (m *TimeEntry) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *TimeEntry) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *TimeEntry) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *TimeEntry) GetState() TimeEntry_Status {
	if m != nil {
		return m.State
	}
	return TimeEntry_ClockedOut
}

func init() {
	proto.RegisterType((*TimeEntry)(nil), "api.TimeEntry")
	proto.RegisterEnum("api.TimeEntry_Status", TimeEntry_Status_name, TimeEntry_Status_value)
}
func (m *TimeEntry) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimeEntry) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ID != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTimeentry(dAtA, i, uint64(m.ID))
	}
	if len(m.UserID) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintTimeentry(dAtA, i, uint64(len(m.UserID)))
		i += copy(dAtA[i:], m.UserID)
	}
	if m.Timestamp != 0 {
		dAtA[i] = 0x18
		i++
		i = encodeVarintTimeentry(dAtA, i, uint64(m.Timestamp))
	}
	if m.State != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintTimeentry(dAtA, i, uint64(m.State))
	}
	return i, nil
}

func encodeVarintTimeentry(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TimeEntry) Size() (n int) {
	var l int
	_ = l
	if m.ID != 0 {
		n += 1 + sovTimeentry(uint64(m.ID))
	}
	l = len(m.UserID)
	if l > 0 {
		n += 1 + l + sovTimeentry(uint64(l))
	}
	if m.Timestamp != 0 {
		n += 1 + sovTimeentry(uint64(m.Timestamp))
	}
	if m.State != 0 {
		n += 1 + sovTimeentry(uint64(m.State))
	}
	return n
}

func sovTimeentry(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTimeentry(x uint64) (n int) {
	return sovTimeentry(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TimeEntry) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimeentry
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TimeEntry: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeEntry: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			m.ID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeentry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ID |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeentry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTimeentry
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeentry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimeentry
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= (TimeEntry_Status(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTimeentry(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimeentry
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
func skipTimeentry(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimeentry
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
					return 0, ErrIntOverflowTimeentry
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTimeentry
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTimeentry
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTimeentry
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTimeentry(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTimeentry = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimeentry   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("timeentry.proto", fileDescriptorTimeentry) }

var fileDescriptorTimeentry = []byte{
	// 199 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0xc9, 0xcc, 0x4d,
	0x4d, 0xcd, 0x2b, 0x29, 0xaa, 0xd4, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4e, 0x2c, 0xc8,
	0x54, 0x5a, 0xce, 0xc8, 0xc5, 0x19, 0x92, 0x99, 0x9b, 0xea, 0x0a, 0x92, 0x10, 0xe2, 0xe3, 0x62,
	0xf2, 0x74, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x62, 0xf2, 0x74, 0x11, 0x12, 0xe3, 0x62,
	0x0b, 0x2d, 0x4e, 0x2d, 0xf2, 0x74, 0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0xf2, 0x84,
	0x64, 0x20, 0x9a, 0x8a, 0x4b, 0x12, 0x73, 0x0b, 0x24, 0x98, 0x15, 0x18, 0x35, 0x98, 0x83, 0x10,
	0x02, 0x42, 0xda, 0x5c, 0xac, 0xc1, 0x25, 0x89, 0x25, 0xa9, 0x12, 0x2c, 0x0a, 0x8c, 0x1a, 0x7c,
	0x46, 0xa2, 0x7a, 0x89, 0x05, 0x99, 0x7a, 0x70, 0x4b, 0xf4, 0x40, 0x72, 0xa5, 0xc5, 0x41, 0x10,
	0x35, 0x4a, 0xea, 0x5c, 0x6c, 0x10, 0x01, 0x21, 0x3e, 0x2e, 0x2e, 0xe7, 0x9c, 0xfc, 0xe4, 0xec,
	0xd4, 0x14, 0xff, 0xd2, 0x12, 0x01, 0x06, 0x21, 0x5e, 0x2e, 0x4e, 0x28, 0xdf, 0x33, 0x4f, 0x80,
	0xd1, 0x49, 0xe0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c,
	0xf0, 0x58, 0x8e, 0x21, 0x89, 0x0d, 0xec, 0x0f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99,
	0x00, 0xae, 0x66, 0xda, 0x00, 0x00, 0x00,
}