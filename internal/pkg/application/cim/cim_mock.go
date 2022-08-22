// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package cim

import (
	"context"
	"github.com/diwise/context-broker/pkg/ngsild"
	"github.com/diwise/context-broker/pkg/ngsild/types"
	"sync"
)

// Ensure, that ContextInformationManagerMock does implement ContextInformationManager.
// If this is not the case, regenerate this file with moq.
var _ ContextInformationManager = &ContextInformationManagerMock{}

// ContextInformationManagerMock is a mock implementation of ContextInformationManager.
//
//	func TestSomethingThatUsesContextInformationManager(t *testing.T) {
//
//		// make and configure a mocked ContextInformationManager
//		mockedContextInformationManager := &ContextInformationManagerMock{
//			CreateEntityFunc: func(ctx context.Context, tenant string, entity types.Entity, headers map[string][]string) (*ngsild.CreateEntityResult, error) {
//				panic("mock out the CreateEntity method")
//			},
//			QueryEntitiesFunc: func(ctx context.Context, tenant string, entityTypes []string, entityAttributes []string, query string, headers map[string][]string) (*ngsild.QueryEntitiesResult, error) {
//				panic("mock out the QueryEntities method")
//			},
//			RetrieveEntityFunc: func(ctx context.Context, tenant string, entityID string, headers map[string][]string) (types.Entity, error) {
//				panic("mock out the RetrieveEntity method")
//			},
//			StartFunc: func() error {
//				panic("mock out the Start method")
//			},
//			StopFunc: func() error {
//				panic("mock out the Stop method")
//			},
//			UpdateEntityAttributesFunc: func(ctx context.Context, tenant string, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.UpdateEntityAttributesResult, error) {
//				panic("mock out the UpdateEntityAttributes method")
//			},
//		}
//
//		// use mockedContextInformationManager in code that requires ContextInformationManager
//		// and then make assertions.
//
//	}
type ContextInformationManagerMock struct {
	// CreateEntityFunc mocks the CreateEntity method.
	CreateEntityFunc func(ctx context.Context, tenant string, entity types.Entity, headers map[string][]string) (*ngsild.CreateEntityResult, error)

	// QueryEntitiesFunc mocks the QueryEntities method.
	QueryEntitiesFunc func(ctx context.Context, tenant string, entityTypes []string, entityAttributes []string, query string, headers map[string][]string) (*ngsild.QueryEntitiesResult, error)

	// RetrieveEntityFunc mocks the RetrieveEntity method.
	RetrieveEntityFunc func(ctx context.Context, tenant string, entityID string, headers map[string][]string) (types.Entity, error)

	// StartFunc mocks the Start method.
	StartFunc func() error

	// StopFunc mocks the Stop method.
	StopFunc func() error

	// UpdateEntityAttributesFunc mocks the UpdateEntityAttributes method.
	UpdateEntityAttributesFunc func(ctx context.Context, tenant string, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.UpdateEntityAttributesResult, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateEntity holds details about calls to the CreateEntity method.
		CreateEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// Entity is the entity argument value.
			Entity types.Entity
			// Headers is the headers argument value.
			Headers map[string][]string
		}
		// QueryEntities holds details about calls to the QueryEntities method.
		QueryEntities []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// EntityTypes is the entityTypes argument value.
			EntityTypes []string
			// EntityAttributes is the entityAttributes argument value.
			EntityAttributes []string
			// Query is the query argument value.
			Query string
			// Headers is the headers argument value.
			Headers map[string][]string
		}
		// RetrieveEntity holds details about calls to the RetrieveEntity method.
		RetrieveEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// EntityID is the entityID argument value.
			EntityID string
			// Headers is the headers argument value.
			Headers map[string][]string
		}
		// Start holds details about calls to the Start method.
		Start []struct {
		}
		// Stop holds details about calls to the Stop method.
		Stop []struct {
		}
		// UpdateEntityAttributes holds details about calls to the UpdateEntityAttributes method.
		UpdateEntityAttributes []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Tenant is the tenant argument value.
			Tenant string
			// EntityID is the entityID argument value.
			EntityID string
			// Fragment is the fragment argument value.
			Fragment types.EntityFragment
			// Headers is the headers argument value.
			Headers map[string][]string
		}
	}
	lockCreateEntity           sync.RWMutex
	lockQueryEntities          sync.RWMutex
	lockRetrieveEntity         sync.RWMutex
	lockStart                  sync.RWMutex
	lockStop                   sync.RWMutex
	lockUpdateEntityAttributes sync.RWMutex
}

