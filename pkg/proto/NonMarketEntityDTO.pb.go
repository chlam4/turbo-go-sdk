// Code generated by protoc-gen-go. DO NOT EDIT.
// source: NonMarketEntityDTO.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Enumerate non market entities
type NonMarketEntityDTO_NonMarketEntityType int32

const (
	NonMarketEntityDTO_CLOUD_SERVICE NonMarketEntityDTO_NonMarketEntityType = 0
	NonMarketEntityDTO_WORKFLOW      NonMarketEntityDTO_NonMarketEntityType = 1
	NonMarketEntityDTO_ACCOUNT       NonMarketEntityDTO_NonMarketEntityType = 2
)

var NonMarketEntityDTO_NonMarketEntityType_name = map[int32]string{
	0: "CLOUD_SERVICE",
	1: "WORKFLOW",
	2: "ACCOUNT",
}

var NonMarketEntityDTO_NonMarketEntityType_value = map[string]int32{
	"CLOUD_SERVICE": 0,
	"WORKFLOW":      1,
	"ACCOUNT":       2,
}

func (x NonMarketEntityDTO_NonMarketEntityType) Enum() *NonMarketEntityDTO_NonMarketEntityType {
	p := new(NonMarketEntityDTO_NonMarketEntityType)
	*p = x
	return p
}

func (x NonMarketEntityDTO_NonMarketEntityType) String() string {
	return proto.EnumName(NonMarketEntityDTO_NonMarketEntityType_name, int32(x))
}

func (x *NonMarketEntityDTO_NonMarketEntityType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NonMarketEntityDTO_NonMarketEntityType_value, data, "NonMarketEntityDTO_NonMarketEntityType")
	if err != nil {
		return err
	}
	*x = NonMarketEntityDTO_NonMarketEntityType(value)
	return nil
}

func (NonMarketEntityDTO_NonMarketEntityType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 0}
}

type NonMarketEntityDTO_CloudServiceData_PriceModel int32

const (
	// Pay for compute capacity by per hour or per second depending
	// on which instances are run.
	NonMarketEntityDTO_CloudServiceData_ON_DEMAND NonMarketEntityDTO_CloudServiceData_PriceModel = 0
	// Pay a discounted price for reserved instances with a significant
	// discount (up to 75%) compared to On-Demand instance pricing.
	NonMarketEntityDTO_CloudServiceData_RESERVED NonMarketEntityDTO_CloudServiceData_PriceModel = 1
	// Pay the Spot price that is in effect for the time period the
	// instances are running and prices are set by Amazon EC2 and
	// adjusted gradually based on long-term trends in supply and demand.
	NonMarketEntityDTO_CloudServiceData_SPOT NonMarketEntityDTO_CloudServiceData_PriceModel = 2
)

var NonMarketEntityDTO_CloudServiceData_PriceModel_name = map[int32]string{
	0: "ON_DEMAND",
	1: "RESERVED",
	2: "SPOT",
}

var NonMarketEntityDTO_CloudServiceData_PriceModel_value = map[string]int32{
	"ON_DEMAND": 0,
	"RESERVED":  1,
	"SPOT":      2,
}

func (x NonMarketEntityDTO_CloudServiceData_PriceModel) Enum() *NonMarketEntityDTO_CloudServiceData_PriceModel {
	p := new(NonMarketEntityDTO_CloudServiceData_PriceModel)
	*p = x
	return p
}

func (x NonMarketEntityDTO_CloudServiceData_PriceModel) String() string {
	return proto.EnumName(NonMarketEntityDTO_CloudServiceData_PriceModel_name, int32(x))
}

func (x *NonMarketEntityDTO_CloudServiceData_PriceModel) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(NonMarketEntityDTO_CloudServiceData_PriceModel_value, data, "NonMarketEntityDTO_CloudServiceData_PriceModel")
	if err != nil {
		return err
	}
	*x = NonMarketEntityDTO_CloudServiceData_PriceModel(value)
	return nil
}

func (NonMarketEntityDTO_CloudServiceData_PriceModel) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 0, 0}
}

