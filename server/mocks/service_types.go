// Code generated by MockGen. DO NOT EDIT.
// Source: ./core/service_types.go
//
// Generated by this command:
//
//	mockgen -source=./core/service_types.go -package=mocks
//
// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	core "main/core"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"
)

// MockLLM is a mock of LLM interface.
type MockLLM struct {
	ctrl     *gomock.Controller
	recorder *MockLLMMockRecorder
}

// MockLLMMockRecorder is the mock recorder for MockLLM.
type MockLLMMockRecorder struct {
	mock *MockLLM
}

// NewMockLLM creates a new mock instance.
func NewMockLLM(ctrl *gomock.Controller) *MockLLM {
	mock := &MockLLM{ctrl: ctrl}
	mock.recorder = &MockLLMMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLLM) EXPECT() *MockLLMMockRecorder {
	return m.recorder
}

// Completion mocks base method.
func (m *MockLLM) Completion(ctx context.Context, prompt string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Completion", ctx, prompt)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Completion indicates an expected call of Completion.
func (mr *MockLLMMockRecorder) Completion(ctx, prompt any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Completion", reflect.TypeOf((*MockLLM)(nil).Completion), ctx, prompt)
}

// Embedding mocks base method.
func (m *MockLLM) Embedding(ctx context.Context, text string) ([]float32, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Embedding", ctx, text)
	ret0, _ := ret[0].([]float32)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Embedding indicates an expected call of Embedding.
func (mr *MockLLMMockRecorder) Embedding(ctx, text any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Embedding", reflect.TypeOf((*MockLLM)(nil).Embedding), ctx, text)
}

// MockSTT is a mock of STT interface.
type MockSTT struct {
	ctrl     *gomock.Controller
	recorder *MockSTTMockRecorder
}

// MockSTTMockRecorder is the mock recorder for MockSTT.
type MockSTTMockRecorder struct {
	mock *MockSTT
}

// NewMockSTT creates a new mock instance.
func NewMockSTT(ctrl *gomock.Controller) *MockSTT {
	mock := &MockSTT{ctrl: ctrl}
	mock.recorder = &MockSTTMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSTT) EXPECT() *MockSTTMockRecorder {
	return m.recorder
}

// SpeechToText mocks base method.
func (m *MockSTT) SpeechToText(ctx context.Context, wavBytes []byte) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpeechToText", ctx, wavBytes)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SpeechToText indicates an expected call of SpeechToText.
func (mr *MockSTTMockRecorder) SpeechToText(ctx, wavBytes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpeechToText", reflect.TypeOf((*MockSTT)(nil).SpeechToText), ctx, wavBytes)
}

// MockTuneIn is a mock of TuneIn interface.
type MockTuneIn struct {
	ctrl     *gomock.Controller
	recorder *MockTuneInMockRecorder
}

// MockTuneInMockRecorder is the mock recorder for MockTuneIn.
type MockTuneInMockRecorder struct {
	mock *MockTuneIn
}

// NewMockTuneIn creates a new mock instance.
func NewMockTuneIn(ctrl *gomock.Controller) *MockTuneIn {
	mock := &MockTuneIn{ctrl: ctrl}
	mock.recorder = &MockTuneInMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTuneIn) EXPECT() *MockTuneInMockRecorder {
	return m.recorder
}

// GetStreamURL mocks base method.
func (m *MockTuneIn) GetStreamURL(query string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStreamURL", query)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStreamURL indicates an expected call of GetStreamURL.
func (mr *MockTuneInMockRecorder) GetStreamURL(query any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStreamURL", reflect.TypeOf((*MockTuneIn)(nil).GetStreamURL), query)
}

// MockPush is a mock of Push interface.
type MockPush struct {
	ctrl     *gomock.Controller
	recorder *MockPushMockRecorder
}

// MockPushMockRecorder is the mock recorder for MockPush.
type MockPushMockRecorder struct {
	mock *MockPush
}

// NewMockPush creates a new mock instance.
func NewMockPush(ctrl *gomock.Controller) *MockPush {
	mock := &MockPush{ctrl: ctrl}
	mock.recorder = &MockPushMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPush) EXPECT() *MockPushMockRecorder {
	return m.recorder
}

// SendLater mocks base method.
func (m *MockPush) SendLater(ctx context.Context, when time.Time, source, clientId string, message core.PushMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendLater", ctx, when, source, clientId, message)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendLater indicates an expected call of SendLater.
func (mr *MockPushMockRecorder) SendLater(ctx, when, source, clientId, message any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendLater", reflect.TypeOf((*MockPush)(nil).SendLater), ctx, when, source, clientId, message)
}

// SendRecurring mocks base method.
func (m *MockPush) SendRecurring(ctx context.Context, source, clientId, schedule, intent string, info map[string]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendRecurring", ctx, source, clientId, schedule, intent, info)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendRecurring indicates an expected call of SendRecurring.
func (mr *MockPushMockRecorder) SendRecurring(ctx, source, clientId, schedule, intent, info any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendRecurring", reflect.TypeOf((*MockPush)(nil).SendRecurring), ctx, source, clientId, schedule, intent, info)
}

// MockHomeAssistant is a mock of HomeAssistant interface.
type MockHomeAssistant struct {
	ctrl     *gomock.Controller
	recorder *MockHomeAssistantMockRecorder
}

// MockHomeAssistantMockRecorder is the mock recorder for MockHomeAssistant.
type MockHomeAssistantMockRecorder struct {
	mock *MockHomeAssistant
}

// NewMockHomeAssistant creates a new mock instance.
func NewMockHomeAssistant(ctrl *gomock.Controller) *MockHomeAssistant {
	mock := &MockHomeAssistant{ctrl: ctrl}
	mock.recorder = &MockHomeAssistantMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHomeAssistant) EXPECT() *MockHomeAssistantMockRecorder {
	return m.recorder
}

// Groups mocks base method.
func (m *MockHomeAssistant) Groups() []core.HomeAssistantGroup {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Groups")
	ret0, _ := ret[0].([]core.HomeAssistantGroup)
	return ret0
}

// Groups indicates an expected call of Groups.
func (mr *MockHomeAssistantMockRecorder) Groups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Groups", reflect.TypeOf((*MockHomeAssistant)(nil).Groups))
}

// ToggleDeviceState mocks base method.
func (m *MockHomeAssistant) ToggleDeviceState(deviceId string, on bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToggleDeviceState", deviceId, on)
	ret0, _ := ret[0].(error)
	return ret0
}

// ToggleDeviceState indicates an expected call of ToggleDeviceState.
func (mr *MockHomeAssistantMockRecorder) ToggleDeviceState(deviceId, on any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToggleDeviceState", reflect.TypeOf((*MockHomeAssistant)(nil).ToggleDeviceState), deviceId, on)
}

// MockWeather is a mock of Weather interface.
type MockWeather struct {
	ctrl     *gomock.Controller
	recorder *MockWeatherMockRecorder
}

// MockWeatherMockRecorder is the mock recorder for MockWeather.
type MockWeatherMockRecorder struct {
	mock *MockWeather
}

// NewMockWeather creates a new mock instance.
func NewMockWeather(ctrl *gomock.Controller) *MockWeather {
	mock := &MockWeather{ctrl: ctrl}
	mock.recorder = &MockWeatherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeather) EXPECT() *MockWeatherMockRecorder {
	return m.recorder
}

// PredictWeather mocks base method.
func (m *MockWeather) PredictWeather(coord core.Coordinate) (core.WeatherForecast, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PredictWeather", coord)
	ret0, _ := ret[0].(core.WeatherForecast)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PredictWeather indicates an expected call of PredictWeather.
func (mr *MockWeatherMockRecorder) PredictWeather(coord any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PredictWeather", reflect.TypeOf((*MockWeather)(nil).PredictWeather), coord)
}

// MockGeocoder is a mock of Geocoder interface.
type MockGeocoder struct {
	ctrl     *gomock.Controller
	recorder *MockGeocoderMockRecorder
}

// MockGeocoderMockRecorder is the mock recorder for MockGeocoder.
type MockGeocoderMockRecorder struct {
	mock *MockGeocoder
}

// NewMockGeocoder creates a new mock instance.
func NewMockGeocoder(ctrl *gomock.Controller) *MockGeocoder {
	mock := &MockGeocoder{ctrl: ctrl}
	mock.recorder = &MockGeocoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGeocoder) EXPECT() *MockGeocoderMockRecorder {
	return m.recorder
}

// Geocode mocks base method.
func (m *MockGeocoder) Geocode(q string) (core.Coordinate, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Geocode", q)
	ret0, _ := ret[0].(core.Coordinate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Geocode indicates an expected call of Geocode.
func (mr *MockGeocoderMockRecorder) Geocode(q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Geocode", reflect.TypeOf((*MockGeocoder)(nil).Geocode), q)
}