// CreateEntity calls CreateEntityFunc.
func (mock *ContextInformationManagerMock) CreateEntity(ctx context.Context, tenant string, entity types.Entity, headers map[string][]string) (*ngsild.CreateEntityResult, error) {
	if mock.CreateEntityFunc == nil {
		panic("ContextInformationManagerMock.CreateEntityFunc: method is nil but ContextInformationManager.CreateEntity was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Tenant  string
		Entity  types.Entity
		Headers map[string][]string
	}{
		Ctx:     ctx,
		Tenant:  tenant,
		Entity:  entity,
		Headers: headers,
	}
	mock.lockCreateEntity.Lock()
	mock.calls.CreateEntity = append(mock.calls.CreateEntity, callInfo)
	mock.lockCreateEntity.Unlock()
	return mock.CreateEntityFunc(ctx, tenant, entity, headers)
}

// CreateEntityCalls gets all the calls that were made to CreateEntity.
// Check the length with:
//
//	len(mockedContextInformationManager.CreateEntityCalls())
func (mock *ContextInformationManagerMock) CreateEntityCalls() []struct {
	Ctx     context.Context
	Tenant  string
	Entity  types.Entity
	Headers map[string][]string
} {
	var calls []struct {
		Ctx     context.Context
		Tenant  string
		Entity  types.Entity
		Headers map[string][]string
	}
	mock.lockCreateEntity.RLock()
	calls = mock.calls.CreateEntity
	mock.lockCreateEntity.RUnlock()
	return calls
}

// QueryEntities calls QueryEntitiesFunc.
func (mock *ContextInformationManagerMock) QueryEntities(ctx context.Context, tenant string, entityTypes []string, entityAttributes []string, query string, headers map[string][]string) (*ngsild.QueryEntitiesResult, error) {
	if mock.QueryEntitiesFunc == nil {
		panic("ContextInformationManagerMock.QueryEntitiesFunc: method is nil but ContextInformationManager.QueryEntities was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		Tenant           string
		EntityTypes      []string
		EntityAttributes []string
		Query            string
		Headers          map[string][]string
	}{
		Ctx:              ctx,
		Tenant:           tenant,
		EntityTypes:      entityTypes,
		EntityAttributes: entityAttributes,
		Query:            query,
		Headers:          headers,
	}
	mock.lockQueryEntities.Lock()
	mock.calls.QueryEntities = append(mock.calls.QueryEntities, callInfo)
	mock.lockQueryEntities.Unlock()
	return mock.QueryEntitiesFunc(ctx, tenant, entityTypes, entityAttributes, query, headers)
}

// QueryEntitiesCalls gets all the calls that were made to QueryEntities.
// Check the length with:
//
//	len(mockedContextInformationManager.QueryEntitiesCalls())
func (mock *ContextInformationManagerMock) QueryEntitiesCalls() []struct {
	Ctx              context.Context
	Tenant           string
	EntityTypes      []string
	EntityAttributes []string
	Query            string
	Headers          map[string][]string
} {
	var calls []struct {
		Ctx              context.Context
		Tenant           string
		EntityTypes      []string
		EntityAttributes []string
		Query            string
		Headers          map[string][]string
	}
	mock.lockQueryEntities.RLock()
	calls = mock.calls.QueryEntities
	mock.lockQueryEntities.RUnlock()
	return calls
}

// RetrieveEntity calls RetrieveEntityFunc.
func (mock *ContextInformationManagerMock) RetrieveEntity(ctx context.Context, tenant string, entityID string, headers map[string][]string) (types.Entity, error) {
	if mock.RetrieveEntityFunc == nil {
		panic("ContextInformationManagerMock.RetrieveEntityFunc: method is nil but ContextInformationManager.RetrieveEntity was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
		Headers  map[string][]string
	}{
		Ctx:      ctx,
		Tenant:   tenant,
		EntityID: entityID,
		Headers:  headers,
	}
	mock.lockRetrieveEntity.Lock()
	mock.calls.RetrieveEntity = append(mock.calls.RetrieveEntity, callInfo)
	mock.lockRetrieveEntity.Unlock()
	return mock.RetrieveEntityFunc(ctx, tenant, entityID, headers)
}

// RetrieveEntityCalls gets all the calls that were made to RetrieveEntity.
// Check the length with:
//
//	len(mockedContextInformationManager.RetrieveEntityCalls())
func (mock *ContextInformationManagerMock) RetrieveEntityCalls() []struct {
	Ctx      context.Context
	Tenant   string
	EntityID string
	Headers  map[string][]string
} {
	var calls []struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
		Headers  map[string][]string
	}
	mock.lockRetrieveEntity.RLock()
	calls = mock.calls.RetrieveEntity
	mock.lockRetrieveEntity.RUnlock()
	return calls
}

// Start calls StartFunc.
func (mock *ContextInformationManagerMock) Start() error {
	if mock.StartFunc == nil {
		panic("ContextInformationManagerMock.StartFunc: method is nil but ContextInformationManager.Start was just called")
	}
	callInfo := struct {
	}{}
	mock.lockStart.Lock()
	mock.calls.Start = append(mock.calls.Start, callInfo)
	mock.lockStart.Unlock()
	return mock.StartFunc()
}

// StartCalls gets all the calls that were made to Start.
// Check the length with:
//
//	len(mockedContextInformationManager.StartCalls())
func (mock *ContextInformationManagerMock) StartCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStart.RLock()
	calls = mock.calls.Start
	mock.lockStart.RUnlock()
	return calls
}

// Stop calls StopFunc.
func (mock *ContextInformationManagerMock) Stop() error {
	if mock.StopFunc == nil {
		panic("ContextInformationManagerMock.StopFunc: method is nil but ContextInformationManager.Stop was just called")
	}
	callInfo := struct {
	}{}
	mock.lockStop.Lock()
	mock.calls.Stop = append(mock.calls.Stop, callInfo)
	mock.lockStop.Unlock()
	return mock.StopFunc()
}

// StopCalls gets all the calls that were made to Stop.
// Check the length with:
//
//	len(mockedContextInformationManager.StopCalls())
func (mock *ContextInformationManagerMock) StopCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStop.RLock()
	calls = mock.calls.Stop
	mock.lockStop.RUnlock()
	return calls
}

