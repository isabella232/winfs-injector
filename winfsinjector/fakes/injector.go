// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"
)

type Injector struct {
	AddReleaseToMetadataStub        func(releasePath, releaseName, releaseVersion, extractedTileDir string) error
	addReleaseToMetadataMutex       sync.RWMutex
	addReleaseToMetadataArgsForCall []struct {
		releasePath      string
		releaseName      string
		releaseVersion   string
		extractedTileDir string
	}
	addReleaseToMetadataReturns struct {
		result1 error
	}
	addReleaseToMetadataReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Injector) AddReleaseToMetadata(releasePath string, releaseName string, releaseVersion string, extractedTileDir string) error {
	fake.addReleaseToMetadataMutex.Lock()
	ret, specificReturn := fake.addReleaseToMetadataReturnsOnCall[len(fake.addReleaseToMetadataArgsForCall)]
	fake.addReleaseToMetadataArgsForCall = append(fake.addReleaseToMetadataArgsForCall, struct {
		releasePath      string
		releaseName      string
		releaseVersion   string
		extractedTileDir string
	}{releasePath, releaseName, releaseVersion, extractedTileDir})
	fake.recordInvocation("AddReleaseToMetadata", []interface{}{releasePath, releaseName, releaseVersion, extractedTileDir})
	fake.addReleaseToMetadataMutex.Unlock()
	if fake.AddReleaseToMetadataStub != nil {
		return fake.AddReleaseToMetadataStub(releasePath, releaseName, releaseVersion, extractedTileDir)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.addReleaseToMetadataReturns.result1
}

func (fake *Injector) AddReleaseToMetadataCallCount() int {
	fake.addReleaseToMetadataMutex.RLock()
	defer fake.addReleaseToMetadataMutex.RUnlock()
	return len(fake.addReleaseToMetadataArgsForCall)
}

func (fake *Injector) AddReleaseToMetadataArgsForCall(i int) (string, string, string, string) {
	fake.addReleaseToMetadataMutex.RLock()
	defer fake.addReleaseToMetadataMutex.RUnlock()
	return fake.addReleaseToMetadataArgsForCall[i].releasePath, fake.addReleaseToMetadataArgsForCall[i].releaseName, fake.addReleaseToMetadataArgsForCall[i].releaseVersion, fake.addReleaseToMetadataArgsForCall[i].extractedTileDir
}

func (fake *Injector) AddReleaseToMetadataReturns(result1 error) {
	fake.AddReleaseToMetadataStub = nil
	fake.addReleaseToMetadataReturns = struct {
		result1 error
	}{result1}
}

func (fake *Injector) AddReleaseToMetadataReturnsOnCall(i int, result1 error) {
	fake.AddReleaseToMetadataStub = nil
	if fake.addReleaseToMetadataReturnsOnCall == nil {
		fake.addReleaseToMetadataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addReleaseToMetadataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *Injector) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addReleaseToMetadataMutex.RLock()
	defer fake.addReleaseToMetadataMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Injector) recordInvocation(key string, args []interface{}) {
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