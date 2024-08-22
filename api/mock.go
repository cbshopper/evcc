// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/evcc-io/evcc/api (interfaces: Charger,ChargeState,CurrentLimiter,CurrentGetter,PhaseSwitcher,PhaseGetter,Identifier,Meter,MeterEnergy,PhaseCurrents,Vehicle,ChargeRater,Battery,Tariff,BatteryController,Circuit)
//
// Generated by this command:
//
//	mockgen -package api -destination mock.go github.com/evcc-io/evcc/api Charger,ChargeState,CurrentLimiter,CurrentGetter,PhaseSwitcher,PhaseGetter,Identifier,Meter,MeterEnergy,PhaseCurrents,Vehicle,ChargeRater,Battery,Tariff,BatteryController,Circuit
//

// Package api is a generated GoMock package.
package api

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCharger is a mock of Charger interface.
type MockCharger struct {
	ctrl     *gomock.Controller
	recorder *MockChargerMockRecorder
}

// MockChargerMockRecorder is the mock recorder for MockCharger.
type MockChargerMockRecorder struct {
	mock *MockCharger
}

// NewMockCharger creates a new mock instance.
func NewMockCharger(ctrl *gomock.Controller) *MockCharger {
	mock := &MockCharger{ctrl: ctrl}
	mock.recorder = &MockChargerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCharger) EXPECT() *MockChargerMockRecorder {
	return m.recorder
}

// Enable mocks base method.
func (m *MockCharger) Enable(arg0 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enable", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enable indicates an expected call of Enable.
func (mr *MockChargerMockRecorder) Enable(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enable", reflect.TypeOf((*MockCharger)(nil).Enable), arg0)
}

// Enabled mocks base method.
func (m *MockCharger) Enabled() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Enabled indicates an expected call of Enabled.
func (mr *MockChargerMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockCharger)(nil).Enabled))
}

// MaxCurrent mocks base method.
func (m *MockCharger) MaxCurrent(arg0 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxCurrent", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// MaxCurrent indicates an expected call of MaxCurrent.
func (mr *MockChargerMockRecorder) MaxCurrent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxCurrent", reflect.TypeOf((*MockCharger)(nil).MaxCurrent), arg0)
}

// Status mocks base method.
func (m *MockCharger) Status() (ChargeStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(ChargeStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockChargerMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockCharger)(nil).Status))
}

// MockChargeState is a mock of ChargeState interface.
type MockChargeState struct {
	ctrl     *gomock.Controller
	recorder *MockChargeStateMockRecorder
}

// MockChargeStateMockRecorder is the mock recorder for MockChargeState.
type MockChargeStateMockRecorder struct {
	mock *MockChargeState
}

// NewMockChargeState creates a new mock instance.
func NewMockChargeState(ctrl *gomock.Controller) *MockChargeState {
	mock := &MockChargeState{ctrl: ctrl}
	mock.recorder = &MockChargeStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChargeState) EXPECT() *MockChargeStateMockRecorder {
	return m.recorder
}

