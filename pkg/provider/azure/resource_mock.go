/*
Copyright 2023 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package azure

import (
	"context"
	"reflect"

	"go.uber.org/mock/gomock"
	clientcache "k8s.io/client-go/tools/cache"

	"sigs.k8s.io/cloud-provider-azure/pkg/cache"
)

// MockResource is a mock of Resource interface
type MockResource struct {
	ctrl     *gomock.Controller
	recorder *MockResourceMockRecorder
}

// MockResourceMockRecorder is the mock recorder for MockResource
type MockResourceMockRecorder struct {
	mock *MockResource
}

// NewMockResource creates a new mock instance
func NewMockResource(ctrl *gomock.Controller) *MockResource {
	mock := &MockResource{ctrl: ctrl}
	mock.recorder = &MockResourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResource) EXPECT() *MockResourceMockRecorder {
	return m.recorder
}

// GetWithDeepCopy mocks base method
func (m *MockResource) GetWithDeepCopy(ctx context.Context, key string, crt cache.AzureCacheReadType) (interface{}, error) {
	ret := m.ctrl.Call(m, "GetWithDeepCopy", ctx, key, crt)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWithDeepCopy indicates an expected call of GetWithDeepCopy
func (mr *MockResourceMockRecorder) GetWithDeepCopy(ctx, key, crt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWithDeepCopy", reflect.TypeOf((*MockResource)(nil).GetWithDeepCopy), ctx, key, crt)
}

// Get mocks base method
func (m *MockResource) Get(ctx context.Context, key string, crt cache.AzureCacheReadType) (interface{}, error) {
	ret := m.ctrl.Call(m, "Get", ctx, key, crt)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockResourceMockRecorder) Get(ctx, key, crt interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockResource)(nil).Get), ctx, key, crt)
}

// Delete mocks base method
func (m *MockResource) Delete(key string) error {
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockResourceMockRecorder) Delete(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockResource)(nil).Delete), key)
}

// Set mocks base method
func (m *MockResource) Set(key string, value interface{}) {
	m.ctrl.Call(m, "Set", key, value)
}

// Set indicates an expected call of Set
func (mr *MockResourceMockRecorder) Set(key, value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockResource)(nil).Set), key, value)
}

// Update mocks base method
func (m *MockResource) Update(key string, value interface{}) {
	m.ctrl.Call(m, "Update", key, value)
}

// Update indicates an expected call of Update
func (mr *MockResourceMockRecorder) Update(key, value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockResource)(nil).Update), key, value)
}

// Lock mocks base method
func (m *MockResource) Lock() {
	m.ctrl.Call(m, "Lock")
}

// Lock indicates an expected call of Lock
func (mr *MockResourceMockRecorder) Lock() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockResource)(nil).Lock))
}

// Unlock mocks base method
func (m *MockResource) Unlock() {
	m.ctrl.Call(m, "Unlock")
}

// Unlock indicates an expected call of Unlock
func (mr *MockResourceMockRecorder) Unlock() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockResource)(nil).Unlock))
}

// GetStore mocks base method
func (m *MockResource) GetStore() clientcache.Store {
	ret := m.ctrl.Call(m, "GetStore")
	ret0, _ := ret[0].(clientcache.Store)
	return ret0
}

// GetStore indicates an expected call of GetStore
func (mr *MockResourceMockRecorder) GetStore() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStore", reflect.TypeOf((*MockResource)(nil).GetStore))
}