// Code generated by MockGen. DO NOT EDIT.
// Source: repository/wallet_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	model "bootcamp/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIWalletRepository is a mock of IWalletRepository interface.
type MockIWalletRepository struct {
	ctrl     *gomock.Controller
	recorder *MockIWalletRepositoryMockRecorder
}

// MockIWalletRepositoryMockRecorder is the mock recorder for MockIWalletRepository.
type MockIWalletRepositoryMockRecorder struct {
	mock *MockIWalletRepository
}

// NewMockIWalletRepository creates a new mock instance.
func NewMockIWalletRepository(ctrl *gomock.Controller) *MockIWalletRepository {
	mock := &MockIWalletRepository{ctrl: ctrl}
	mock.recorder = &MockIWalletRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWalletRepository) EXPECT() *MockIWalletRepositoryMockRecorder {
	return m.recorder
}

// GetAllWallets mocks base method.
func (m *MockIWalletRepository) GetAllWallets() (model.WalletsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllWallets")
	ret0, _ := ret[0].(model.WalletsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllWallets indicates an expected call of GetAllWallets.
func (mr *MockIWalletRepositoryMockRecorder) GetAllWallets() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllWallets", reflect.TypeOf((*MockIWalletRepository)(nil).GetAllWallets))
}

// GetBalanceOfUser mocks base method.
func (m *MockIWalletRepository) GetBalanceOfUser(username string) (model.BalanceResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBalanceOfUser", username)
	ret0, _ := ret[0].(model.BalanceResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBalanceOfUser indicates an expected call of GetBalanceOfUser.
func (mr *MockIWalletRepositoryMockRecorder) GetBalanceOfUser(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBalanceOfUser", reflect.TypeOf((*MockIWalletRepository)(nil).GetBalanceOfUser), username)
}

// PostWallet mocks base method.
func (m *MockIWalletRepository) PostWallet(username string, balance int) (model.WalletResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostWallet", username, balance)
	ret0, _ := ret[0].(model.WalletResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostWallet indicates an expected call of PostWallet.
func (mr *MockIWalletRepositoryMockRecorder) PostWallet(username, balance interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostWallet", reflect.TypeOf((*MockIWalletRepository)(nil).PostWallet), username, balance)
}

// PutWallet mocks base method.
func (m *MockIWalletRepository) PutWallet(username string) (model.WalletResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PutWallet", username)
	ret0, _ := ret[0].(model.WalletResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PutWallet indicates an expected call of PutWallet.
func (mr *MockIWalletRepositoryMockRecorder) PutWallet(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PutWallet", reflect.TypeOf((*MockIWalletRepository)(nil).PutWallet), username)
}