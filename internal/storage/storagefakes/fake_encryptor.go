// Code generated by counterfeiter. DO NOT EDIT.
package storagefakes

import (
	"sync"

	"github.com/cloudfoundry/cloud-service-broker/internal/storage"
)

type FakeEncryptor struct {
	DecryptStub        func([]byte) ([]byte, error)
	decryptMutex       sync.RWMutex
	decryptArgsForCall []struct {
		arg1 []byte
	}
	decryptReturns struct {
		result1 []byte
		result2 error
	}
	decryptReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	EncryptStub        func([]byte) ([]byte, error)
	encryptMutex       sync.RWMutex
	encryptArgsForCall []struct {
		arg1 []byte
	}
	encryptReturns struct {
		result1 []byte
		result2 error
	}
	encryptReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEncryptor) Decrypt(arg1 []byte) ([]byte, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.decryptMutex.Lock()
	ret, specificReturn := fake.decryptReturnsOnCall[len(fake.decryptArgsForCall)]
	fake.decryptArgsForCall = append(fake.decryptArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.DecryptStub
	fakeReturns := fake.decryptReturns
	fake.recordInvocation("Decrypt", []interface{}{arg1Copy})
	fake.decryptMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEncryptor) DecryptCallCount() int {
	fake.decryptMutex.RLock()
	defer fake.decryptMutex.RUnlock()
	return len(fake.decryptArgsForCall)
}

func (fake *FakeEncryptor) DecryptCalls(stub func([]byte) ([]byte, error)) {
	fake.decryptMutex.Lock()
	defer fake.decryptMutex.Unlock()
	fake.DecryptStub = stub
}

func (fake *FakeEncryptor) DecryptArgsForCall(i int) []byte {
	fake.decryptMutex.RLock()
	defer fake.decryptMutex.RUnlock()
	argsForCall := fake.decryptArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEncryptor) DecryptReturns(result1 []byte, result2 error) {
	fake.decryptMutex.Lock()
	defer fake.decryptMutex.Unlock()
	fake.DecryptStub = nil
	fake.decryptReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeEncryptor) DecryptReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.decryptMutex.Lock()
	defer fake.decryptMutex.Unlock()
	fake.DecryptStub = nil
	if fake.decryptReturnsOnCall == nil {
		fake.decryptReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.decryptReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeEncryptor) Encrypt(arg1 []byte) ([]byte, error) {
	var arg1Copy []byte
	if arg1 != nil {
		arg1Copy = make([]byte, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.encryptMutex.Lock()
	ret, specificReturn := fake.encryptReturnsOnCall[len(fake.encryptArgsForCall)]
	fake.encryptArgsForCall = append(fake.encryptArgsForCall, struct {
		arg1 []byte
	}{arg1Copy})
	stub := fake.EncryptStub
	fakeReturns := fake.encryptReturns
	fake.recordInvocation("Encrypt", []interface{}{arg1Copy})
	fake.encryptMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEncryptor) EncryptCallCount() int {
	fake.encryptMutex.RLock()
	defer fake.encryptMutex.RUnlock()
	return len(fake.encryptArgsForCall)
}

func (fake *FakeEncryptor) EncryptCalls(stub func([]byte) ([]byte, error)) {
	fake.encryptMutex.Lock()
	defer fake.encryptMutex.Unlock()
	fake.EncryptStub = stub
}

func (fake *FakeEncryptor) EncryptArgsForCall(i int) []byte {
	fake.encryptMutex.RLock()
	defer fake.encryptMutex.RUnlock()
	argsForCall := fake.encryptArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEncryptor) EncryptReturns(result1 []byte, result2 error) {
	fake.encryptMutex.Lock()
	defer fake.encryptMutex.Unlock()
	fake.EncryptStub = nil
	fake.encryptReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeEncryptor) EncryptReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.encryptMutex.Lock()
	defer fake.encryptMutex.Unlock()
	fake.EncryptStub = nil
	if fake.encryptReturnsOnCall == nil {
		fake.encryptReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.encryptReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *FakeEncryptor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.decryptMutex.RLock()
	defer fake.decryptMutex.RUnlock()
	fake.encryptMutex.RLock()
	defer fake.encryptMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEncryptor) recordInvocation(key string, args []interface{}) {
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

var _ storage.Encryptor = new(FakeEncryptor)