// UpdateEntityAttributes calls UpdateEntityAttributesFunc.
func (mock *ContextInformationManagerMock) UpdateEntityAttributes(ctx context.Context, tenant string, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.UpdateEntityAttributesResult, error) {
	if mock.UpdateEntityAttributesFunc == nil {
		panic("ContextInformationManagerMock.UpdateEntityAttributesFunc: method is nil but ContextInformationManager.UpdateEntityAttributes was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
		Fragment types.EntityFragment
		Headers  map[string][]string
	}{
		Ctx:      ctx,
		Tenant:   tenant,
		EntityID: entityID,
		Fragment: fragment,
		Headers:  headers,
	}
	mock.lockUpdateEntityAttributes.Lock()
	mock.calls.UpdateEntityAttributes = append(mock.calls.UpdateEntityAttributes, callInfo)
	mock.lockUpdateEntityAttributes.Unlock()
	return mock.UpdateEntityAttributesFunc(ctx, tenant, entityID, fragment, headers)
}

// UpdateEntityAttributesCalls gets all the calls that were made to UpdateEntityAttributes.
// Check the length with:
//
//	len(mockedContextInformationManager.UpdateEntityAttributesCalls())
func (mock *ContextInformationManagerMock) UpdateEntityAttributesCalls() []struct {
	Ctx      context.Context
	Tenant   string
	EntityID string
	Fragment types.EntityFragment
	Headers  map[string][]string
} {
	var calls []struct {
		Ctx      context.Context
		Tenant   string
		EntityID string
		Fragment types.EntityFragment
		Headers  map[string][]string
	}
	mock.lockUpdateEntityAttributes.RLock()
	calls = mock.calls.UpdateEntityAttributes
	mock.lockUpdateEntityAttributes.RUnlock()
	return calls
}
