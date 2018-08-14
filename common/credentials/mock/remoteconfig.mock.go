// Code generated by MockGen. DO NOT EDIT.
// Source: remoteconfig.go

// Package mock_credentials is a generated GoMock package.
package mock_credentials

import (
	gomock "github.com/golang/mock/gomock"
	consul "github.com/shankj3/go-til/consul"
	vault "github.com/shankj3/go-til/vault"
	pb "github.com/shankj3/ocelot/models/pb"
	storage "github.com/shankj3/ocelot/storage"
	reflect "reflect"
	"github.com/shankj3/ocelot/common/credentials"
)

// MockStorageCred is a mock of StorageCred interface
type MockStorageCred struct {
	ctrl     *gomock.Controller
	recorder *MockStorageCredMockRecorder
}

// MockStorageCredMockRecorder is the mock recorder for MockStorageCred
type MockStorageCredMockRecorder struct {
	mock *MockStorageCred
}

// NewMockStorageCred creates a new mock instance
func NewMockStorageCred(ctrl *gomock.Controller) *MockStorageCred {
	mock := &MockStorageCred{ctrl: ctrl}
	mock.recorder = &MockStorageCredMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockStorageCred) EXPECT() *MockStorageCredMockRecorder {
	return m.recorder
}

// GetStorageCreds mocks base method
func (m *MockStorageCred) GetStorageCreds(typ storage.Dest) (*credentials.StorageCreds, error) {
	ret := m.ctrl.Call(m, "GetStorageCreds", typ)
	ret0, _ := ret[0].(*credentials.StorageCreds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageCreds indicates an expected call of GetStorageCreds
func (mr *MockStorageCredMockRecorder) GetStorageCreds(typ interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageCreds", reflect.TypeOf((*MockStorageCred)(nil).GetStorageCreds), typ)
}

// GetStorageType mocks base method
func (m *MockStorageCred) GetStorageType() (storage.Dest, error) {
	ret := m.ctrl.Call(m, "GetStorageType")
	ret0, _ := ret[0].(storage.Dest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageType indicates an expected call of GetStorageType
func (mr *MockStorageCredMockRecorder) GetStorageType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageType", reflect.TypeOf((*MockStorageCred)(nil).GetStorageType))
}

// GetOcelotStorage mocks base method
func (m *MockStorageCred) GetOcelotStorage() (storage.OcelotStorage, error) {
	ret := m.ctrl.Call(m, "GetOcelotStorage")
	ret0, _ := ret[0].(storage.OcelotStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOcelotStorage indicates an expected call of GetOcelotStorage
func (mr *MockStorageCredMockRecorder) GetOcelotStorage() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOcelotStorage", reflect.TypeOf((*MockStorageCred)(nil).GetOcelotStorage))
}

// MockHealthyMaintainer is a mock of HealthyMaintainer interface
type MockHealthyMaintainer struct {
	ctrl     *gomock.Controller
	recorder *MockHealthyMaintainerMockRecorder
}

// MockHealthyMaintainerMockRecorder is the mock recorder for MockHealthyMaintainer
type MockHealthyMaintainerMockRecorder struct {
	mock *MockHealthyMaintainer
}

// NewMockHealthyMaintainer creates a new mock instance
func NewMockHealthyMaintainer(ctrl *gomock.Controller) *MockHealthyMaintainer {
	mock := &MockHealthyMaintainer{ctrl: ctrl}
	mock.recorder = &MockHealthyMaintainerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHealthyMaintainer) EXPECT() *MockHealthyMaintainerMockRecorder {
	return m.recorder
}

// Reconnect mocks base method
func (m *MockHealthyMaintainer) Reconnect() error {
	ret := m.ctrl.Call(m, "Reconnect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconnect indicates an expected call of Reconnect
func (mr *MockHealthyMaintainerMockRecorder) Reconnect() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconnect", reflect.TypeOf((*MockHealthyMaintainer)(nil).Reconnect))
}

// Healthy mocks base method
func (m *MockHealthyMaintainer) Healthy() bool {
	ret := m.ctrl.Call(m, "Healthy")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Healthy indicates an expected call of Healthy
func (mr *MockHealthyMaintainerMockRecorder) Healthy() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Healthy", reflect.TypeOf((*MockHealthyMaintainer)(nil).Healthy))
}

// MockCVRemoteConfig is a mock of CVRemoteConfig interface
type MockCVRemoteConfig struct {
	ctrl     *gomock.Controller
	recorder *MockCVRemoteConfigMockRecorder
}

// MockCVRemoteConfigMockRecorder is the mock recorder for MockCVRemoteConfig
type MockCVRemoteConfigMockRecorder struct {
	mock *MockCVRemoteConfig
}

// NewMockCVRemoteConfig creates a new mock instance
func NewMockCVRemoteConfig(ctrl *gomock.Controller) *MockCVRemoteConfig {
	mock := &MockCVRemoteConfig{ctrl: ctrl}
	mock.recorder = &MockCVRemoteConfigMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCVRemoteConfig) EXPECT() *MockCVRemoteConfigMockRecorder {
	return m.recorder
}

// GetConsul mocks base method
func (m *MockCVRemoteConfig) GetConsul() consul.Consuletty {
	ret := m.ctrl.Call(m, "GetConsul")
	ret0, _ := ret[0].(consul.Consuletty)
	return ret0
}

// GetConsul indicates an expected call of GetConsul
func (mr *MockCVRemoteConfigMockRecorder) GetConsul() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConsul", reflect.TypeOf((*MockCVRemoteConfig)(nil).GetConsul))
}

// SetConsul mocks base method
func (m *MockCVRemoteConfig) SetConsul(consul consul.Consuletty) {
	m.ctrl.Call(m, "SetConsul", consul)
}

// SetConsul indicates an expected call of SetConsul
func (mr *MockCVRemoteConfigMockRecorder) SetConsul(consul interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetConsul", reflect.TypeOf((*MockCVRemoteConfig)(nil).SetConsul), consul)
}

// GetVault mocks base method
func (m *MockCVRemoteConfig) GetVault() vault.Vaulty {
	ret := m.ctrl.Call(m, "GetVault")
	ret0, _ := ret[0].(vault.Vaulty)
	return ret0
}

// GetVault indicates an expected call of GetVault
func (mr *MockCVRemoteConfigMockRecorder) GetVault() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVault", reflect.TypeOf((*MockCVRemoteConfig)(nil).GetVault))
}

// SetVault mocks base method
func (m *MockCVRemoteConfig) SetVault(vault vault.Vaulty) {
	m.ctrl.Call(m, "SetVault", vault)
}

// SetVault indicates an expected call of SetVault
func (mr *MockCVRemoteConfigMockRecorder) SetVault(vault interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVault", reflect.TypeOf((*MockCVRemoteConfig)(nil).SetVault), vault)
}

// AddSSHKey mocks base method
func (m *MockCVRemoteConfig) AddSSHKey(path string, sshKeyFile []byte) error {
	ret := m.ctrl.Call(m, "AddSSHKey", path, sshKeyFile)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddSSHKey indicates an expected call of AddSSHKey
func (mr *MockCVRemoteConfigMockRecorder) AddSSHKey(path, sshKeyFile interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddSSHKey", reflect.TypeOf((*MockCVRemoteConfig)(nil).AddSSHKey), path, sshKeyFile)
}

// CheckSSHKeyExists mocks base method
func (m *MockCVRemoteConfig) CheckSSHKeyExists(path string) error {
	ret := m.ctrl.Call(m, "CheckSSHKeyExists", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckSSHKeyExists indicates an expected call of CheckSSHKeyExists
func (mr *MockCVRemoteConfigMockRecorder) CheckSSHKeyExists(path interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSSHKeyExists", reflect.TypeOf((*MockCVRemoteConfig)(nil).CheckSSHKeyExists), path)
}

// GetPassword mocks base method
func (m *MockCVRemoteConfig) GetPassword(scType pb.SubCredType, acctName string, ocyCredType pb.CredType, identifier string) (string, error) {
	ret := m.ctrl.Call(m, "GetPassword", scType, acctName, ocyCredType, identifier)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPassword indicates an expected call of GetPassword
func (mr *MockCVRemoteConfigMockRecorder) GetPassword(scType, acctName, ocyCredType, identifier interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPassword", reflect.TypeOf((*MockCVRemoteConfig)(nil).GetPassword), scType, acctName, ocyCredType, identifier)
}

// DeleteCred mocks base method
func (m *MockCVRemoteConfig) DeleteCred(store storage.CredTable, anyCred pb.OcyCredder) error {
	ret := m.ctrl.Call(m, "DeleteCred", store, anyCred)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteCred indicates an expected call of DeleteCred
func (mr *MockCVRemoteConfigMockRecorder) DeleteCred(store, anyCred interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCred", reflect.TypeOf((*MockCVRemoteConfig)(nil).DeleteCred), store, anyCred)
}

// GetCredsByType mocks base method
func (m *MockCVRemoteConfig) GetCredsByType(store storage.CredTable, ctype pb.CredType, hideSecret bool) ([]pb.OcyCredder, error) {
	ret := m.ctrl.Call(m, "GetCredsByType", store, ctype, hideSecret)
	ret0, _ := ret[0].([]pb.OcyCredder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredsByType indicates an expected call of GetCredsByType
func (mr *MockCVRemoteConfigMockRecorder) GetCredsByType(store, ctype, hideSecret interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredsByType", reflect.TypeOf((*MockCVRemoteConfig)(nil).GetCredsByType), store, ctype, hideSecret)
}

// GetAllCreds mocks base method
func (m *MockCVRemoteConfig) GetAllCreds(store storage.CredTable, hideSecret bool) ([]pb.OcyCredder, error) {
	ret := m.ctrl.Call(m, "GetAllCreds", store, hideSecret)
	ret0, _ := ret[0].([]pb.OcyCredder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCreds indicates an expected call of GetAllCreds
func (mr *MockCVRemoteConfigMockRecorder) GetAllCreds(store, hideSecret interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCreds", reflect.TypeOf((*MockCVRemoteConfig)(nil).GetAllCreds), store, hideSecret)
}

// GetCred mocks base method
func (m *MockCVRemoteConfig) GetCred(store storage.CredTable, subCredType pb.SubCredType, identifier, accountName string, hideSecret bool) (pb.OcyCredder, error) {
	ret := m.ctrl.Call(m, "GetCred", store, subCredType, identifier, accountName, hideSecret)
	ret0, _ := ret[0].(pb.OcyCredder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCred indicates an expected call of GetCred
func (mr *MockCVRemoteConfigMockRecorder) GetCred(store, subCredType, identifier, accountName, hideSecret interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCred", reflect.TypeOf((*MockCVRemoteConfig)(nil).GetCred), store, subCredType, identifier, accountName, hideSecret)
}

// GetCredsBySubTypeAndAcct mocks base method
func (m *MockCVRemoteConfig) GetCredsBySubTypeAndAcct(store storage.CredTable, stype pb.SubCredType, accountName string, hideSecret bool) ([]pb.OcyCredder, error) {
	ret := m.ctrl.Call(m, "GetCredsBySubTypeAndAcct", store, stype, accountName, hideSecret)
	ret0, _ := ret[0].([]pb.OcyCredder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredsBySubTypeAndAcct indicates an expected call of GetCredsBySubTypeAndAcct
func (mr *MockCVRemoteConfigMockRecorder) GetCredsBySubTypeAndAcct(store, stype, accountName, hideSecret interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredsBySubTypeAndAcct", reflect.TypeOf((*MockCVRemoteConfig)(nil).GetCredsBySubTypeAndAcct), store, stype, accountName, hideSecret)
}

// AddCreds mocks base method
func (m *MockCVRemoteConfig) AddCreds(store storage.CredTable, anyCred pb.OcyCredder, overwriteOk bool) error {
	ret := m.ctrl.Call(m, "AddCreds", store, anyCred, overwriteOk)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCreds indicates an expected call of AddCreds
func (mr *MockCVRemoteConfigMockRecorder) AddCreds(store, anyCred, overwriteOk interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCreds", reflect.TypeOf((*MockCVRemoteConfig)(nil).AddCreds), store, anyCred, overwriteOk)
}

// UpdateCreds mocks base method
func (m *MockCVRemoteConfig) UpdateCreds(store storage.CredTable, anyCred pb.OcyCredder) error {
	ret := m.ctrl.Call(m, "UpdateCreds", store, anyCred)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCreds indicates an expected call of UpdateCreds
func (mr *MockCVRemoteConfigMockRecorder) UpdateCreds(store, anyCred interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCreds", reflect.TypeOf((*MockCVRemoteConfig)(nil).UpdateCreds), store, anyCred)
}

// Reconnect mocks base method
func (m *MockCVRemoteConfig) Reconnect() error {
	ret := m.ctrl.Call(m, "Reconnect")
	ret0, _ := ret[0].(error)
	return ret0
}

// Reconnect indicates an expected call of Reconnect
func (mr *MockCVRemoteConfigMockRecorder) Reconnect() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconnect", reflect.TypeOf((*MockCVRemoteConfig)(nil).Reconnect))
}

// Healthy mocks base method
func (m *MockCVRemoteConfig) Healthy() bool {
	ret := m.ctrl.Call(m, "Healthy")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Healthy indicates an expected call of Healthy
func (mr *MockCVRemoteConfigMockRecorder) Healthy() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Healthy", reflect.TypeOf((*MockCVRemoteConfig)(nil).Healthy))
}

// GetStorageCreds mocks base method
func (m *MockCVRemoteConfig) GetStorageCreds(typ storage.Dest) (*credentials.StorageCreds, error) {
	ret := m.ctrl.Call(m, "GetStorageCreds", typ)
	ret0, _ := ret[0].(*credentials.StorageCreds)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageCreds indicates an expected call of GetStorageCreds
func (mr *MockCVRemoteConfigMockRecorder) GetStorageCreds(typ interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageCreds", reflect.TypeOf((*MockCVRemoteConfig)(nil).GetStorageCreds), typ)
}

// GetStorageType mocks base method
func (m *MockCVRemoteConfig) GetStorageType() (storage.Dest, error) {
	ret := m.ctrl.Call(m, "GetStorageType")
	ret0, _ := ret[0].(storage.Dest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStorageType indicates an expected call of GetStorageType
func (mr *MockCVRemoteConfigMockRecorder) GetStorageType() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStorageType", reflect.TypeOf((*MockCVRemoteConfig)(nil).GetStorageType))
}

// GetOcelotStorage mocks base method
func (m *MockCVRemoteConfig) GetOcelotStorage() (storage.OcelotStorage, error) {
	ret := m.ctrl.Call(m, "GetOcelotStorage")
	ret0, _ := ret[0].(storage.OcelotStorage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOcelotStorage indicates an expected call of GetOcelotStorage
func (mr *MockCVRemoteConfigMockRecorder) GetOcelotStorage() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOcelotStorage", reflect.TypeOf((*MockCVRemoteConfig)(nil).GetOcelotStorage))
}

// MockInsecureCredStorage is a mock of InsecureCredStorage interface
type MockInsecureCredStorage struct {
	ctrl     *gomock.Controller
	recorder *MockInsecureCredStorageMockRecorder
}

// MockInsecureCredStorageMockRecorder is the mock recorder for MockInsecureCredStorage
type MockInsecureCredStorageMockRecorder struct {
	mock *MockInsecureCredStorage
}

// NewMockInsecureCredStorage creates a new mock instance
func NewMockInsecureCredStorage(ctrl *gomock.Controller) *MockInsecureCredStorage {
	mock := &MockInsecureCredStorage{ctrl: ctrl}
	mock.recorder = &MockInsecureCredStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInsecureCredStorage) EXPECT() *MockInsecureCredStorageMockRecorder {
	return m.recorder
}

// GetCredsByType mocks base method
func (m *MockInsecureCredStorage) GetCredsByType(store storage.CredTable, ctype pb.CredType, hideSecret bool) ([]pb.OcyCredder, error) {
	ret := m.ctrl.Call(m, "GetCredsByType", store, ctype, hideSecret)
	ret0, _ := ret[0].([]pb.OcyCredder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredsByType indicates an expected call of GetCredsByType
func (mr *MockInsecureCredStorageMockRecorder) GetCredsByType(store, ctype, hideSecret interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredsByType", reflect.TypeOf((*MockInsecureCredStorage)(nil).GetCredsByType), store, ctype, hideSecret)
}

// GetAllCreds mocks base method
func (m *MockInsecureCredStorage) GetAllCreds(store storage.CredTable, hideSecret bool) ([]pb.OcyCredder, error) {
	ret := m.ctrl.Call(m, "GetAllCreds", store, hideSecret)
	ret0, _ := ret[0].([]pb.OcyCredder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCreds indicates an expected call of GetAllCreds
func (mr *MockInsecureCredStorageMockRecorder) GetAllCreds(store, hideSecret interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCreds", reflect.TypeOf((*MockInsecureCredStorage)(nil).GetAllCreds), store, hideSecret)
}

// GetCred mocks base method
func (m *MockInsecureCredStorage) GetCred(store storage.CredTable, subCredType pb.SubCredType, identifier, accountName string, hideSecret bool) (pb.OcyCredder, error) {
	ret := m.ctrl.Call(m, "GetCred", store, subCredType, identifier, accountName, hideSecret)
	ret0, _ := ret[0].(pb.OcyCredder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCred indicates an expected call of GetCred
func (mr *MockInsecureCredStorageMockRecorder) GetCred(store, subCredType, identifier, accountName, hideSecret interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCred", reflect.TypeOf((*MockInsecureCredStorage)(nil).GetCred), store, subCredType, identifier, accountName, hideSecret)
}

// GetCredsBySubTypeAndAcct mocks base method
func (m *MockInsecureCredStorage) GetCredsBySubTypeAndAcct(store storage.CredTable, stype pb.SubCredType, accountName string, hideSecret bool) ([]pb.OcyCredder, error) {
	ret := m.ctrl.Call(m, "GetCredsBySubTypeAndAcct", store, stype, accountName, hideSecret)
	ret0, _ := ret[0].([]pb.OcyCredder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCredsBySubTypeAndAcct indicates an expected call of GetCredsBySubTypeAndAcct
func (mr *MockInsecureCredStorageMockRecorder) GetCredsBySubTypeAndAcct(store, stype, accountName, hideSecret interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCredsBySubTypeAndAcct", reflect.TypeOf((*MockInsecureCredStorage)(nil).GetCredsBySubTypeAndAcct), store, stype, accountName, hideSecret)
}

// AddCreds mocks base method
func (m *MockInsecureCredStorage) AddCreds(store storage.CredTable, anyCred pb.OcyCredder, overwriteOk bool) error {
	ret := m.ctrl.Call(m, "AddCreds", store, anyCred, overwriteOk)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCreds indicates an expected call of AddCreds
func (mr *MockInsecureCredStorageMockRecorder) AddCreds(store, anyCred, overwriteOk interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCreds", reflect.TypeOf((*MockInsecureCredStorage)(nil).AddCreds), store, anyCred, overwriteOk)
}

// UpdateCreds mocks base method
func (m *MockInsecureCredStorage) UpdateCreds(store storage.CredTable, anyCred pb.OcyCredder) error {
	ret := m.ctrl.Call(m, "UpdateCreds", store, anyCred)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateCreds indicates an expected call of UpdateCreds
func (mr *MockInsecureCredStorageMockRecorder) UpdateCreds(store, anyCred interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCreds", reflect.TypeOf((*MockInsecureCredStorage)(nil).UpdateCreds), store, anyCred)
}
