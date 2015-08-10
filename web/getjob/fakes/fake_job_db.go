// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/db"
	"github.com/concourse/atc/web/getjob"
)

type FakeJobDB struct {
	GetConfigStub        func() (atc.Config, db.ConfigVersion, error)
	getConfigMutex       sync.RWMutex
	getConfigArgsForCall []struct{}
	getConfigReturns     struct {
		result1 atc.Config
		result2 db.ConfigVersion
		result3 error
	}
	GetJobStub        func(string) (db.SavedJob, error)
	getJobMutex       sync.RWMutex
	getJobArgsForCall []struct {
		arg1 string
	}
	getJobReturns struct {
		result1 db.SavedJob
		result2 error
	}
	GetAllJobBuildsStub        func(job string) ([]db.Build, error)
	getAllJobBuildsMutex       sync.RWMutex
	getAllJobBuildsArgsForCall []struct {
		job string
	}
	getAllJobBuildsReturns struct {
		result1 []db.Build
		result2 error
	}
	GetCurrentBuildStub        func(job string) (db.Build, error)
	getCurrentBuildMutex       sync.RWMutex
	getCurrentBuildArgsForCall []struct {
		job string
	}
	getCurrentBuildReturns struct {
		result1 db.Build
		result2 error
	}
	GetPipelineNameStub        func() string
	getPipelineNameMutex       sync.RWMutex
	getPipelineNameArgsForCall []struct{}
	getPipelineNameReturns     struct {
		result1 string
	}
	GetBuildResourcesStub        func(buildID int) ([]db.BuildInput, []db.BuildOutput, error)
	getBuildResourcesMutex       sync.RWMutex
	getBuildResourcesArgsForCall []struct {
		buildID int
	}
	getBuildResourcesReturns struct {
		result1 []db.BuildInput
		result2 []db.BuildOutput
		result3 error
	}
}

func (fake *FakeJobDB) GetConfig() (atc.Config, db.ConfigVersion, error) {
	fake.getConfigMutex.Lock()
	fake.getConfigArgsForCall = append(fake.getConfigArgsForCall, struct{}{})
	fake.getConfigMutex.Unlock()
	if fake.GetConfigStub != nil {
		return fake.GetConfigStub()
	} else {
		return fake.getConfigReturns.result1, fake.getConfigReturns.result2, fake.getConfigReturns.result3
	}
}

func (fake *FakeJobDB) GetConfigCallCount() int {
	fake.getConfigMutex.RLock()
	defer fake.getConfigMutex.RUnlock()
	return len(fake.getConfigArgsForCall)
}

