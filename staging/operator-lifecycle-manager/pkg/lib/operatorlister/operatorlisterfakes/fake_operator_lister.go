// Code generated by counterfeiter. DO NOT EDIT.
package operatorlisterfakes

import (
	"sync"

	"github.com/operator-framework/operator-lifecycle-manager/pkg/lib/operatorlister"
)

type FakeOperatorLister struct {
	APIExtensionsV1Stub        func() operatorlister.APIExtensionsV1Lister
	aPIExtensionsV1Mutex       sync.RWMutex
	aPIExtensionsV1ArgsForCall []struct {
	}
	aPIExtensionsV1Returns struct {
		result1 operatorlister.APIExtensionsV1Lister
	}
	aPIExtensionsV1ReturnsOnCall map[int]struct {
		result1 operatorlister.APIExtensionsV1Lister
	}
	APIRegistrationV1Stub        func() operatorlister.APIRegistrationV1Lister
	aPIRegistrationV1Mutex       sync.RWMutex
	aPIRegistrationV1ArgsForCall []struct {
	}
	aPIRegistrationV1Returns struct {
		result1 operatorlister.APIRegistrationV1Lister
	}
	aPIRegistrationV1ReturnsOnCall map[int]struct {
		result1 operatorlister.APIRegistrationV1Lister
	}
	AppsV1Stub        func() operatorlister.AppsV1Lister
	appsV1Mutex       sync.RWMutex
	appsV1ArgsForCall []struct {
	}
	appsV1Returns struct {
		result1 operatorlister.AppsV1Lister
	}
	appsV1ReturnsOnCall map[int]struct {
		result1 operatorlister.AppsV1Lister
	}
	CoreV1Stub        func() operatorlister.CoreV1Lister
	coreV1Mutex       sync.RWMutex
	coreV1ArgsForCall []struct {
	}
	coreV1Returns struct {
		result1 operatorlister.CoreV1Lister
	}
	coreV1ReturnsOnCall map[int]struct {
		result1 operatorlister.CoreV1Lister
	}
	OperatorsV1Stub        func() operatorlister.OperatorsV1Lister
	operatorsV1Mutex       sync.RWMutex
	operatorsV1ArgsForCall []struct {
	}
	operatorsV1Returns struct {
		result1 operatorlister.OperatorsV1Lister
	}
	operatorsV1ReturnsOnCall map[int]struct {
		result1 operatorlister.OperatorsV1Lister
	}
	OperatorsV1alpha1Stub        func() operatorlister.OperatorsV1alpha1Lister
	operatorsV1alpha1Mutex       sync.RWMutex
	operatorsV1alpha1ArgsForCall []struct {
	}
	operatorsV1alpha1Returns struct {
		result1 operatorlister.OperatorsV1alpha1Lister
	}
	operatorsV1alpha1ReturnsOnCall map[int]struct {
		result1 operatorlister.OperatorsV1alpha1Lister
	}
	OperatorsV2Stub        func() operatorlister.OperatorsV2Lister
	operatorsV2Mutex       sync.RWMutex
	operatorsV2ArgsForCall []struct {
	}
	operatorsV2Returns struct {
		result1 operatorlister.OperatorsV2Lister
	}
	operatorsV2ReturnsOnCall map[int]struct {
		result1 operatorlister.OperatorsV2Lister
	}
	RbacV1Stub        func() operatorlister.RbacV1Lister
	rbacV1Mutex       sync.RWMutex
	rbacV1ArgsForCall []struct {
	}
	rbacV1Returns struct {
		result1 operatorlister.RbacV1Lister
	}
	rbacV1ReturnsOnCall map[int]struct {
		result1 operatorlister.RbacV1Lister
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOperatorLister) APIExtensionsV1() operatorlister.APIExtensionsV1Lister {
	fake.aPIExtensionsV1Mutex.Lock()
	ret, specificReturn := fake.aPIExtensionsV1ReturnsOnCall[len(fake.aPIExtensionsV1ArgsForCall)]
	fake.aPIExtensionsV1ArgsForCall = append(fake.aPIExtensionsV1ArgsForCall, struct {
	}{})
	stub := fake.APIExtensionsV1Stub
	fakeReturns := fake.aPIExtensionsV1Returns
	fake.recordInvocation("APIExtensionsV1", []interface{}{})
	fake.aPIExtensionsV1Mutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOperatorLister) APIExtensionsV1CallCount() int {
	fake.aPIExtensionsV1Mutex.RLock()
	defer fake.aPIExtensionsV1Mutex.RUnlock()
	return len(fake.aPIExtensionsV1ArgsForCall)
}

func (fake *FakeOperatorLister) APIExtensionsV1Calls(stub func() operatorlister.APIExtensionsV1Lister) {
	fake.aPIExtensionsV1Mutex.Lock()
	defer fake.aPIExtensionsV1Mutex.Unlock()
	fake.APIExtensionsV1Stub = stub
}

func (fake *FakeOperatorLister) APIExtensionsV1Returns(result1 operatorlister.APIExtensionsV1Lister) {
	fake.aPIExtensionsV1Mutex.Lock()
	defer fake.aPIExtensionsV1Mutex.Unlock()
	fake.APIExtensionsV1Stub = nil
	fake.aPIExtensionsV1Returns = struct {
		result1 operatorlister.APIExtensionsV1Lister
	}{result1}
}

func (fake *FakeOperatorLister) APIExtensionsV1ReturnsOnCall(i int, result1 operatorlister.APIExtensionsV1Lister) {
	fake.aPIExtensionsV1Mutex.Lock()
	defer fake.aPIExtensionsV1Mutex.Unlock()
	fake.APIExtensionsV1Stub = nil
	if fake.aPIExtensionsV1ReturnsOnCall == nil {
		fake.aPIExtensionsV1ReturnsOnCall = make(map[int]struct {
			result1 operatorlister.APIExtensionsV1Lister
		})
	}
	fake.aPIExtensionsV1ReturnsOnCall[i] = struct {
		result1 operatorlister.APIExtensionsV1Lister
	}{result1}
}

func (fake *FakeOperatorLister) APIRegistrationV1() operatorlister.APIRegistrationV1Lister {
	fake.aPIRegistrationV1Mutex.Lock()
	ret, specificReturn := fake.aPIRegistrationV1ReturnsOnCall[len(fake.aPIRegistrationV1ArgsForCall)]
	fake.aPIRegistrationV1ArgsForCall = append(fake.aPIRegistrationV1ArgsForCall, struct {
	}{})
	stub := fake.APIRegistrationV1Stub
	fakeReturns := fake.aPIRegistrationV1Returns
	fake.recordInvocation("APIRegistrationV1", []interface{}{})
	fake.aPIRegistrationV1Mutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOperatorLister) APIRegistrationV1CallCount() int {
	fake.aPIRegistrationV1Mutex.RLock()
	defer fake.aPIRegistrationV1Mutex.RUnlock()
	return len(fake.aPIRegistrationV1ArgsForCall)
}

func (fake *FakeOperatorLister) APIRegistrationV1Calls(stub func() operatorlister.APIRegistrationV1Lister) {
	fake.aPIRegistrationV1Mutex.Lock()
	defer fake.aPIRegistrationV1Mutex.Unlock()
	fake.APIRegistrationV1Stub = stub
}

func (fake *FakeOperatorLister) APIRegistrationV1Returns(result1 operatorlister.APIRegistrationV1Lister) {
	fake.aPIRegistrationV1Mutex.Lock()
	defer fake.aPIRegistrationV1Mutex.Unlock()
	fake.APIRegistrationV1Stub = nil
	fake.aPIRegistrationV1Returns = struct {
		result1 operatorlister.APIRegistrationV1Lister
	}{result1}
}

func (fake *FakeOperatorLister) APIRegistrationV1ReturnsOnCall(i int, result1 operatorlister.APIRegistrationV1Lister) {
	fake.aPIRegistrationV1Mutex.Lock()
	defer fake.aPIRegistrationV1Mutex.Unlock()
	fake.APIRegistrationV1Stub = nil
	if fake.aPIRegistrationV1ReturnsOnCall == nil {
		fake.aPIRegistrationV1ReturnsOnCall = make(map[int]struct {
			result1 operatorlister.APIRegistrationV1Lister
		})
	}
	fake.aPIRegistrationV1ReturnsOnCall[i] = struct {
		result1 operatorlister.APIRegistrationV1Lister
	}{result1}
}

func (fake *FakeOperatorLister) AppsV1() operatorlister.AppsV1Lister {
	fake.appsV1Mutex.Lock()
	ret, specificReturn := fake.appsV1ReturnsOnCall[len(fake.appsV1ArgsForCall)]
	fake.appsV1ArgsForCall = append(fake.appsV1ArgsForCall, struct {
	}{})
	stub := fake.AppsV1Stub
	fakeReturns := fake.appsV1Returns
	fake.recordInvocation("AppsV1", []interface{}{})
	fake.appsV1Mutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOperatorLister) AppsV1CallCount() int {
	fake.appsV1Mutex.RLock()
	defer fake.appsV1Mutex.RUnlock()
	return len(fake.appsV1ArgsForCall)
}

func (fake *FakeOperatorLister) AppsV1Calls(stub func() operatorlister.AppsV1Lister) {
	fake.appsV1Mutex.Lock()
	defer fake.appsV1Mutex.Unlock()
	fake.AppsV1Stub = stub
}

func (fake *FakeOperatorLister) AppsV1Returns(result1 operatorlister.AppsV1Lister) {
	fake.appsV1Mutex.Lock()
	defer fake.appsV1Mutex.Unlock()
	fake.AppsV1Stub = nil
	fake.appsV1Returns = struct {
		result1 operatorlister.AppsV1Lister
	}{result1}
}

func (fake *FakeOperatorLister) AppsV1ReturnsOnCall(i int, result1 operatorlister.AppsV1Lister) {
	fake.appsV1Mutex.Lock()
	defer fake.appsV1Mutex.Unlock()
	fake.AppsV1Stub = nil
	if fake.appsV1ReturnsOnCall == nil {
		fake.appsV1ReturnsOnCall = make(map[int]struct {
			result1 operatorlister.AppsV1Lister
		})
	}
	fake.appsV1ReturnsOnCall[i] = struct {
		result1 operatorlister.AppsV1Lister
	}{result1}
}

func (fake *FakeOperatorLister) CoreV1() operatorlister.CoreV1Lister {
	fake.coreV1Mutex.Lock()
	ret, specificReturn := fake.coreV1ReturnsOnCall[len(fake.coreV1ArgsForCall)]
	fake.coreV1ArgsForCall = append(fake.coreV1ArgsForCall, struct {
	}{})
	stub := fake.CoreV1Stub
	fakeReturns := fake.coreV1Returns
	fake.recordInvocation("CoreV1", []interface{}{})
	fake.coreV1Mutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOperatorLister) CoreV1CallCount() int {
	fake.coreV1Mutex.RLock()
	defer fake.coreV1Mutex.RUnlock()
	return len(fake.coreV1ArgsForCall)
}

func (fake *FakeOperatorLister) CoreV1Calls(stub func() operatorlister.CoreV1Lister) {
	fake.coreV1Mutex.Lock()
	defer fake.coreV1Mutex.Unlock()
	fake.CoreV1Stub = stub
}

func (fake *FakeOperatorLister) CoreV1Returns(result1 operatorlister.CoreV1Lister) {
	fake.coreV1Mutex.Lock()
	defer fake.coreV1Mutex.Unlock()
	fake.CoreV1Stub = nil
	fake.coreV1Returns = struct {
		result1 operatorlister.CoreV1Lister
	}{result1}
}

func (fake *FakeOperatorLister) CoreV1ReturnsOnCall(i int, result1 operatorlister.CoreV1Lister) {
	fake.coreV1Mutex.Lock()
	defer fake.coreV1Mutex.Unlock()
	fake.CoreV1Stub = nil
	if fake.coreV1ReturnsOnCall == nil {
		fake.coreV1ReturnsOnCall = make(map[int]struct {
			result1 operatorlister.CoreV1Lister
		})
	}
	fake.coreV1ReturnsOnCall[i] = struct {
		result1 operatorlister.CoreV1Lister
	}{result1}
}

func (fake *FakeOperatorLister) OperatorsV1() operatorlister.OperatorsV1Lister {
	fake.operatorsV1Mutex.Lock()
	ret, specificReturn := fake.operatorsV1ReturnsOnCall[len(fake.operatorsV1ArgsForCall)]
	fake.operatorsV1ArgsForCall = append(fake.operatorsV1ArgsForCall, struct {
	}{})
	stub := fake.OperatorsV1Stub
	fakeReturns := fake.operatorsV1Returns
	fake.recordInvocation("OperatorsV1", []interface{}{})
	fake.operatorsV1Mutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOperatorLister) OperatorsV1CallCount() int {
	fake.operatorsV1Mutex.RLock()
	defer fake.operatorsV1Mutex.RUnlock()
	return len(fake.operatorsV1ArgsForCall)
}

func (fake *FakeOperatorLister) OperatorsV1Calls(stub func() operatorlister.OperatorsV1Lister) {
	fake.operatorsV1Mutex.Lock()
	defer fake.operatorsV1Mutex.Unlock()
	fake.OperatorsV1Stub = stub
}

func (fake *FakeOperatorLister) OperatorsV1Returns(result1 operatorlister.OperatorsV1Lister) {
	fake.operatorsV1Mutex.Lock()
	defer fake.operatorsV1Mutex.Unlock()
	fake.OperatorsV1Stub = nil
	fake.operatorsV1Returns = struct {
		result1 operatorlister.OperatorsV1Lister
	}{result1}
}

func (fake *FakeOperatorLister) OperatorsV1ReturnsOnCall(i int, result1 operatorlister.OperatorsV1Lister) {
	fake.operatorsV1Mutex.Lock()
	defer fake.operatorsV1Mutex.Unlock()
	fake.OperatorsV1Stub = nil
	if fake.operatorsV1ReturnsOnCall == nil {
		fake.operatorsV1ReturnsOnCall = make(map[int]struct {
			result1 operatorlister.OperatorsV1Lister
		})
	}
	fake.operatorsV1ReturnsOnCall[i] = struct {
		result1 operatorlister.OperatorsV1Lister
	}{result1}
}

func (fake *FakeOperatorLister) OperatorsV1alpha1() operatorlister.OperatorsV1alpha1Lister {
	fake.operatorsV1alpha1Mutex.Lock()
	ret, specificReturn := fake.operatorsV1alpha1ReturnsOnCall[len(fake.operatorsV1alpha1ArgsForCall)]
	fake.operatorsV1alpha1ArgsForCall = append(fake.operatorsV1alpha1ArgsForCall, struct {
	}{})
	stub := fake.OperatorsV1alpha1Stub
	fakeReturns := fake.operatorsV1alpha1Returns
	fake.recordInvocation("OperatorsV1alpha1", []interface{}{})
	fake.operatorsV1alpha1Mutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOperatorLister) OperatorsV1alpha1CallCount() int {
	fake.operatorsV1alpha1Mutex.RLock()
	defer fake.operatorsV1alpha1Mutex.RUnlock()
	return len(fake.operatorsV1alpha1ArgsForCall)
}

func (fake *FakeOperatorLister) OperatorsV1alpha1Calls(stub func() operatorlister.OperatorsV1alpha1Lister) {
	fake.operatorsV1alpha1Mutex.Lock()
	defer fake.operatorsV1alpha1Mutex.Unlock()
	fake.OperatorsV1alpha1Stub = stub
}

func (fake *FakeOperatorLister) OperatorsV1alpha1Returns(result1 operatorlister.OperatorsV1alpha1Lister) {
	fake.operatorsV1alpha1Mutex.Lock()
	defer fake.operatorsV1alpha1Mutex.Unlock()
	fake.OperatorsV1alpha1Stub = nil
	fake.operatorsV1alpha1Returns = struct {
		result1 operatorlister.OperatorsV1alpha1Lister
	}{result1}
}

func (fake *FakeOperatorLister) OperatorsV1alpha1ReturnsOnCall(i int, result1 operatorlister.OperatorsV1alpha1Lister) {
	fake.operatorsV1alpha1Mutex.Lock()
	defer fake.operatorsV1alpha1Mutex.Unlock()
	fake.OperatorsV1alpha1Stub = nil
	if fake.operatorsV1alpha1ReturnsOnCall == nil {
		fake.operatorsV1alpha1ReturnsOnCall = make(map[int]struct {
			result1 operatorlister.OperatorsV1alpha1Lister
		})
	}
	fake.operatorsV1alpha1ReturnsOnCall[i] = struct {
		result1 operatorlister.OperatorsV1alpha1Lister
	}{result1}
}

func (fake *FakeOperatorLister) OperatorsV2() operatorlister.OperatorsV2Lister {
	fake.operatorsV2Mutex.Lock()
	ret, specificReturn := fake.operatorsV2ReturnsOnCall[len(fake.operatorsV2ArgsForCall)]
	fake.operatorsV2ArgsForCall = append(fake.operatorsV2ArgsForCall, struct {
	}{})
	stub := fake.OperatorsV2Stub
	fakeReturns := fake.operatorsV2Returns
	fake.recordInvocation("OperatorsV2", []interface{}{})
	fake.operatorsV2Mutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOperatorLister) OperatorsV2CallCount() int {
	fake.operatorsV2Mutex.RLock()
	defer fake.operatorsV2Mutex.RUnlock()
	return len(fake.operatorsV2ArgsForCall)
}

func (fake *FakeOperatorLister) OperatorsV2Calls(stub func() operatorlister.OperatorsV2Lister) {
	fake.operatorsV2Mutex.Lock()
	defer fake.operatorsV2Mutex.Unlock()
	fake.OperatorsV2Stub = stub
}

func (fake *FakeOperatorLister) OperatorsV2Returns(result1 operatorlister.OperatorsV2Lister) {
	fake.operatorsV2Mutex.Lock()
	defer fake.operatorsV2Mutex.Unlock()
	fake.OperatorsV2Stub = nil
	fake.operatorsV2Returns = struct {
		result1 operatorlister.OperatorsV2Lister
	}{result1}
}

func (fake *FakeOperatorLister) OperatorsV2ReturnsOnCall(i int, result1 operatorlister.OperatorsV2Lister) {
	fake.operatorsV2Mutex.Lock()
	defer fake.operatorsV2Mutex.Unlock()
	fake.OperatorsV2Stub = nil
	if fake.operatorsV2ReturnsOnCall == nil {
		fake.operatorsV2ReturnsOnCall = make(map[int]struct {
			result1 operatorlister.OperatorsV2Lister
		})
	}
	fake.operatorsV2ReturnsOnCall[i] = struct {
		result1 operatorlister.OperatorsV2Lister
	}{result1}
}

func (fake *FakeOperatorLister) RbacV1() operatorlister.RbacV1Lister {
	fake.rbacV1Mutex.Lock()
	ret, specificReturn := fake.rbacV1ReturnsOnCall[len(fake.rbacV1ArgsForCall)]
	fake.rbacV1ArgsForCall = append(fake.rbacV1ArgsForCall, struct {
	}{})
	stub := fake.RbacV1Stub
	fakeReturns := fake.rbacV1Returns
	fake.recordInvocation("RbacV1", []interface{}{})
	fake.rbacV1Mutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeOperatorLister) RbacV1CallCount() int {
	fake.rbacV1Mutex.RLock()
	defer fake.rbacV1Mutex.RUnlock()
	return len(fake.rbacV1ArgsForCall)
}

func (fake *FakeOperatorLister) RbacV1Calls(stub func() operatorlister.RbacV1Lister) {
	fake.rbacV1Mutex.Lock()
	defer fake.rbacV1Mutex.Unlock()
	fake.RbacV1Stub = stub
}

func (fake *FakeOperatorLister) RbacV1Returns(result1 operatorlister.RbacV1Lister) {
	fake.rbacV1Mutex.Lock()
	defer fake.rbacV1Mutex.Unlock()
	fake.RbacV1Stub = nil
	fake.rbacV1Returns = struct {
		result1 operatorlister.RbacV1Lister
	}{result1}
}

func (fake *FakeOperatorLister) RbacV1ReturnsOnCall(i int, result1 operatorlister.RbacV1Lister) {
	fake.rbacV1Mutex.Lock()
	defer fake.rbacV1Mutex.Unlock()
	fake.RbacV1Stub = nil
	if fake.rbacV1ReturnsOnCall == nil {
		fake.rbacV1ReturnsOnCall = make(map[int]struct {
			result1 operatorlister.RbacV1Lister
		})
	}
	fake.rbacV1ReturnsOnCall[i] = struct {
		result1 operatorlister.RbacV1Lister
	}{result1}
}

func (fake *FakeOperatorLister) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.aPIExtensionsV1Mutex.RLock()
	defer fake.aPIExtensionsV1Mutex.RUnlock()
	fake.aPIRegistrationV1Mutex.RLock()
	defer fake.aPIRegistrationV1Mutex.RUnlock()
	fake.appsV1Mutex.RLock()
	defer fake.appsV1Mutex.RUnlock()
	fake.coreV1Mutex.RLock()
	defer fake.coreV1Mutex.RUnlock()
	fake.operatorsV1Mutex.RLock()
	defer fake.operatorsV1Mutex.RUnlock()
	fake.operatorsV1alpha1Mutex.RLock()
	defer fake.operatorsV1alpha1Mutex.RUnlock()
	fake.operatorsV2Mutex.RLock()
	defer fake.operatorsV2Mutex.RUnlock()
	fake.rbacV1Mutex.RLock()
	defer fake.rbacV1Mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOperatorLister) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ operatorlister.OperatorLister = new(FakeOperatorLister)