// Status mocks base method.
func (m *MockChargeState) Status() (ChargeStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Status")
	ret0, _ := ret[0].(ChargeStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Status indicates an expected call of Status.
func (mr *MockChargeStateMockRecorder) Status() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Status", reflect.TypeOf((*MockChargeState)(nil).Status))
}

// MockCurrentLimiter is a mock of CurrentLimiter interface.
type MockCurrentLimiter struct {
	ctrl     *gomock.Controller
	recorder *MockCurrentLimiterMockRecorder
}

// MockCurrentLimiterMockRecorder is the mock recorder for MockCurrentLimiter.
type MockCurrentLimiterMockRecorder struct {
	mock *MockCurrentLimiter
}

// NewMockCurrentLimiter creates a new mock instance.
func NewMockCurrentLimiter(ctrl *gomock.Controller) *MockCurrentLimiter {
	mock := &MockCurrentLimiter{ctrl: ctrl}
	mock.recorder = &MockCurrentLimiterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCurrentLimiter) EXPECT() *MockCurrentLimiterMockRecorder {
	return m.recorder
}

// GetMinMaxCurrent mocks base method.
func (m *MockCurrentLimiter) GetMinMaxCurrent() (float64, float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMinMaxCurrent")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(float64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetMinMaxCurrent indicates an expected call of GetMinMaxCurrent.
func (mr *MockCurrentLimiterMockRecorder) GetMinMaxCurrent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMinMaxCurrent", reflect.TypeOf((*MockCurrentLimiter)(nil).GetMinMaxCurrent))
}

// MockCurrentGetter is a mock of CurrentGetter interface.
type MockCurrentGetter struct {
	ctrl     *gomock.Controller
	recorder *MockCurrentGetterMockRecorder
}

// MockCurrentGetterMockRecorder is the mock recorder for MockCurrentGetter.
type MockCurrentGetterMockRecorder struct {
	mock *MockCurrentGetter
}

// NewMockCurrentGetter creates a new mock instance.
func NewMockCurrentGetter(ctrl *gomock.Controller) *MockCurrentGetter {
	mock := &MockCurrentGetter{ctrl: ctrl}
	mock.recorder = &MockCurrentGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCurrentGetter) EXPECT() *MockCurrentGetterMockRecorder {
	return m.recorder
}

// GetMaxCurrent mocks base method.
func (m *MockCurrentGetter) GetMaxCurrent() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxCurrent")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMaxCurrent indicates an expected call of GetMaxCurrent.
func (mr *MockCurrentGetterMockRecorder) GetMaxCurrent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxCurrent", reflect.TypeOf((*MockCurrentGetter)(nil).GetMaxCurrent))
}

// MockPhaseSwitcher is a mock of PhaseSwitcher interface.
type MockPhaseSwitcher struct {
	ctrl     *gomock.Controller
	recorder *MockPhaseSwitcherMockRecorder
}

// MockPhaseSwitcherMockRecorder is the mock recorder for MockPhaseSwitcher.
type MockPhaseSwitcherMockRecorder struct {
	mock *MockPhaseSwitcher
}

// NewMockPhaseSwitcher creates a new mock instance.
func NewMockPhaseSwitcher(ctrl *gomock.Controller) *MockPhaseSwitcher {
	mock := &MockPhaseSwitcher{ctrl: ctrl}
	mock.recorder = &MockPhaseSwitcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhaseSwitcher) EXPECT() *MockPhaseSwitcherMockRecorder {
	return m.recorder
}

// Phases1p3p mocks base method.
func (m *MockPhaseSwitcher) Phases1p3p(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Phases1p3p", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Phases1p3p indicates an expected call of Phases1p3p.
func (mr *MockPhaseSwitcherMockRecorder) Phases1p3p(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Phases1p3p", reflect.TypeOf((*MockPhaseSwitcher)(nil).Phases1p3p), arg0)
}

// MockPhaseGetter is a mock of PhaseGetter interface.
type MockPhaseGetter struct {
	ctrl     *gomock.Controller
	recorder *MockPhaseGetterMockRecorder
}

// MockPhaseGetterMockRecorder is the mock recorder for MockPhaseGetter.
type MockPhaseGetterMockRecorder struct {
	mock *MockPhaseGetter
}

// NewMockPhaseGetter creates a new mock instance.
func NewMockPhaseGetter(ctrl *gomock.Controller) *MockPhaseGetter {
	mock := &MockPhaseGetter{ctrl: ctrl}
	mock.recorder = &MockPhaseGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhaseGetter) EXPECT() *MockPhaseGetterMockRecorder {
	return m.recorder
}

// GetPhases mocks base method.
func (m *MockPhaseGetter) GetPhases() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPhases")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPhases indicates an expected call of GetPhases.
func (mr *MockPhaseGetterMockRecorder) GetPhases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPhases", reflect.TypeOf((*MockPhaseGetter)(nil).GetPhases))
}

// MockIdentifier is a mock of Identifier interface.
type MockIdentifier struct {
	ctrl     *gomock.Controller
	recorder *MockIdentifierMockRecorder
}

// MockIdentifierMockRecorder is the mock recorder for MockIdentifier.
type MockIdentifierMockRecorder struct {
	mock *MockIdentifier
}

// NewMockIdentifier creates a new mock instance.
func NewMockIdentifier(ctrl *gomock.Controller) *MockIdentifier {
	mock := &MockIdentifier{ctrl: ctrl}
	mock.recorder = &MockIdentifierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIdentifier) EXPECT() *MockIdentifierMockRecorder {
	return m.recorder
}