func (fake *FakeJobDB) GetConfigReturns(result1 atc.Config, result2 db.ConfigVersion, result3 error) {
	fake.GetConfigStub = nil
	fake.getConfigReturns = struct {
		result1 atc.Config
		result2 db.ConfigVersion
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeJobDB) GetJob(arg1 string) (db.SavedJob, error) {
	fake.getJobMutex.Lock()
	fake.getJobArgsForCall = append(fake.getJobArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.getJobMutex.Unlock()
	if fake.GetJobStub != nil {
		return fake.GetJobStub(arg1)
	} else {
		return fake.getJobReturns.result1, fake.getJobReturns.result2
	}
}

func (fake *FakeJobDB) GetJobCallCount() int {
	fake.getJobMutex.RLock()
	defer fake.getJobMutex.RUnlock()
	return len(fake.getJobArgsForCall)
}

func (fake *FakeJobDB) GetJobArgsForCall(i int) string {
	fake.getJobMutex.RLock()
	defer fake.getJobMutex.RUnlock()
	return fake.getJobArgsForCall[i].arg1
}

func (fake *FakeJobDB) GetJobReturns(result1 db.SavedJob, result2 error) {
	fake.GetJobStub = nil
	fake.getJobReturns = struct {
		result1 db.SavedJob
		result2 error
	}{result1, result2}
}

func (fake *FakeJobDB) GetAllJobBuilds(job string) ([]db.Build, error) {
	fake.getAllJobBuildsMutex.Lock()
	fake.getAllJobBuildsArgsForCall = append(fake.getAllJobBuildsArgsForCall, struct {
		job string
	}{job})
	fake.getAllJobBuildsMutex.Unlock()
	if fake.GetAllJobBuildsStub != nil {
		return fake.GetAllJobBuildsStub(job)
	} else {
		return fake.getAllJobBuildsReturns.result1, fake.getAllJobBuildsReturns.result2
	}
}

func (fake *FakeJobDB) GetAllJobBuildsCallCount() int {
	fake.getAllJobBuildsMutex.RLock()
	defer fake.getAllJobBuildsMutex.RUnlock()
	return len(fake.getAllJobBuildsArgsForCall)
}

func (fake *FakeJobDB) GetAllJobBuildsArgsForCall(i int) string {
	fake.getAllJobBuildsMutex.RLock()
	defer fake.getAllJobBuildsMutex.RUnlock()
	return fake.getAllJobBuildsArgsForCall[i].job
}

func (fake *FakeJobDB) GetAllJobBuildsReturns(result1 []db.Build, result2 error) {
	fake.GetAllJobBuildsStub = nil
	fake.getAllJobBuildsReturns = struct {
		result1 []db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeJobDB) GetCurrentBuild(job string) (db.Build, error) {
	fake.getCurrentBuildMutex.Lock()
	fake.getCurrentBuildArgsForCall = append(fake.getCurrentBuildArgsForCall, struct {
		job string
	}{job})
	fake.getCurrentBuildMutex.Unlock()
	if fake.GetCurrentBuildStub != nil {
		return fake.GetCurrentBuildStub(job)
	} else {
		return fake.getCurrentBuildReturns.result1, fake.getCurrentBuildReturns.result2
	}
}

func (fake *FakeJobDB) GetCurrentBuildCallCount() int {
	fake.getCurrentBuildMutex.RLock()
	defer fake.getCurrentBuildMutex.RUnlock()
	return len(fake.getCurrentBuildArgsForCall)
}

func (fake *FakeJobDB) GetCurrentBuildArgsForCall(i int) string {
	fake.getCurrentBuildMutex.RLock()
	defer fake.getCurrentBuildMutex.RUnlock()
	return fake.getCurrentBuildArgsForCall[i].job
}

func (fake *FakeJobDB) GetCurrentBuildReturns(result1 db.Build, result2 error) {
	fake.GetCurrentBuildStub = nil
	fake.getCurrentBuildReturns = struct {
		result1 db.Build
		result2 error
	}{result1, result2}
}

func (fake *FakeJobDB) GetPipelineName() string {
	fake.getPipelineNameMutex.Lock()
	fake.getPipelineNameArgsForCall = append(fake.getPipelineNameArgsForCall, struct{}{})
	fake.getPipelineNameMutex.Unlock()
	if fake.GetPipelineNameStub != nil {
		return fake.GetPipelineNameStub()
	} else {
		return fake.getPipelineNameReturns.result1
	}
}

func (fake *FakeJobDB) GetPipelineNameCallCount() int {
	fake.getPipelineNameMutex.RLock()
	defer fake.getPipelineNameMutex.RUnlock()
	return len(fake.getPipelineNameArgsForCall)
}

func (fake *FakeJobDB) GetPipelineNameReturns(result1 string) {
	fake.GetPipelineNameStub = nil
	fake.getPipelineNameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeJobDB) GetBuildResources(buildID int) ([]db.BuildInput, []db.BuildOutput, error) {
	fake.getBuildResourcesMutex.Lock()
	fake.getBuildResourcesArgsForCall = append(fake.getBuildResourcesArgsForCall, struct {
		buildID int
	}{buildID})
	fake.getBuildResourcesMutex.Unlock()
	if fake.GetBuildResourcesStub != nil {
		return fake.GetBuildResourcesStub(buildID)
	} else {
		return fake.getBuildResourcesReturns.result1, fake.getBuildResourcesReturns.result2, fake.getBuildResourcesReturns.result3
	}
}

func (fake *FakeJobDB) GetBuildResourcesCallCount() int {
	fake.getBuildResourcesMutex.RLock()
	defer fake.getBuildResourcesMutex.RUnlock()
	return len(fake.getBuildResourcesArgsForCall)
}

func (fake *FakeJobDB) GetBuildResourcesArgsForCall(i int) int {
	fake.getBuildResourcesMutex.RLock()
	defer fake.getBuildResourcesMutex.RUnlock()
	return fake.getBuildResourcesArgsForCall[i].buildID
}

func (fake *FakeJobDB) GetBuildResourcesReturns(result1 []db.BuildInput, result2 []db.BuildOutput, result3 error) {
	fake.GetBuildResourcesStub = nil
	fake.getBuildResourcesReturns = struct {
		result1 []db.BuildInput
		result2 []db.BuildOutput
		result3 error
	}{result1, result2, result3}
}

var _ getjob.JobDB = new(FakeJobDB)
