// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.23.2
// source: tensorflow/tsl/profiler/protobuf/xplane.proto

package for_profiler_protos_go_proto

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

// A container of parallel XPlanes, generated by one or more profiling sources.
// Next ID: 5
type XSpace struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Planes []*XPlane `protobuf:"bytes,1,rep,name=planes,proto3" json:"planes,omitempty"`
	// Errors (if any) in the generation of planes.
	Errors []string `protobuf:"bytes,2,rep,name=errors,proto3" json:"errors,omitempty"`
	// Warnings (if any) in the generation of planes;
	Warnings []string `protobuf:"bytes,3,rep,name=warnings,proto3" json:"warnings,omitempty"`
	// List of hostnames that XPlanes are generated from.
	Hostnames []string `protobuf:"bytes,4,rep,name=hostnames,proto3" json:"hostnames,omitempty"`
}

func (x *XSpace) Reset() {
	*x = XSpace{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XSpace) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XSpace) ProtoMessage() {}

func (x *XSpace) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XSpace.ProtoReflect.Descriptor instead.
func (*XSpace) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDescGZIP(), []int{0}
}

func (x *XSpace) GetPlanes() []*XPlane {
	if x != nil {
		return x.Planes
	}
	return nil
}

func (x *XSpace) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

func (x *XSpace) GetWarnings() []string {
	if x != nil {
		return x.Warnings
	}
	return nil
}

func (x *XSpace) GetHostnames() []string {
	if x != nil {
		return x.Hostnames
	}
	return nil
}

// An XPlane is a container of parallel timelines (XLines), generated by a
// profiling source or by post-processing one or more XPlanes.
// Next ID: 7
type XPlane struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of this XPlane.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Parallel timelines grouped in this plane. XLines with the same id
	// are effectively the same timeline.
	Lines []*XLine `protobuf:"bytes,3,rep,name=lines,proto3" json:"lines,omitempty"`
	// XEventMetadata map, each entry uses the XEventMetadata.id as key. This map
	// should be used for events that share the same ID over the whole XPlane.
	EventMetadata map[int64]*XEventMetadata `protobuf:"bytes,4,rep,name=event_metadata,json=eventMetadata,proto3" json:"event_metadata,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// XStatMetadata map, each entry uses the XStatMetadata.id as key. This map
	// should be used for stats that share the same ID over the whole XPlane.
	StatMetadata map[int64]*XStatMetadata `protobuf:"bytes,5,rep,name=stat_metadata,json=statMetadata,proto3" json:"stat_metadata,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// XStats associated with this plane, e.g. device capabilities.
	// Each of these XStats should have a different metadata_id.
	Stats []*XStat `protobuf:"bytes,6,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *XPlane) Reset() {
	*x = XPlane{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XPlane) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XPlane) ProtoMessage() {}

func (x *XPlane) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XPlane.ProtoReflect.Descriptor instead.
func (*XPlane) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDescGZIP(), []int{1}
}

func (x *XPlane) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *XPlane) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *XPlane) GetLines() []*XLine {
	if x != nil {
		return x.Lines
	}
	return nil
}

func (x *XPlane) GetEventMetadata() map[int64]*XEventMetadata {
	if x != nil {
		return x.EventMetadata
	}
	return nil
}

func (x *XPlane) GetStatMetadata() map[int64]*XStatMetadata {
	if x != nil {
		return x.StatMetadata
	}
	return nil
}

func (x *XPlane) GetStats() []*XStat {
	if x != nil {
		return x.Stats
	}
	return nil
}

// An XLine is a timeline of trace events (XEvents).
// Next ID: 12
type XLine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Id of this line, can be repeated within an XPlane. All XLines with the
	// same id are effectively the same timeline.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Display id of this line. Multiple lines with the same display_id are
	// grouped together in the same trace viewer row.
	DisplayId int64 `protobuf:"varint,10,opt,name=display_id,json=displayId,proto3" json:"display_id,omitempty"`
	// Name of this XLine.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Name of this XLine to display in trace viewer.
	DisplayName string `protobuf:"bytes,11,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Start time of this line in nanoseconds since the UNIX epoch.
	// XEvent.offset_ps is relative to this timestamp.
	TimestampNs int64 `protobuf:"varint,3,opt,name=timestamp_ns,json=timestampNs,proto3" json:"timestamp_ns,omitempty"`
	// Profiling duration for this line in picoseconds.
	DurationPs int64 `protobuf:"varint,9,opt,name=duration_ps,json=durationPs,proto3" json:"duration_ps,omitempty"`
	// XEvents within the same XLine should not overlap in time, but they can be
	// nested.
	Events []*XEvent `protobuf:"bytes,4,rep,name=events,proto3" json:"events,omitempty"`
}