// Identify mocks base method.
func (m *MockIdentifier) Identify() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identify")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Identify indicates an expected call of Identify.
func (mr *MockIdentifierMockRecorder) Identify() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identify", reflect.TypeOf((*MockIdentifier)(nil).Identify))
}

// MockMeter is a mock of Meter interface.
type MockMeter struct {
	ctrl     *gomock.Controller
	recorder *MockMeterMockRecorder
}

// MockMeterMockRecorder is the mock recorder for MockMeter.
type MockMeterMockRecorder struct {
	mock *MockMeter
}

// NewMockMeter creates a new mock instance.
func NewMockMeter(ctrl *gomock.Controller) *MockMeter {
	mock := &MockMeter{ctrl: ctrl}
	mock.recorder = &MockMeterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeter) EXPECT() *MockMeterMockRecorder {
	return m.recorder
}

// CurrentPower mocks base method.
func (m *MockMeter) CurrentPower() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentPower")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentPower indicates an expected call of CurrentPower.
func (mr *MockMeterMockRecorder) CurrentPower() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentPower", reflect.TypeOf((*MockMeter)(nil).CurrentPower))
}

// MockMeterEnergy is a mock of MeterEnergy interface.
type MockMeterEnergy struct {
	ctrl     *gomock.Controller
	recorder *MockMeterEnergyMockRecorder
}

// MockMeterEnergyMockRecorder is the mock recorder for MockMeterEnergy.
type MockMeterEnergyMockRecorder struct {
	mock *MockMeterEnergy
}

// NewMockMeterEnergy creates a new mock instance.
func NewMockMeterEnergy(ctrl *gomock.Controller) *MockMeterEnergy {
	mock := &MockMeterEnergy{ctrl: ctrl}
	mock.recorder = &MockMeterEnergyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMeterEnergy) EXPECT() *MockMeterEnergyMockRecorder {
	return m.recorder
}

// TotalEnergy mocks base method.
func (m *MockMeterEnergy) TotalEnergy() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TotalEnergy")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TotalEnergy indicates an expected call of TotalEnergy.
func (mr *MockMeterEnergyMockRecorder) TotalEnergy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TotalEnergy", reflect.TypeOf((*MockMeterEnergy)(nil).TotalEnergy))
}

// MockPhaseCurrents is a mock of PhaseCurrents interface.
type MockPhaseCurrents struct {
	ctrl     *gomock.Controller
	recorder *MockPhaseCurrentsMockRecorder
}

// MockPhaseCurrentsMockRecorder is the mock recorder for MockPhaseCurrents.
type MockPhaseCurrentsMockRecorder struct {
	mock *MockPhaseCurrents
}

// NewMockPhaseCurrents creates a new mock instance.
func NewMockPhaseCurrents(ctrl *gomock.Controller) *MockPhaseCurrents {
	mock := &MockPhaseCurrents{ctrl: ctrl}
	mock.recorder = &MockPhaseCurrentsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPhaseCurrents) EXPECT() *MockPhaseCurrentsMockRecorder {
	return m.recorder
}

// Currents mocks base method.
func (m *MockPhaseCurrents) Currents() (float64, float64, float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Currents")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(float64)
	ret2, _ := ret[2].(float64)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Currents indicates an expected call of Currents.
func (mr *MockPhaseCurrentsMockRecorder) Currents() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Currents", reflect.TypeOf((*MockPhaseCurrents)(nil).Currents))
}

// MockVehicle is a mock of Vehicle interface.
type MockVehicle struct {
	ctrl     *gomock.Controller
	recorder *MockVehicleMockRecorder
}

// MockVehicleMockRecorder is the mock recorder for MockVehicle.
type MockVehicleMockRecorder struct {
	mock *MockVehicle
}

// NewMockVehicle creates a new mock instance.
func NewMockVehicle(ctrl *gomock.Controller) *MockVehicle {
	mock := &MockVehicle{ctrl: ctrl}
	mock.recorder = &MockVehicleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVehicle) EXPECT() *MockVehicleMockRecorder {
	return m.recorder
}

// Capacity mocks base method.
func (m *MockVehicle) Capacity() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Capacity")
	ret0, _ := ret[0].(float64)
	return ret0
}

