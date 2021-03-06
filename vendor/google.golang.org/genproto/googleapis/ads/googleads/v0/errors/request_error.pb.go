// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v0/errors/request_error.proto

package errors // import "google.golang.org/genproto/googleapis/ads/googleads/v0/errors"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enum describing possible request errors.
type RequestErrorEnum_RequestError int32

const (
	// Enum unspecified.
	RequestErrorEnum_UNSPECIFIED RequestErrorEnum_RequestError = 0
	// The received error code is not known in this version.
	RequestErrorEnum_UNKNOWN RequestErrorEnum_RequestError = 1
	// Resource name is required for this request.
	RequestErrorEnum_RESOURCE_NAME_MISSING RequestErrorEnum_RequestError = 3
	// Resource name provided is malformed.
	RequestErrorEnum_RESOURCE_NAME_MALFORMED RequestErrorEnum_RequestError = 4
	// Resource name provided is malformed.
	RequestErrorEnum_BAD_RESOURCE_ID RequestErrorEnum_RequestError = 17
	// Customer ID is invalid.
	RequestErrorEnum_INVALID_CUSTOMER_ID RequestErrorEnum_RequestError = 16
	// Mutate operation should have either create, update, or remove specified.
	RequestErrorEnum_OPERATION_REQUIRED RequestErrorEnum_RequestError = 5
	// Requested resource not found.
	RequestErrorEnum_RESOURCE_NOT_FOUND RequestErrorEnum_RequestError = 6
	// Next page token specified in user request is invalid.
	RequestErrorEnum_INVALID_PAGE_TOKEN RequestErrorEnum_RequestError = 7
	// Next page token specified in user request has expired.
	RequestErrorEnum_EXPIRED_PAGE_TOKEN RequestErrorEnum_RequestError = 8
	// Required field is missing.
	RequestErrorEnum_REQUIRED_FIELD_MISSING RequestErrorEnum_RequestError = 9
	// The field cannot be modified because it's immutable. It's also possible
	// that the field can be modified using 'create' operation but not 'update'.
	RequestErrorEnum_IMMUTABLE_FIELD RequestErrorEnum_RequestError = 11
	// Received too many entries in request.
	RequestErrorEnum_TOO_MANY_MUTATE_OPERATIONS RequestErrorEnum_RequestError = 13
	// Request cannot be executed by a manager account.
	RequestErrorEnum_CANNOT_BE_EXECUTED_BY_MANAGER_ACCOUNT RequestErrorEnum_RequestError = 14
	// Mutate request was attempting to modify a readonly field.
	// For instance, Budget fields can be requested for Ad Group,
	// but are read-only for adGroups:mutate.
	RequestErrorEnum_CANNOT_MODIFY_FOREIGN_FIELD RequestErrorEnum_RequestError = 15
	// Enum value is not permitted.
	RequestErrorEnum_INVALID_ENUM_VALUE RequestErrorEnum_RequestError = 18
	// The developer-token parameter is required for all requests.
	RequestErrorEnum_DEVELOPER_TOKEN_PARAMETER_MISSING RequestErrorEnum_RequestError = 19
	// The login-customer-id parameter is required for this request.
	RequestErrorEnum_LOGIN_CUSTOMER_ID_PARAMETER_MISSING RequestErrorEnum_RequestError = 20
)

var RequestErrorEnum_RequestError_name = map[int32]string{
	0:  "UNSPECIFIED",
	1:  "UNKNOWN",
	3:  "RESOURCE_NAME_MISSING",
	4:  "RESOURCE_NAME_MALFORMED",
	17: "BAD_RESOURCE_ID",
	16: "INVALID_CUSTOMER_ID",
	5:  "OPERATION_REQUIRED",
	6:  "RESOURCE_NOT_FOUND",
	7:  "INVALID_PAGE_TOKEN",
	8:  "EXPIRED_PAGE_TOKEN",
	9:  "REQUIRED_FIELD_MISSING",
	11: "IMMUTABLE_FIELD",
	13: "TOO_MANY_MUTATE_OPERATIONS",
	14: "CANNOT_BE_EXECUTED_BY_MANAGER_ACCOUNT",
	15: "CANNOT_MODIFY_FOREIGN_FIELD",
	18: "INVALID_ENUM_VALUE",
	19: "DEVELOPER_TOKEN_PARAMETER_MISSING",
	20: "LOGIN_CUSTOMER_ID_PARAMETER_MISSING",
}
var RequestErrorEnum_RequestError_value = map[string]int32{
	"UNSPECIFIED":                           0,
	"UNKNOWN":                               1,
	"RESOURCE_NAME_MISSING":                 3,
	"RESOURCE_NAME_MALFORMED":               4,
	"BAD_RESOURCE_ID":                       17,
	"INVALID_CUSTOMER_ID":                   16,
	"OPERATION_REQUIRED":                    5,
	"RESOURCE_NOT_FOUND":                    6,
	"INVALID_PAGE_TOKEN":                    7,
	"EXPIRED_PAGE_TOKEN":                    8,
	"REQUIRED_FIELD_MISSING":                9,
	"IMMUTABLE_FIELD":                       11,
	"TOO_MANY_MUTATE_OPERATIONS":            13,
	"CANNOT_BE_EXECUTED_BY_MANAGER_ACCOUNT": 14,
	"CANNOT_MODIFY_FOREIGN_FIELD":           15,
	"INVALID_ENUM_VALUE":                    18,
	"DEVELOPER_TOKEN_PARAMETER_MISSING":     19,
	"LOGIN_CUSTOMER_ID_PARAMETER_MISSING":   20,
}

