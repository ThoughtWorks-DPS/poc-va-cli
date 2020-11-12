package cmd

import (
	"github.com/golang/mock/gomock"
reflect "reflect"
)

// MockRestClient is a mock of RestClient interface
type MockRestClient struct {
	ctrl     *gomock.Controller
	recorder *MockRestClientMockRecorder
}

// MockRestClientMockRecorder is the mock recorder for MockRestClient
type MockRestClientMockRecorder struct {
	mock *MockRestClient
}

// NewMockRestClient creates a new mock instance
func NewMockRestClient(ctrl *gomock.Controller) *MockRestClient {
	mock := &MockRestClient{ctrl: ctrl}
	mock.recorder = &MockRestClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRestClient) EXPECT() *MockRestClientMockRecorder {
	return m.recorder
}

// getHello mocks base method
func (m *MockRestClient) getHello(url string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getHello", url)
	ret0, _ := ret[0].(string)
	return ret0
}

// getHello indicates an expected call of getHello
func (mr *MockRestClientMockRecorder) getHello(url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getHello", reflect.TypeOf((*MockRestClient)(nil).getHello), url)
}
