// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        (unknown)
// source: codemen.org/inventory/gunk/v1/salesReport/all.proto

package salereportpb

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SalesReport struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId        int32  `protobuf:"varint,1,opt,name=ProductId,proto3" json:"ProductId,omitempty"`
	ProductName      string `protobuf:"bytes,2,opt,name=ProductName,proto3" json:"ProductName,omitempty"`
	PurchaseQuantity int32  `protobuf:"varint,3,opt,name=PurchaseQuantity,proto3" json:"PurchaseQuantity,omitempty"`
	SellQuantity     int32  `protobuf:"varint,4,opt,name=SellQuantity,proto3" json:"SellQuantity,omitempty"`
	StockQuantity    int32  `protobuf:"varint,5,opt,name=StockQuantity,proto3" json:"StockQuantity,omitempty"`
}

func (x *SalesReport) Reset() {
	*x = SalesReport{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codemen_org_inventory_gunk_v1_salesReport_all_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SalesReport) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SalesReport) ProtoMessage() {}

func (x *SalesReport) ProtoReflect() protoreflect.Message {
	mi := &file_codemen_org_inventory_gunk_v1_salesReport_all_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SalesReport.ProtoReflect.Descriptor instead.
func (*SalesReport) Descriptor() ([]byte, []int) {
	return file_codemen_org_inventory_gunk_v1_salesReport_all_proto_rawDescGZIP(), []int{0}
}

func (x *SalesReport) GetProductId() int32 {
	if x != nil {
		return x.ProductId
	}
	return 0
}

func (x *SalesReport) GetProductName() string {
	if x != nil {
		return x.ProductName
	}
	return ""
}

func (x *SalesReport) GetPurchaseQuantity() int32 {
	if x != nil {
		return x.PurchaseQuantity
	}
	return 0
}

func (x *SalesReport) GetSellQuantity() int32 {
	if x != nil {
		return x.SellQuantity
	}
	return 0
}

func (x *SalesReport) GetStockQuantity() int32 {
	if x != nil {
		return x.StockQuantity
	}
	return 0
}

type ListSalesReportRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SearchTerm string `protobuf:"bytes,1,opt,name=SearchTerm,proto3" json:"SearchTerm,omitempty"`
	Offset     int32  `protobuf:"varint,2,opt,name=Offset,proto3" json:"Offset,omitempty"`
	Limit      int32  `protobuf:"varint,3,opt,name=Limit,proto3" json:"Limit,omitempty"`
}

func (x *ListSalesReportRequest) Reset() {
	*x = ListSalesReportRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codemen_org_inventory_gunk_v1_salesReport_all_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSalesReportRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSalesReportRequest) ProtoMessage() {}

