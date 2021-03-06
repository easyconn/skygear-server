// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/skygeario/skygear-server/pkg/server/skydb (interfaces: Database)

package mock_skydb

import (
	gomock "github.com/golang/mock/gomock"
	skydb "github.com/skygeario/skygear-server/pkg/server/skydb"
)

// MockDatabase is a mock of Database interface
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return _m.recorder
}

// Conn mocks base method
func (_m *MockDatabase) Conn() skydb.Conn {
	ret := _m.ctrl.Call(_m, "Conn")
	ret0, _ := ret[0].(skydb.Conn)
	return ret0
}

// Conn indicates an expected call of Conn
func (_mr *MockDatabaseMockRecorder) Conn() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Conn")
}

// DatabaseType mocks base method
func (_m *MockDatabase) DatabaseType() skydb.DatabaseType {
	ret := _m.ctrl.Call(_m, "DatabaseType")
	ret0, _ := ret[0].(skydb.DatabaseType)
	return ret0
}

// DatabaseType indicates an expected call of DatabaseType
func (_mr *MockDatabaseMockRecorder) DatabaseType() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DatabaseType")
}

// Delete mocks base method
func (_m *MockDatabase) Delete(_param0 skydb.RecordID) error {
	ret := _m.ctrl.Call(_m, "Delete", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (_mr *MockDatabaseMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}

// DeleteSchema mocks base method
func (_m *MockDatabase) DeleteSchema(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteSchema", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSchema indicates an expected call of DeleteSchema
func (_mr *MockDatabaseMockRecorder) DeleteSchema(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteSchema", arg0, arg1)
}

// DeleteSubscription mocks base method
func (_m *MockDatabase) DeleteSubscription(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteSubscription", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSubscription indicates an expected call of DeleteSubscription
func (_mr *MockDatabaseMockRecorder) DeleteSubscription(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteSubscription", arg0, arg1)
}

// Extend mocks base method
func (_m *MockDatabase) Extend(_param0 string, _param1 skydb.RecordSchema) (bool, error) {
	ret := _m.ctrl.Call(_m, "Extend", _param0, _param1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Extend indicates an expected call of Extend
func (_mr *MockDatabaseMockRecorder) Extend(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Extend", arg0, arg1)
}

// Get mocks base method
func (_m *MockDatabase) Get(_param0 skydb.RecordID, _param1 *skydb.Record) error {
	ret := _m.ctrl.Call(_m, "Get", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Get indicates an expected call of Get
func (_mr *MockDatabaseMockRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1)
}

// GetByIDs mocks base method
func (_m *MockDatabase) GetByIDs(_param0 []skydb.RecordID) (*skydb.Rows, error) {
	ret := _m.ctrl.Call(_m, "GetByIDs", _param0)
	ret0, _ := ret[0].(*skydb.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByIDs indicates an expected call of GetByIDs
func (_mr *MockDatabaseMockRecorder) GetByIDs(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetByIDs", arg0)
}

// GetMatchingSubscriptions mocks base method
func (_m *MockDatabase) GetMatchingSubscriptions(_param0 *skydb.Record) []skydb.Subscription {
	ret := _m.ctrl.Call(_m, "GetMatchingSubscriptions", _param0)
	ret0, _ := ret[0].([]skydb.Subscription)
	return ret0
}

// GetMatchingSubscriptions indicates an expected call of GetMatchingSubscriptions
func (_mr *MockDatabaseMockRecorder) GetMatchingSubscriptions(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetMatchingSubscriptions", arg0)
}

// GetRecordSchemas mocks base method
func (_m *MockDatabase) GetRecordSchemas() (map[string]skydb.RecordSchema, error) {
	ret := _m.ctrl.Call(_m, "GetRecordSchemas")
	ret0, _ := ret[0].(map[string]skydb.RecordSchema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRecordSchemas indicates an expected call of GetRecordSchemas
func (_mr *MockDatabaseMockRecorder) GetRecordSchemas() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetRecordSchemas")
}

// GetSchema mocks base method
func (_m *MockDatabase) GetSchema(_param0 string) (skydb.RecordSchema, error) {
	ret := _m.ctrl.Call(_m, "GetSchema", _param0)
	ret0, _ := ret[0].(skydb.RecordSchema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSchema indicates an expected call of GetSchema
func (_mr *MockDatabaseMockRecorder) GetSchema(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSchema", arg0)
}

// GetSubscription mocks base method
func (_m *MockDatabase) GetSubscription(_param0 string, _param1 string, _param2 *skydb.Subscription) error {
	ret := _m.ctrl.Call(_m, "GetSubscription", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// GetSubscription indicates an expected call of GetSubscription
func (_mr *MockDatabaseMockRecorder) GetSubscription(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSubscription", arg0, arg1, arg2)
}

// GetSubscriptionsByDeviceID mocks base method
func (_m *MockDatabase) GetSubscriptionsByDeviceID(_param0 string) []skydb.Subscription {
	ret := _m.ctrl.Call(_m, "GetSubscriptionsByDeviceID", _param0)
	ret0, _ := ret[0].([]skydb.Subscription)
	return ret0
}

// GetSubscriptionsByDeviceID indicates an expected call of GetSubscriptionsByDeviceID
func (_mr *MockDatabaseMockRecorder) GetSubscriptionsByDeviceID(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetSubscriptionsByDeviceID", arg0)
}

// ID mocks base method
func (_m *MockDatabase) ID() string {
	ret := _m.ctrl.Call(_m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID
func (_mr *MockDatabaseMockRecorder) ID() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ID")
}

// IsReadOnly mocks base method
func (_m *MockDatabase) IsReadOnly() bool {
	ret := _m.ctrl.Call(_m, "IsReadOnly")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsReadOnly indicates an expected call of IsReadOnly
func (_mr *MockDatabaseMockRecorder) IsReadOnly() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IsReadOnly")
}

// Query mocks base method
func (_m *MockDatabase) Query(_param0 *skydb.Query) (*skydb.Rows, error) {
	ret := _m.ctrl.Call(_m, "Query", _param0)
	ret0, _ := ret[0].(*skydb.Rows)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Query indicates an expected call of Query
func (_mr *MockDatabaseMockRecorder) Query(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Query", arg0)
}

// QueryCount mocks base method
func (_m *MockDatabase) QueryCount(_param0 *skydb.Query) (uint64, error) {
	ret := _m.ctrl.Call(_m, "QueryCount", _param0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// QueryCount indicates an expected call of QueryCount
func (_mr *MockDatabaseMockRecorder) QueryCount(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "QueryCount", arg0)
}

// RemoteColumnTypes mocks base method
func (_m *MockDatabase) RemoteColumnTypes(_param0 string) (skydb.RecordSchema, error) {
	ret := _m.ctrl.Call(_m, "RemoteColumnTypes", _param0)
	ret0, _ := ret[0].(skydb.RecordSchema)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoteColumnTypes indicates an expected call of RemoteColumnTypes
func (_mr *MockDatabaseMockRecorder) RemoteColumnTypes(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RemoteColumnTypes", arg0)
}

// RenameSchema mocks base method
func (_m *MockDatabase) RenameSchema(_param0 string, _param1 string, _param2 string) error {
	ret := _m.ctrl.Call(_m, "RenameSchema", _param0, _param1, _param2)
	ret0, _ := ret[0].(error)
	return ret0
}

// RenameSchema indicates an expected call of RenameSchema
func (_mr *MockDatabaseMockRecorder) RenameSchema(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RenameSchema", arg0, arg1, arg2)
}

// Save mocks base method
func (_m *MockDatabase) Save(_param0 *skydb.Record) error {
	ret := _m.ctrl.Call(_m, "Save", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (_mr *MockDatabaseMockRecorder) Save(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Save", arg0)
}

// SaveSubscription mocks base method
func (_m *MockDatabase) SaveSubscription(_param0 *skydb.Subscription) error {
	ret := _m.ctrl.Call(_m, "SaveSubscription", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveSubscription indicates an expected call of SaveSubscription
func (_mr *MockDatabaseMockRecorder) SaveSubscription(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "SaveSubscription", arg0)
}

// TableName mocks base method
func (_m *MockDatabase) TableName(_param0 string) string {
	ret := _m.ctrl.Call(_m, "TableName", _param0)
	ret0, _ := ret[0].(string)
	return ret0
}

// TableName indicates an expected call of TableName
func (_mr *MockDatabaseMockRecorder) TableName(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "TableName", arg0)
}

// UserRecordType mocks base method
func (_m *MockDatabase) UserRecordType() string {
	ret := _m.ctrl.Call(_m, "UserRecordType")
	ret0, _ := ret[0].(string)
	return ret0
}

// UserRecordType indicates an expected call of UserRecordType
func (_mr *MockDatabaseMockRecorder) UserRecordType() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UserRecordType")
}
