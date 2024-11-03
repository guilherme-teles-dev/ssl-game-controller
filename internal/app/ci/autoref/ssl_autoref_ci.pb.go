// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.3
// source: ssl_autoref_ci.proto

package autoref

import (
	state "github.com/RoboCup-SSL/ssl-game-controller/internal/app/state"
	tracker "github.com/RoboCup-SSL/ssl-game-controller/internal/app/tracker"
	vision "github.com/RoboCup-SSL/ssl-game-controller/internal/app/vision"
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

// The AutoRefCiInput contains all packets/messages that would otherwise be received through multicast by the auto-referee
// It may contain either a raw or a tracked SSL-vision packet. If both are given, the implementation may choose either one.
type AutoRefCiInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Latest referee message
	RefereeMessage *state.Referee `protobuf:"bytes,1,opt,name=referee_message,json=refereeMessage" json:"referee_message,omitempty"`
	// A tracked SSL-Vision packet to be processed without filtering
	TrackerWrapperPacket *tracker.TrackerWrapperPacket `protobuf:"bytes,2,opt,name=tracker_wrapper_packet,json=trackerWrapperPacket" json:"tracker_wrapper_packet,omitempty"`
	// A list of unfiltered SSL-Vision packets (for multiple cameras) to be filtered and processed
	Detection []*vision.SSL_DetectionFrame `protobuf:"bytes,3,rep,name=detection" json:"detection,omitempty"`
	// Current geometry data, to be sent at least once at the beginning of the connection
	Geometry *vision.SSL_GeometryData `protobuf:"bytes,4,opt,name=geometry" json:"geometry,omitempty"`
}

func (x *AutoRefCiInput) Reset() {
	*x = AutoRefCiInput{}
	mi := &file_ssl_autoref_ci_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AutoRefCiInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoRefCiInput) ProtoMessage() {}

func (x *AutoRefCiInput) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_autoref_ci_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoRefCiInput.ProtoReflect.Descriptor instead.
func (*AutoRefCiInput) Descriptor() ([]byte, []int) {
	return file_ssl_autoref_ci_proto_rawDescGZIP(), []int{0}
}

func (x *AutoRefCiInput) GetRefereeMessage() *state.Referee {
	if x != nil {
		return x.RefereeMessage
	}
	return nil
}

func (x *AutoRefCiInput) GetTrackerWrapperPacket() *tracker.TrackerWrapperPacket {
	if x != nil {
		return x.TrackerWrapperPacket
	}
	return nil
}

func (x *AutoRefCiInput) GetDetection() []*vision.SSL_DetectionFrame {
	if x != nil {
		return x.Detection
	}
	return nil
}

func (x *AutoRefCiInput) GetGeometry() *vision.SSL_GeometryData {
	if x != nil {
		return x.Geometry
	}
	return nil
}

// The AutoRefCiOutput contains any new data created by the auto-referee for further processing
type AutoRefCiOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A resulting tracked SSL-Vision packet for input into the ssl-game-controller.
	// The auto-referee will either generate it from the unfiltered SSL-Vision packets
	// or simply return the tracked packet from the input.
	TrackerWrapperPacket *tracker.TrackerWrapperPacket `protobuf:"bytes,1,opt,name=tracker_wrapper_packet,json=trackerWrapperPacket" json:"tracker_wrapper_packet,omitempty"`
}

func (x *AutoRefCiOutput) Reset() {
	*x = AutoRefCiOutput{}
	mi := &file_ssl_autoref_ci_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AutoRefCiOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoRefCiOutput) ProtoMessage() {}

func (x *AutoRefCiOutput) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_autoref_ci_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoRefCiOutput.ProtoReflect.Descriptor instead.
func (*AutoRefCiOutput) Descriptor() ([]byte, []int) {
	return file_ssl_autoref_ci_proto_rawDescGZIP(), []int{1}
}

func (x *AutoRefCiOutput) GetTrackerWrapperPacket() *tracker.TrackerWrapperPacket {
	if x != nil {
		return x.TrackerWrapperPacket
	}
	return nil
}

var File_ssl_autoref_ci_proto protoreflect.FileDescriptor