// The NonMarketEntityDTO message represents an Entity discovered in the target that your probe is
// monitoring and this entity does not participate in the market.
//
// Each entity must have a unique ID to identify it in the Operations Manager repository.
// Many targets provide unique IDs for their entities, or you can generate your own.
// To guarantee that it's unique, you can give the ID a prefix that identifies your
// probe and the given target.
//
// Specify entity type by setting a 'NonMarketEntityType' value to the 'entityType' field.
//
// The 'displayName' value appears in the product GUI and in reports to identify the entity.
//
type NonMarketEntityDTO struct {
	EntityType  *NonMarketEntityDTO_NonMarketEntityType `protobuf:"varint,1,req,name=entityType,enum=common_dto.NonMarketEntityDTO_NonMarketEntityType" json:"entityType,omitempty"`
	Id          *string                                 `protobuf:"bytes,2,req,name=id" json:"id,omitempty"`
	DisplayName *string                                 `protobuf:"bytes,3,opt,name=displayName" json:"displayName,omitempty"`
	Description *string                                 `protobuf:"bytes,4,opt,name=description" json:"description,omitempty"`
	// Collection of entity type's specific data.
	//
	// Types that are valid to be assigned to EntityData:
	//	*NonMarketEntityDTO_CloudServiceData_
	//	*NonMarketEntityDTO_WorkflowData_
	EntityData           isNonMarketEntityDTO_EntityData `protobuf_oneof:"entity_data"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *NonMarketEntityDTO) Reset()         { *m = NonMarketEntityDTO{} }
func (m *NonMarketEntityDTO) String() string { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO) ProtoMessage()    {}
func (*NonMarketEntityDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0}
}

func (m *NonMarketEntityDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonMarketEntityDTO.Unmarshal(m, b)
}
func (m *NonMarketEntityDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonMarketEntityDTO.Marshal(b, m, deterministic)
}
func (m *NonMarketEntityDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonMarketEntityDTO.Merge(m, src)
}
func (m *NonMarketEntityDTO) XXX_Size() int {
	return xxx_messageInfo_NonMarketEntityDTO.Size(m)
}
func (m *NonMarketEntityDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_NonMarketEntityDTO.DiscardUnknown(m)
}

var xxx_messageInfo_NonMarketEntityDTO proto.InternalMessageInfo

func (m *NonMarketEntityDTO) GetEntityType() NonMarketEntityDTO_NonMarketEntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return NonMarketEntityDTO_CLOUD_SERVICE
}

func (m *NonMarketEntityDTO) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *NonMarketEntityDTO) GetDisplayName() string {
	if m != nil && m.DisplayName != nil {
		return *m.DisplayName
	}
	return ""
}

func (m *NonMarketEntityDTO) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

type isNonMarketEntityDTO_EntityData interface {
	isNonMarketEntityDTO_EntityData()
}

type NonMarketEntityDTO_CloudServiceData_ struct {
	CloudServiceData *NonMarketEntityDTO_CloudServiceData `protobuf:"bytes,500,opt,name=cloud_service_data,json=cloudServiceData,oneof"`
}

type NonMarketEntityDTO_WorkflowData_ struct {
	WorkflowData *NonMarketEntityDTO_WorkflowData `protobuf:"bytes,501,opt,name=workflow_data,json=workflowData,oneof"`
}

func (*NonMarketEntityDTO_CloudServiceData_) isNonMarketEntityDTO_EntityData() {}

func (*NonMarketEntityDTO_WorkflowData_) isNonMarketEntityDTO_EntityData() {}

func (m *NonMarketEntityDTO) GetEntityData() isNonMarketEntityDTO_EntityData {
	if m != nil {
		return m.EntityData
	}
	return nil
}

func (m *NonMarketEntityDTO) GetCloudServiceData() *NonMarketEntityDTO_CloudServiceData {
	if x, ok := m.GetEntityData().(*NonMarketEntityDTO_CloudServiceData_); ok {
		return x.CloudServiceData
	}
	return nil
}

func (m *NonMarketEntityDTO) GetWorkflowData() *NonMarketEntityDTO_WorkflowData {
	if x, ok := m.GetEntityData().(*NonMarketEntityDTO_WorkflowData_); ok {
		return x.WorkflowData
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*NonMarketEntityDTO) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _NonMarketEntityDTO_OneofMarshaler, _NonMarketEntityDTO_OneofUnmarshaler, _NonMarketEntityDTO_OneofSizer, []interface{}{
		(*NonMarketEntityDTO_CloudServiceData_)(nil),
		(*NonMarketEntityDTO_WorkflowData_)(nil),
	}
}

func _NonMarketEntityDTO_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*NonMarketEntityDTO)
	// entity_data
	switch x := m.EntityData.(type) {
	case *NonMarketEntityDTO_CloudServiceData_:
		b.EncodeVarint(500<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.CloudServiceData); err != nil {
			return err
		}
	case *NonMarketEntityDTO_WorkflowData_:
		b.EncodeVarint(501<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.WorkflowData); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("NonMarketEntityDTO.EntityData has unexpected type %T", x)
	}
	return nil
}

func _NonMarketEntityDTO_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*NonMarketEntityDTO)
	switch tag {
	case 500: // entity_data.cloud_service_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NonMarketEntityDTO_CloudServiceData)
		err := b.DecodeMessage(msg)
		m.EntityData = &NonMarketEntityDTO_CloudServiceData_{msg}
		return true, err
	case 501: // entity_data.workflow_data
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(NonMarketEntityDTO_WorkflowData)
		err := b.DecodeMessage(msg)
		m.EntityData = &NonMarketEntityDTO_WorkflowData_{msg}
		return true, err
	default:
		return false, nil
	}
}

func _NonMarketEntityDTO_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*NonMarketEntityDTO)
	// entity_data
	switch x := m.EntityData.(type) {
	case *NonMarketEntityDTO_CloudServiceData_:
		s := proto.Size(x.CloudServiceData)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *NonMarketEntityDTO_WorkflowData_:
		s := proto.Size(x.WorkflowData)
		n += 2 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type NonMarketEntityDTO_CloudServiceData struct {
	CloudProvider        *string                                          `protobuf:"bytes,1,req,name=cloud_provider,json=cloudProvider" json:"cloud_provider,omitempty"`
	BillingData          *NonMarketEntityDTO_CloudServiceData_BillingData `protobuf:"bytes,3,opt,name=billing_data,json=billingData" json:"billing_data,omitempty"`
	AccountId            *string                                          `protobuf:"bytes,4,opt,name=account_id,json=accountId" json:"account_id,omitempty"`
	PriceModels          []NonMarketEntityDTO_CloudServiceData_PriceModel `protobuf:"varint,5,rep,name=price_models,json=priceModels,enum=common_dto.NonMarketEntityDTO_CloudServiceData_PriceModel" json:"price_models,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                         `json:"-"`
	XXX_unrecognized     []byte                                           `json:"-"`
	XXX_sizecache        int32                                            `json:"-"`
}

