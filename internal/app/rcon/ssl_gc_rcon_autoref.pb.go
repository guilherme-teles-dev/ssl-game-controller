// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.22.3
// source: ssl_gc_rcon_autoref.proto

package rcon

import (
	state "github.com/RoboCup-SSL/ssl-game-controller/internal/app/state"
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

// AutoRefRegistration is the first message that a client must send to the controller to identify itself
type AutoRefRegistration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// identifier is a unique name of the client
	Identifier *string `protobuf:"bytes,1,req,name=identifier" json:"identifier,omitempty"`
	// signature can optionally be specified to enable secure communication
	Signature *Signature `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
}

func (x *AutoRefRegistration) Reset() {
	*x = AutoRefRegistration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_rcon_autoref_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoRefRegistration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoRefRegistration) ProtoMessage() {}

func (x *AutoRefRegistration) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_rcon_autoref_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoRefRegistration.ProtoReflect.Descriptor instead.
func (*AutoRefRegistration) Descriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_autoref_proto_rawDescGZIP(), []int{0}
}

func (x *AutoRefRegistration) GetIdentifier() string {
	if x != nil && x.Identifier != nil {
		return *x.Identifier
	}
	return ""
}

func (x *AutoRefRegistration) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

// AutoRefToController is the wrapper message for all subsequent messages from the autoRef to the controller
type AutoRefToController struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// signature can optionally be specified to enable secure communication
	Signature *Signature `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	// game_event is an optional event that the autoRef detected during the game
	GameEvent *state.GameEvent `protobuf:"bytes,2,opt,name=game_event,json=gameEvent" json:"game_event,omitempty"`
}

func (x *AutoRefToController) Reset() {
	*x = AutoRefToController{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_rcon_autoref_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoRefToController) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoRefToController) ProtoMessage() {}

func (x *AutoRefToController) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_rcon_autoref_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoRefToController.ProtoReflect.Descriptor instead.
func (*AutoRefToController) Descriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_autoref_proto_rawDescGZIP(), []int{1}
}

func (x *AutoRefToController) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *AutoRefToController) GetGameEvent() *state.GameEvent {
	if x != nil {
		return x.GameEvent
	}
	return nil
}

// ControllerToAutoRef is the wrapper message for all messages from controller to autoRef
type ControllerToAutoRef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//
	//	*ControllerToAutoRef_ControllerReply
	Msg isControllerToAutoRef_Msg `protobuf_oneof:"msg"`
}

func (x *ControllerToAutoRef) Reset() {
	*x = ControllerToAutoRef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_rcon_autoref_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerToAutoRef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerToAutoRef) ProtoMessage() {}

func (x *ControllerToAutoRef) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_rcon_autoref_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerToAutoRef.ProtoReflect.Descriptor instead.
func (*ControllerToAutoRef) Descriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_autoref_proto_rawDescGZIP(), []int{2}
}

func (m *ControllerToAutoRef) GetMsg() isControllerToAutoRef_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *ControllerToAutoRef) GetControllerReply() *ControllerReply {
	if x, ok := x.GetMsg().(*ControllerToAutoRef_ControllerReply); ok {
		return x.ControllerReply
	}
	return nil
}

type isControllerToAutoRef_Msg interface {
	isControllerToAutoRef_Msg()
}

type ControllerToAutoRef_ControllerReply struct {
	// a reply from the controller
	ControllerReply *ControllerReply `protobuf:"bytes,1,opt,name=controller_reply,json=controllerReply,oneof"`
}

func (*ControllerToAutoRef_ControllerReply) isControllerToAutoRef_Msg() {}

var File_ssl_gc_rcon_autoref_proto protoreflect.FileDescriptor