// Capacity indicates an expected call of Capacity.
func (mr *MockVehicleMockRecorder) Capacity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Capacity", reflect.TypeOf((*MockVehicle)(nil).Capacity))
}

// Features mocks base method.
func (m *MockVehicle) Features() []Feature {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Features")
	ret0, _ := ret[0].([]Feature)
	return ret0
}

// Features indicates an expected call of Features.
func (mr *MockVehicleMockRecorder) Features() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Features", reflect.TypeOf((*MockVehicle)(nil).Features))
}

// Icon mocks base method.
func (m *MockVehicle) Icon() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Icon")
	ret0, _ := ret[0].(string)
	return ret0
}

// Icon indicates an expected call of Icon.
func (mr *MockVehicleMockRecorder) Icon() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Icon", reflect.TypeOf((*MockVehicle)(nil).Icon))
}

// Identifiers mocks base method.
func (m *MockVehicle) Identifiers() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identifiers")
	ret0, _ := ret[0].([]string)
	return ret0
}

// Identifiers indicates an expected call of Identifiers.
func (mr *MockVehicleMockRecorder) Identifiers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identifiers", reflect.TypeOf((*MockVehicle)(nil).Identifiers))
}

// OnIdentified mocks base method.
func (m *MockVehicle) OnIdentified() ActionConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OnIdentified")
	ret0, _ := ret[0].(ActionConfig)
	return ret0
}

// OnIdentified indicates an expected call of OnIdentified.
func (mr *MockVehicleMockRecorder) OnIdentified() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnIdentified", reflect.TypeOf((*MockVehicle)(nil).OnIdentified))
}

// Phases mocks base method.
func (m *MockVehicle) Phases() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Phases")
	ret0, _ := ret[0].(int)
	return ret0
}

// Phases indicates an expected call of Phases.
func (mr *MockVehicleMockRecorder) Phases() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Phases", reflect.TypeOf((*MockVehicle)(nil).Phases))
}

// SetTitle mocks base method.
func (m *MockVehicle) SetTitle(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTitle", arg0)
}

// SetTitle indicates an expected call of SetTitle.
func (mr *MockVehicleMockRecorder) SetTitle(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTitle", reflect.TypeOf((*MockVehicle)(nil).SetTitle), arg0)
}

// Soc mocks base method.
func (m *MockVehicle) Soc() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Soc")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Soc indicates an expected call of Soc.
func (mr *MockVehicleMockRecorder) Soc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Soc", reflect.TypeOf((*MockVehicle)(nil).Soc))
}

// Title mocks base method.
func (m *MockVehicle) Title() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Title")
	ret0, _ := ret[0].(string)
	return ret0
}

// Title indicates an expected call of Title.
func (mr *MockVehicleMockRecorder) Title() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Title", reflect.TypeOf((*MockVehicle)(nil).Title))
}

// MockChargeRater is a mock of ChargeRater interface.
type MockChargeRater struct {
	ctrl     *gomock.Controller
	recorder *MockChargeRaterMockRecorder
}

// MockChargeRaterMockRecorder is the mock recorder for MockChargeRater.
type MockChargeRaterMockRecorder struct {
	mock *MockChargeRater
}

// NewMockChargeRater creates a new mock instance.
func NewMockChargeRater(ctrl *gomock.Controller) *MockChargeRater {
	mock := &MockChargeRater{ctrl: ctrl}
	mock.recorder = &MockChargeRaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChargeRater) EXPECT() *MockChargeRaterMockRecorder {
	return m.recorder
}

// ChargedEnergy mocks base method.
func (m *MockChargeRater) ChargedEnergy() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ChargedEnergy")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ChargedEnergy indicates an expected call of ChargedEnergy.
func (mr *MockChargeRaterMockRecorder) ChargedEnergy() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ChargedEnergy", reflect.TypeOf((*MockChargeRater)(nil).ChargedEnergy))
}

// MockBattery is a mock of Battery interface.
type MockBattery struct {
	ctrl     *gomock.Controller
	recorder *MockBatteryMockRecorder
}

// MockBatteryMockRecorder is the mock recorder for MockBattery.
type MockBatteryMockRecorder struct {
	mock *MockBattery
}

// NewMockBattery creates a new mock instance.
func NewMockBattery(ctrl *gomock.Controller) *MockBattery {
	mock := &MockBattery{ctrl: ctrl}
	mock.recorder = &MockBatteryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBattery) EXPECT() *MockBatteryMockRecorder {
	return m.recorder
}