func (x *ListSalesReportRequest) ProtoReflect() protoreflect.Message {
	mi := &file_codemen_org_inventory_gunk_v1_salesReport_all_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSalesReportRequest.ProtoReflect.Descriptor instead.
func (*ListSalesReportRequest) Descriptor() ([]byte, []int) {
	return file_codemen_org_inventory_gunk_v1_salesReport_all_proto_rawDescGZIP(), []int{1}
}

func (x *ListSalesReportRequest) GetSearchTerm() string {
	if x != nil {
		return x.SearchTerm
	}
	return ""
}

func (x *ListSalesReportRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *ListSalesReportRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type ListSalesReportResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SalesReports []*SalesReport `protobuf:"bytes,1,rep,name=SalesReports,proto3" json:"SalesReports,omitempty"`
}

func (x *ListSalesReportResponse) Reset() {
	*x = ListSalesReportResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_codemen_org_inventory_gunk_v1_salesReport_all_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSalesReportResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSalesReportResponse) ProtoMessage() {}

func (x *ListSalesReportResponse) ProtoReflect() protoreflect.Message {
	mi := &file_codemen_org_inventory_gunk_v1_salesReport_all_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSalesReportResponse.ProtoReflect.Descriptor instead.
func (*ListSalesReportResponse) Descriptor() ([]byte, []int) {
	return file_codemen_org_inventory_gunk_v1_salesReport_all_proto_rawDescGZIP(), []int{2}
}

func (x *ListSalesReportResponse) GetSalesReports() []*SalesReport {
	if x != nil {
		return x.SalesReports
	}
	return nil
}

var File_codemen_org_inventory_gunk_v1_salesReport_all_proto protoreflect.FileDescriptor

var file_codemen_org_inventory_gunk_v1_salesReport_all_proto_rawDesc = []byte{
	0x0a, 0x33, 0x63, 0x6f, 0x64, 0x65, 0x6d, 0x65, 0x6e, 0x2e, 0x6f, 0x72, 0x67, 0x2f, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x67, 0x75, 0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f,
	0x73, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x6c, 0x6c, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x61, 0x6c, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x70, 0x62, 0x22, 0xc0, 0x01, 0x0a, 0x0b, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x12, 0x1d, 0x0a, 0x09, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00,
	0x50, 0x00, 0x12, 0x1f, 0x0a, 0x0b, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30,
	0x00, 0x50, 0x00, 0x12, 0x24, 0x0a, 0x10, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x51,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08,
	0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x20, 0x0a, 0x0c, 0x53, 0x65, 0x6c,
	0x6c, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x42,
	0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x21, 0x0a, 0x0d, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06,
	0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22, 0x77, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x61,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x1e, 0x0a, 0x0a, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x54, 0x65, 0x72, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00,
	0x12, 0x1a, 0x0a, 0x06, 0x4f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x42, 0x0a, 0x08, 0x00, 0x18, 0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x12, 0x19, 0x0a, 0x05,
	0x4c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x22,
	0x5e, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x53, 0x61,
	0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0x2e,
	0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x42, 0x0a, 0x08, 0x00, 0x18,
	0x00, 0x28, 0x00, 0x30, 0x00, 0x50, 0x00, 0x3a, 0x06, 0x08, 0x00, 0x10, 0x00, 0x18, 0x00, 0x32,
	0x7f, 0x0a, 0x0c, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x6a, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x12, 0x24, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x70,
	0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x73, 0x61, 0x6c, 0x65, 0x72,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x61, 0x6c, 0x65,
	0x73, 0x52, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x06, 0x88, 0x02, 0x00, 0x90, 0x02, 0x00, 0x28, 0x00, 0x30, 0x00, 0x1a, 0x03, 0x88, 0x02, 0x00,
	0x42, 0x51, 0x48, 0x01, 0x50, 0x00, 0x5a, 0x36, 0x63, 0x6f, 0x64, 0x65, 0x6d, 0x65, 0x6e, 0x2e,
	0x6f, 0x72, 0x67, 0x2f, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x67, 0x75,
	0x6e, 0x6b, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x61, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x3b, 0x73, 0x61, 0x6c, 0x65, 0x72, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x70, 0x62, 0x80, 0x01,
	0x00, 0x88, 0x01, 0x00, 0x90, 0x01, 0x00, 0xb8, 0x01, 0x00, 0xd8, 0x01, 0x00, 0xf8, 0x01, 0x01,
	0xd0, 0x02, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_codemen_org_inventory_gunk_v1_salesReport_all_proto_rawDescOnce sync.Once
	file_codemen_org_inventory_gunk_v1_salesReport_all_proto_rawDescData = file_codemen_org_inventory_gunk_v1_salesReport_all_proto_rawDesc
)

func file_codemen_org_inventory_gunk_v1_salesReport_all_proto_rawDescGZIP() []byte {
	file_codemen_org_inventory_gunk_v1_salesReport_all_proto_rawDescOnce.Do(func() {
		file_codemen_org_inventory_gunk_v1_salesReport_all_proto_rawDescData = protoimpl.X.CompressGZIP(file_codemen_org_inventory_gunk_v1_salesReport_all_proto_rawDescData)
	})
	return file_codemen_org_inventory_gunk_v1_salesReport_all_proto_rawDescData
}

var (
	file_codemen_org_inventory_gunk_v1_salesReport_all_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
	file_codemen_org_inventory_gunk_v1_salesReport_all_proto_goTypes  = []interface{}{
		(*SalesReport)(nil),             // 0: salereportpb.SalesReport
		(*ListSalesReportRequest)(nil),  // 1: salereportpb.ListSalesReportRequest
		(*ListSalesReportResponse)(nil), // 2: salereportpb.ListSalesReportResponse
	}
)

var file_codemen_org_inventory_gunk_v1_salesReport_all_proto_depIdxs = []int32{
	0, // 0: salereportpb.ListSalesReportResponse.SalesReports:type_name -> salereportpb.SalesReport
	1, // 1: salereportpb.SalesService.ListSalesReport:input_type -> salereportpb.ListSalesReportRequest
	2, // 2: salereportpb.SalesService.ListSalesReport:output_type -> salereportpb.ListSalesReportResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_codemen_org_inventory_gunk_v1_salesReport_all_proto_init() }
func file_codemen_org_inventory_gunk_v1_salesReport_all_proto_init() {
	if File_codemen_org_inventory_gunk_v1_salesReport_all_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_codemen_org_inventory_gunk_v1_salesReport_all_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SalesReport); i {
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
		file_codemen_org_inventory_gunk_v1_salesReport_all_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSalesReportRequest); i {
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
		file_codemen_org_inventory_gunk_v1_salesReport_all_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSalesReportResponse); i {
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
			RawDescriptor: file_codemen_org_inventory_gunk_v1_salesReport_all_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_codemen_org_inventory_gunk_v1_salesReport_all_proto_goTypes,
		DependencyIndexes: file_codemen_org_inventory_gunk_v1_salesReport_all_proto_depIdxs,
		MessageInfos:      file_codemen_org_inventory_gunk_v1_salesReport_all_proto_msgTypes,
	}.Build()
	File_codemen_org_inventory_gunk_v1_salesReport_all_proto = out.File
	file_codemen_org_inventory_gunk_v1_salesReport_all_proto_rawDesc = nil
	file_codemen_org_inventory_gunk_v1_salesReport_all_proto_goTypes = nil
	file_codemen_org_inventory_gunk_v1_salesReport_all_proto_depIdxs = nil
}