func (m *NonMarketEntityDTO_CloudServiceData) Reset()         { *m = NonMarketEntityDTO_CloudServiceData{} }
func (m *NonMarketEntityDTO_CloudServiceData) String() string { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO_CloudServiceData) ProtoMessage()    {}
func (*NonMarketEntityDTO_CloudServiceData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 0}
}

func (m *NonMarketEntityDTO_CloudServiceData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData.Unmarshal(m, b)
}
func (m *NonMarketEntityDTO_CloudServiceData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData.Marshal(b, m, deterministic)
}
func (m *NonMarketEntityDTO_CloudServiceData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonMarketEntityDTO_CloudServiceData.Merge(m, src)
}
func (m *NonMarketEntityDTO_CloudServiceData) XXX_Size() int {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData.Size(m)
}
func (m *NonMarketEntityDTO_CloudServiceData) XXX_DiscardUnknown() {
	xxx_messageInfo_NonMarketEntityDTO_CloudServiceData.DiscardUnknown(m)
}

var xxx_messageInfo_NonMarketEntityDTO_CloudServiceData proto.InternalMessageInfo

func (m *NonMarketEntityDTO_CloudServiceData) GetCloudProvider() string {
	if m != nil && m.CloudProvider != nil {
		return *m.CloudProvider
	}
	return ""
}

func (m *NonMarketEntityDTO_CloudServiceData) GetBillingData() *NonMarketEntityDTO_CloudServiceData_BillingData {
	if m != nil {
		return m.BillingData
	}
	return nil
}

func (m *NonMarketEntityDTO_CloudServiceData) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *NonMarketEntityDTO_CloudServiceData) GetPriceModels() []NonMarketEntityDTO_CloudServiceData_PriceModel {
	if m != nil {
		return m.PriceModels
	}
	return nil
}

