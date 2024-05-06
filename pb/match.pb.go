// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v4.25.2
// source: match.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LikeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LikedId string `protobuf:"bytes,1,opt,name=likedId,proto3" json:"likedId,omitempty"`
	UserId  string `protobuf:"bytes,2,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *LikeRequest) Reset() {
	*x = LikeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LikeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LikeRequest) ProtoMessage() {}

func (x *LikeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LikeRequest.ProtoReflect.Descriptor instead.
func (*LikeRequest) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{0}
}

func (x *LikeRequest) GetLikedId() string {
	if x != nil {
		return x.LikedId
	}
	return ""
}

func (x *LikeRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type GetByUserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetByUserId) Reset() {
	*x = GetByUserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByUserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByUserId) ProtoMessage() {}

func (x *GetByUserId) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByUserId.ProtoReflect.Descriptor instead.
func (*GetByUserId) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{1}
}

func (x *GetByUserId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type MatchResposne struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	MatchId string `protobuf:"bytes,2,opt,name=matchId,proto3" json:"matchId,omitempty"`
	UserId  string `protobuf:"bytes,3,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *MatchResposne) Reset() {
	*x = MatchResposne{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MatchResposne) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MatchResposne) ProtoMessage() {}

func (x *MatchResposne) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MatchResposne.ProtoReflect.Descriptor instead.
func (*MatchResposne) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{2}
}

func (x *MatchResposne) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MatchResposne) GetMatchId() string {
	if x != nil {
		return x.MatchId
	}
	return ""
}

func (x *MatchResposne) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type NoPara struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NoPara) Reset() {
	*x = NoPara{}
	if protoimpl.UnsafeEnabled {
		mi := &file_match_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NoPara) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NoPara) ProtoMessage() {}

func (x *NoPara) ProtoReflect() protoreflect.Message {
	mi := &file_match_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NoPara.ProtoReflect.Descriptor instead.
func (*NoPara) Descriptor() ([]byte, []int) {
	return file_match_proto_rawDescGZIP(), []int{3}
}

var File_match_proto protoreflect.FileDescriptor

var file_match_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x22, 0x3f, 0x0a, 0x0b, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x69, 0x6b, 0x65, 0x64, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x1d, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x0d, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x73, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x08, 0x0a, 0x06, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61,
	0x32, 0xc4, 0x01, 0x0a, 0x0c, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x27, 0x0a, 0x04, 0x4c, 0x69, 0x6b, 0x65, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x4c, 0x69, 0x6b, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x12, 0x29, 0x0a, 0x06, 0x55, 0x6e,
	0x6c, 0x69, 0x6b, 0x65, 0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x6b, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4e,
	0x6f, 0x50, 0x61, 0x72, 0x61, 0x12, 0x2a, 0x0a, 0x07, 0x55, 0x6e, 0x4d, 0x61, 0x74, 0x63, 0x68,
	0x12, 0x11, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x1a, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4e, 0x6f, 0x50, 0x61, 0x72,
	0x61, 0x12, 0x34, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x12, 0x11, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x1a, 0x13, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4d, 0x61, 0x74, 0x63, 0x68, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x73, 0x6e, 0x65, 0x30, 0x01, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_match_proto_rawDescOnce sync.Once
	file_match_proto_rawDescData = file_match_proto_rawDesc
)

func file_match_proto_rawDescGZIP() []byte {
	file_match_proto_rawDescOnce.Do(func() {
		file_match_proto_rawDescData = protoimpl.X.CompressGZIP(file_match_proto_rawDescData)
	})
	return file_match_proto_rawDescData
}

var file_match_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_match_proto_goTypes = []interface{}{
	(*LikeRequest)(nil),   // 0: user.LikeRequest
	(*GetByUserId)(nil),   // 1: user.GetByUserId
	(*MatchResposne)(nil), // 2: user.MatchResposne
	(*NoPara)(nil),        // 3: user.NoPara
}
var file_match_proto_depIdxs = []int32{
	0, // 0: user.MatchService.Like:input_type -> user.LikeRequest
	0, // 1: user.MatchService.Unlike:input_type -> user.LikeRequest
	1, // 2: user.MatchService.UnMatch:input_type -> user.GetByUserId
	1, // 3: user.MatchService.GetMatch:input_type -> user.GetByUserId
	3, // 4: user.MatchService.Like:output_type -> user.NoPara
	3, // 5: user.MatchService.Unlike:output_type -> user.NoPara
	3, // 6: user.MatchService.UnMatch:output_type -> user.NoPara
	2, // 7: user.MatchService.GetMatch:output_type -> user.MatchResposne
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_match_proto_init() }
func file_match_proto_init() {
	if File_match_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_match_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LikeRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_match_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByUserId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_match_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MatchResposne); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_match_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NoPara); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_match_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_match_proto_goTypes,
		DependencyIndexes: file_match_proto_depIdxs,
		MessageInfos:      file_match_proto_msgTypes,
	}.Build()
	File_match_proto = out.File
	file_match_proto_rawDesc = nil
	file_match_proto_goTypes = nil
	file_match_proto_depIdxs = nil
}