var file_ssl_gc_rcon_autoref_proto_rawDesc = []byte{
	0x0a, 0x19, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x72, 0x63, 0x6f, 0x6e, 0x5f, 0x61, 0x75,
	0x74, 0x6f, 0x72, 0x65, 0x66, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x73, 0x73, 0x6c,
	0x5f, 0x67, 0x63, 0x5f, 0x67, 0x61, 0x6d, 0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x72, 0x63, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5f, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x6f, 0x52,
	0x65, 0x66, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x18, 0x01, 0x20, 0x02,
	0x28, 0x09, 0x52, 0x0a, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x12, 0x28,
	0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0x76, 0x0a, 0x13, 0x41, 0x75, 0x74, 0x6f,
	0x52, 0x65, 0x66, 0x54, 0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12,
	0x28, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x29, 0x0a, 0x0a, 0x67, 0x61, 0x6d,
	0x65, 0x5f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x47, 0x61, 0x6d, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x52, 0x09, 0x67, 0x61, 0x6d, 0x65, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05,
	0x22, 0x5b, 0x0a, 0x13, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x54, 0x6f,
	0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x66, 0x12, 0x3d, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x42, 0x3e, 0x5a,
	0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x62, 0x6f,
	0x43, 0x75, 0x70, 0x2d, 0x53, 0x53, 0x4c, 0x2f, 0x73, 0x73, 0x6c, 0x2d, 0x67, 0x61, 0x6d, 0x65,
	0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x72, 0x63, 0x6f, 0x6e,
}

var (
	file_ssl_gc_rcon_autoref_proto_rawDescOnce sync.Once
	file_ssl_gc_rcon_autoref_proto_rawDescData = file_ssl_gc_rcon_autoref_proto_rawDesc
)

func file_ssl_gc_rcon_autoref_proto_rawDescGZIP() []byte {
	file_ssl_gc_rcon_autoref_proto_rawDescOnce.Do(func() {
		file_ssl_gc_rcon_autoref_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssl_gc_rcon_autoref_proto_rawDescData)
	})
	return file_ssl_gc_rcon_autoref_proto_rawDescData
}

var file_ssl_gc_rcon_autoref_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_ssl_gc_rcon_autoref_proto_goTypes = []interface{}{
	(*AutoRefRegistration)(nil), // 0: AutoRefRegistration
	(*AutoRefToController)(nil), // 1: AutoRefToController
	(*ControllerToAutoRef)(nil), // 2: ControllerToAutoRef
	(*Signature)(nil),           // 3: Signature
	(*state.GameEvent)(nil),     // 4: GameEvent
	(*ControllerReply)(nil),     // 5: ControllerReply
}
var file_ssl_gc_rcon_autoref_proto_depIdxs = []int32{
	3, // 0: AutoRefRegistration.signature:type_name -> Signature
	3, // 1: AutoRefToController.signature:type_name -> Signature
	4, // 2: AutoRefToController.game_event:type_name -> GameEvent
	5, // 3: ControllerToAutoRef.controller_reply:type_name -> ControllerReply
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_ssl_gc_rcon_autoref_proto_init() }
func file_ssl_gc_rcon_autoref_proto_init() {
	if File_ssl_gc_rcon_autoref_proto != nil {
		return
	}
	file_ssl_gc_rcon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ssl_gc_rcon_autoref_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoRefRegistration); i {
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
		file_ssl_gc_rcon_autoref_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoRefToController); i {
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
		file_ssl_gc_rcon_autoref_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControllerToAutoRef); i {
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
	file_ssl_gc_rcon_autoref_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ControllerToAutoRef_ControllerReply)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ssl_gc_rcon_autoref_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ssl_gc_rcon_autoref_proto_goTypes,
		DependencyIndexes: file_ssl_gc_rcon_autoref_proto_depIdxs,
		MessageInfos:      file_ssl_gc_rcon_autoref_proto_msgTypes,
	}.Build()
	File_ssl_gc_rcon_autoref_proto = out.File
	file_ssl_gc_rcon_autoref_proto_rawDesc = nil
	file_ssl_gc_rcon_autoref_proto_goTypes = nil
	file_ssl_gc_rcon_autoref_proto_depIdxs = nil
}
