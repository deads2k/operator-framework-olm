// Code generated by counterfeiter. DO NOT EDIT.
package operatorlisterfakes

import (
	"sync"

	"github.com/operator-framework/operator-lifecycle-manager/pkg/lib/operatorlister"
	v1 "k8s.io/client-go/listers/rbac/v1"
)

type FakeRbacV1Lister struct {
	ClusterRoleBindingListerStub        func() v1.ClusterRoleBindingLister
	clusterRoleBindingListerMutex       sync.RWMutex
	clusterRoleBindingListerArgsForCall []struct {
	}
	clusterRoleBindingListerReturns struct {
		result1 v1.ClusterRoleBindingLister
	}
	clusterRoleBindingListerReturnsOnCall map[int]struct {
		result1 v1.ClusterRoleBindingLister
	}
	ClusterRoleListerStub        func() v1.ClusterRoleLister
	clusterRoleListerMutex       sync.RWMutex
	clusterRoleListerArgsForCall []struct {
	}
	clusterRoleListerReturns struct {
		result1 v1.ClusterRoleLister
	}
	clusterRoleListerReturnsOnCall map[int]struct {
		result1 v1.ClusterRoleLister
	}
	RegisterClusterRoleBindingListerStub        func(v1.ClusterRoleBindingLister)
	registerClusterRoleBindingListerMutex       sync.RWMutex
	registerClusterRoleBindingListerArgsForCall []struct {
		arg1 v1.ClusterRoleBindingLister
	}
	RegisterClusterRoleListerStub        func(v1.ClusterRoleLister)
	registerClusterRoleListerMutex       sync.RWMutex
	registerClusterRoleListerArgsForCall []struct {
		arg1 v1.ClusterRoleLister
	}
	RegisterRoleBindingListerStub        func(string, v1.RoleBindingLister)
	registerRoleBindingListerMutex       sync.RWMutex
	registerRoleBindingListerArgsForCall []struct {
		arg1 string
		arg2 v1.RoleBindingLister
	}
	RegisterRoleListerStub        func(string, v1.RoleLister)
	registerRoleListerMutex       sync.RWMutex
	registerRoleListerArgsForCall []struct {
		arg1 string
		arg2 v1.RoleLister
	}
	RoleBindingListerStub        func() v1.RoleBindingLister
	roleBindingListerMutex       sync.RWMutex
	roleBindingListerArgsForCall []struct {
	}
	roleBindingListerReturns struct {
		result1 v1.RoleBindingLister
	}
	roleBindingListerReturnsOnCall map[int]struct {
		result1 v1.RoleBindingLister
	}
	RoleListerStub        func() v1.RoleLister
	roleListerMutex       sync.RWMutex
	roleListerArgsForCall []struct {
	}
	roleListerReturns struct {
		result1 v1.RoleLister
	}
	roleListerReturnsOnCall map[int]struct {
		result1 v1.RoleLister
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRbacV1Lister) ClusterRoleBindingLister() v1.ClusterRoleBindingLister {
	fake.clusterRoleBindingListerMutex.Lock()
	ret, specificReturn := fake.clusterRoleBindingListerReturnsOnCall[len(fake.clusterRoleBindingListerArgsForCall)]
	fake.clusterRoleBindingListerArgsForCall = append(fake.clusterRoleBindingListerArgsForCall, struct {
	}{})
	stub := fake.ClusterRoleBindingListerStub
	fakeReturns := fake.clusterRoleBindingListerReturns
	fake.recordInvocation("ClusterRoleBindingLister", []interface{}{})
	fake.clusterRoleBindingListerMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRbacV1Lister) ClusterRoleBindingListerCallCount() int {
	fake.clusterRoleBindingListerMutex.RLock()
	defer fake.clusterRoleBindingListerMutex.RUnlock()
	return len(fake.clusterRoleBindingListerArgsForCall)
}

func (fake *FakeRbacV1Lister) ClusterRoleBindingListerCalls(stub func() v1.ClusterRoleBindingLister) {
	fake.clusterRoleBindingListerMutex.Lock()
	defer fake.clusterRoleBindingListerMutex.Unlock()
	fake.ClusterRoleBindingListerStub = stub
}

func (fake *FakeRbacV1Lister) ClusterRoleBindingListerReturns(result1 v1.ClusterRoleBindingLister) {
	fake.clusterRoleBindingListerMutex.Lock()
	defer fake.clusterRoleBindingListerMutex.Unlock()
	fake.ClusterRoleBindingListerStub = nil
	fake.clusterRoleBindingListerReturns = struct {
		result1 v1.ClusterRoleBindingLister
	}{result1}
}

func (fake *FakeRbacV1Lister) ClusterRoleBindingListerReturnsOnCall(i int, result1 v1.ClusterRoleBindingLister) {
	fake.clusterRoleBindingListerMutex.Lock()
	defer fake.clusterRoleBindingListerMutex.Unlock()
	fake.ClusterRoleBindingListerStub = nil
	if fake.clusterRoleBindingListerReturnsOnCall == nil {
		fake.clusterRoleBindingListerReturnsOnCall = make(map[int]struct {
			result1 v1.ClusterRoleBindingLister
		})
	}
	fake.clusterRoleBindingListerReturnsOnCall[i] = struct {
		result1 v1.ClusterRoleBindingLister
	}{result1}
}

func (fake *FakeRbacV1Lister) ClusterRoleLister() v1.ClusterRoleLister {
	fake.clusterRoleListerMutex.Lock()
	ret, specificReturn := fake.clusterRoleListerReturnsOnCall[len(fake.clusterRoleListerArgsForCall)]
	fake.clusterRoleListerArgsForCall = append(fake.clusterRoleListerArgsForCall, struct {
	}{})
	stub := fake.ClusterRoleListerStub
	fakeReturns := fake.clusterRoleListerReturns
	fake.recordInvocation("ClusterRoleLister", []interface{}{})
	fake.clusterRoleListerMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRbacV1Lister) ClusterRoleListerCallCount() int {
	fake.clusterRoleListerMutex.RLock()
	defer fake.clusterRoleListerMutex.RUnlock()
	return len(fake.clusterRoleListerArgsForCall)
}

func (fake *FakeRbacV1Lister) ClusterRoleListerCalls(stub func() v1.ClusterRoleLister) {
	fake.clusterRoleListerMutex.Lock()
	defer fake.clusterRoleListerMutex.Unlock()
	fake.ClusterRoleListerStub = stub
}

func (fake *FakeRbacV1Lister) ClusterRoleListerReturns(result1 v1.ClusterRoleLister) {
	fake.clusterRoleListerMutex.Lock()
	defer fake.clusterRoleListerMutex.Unlock()
	fake.ClusterRoleListerStub = nil
	fake.clusterRoleListerReturns = struct {
		result1 v1.ClusterRoleLister
	}{result1}
}

func (fake *FakeRbacV1Lister) ClusterRoleListerReturnsOnCall(i int, result1 v1.ClusterRoleLister) {
	fake.clusterRoleListerMutex.Lock()
	defer fake.clusterRoleListerMutex.Unlock()
	fake.ClusterRoleListerStub = nil
	if fake.clusterRoleListerReturnsOnCall == nil {
		fake.clusterRoleListerReturnsOnCall = make(map[int]struct {
			result1 v1.ClusterRoleLister
		})
	}
	fake.clusterRoleListerReturnsOnCall[i] = struct {
		result1 v1.ClusterRoleLister
	}{result1}
}

func (fake *FakeRbacV1Lister) RegisterClusterRoleBindingLister(arg1 v1.ClusterRoleBindingLister) {
	fake.registerClusterRoleBindingListerMutex.Lock()
	fake.registerClusterRoleBindingListerArgsForCall = append(fake.registerClusterRoleBindingListerArgsForCall, struct {
		arg1 v1.ClusterRoleBindingLister
	}{arg1})
	stub := fake.RegisterClusterRoleBindingListerStub
	fake.recordInvocation("RegisterClusterRoleBindingLister", []interface{}{arg1})
	fake.registerClusterRoleBindingListerMutex.Unlock()
	if stub != nil {
		fake.RegisterClusterRoleBindingListerStub(arg1)
	}
}

func (fake *FakeRbacV1Lister) RegisterClusterRoleBindingListerCallCount() int {
	fake.registerClusterRoleBindingListerMutex.RLock()
	defer fake.registerClusterRoleBindingListerMutex.RUnlock()
	return len(fake.registerClusterRoleBindingListerArgsForCall)
}

func (fake *FakeRbacV1Lister) RegisterClusterRoleBindingListerCalls(stub func(v1.ClusterRoleBindingLister)) {
	fake.registerClusterRoleBindingListerMutex.Lock()
	defer fake.registerClusterRoleBindingListerMutex.Unlock()
	fake.RegisterClusterRoleBindingListerStub = stub
}

func (fake *FakeRbacV1Lister) RegisterClusterRoleBindingListerArgsForCall(i int) v1.ClusterRoleBindingLister {
	fake.registerClusterRoleBindingListerMutex.RLock()
	defer fake.registerClusterRoleBindingListerMutex.RUnlock()
	argsForCall := fake.registerClusterRoleBindingListerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRbacV1Lister) RegisterClusterRoleLister(arg1 v1.ClusterRoleLister) {
	fake.registerClusterRoleListerMutex.Lock()
	fake.registerClusterRoleListerArgsForCall = append(fake.registerClusterRoleListerArgsForCall, struct {
		arg1 v1.ClusterRoleLister
	}{arg1})
	stub := fake.RegisterClusterRoleListerStub
	fake.recordInvocation("RegisterClusterRoleLister", []interface{}{arg1})
	fake.registerClusterRoleListerMutex.Unlock()
	if stub != nil {
		fake.RegisterClusterRoleListerStub(arg1)
	}
}

func (fake *FakeRbacV1Lister) RegisterClusterRoleListerCallCount() int {
	fake.registerClusterRoleListerMutex.RLock()
	defer fake.registerClusterRoleListerMutex.RUnlock()
	return len(fake.registerClusterRoleListerArgsForCall)
}

func (fake *FakeRbacV1Lister) RegisterClusterRoleListerCalls(stub func(v1.ClusterRoleLister)) {
	fake.registerClusterRoleListerMutex.Lock()
	defer fake.registerClusterRoleListerMutex.Unlock()
	fake.RegisterClusterRoleListerStub = stub
}

func (fake *FakeRbacV1Lister) RegisterClusterRoleListerArgsForCall(i int) v1.ClusterRoleLister {
	fake.registerClusterRoleListerMutex.RLock()
	defer fake.registerClusterRoleListerMutex.RUnlock()
	argsForCall := fake.registerClusterRoleListerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRbacV1Lister) RegisterRoleBindingLister(arg1 string, arg2 v1.RoleBindingLister) {
	fake.registerRoleBindingListerMutex.Lock()
	fake.registerRoleBindingListerArgsForCall = append(fake.registerRoleBindingListerArgsForCall, struct {
		arg1 string
		arg2 v1.RoleBindingLister
	}{arg1, arg2})
	stub := fake.RegisterRoleBindingListerStub
	fake.recordInvocation("RegisterRoleBindingLister", []interface{}{arg1, arg2})
	fake.registerRoleBindingListerMutex.Unlock()
	if stub != nil {
		fake.RegisterRoleBindingListerStub(arg1, arg2)
	}
}

func (fake *FakeRbacV1Lister) RegisterRoleBindingListerCallCount() int {
	fake.registerRoleBindingListerMutex.RLock()
	defer fake.registerRoleBindingListerMutex.RUnlock()
	return len(fake.registerRoleBindingListerArgsForCall)
}

func (fake *FakeRbacV1Lister) RegisterRoleBindingListerCalls(stub func(string, v1.RoleBindingLister)) {
	fake.registerRoleBindingListerMutex.Lock()
	defer fake.registerRoleBindingListerMutex.Unlock()
	fake.RegisterRoleBindingListerStub = stub
}

func (fake *FakeRbacV1Lister) RegisterRoleBindingListerArgsForCall(i int) (string, v1.RoleBindingLister) {
	fake.registerRoleBindingListerMutex.RLock()
	defer fake.registerRoleBindingListerMutex.RUnlock()
	argsForCall := fake.registerRoleBindingListerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRbacV1Lister) RegisterRoleLister(arg1 string, arg2 v1.RoleLister) {
	fake.registerRoleListerMutex.Lock()
	fake.registerRoleListerArgsForCall = append(fake.registerRoleListerArgsForCall, struct {
		arg1 string
		arg2 v1.RoleLister
	}{arg1, arg2})
	stub := fake.RegisterRoleListerStub
	fake.recordInvocation("RegisterRoleLister", []interface{}{arg1, arg2})
	fake.registerRoleListerMutex.Unlock()
	if stub != nil {
		fake.RegisterRoleListerStub(arg1, arg2)
	}
}

func (fake *FakeRbacV1Lister) RegisterRoleListerCallCount() int {
	fake.registerRoleListerMutex.RLock()
	defer fake.registerRoleListerMutex.RUnlock()
	return len(fake.registerRoleListerArgsForCall)
}

func (fake *FakeRbacV1Lister) RegisterRoleListerCalls(stub func(string, v1.RoleLister)) {
	fake.registerRoleListerMutex.Lock()
	defer fake.registerRoleListerMutex.Unlock()
	fake.RegisterRoleListerStub = stub
}

func (fake *FakeRbacV1Lister) RegisterRoleListerArgsForCall(i int) (string, v1.RoleLister) {
	fake.registerRoleListerMutex.RLock()
	defer fake.registerRoleListerMutex.RUnlock()
	argsForCall := fake.registerRoleListerArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRbacV1Lister) RoleBindingLister() v1.RoleBindingLister {
	fake.roleBindingListerMutex.Lock()
	ret, specificReturn := fake.roleBindingListerReturnsOnCall[len(fake.roleBindingListerArgsForCall)]
	fake.roleBindingListerArgsForCall = append(fake.roleBindingListerArgsForCall, struct {
	}{})
	stub := fake.RoleBindingListerStub
	fakeReturns := fake.roleBindingListerReturns
	fake.recordInvocation("RoleBindingLister", []interface{}{})
	fake.roleBindingListerMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRbacV1Lister) RoleBindingListerCallCount() int {
	fake.roleBindingListerMutex.RLock()
	defer fake.roleBindingListerMutex.RUnlock()
	return len(fake.roleBindingListerArgsForCall)
}

func (fake *FakeRbacV1Lister) RoleBindingListerCalls(stub func() v1.RoleBindingLister) {
	fake.roleBindingListerMutex.Lock()
	defer fake.roleBindingListerMutex.Unlock()
	fake.RoleBindingListerStub = stub
}

func (fake *FakeRbacV1Lister) RoleBindingListerReturns(result1 v1.RoleBindingLister) {
	fake.roleBindingListerMutex.Lock()
	defer fake.roleBindingListerMutex.Unlock()
	fake.RoleBindingListerStub = nil
	fake.roleBindingListerReturns = struct {
		result1 v1.RoleBindingLister
	}{result1}
}

func (fake *FakeRbacV1Lister) RoleBindingListerReturnsOnCall(i int, result1 v1.RoleBindingLister) {
	fake.roleBindingListerMutex.Lock()
	defer fake.roleBindingListerMutex.Unlock()
	fake.RoleBindingListerStub = nil
	if fake.roleBindingListerReturnsOnCall == nil {
		fake.roleBindingListerReturnsOnCall = make(map[int]struct {
			result1 v1.RoleBindingLister
		})
	}
	fake.roleBindingListerReturnsOnCall[i] = struct {
		result1 v1.RoleBindingLister
	}{result1}
}

func (fake *FakeRbacV1Lister) RoleLister() v1.RoleLister {
	fake.roleListerMutex.Lock()
	ret, specificReturn := fake.roleListerReturnsOnCall[len(fake.roleListerArgsForCall)]
	fake.roleListerArgsForCall = append(fake.roleListerArgsForCall, struct {
	}{})
	stub := fake.RoleListerStub
	fakeReturns := fake.roleListerReturns
	fake.recordInvocation("RoleLister", []interface{}{})
	fake.roleListerMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeRbacV1Lister) RoleListerCallCount() int {
	fake.roleListerMutex.RLock()
	defer fake.roleListerMutex.RUnlock()
	return len(fake.roleListerArgsForCall)
}

func (fake *FakeRbacV1Lister) RoleListerCalls(stub func() v1.RoleLister) {
	fake.roleListerMutex.Lock()
	defer fake.roleListerMutex.Unlock()
	fake.RoleListerStub = stub
}

func (fake *FakeRbacV1Lister) RoleListerReturns(result1 v1.RoleLister) {
	fake.roleListerMutex.Lock()
	defer fake.roleListerMutex.Unlock()
	fake.RoleListerStub = nil
	fake.roleListerReturns = struct {
		result1 v1.RoleLister
	}{result1}
}

func (fake *FakeRbacV1Lister) RoleListerReturnsOnCall(i int, result1 v1.RoleLister) {
	fake.roleListerMutex.Lock()
	defer fake.roleListerMutex.Unlock()
	fake.RoleListerStub = nil
	if fake.roleListerReturnsOnCall == nil {
		fake.roleListerReturnsOnCall = make(map[int]struct {
			result1 v1.RoleLister
		})
	}
	fake.roleListerReturnsOnCall[i] = struct {
		result1 v1.RoleLister
	}{result1}
}

func (fake *FakeRbacV1Lister) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.clusterRoleBindingListerMutex.RLock()
	defer fake.clusterRoleBindingListerMutex.RUnlock()
	fake.clusterRoleListerMutex.RLock()
	defer fake.clusterRoleListerMutex.RUnlock()
	fake.registerClusterRoleBindingListerMutex.RLock()
	defer fake.registerClusterRoleBindingListerMutex.RUnlock()
	fake.registerClusterRoleListerMutex.RLock()
	defer fake.registerClusterRoleListerMutex.RUnlock()
	fake.registerRoleBindingListerMutex.RLock()
	defer fake.registerRoleBindingListerMutex.RUnlock()
	fake.registerRoleListerMutex.RLock()
	defer fake.registerRoleListerMutex.RUnlock()
	fake.roleBindingListerMutex.RLock()
	defer fake.roleBindingListerMutex.RUnlock()
	fake.roleListerMutex.RLock()
	defer fake.roleListerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRbacV1Lister) recordInvocation(key string, args []interface{}) {
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

var _ operatorlister.RbacV1Lister = new(FakeRbacV1Lister)
