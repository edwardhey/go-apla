// Code generated by mockery v1.0.0
package crypto

import mock "github.com/stretchr/testify/mock"
import model "github.com/AplaProject/go-apla/tools/update_server/model"

// MockSigner is an autogenerated mock type for the Signer type
type MockSigner struct {
	mock.Mock
}

// CheckSign provides a mock function with given fields: build, public
func (_m *MockSigner) CheckSign(build model.Build, public []byte) (bool, error) {
	ret := _m.Called(build, public)

	var r0 bool
	if rf, ok := ret.Get(0).(func(model.Build, []byte) bool); ok {
		r0 = rf(build, public)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Build, []byte) error); ok {
		r1 = rf(build, public)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MakeSign provides a mock function with given fields: build
func (_m *MockSigner) MakeSign(build model.Build) ([]byte, error) {
	ret := _m.Called(build)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(model.Build) []byte); ok {
		r0 = rf(build)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(model.Build) error); ok {
		r1 = rf(build)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
