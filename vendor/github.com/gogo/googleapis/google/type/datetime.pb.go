// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/type/datetime.proto

package _type

import (
	bytes "bytes"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	types "github.com/gogo/protobuf/types"
	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"
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

// Represents civil time in one of a few possible ways:
//
//  * When utc_offset is set and time_zone is unset: a civil time on a calendar
//    day with a particular offset from UTC.
//  * When time_zone is set and utc_offset is unset: a civil time on a calendar
//    day in a particular time zone.
//  * When neither time_zone nor utc_offset is set: a civil time on a calendar
//    day in local time.
//
// The date is relative to the Proleptic Gregorian Calendar.
//
// If year is 0, the DateTime is considered not to have a specific year. month
// and day must have valid, non-zero values.
//
// This type is more flexible than some applications may want. Make sure to
// document and validate your application's limitations.
type DateTime struct {
	// Optional. Year of date. Must be from 1 to 9999, or 0 if specifying a
	// datetime without a year.
	Year int32 `protobuf:"varint,1,opt,name=year,proto3" json:"year,omitempty"`
	// Required. Month of year. Must be from 1 to 12.
	Month int32 `protobuf:"varint,2,opt,name=month,proto3" json:"month,omitempty"`
	// Required. Day of month. Must be from 1 to 31 and valid for the year and
	// month.
	Day int32 `protobuf:"varint,3,opt,name=day,proto3" json:"day,omitempty"`
	// Required. Hours of day in 24 hour format. Should be from 0 to 23. An API
	// may choose to allow the value "24:00:00" for scenarios like business
	// closing time.
	Hours int32 `protobuf:"varint,4,opt,name=hours,proto3" json:"hours,omitempty"`
	// Required. Minutes of hour of day. Must be from 0 to 59.
	Minutes int32 `protobuf:"varint,5,opt,name=minutes,proto3" json:"minutes,omitempty"`
	// Required. Seconds of minutes of the time. Must normally be from 0 to 59. An
	// API may allow the value 60 if it allows leap-seconds.
	Seconds int32 `protobuf:"varint,6,opt,name=seconds,proto3" json:"seconds,omitempty"`
	// Required. Fractions of seconds in nanoseconds. Must be from 0 to
	// 999,999,999.
	Nanos int32 `protobuf:"varint,7,opt,name=nanos,proto3" json:"nanos,omitempty"`
	// Optional. Specifies either the UTC offset or the time zone of the DateTime.
	// Choose carefully between them, considering that time zone data may change
	// in the future (for example, a country modifies their DST start/end dates,
	// and future DateTimes in the affected range had already been stored).
	// If omitted, the DateTime is considered to be in local time.
	//
	// Types that are valid to be assigned to TimeOffset:
	//	*DateTime_UtcOffset
	//	*DateTime_TimeZone
	TimeOffset           isDateTime_TimeOffset `protobuf_oneof:"time_offset"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *DateTime) Reset()      { *m = DateTime{} }
func (*DateTime) ProtoMessage() {}
func (*DateTime) Descriptor() ([]byte, []int) {
	return fileDescriptor_247e8eac669493f4, []int{0}
}
func (m *DateTime) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DateTime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DateTime.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DateTime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DateTime.Merge(m, src)
}
func (m *DateTime) XXX_Size() int {
	return m.Size()
}
func (m *DateTime) XXX_DiscardUnknown() {
	xxx_messageInfo_DateTime.DiscardUnknown(m)
}

var xxx_messageInfo_DateTime proto.InternalMessageInfo

type isDateTime_TimeOffset interface {
	isDateTime_TimeOffset()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type DateTime_UtcOffset struct {
	UtcOffset *types.Duration `protobuf:"bytes,8,opt,name=utc_offset,json=utcOffset,proto3,oneof"`
}
type DateTime_TimeZone struct {
	TimeZone *TimeZone `protobuf:"bytes,9,opt,name=time_zone,json=timeZone,proto3,oneof"`
}

func (*DateTime_UtcOffset) isDateTime_TimeOffset() {}
func (*DateTime_TimeZone) isDateTime_TimeOffset()  {}

func (m *DateTime) GetTimeOffset() isDateTime_TimeOffset {
	if m != nil {
		return m.TimeOffset
	}
	return nil
}

func (m *DateTime) GetYear() int32 {
	if m != nil {
		return m.Year
	}
	return 0
}

func (m *DateTime) GetMonth() int32 {
	if m != nil {
		return m.Month
	}
	return 0
}

func (m *DateTime) GetDay() int32 {
	if m != nil {
		return m.Day
	}
	return 0
}

func (m *DateTime) GetHours() int32 {
	if m != nil {
		return m.Hours
	}
	return 0
}

func (m *DateTime) GetMinutes() int32 {
	if m != nil {
		return m.Minutes
	}
	return 0
}

func (m *DateTime) GetSeconds() int32 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *DateTime) GetNanos() int32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

func (m *DateTime) GetUtcOffset() *types.Duration {
	if x, ok := m.GetTimeOffset().(*DateTime_UtcOffset); ok {
		return x.UtcOffset
	}
	return nil
}

func (m *DateTime) GetTimeZone() *TimeZone {
	if x, ok := m.GetTimeOffset().(*DateTime_TimeZone); ok {
		return x.TimeZone
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DateTime) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DateTime_UtcOffset)(nil),
		(*DateTime_TimeZone)(nil),
	}
}

func (*DateTime) XXX_MessageName() string {
	return "google.type.DateTime"
}

// Represents a time zone from the
// [IANA Time Zone Database](https://www.iana.org/time-zones).
type TimeZone struct {
	// IANA Time Zone Database time zone, e.g. "America/New_York".
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Optional. IANA Time Zone Database version number, e.g. "2019a".
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeZone) Reset()      { *m = TimeZone{} }
func (*TimeZone) ProtoMessage() {}
func (*TimeZone) Descriptor() ([]byte, []int) {
	return fileDescriptor_247e8eac669493f4, []int{1}
}
func (m *TimeZone) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TimeZone) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TimeZone.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TimeZone) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeZone.Merge(m, src)
}
func (m *TimeZone) XXX_Size() int {
	return m.Size()
}
func (m *TimeZone) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeZone.DiscardUnknown(m)
}

var xxx_messageInfo_TimeZone proto.InternalMessageInfo

func (m *TimeZone) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *TimeZone) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (*TimeZone) XXX_MessageName() string {
	return "google.type.TimeZone"
}
func init() {
	proto.RegisterType((*DateTime)(nil), "google.type.DateTime")
	proto.RegisterType((*TimeZone)(nil), "google.type.TimeZone")
}

func init() { proto.RegisterFile("google/type/datetime.proto", fileDescriptor_247e8eac669493f4) }

var fileDescriptor_247e8eac669493f4 = []byte{
	// 371 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x51, 0x3d, 0x6f, 0xea, 0x30,
	0x14, 0x8d, 0xc3, 0x57, 0x62, 0xc4, 0x7b, 0x4f, 0xd6, 0xab, 0xe4, 0x32, 0x58, 0x88, 0x89, 0x29,
	0x48, 0x2d, 0x53, 0x47, 0x84, 0x54, 0xb6, 0xa2, 0x88, 0x09, 0xa9, 0x42, 0x21, 0x31, 0x10, 0xa9,
	0xf1, 0x45, 0x89, 0x53, 0x89, 0x4e, 0xfd, 0x2d, 0x5d, 0xda, 0x9f, 0xd2, 0xb1, 0x63, 0xc7, 0x92,
	0x2e, 0x1d, 0x19, 0x19, 0x2b, 0xdb, 0x89, 0xc4, 0x76, 0xcf, 0xc7, 0xbd, 0xc9, 0x39, 0xc6, 0xdd,
	0x0d, 0xc0, 0xe6, 0x81, 0x0f, 0xe5, 0x7e, 0xc7, 0x87, 0x51, 0x20, 0xb9, 0x8c, 0x13, 0xee, 0xed,
	0x52, 0x90, 0x40, 0xda, 0x46, 0xf3, 0x94, 0xd6, 0x65, 0xa5, 0x51, 0x4b, 0xab, 0x7c, 0x3d, 0x8c,
	0xf2, 0x34, 0x90, 0x31, 0x08, 0x63, 0xee, 0xbf, 0xda, 0xd8, 0x99, 0x04, 0x92, 0xcf, 0xe3, 0x84,
	0x13, 0x82, 0xeb, 0x7b, 0x1e, 0xa4, 0x14, 0xf5, 0xd0, 0xa0, 0xe1, 0xeb, 0x99, 0xfc, 0xc7, 0x8d,
	0x04, 0x84, 0xdc, 0x52, 0x5b, 0x93, 0x06, 0x90, 0x7f, 0xb8, 0x16, 0x05, 0x7b, 0x5a, 0xd3, 0x9c,
	0x1a, 0x95, 0x6f, 0x0b, 0x79, 0x9a, 0xd1, 0xba, 0xf1, 0x69, 0x40, 0x28, 0x6e, 0x25, 0xb1, 0xc8,
	0x25, 0xcf, 0x68, 0x43, 0xf3, 0x15, 0x54, 0x4a, 0xc6, 0x43, 0x10, 0x51, 0x46, 0x9b, 0x46, 0x29,
	0xa1, 0xba, 0x24, 0x02, 0x01, 0x19, 0x6d, 0x99, 0x4b, 0x1a, 0x90, 0x1b, 0x8c, 0x73, 0x19, 0x2e,
	0x61, 0xbd, 0xce, 0xb8, 0xa4, 0x4e, 0x0f, 0x0d, 0xda, 0x57, 0x97, 0x5e, 0x19, 0xb5, 0x4a, 0xe7,
	0x4d, 0xca, 0x74, 0x53, 0xcb, 0x77, 0x73, 0x19, 0xde, 0x69, 0x37, 0x19, 0x61, 0x57, 0xf5, 0xb3,
	0x7c, 0x02, 0xc1, 0xa9, 0xab, 0x57, 0x2f, 0xbc, 0xb3, 0x96, 0x3c, 0x95, 0x7e, 0x01, 0x82, 0x4f,
	0x2d, 0xdf, 0x91, 0xe5, 0x3c, 0xee, 0xe0, 0xb6, 0xde, 0x32, 0x9f, 0xec, 0x8f, 0xb0, 0x53, 0xd9,
	0xc8, 0x1f, 0x6c, 0xc7, 0x91, 0xae, 0xc9, 0xf5, 0xed, 0x38, 0x52, 0x61, 0x1e, 0x79, 0x9a, 0xc5,
	0x20, 0x74, 0x4d, 0xae, 0x5f, 0xc1, 0xf1, 0xfd, 0xe7, 0x81, 0x59, 0xc7, 0x03, 0x43, 0xa7, 0x03,
	0x43, 0xcf, 0x05, 0x43, 0x6f, 0x05, 0x43, 0xef, 0x05, 0x43, 0x1f, 0x05, 0x43, 0x5f, 0x05, 0x43,
	0x3f, 0x05, 0xb3, 0x8e, 0x8a, 0xfb, 0x66, 0x08, 0xff, 0x0d, 0x21, 0x39, 0xff, 0xb1, 0x71, 0xa7,
	0x7a, 0x9b, 0x99, 0x4a, 0x38, 0x43, 0x8b, 0xba, 0xa2, 0x4f, 0x08, 0xbd, 0xd8, 0xb5, 0xdb, 0xf9,
	0x6c, 0xd5, 0xd4, 0xc9, 0xaf, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x43, 0xbb, 0x91, 0xb7, 0x10,
	0x02, 0x00, 0x00,
}

func (this *DateTime) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DateTime)
	if !ok {
		that2, ok := that.(DateTime)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Year != that1.Year {
		return false
	}
	if this.Month != that1.Month {
		return false
	}
	if this.Day != that1.Day {
		return false
	}
	if this.Hours != that1.Hours {
		return false
	}
	if this.Minutes != that1.Minutes {
		return false
	}
	if this.Seconds != that1.Seconds {
		return false
	}
	if this.Nanos != that1.Nanos {
		return false
	}
	if that1.TimeOffset == nil {
		if this.TimeOffset != nil {
			return false
		}
	} else if this.TimeOffset == nil {
		return false
	} else if !this.TimeOffset.Equal(that1.TimeOffset) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DateTime_UtcOffset) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DateTime_UtcOffset)
	if !ok {
		that2, ok := that.(DateTime_UtcOffset)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.UtcOffset.Equal(that1.UtcOffset) {
		return false
	}
	return true
}
func (this *DateTime_TimeZone) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DateTime_TimeZone)
	if !ok {
		that2, ok := that.(DateTime_TimeZone)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.TimeZone.Equal(that1.TimeZone) {
		return false
	}
	return true
}
func (this *TimeZone) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TimeZone)
	if !ok {
		that2, ok := that.(TimeZone)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.Version != that1.Version {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *DateTime) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 13)
	s = append(s, "&_type.DateTime{")
	s = append(s, "Year: "+fmt.Sprintf("%#v", this.Year)+",\n")
	s = append(s, "Month: "+fmt.Sprintf("%#v", this.Month)+",\n")
	s = append(s, "Day: "+fmt.Sprintf("%#v", this.Day)+",\n")
	s = append(s, "Hours: "+fmt.Sprintf("%#v", this.Hours)+",\n")
	s = append(s, "Minutes: "+fmt.Sprintf("%#v", this.Minutes)+",\n")
	s = append(s, "Seconds: "+fmt.Sprintf("%#v", this.Seconds)+",\n")
	s = append(s, "Nanos: "+fmt.Sprintf("%#v", this.Nanos)+",\n")
	if this.TimeOffset != nil {
		s = append(s, "TimeOffset: "+fmt.Sprintf("%#v", this.TimeOffset)+",\n")
	}
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func (this *DateTime_UtcOffset) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&_type.DateTime_UtcOffset{` +
		`UtcOffset:` + fmt.Sprintf("%#v", this.UtcOffset) + `}`}, ", ")
	return s
}
func (this *DateTime_TimeZone) GoString() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&_type.DateTime_TimeZone{` +
		`TimeZone:` + fmt.Sprintf("%#v", this.TimeZone) + `}`}, ", ")
	return s
}
func (this *TimeZone) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 6)
	s = append(s, "&_type.TimeZone{")
	s = append(s, "Id: "+fmt.Sprintf("%#v", this.Id)+",\n")
	s = append(s, "Version: "+fmt.Sprintf("%#v", this.Version)+",\n")
	if this.XXX_unrecognized != nil {
		s = append(s, "XXX_unrecognized:"+fmt.Sprintf("%#v", this.XXX_unrecognized)+",\n")
	}
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringDatetime(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *DateTime) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DateTime) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DateTime) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.TimeOffset != nil {
		{
			size := m.TimeOffset.Size()
			i -= size
			if _, err := m.TimeOffset.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	if m.Nanos != 0 {
		i = encodeVarintDatetime(dAtA, i, uint64(m.Nanos))
		i--
		dAtA[i] = 0x38
	}
	if m.Seconds != 0 {
		i = encodeVarintDatetime(dAtA, i, uint64(m.Seconds))
		i--
		dAtA[i] = 0x30
	}
	if m.Minutes != 0 {
		i = encodeVarintDatetime(dAtA, i, uint64(m.Minutes))
		i--
		dAtA[i] = 0x28
	}
	if m.Hours != 0 {
		i = encodeVarintDatetime(dAtA, i, uint64(m.Hours))
		i--
		dAtA[i] = 0x20
	}
	if m.Day != 0 {
		i = encodeVarintDatetime(dAtA, i, uint64(m.Day))
		i--
		dAtA[i] = 0x18
	}
	if m.Month != 0 {
		i = encodeVarintDatetime(dAtA, i, uint64(m.Month))
		i--
		dAtA[i] = 0x10
	}
	if m.Year != 0 {
		i = encodeVarintDatetime(dAtA, i, uint64(m.Year))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DateTime_UtcOffset) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *DateTime_UtcOffset) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.UtcOffset != nil {
		{
			size, err := m.UtcOffset.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDatetime(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x42
	}
	return len(dAtA) - i, nil
}
func (m *DateTime_TimeZone) MarshalTo(dAtA []byte) (int, error) {
	return m.MarshalToSizedBuffer(dAtA[:m.Size()])
}

func (m *DateTime_TimeZone) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.TimeZone != nil {
		{
			size, err := m.TimeZone.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDatetime(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x4a
	}
	return len(dAtA) - i, nil
}
func (m *TimeZone) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TimeZone) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TimeZone) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Version) > 0 {
		i -= len(m.Version)
		copy(dAtA[i:], m.Version)
		i = encodeVarintDatetime(dAtA, i, uint64(len(m.Version)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDatetime(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDatetime(dAtA []byte, offset int, v uint64) int {
	offset -= sovDatetime(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func NewPopulatedDateTime(r randyDatetime, easy bool) *DateTime {
	this := &DateTime{}
	this.Year = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Year *= -1
	}
	this.Month = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Month *= -1
	}
	this.Day = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Day *= -1
	}
	this.Hours = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Hours *= -1
	}
	this.Minutes = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Minutes *= -1
	}
	this.Seconds = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Seconds *= -1
	}
	this.Nanos = int32(r.Int31())
	if r.Intn(2) == 0 {
		this.Nanos *= -1
	}
	oneofNumber_TimeOffset := []int32{8, 9}[r.Intn(2)]
	switch oneofNumber_TimeOffset {
	case 8:
		this.TimeOffset = NewPopulatedDateTime_UtcOffset(r, easy)
	case 9:
		this.TimeOffset = NewPopulatedDateTime_TimeZone(r, easy)
	}
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedDatetime(r, 10)
	}
	return this
}

func NewPopulatedDateTime_UtcOffset(r randyDatetime, easy bool) *DateTime_UtcOffset {
	this := &DateTime_UtcOffset{}
	this.UtcOffset = types.NewPopulatedDuration(r, easy)
	return this
}
func NewPopulatedDateTime_TimeZone(r randyDatetime, easy bool) *DateTime_TimeZone {
	this := &DateTime_TimeZone{}
	this.TimeZone = NewPopulatedTimeZone(r, easy)
	return this
}
func NewPopulatedTimeZone(r randyDatetime, easy bool) *TimeZone {
	this := &TimeZone{}
	this.Id = string(randStringDatetime(r))
	this.Version = string(randStringDatetime(r))
	if !easy && r.Intn(10) != 0 {
		this.XXX_unrecognized = randUnrecognizedDatetime(r, 3)
	}
	return this
}

type randyDatetime interface {
	Float32() float32
	Float64() float64
	Int63() int64
	Int31() int32
	Uint32() uint32
	Intn(n int) int
}

func randUTF8RuneDatetime(r randyDatetime) rune {
	ru := r.Intn(62)
	if ru < 10 {
		return rune(ru + 48)
	} else if ru < 36 {
		return rune(ru + 55)
	}
	return rune(ru + 61)
}
func randStringDatetime(r randyDatetime) string {
	v1 := r.Intn(100)
	tmps := make([]rune, v1)
	for i := 0; i < v1; i++ {
		tmps[i] = randUTF8RuneDatetime(r)
	}
	return string(tmps)
}
func randUnrecognizedDatetime(r randyDatetime, maxFieldNumber int) (dAtA []byte) {
	l := r.Intn(5)
	for i := 0; i < l; i++ {
		wire := r.Intn(4)
		if wire == 3 {
			wire = 5
		}
		fieldNumber := maxFieldNumber + r.Intn(100)
		dAtA = randFieldDatetime(dAtA, r, fieldNumber, wire)
	}
	return dAtA
}
func randFieldDatetime(dAtA []byte, r randyDatetime, fieldNumber int, wire int) []byte {
	key := uint32(fieldNumber)<<3 | uint32(wire)
	switch wire {
	case 0:
		dAtA = encodeVarintPopulateDatetime(dAtA, uint64(key))
		v2 := r.Int63()
		if r.Intn(2) == 0 {
			v2 *= -1
		}
		dAtA = encodeVarintPopulateDatetime(dAtA, uint64(v2))
	case 1:
		dAtA = encodeVarintPopulateDatetime(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	case 2:
		dAtA = encodeVarintPopulateDatetime(dAtA, uint64(key))
		ll := r.Intn(100)
		dAtA = encodeVarintPopulateDatetime(dAtA, uint64(ll))
		for j := 0; j < ll; j++ {
			dAtA = append(dAtA, byte(r.Intn(256)))
		}
	default:
		dAtA = encodeVarintPopulateDatetime(dAtA, uint64(key))
		dAtA = append(dAtA, byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)), byte(r.Intn(256)))
	}
	return dAtA
}
func encodeVarintPopulateDatetime(dAtA []byte, v uint64) []byte {
	for v >= 1<<7 {
		dAtA = append(dAtA, uint8(uint64(v)&0x7f|0x80))
		v >>= 7
	}
	dAtA = append(dAtA, uint8(v))
	return dAtA
}
func (m *DateTime) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Year != 0 {
		n += 1 + sovDatetime(uint64(m.Year))
	}
	if m.Month != 0 {
		n += 1 + sovDatetime(uint64(m.Month))
	}
	if m.Day != 0 {
		n += 1 + sovDatetime(uint64(m.Day))
	}
	if m.Hours != 0 {
		n += 1 + sovDatetime(uint64(m.Hours))
	}
	if m.Minutes != 0 {
		n += 1 + sovDatetime(uint64(m.Minutes))
	}
	if m.Seconds != 0 {
		n += 1 + sovDatetime(uint64(m.Seconds))
	}
	if m.Nanos != 0 {
		n += 1 + sovDatetime(uint64(m.Nanos))
	}
	if m.TimeOffset != nil {
		n += m.TimeOffset.Size()
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *DateTime_UtcOffset) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.UtcOffset != nil {
		l = m.UtcOffset.Size()
		n += 1 + l + sovDatetime(uint64(l))
	}
	return n
}
func (m *DateTime_TimeZone) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TimeZone != nil {
		l = m.TimeZone.Size()
		n += 1 + l + sovDatetime(uint64(l))
	}
	return n
}
func (m *TimeZone) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDatetime(uint64(l))
	}
	l = len(m.Version)
	if l > 0 {
		n += 1 + l + sovDatetime(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovDatetime(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDatetime(x uint64) (n int) {
	return sovDatetime(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *DateTime) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DateTime{`,
		`Year:` + fmt.Sprintf("%v", this.Year) + `,`,
		`Month:` + fmt.Sprintf("%v", this.Month) + `,`,
		`Day:` + fmt.Sprintf("%v", this.Day) + `,`,
		`Hours:` + fmt.Sprintf("%v", this.Hours) + `,`,
		`Minutes:` + fmt.Sprintf("%v", this.Minutes) + `,`,
		`Seconds:` + fmt.Sprintf("%v", this.Seconds) + `,`,
		`Nanos:` + fmt.Sprintf("%v", this.Nanos) + `,`,
		`TimeOffset:` + fmt.Sprintf("%v", this.TimeOffset) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DateTime_UtcOffset) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DateTime_UtcOffset{`,
		`UtcOffset:` + strings.Replace(fmt.Sprintf("%v", this.UtcOffset), "Duration", "types.Duration", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *DateTime_TimeZone) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&DateTime_TimeZone{`,
		`TimeZone:` + strings.Replace(fmt.Sprintf("%v", this.TimeZone), "TimeZone", "TimeZone", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *TimeZone) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&TimeZone{`,
		`Id:` + fmt.Sprintf("%v", this.Id) + `,`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`XXX_unrecognized:` + fmt.Sprintf("%v", this.XXX_unrecognized) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringDatetime(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *DateTime) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDatetime
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
			return fmt.Errorf("proto: DateTime: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DateTime: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Year", wireType)
			}
			m.Year = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatetime
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Year |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Month", wireType)
			}
			m.Month = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatetime
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Month |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Day", wireType)
			}
			m.Day = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatetime
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Day |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hours", wireType)
			}
			m.Hours = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatetime
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Hours |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Minutes", wireType)
			}
			m.Minutes = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatetime
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Minutes |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seconds", wireType)
			}
			m.Seconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatetime
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seconds |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nanos", wireType)
			}
			m.Nanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatetime
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nanos |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UtcOffset", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatetime
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
				return ErrInvalidLengthDatetime
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDatetime
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &types.Duration{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.TimeOffset = &DateTime_UtcOffset{v}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeZone", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatetime
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
				return ErrInvalidLengthDatetime
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDatetime
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &TimeZone{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.TimeOffset = &DateTime_TimeZone{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDatetime(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDatetime
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDatetime
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TimeZone) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDatetime
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
			return fmt.Errorf("proto: TimeZone: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TimeZone: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatetime
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDatetime
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDatetime
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDatetime
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDatetime
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDatetime
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Version = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDatetime(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDatetime
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDatetime
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDatetime(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDatetime
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
					return 0, ErrIntOverflowDatetime
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
					return 0, ErrIntOverflowDatetime
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
				return 0, ErrInvalidLengthDatetime
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthDatetime
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowDatetime
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
				next, err := skipDatetime(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthDatetime
				}
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
	ErrInvalidLengthDatetime = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDatetime   = fmt.Errorf("proto: integer overflow")
)