type NonMarketEntityDTO_CloudServiceData_BillingData struct {
	VirtualMachines      []*EntityDTO `protobuf:"bytes,1,rep,name=virtual_machines,json=virtualMachines" json:"virtual_machines,omitempty"`
	ReservedInstances    []*EntityDTO `protobuf:"bytes,2,rep,name=reserved_instances,json=reservedInstances" json:"reserved_instances,omitempty"`
	BillingDataUuid      *string      `protobuf:"bytes,3,opt,name=billingDataUuid" json:"billingDataUuid,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) Reset() {
	*m = NonMarketEntityDTO_CloudServiceData_BillingData{}
}
func (m *NonMarketEntityDTO_CloudServiceData_BillingData) String() string {
	return proto.CompactTextString(m)
}
func (*NonMarketEntityDTO_CloudServiceData_BillingData) ProtoMessage() {}
func (*NonMarketEntityDTO_CloudServiceData_BillingData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 0, 0}
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData.Unmarshal(m, b)
}
func (m *NonMarketEntityDTO_CloudServiceData_BillingData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData.Marshal(b, m, deterministic)
}
func (m *NonMarketEntityDTO_CloudServiceData_BillingData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData.Merge(m, src)
}
func (m *NonMarketEntityDTO_CloudServiceData_BillingData) XXX_Size() int {
	return xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData.Size(m)
}
func (m *NonMarketEntityDTO_CloudServiceData_BillingData) XXX_DiscardUnknown() {
	xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData.DiscardUnknown(m)
}

var xxx_messageInfo_NonMarketEntityDTO_CloudServiceData_BillingData proto.InternalMessageInfo

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) GetVirtualMachines() []*EntityDTO {
	if m != nil {
		return m.VirtualMachines
	}
	return nil
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) GetReservedInstances() []*EntityDTO {
	if m != nil {
		return m.ReservedInstances
	}
	return nil
}

func (m *NonMarketEntityDTO_CloudServiceData_BillingData) GetBillingDataUuid() string {
	if m != nil && m.BillingDataUuid != nil {
		return *m.BillingDataUuid
	}
	return ""
}

type NonMarketEntityDTO_WorkflowData struct {
	Param                []*NonMarketEntityDTO_Parameter `protobuf:"bytes,1,rep,name=param" json:"param,omitempty"`
	Property             []*NonMarketEntityDTO_Property  `protobuf:"bytes,2,rep,name=property" json:"property,omitempty"`
	EntityType           *EntityDTO_EntityType           `protobuf:"varint,3,opt,name=entityType,enum=common_dto.EntityDTO_EntityType" json:"entityType,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *NonMarketEntityDTO_WorkflowData) Reset()         { *m = NonMarketEntityDTO_WorkflowData{} }
func (m *NonMarketEntityDTO_WorkflowData) String() string { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO_WorkflowData) ProtoMessage()    {}
func (*NonMarketEntityDTO_WorkflowData) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 1}
}

func (m *NonMarketEntityDTO_WorkflowData) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonMarketEntityDTO_WorkflowData.Unmarshal(m, b)
}
func (m *NonMarketEntityDTO_WorkflowData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonMarketEntityDTO_WorkflowData.Marshal(b, m, deterministic)
}
func (m *NonMarketEntityDTO_WorkflowData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonMarketEntityDTO_WorkflowData.Merge(m, src)
}
func (m *NonMarketEntityDTO_WorkflowData) XXX_Size() int {
	return xxx_messageInfo_NonMarketEntityDTO_WorkflowData.Size(m)
}
func (m *NonMarketEntityDTO_WorkflowData) XXX_DiscardUnknown() {
	xxx_messageInfo_NonMarketEntityDTO_WorkflowData.DiscardUnknown(m)
}

var xxx_messageInfo_NonMarketEntityDTO_WorkflowData proto.InternalMessageInfo

func (m *NonMarketEntityDTO_WorkflowData) GetParam() []*NonMarketEntityDTO_Parameter {
	if m != nil {
		return m.Param
	}
	return nil
}

func (m *NonMarketEntityDTO_WorkflowData) GetProperty() []*NonMarketEntityDTO_Property {
	if m != nil {
		return m.Property
	}
	return nil
}

func (m *NonMarketEntityDTO_WorkflowData) GetEntityType() EntityDTO_EntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return EntityDTO_SWITCH
}

