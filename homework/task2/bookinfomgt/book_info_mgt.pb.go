// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: bookinfomgt/book_info_mgt.proto

package bookinfomgt

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

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookinfomgt_book_info_mgt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_bookinfomgt_book_info_mgt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_bookinfomgt_book_info_mgt_proto_rawDescGZIP(), []int{0}
}

func (x *Book) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type AddBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *AddBookResponse) Reset() {
	*x = AddBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookinfomgt_book_info_mgt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookResponse) ProtoMessage() {}

func (x *AddBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookinfomgt_book_info_mgt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookResponse.ProtoReflect.Descriptor instead.
func (*AddBookResponse) Descriptor() ([]byte, []int) {
	return file_bookinfomgt_book_info_mgt_proto_rawDescGZIP(), []int{1}
}

func (x *AddBookResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

type QueryByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *QueryByIdRequest) Reset() {
	*x = QueryByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookinfomgt_book_info_mgt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryByIdRequest) ProtoMessage() {}

func (x *QueryByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookinfomgt_book_info_mgt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryByIdRequest.ProtoReflect.Descriptor instead.
func (*QueryByIdRequest) Descriptor() ([]byte, []int) {
	return file_bookinfomgt_book_info_mgt_proto_rawDescGZIP(), []int{2}
}

func (x *QueryByIdRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type BookList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *BookList) Reset() {
	*x = BookList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookinfomgt_book_info_mgt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BookList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BookList) ProtoMessage() {}

func (x *BookList) ProtoReflect() protoreflect.Message {
	mi := &file_bookinfomgt_book_info_mgt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BookList.ProtoReflect.Descriptor instead.
func (*BookList) Descriptor() ([]byte, []int) {
	return file_bookinfomgt_book_info_mgt_proto_rawDescGZIP(), []int{3}
}

func (x *BookList) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type QueryByNameRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *QueryByNameRequest) Reset() {
	*x = QueryByNameRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookinfomgt_book_info_mgt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueryByNameRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryByNameRequest) ProtoMessage() {}

func (x *QueryByNameRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookinfomgt_book_info_mgt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryByNameRequest.ProtoReflect.Descriptor instead.
func (*QueryByNameRequest) Descriptor() ([]byte, []int) {
	return file_bookinfomgt_book_info_mgt_proto_rawDescGZIP(), []int{4}
}

func (x *QueryByNameRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookinfomgt_book_info_mgt_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookinfomgt_book_info_mgt_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_bookinfomgt_book_info_mgt_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ok bool `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
}

func (x *DeleteResponse) Reset() {
	*x = DeleteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookinfomgt_book_info_mgt_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteResponse) ProtoMessage() {}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookinfomgt_book_info_mgt_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteResponse.ProtoReflect.Descriptor instead.
func (*DeleteResponse) Descriptor() ([]byte, []int) {
	return file_bookinfomgt_book_info_mgt_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteResponse) GetOk() bool {
	if x != nil {
		return x.Ok
	}
	return false
}

var File_bookinfomgt_book_info_mgt_proto protoreflect.FileDescriptor

var file_bookinfomgt_book_info_mgt_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x6d, 0x67, 0x74, 0x2f, 0x62, 0x6f,
	0x6f, 0x6b, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x5f, 0x6d, 0x67, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0b, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x6d, 0x67, 0x74, 0x22, 0x2a,
	0x0a, 0x04, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x41, 0x64,
	0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a,
	0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x02, 0x6f, 0x6b, 0x22, 0x22, 0x0a,
	0x10, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x33, 0x0a, 0x08, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x27, 0x0a,
	0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62,
	0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x6d, 0x67, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52,
	0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x22, 0x28, 0x0a, 0x12, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42,
	0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x22, 0x1f, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x20, 0x0a, 0x0e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x6f, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x02, 0x6f, 0x6b, 0x32, 0x9a, 0x02, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x49, 0x6e, 0x66, 0x6f,
	0x4d, 0x67, 0x74, 0x12, 0x3c, 0x0a, 0x07, 0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x11,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x6d, 0x67, 0x74, 0x2e, 0x42, 0x6f, 0x6f,
	0x6b, 0x1a, 0x1c, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x6d, 0x67, 0x74, 0x2e,
	0x41, 0x64, 0x64, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3f, 0x0a, 0x09, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x12, 0x1d,
	0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x6d, 0x67, 0x74, 0x2e, 0x51, 0x75, 0x65,
	0x72, 0x79, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e,
	0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x6d, 0x67, 0x74, 0x2e, 0x42, 0x6f, 0x6f, 0x6b,
	0x22, 0x00, 0x12, 0x47, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x6d, 0x67, 0x74, 0x2e,
	0x51, 0x75, 0x65, 0x72, 0x79, 0x42, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x15, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x6d, 0x67, 0x74,
	0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x06, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1a, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x66, 0x6f,
	0x6d, 0x67, 0x74, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x6d, 0x67, 0x74, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x1c, 0x5a, 0x1a, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x2f, 0x62, 0x6f, 0x6f, 0x6b, 0x69, 0x6e, 0x66, 0x6f, 0x6d, 0x67, 0x74, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bookinfomgt_book_info_mgt_proto_rawDescOnce sync.Once
	file_bookinfomgt_book_info_mgt_proto_rawDescData = file_bookinfomgt_book_info_mgt_proto_rawDesc
)

func file_bookinfomgt_book_info_mgt_proto_rawDescGZIP() []byte {
	file_bookinfomgt_book_info_mgt_proto_rawDescOnce.Do(func() {
		file_bookinfomgt_book_info_mgt_proto_rawDescData = protoimpl.X.CompressGZIP(file_bookinfomgt_book_info_mgt_proto_rawDescData)
	})
	return file_bookinfomgt_book_info_mgt_proto_rawDescData
}

var file_bookinfomgt_book_info_mgt_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_bookinfomgt_book_info_mgt_proto_goTypes = []interface{}{
	(*Book)(nil),               // 0: bookinfomgt.Book
	(*AddBookResponse)(nil),    // 1: bookinfomgt.AddBookResponse
	(*QueryByIdRequest)(nil),   // 2: bookinfomgt.QueryByIdRequest
	(*BookList)(nil),           // 3: bookinfomgt.BookList
	(*QueryByNameRequest)(nil), // 4: bookinfomgt.QueryByNameRequest
	(*DeleteRequest)(nil),      // 5: bookinfomgt.DeleteRequest
	(*DeleteResponse)(nil),     // 6: bookinfomgt.DeleteResponse
}
var file_bookinfomgt_book_info_mgt_proto_depIdxs = []int32{
	0, // 0: bookinfomgt.BookList.books:type_name -> bookinfomgt.Book
	0, // 1: bookinfomgt.BookInfoMgt.AddBook:input_type -> bookinfomgt.Book
	2, // 2: bookinfomgt.BookInfoMgt.QueryById:input_type -> bookinfomgt.QueryByIdRequest
	4, // 3: bookinfomgt.BookInfoMgt.QueryByName:input_type -> bookinfomgt.QueryByNameRequest
	5, // 4: bookinfomgt.BookInfoMgt.Delete:input_type -> bookinfomgt.DeleteRequest
	1, // 5: bookinfomgt.BookInfoMgt.AddBook:output_type -> bookinfomgt.AddBookResponse
	0, // 6: bookinfomgt.BookInfoMgt.QueryById:output_type -> bookinfomgt.Book
	3, // 7: bookinfomgt.BookInfoMgt.QueryByName:output_type -> bookinfomgt.BookList
	6, // 8: bookinfomgt.BookInfoMgt.Delete:output_type -> bookinfomgt.DeleteResponse
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_bookinfomgt_book_info_mgt_proto_init() }
func file_bookinfomgt_book_info_mgt_proto_init() {
	if File_bookinfomgt_book_info_mgt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bookinfomgt_book_info_mgt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
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
		file_bookinfomgt_book_info_mgt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddBookResponse); i {
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
		file_bookinfomgt_book_info_mgt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryByIdRequest); i {
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
		file_bookinfomgt_book_info_mgt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BookList); i {
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
		file_bookinfomgt_book_info_mgt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueryByNameRequest); i {
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
		file_bookinfomgt_book_info_mgt_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRequest); i {
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
		file_bookinfomgt_book_info_mgt_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteResponse); i {
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
			RawDescriptor: file_bookinfomgt_book_info_mgt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_bookinfomgt_book_info_mgt_proto_goTypes,
		DependencyIndexes: file_bookinfomgt_book_info_mgt_proto_depIdxs,
		MessageInfos:      file_bookinfomgt_book_info_mgt_proto_msgTypes,
	}.Build()
	File_bookinfomgt_book_info_mgt_proto = out.File
	file_bookinfomgt_book_info_mgt_proto_rawDesc = nil
	file_bookinfomgt_book_info_mgt_proto_goTypes = nil
	file_bookinfomgt_book_info_mgt_proto_depIdxs = nil
}
