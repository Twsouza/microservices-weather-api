// Code generated by mockery v2.46.3. DO NOT EDIT.

package services

import mock "github.com/stretchr/testify/mock"

// MockWeatherService is an autogenerated mock type for the WeatherService type
type MockWeatherService struct {
	mock.Mock
}

type MockWeatherService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWeatherService) EXPECT() *MockWeatherService_Expecter {
	return &MockWeatherService_Expecter{mock: &_m.Mock}
}

// CalculateTemperature provides a mock function with given fields: tempC
func (_m *MockWeatherService) CalculateTemperature(tempC float64) (float64, float64) {
	ret := _m.Called(tempC)

	if len(ret) == 0 {
		panic("no return value specified for CalculateTemperature")
	}

	var r0 float64
	var r1 float64
	if rf, ok := ret.Get(0).(func(float64) (float64, float64)); ok {
		return rf(tempC)
	}
	if rf, ok := ret.Get(0).(func(float64) float64); ok {
		r0 = rf(tempC)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(float64) float64); ok {
		r1 = rf(tempC)
	} else {
		r1 = ret.Get(1).(float64)
	}

	return r0, r1
}

// MockWeatherService_CalculateTemperature_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CalculateTemperature'
type MockWeatherService_CalculateTemperature_Call struct {
	*mock.Call
}

// CalculateTemperature is a helper method to define mock.On call
//   - tempC float64
func (_e *MockWeatherService_Expecter) CalculateTemperature(tempC interface{}) *MockWeatherService_CalculateTemperature_Call {
	return &MockWeatherService_CalculateTemperature_Call{Call: _e.mock.On("CalculateTemperature", tempC)}
}

func (_c *MockWeatherService_CalculateTemperature_Call) Run(run func(tempC float64)) *MockWeatherService_CalculateTemperature_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(float64))
	})
	return _c
}

func (_c *MockWeatherService_CalculateTemperature_Call) Return(fahrenheit float64, kelvin float64) *MockWeatherService_CalculateTemperature_Call {
	_c.Call.Return(fahrenheit, kelvin)
	return _c
}

func (_c *MockWeatherService_CalculateTemperature_Call) RunAndReturn(run func(float64) (float64, float64)) *MockWeatherService_CalculateTemperature_Call {
	_c.Call.Return(run)
	return _c
}

// GetTemperatureByLocation provides a mock function with given fields: location
func (_m *MockWeatherService) GetTemperatureByLocation(location string) (float64, error) {
	ret := _m.Called(location)

	if len(ret) == 0 {
		panic("no return value specified for GetTemperatureByLocation")
	}

	var r0 float64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (float64, error)); ok {
		return rf(location)
	}
	if rf, ok := ret.Get(0).(func(string) float64); ok {
		r0 = rf(location)
	} else {
		r0 = ret.Get(0).(float64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(location)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWeatherService_GetTemperatureByLocation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTemperatureByLocation'
type MockWeatherService_GetTemperatureByLocation_Call struct {
	*mock.Call
}

// GetTemperatureByLocation is a helper method to define mock.On call
//   - location string
func (_e *MockWeatherService_Expecter) GetTemperatureByLocation(location interface{}) *MockWeatherService_GetTemperatureByLocation_Call {
	return &MockWeatherService_GetTemperatureByLocation_Call{Call: _e.mock.On("GetTemperatureByLocation", location)}
}

func (_c *MockWeatherService_GetTemperatureByLocation_Call) Run(run func(location string)) *MockWeatherService_GetTemperatureByLocation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockWeatherService_GetTemperatureByLocation_Call) Return(_a0 float64, _a1 error) *MockWeatherService_GetTemperatureByLocation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWeatherService_GetTemperatureByLocation_Call) RunAndReturn(run func(string) (float64, error)) *MockWeatherService_GetTemperatureByLocation_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWeatherService creates a new instance of MockWeatherService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWeatherService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWeatherService {
	mock := &MockWeatherService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