type NonMarketEntityDTO_Parameter struct {
	Name                 *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Description          *string  `protobuf:"bytes,2,opt,name=description" json:"description,omitempty"`
	Type                 *string  `protobuf:"bytes,3,req,name=type,def=String" json:"type,omitempty"`
	Mandatory            *bool    `protobuf:"varint,4,opt,name=mandatory,def=1" json:"mandatory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NonMarketEntityDTO_Parameter) Reset()         { *m = NonMarketEntityDTO_Parameter{} }
func (m *NonMarketEntityDTO_Parameter) String() string { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO_Parameter) ProtoMessage()    {}
func (*NonMarketEntityDTO_Parameter) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 2}
}

func (m *NonMarketEntityDTO_Parameter) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonMarketEntityDTO_Parameter.Unmarshal(m, b)
}
func (m *NonMarketEntityDTO_Parameter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonMarketEntityDTO_Parameter.Marshal(b, m, deterministic)
}
func (m *NonMarketEntityDTO_Parameter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonMarketEntityDTO_Parameter.Merge(m, src)
}
func (m *NonMarketEntityDTO_Parameter) XXX_Size() int {
	return xxx_messageInfo_NonMarketEntityDTO_Parameter.Size(m)
}
func (m *NonMarketEntityDTO_Parameter) XXX_DiscardUnknown() {
	xxx_messageInfo_NonMarketEntityDTO_Parameter.DiscardUnknown(m)
}

var xxx_messageInfo_NonMarketEntityDTO_Parameter proto.InternalMessageInfo

const Default_NonMarketEntityDTO_Parameter_Type string = "String"
const Default_NonMarketEntityDTO_Parameter_Mandatory bool = true

func (m *NonMarketEntityDTO_Parameter) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *NonMarketEntityDTO_Parameter) GetDescription() string {
	if m != nil && m.Description != nil {
		return *m.Description
	}
	return ""
}

func (m *NonMarketEntityDTO_Parameter) GetType() string {
	if m != nil && m.Type != nil {
		return *m.Type
	}
	return Default_NonMarketEntityDTO_Parameter_Type
}

func (m *NonMarketEntityDTO_Parameter) GetMandatory() bool {
	if m != nil && m.Mandatory != nil {
		return *m.Mandatory
	}
	return Default_NonMarketEntityDTO_Parameter_Mandatory
}

type NonMarketEntityDTO_Property struct {
	Name                 *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value                *string  `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NonMarketEntityDTO_Property) Reset()         { *m = NonMarketEntityDTO_Property{} }
func (m *NonMarketEntityDTO_Property) String() string { return proto.CompactTextString(m) }
func (*NonMarketEntityDTO_Property) ProtoMessage()    {}
func (*NonMarketEntityDTO_Property) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{0, 3}
}

func (m *NonMarketEntityDTO_Property) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NonMarketEntityDTO_Property.Unmarshal(m, b)
}
func (m *NonMarketEntityDTO_Property) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NonMarketEntityDTO_Property.Marshal(b, m, deterministic)
}
func (m *NonMarketEntityDTO_Property) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NonMarketEntityDTO_Property.Merge(m, src)
}
func (m *NonMarketEntityDTO_Property) XXX_Size() int {
	return xxx_messageInfo_NonMarketEntityDTO_Property.Size(m)
}
func (m *NonMarketEntityDTO_Property) XXX_DiscardUnknown() {
	xxx_messageInfo_NonMarketEntityDTO_Property.DiscardUnknown(m)
}

var xxx_messageInfo_NonMarketEntityDTO_Property proto.InternalMessageInfo

func (m *NonMarketEntityDTO_Property) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *NonMarketEntityDTO_Property) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

// Generic cost data struct used for sending cost/spend from probe to the platform.
type CostDataDTO struct {
	// UUID of the entity (Account/CloudService) or profile (VMTemplate/DBTemplate)
	Id *string `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	// Whether the cost applies to an entity (default), or a profile.
	AppliesProfile *bool `protobuf:"varint,2,req,name=applies_profile,json=appliesProfile" json:"applies_profile,omitempty"`
	// Type of the cost data entity or the entity that the profile applies to.
	EntityType *EntityDTO_EntityType `protobuf:"varint,3,req,name=entity_type,json=entityType,enum=common_dto.EntityDTO_EntityType" json:"entity_type,omitempty"`
	// Business account id
	AccountId *string `protobuf:"bytes,4,req,name=account_id,json=accountId" json:"account_id,omitempty"`
	// Cost of the associated entity (e.g template/VM etc.)
	// If representing spend (top-down), then how much was spent for this account/service/template,
	// parsed from the bill.
	Cost                 *float32 `protobuf:"fixed32,5,req,name=cost" json:"cost,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CostDataDTO) Reset()         { *m = CostDataDTO{} }
func (m *CostDataDTO) String() string { return proto.CompactTextString(m) }
func (*CostDataDTO) ProtoMessage()    {}
func (*CostDataDTO) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7ef4e88daee47ae, []int{1}
}