// Soc mocks base method.
func (m *MockBattery) Soc() (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Soc")
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Soc indicates an expected call of Soc.
func (mr *MockBatteryMockRecorder) Soc() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Soc", reflect.TypeOf((*MockBattery)(nil).Soc))
}

// MockTariff is a mock of Tariff interface.
type MockTariff struct {
	ctrl     *gomock.Controller
	recorder *MockTariffMockRecorder
}

// MockTariffMockRecorder is the mock recorder for MockTariff.
type MockTariffMockRecorder struct {
	mock *MockTariff
}

// NewMockTariff creates a new mock instance.
func NewMockTariff(ctrl *gomock.Controller) *MockTariff {
	mock := &MockTariff{ctrl: ctrl}
	mock.recorder = &MockTariffMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTariff) EXPECT() *MockTariffMockRecorder {
	return m.recorder
}

// Rates mocks base method.
func (m *MockTariff) Rates() (Rates, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rates")
	ret0, _ := ret[0].(Rates)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Rates indicates an expected call of Rates.
func (mr *MockTariffMockRecorder) Rates() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rates", reflect.TypeOf((*MockTariff)(nil).Rates))
}

// Type mocks base method.
func (m *MockTariff) Type() TariffType {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Type")
	ret0, _ := ret[0].(TariffType)
	return ret0
}

// Type indicates an expected call of Type.
func (mr *MockTariffMockRecorder) Type() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Type", reflect.TypeOf((*MockTariff)(nil).Type))
}

// MockBatteryController is a mock of BatteryController interface.
type MockBatteryController struct {
	ctrl     *gomock.Controller
	recorder *MockBatteryControllerMockRecorder
}

// MockBatteryControllerMockRecorder is the mock recorder for MockBatteryController.
type MockBatteryControllerMockRecorder struct {
	mock *MockBatteryController
}

// NewMockBatteryController creates a new mock instance.
func NewMockBatteryController(ctrl *gomock.Controller) *MockBatteryController {
	mock := &MockBatteryController{ctrl: ctrl}
	mock.recorder = &MockBatteryControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBatteryController) EXPECT() *MockBatteryControllerMockRecorder {
	return m.recorder
}

// SetBatteryMode mocks base method.
func (m *MockBatteryController) SetBatteryMode(arg0 BatteryMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetBatteryMode", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetBatteryMode indicates an expected call of SetBatteryMode.
func (mr *MockBatteryControllerMockRecorder) SetBatteryMode(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBatteryMode", reflect.TypeOf((*MockBatteryController)(nil).SetBatteryMode), arg0)
}

// MockCircuit is a mock of Circuit interface.
type MockCircuit struct {
	ctrl     *gomock.Controller
	recorder *MockCircuitMockRecorder
}

// MockCircuitMockRecorder is the mock recorder for MockCircuit.
type MockCircuitMockRecorder struct {
	mock *MockCircuit
}

// NewMockCircuit creates a new mock instance.
func NewMockCircuit(ctrl *gomock.Controller) *MockCircuit {
	mock := &MockCircuit{ctrl: ctrl}
	mock.recorder = &MockCircuitMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCircuit) EXPECT() *MockCircuitMockRecorder {
	return m.recorder
}

// GetChargePower mocks base method.
func (m *MockCircuit) GetChargePower() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChargePower")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetChargePower indicates an expected call of GetChargePower.
func (mr *MockCircuitMockRecorder) GetChargePower() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChargePower", reflect.TypeOf((*MockCircuit)(nil).GetChargePower))
}

// GetMaxCurrent mocks base method.
func (m *MockCircuit) GetMaxCurrent() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxCurrent")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetMaxCurrent indicates an expected call of GetMaxCurrent.
func (mr *MockCircuitMockRecorder) GetMaxCurrent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxCurrent", reflect.TypeOf((*MockCircuit)(nil).GetMaxCurrent))
}

// GetMaxPhaseCurrent mocks base method.
func (m *MockCircuit) GetMaxPhaseCurrent() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxPhaseCurrent")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetMaxPhaseCurrent indicates an expected call of GetMaxPhaseCurrent.
func (mr *MockCircuitMockRecorder) GetMaxPhaseCurrent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxPhaseCurrent", reflect.TypeOf((*MockCircuit)(nil).GetMaxPhaseCurrent))
}