func (x RequestErrorEnum_RequestError) String() string {
	return proto.EnumName(RequestErrorEnum_RequestError_name, int32(x))
}
func (RequestErrorEnum_RequestError) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_request_error_7fde1cab6fa86182, []int{0, 0}
}

// Container for enum describing possible request errors.
type RequestErrorEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestErrorEnum) Reset()         { *m = RequestErrorEnum{} }
func (m *RequestErrorEnum) String() string { return proto.CompactTextString(m) }
func (*RequestErrorEnum) ProtoMessage()    {}
func (*RequestErrorEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_request_error_7fde1cab6fa86182, []int{0}
}
func (m *RequestErrorEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestErrorEnum.Unmarshal(m, b)
}
func (m *RequestErrorEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestErrorEnum.Marshal(b, m, deterministic)
}
func (dst *RequestErrorEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestErrorEnum.Merge(dst, src)
}
func (m *RequestErrorEnum) XXX_Size() int {
	return xxx_messageInfo_RequestErrorEnum.Size(m)
}
func (m *RequestErrorEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestErrorEnum.DiscardUnknown(m)
}

var xxx_messageInfo_RequestErrorEnum proto.InternalMessageInfo

func init() {
	proto.RegisterType((*RequestErrorEnum)(nil), "google.ads.googleads.v0.errors.RequestErrorEnum")
	proto.RegisterEnum("google.ads.googleads.v0.errors.RequestErrorEnum_RequestError", RequestErrorEnum_RequestError_name, RequestErrorEnum_RequestError_value)
}

func init() {
	proto.RegisterFile("google/ads/googleads/v0/errors/request_error.proto", fileDescriptor_request_error_7fde1cab6fa86182)
}

