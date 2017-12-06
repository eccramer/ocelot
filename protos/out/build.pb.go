// Code generated by protoc-gen-go. DO NOT EDIT.
// source: build.proto

/*
Package protos is a generated protocol buffer package.

It is generated from these files:
	build.proto
	common.proto
	commonevententities.proto
	projectrootdir.proto
	respositories.proto
	webhook.proto

It has these top-level messages:
	BuildConfig
	Stage
	PushBuildBundle
	PRBuildBundle
	LinkUrl
	LinkAndName
	Links
	Owner
	Repository
	PullRequestEntity
	PRInfo
	Project
	Changeset
	Commit
	RepoSourceFile
	PaginatedRootDirs
	PaginatedRepository
	RepoPush
	PullRequest
	PullRequestApproved
	CreateWebhook
	GetWebhooks
	Webhooks
*/
package protos

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

// this is a direct translation of the ocelot.yaml file
type BuildConfig struct {
	Image    string            `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Packages []string          `protobuf:"bytes,2,rep,name=packages" json:"packages,omitempty"`
	Env      map[string]string `protobuf:"bytes,3,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Before   *Stage            `protobuf:"bytes,4,opt,name=before" json:"before,omitempty"`
	After    *Stage            `protobuf:"bytes,5,opt,name=after" json:"after,omitempty"`
	Build    *Stage            `protobuf:"bytes,6,opt,name=build" json:"build,omitempty"`
	Test     *Stage            `protobuf:"bytes,7,opt,name=test" json:"test,omitempty"`
	Deploy   *Stage            `protobuf:"bytes,8,opt,name=deploy" json:"deploy,omitempty"`
}

func (m *BuildConfig) Reset()                    { *m = BuildConfig{} }
func (m *BuildConfig) String() string            { return proto.CompactTextString(m) }
func (*BuildConfig) ProtoMessage()               {}
func (*BuildConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BuildConfig) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *BuildConfig) GetPackages() []string {
	if m != nil {
		return m.Packages
	}
	return nil
}

func (m *BuildConfig) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *BuildConfig) GetBefore() *Stage {
	if m != nil {
		return m.Before
	}
	return nil
}

func (m *BuildConfig) GetAfter() *Stage {
	if m != nil {
		return m.After
	}
	return nil
}

func (m *BuildConfig) GetBuild() *Stage {
	if m != nil {
		return m.Build
	}
	return nil
}

func (m *BuildConfig) GetTest() *Stage {
	if m != nil {
		return m.Test
	}
	return nil
}

func (m *BuildConfig) GetDeploy() *Stage {
	if m != nil {
		return m.Deploy
	}
	return nil
}

