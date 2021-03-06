// This file was generated by counterfeiter
package authfakes

import (
	"net/http"
	"sync"

	"github.com/concourse/atc/auth"
)

type FakeUserContextReader struct {
	GetTeamStub        func(r *http.Request) (string, int, bool, bool)
	getTeamMutex       sync.RWMutex
	getTeamArgsForCall []struct {
		r *http.Request
	}
	getTeamReturns struct {
		result1 string
		result2 int
		result3 bool
		result4 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUserContextReader) GetTeam(r *http.Request) (string, int, bool, bool) {
	fake.getTeamMutex.Lock()
	fake.getTeamArgsForCall = append(fake.getTeamArgsForCall, struct {
		r *http.Request
	}{r})
	fake.recordInvocation("GetTeam", []interface{}{r})
	fake.getTeamMutex.Unlock()
	if fake.GetTeamStub != nil {
		return fake.GetTeamStub(r)
	} else {
		return fake.getTeamReturns.result1, fake.getTeamReturns.result2, fake.getTeamReturns.result3, fake.getTeamReturns.result4
	}
}

func (fake *FakeUserContextReader) GetTeamCallCount() int {
	fake.getTeamMutex.RLock()
	defer fake.getTeamMutex.RUnlock()
	return len(fake.getTeamArgsForCall)
}

func (fake *FakeUserContextReader) GetTeamArgsForCall(i int) *http.Request {
	fake.getTeamMutex.RLock()
	defer fake.getTeamMutex.RUnlock()
	return fake.getTeamArgsForCall[i].r
}

func (fake *FakeUserContextReader) GetTeamReturns(result1 string, result2 int, result3 bool, result4 bool) {
	fake.GetTeamStub = nil
	fake.getTeamReturns = struct {
		result1 string
		result2 int
		result3 bool
		result4 bool
	}{result1, result2, result3, result4}
}

func (fake *FakeUserContextReader) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTeamMutex.RLock()
	defer fake.getTeamMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeUserContextReader) recordInvocation(key string, args []interface{}) {
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

var _ auth.UserContextReader = new(FakeUserContextReader)
