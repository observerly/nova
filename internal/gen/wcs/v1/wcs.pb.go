//***************************************************************************************************************

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/nova/proto/solve/v1
//	@license	Copyright © 2021-2025 observerly

//***************************************************************************************************************

// protolint:disable

//***************************************************************************************************************

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: wcs/v1/wcs.proto

//***************************************************************************************************************

package wcsv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// WCS represents World Coordinate System parameters.
type WCS struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Number of world coordinate axes
	WCAXES int32 `protobuf:"varint,1,opt,name=WCAXES,proto3" json:"WCAXES,omitempty"` // Default: 2
	// Reference pixel X
	CRPIX1 float64 `protobuf:"fixed64,2,opt,name=CRPIX1,proto3" json:"CRPIX1,omitempty"`
	// Reference pixel Y
	CRPIX2 float64 `protobuf:"fixed64,3,opt,name=CRPIX2,proto3" json:"CRPIX2,omitempty"`
	// Reference RA (Default: 0.0)
	CRVAL1 float64 `protobuf:"fixed64,4,opt,name=CRVAL1,proto3" json:"CRVAL1,omitempty"`
	// Reference Dec (Default: 0.0)
	CRVAL2 float64 `protobuf:"fixed64,5,opt,name=CRVAL2,proto3" json:"CRVAL2,omitempty"`
	// Coordinate type for axis 1, typically RA with TAN projection (Default: "RA---TAN")
	CTYPE1 string `protobuf:"bytes,6,opt,name=CTYPE1,proto3" json:"CTYPE1,omitempty"`
	// Coordinate type for axis 2, typically DEC with TAN projection (Default: "DEC--TAN")
	CTYPE2 string `protobuf:"bytes,7,opt,name=CTYPE2,proto3" json:"CTYPE2,omitempty"`
	// Coordinate increment for axis 1
	CDELT1 float64 `protobuf:"fixed64,8,opt,name=CDELT1,proto3" json:"CDELT1,omitempty"`
	// Coordinate increment for axis 2
	CDELT2 float64 `protobuf:"fixed64,9,opt,name=CDELT2,proto3" json:"CDELT2,omitempty"`
	// Coordinate unit for axis 1 (Default: "deg")
	CUNIT1 string `protobuf:"bytes,10,opt,name=CUNIT1,proto3" json:"CUNIT1,omitempty"`
	// Coordinate unit for axis 2 (Default: "deg")
	CUNIT2 string `protobuf:"bytes,11,opt,name=CUNIT2,proto3" json:"CUNIT2,omitempty"`
	// Affine transform parameter A
	CD1_1 float64 `protobuf:"fixed64,12,opt,name=CD1_1,json=CD11,proto3" json:"CD1_1,omitempty"`
	// Affine transform parameter B
	CD1_2 float64 `protobuf:"fixed64,13,opt,name=CD1_2,json=CD12,proto3" json:"CD1_2,omitempty"`
	// Affine transform parameter C
	CD2_1 float64 `protobuf:"fixed64,14,opt,name=CD2_1,json=CD21,proto3" json:"CD2_1,omitempty"`
	// Affine transform parameter D
	CD2_2 float64 `protobuf:"fixed64,15,opt,name=CD2_2,json=CD22,proto3" json:"CD2_2,omitempty"`
	// Affine translation parameter e (optional)
	E float64 `protobuf:"fixed64,16,opt,name=E,proto3" json:"E,omitempty"`
	// Affine translation parameter f (optional)
	F float64 `protobuf:"fixed64,17,opt,name=F,proto3" json:"F,omitempty"`
	// SIP forward transformation (distortion) coefficients
	FSIP *SIP2DForwardParameters `protobuf:"bytes,18,opt,name=FSIP,proto3" json:"FSIP,omitempty"`
	// SIP inverse transformation (distortion) coefficients
	ISIP          *SIP2DInverseParameters `protobuf:"bytes,19,opt,name=ISIP,proto3" json:"ISIP,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WCS) Reset() {
	*x = WCS{}
	mi := &file_wcs_v1_wcs_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WCS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WCS) ProtoMessage() {}

func (x *WCS) ProtoReflect() protoreflect.Message {
	mi := &file_wcs_v1_wcs_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WCS.ProtoReflect.Descriptor instead.
func (*WCS) Descriptor() ([]byte, []int) {
	return file_wcs_v1_wcs_proto_rawDescGZIP(), []int{0}
}

func (x *WCS) GetWCAXES() int32 {
	if x != nil {
		return x.WCAXES
	}
	return 0
}

func (x *WCS) GetCRPIX1() float64 {
	if x != nil {
		return x.CRPIX1
	}
	return 0
}

func (x *WCS) GetCRPIX2() float64 {
	if x != nil {
		return x.CRPIX2
	}
	return 0
}

func (x *WCS) GetCRVAL1() float64 {
	if x != nil {
		return x.CRVAL1
	}
	return 0
}

func (x *WCS) GetCRVAL2() float64 {
	if x != nil {
		return x.CRVAL2
	}
	return 0
}

func (x *WCS) GetCTYPE1() string {
	if x != nil {
		return x.CTYPE1
	}
	return ""
}

func (x *WCS) GetCTYPE2() string {
	if x != nil {
		return x.CTYPE2
	}
	return ""
}

func (x *WCS) GetCDELT1() float64 {
	if x != nil {
		return x.CDELT1
	}
	return 0
}

func (x *WCS) GetCDELT2() float64 {
	if x != nil {
		return x.CDELT2
	}
	return 0
}

func (x *WCS) GetCUNIT1() string {
	if x != nil {
		return x.CUNIT1
	}
	return ""
}

func (x *WCS) GetCUNIT2() string {
	if x != nil {
		return x.CUNIT2
	}
	return ""
}

func (x *WCS) GetCD1_1() float64 {
	if x != nil {
		return x.CD1_1
	}
	return 0
}

func (x *WCS) GetCD1_2() float64 {
	if x != nil {
		return x.CD1_2
	}
	return 0
}

func (x *WCS) GetCD2_1() float64 {
	if x != nil {
		return x.CD2_1
	}
	return 0
}

func (x *WCS) GetCD2_2() float64 {
	if x != nil {
		return x.CD2_2
	}
	return 0
}

func (x *WCS) GetE() float64 {
	if x != nil {
		return x.E
	}
	return 0
}

func (x *WCS) GetF() float64 {
	if x != nil {
		return x.F
	}
	return 0
}

func (x *WCS) GetFSIP() *SIP2DForwardParameters {
	if x != nil {
		return x.FSIP
	}
	return nil
}

func (x *WCS) GetISIP() *SIP2DInverseParameters {
	if x != nil {
		return x.ISIP
	}
	return nil
}

// SIP2DForwardParameters contains polynomial coefficients for forward SIP transformation.
type SIP2DForwardParameters struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Order of the A polynomial
	AOrder int32 `protobuf:"varint,1,opt,name=AOrder,proto3" json:"AOrder,omitempty"`
	// Polynomial coefficients for A
	APower map[string]float64 `protobuf:"bytes,2,rep,name=APower,proto3" json:"APower,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	// Order of the B polynomial
	BOrder int32 `protobuf:"varint,3,opt,name=BOrder,proto3" json:"BOrder,omitempty"`
	// Polynomial coefficients for B
	BPower        map[string]float64 `protobuf:"bytes,4,rep,name=BPower,proto3" json:"BPower,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SIP2DForwardParameters) Reset() {
	*x = SIP2DForwardParameters{}
	mi := &file_wcs_v1_wcs_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SIP2DForwardParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SIP2DForwardParameters) ProtoMessage() {}

func (x *SIP2DForwardParameters) ProtoReflect() protoreflect.Message {
	mi := &file_wcs_v1_wcs_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SIP2DForwardParameters.ProtoReflect.Descriptor instead.
func (*SIP2DForwardParameters) Descriptor() ([]byte, []int) {
	return file_wcs_v1_wcs_proto_rawDescGZIP(), []int{1}
}

func (x *SIP2DForwardParameters) GetAOrder() int32 {
	if x != nil {
		return x.AOrder
	}
	return 0
}

func (x *SIP2DForwardParameters) GetAPower() map[string]float64 {
	if x != nil {
		return x.APower
	}
	return nil
}

func (x *SIP2DForwardParameters) GetBOrder() int32 {
	if x != nil {
		return x.BOrder
	}
	return 0
}

func (x *SIP2DForwardParameters) GetBPower() map[string]float64 {
	if x != nil {
		return x.BPower
	}
	return nil
}

// SIP2DInverseParameters contains polynomial coefficients for inverse SIP transformation.
type SIP2DInverseParameters struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Order of the A' polynomial
	APOrder int32 `protobuf:"varint,1,opt,name=APOrder,proto3" json:"APOrder,omitempty"`
	// Polynomial coefficients for A'
	APPower map[string]float64 `protobuf:"bytes,2,rep,name=APPower,proto3" json:"APPower,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	// Order of the B' polynomial
	BPOrder int32 `protobuf:"varint,3,opt,name=BPOrder,proto3" json:"BPOrder,omitempty"`
	// Polynomial coefficients for B'
	BPPower       map[string]float64 `protobuf:"bytes,4,rep,name=BPPower,proto3" json:"BPPower,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"fixed64,2,opt,name=value"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SIP2DInverseParameters) Reset() {
	*x = SIP2DInverseParameters{}
	mi := &file_wcs_v1_wcs_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SIP2DInverseParameters) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SIP2DInverseParameters) ProtoMessage() {}

func (x *SIP2DInverseParameters) ProtoReflect() protoreflect.Message {
	mi := &file_wcs_v1_wcs_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SIP2DInverseParameters.ProtoReflect.Descriptor instead.
func (*SIP2DInverseParameters) Descriptor() ([]byte, []int) {
	return file_wcs_v1_wcs_proto_rawDescGZIP(), []int{2}
}

func (x *SIP2DInverseParameters) GetAPOrder() int32 {
	if x != nil {
		return x.APOrder
	}
	return 0
}

func (x *SIP2DInverseParameters) GetAPPower() map[string]float64 {
	if x != nil {
		return x.APPower
	}
	return nil
}

func (x *SIP2DInverseParameters) GetBPOrder() int32 {
	if x != nil {
		return x.BPOrder
	}
	return 0
}

func (x *SIP2DInverseParameters) GetBPPower() map[string]float64 {
	if x != nil {
		return x.BPPower
	}
	return nil
}

var File_wcs_v1_wcs_proto protoreflect.FileDescriptor

var file_wcs_v1_wcs_proto_rawDesc = string([]byte{
	0x0a, 0x10, 0x77, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x77, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x22, 0xe5, 0x03, 0x0a, 0x03, 0x57,
	0x43, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x57, 0x43, 0x41, 0x58, 0x45, 0x53, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x57, 0x43, 0x41, 0x58, 0x45, 0x53, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x52,
	0x50, 0x49, 0x58, 0x31, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x43, 0x52, 0x50, 0x49,
	0x58, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x52, 0x50, 0x49, 0x58, 0x32, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x43, 0x52, 0x50, 0x49, 0x58, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x52,
	0x56, 0x41, 0x4c, 0x31, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x43, 0x52, 0x56, 0x41,
	0x4c, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x52, 0x56, 0x41, 0x4c, 0x32, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x43, 0x52, 0x56, 0x41, 0x4c, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x54,
	0x59, 0x50, 0x45, 0x31, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x54, 0x59, 0x50,
	0x45, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x54, 0x59, 0x50, 0x45, 0x32, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x43, 0x54, 0x59, 0x50, 0x45, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x44,
	0x45, 0x4c, 0x54, 0x31, 0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x43, 0x44, 0x45, 0x4c,
	0x54, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x44, 0x45, 0x4c, 0x54, 0x32, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x01, 0x52, 0x06, 0x43, 0x44, 0x45, 0x4c, 0x54, 0x32, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x55,
	0x4e, 0x49, 0x54, 0x31, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x43, 0x55, 0x4e, 0x49,
	0x54, 0x31, 0x12, 0x16, 0x0a, 0x06, 0x43, 0x55, 0x4e, 0x49, 0x54, 0x32, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x43, 0x55, 0x4e, 0x49, 0x54, 0x32, 0x12, 0x13, 0x0a, 0x05, 0x43, 0x44,
	0x31, 0x5f, 0x31, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x43, 0x44, 0x31, 0x31, 0x12,
	0x13, 0x0a, 0x05, 0x43, 0x44, 0x31, 0x5f, 0x32, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x43, 0x44, 0x31, 0x32, 0x12, 0x13, 0x0a, 0x05, 0x43, 0x44, 0x32, 0x5f, 0x31, 0x18, 0x0e, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x04, 0x43, 0x44, 0x32, 0x31, 0x12, 0x13, 0x0a, 0x05, 0x43, 0x44, 0x32,
	0x5f, 0x32, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x43, 0x44, 0x32, 0x32, 0x12, 0x0c,
	0x0a, 0x01, 0x45, 0x18, 0x10, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x45, 0x12, 0x0c, 0x0a, 0x01,
	0x46, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x01, 0x46, 0x12, 0x32, 0x0a, 0x04, 0x46, 0x53,
	0x49, 0x50, 0x18, 0x12, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x77, 0x63, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x53, 0x49, 0x50, 0x32, 0x44, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x04, 0x46, 0x53, 0x49, 0x50, 0x12, 0x32,
	0x0a, 0x04, 0x49, 0x53, 0x49, 0x50, 0x18, 0x13, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x77,
	0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x49, 0x50, 0x32, 0x44, 0x49, 0x6e, 0x76, 0x65, 0x72,
	0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x52, 0x04, 0x49, 0x53,
	0x49, 0x50, 0x22, 0xc6, 0x02, 0x0a, 0x16, 0x53, 0x49, 0x50, 0x32, 0x44, 0x46, 0x6f, 0x72, 0x77,
	0x61, 0x72, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x41, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x41,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x42, 0x0a, 0x06, 0x41, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x77, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53,
	0x49, 0x50, 0x32, 0x44, 0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d,
	0x65, 0x74, 0x65, 0x72, 0x73, 0x2e, 0x41, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x06, 0x41, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x42, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x42, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x42, 0x0a, 0x06, 0x42, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2a, 0x2e, 0x77, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x49, 0x50, 0x32, 0x44,
	0x46, 0x6f, 0x72, 0x77, 0x61, 0x72, 0x64, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x42, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x06, 0x42,
	0x50, 0x6f, 0x77, 0x65, 0x72, 0x1a, 0x39, 0x0a, 0x0b, 0x41, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x1a, 0x39, 0x0a, 0x0b, 0x42, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12,
	0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xd2, 0x02, 0x0a, 0x16,
	0x53, 0x49, 0x50, 0x32, 0x44, 0x49, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61,
	0x6d, 0x65, 0x74, 0x65, 0x72, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x41, 0x50, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x41, 0x50, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x12, 0x45, 0x0a, 0x07, 0x41, 0x50, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x2b, 0x2e, 0x77, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x49, 0x50, 0x32, 0x44,
	0x49, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65, 0x72,
	0x73, 0x2e, 0x41, 0x50, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x41, 0x50, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x42, 0x50, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x42, 0x50, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x45, 0x0a, 0x07, 0x42, 0x50, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x18, 0x04, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x77, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x49, 0x50, 0x32,
	0x44, 0x49, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x65, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x65, 0x74, 0x65,
	0x72, 0x73, 0x2e, 0x42, 0x50, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x07, 0x42, 0x50, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x1a, 0x3a, 0x0a, 0x0c, 0x41, 0x50, 0x50, 0x6f,
	0x77, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x3a, 0x02, 0x38, 0x01, 0x1a, 0x3a, 0x0a, 0x0c, 0x42, 0x50, 0x50, 0x6f, 0x77, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x42, 0x6f, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x2e, 0x77, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x08,
	0x57, 0x63, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1e, 0x6e, 0x6f, 0x76, 0x61,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x77, 0x63,
	0x73, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x63, 0x73, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x57, 0x58, 0x58,
	0xaa, 0x02, 0x06, 0x57, 0x63, 0x73, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x06, 0x57, 0x63, 0x73, 0x5c,
	0x56, 0x31, 0xe2, 0x02, 0x12, 0x57, 0x63, 0x73, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x57, 0x63, 0x73, 0x3a, 0x3a, 0x56,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_wcs_v1_wcs_proto_rawDescOnce sync.Once
	file_wcs_v1_wcs_proto_rawDescData []byte
)

func file_wcs_v1_wcs_proto_rawDescGZIP() []byte {
	file_wcs_v1_wcs_proto_rawDescOnce.Do(func() {
		file_wcs_v1_wcs_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_wcs_v1_wcs_proto_rawDesc), len(file_wcs_v1_wcs_proto_rawDesc)))
	})
	return file_wcs_v1_wcs_proto_rawDescData
}

var file_wcs_v1_wcs_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_wcs_v1_wcs_proto_goTypes = []any{
	(*WCS)(nil),                    // 0: wcs.v1.WCS
	(*SIP2DForwardParameters)(nil), // 1: wcs.v1.SIP2DForwardParameters
	(*SIP2DInverseParameters)(nil), // 2: wcs.v1.SIP2DInverseParameters
	nil,                            // 3: wcs.v1.SIP2DForwardParameters.APowerEntry
	nil,                            // 4: wcs.v1.SIP2DForwardParameters.BPowerEntry
	nil,                            // 5: wcs.v1.SIP2DInverseParameters.APPowerEntry
	nil,                            // 6: wcs.v1.SIP2DInverseParameters.BPPowerEntry
}
var file_wcs_v1_wcs_proto_depIdxs = []int32{
	1, // 0: wcs.v1.WCS.FSIP:type_name -> wcs.v1.SIP2DForwardParameters
	2, // 1: wcs.v1.WCS.ISIP:type_name -> wcs.v1.SIP2DInverseParameters
	3, // 2: wcs.v1.SIP2DForwardParameters.APower:type_name -> wcs.v1.SIP2DForwardParameters.APowerEntry
	4, // 3: wcs.v1.SIP2DForwardParameters.BPower:type_name -> wcs.v1.SIP2DForwardParameters.BPowerEntry
	5, // 4: wcs.v1.SIP2DInverseParameters.APPower:type_name -> wcs.v1.SIP2DInverseParameters.APPowerEntry
	6, // 5: wcs.v1.SIP2DInverseParameters.BPPower:type_name -> wcs.v1.SIP2DInverseParameters.BPPowerEntry
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_wcs_v1_wcs_proto_init() }
func file_wcs_v1_wcs_proto_init() {
	if File_wcs_v1_wcs_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_wcs_v1_wcs_proto_rawDesc), len(file_wcs_v1_wcs_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_wcs_v1_wcs_proto_goTypes,
		DependencyIndexes: file_wcs_v1_wcs_proto_depIdxs,
		MessageInfos:      file_wcs_v1_wcs_proto_msgTypes,
	}.Build()
	File_wcs_v1_wcs_proto = out.File
	file_wcs_v1_wcs_proto_goTypes = nil
	file_wcs_v1_wcs_proto_depIdxs = nil
}