func (m *CostDataDTO) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CostDataDTO.Unmarshal(m, b)
}
func (m *CostDataDTO) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CostDataDTO.Marshal(b, m, deterministic)
}
func (m *CostDataDTO) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CostDataDTO.Merge(m, src)
}
func (m *CostDataDTO) XXX_Size() int {
	return xxx_messageInfo_CostDataDTO.Size(m)
}
func (m *CostDataDTO) XXX_DiscardUnknown() {
	xxx_messageInfo_CostDataDTO.DiscardUnknown(m)
}

var xxx_messageInfo_CostDataDTO proto.InternalMessageInfo

func (m *CostDataDTO) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *CostDataDTO) GetAppliesProfile() bool {
	if m != nil && m.AppliesProfile != nil {
		return *m.AppliesProfile
	}
	return false
}

func (m *CostDataDTO) GetEntityType() EntityDTO_EntityType {
	if m != nil && m.EntityType != nil {
		return *m.EntityType
	}
	return EntityDTO_SWITCH
}

func (m *CostDataDTO) GetAccountId() string {
	if m != nil && m.AccountId != nil {
		return *m.AccountId
	}
	return ""
}

func (m *CostDataDTO) GetCost() float32 {
	if m != nil && m.Cost != nil {
		return *m.Cost
	}
	return 0
}

func init() {
	proto.RegisterEnum("common_dto.NonMarketEntityDTO_NonMarketEntityType", NonMarketEntityDTO_NonMarketEntityType_name, NonMarketEntityDTO_NonMarketEntityType_value)
	proto.RegisterEnum("common_dto.NonMarketEntityDTO_CloudServiceData_PriceModel", NonMarketEntityDTO_CloudServiceData_PriceModel_name, NonMarketEntityDTO_CloudServiceData_PriceModel_value)
	proto.RegisterType((*NonMarketEntityDTO)(nil), "common_dto.NonMarketEntityDTO")
	proto.RegisterType((*NonMarketEntityDTO_CloudServiceData)(nil), "common_dto.NonMarketEntityDTO.CloudServiceData")
	proto.RegisterType((*NonMarketEntityDTO_CloudServiceData_BillingData)(nil), "common_dto.NonMarketEntityDTO.CloudServiceData.BillingData")
	proto.RegisterType((*NonMarketEntityDTO_WorkflowData)(nil), "common_dto.NonMarketEntityDTO.WorkflowData")
	proto.RegisterType((*NonMarketEntityDTO_Parameter)(nil), "common_dto.NonMarketEntityDTO.Parameter")
	proto.RegisterType((*NonMarketEntityDTO_Property)(nil), "common_dto.NonMarketEntityDTO.Property")
	proto.RegisterType((*CostDataDTO)(nil), "common_dto.CostDataDTO")
}

func init() { proto.RegisterFile("NonMarketEntityDTO.proto", fileDescriptor_a7ef4e88daee47ae) }