func (x *XLine) Reset() {
	*x = XLine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XLine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XLine) ProtoMessage() {}

func (x *XLine) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XLine.ProtoReflect.Descriptor instead.
func (*XLine) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDescGZIP(), []int{2}
}

func (x *XLine) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *XLine) GetDisplayId() int64 {
	if x != nil {
		return x.DisplayId
	}
	return 0
}

func (x *XLine) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *XLine) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *XLine) GetTimestampNs() int64 {
	if x != nil {
		return x.TimestampNs
	}
	return 0
}

func (x *XLine) GetDurationPs() int64 {
	if x != nil {
		return x.DurationPs
	}
	return 0
}

func (x *XLine) GetEvents() []*XEvent {
	if x != nil {
		return x.Events
	}
	return nil
}

// An XEvent is a trace event, optionally annotated with XStats.
// Next ID: 6
type XEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// XEventMetadata.id of corresponding metadata.
	MetadataId int64 `protobuf:"varint,1,opt,name=metadata_id,json=metadataId,proto3" json:"metadata_id,omitempty"`
	// Types that are assignable to Data:
	//
	//	*XEvent_OffsetPs
	//	*XEvent_NumOccurrences
	Data isXEvent_Data `protobuf_oneof:"data"`
	// Duration of the event in picoseconds. Can be zero for an instant event.
	DurationPs int64 `protobuf:"varint,3,opt,name=duration_ps,json=durationPs,proto3" json:"duration_ps,omitempty"`
	// XStats associated with the event.
	// Each of these XStats should have a different metadata_id.
	Stats []*XStat `protobuf:"bytes,4,rep,name=stats,proto3" json:"stats,omitempty"`
}

func (x *XEvent) Reset() {
	*x = XEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XEvent) ProtoMessage() {}

func (x *XEvent) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XEvent.ProtoReflect.Descriptor instead.
func (*XEvent) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDescGZIP(), []int{3}
}

func (x *XEvent) GetMetadataId() int64 {
	if x != nil {
		return x.MetadataId
	}
	return 0
}

func (m *XEvent) GetData() isXEvent_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *XEvent) GetOffsetPs() int64 {
	if x, ok := x.GetData().(*XEvent_OffsetPs); ok {
		return x.OffsetPs
	}
	return 0
}

func (x *XEvent) GetNumOccurrences() int64 {
	if x, ok := x.GetData().(*XEvent_NumOccurrences); ok {
		return x.NumOccurrences
	}
	return 0
}

func (x *XEvent) GetDurationPs() int64 {
	if x != nil {
		return x.DurationPs
	}
	return 0
}

func (x *XEvent) GetStats() []*XStat {
	if x != nil {
		return x.Stats
	}
	return nil
}

type isXEvent_Data interface {
	isXEvent_Data()
}

type XEvent_OffsetPs struct {
	// Start time of the event in picoseconds, as offset from
	// XLine.timestamp_ns().
	OffsetPs int64 `protobuf:"varint,2,opt,name=offset_ps,json=offsetPs,proto3,oneof"`
}

type XEvent_NumOccurrences struct {
	// Number of occurrences of the event, if aggregated.
	NumOccurrences int64 `protobuf:"varint,5,opt,name=num_occurrences,json=numOccurrences,proto3,oneof"`
}

func (*XEvent_OffsetPs) isXEvent_Data() {}

func (*XEvent_NumOccurrences) isXEvent_Data() {}

// An XStat is a named value associated with an XEvent, e.g., a performance
// counter value, a metric computed by a formula applied over nested XEvents
// and XStats.
// Next ID: 8
type XStat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// XStatMetadata.id of corresponding metadata.
	MetadataId int64 `protobuf:"varint,1,opt,name=metadata_id,json=metadataId,proto3" json:"metadata_id,omitempty"`
	// Value of this stat.
	//
	// Types that are assignable to Value:
	//
	//	*XStat_DoubleValue
	//	*XStat_Uint64Value
	//	*XStat_Int64Value
	//	*XStat_StrValue
	//	*XStat_BytesValue
	//	*XStat_RefValue
	Value isXStat_Value `protobuf_oneof:"value"`
}

func (x *XStat) Reset() {
	*x = XStat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XStat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XStat) ProtoMessage() {}

