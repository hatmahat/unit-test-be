// Code generated by MockGen. DO NOT EDIT.
// Source: be-4-rs-crud/src/db/sqlc (interfaces: Store)

// Package mock_sqlc is a generated GoMock package.
package mock_sqlc

import (
	db "be-4-rs-crud/src/db/sqlc"
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// CreateAccount mocks base method.
func (m *MockStore) CreateAccount(arg0 context.Context, arg1 db.CreateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateAccount indicates an expected call of CreateAccount.
func (mr *MockStoreMockRecorder) CreateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAccount", reflect.TypeOf((*MockStore)(nil).CreateAccount), arg0, arg1)
}

// CreatePaymentName mocks base method.
func (m *MockStore) CreatePaymentName(arg0 context.Context, arg1 string) (db.PaymentBy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePaymentName", arg0, arg1)
	ret0, _ := ret[0].(db.PaymentBy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePaymentName indicates an expected call of CreatePaymentName.
func (mr *MockStoreMockRecorder) CreatePaymentName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePaymentName", reflect.TypeOf((*MockStore)(nil).CreatePaymentName), arg0, arg1)
}

// CreateTransaction mocks base method.
func (m *MockStore) CreateTransaction(arg0 context.Context, arg1 db.CreateTransactionParams) (db.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTransaction", arg0, arg1)
	ret0, _ := ret[0].(db.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTransaction indicates an expected call of CreateTransaction.
func (mr *MockStoreMockRecorder) CreateTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTransaction", reflect.TypeOf((*MockStore)(nil).CreateTransaction), arg0, arg1)
}

// DeleteAccount mocks base method.
func (m *MockStore) DeleteAccount(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAccount", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAccount indicates an expected call of DeleteAccount.
func (mr *MockStoreMockRecorder) DeleteAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAccount", reflect.TypeOf((*MockStore)(nil).DeleteAccount), arg0, arg1)
}

// DeletePaymenBy mocks base method.
func (m *MockStore) DeletePaymenBy(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePaymenBy", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePaymenBy indicates an expected call of DeletePaymenBy.
func (mr *MockStoreMockRecorder) DeletePaymenBy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePaymenBy", reflect.TypeOf((*MockStore)(nil).DeletePaymenBy), arg0, arg1)
}

// DeleteTransactionById mocks base method.
func (m *MockStore) DeleteTransactionById(arg0 context.Context, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTransactionById", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTransactionById indicates an expected call of DeleteTransactionById.
func (mr *MockStoreMockRecorder) DeleteTransactionById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTransactionById", reflect.TypeOf((*MockStore)(nil).DeleteTransactionById), arg0, arg1)
}

// GetAccountById mocks base method.
func (m *MockStore) GetAccountById(arg0 context.Context, arg1 int64) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountById", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountById indicates an expected call of GetAccountById.
func (mr *MockStoreMockRecorder) GetAccountById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountById", reflect.TypeOf((*MockStore)(nil).GetAccountById), arg0, arg1)
}

// GetPaymentBy mocks base method.
func (m *MockStore) GetPaymentBy(arg0 context.Context, arg1 db.GetPaymentByParams) ([]db.PaymentBy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentBy", arg0, arg1)
	ret0, _ := ret[0].([]db.PaymentBy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentBy indicates an expected call of GetPaymentBy.
func (mr *MockStoreMockRecorder) GetPaymentBy(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentBy", reflect.TypeOf((*MockStore)(nil).GetPaymentBy), arg0, arg1)
}

// GetPaymentById mocks base method.
func (m *MockStore) GetPaymentById(arg0 context.Context, arg1 int64) (db.PaymentBy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPaymentById", arg0, arg1)
	ret0, _ := ret[0].(db.PaymentBy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPaymentById indicates an expected call of GetPaymentById.
func (mr *MockStoreMockRecorder) GetPaymentById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPaymentById", reflect.TypeOf((*MockStore)(nil).GetPaymentById), arg0, arg1)
}

// GetTransactions mocks base method.
func (m *MockStore) GetTransactions(arg0 context.Context, arg1 db.GetTransactionsParams) ([]db.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions", arg0, arg1)
	ret0, _ := ret[0].([]db.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactions indicates an expected call of GetTransactions.
func (mr *MockStoreMockRecorder) GetTransactions(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockStore)(nil).GetTransactions), arg0, arg1)
}

// GetTransactionsById mocks base method.
func (m *MockStore) GetTransactionsById(arg0 context.Context, arg1 int64) (db.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactionsById", arg0, arg1)
	ret0, _ := ret[0].(db.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactionsById indicates an expected call of GetTransactionsById.
func (mr *MockStoreMockRecorder) GetTransactionsById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactionsById", reflect.TypeOf((*MockStore)(nil).GetTransactionsById), arg0, arg1)
}

// ListAccounts mocks base method.
func (m *MockStore) ListAccounts(arg0 context.Context, arg1 db.ListAccountsParams) ([]db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAccounts", arg0, arg1)
	ret0, _ := ret[0].([]db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAccounts indicates an expected call of ListAccounts.
func (mr *MockStoreMockRecorder) ListAccounts(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAccounts", reflect.TypeOf((*MockStore)(nil).ListAccounts), arg0, arg1)
}

// TransactionTx mocks base method.
func (m *MockStore) TransactionTx(arg0 context.Context, arg1 db.TransactionParams) (db.TransactionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionTx", arg0, arg1)
	ret0, _ := ret[0].(db.TransactionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionTx indicates an expected call of TransactionTx.
func (mr *MockStoreMockRecorder) TransactionTx(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionTx", reflect.TypeOf((*MockStore)(nil).TransactionTx), arg0, arg1)
}

// UpdateAccount mocks base method.
func (m *MockStore) UpdateAccount(arg0 context.Context, arg1 db.UpdateAccountParams) (db.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", arg0, arg1)
	ret0, _ := ret[0].(db.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockStoreMockRecorder) UpdateAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockStore)(nil).UpdateAccount), arg0, arg1)
}

// UpdatePaymentNameById mocks base method.
func (m *MockStore) UpdatePaymentNameById(arg0 context.Context, arg1 db.UpdatePaymentNameByIdParams) (db.PaymentBy, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePaymentNameById", arg0, arg1)
	ret0, _ := ret[0].(db.PaymentBy)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePaymentNameById indicates an expected call of UpdatePaymentNameById.
func (mr *MockStoreMockRecorder) UpdatePaymentNameById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePaymentNameById", reflect.TypeOf((*MockStore)(nil).UpdatePaymentNameById), arg0, arg1)
}

// UpdateTransactionById mocks base method.
func (m *MockStore) UpdateTransactionById(arg0 context.Context, arg1 db.UpdateTransactionByIdParams) (db.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTransactionById", arg0, arg1)
	ret0, _ := ret[0].(db.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTransactionById indicates an expected call of UpdateTransactionById.
func (mr *MockStoreMockRecorder) UpdateTransactionById(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTransactionById", reflect.TypeOf((*MockStore)(nil).UpdateTransactionById), arg0, arg1)
}