type Stage struct {
	Env    map[string]string `protobuf:"bytes,1,rep,name=env" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Script []string          `protobuf:"bytes,2,rep,name=script" json:"script,omitempty"`
}

func (m *Stage) Reset()                    { *m = Stage{} }
func (m *Stage) String() string            { return proto.CompactTextString(m) }
func (*Stage) ProtoMessage()               {}
func (*Stage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Stage) GetEnv() map[string]string {
	if m != nil {
		return m.Env
	}
	return nil
}

func (m *Stage) GetScript() []string {
	if m != nil {
		return m.Script
	}
	return nil
}

type PushBuildBundle struct {
	Config       *BuildConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	PushData     *RepoPush    `protobuf:"bytes,2,opt,name=pushData" json:"pushData,omitempty"`
	VaultToken   string       `protobuf:"bytes,3,opt,name=vaultToken" json:"vaultToken,omitempty"`
	CheckoutHash string       `protobuf:"bytes,4,opt,name=checkoutHash" json:"checkoutHash,omitempty"`
}

func (m *PushBuildBundle) Reset()                    { *m = PushBuildBundle{} }
func (m *PushBuildBundle) String() string            { return proto.CompactTextString(m) }
func (*PushBuildBundle) ProtoMessage()               {}
func (*PushBuildBundle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PushBuildBundle) GetConfig() *BuildConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *PushBuildBundle) GetPushData() *RepoPush {
	if m != nil {
		return m.PushData
	}
	return nil
}

func (m *PushBuildBundle) GetVaultToken() string {
	if m != nil {
		return m.VaultToken
	}
	return ""
}

func (m *PushBuildBundle) GetCheckoutHash() string {
	if m != nil {
		return m.CheckoutHash
	}
	return ""
}

type PRBuildBundle struct {
	Config       *BuildConfig `protobuf:"bytes,1,opt,name=config" json:"config,omitempty"`
	PrData       *PullRequest `protobuf:"bytes,2,opt,name=prData" json:"prData,omitempty"`
	VaultToken   string       `protobuf:"bytes,3,opt,name=vaultToken" json:"vaultToken,omitempty"`
	CheckoutHash string       `protobuf:"bytes,4,opt,name=checkoutHash" json:"checkoutHash,omitempty"`
}

func (m *PRBuildBundle) Reset()                    { *m = PRBuildBundle{} }
func (m *PRBuildBundle) String() string            { return proto.CompactTextString(m) }
func (*PRBuildBundle) ProtoMessage()               {}
func (*PRBuildBundle) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *PRBuildBundle) GetConfig() *BuildConfig {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *PRBuildBundle) GetPrData() *PullRequest {
	if m != nil {
		return m.PrData
	}
	return nil
}

func (m *PRBuildBundle) GetVaultToken() string {
	if m != nil {
		return m.VaultToken
	}
	return ""
}

func (m *PRBuildBundle) GetCheckoutHash() string {
	if m != nil {
		return m.CheckoutHash
	}
	return ""
}

func init() {
	proto.RegisterType((*BuildConfig)(nil), "protos.BuildConfig")
	proto.RegisterType((*Stage)(nil), "protos.Stage")
	proto.RegisterType((*PushBuildBundle)(nil), "protos.PushBuildBundle")
	proto.RegisterType((*PRBuildBundle)(nil), "protos.PRBuildBundle")
}

func init() { proto.RegisterFile("build.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 410 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0xd1, 0x8a, 0xd3, 0x40,
	0x14, 0x86, 0x99, 0x66, 0x1b, 0xdb, 0x13, 0x8b, 0xcb, 0x28, 0xcb, 0x50, 0x44, 0x6a, 0x44, 0x08,
	0xac, 0xe4, 0x22, 0x82, 0x88, 0x97, 0xab, 0x0b, 0x5e, 0x96, 0xd1, 0x17, 0x98, 0xa4, 0xa7, 0x49,
	0x48, 0xcc, 0xc4, 0xcc, 0x4c, 0xa4, 0x97, 0x3e, 0x90, 0xe0, 0x33, 0xf9, 0x24, 0x92, 0x99, 0xb4,
	0x74, 0x25, 0x57, 0xc5, 0xab, 0xe4, 0x9c, 0xf3, 0xe5, 0xcf, 0x3f, 0xff, 0x49, 0x20, 0x48, 0x4d,
	0x59, 0xef, 0xe2, 0xb6, 0x93, 0x5a, 0x52, 0xdf, 0x5e, 0xd4, 0x7a, 0xf5, 0x03, 0xd3, 0x42, 0xca,
	0xca, 0xb5, 0xc3, 0x3f, 0x33, 0x08, 0xee, 0x06, 0xec, 0xa3, 0x6c, 0xf6, 0x65, 0x4e, 0x9f, 0xc1,
	0xbc, 0xfc, 0x26, 0x72, 0x64, 0x64, 0x43, 0xa2, 0x25, 0x77, 0x05, 0x5d, 0xc3, 0xa2, 0x15, 0x59,
	0x25, 0x72, 0x54, 0x6c, 0xb6, 0xf1, 0xa2, 0x25, 0x3f, 0xd5, 0x34, 0x06, 0x0f, 0x9b, 0x9e, 0x79,
	0x1b, 0x2f, 0x0a, 0x92, 0xe7, 0x4e, 0x56, 0xc5, 0x67, 0x9a, 0xf1, 0x7d, 0xd3, 0xdf, 0x37, 0xba,
	0x3b, 0xf0, 0x01, 0xa4, 0xaf, 0xc1, 0x4f, 0x71, 0x2f, 0x3b, 0x64, 0x57, 0x1b, 0x12, 0x05, 0xc9,
	0xea, 0xf8, 0xc8, 0x17, 0x2d, 0x72, 0xe4, 0xe3, 0x90, 0xbe, 0x82, 0xb9, 0xd8, 0x6b, 0xec, 0xd8,
	0x7c, 0x8a, 0x72, 0xb3, 0x01, 0xb2, 0x67, 0x64, 0xfe, 0x24, 0x64, 0x67, 0xf4, 0x25, 0x5c, 0x69,
	0x54, 0x9a, 0x3d, 0x9a, 0x62, 0xec, 0x68, 0xf0, 0xb4, 0xc3, 0xb6, 0x96, 0x07, 0xb6, 0x98, 0xf4,
	0xe4, 0x86, 0xeb, 0x77, 0xb0, 0x38, 0x9e, 0x85, 0x5e, 0x83, 0x57, 0xe1, 0x61, 0x8c, 0x69, 0xb8,
	0x1d, 0xa2, 0xeb, 0x45, 0x6d, 0x90, 0xcd, 0x5c, 0x74, 0xb6, 0xf8, 0x30, 0x7b, 0x4f, 0xc2, 0x9f,
	0x04, 0xe6, 0x56, 0x89, 0x46, 0x2e, 0x2c, 0x62, 0xc3, 0xba, 0x79, 0xf0, 0x96, 0x7f, 0x62, 0xba,
	0x01, 0x5f, 0x65, 0x5d, 0xd9, 0xea, 0x31, 0xf0, 0xb1, 0xba, 0xd8, 0xc3, 0x6f, 0x02, 0x4f, 0xb6,
	0x46, 0x15, 0x76, 0x31, 0x77, 0xa6, 0xd9, 0xd5, 0x48, 0x6f, 0xc1, 0xcf, 0xec, 0x8a, 0xac, 0x44,
	0x90, 0x3c, 0x9d, 0xd8, 0x1e, 0x1f, 0x11, 0xfa, 0x06, 0x16, 0xad, 0x51, 0xc5, 0x27, 0xa1, 0x85,
	0x55, 0x0f, 0x92, 0xeb, 0x23, 0xce, 0xb1, 0x95, 0x83, 0x36, 0x3f, 0x11, 0xf4, 0x05, 0x40, 0x2f,
	0x4c, 0xad, 0xbf, 0xca, 0x0a, 0x1b, 0xe6, 0x59, 0x37, 0x67, 0x1d, 0x1a, 0xc2, 0xe3, 0xac, 0xc0,
	0xac, 0x92, 0x46, 0x7f, 0x16, 0xaa, 0xb0, 0xdf, 0xc2, 0x92, 0x3f, 0xe8, 0x85, 0xbf, 0x08, 0xac,
	0xb6, 0xfc, 0x62, 0xc3, 0xb7, 0xe0, 0xb7, 0xdd, 0x99, 0xdd, 0x13, 0xbc, 0x35, 0x75, 0xcd, 0xf1,
	0xbb, 0x41, 0xa5, 0xf9, 0x88, 0xfc, 0x0f, 0xbf, 0xa9, 0xfb, 0xc5, 0xde, 0xfe, 0x0d, 0x00, 0x00,
	0xff, 0xff, 0x12, 0x72, 0xd0, 0x1b, 0x78, 0x03, 0x00, 0x00,
}
