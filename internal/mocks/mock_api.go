package mock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	apiToken "github.com/nikaydo/grpc-contract/gen/apiToken"
	grpc "google.golang.org/grpc"
)

// MockApiTokenClient is a mock of ApiTokenClient interface.
type MockApiTokenClient struct {
	ctrl     *gomock.Controller
	recorder *MockApiTokenClientMockRecorder
}

// MockApiTokenClientMockRecorder is the mock recorder for MockApiTokenClient.
type MockApiTokenClientMockRecorder struct {
	mock *MockApiTokenClient
}

// NewMockApiTokenClient creates a new mock instance.
func NewMockApiTokenClient(ctrl *gomock.Controller) *MockApiTokenClient {
	mock := &MockApiTokenClient{ctrl: ctrl}
	mock.recorder = &MockApiTokenClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApiTokenClient) EXPECT() *MockApiTokenClientMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockApiTokenClient) Create(arg0 context.Context, arg1 *apiToken.CreateRequest, arg2 ...grpc.CallOption) (*apiToken.CreateResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Create", varargs...)
	ret0, _ := ret[0].(*apiToken.CreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockApiTokenClientMockRecorder) Create(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockApiTokenClient)(nil).Create), varargs...)
}

// Delete mocks base method.
func (m *MockApiTokenClient) Delete(arg0 context.Context, arg1 *apiToken.DeleteRequest, arg2 ...grpc.CallOption) (*apiToken.DeleteResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Delete", varargs...)
	ret0, _ := ret[0].(*apiToken.DeleteResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete.
func (mr *MockApiTokenClientMockRecorder) Delete(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockApiTokenClient)(nil).Delete), varargs...)
}

// Get mocks base method.
func (m *MockApiTokenClient) Get(arg0 context.Context, arg1 *apiToken.GetRequest, arg2 ...grpc.CallOption) (*apiToken.GetResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Get", varargs...)
	ret0, _ := ret[0].(*apiToken.GetResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockApiTokenClientMockRecorder) Get(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockApiTokenClient)(nil).Get), varargs...)
}

// Verify mocks base method.
func (m *MockApiTokenClient) Verify(arg0 context.Context, arg1 *apiToken.VerifyRequest, arg2 ...grpc.CallOption) (*apiToken.VerifyResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Verify", varargs...)
	ret0, _ := ret[0].(*apiToken.VerifyResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Verify indicates an expected call of Verify.
func (mr *MockApiTokenClientMockRecorder) Verify(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Verify", reflect.TypeOf((*MockApiTokenClient)(nil).Verify), varargs...)
}
