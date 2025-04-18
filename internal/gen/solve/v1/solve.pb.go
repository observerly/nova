//***************************************************************************************************************

//	@author		Michael Roberts <michael@observerly.com>
//	@package	@observerly/nova/proto/solve/v1
//	@license	Copyright © 2021-2025 observerly

//***************************************************************************************************************

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        (unknown)
// source: solve/v1/solve.proto

//***************************************************************************************************************

package solvev1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	v1 "nova/internal/gen/wcs/v1"
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

type SolveForWCSFITSHandlerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Owner         string                 `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	BucketName    string                 `protobuf:"bytes,2,opt,name=bucket_name,json=bucketName,proto3" json:"bucket_name,omitempty"`
	Location      string                 `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Ra            *float64               `protobuf:"fixed64,4,opt,name=ra,proto3,oneof" json:"ra,omitempty"`
	Dec           *float64               `protobuf:"fixed64,5,opt,name=dec,proto3,oneof" json:"dec,omitempty"`
	PixelScaleX   *float64               `protobuf:"fixed64,6,opt,name=pixel_scale_x,json=pixelScaleX,proto3,oneof" json:"pixel_scale_x,omitempty"`
	PixelScaleY   *float64               `protobuf:"fixed64,7,opt,name=pixel_scale_y,json=pixelScaleY,proto3,oneof" json:"pixel_scale_y,omitempty"`
	Radius        *float64               `protobuf:"fixed64,8,opt,name=radius,proto3,oneof" json:"radius,omitempty"`
	Sigma         *float64               `protobuf:"fixed64,9,opt,name=sigma,proto3,oneof" json:"sigma,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SolveForWCSFITSHandlerRequest) Reset() {
	*x = SolveForWCSFITSHandlerRequest{}
	mi := &file_solve_v1_solve_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SolveForWCSFITSHandlerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolveForWCSFITSHandlerRequest) ProtoMessage() {}

func (x *SolveForWCSFITSHandlerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_solve_v1_solve_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolveForWCSFITSHandlerRequest.ProtoReflect.Descriptor instead.
func (*SolveForWCSFITSHandlerRequest) Descriptor() ([]byte, []int) {
	return file_solve_v1_solve_proto_rawDescGZIP(), []int{0}
}

func (x *SolveForWCSFITSHandlerRequest) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *SolveForWCSFITSHandlerRequest) GetBucketName() string {
	if x != nil {
		return x.BucketName
	}
	return ""
}

func (x *SolveForWCSFITSHandlerRequest) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *SolveForWCSFITSHandlerRequest) GetRa() float64 {
	if x != nil && x.Ra != nil {
		return *x.Ra
	}
	return 0
}

func (x *SolveForWCSFITSHandlerRequest) GetDec() float64 {
	if x != nil && x.Dec != nil {
		return *x.Dec
	}
	return 0
}

func (x *SolveForWCSFITSHandlerRequest) GetPixelScaleX() float64 {
	if x != nil && x.PixelScaleX != nil {
		return *x.PixelScaleX
	}
	return 0
}

func (x *SolveForWCSFITSHandlerRequest) GetPixelScaleY() float64 {
	if x != nil && x.PixelScaleY != nil {
		return *x.PixelScaleY
	}
	return 0
}

func (x *SolveForWCSFITSHandlerRequest) GetRadius() float64 {
	if x != nil && x.Radius != nil {
		return *x.Radius
	}
	return 0
}

func (x *SolveForWCSFITSHandlerRequest) GetSigma() float64 {
	if x != nil && x.Sigma != nil {
		return *x.Sigma
	}
	return 0
}

type SolveForWCSFITSHandlerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Wcs           *v1.WCS                `protobuf:"bytes,1,opt,name=wcs,proto3" json:"wcs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SolveForWCSFITSHandlerResponse) Reset() {
	*x = SolveForWCSFITSHandlerResponse{}
	mi := &file_solve_v1_solve_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SolveForWCSFITSHandlerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolveForWCSFITSHandlerResponse) ProtoMessage() {}

func (x *SolveForWCSFITSHandlerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_solve_v1_solve_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolveForWCSFITSHandlerResponse.ProtoReflect.Descriptor instead.
func (*SolveForWCSFITSHandlerResponse) Descriptor() ([]byte, []int) {
	return file_solve_v1_solve_proto_rawDescGZIP(), []int{1}
}

func (x *SolveForWCSFITSHandlerResponse) GetWcs() *v1.WCS {
	if x != nil {
		return x.Wcs
	}
	return nil
}

var File_solve_v1_solve_proto protoreflect.FileDescriptor

var file_solve_v1_solve_proto_rawDesc = string([]byte{
	0x0a, 0x14, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x2e, 0x76, 0x31,
	0x1a, 0x10, 0x77, 0x63, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0xf0, 0x02, 0x0a, 0x1d, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x46, 0x6f, 0x72, 0x57,
	0x43, 0x53, 0x46, 0x49, 0x54, 0x53, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x0b, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x62, 0x75, 0x63, 0x6b, 0x65, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x02, 0x72, 0x61, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x02, 0x72, 0x61, 0x88, 0x01, 0x01, 0x12, 0x15, 0x0a, 0x03,
	0x64, 0x65, 0x63, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x48, 0x01, 0x52, 0x03, 0x64, 0x65, 0x63,
	0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0d, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x5f, 0x73, 0x63, 0x61,
	0x6c, 0x65, 0x5f, 0x78, 0x18, 0x06, 0x20, 0x01, 0x28, 0x01, 0x48, 0x02, 0x52, 0x0b, 0x70, 0x69,
	0x78, 0x65, 0x6c, 0x53, 0x63, 0x61, 0x6c, 0x65, 0x58, 0x88, 0x01, 0x01, 0x12, 0x27, 0x0a, 0x0d,
	0x70, 0x69, 0x78, 0x65, 0x6c, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x5f, 0x79, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x03, 0x52, 0x0b, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x53, 0x63, 0x61, 0x6c,
	0x65, 0x59, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x01, 0x48, 0x04, 0x52, 0x06, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x88,
	0x01, 0x01, 0x12, 0x19, 0x0a, 0x05, 0x73, 0x69, 0x67, 0x6d, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x05, 0x52, 0x05, 0x73, 0x69, 0x67, 0x6d, 0x61, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a,
	0x03, 0x5f, 0x72, 0x61, 0x42, 0x06, 0x0a, 0x04, 0x5f, 0x64, 0x65, 0x63, 0x42, 0x10, 0x0a, 0x0e,
	0x5f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x5f, 0x78, 0x42, 0x10,
	0x0a, 0x0e, 0x5f, 0x70, 0x69, 0x78, 0x65, 0x6c, 0x5f, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x5f, 0x79,
	0x42, 0x09, 0x0a, 0x07, 0x5f, 0x72, 0x61, 0x64, 0x69, 0x75, 0x73, 0x42, 0x08, 0x0a, 0x06, 0x5f,
	0x73, 0x69, 0x67, 0x6d, 0x61, 0x22, 0x3f, 0x0a, 0x1e, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x46, 0x6f,
	0x72, 0x57, 0x43, 0x53, 0x46, 0x49, 0x54, 0x53, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x03, 0x77, 0x63, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x77, 0x63, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x43,
	0x53, 0x52, 0x03, 0x77, 0x63, 0x73, 0x32, 0x7d, 0x0a, 0x0c, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x16, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x46,
	0x6f, 0x72, 0x57, 0x43, 0x53, 0x46, 0x49, 0x54, 0x53, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72,
	0x12, 0x27, 0x2e, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x6c, 0x76,
	0x65, 0x46, 0x6f, 0x72, 0x57, 0x43, 0x53, 0x46, 0x49, 0x54, 0x53, 0x48, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x28, 0x2e, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x46, 0x6f, 0x72, 0x57, 0x43, 0x53,
	0x46, 0x49, 0x54, 0x53, 0x48, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x7f, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x6f, 0x6c,
	0x76, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x22, 0x6e, 0x6f, 0x76, 0x61, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x2f, 0x76, 0x31, 0x3b,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x08,
	0x53, 0x6f, 0x6c, 0x76, 0x65, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x08, 0x53, 0x6f, 0x6c, 0x76, 0x65,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x14, 0x53, 0x6f, 0x6c, 0x76, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x53, 0x6f, 0x6c,
	0x76, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_solve_v1_solve_proto_rawDescOnce sync.Once
	file_solve_v1_solve_proto_rawDescData []byte
)

func file_solve_v1_solve_proto_rawDescGZIP() []byte {
	file_solve_v1_solve_proto_rawDescOnce.Do(func() {
		file_solve_v1_solve_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_solve_v1_solve_proto_rawDesc), len(file_solve_v1_solve_proto_rawDesc)))
	})
	return file_solve_v1_solve_proto_rawDescData
}

var file_solve_v1_solve_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_solve_v1_solve_proto_goTypes = []any{
	(*SolveForWCSFITSHandlerRequest)(nil),  // 0: solve.v1.SolveForWCSFITSHandlerRequest
	(*SolveForWCSFITSHandlerResponse)(nil), // 1: solve.v1.SolveForWCSFITSHandlerResponse
	(*v1.WCS)(nil),                         // 2: wcs.v1.WCS
}
var file_solve_v1_solve_proto_depIdxs = []int32{
	2, // 0: solve.v1.SolveForWCSFITSHandlerResponse.wcs:type_name -> wcs.v1.WCS
	0, // 1: solve.v1.SolveService.SolveForWCSFITSHandler:input_type -> solve.v1.SolveForWCSFITSHandlerRequest
	1, // 2: solve.v1.SolveService.SolveForWCSFITSHandler:output_type -> solve.v1.SolveForWCSFITSHandlerResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_solve_v1_solve_proto_init() }
func file_solve_v1_solve_proto_init() {
	if File_solve_v1_solve_proto != nil {
		return
	}
	file_solve_v1_solve_proto_msgTypes[0].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_solve_v1_solve_proto_rawDesc), len(file_solve_v1_solve_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_solve_v1_solve_proto_goTypes,
		DependencyIndexes: file_solve_v1_solve_proto_depIdxs,
		MessageInfos:      file_solve_v1_solve_proto_msgTypes,
	}.Build()
	File_solve_v1_solve_proto = out.File
	file_solve_v1_solve_proto_goTypes = nil
	file_solve_v1_solve_proto_depIdxs = nil
}