func (x *XStat) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XStat.ProtoReflect.Descriptor instead.
func (*XStat) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDescGZIP(), []int{4}
}

func (x *XStat) GetMetadataId() int64 {
	if x != nil {
		return x.MetadataId
	}
	return 0
}

func (m *XStat) GetValue() isXStat_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *XStat) GetDoubleValue() float64 {
	if x, ok := x.GetValue().(*XStat_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (x *XStat) GetUint64Value() uint64 {
	if x, ok := x.GetValue().(*XStat_Uint64Value); ok {
		return x.Uint64Value
	}
	return 0
}

func (x *XStat) GetInt64Value() int64 {
	if x, ok := x.GetValue().(*XStat_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (x *XStat) GetStrValue() string {
	if x, ok := x.GetValue().(*XStat_StrValue); ok {
		return x.StrValue
	}
	return ""
}

func (x *XStat) GetBytesValue() []byte {
	if x, ok := x.GetValue().(*XStat_BytesValue); ok {
		return x.BytesValue
	}
	return nil
}

func (x *XStat) GetRefValue() uint64 {
	if x, ok := x.GetValue().(*XStat_RefValue); ok {
		return x.RefValue
	}
	return 0
}

type isXStat_Value interface {
	isXStat_Value()
}

type XStat_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,2,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type XStat_Uint64Value struct {
	Uint64Value uint64 `protobuf:"varint,3,opt,name=uint64_value,json=uint64Value,proto3,oneof"`
}

type XStat_Int64Value struct {
	Int64Value int64 `protobuf:"varint,4,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type XStat_StrValue struct {
	StrValue string `protobuf:"bytes,5,opt,name=str_value,json=strValue,proto3,oneof"`
}

type XStat_BytesValue struct {
	BytesValue []byte `protobuf:"bytes,6,opt,name=bytes_value,json=bytesValue,proto3,oneof"`
}

type XStat_RefValue struct {
	// A string value that stored in XStatMetadata::name.
	RefValue uint64 `protobuf:"varint,7,opt,name=ref_value,json=refValue,proto3,oneof"`
}

func (*XStat_DoubleValue) isXStat_Value() {}

func (*XStat_Uint64Value) isXStat_Value() {}

func (*XStat_Int64Value) isXStat_Value() {}

func (*XStat_StrValue) isXStat_Value() {}

func (*XStat_BytesValue) isXStat_Value() {}

func (*XStat_RefValue) isXStat_Value() {}

// Metadata for an XEvent, corresponds to an event type and is shared by
// all XEvents with the same metadata_id.
// Next ID: 7
type XEventMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// XPlane.event_metadata map key.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the event.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Name of the event shown in trace viewer.
	DisplayName string `protobuf:"bytes,4,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// Additional metadata in serialized format.
	Metadata []byte `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// XStats that are constant for all XEvents with the same metadata_id.
	// Each of these XStats should have a different metadata_id.
	Stats []*XStat `protobuf:"bytes,5,rep,name=stats,proto3" json:"stats,omitempty"`
	// XPlane.event_metadata map key for children events.
	ChildId []int64 `protobuf:"varint,6,rep,packed,name=child_id,json=childId,proto3" json:"child_id,omitempty"`
}

func (x *XEventMetadata) Reset() {
	*x = XEventMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XEventMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XEventMetadata) ProtoMessage() {}

func (x *XEventMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XEventMetadata.ProtoReflect.Descriptor instead.
func (*XEventMetadata) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDescGZIP(), []int{5}
}

func (x *XEventMetadata) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *XEventMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *XEventMetadata) GetDisplayName() string {
	if x != nil {
		return x.DisplayName
	}
	return ""
}

func (x *XEventMetadata) GetMetadata() []byte {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *XEventMetadata) GetStats() []*XStat {
	if x != nil {
		return x.Stats
	}
	return nil
}

func (x *XEventMetadata) GetChildId() []int64 {
	if x != nil {
		return x.ChildId
	}
	return nil
}

// Metadata for an XStat, corresponds to a stat type and is shared by all
// XStats with the same metadata_id.
// Next ID: 4
type XStatMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// XPlane.stat_metadata map key.
	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	// Name of the stat (should be short).
	// Two XStatMetadata with different id should have different names.
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Description of the stat (might be long).
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *XStatMetadata) Reset() {
	*x = XStatMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *XStatMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*XStatMetadata) ProtoMessage() {}

func (x *XStatMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use XStatMetadata.ProtoReflect.Descriptor instead.
func (*XStatMetadata) Descriptor() ([]byte, []int) {
	return file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDescGZIP(), []int{6}
}

func (x *XStatMetadata) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *XStatMetadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *XStatMetadata) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_tensorflow_tsl_profiler_protobuf_xplane_proto protoreflect.FileDescriptor

var file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x73, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x78, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x13, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x72, 0x22, 0x8f, 0x01, 0x0a, 0x06, 0x58, 0x53, 0x70, 0x61, 0x63, 0x65, 0x12,
	0x33, 0x0a, 0x06, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x58, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x52, 0x06, 0x70, 0x6c,
	0x61, 0x6e, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08,
	0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x08,
	0x77, 0x61, 0x72, 0x6e, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x68, 0x6f, 0x73, 0x74,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x68, 0x6f, 0x73,
	0x74, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x22, 0x87, 0x04, 0x0a, 0x06, 0x58, 0x50, 0x6c, 0x61, 0x6e,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x30, 0x0a, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x58, 0x4c, 0x69, 0x6e, 0x65,
	0x52, 0x05, 0x6c, 0x69, 0x6e, 0x65, 0x73, 0x12, 0x55, 0x0a, 0x0e, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x58, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x2e, 0x45, 0x76, 0x65,
	0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x0d, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x52,
	0x0a, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x58, 0x50, 0x6c, 0x61,
	0x6e, 0x65, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x0c, 0x73, 0x74, 0x61, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x58, 0x53, 0x74, 0x61, 0x74, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x73, 0x1a, 0x65, 0x0a, 0x12, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x39, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x74, 0x65,
	0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x72, 0x2e, 0x58, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x63, 0x0a, 0x11, 0x53,
	0x74, 0x61, 0x74, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x38, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x58, 0x53, 0x74, 0x61, 0x74, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01,
	0x22, 0xfe, 0x01, 0x0a, 0x05, 0x58, 0x4c, 0x69, 0x6e, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x69,
	0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a,
	0x0c, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0b, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x5f, 0x6e, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x4e, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x70, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x50, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f,
	0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x58, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x52, 0x06, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x4a, 0x04, 0x08, 0x05, 0x10, 0x06, 0x4a,
	0x04, 0x08, 0x06, 0x10, 0x07, 0x4a, 0x04, 0x08, 0x07, 0x10, 0x08, 0x4a, 0x04, 0x08, 0x08, 0x10,
	0x09, 0x22, 0xce, 0x01, 0x0a, 0x06, 0x58, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49, 0x64, 0x12, 0x1d, 0x0a,
	0x09, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x5f, 0x70, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x08, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x50, 0x73, 0x12, 0x29, 0x0a, 0x0f,
	0x6e, 0x75, 0x6d, 0x5f, 0x6f, 0x63, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0e, 0x6e, 0x75, 0x6d, 0x4f, 0x63, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x64, 0x75,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74,
	0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72,
	0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e, 0x58, 0x53,
	0x74, 0x61, 0x74, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x22, 0xff, 0x01, 0x0a, 0x05, 0x58, 0x53, 0x74, 0x61, 0x74, 0x12, 0x1f, 0x0a, 0x0b,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x49, 0x64, 0x12, 0x23, 0x0a,
	0x0c, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x48, 0x00, 0x52, 0x0b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x23, 0x0a, 0x0c, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x48, 0x00, 0x52, 0x0b, 0x75, 0x69, 0x6e, 0x74,
	0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x69, 0x6e, 0x74, 0x36, 0x34,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0a,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09, 0x73, 0x74,
	0x72, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52,
	0x08, 0x73, 0x74, 0x72, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x21, 0x0a, 0x0b, 0x62, 0x79, 0x74,
	0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x48, 0x00,
	0x52, 0x0a, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1d, 0x0a, 0x09,
	0x72, 0x65, 0x66, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x48,
	0x00, 0x52, 0x08, 0x72, 0x65, 0x66, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x42, 0x07, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0xc0, 0x01, 0x0a, 0x0e, 0x58, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x64,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x30, 0x0a, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x65, 0x6e, 0x73,
	0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2e,
	0x58, 0x53, 0x74, 0x61, 0x74, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x73, 0x12, 0x19, 0x0a, 0x08,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x03, 0x28, 0x03, 0x52, 0x07,
	0x63, 0x68, 0x69, 0x6c, 0x64, 0x49, 0x64, 0x22, 0x55, 0x0a, 0x0d, 0x58, 0x53, 0x74, 0x61, 0x74,
	0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x66,
	0x5a, 0x61, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x65, 0x6e,
	0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c,
	0x6f, 0x77, 0x2f, 0x74, 0x65, 0x6e, 0x73, 0x6f, 0x72, 0x66, 0x6c, 0x6f, 0x77, 0x2f, 0x67, 0x6f,
	0x2f, 0x74, 0x73, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x66, 0x6f, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x72, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5f, 0x67, 0x6f, 0x5f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0xf8, 0x01, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDescOnce sync.Once
	file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDescData = file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDesc
)

func file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDescGZIP() []byte {
	file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDescOnce.Do(func() {
		file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDescData = protoimpl.X.CompressGZIP(file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDescData)
	})
	return file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDescData
}

var file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_tensorflow_tsl_profiler_protobuf_xplane_proto_goTypes = []interface{}{
	(*XSpace)(nil),         // 0: tensorflow.profiler.XSpace
	(*XPlane)(nil),         // 1: tensorflow.profiler.XPlane
	(*XLine)(nil),          // 2: tensorflow.profiler.XLine
	(*XEvent)(nil),         // 3: tensorflow.profiler.XEvent
	(*XStat)(nil),          // 4: tensorflow.profiler.XStat
	(*XEventMetadata)(nil), // 5: tensorflow.profiler.XEventMetadata
	(*XStatMetadata)(nil),  // 6: tensorflow.profiler.XStatMetadata
	nil,                    // 7: tensorflow.profiler.XPlane.EventMetadataEntry
	nil,                    // 8: tensorflow.profiler.XPlane.StatMetadataEntry
}
var file_tensorflow_tsl_profiler_protobuf_xplane_proto_depIdxs = []int32{
	1,  // 0: tensorflow.profiler.XSpace.planes:type_name -> tensorflow.profiler.XPlane
	2,  // 1: tensorflow.profiler.XPlane.lines:type_name -> tensorflow.profiler.XLine
	7,  // 2: tensorflow.profiler.XPlane.event_metadata:type_name -> tensorflow.profiler.XPlane.EventMetadataEntry
	8,  // 3: tensorflow.profiler.XPlane.stat_metadata:type_name -> tensorflow.profiler.XPlane.StatMetadataEntry
	4,  // 4: tensorflow.profiler.XPlane.stats:type_name -> tensorflow.profiler.XStat
	3,  // 5: tensorflow.profiler.XLine.events:type_name -> tensorflow.profiler.XEvent
	4,  // 6: tensorflow.profiler.XEvent.stats:type_name -> tensorflow.profiler.XStat
	4,  // 7: tensorflow.profiler.XEventMetadata.stats:type_name -> tensorflow.profiler.XStat
	5,  // 8: tensorflow.profiler.XPlane.EventMetadataEntry.value:type_name -> tensorflow.profiler.XEventMetadata
	6,  // 9: tensorflow.profiler.XPlane.StatMetadataEntry.value:type_name -> tensorflow.profiler.XStatMetadata
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_tensorflow_tsl_profiler_protobuf_xplane_proto_init() }
func file_tensorflow_tsl_profiler_protobuf_xplane_proto_init() {
	if File_tensorflow_tsl_profiler_protobuf_xplane_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XSpace); i {
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
		file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XPlane); i {
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
		file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XLine); i {
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
		file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XEvent); i {
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
		file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XStat); i {
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
		file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XEventMetadata); i {
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
		file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*XStatMetadata); i {
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
	file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*XEvent_OffsetPs)(nil),
		(*XEvent_NumOccurrences)(nil),
	}
	file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*XStat_DoubleValue)(nil),
		(*XStat_Uint64Value)(nil),
		(*XStat_Int64Value)(nil),
		(*XStat_StrValue)(nil),
		(*XStat_BytesValue)(nil),
		(*XStat_RefValue)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tensorflow_tsl_profiler_protobuf_xplane_proto_goTypes,
		DependencyIndexes: file_tensorflow_tsl_profiler_protobuf_xplane_proto_depIdxs,
		MessageInfos:      file_tensorflow_tsl_profiler_protobuf_xplane_proto_msgTypes,
	}.Build()
	File_tensorflow_tsl_profiler_protobuf_xplane_proto = out.File
	file_tensorflow_tsl_profiler_protobuf_xplane_proto_rawDesc = nil
	file_tensorflow_tsl_profiler_protobuf_xplane_proto_goTypes = nil
	file_tensorflow_tsl_profiler_protobuf_xplane_proto_depIdxs = nil
}