// GetMaxPower mocks base method.
func (m *MockCircuit) GetMaxPower() float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMaxPower")
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetMaxPower indicates an expected call of GetMaxPower.
func (mr *MockCircuitMockRecorder) GetMaxPower() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMaxPower", reflect.TypeOf((*MockCircuit)(nil).GetMaxPower))
}

// GetParent mocks base method.
func (m *MockCircuit) GetParent() Circuit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetParent")
	ret0, _ := ret[0].(Circuit)
	return ret0
}

// GetParent indicates an expected call of GetParent.
func (mr *MockCircuitMockRecorder) GetParent() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetParent", reflect.TypeOf((*MockCircuit)(nil).GetParent))
}

// GetTitle mocks base method.
func (m *MockCircuit) GetTitle() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTitle")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetTitle indicates an expected call of GetTitle.
func (mr *MockCircuitMockRecorder) GetTitle() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTitle", reflect.TypeOf((*MockCircuit)(nil).GetTitle))
}

// HasMeter mocks base method.
func (m *MockCircuit) HasMeter() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasMeter")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasMeter indicates an expected call of HasMeter.
func (mr *MockCircuitMockRecorder) HasMeter() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasMeter", reflect.TypeOf((*MockCircuit)(nil).HasMeter))
}

// RegisterChild mocks base method.
func (m *MockCircuit) RegisterChild(arg0 Circuit) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterChild", arg0)
}

// RegisterChild indicates an expected call of RegisterChild.
func (mr *MockCircuitMockRecorder) RegisterChild(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterChild", reflect.TypeOf((*MockCircuit)(nil).RegisterChild), arg0)
}

// SetMaxCurrent mocks base method.
func (m *MockCircuit) SetMaxCurrent(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxCurrent", arg0)
}

// SetMaxCurrent indicates an expected call of SetMaxCurrent.
func (mr *MockCircuitMockRecorder) SetMaxCurrent(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxCurrent", reflect.TypeOf((*MockCircuit)(nil).SetMaxCurrent), arg0)
}

// SetMaxPower mocks base method.
func (m *MockCircuit) SetMaxPower(arg0 float64) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetMaxPower", arg0)
}

// SetMaxPower indicates an expected call of SetMaxPower.
func (mr *MockCircuitMockRecorder) SetMaxPower(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxPower", reflect.TypeOf((*MockCircuit)(nil).SetMaxPower), arg0)
}

// SetTitle mocks base method.
func (m *MockCircuit) SetTitle(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTitle", arg0)
}

// SetTitle indicates an expected call of SetTitle.
func (mr *MockCircuitMockRecorder) SetTitle(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTitle", reflect.TypeOf((*MockCircuit)(nil).SetTitle), arg0)
}

// Update mocks base method.
func (m *MockCircuit) Update(arg0 []CircuitLoad) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockCircuitMockRecorder) Update(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCircuit)(nil).Update), arg0)
}

// ValidateCurrent mocks base method.
func (m *MockCircuit) ValidateCurrent(arg0, arg1 float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateCurrent", arg0, arg1)
	ret0, _ := ret[0].(float64)
	return ret0
}

// ValidateCurrent indicates an expected call of ValidateCurrent.
func (mr *MockCircuitMockRecorder) ValidateCurrent(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateCurrent", reflect.TypeOf((*MockCircuit)(nil).ValidateCurrent), arg0, arg1)
}

// ValidatePower mocks base method.
func (m *MockCircuit) ValidatePower(arg0, arg1 float64) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatePower", arg0, arg1)
	ret0, _ := ret[0].(float64)
	return ret0
}

// ValidatePower indicates an expected call of ValidatePower.
func (mr *MockCircuitMockRecorder) ValidatePower(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatePower", reflect.TypeOf((*MockCircuit)(nil).ValidatePower), arg0, arg1)
}

// Wrap mocks base method.
func (m *MockCircuit) Wrap(arg0 Circuit) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wrap", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wrap indicates an expected call of Wrap.
func (mr *MockCircuitMockRecorder) Wrap(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wrap", reflect.TypeOf((*MockCircuit)(nil).Wrap), arg0)
}