var fileDescriptor_request_error_7fde1cab6fa86182 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x52, 0x51, 0x6f, 0xd3, 0x3c,
	0x14, 0xfd, 0xba, 0xee, 0xdb, 0xc0, 0x05, 0xea, 0xb9, 0xb0, 0x89, 0x4d, 0x2a, 0xa2, 0x68, 0x42,
	0xbc, 0xa4, 0x15, 0x3c, 0xf2, 0xe4, 0xc4, 0x37, 0x91, 0xb5, 0xc4, 0x0e, 0x4e, 0x5c, 0x56, 0x54,
	0xc9, 0x2a, 0xb4, 0x8a, 0x90, 0xb6, 0x66, 0x24, 0xdb, 0x7e, 0x05, 0xbf, 0x82, 0x47, 0xfe, 0x08,
	0x12, 0x7f, 0x88, 0x57, 0xe4, 0xba, 0x0d, 0x15, 0x02, 0x9e, 0x72, 0x73, 0xee, 0x39, 0xf7, 0xdc,
	0x6b, 0x1d, 0xf4, 0xb2, 0x28, 0xcb, 0xe2, 0x62, 0x31, 0x9c, 0xcd, 0xeb, 0xa1, 0x2b, 0x6d, 0x75,
	0x3b, 0x1a, 0x2e, 0xaa, 0xaa, 0xac, 0xea, 0x61, 0xb5, 0xf8, 0x74, 0xb3, 0xa8, 0xaf, 0xcd, 0xea,
	0xd7, 0xbb, 0xaa, 0xca, 0xeb, 0x92, 0xf4, 0x1d, 0xd1, 0x9b, 0xcd, 0x6b, 0xaf, 0xd1, 0x78, 0xb7,
	0x23, 0xcf, 0x69, 0x06, 0x9f, 0x77, 0x11, 0x56, 0x4e, 0x07, 0x16, 0x81, 0xe5, 0xcd, 0xe5, 0xe0,
	0x47, 0x1b, 0xdd, 0xdb, 0x06, 0x49, 0x17, 0x75, 0xb4, 0xc8, 0x52, 0x08, 0x78, 0xc8, 0x81, 0xe1,
	0xff, 0x48, 0x07, 0xed, 0x6b, 0x71, 0x26, 0xe4, 0x5b, 0x81, 0x5b, 0xe4, 0x31, 0x7a, 0xa4, 0x20,
	0x93, 0x5a, 0x05, 0x60, 0x04, 0x4d, 0xc0, 0x24, 0x3c, 0xcb, 0xb8, 0x88, 0x70, 0x9b, 0x9c, 0xa0,
	0xa3, 0xdf, 0x5a, 0x34, 0x0e, 0xa5, 0x4a, 0x80, 0xe1, 0x5d, 0xd2, 0x43, 0x5d, 0x9f, 0x32, 0xd3,
	0x10, 0x38, 0xc3, 0x07, 0xe4, 0x08, 0xf5, 0xb8, 0x18, 0xd3, 0x98, 0x33, 0x13, 0xe8, 0x2c, 0x97,
	0x09, 0x28, 0xdb, 0xc0, 0xe4, 0x10, 0x11, 0x99, 0x82, 0xa2, 0x39, 0x97, 0xc2, 0x28, 0x78, 0xa3,
	0xb9, 0x02, 0x86, 0xff, 0xb7, 0xf8, 0x2f, 0x0b, 0x99, 0x9b, 0x50, 0x6a, 0xc1, 0xf0, 0x9e, 0xc5,
	0x37, 0x83, 0x52, 0x1a, 0x81, 0xc9, 0xe5, 0x19, 0x08, 0xbc, 0x6f, 0x71, 0x38, 0x4f, 0xad, 0x78,
	0x1b, 0xbf, 0x43, 0x8e, 0xd1, 0xe1, 0x66, 0xaa, 0x09, 0x39, 0xc4, 0xac, 0x39, 0xe3, 0xae, 0xdd,
	0x94, 0x27, 0x89, 0xce, 0xa9, 0x1f, 0x83, 0x6b, 0xe2, 0x0e, 0xe9, 0xa3, 0xe3, 0x5c, 0x4a, 0x93,
	0x50, 0x31, 0x31, 0xb6, 0x97, 0x83, 0x69, 0x16, 0xcc, 0xf0, 0x7d, 0xf2, 0x02, 0x9d, 0x06, 0x54,
	0xd8, 0x95, 0x7c, 0x30, 0x70, 0x0e, 0x81, 0xce, 0x81, 0x19, 0x7f, 0x62, 0x15, 0x34, 0x02, 0x65,
	0x68, 0x10, 0x48, 0x2d, 0x72, 0xfc, 0x80, 0x3c, 0x41, 0x27, 0x6b, 0x6a, 0x22, 0x19, 0x0f, 0x27,
	0x26, 0x94, 0x0a, 0x78, 0x24, 0xd6, 0x5e, 0xdd, 0xed, 0x63, 0x40, 0xe8, 0xc4, 0x8c, 0x69, 0xac,
	0x01, 0x13, 0x72, 0x8a, 0x9e, 0x32, 0x18, 0x43, 0x6c, 0x8d, 0xdd, 0x25, 0x26, 0xa5, 0x8a, 0x26,
	0x90, 0x83, 0x6a, 0xf6, 0xef, 0x91, 0xe7, 0xe8, 0x59, 0x2c, 0x23, 0x2e, 0xb6, 0x9f, 0xf4, 0x0f,
	0xc4, 0x87, 0xfe, 0xb7, 0x16, 0x1a, 0x7c, 0x28, 0x2f, 0xbd, 0x7f, 0xa7, 0xc6, 0x3f, 0xd8, 0x4e,
	0x47, 0x6a, 0x83, 0x96, 0xb6, 0xde, 0xb1, 0xb5, 0xa8, 0x28, 0x2f, 0x66, 0xcb, 0xc2, 0x2b, 0xab,
	0x62, 0x58, 0x2c, 0x96, 0xab, 0x18, 0x6e, 0xe2, 0x7a, 0xf5, 0xb1, 0xfe, 0x5b, 0x7a, 0x5f, 0xbb,
	0xcf, 0x97, 0x9d, 0x76, 0x44, 0xe9, 0xd7, 0x9d, 0x7e, 0xe4, 0x86, 0xd1, 0x79, 0xed, 0xb9, 0xd2,
	0x56, 0xe3, 0x91, 0xb7, 0xb2, 0xac, 0xbf, 0x6f, 0x08, 0x53, 0x3a, 0xaf, 0xa7, 0x0d, 0x61, 0x3a,
	0x1e, 0x4d, 0x1d, 0xe1, 0xfd, 0xde, 0xca, 0xf8, 0xd5, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x58,
	0xab, 0xbb, 0xa7, 0x35, 0x03, 0x00, 0x00,
}
