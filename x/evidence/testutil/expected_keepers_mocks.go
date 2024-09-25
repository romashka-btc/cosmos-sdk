// Code generated by MockGen. DO NOT EDIT.
// Source: x/evidence/types/expected_keepers.go

// Package testutil is a generated GoMock package.
package testutil

import (
	context "context"
	reflect "reflect"
	time "time"

	stakingv1beta1 "cosmossdk.io/api/cosmos/staking/v1beta1"
	math "cosmossdk.io/math"
	types "github.com/cosmos/cosmos-sdk/crypto/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockStakingKeeper is a mock of StakingKeeper interface.
type MockStakingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockStakingKeeperMockRecorder
}

// MockStakingKeeperMockRecorder is the mock recorder for MockStakingKeeper.
type MockStakingKeeperMockRecorder struct {
	mock *MockStakingKeeper
}

// NewMockStakingKeeper creates a new mock instance.
func NewMockStakingKeeper(ctrl *gomock.Controller) *MockStakingKeeper {
	mock := &MockStakingKeeper{ctrl: ctrl}
	mock.recorder = &MockStakingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStakingKeeper) EXPECT() *MockStakingKeeperMockRecorder {
	return m.recorder
}

// ValidatorByConsAddr mocks base method.
func (m *MockStakingKeeper) ValidatorByConsAddr(arg0 context.Context, arg1 types0.ConsAddress) (types0.ValidatorI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorByConsAddr", arg0, arg1)
	ret0, _ := ret[0].(types0.ValidatorI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ValidatorByConsAddr indicates an expected call of ValidatorByConsAddr.
func (mr *MockStakingKeeperMockRecorder) ValidatorByConsAddr(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorByConsAddr", reflect.TypeOf((*MockStakingKeeper)(nil).ValidatorByConsAddr), arg0, arg1)
}

// MockSlashingKeeper is a mock of SlashingKeeper interface.
type MockSlashingKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockSlashingKeeperMockRecorder
}

// MockSlashingKeeperMockRecorder is the mock recorder for MockSlashingKeeper.
type MockSlashingKeeperMockRecorder struct {
	mock *MockSlashingKeeper
}

// NewMockSlashingKeeper creates a new mock instance.
func NewMockSlashingKeeper(ctrl *gomock.Controller) *MockSlashingKeeper {
	mock := &MockSlashingKeeper{ctrl: ctrl}
	mock.recorder = &MockSlashingKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSlashingKeeper) EXPECT() *MockSlashingKeeperMockRecorder {
	return m.recorder
}

// GetPubkey mocks base method.
func (m *MockSlashingKeeper) GetPubkey(arg0 context.Context, arg1 types.Address) (types.PubKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPubkey", arg0, arg1)
	ret0, _ := ret[0].(types.PubKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPubkey indicates an expected call of GetPubkey.
func (mr *MockSlashingKeeperMockRecorder) GetPubkey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPubkey", reflect.TypeOf((*MockSlashingKeeper)(nil).GetPubkey), arg0, arg1)
}

// HasValidatorSigningInfo mocks base method.
func (m *MockSlashingKeeper) HasValidatorSigningInfo(arg0 context.Context, arg1 types0.ConsAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasValidatorSigningInfo", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasValidatorSigningInfo indicates an expected call of HasValidatorSigningInfo.
func (mr *MockSlashingKeeperMockRecorder) HasValidatorSigningInfo(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasValidatorSigningInfo", reflect.TypeOf((*MockSlashingKeeper)(nil).HasValidatorSigningInfo), arg0, arg1)
}

// IsTombstoned mocks base method.
func (m *MockSlashingKeeper) IsTombstoned(arg0 context.Context, arg1 types0.ConsAddress) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTombstoned", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTombstoned indicates an expected call of IsTombstoned.
func (mr *MockSlashingKeeperMockRecorder) IsTombstoned(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTombstoned", reflect.TypeOf((*MockSlashingKeeper)(nil).IsTombstoned), arg0, arg1)
}

// Jail mocks base method.
func (m *MockSlashingKeeper) Jail(arg0 context.Context, arg1 types0.ConsAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Jail", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Jail indicates an expected call of Jail.
func (mr *MockSlashingKeeperMockRecorder) Jail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Jail", reflect.TypeOf((*MockSlashingKeeper)(nil).Jail), arg0, arg1)
}

// JailUntil mocks base method.
func (m *MockSlashingKeeper) JailUntil(arg0 context.Context, arg1 types0.ConsAddress, arg2 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "JailUntil", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// JailUntil indicates an expected call of JailUntil.
func (mr *MockSlashingKeeperMockRecorder) JailUntil(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "JailUntil", reflect.TypeOf((*MockSlashingKeeper)(nil).JailUntil), arg0, arg1, arg2)
}

// Slash mocks base method.
func (m *MockSlashingKeeper) Slash(arg0 context.Context, arg1 types0.ConsAddress, arg2 math.LegacyDec, arg3, arg4 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Slash", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// Slash indicates an expected call of Slash.
func (mr *MockSlashingKeeperMockRecorder) Slash(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Slash", reflect.TypeOf((*MockSlashingKeeper)(nil).Slash), arg0, arg1, arg2, arg3, arg4)
}

// SlashFractionDoubleSign mocks base method.
func (m *MockSlashingKeeper) SlashFractionDoubleSign(arg0 context.Context) (math.LegacyDec, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SlashFractionDoubleSign", arg0)
	ret0, _ := ret[0].(math.LegacyDec)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SlashFractionDoubleSign indicates an expected call of SlashFractionDoubleSign.
func (mr *MockSlashingKeeperMockRecorder) SlashFractionDoubleSign(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SlashFractionDoubleSign", reflect.TypeOf((*MockSlashingKeeper)(nil).SlashFractionDoubleSign), arg0)
}

// SlashWithInfractionReason mocks base method.
func (m *MockSlashingKeeper) SlashWithInfractionReason(arg0 context.Context, arg1 types0.ConsAddress, arg2 math.LegacyDec, arg3, arg4 int64, arg5 stakingv1beta1.Infraction) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SlashWithInfractionReason", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// SlashWithInfractionReason indicates an expected call of SlashWithInfractionReason.
func (mr *MockSlashingKeeperMockRecorder) SlashWithInfractionReason(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SlashWithInfractionReason", reflect.TypeOf((*MockSlashingKeeper)(nil).SlashWithInfractionReason), arg0, arg1, arg2, arg3, arg4, arg5)
}

// Tombstone mocks base method.
func (m *MockSlashingKeeper) Tombstone(arg0 context.Context, arg1 types0.ConsAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tombstone", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Tombstone indicates an expected call of Tombstone.
func (mr *MockSlashingKeeperMockRecorder) Tombstone(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tombstone", reflect.TypeOf((*MockSlashingKeeper)(nil).Tombstone), arg0, arg1)
}

// MockAccountKeeper is a mock of AccountKeeper interface.
type MockAccountKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockAccountKeeperMockRecorder
}

// MockAccountKeeperMockRecorder is the mock recorder for MockAccountKeeper.
type MockAccountKeeperMockRecorder struct {
	mock *MockAccountKeeper
}

// NewMockAccountKeeper creates a new mock instance.
func NewMockAccountKeeper(ctrl *gomock.Controller) *MockAccountKeeper {
	mock := &MockAccountKeeper{ctrl: ctrl}
	mock.recorder = &MockAccountKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountKeeper) EXPECT() *MockAccountKeeperMockRecorder {
	return m.recorder
}

// SetAccount mocks base method.
func (m *MockAccountKeeper) SetAccount(ctx context.Context, acc types0.AccountI) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAccount", ctx, acc)
}

// SetAccount indicates an expected call of SetAccount.
func (mr *MockAccountKeeperMockRecorder) SetAccount(ctx, acc interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAccount", reflect.TypeOf((*MockAccountKeeper)(nil).SetAccount), ctx, acc)
}

// MockConsensusKeeper is a mock of ConsensusKeeper interface.
type MockConsensusKeeper struct {
	ctrl     *gomock.Controller
	recorder *MockConsensusKeeperMockRecorder
}

// MockConsensusKeeperMockRecorder is the mock recorder for MockConsensusKeeper.
type MockConsensusKeeperMockRecorder struct {
	mock *MockConsensusKeeper
}

// NewMockConsensusKeeper creates a new mock instance.
func NewMockConsensusKeeper(ctrl *gomock.Controller) *MockConsensusKeeper {
	mock := &MockConsensusKeeper{ctrl: ctrl}
	mock.recorder = &MockConsensusKeeperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockConsensusKeeper) EXPECT() *MockConsensusKeeperMockRecorder {
	return m.recorder
}

// EvidenceParams mocks base method.
func (m *MockConsensusKeeper) EvidenceParams(arg0 context.Context) (int64, time.Duration, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EvidenceParams", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(time.Duration)
	ret2, _ := ret[2].(uint64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// EvidenceParams indicates an expected call of EvidenceParams.
func (mr *MockConsensusKeeperMockRecorder) EvidenceParams(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvidenceParams", reflect.TypeOf((*MockConsensusKeeper)(nil).EvidenceParams), arg0)
}