var file_ssl_autoref_ci_proto_rawDesc = []byte{
	0x0a, 0x14, 0x73, 0x73, 0x6c, 0x5f, 0x61, 0x75, 0x74, 0x6f, 0x72, 0x65, 0x66, 0x5f, 0x63, 0x69,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x73, 0x73, 0x6c, 0x5f, 0x76, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x73, 0x73, 0x6c, 0x5f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x73, 0x73, 0x6c, 0x5f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e,
	0x5f, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x65, 0x5f,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf2, 0x01,
	0x0a, 0x0e, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x66, 0x43, 0x69, 0x49, 0x6e, 0x70, 0x75, 0x74,
	0x12, 0x31, 0x0a, 0x0f, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x65, 0x5f, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x08, 0x2e, 0x52, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x65, 0x52, 0x0e, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x4b, 0x0a, 0x16, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x5f, 0x77,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x57, 0x72, 0x61,
	0x70, 0x70, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x14, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x65, 0x72, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74,
	0x12, 0x31, 0x0a, 0x09, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x53, 0x53, 0x4c, 0x5f, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x52, 0x09, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x53, 0x53, 0x4c, 0x5f, 0x47, 0x65, 0x6f, 0x6d,
	0x65, 0x74, 0x72, 0x79, 0x44, 0x61, 0x74, 0x61, 0x52, 0x08, 0x67, 0x65, 0x6f, 0x6d, 0x65, 0x74,
	0x72, 0x79, 0x22, 0x5e, 0x0a, 0x0f, 0x41, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x66, 0x43, 0x69, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x4b, 0x0a, 0x16, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72,
	0x5f, 0x77, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x54, 0x72, 0x61, 0x63, 0x6b, 0x65, 0x72, 0x57,
	0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b, 0x65, 0x74, 0x52, 0x14, 0x74, 0x72,
	0x61, 0x63, 0x6b, 0x65, 0x72, 0x57, 0x72, 0x61, 0x70, 0x70, 0x65, 0x72, 0x50, 0x61, 0x63, 0x6b,
	0x65, 0x74, 0x42, 0x44, 0x5a, 0x42, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x52, 0x6f, 0x62, 0x6f, 0x43, 0x75, 0x70, 0x2d, 0x53, 0x53, 0x4c, 0x2f, 0x73, 0x73, 0x6c,
	0x2d, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72,
	0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70, 0x2f, 0x63, 0x69,
	0x2f, 0x61, 0x75, 0x74, 0x6f, 0x72, 0x65, 0x66,
}

var (
	file_ssl_autoref_ci_proto_rawDescOnce sync.Once
	file_ssl_autoref_ci_proto_rawDescData = file_ssl_autoref_ci_proto_rawDesc
)

func file_ssl_autoref_ci_proto_rawDescGZIP() []byte {
	file_ssl_autoref_ci_proto_rawDescOnce.Do(func() {
		file_ssl_autoref_ci_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssl_autoref_ci_proto_rawDescData)
	})
	return file_ssl_autoref_ci_proto_rawDescData
}

var file_ssl_autoref_ci_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_ssl_autoref_ci_proto_goTypes = []any{
	(*AutoRefCiInput)(nil),               // 0: AutoRefCiInput
	(*AutoRefCiOutput)(nil),              // 1: AutoRefCiOutput
	(*state.Referee)(nil),                // 2: Referee
	(*tracker.TrackerWrapperPacket)(nil), // 3: TrackerWrapperPacket
	(*vision.SSL_DetectionFrame)(nil),    // 4: SSL_DetectionFrame
	(*vision.SSL_GeometryData)(nil),      // 5: SSL_GeometryData
}
var file_ssl_autoref_ci_proto_depIdxs = []int32{
	2, // 0: AutoRefCiInput.referee_message:type_name -> Referee
	3, // 1: AutoRefCiInput.tracker_wrapper_packet:type_name -> TrackerWrapperPacket
	4, // 2: AutoRefCiInput.detection:type_name -> SSL_DetectionFrame
	5, // 3: AutoRefCiInput.geometry:type_name -> SSL_GeometryData
	3, // 4: AutoRefCiOutput.tracker_wrapper_packet:type_name -> TrackerWrapperPacket
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_ssl_autoref_ci_proto_init() }
func file_ssl_autoref_ci_proto_init() {
	if File_ssl_autoref_ci_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ssl_autoref_ci_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ssl_autoref_ci_proto_goTypes,
		DependencyIndexes: file_ssl_autoref_ci_proto_depIdxs,
		MessageInfos:      file_ssl_autoref_ci_proto_msgTypes,
	}.Build()
	File_ssl_autoref_ci_proto = out.File
	file_ssl_autoref_ci_proto_rawDesc = nil
	file_ssl_autoref_ci_proto_goTypes = nil
	file_ssl_autoref_ci_proto_depIdxs = nil
}