var fileDescriptor_a7ef4e88daee47ae = []byte{
	// 757 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0x51, 0x6f, 0xe3, 0x44,
	0x10, 0x3e, 0xbb, 0x29, 0xc4, 0xe3, 0x24, 0xf5, 0x2d, 0x20, 0x59, 0x91, 0x10, 0x56, 0x24, 0x74,
	0x96, 0x90, 0x7c, 0x52, 0xe0, 0xa9, 0x48, 0xe8, 0x5a, 0x3b, 0x88, 0x8a, 0x4b, 0x1c, 0x6d, 0x52,
	0xfa, 0x04, 0x66, 0x6b, 0x6f, 0x8f, 0xd5, 0xd9, 0x5e, 0x6b, 0xbd, 0x49, 0x95, 0x47, 0xfe, 0x00,
	0xbf, 0x08, 0x89, 0x5f, 0xc1, 0x9f, 0xe0, 0x19, 0xde, 0x91, 0xbd, 0xae, 0x93, 0xa6, 0xd1, 0x45,
	0x7d, 0x1b, 0x7f, 0x9e, 0xf9, 0x66, 0xe6, 0xdb, 0x99, 0x01, 0x7b, 0xc6, 0xf3, 0x29, 0x11, 0xef,
	0xa9, 0x9c, 0xe4, 0x92, 0xc9, 0x4d, 0xb0, 0x0c, 0xbd, 0x42, 0x70, 0xc9, 0x11, 0xc4, 0x3c, 0xcb,
	0x78, 0x1e, 0x25, 0x92, 0x0f, 0xcf, 0xfc, 0xda, 0x6e, 0x7f, 0x8e, 0xfe, 0x01, 0x40, 0x4f, 0x23,
	0x11, 0x06, 0xa0, 0xf5, 0xc7, 0x72, 0x53, 0x50, 0x5b, 0x73, 0x74, 0x77, 0x30, 0x1e, 0x7b, 0x5b,
	0x22, 0xef, 0x40, 0xb6, 0x3d, 0xa8, 0x8a, 0xc4, 0x3b, 0x2c, 0x68, 0x00, 0x3a, 0x4b, 0x6c, 0xdd,
	0xd1, 0x5d, 0x03, 0xeb, 0x2c, 0x41, 0x0e, 0x98, 0x09, 0x2b, 0x8b, 0x94, 0x6c, 0x66, 0x24, 0xa3,
	0xf6, 0x89, 0xa3, 0xb9, 0x06, 0xde, 0x85, 0x6a, 0x0f, 0x5a, 0xc6, 0x82, 0x15, 0x92, 0xf1, 0xdc,
	0xee, 0x34, 0x1e, 0x5b, 0x08, 0xfd, 0x0a, 0x28, 0x4e, 0xf9, 0x2a, 0x89, 0x4a, 0x2a, 0xd6, 0x2c,
	0xa6, 0x51, 0x42, 0x24, 0xb1, 0xff, 0xad, 0xb8, 0xcc, 0xf1, 0xeb, 0x23, 0x05, 0xfb, 0x55, 0xe4,
	0x42, 0x05, 0x06, 0x44, 0x92, 0x1f, 0x5e, 0x60, 0x2b, 0xde, 0xc3, 0xd0, 0x02, 0xfa, 0xf7, 0x5c,
	0xbc, 0xbf, 0x4b, 0xf9, 0xbd, 0x22, 0xff, 0x4f, 0x91, 0x7f, 0x75, 0x84, 0xfc, 0xa6, 0x09, 0x6a,
	0x88, 0x7b, 0xf7, 0x3b, 0xdf, 0xc3, 0x3f, 0x3a, 0x60, 0xed, 0x67, 0x47, 0x5f, 0xc2, 0x40, 0xf5,
	0x52, 0x08, 0xbe, 0x66, 0x09, 0x15, 0xb5, 0xee, 0x06, 0xee, 0xd7, 0xe8, 0xbc, 0x01, 0xd1, 0x2f,
	0xd0, 0xbb, 0x65, 0x69, 0xca, 0xf2, 0x77, 0xaa, 0x1e, 0x55, 0xce, 0xb7, 0xcf, 0xec, 0xd5, 0xbb,
	0x54, 0x1c, 0x95, 0x8d, 0xcd, 0xdb, 0xed, 0x07, 0xfa, 0x1c, 0x80, 0xc4, 0x31, 0x5f, 0xe5, 0x32,
	0x62, 0x49, 0xa3, 0xb9, 0xd1, 0x20, 0x57, 0x09, 0xfa, 0x19, 0x7a, 0x85, 0xa8, 0x94, 0xce, 0x78,
	0x42, 0xd3, 0xd2, 0x3e, 0x75, 0x4e, 0xdc, 0xc1, 0xf8, 0xfc, 0xb9, 0xe9, 0xe7, 0x15, 0xc7, 0xb4,
	0xa2, 0xc0, 0x66, 0xd1, 0xda, 0xe5, 0xf0, 0x2f, 0x0d, 0xcc, 0x9d, 0xd2, 0xd0, 0x1b, 0xb0, 0xd6,
	0x4c, 0xc8, 0x15, 0x49, 0xa3, 0x8c, 0xc4, 0xbf, 0xb1, 0x9c, 0x96, 0xb6, 0xe6, 0x9c, 0xb8, 0xe6,
	0xf8, 0xb3, 0xdd, 0x94, 0x6d, 0x26, 0x7c, 0xd6, 0xb8, 0x4f, 0x1b, 0x6f, 0x14, 0x00, 0x12, 0xb4,
	0x1a, 0x0f, 0x9a, 0x44, 0x2c, 0x2f, 0x25, 0xc9, 0x63, 0x5a, 0xda, 0xfa, 0x87, 0x38, 0x5e, 0x3e,
	0x04, 0x5c, 0x3d, 0xf8, 0x23, 0x17, 0xce, 0x76, 0x44, 0xba, 0x5e, 0xb1, 0xa4, 0x19, 0xd8, 0x7d,
	0x78, 0xf4, 0x35, 0xc0, 0xb6, 0x39, 0xd4, 0x07, 0x23, 0x9c, 0x45, 0xc1, 0x64, 0x7a, 0x31, 0x0b,
	0xac, 0x17, 0xa8, 0x07, 0x5d, 0x3c, 0x59, 0x4c, 0xf0, 0x4f, 0x93, 0xc0, 0xd2, 0x50, 0x17, 0x3a,
	0x8b, 0x79, 0xb8, 0xb4, 0xf4, 0xe1, 0xdf, 0x1a, 0xf4, 0x76, 0x27, 0x06, 0x7d, 0x07, 0xa7, 0x05,
	0x11, 0x24, 0x6b, 0x9a, 0x75, 0x8f, 0xe8, 0x3b, 0xaf, 0x7c, 0xa9, 0xa4, 0x02, 0xab, 0x30, 0xe4,
	0x43, 0xb7, 0x10, 0xbc, 0xa0, 0x42, 0x6e, 0x9a, 0x5e, 0x5f, 0x1d, 0xa3, 0x68, 0xdc, 0x71, 0x1b,
	0x88, 0xde, 0x3c, 0xba, 0x02, 0x55, 0xbf, 0x83, 0xb1, 0x73, 0x50, 0x32, 0xef, 0xf0, 0xce, 0x0f,
	0x7f, 0xd7, 0xc0, 0x68, 0x6b, 0x43, 0x08, 0x3a, 0x79, 0xb5, 0xea, 0x6a, 0xae, 0x6b, 0x7b, 0x7f,
	0xc7, 0xf5, 0xa7, 0x3b, 0x3e, 0x84, 0x8e, 0x54, 0xf9, 0x75, 0xd7, 0x38, 0xff, 0x68, 0x21, 0x05,
	0xcb, 0xdf, 0xe1, 0x1a, 0x43, 0x23, 0x30, 0x32, 0x92, 0x27, 0x44, 0x72, 0xb1, 0xa9, 0x67, 0xb5,
	0x7b, 0xde, 0x91, 0x62, 0x45, 0xf1, 0x16, 0x1e, 0x7e, 0x03, 0xdd, 0x87, 0xde, 0x0e, 0x56, 0xf0,
	0x29, 0x9c, 0xae, 0x49, 0xba, 0xa2, 0xcd, 0x69, 0x52, 0x1f, 0x23, 0x1f, 0x3e, 0x39, 0x70, 0xd0,
	0xd0, 0x4b, 0xe8, 0xfb, 0x6f, 0xc3, 0xeb, 0x20, 0xaa, 0x1e, 0xf1, 0xca, 0x9f, 0xa8, 0x37, 0xbd,
	0x09, 0xf1, 0x8f, 0xdf, 0xbf, 0x0d, 0x6f, 0x2c, 0x0d, 0x99, 0xf0, 0xf1, 0x85, 0xef, 0x87, 0xd7,
	0xb3, 0xa5, 0xa5, 0x5f, 0xf6, 0xc1, 0x54, 0x62, 0xd4, 0xab, 0x3a, 0xfa, 0x53, 0x03, 0xd3, 0xe7,
	0xa5, 0xac, 0x5e, 0xb8, 0xba, 0xb2, 0xea, 0x22, 0x6a, 0xed, 0x45, 0x7c, 0x05, 0x67, 0xa4, 0x28,
	0x52, 0x46, 0xcb, 0xea, 0x06, 0xdc, 0xb1, 0x54, 0xd5, 0xd4, 0xc5, 0x83, 0x06, 0x9e, 0x2b, 0x14,
	0x5d, 0xb4, 0xbc, 0xad, 0x32, 0xcf, 0x7c, 0x99, 0x27, 0x6b, 0xae, 0x3f, 0x5e, 0x73, 0x04, 0x9d,
	0x98, 0x97, 0xd2, 0x3e, 0x75, 0x74, 0x57, 0xc7, 0xb5, 0x7d, 0xf9, 0x1a, 0xbe, 0x88, 0x79, 0xe6,
	0xad, 0x33, 0xb9, 0x12, 0xb7, 0xdc, 0x2b, 0x52, 0x22, 0xef, 0xb8, 0xc8, 0x9a, 0xb4, 0x5e, 0x22,
	0xf9, 0x65, 0xaf, 0xd5, 0x2c, 0x58, 0x86, 0xff, 0x07, 0x00, 0x00, 0xff, 0xff, 0x9a, 0xfe, 0xe3,
	0x46, 0x94, 0x06, 0x00, 0x00,
}
