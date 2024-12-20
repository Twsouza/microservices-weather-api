// Code generated by mockery v2.46.3. DO NOT EDIT.

package services

import (
	context "context"
	dto "servicea/internal/dto"

	mock "github.com/stretchr/testify/mock"
)

// MockWeatherApiService is an autogenerated mock type for the WeatherApiService type
type MockWeatherApiService struct {
	mock.Mock
}

type MockWeatherApiService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockWeatherApiService) EXPECT() *MockWeatherApiService_Expecter {
	return &MockWeatherApiService_Expecter{mock: &_m.Mock}
}

// GetTemperaturesByZipCode provides a mock function with given fields: ctx, zipCode
func (_m *MockWeatherApiService) GetTemperaturesByZipCode(ctx context.Context, zipCode string) (*dto.WeatherResponse, error) {
	ret := _m.Called(ctx, zipCode)

	if len(ret) == 0 {
		panic("no return value specified for GetTemperaturesByZipCode")
	}

	var r0 *dto.WeatherResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*dto.WeatherResponse, error)); ok {
		return rf(ctx, zipCode)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *dto.WeatherResponse); ok {
		r0 = rf(ctx, zipCode)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dto.WeatherResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, zipCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockWeatherApiService_GetTemperaturesByZipCode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTemperaturesByZipCode'
type MockWeatherApiService_GetTemperaturesByZipCode_Call struct {
	*mock.Call
}

// GetTemperaturesByZipCode is a helper method to define mock.On call
//   - ctx context.Context
//   - zipCode string
func (_e *MockWeatherApiService_Expecter) GetTemperaturesByZipCode(ctx interface{}, zipCode interface{}) *MockWeatherApiService_GetTemperaturesByZipCode_Call {
	return &MockWeatherApiService_GetTemperaturesByZipCode_Call{Call: _e.mock.On("GetTemperaturesByZipCode", ctx, zipCode)}
}

func (_c *MockWeatherApiService_GetTemperaturesByZipCode_Call) Run(run func(ctx context.Context, zipCode string)) *MockWeatherApiService_GetTemperaturesByZipCode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockWeatherApiService_GetTemperaturesByZipCode_Call) Return(_a0 *dto.WeatherResponse, _a1 error) *MockWeatherApiService_GetTemperaturesByZipCode_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockWeatherApiService_GetTemperaturesByZipCode_Call) RunAndReturn(run func(context.Context, string) (*dto.WeatherResponse, error)) *MockWeatherApiService_GetTemperaturesByZipCode_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockWeatherApiService creates a new instance of MockWeatherApiService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockWeatherApiService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockWeatherApiService {
	mock := &MockWeatherApiService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
