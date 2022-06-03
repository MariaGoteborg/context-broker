// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package test

import (
	"context"
	"github.com/diwise/context-broker/pkg/ngsild"
	"github.com/diwise/context-broker/pkg/ngsild/client"
	"github.com/diwise/context-broker/pkg/ngsild/types"
	"sync"
)

// Ensure, that ContextBrokerClientMock does implement ContextBrokerClient.
// If this is not the case, regenerate this file with moq.
var _ client.ContextBrokerClient = &ContextBrokerClientMock{}

// ContextBrokerClientMock is a mock implementation of ContextBrokerClient.
//
// 	func TestSomethingThatUsesContextBrokerClient(t *testing.T) {
//
// 		// make and configure a mocked ContextBrokerClient
// 		mockedContextBrokerClient := &ContextBrokerClientMock{
// 			CreateEntityFunc: func(ctx context.Context, entity types.Entity, headers map[string][]string) (*ngsild.CreateEntityResult, error) {
// 				panic("mock out the CreateEntity method")
// 			},
// 			QueryEntitiesFunc: func(ctx context.Context, entityTypes []string, entityAttributes []string, query string, headers map[string][]string) (*ngsild.QueryEntitiesResult, error) {
// 				panic("mock out the QueryEntities method")
// 			},
// 			RetrieveEntityFunc: func(ctx context.Context, entityID string, headers map[string][]string) (types.Entity, error) {
// 				panic("mock out the RetrieveEntity method")
// 			},
// 			UpdateEntityAttributesFunc: func(ctx context.Context, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.UpdateEntityAttributesResult, error) {
// 				panic("mock out the UpdateEntityAttributes method")
// 			},
// 		}
//
// 		// use mockedContextBrokerClient in code that requires ContextBrokerClient
// 		// and then make assertions.
//
// 	}
type ContextBrokerClientMock struct {
	// CreateEntityFunc mocks the CreateEntity method.
	CreateEntityFunc func(ctx context.Context, entity types.Entity, headers map[string][]string) (*ngsild.CreateEntityResult, error)

	// QueryEntitiesFunc mocks the QueryEntities method.
	QueryEntitiesFunc func(ctx context.Context, entityTypes []string, entityAttributes []string, query string, headers map[string][]string) (*ngsild.QueryEntitiesResult, error)

	// RetrieveEntityFunc mocks the RetrieveEntity method.
	RetrieveEntityFunc func(ctx context.Context, entityID string, headers map[string][]string) (types.Entity, error)

	// UpdateEntityAttributesFunc mocks the UpdateEntityAttributes method.
	UpdateEntityAttributesFunc func(ctx context.Context, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.UpdateEntityAttributesResult, error)

	// calls tracks calls to the methods.
	calls struct {
		// CreateEntity holds details about calls to the CreateEntity method.
		CreateEntity []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Entity is the entity argument value.
			Entity types.Entity
			// Headers is the headers argument value.
			Headers map[string][]string
		}
		// QueryEntities holds details about calls to the QueryEntities method.
		QueryEntities []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
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
			// EntityID is the entityID argument value.
			EntityID string
			// Headers is the headers argument value.
			Headers map[string][]string
		}
		// UpdateEntityAttributes holds details about calls to the UpdateEntityAttributes method.
		UpdateEntityAttributes []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
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
	lockUpdateEntityAttributes sync.RWMutex
}

// CreateEntity calls CreateEntityFunc.
func (mock *ContextBrokerClientMock) CreateEntity(ctx context.Context, entity types.Entity, headers map[string][]string) (*ngsild.CreateEntityResult, error) {
	if mock.CreateEntityFunc == nil {
		panic("ContextBrokerClientMock.CreateEntityFunc: method is nil but ContextBrokerClient.CreateEntity was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Entity  types.Entity
		Headers map[string][]string
	}{
		Ctx:     ctx,
		Entity:  entity,
		Headers: headers,
	}
	mock.lockCreateEntity.Lock()
	mock.calls.CreateEntity = append(mock.calls.CreateEntity, callInfo)
	mock.lockCreateEntity.Unlock()
	return mock.CreateEntityFunc(ctx, entity, headers)
}

// CreateEntityCalls gets all the calls that were made to CreateEntity.
// Check the length with:
//     len(mockedContextBrokerClient.CreateEntityCalls())
func (mock *ContextBrokerClientMock) CreateEntityCalls() []struct {
	Ctx     context.Context
	Entity  types.Entity
	Headers map[string][]string
} {
	var calls []struct {
		Ctx     context.Context
		Entity  types.Entity
		Headers map[string][]string
	}
	mock.lockCreateEntity.RLock()
	calls = mock.calls.CreateEntity
	mock.lockCreateEntity.RUnlock()
	return calls
}

// QueryEntities calls QueryEntitiesFunc.
func (mock *ContextBrokerClientMock) QueryEntities(ctx context.Context, entityTypes []string, entityAttributes []string, query string, headers map[string][]string) (*ngsild.QueryEntitiesResult, error) {
	if mock.QueryEntitiesFunc == nil {
		panic("ContextBrokerClientMock.QueryEntitiesFunc: method is nil but ContextBrokerClient.QueryEntities was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		EntityTypes      []string
		EntityAttributes []string
		Query            string
		Headers          map[string][]string
	}{
		Ctx:              ctx,
		EntityTypes:      entityTypes,
		EntityAttributes: entityAttributes,
		Query:            query,
		Headers:          headers,
	}
	mock.lockQueryEntities.Lock()
	mock.calls.QueryEntities = append(mock.calls.QueryEntities, callInfo)
	mock.lockQueryEntities.Unlock()
	return mock.QueryEntitiesFunc(ctx, entityTypes, entityAttributes, query, headers)
}

// QueryEntitiesCalls gets all the calls that were made to QueryEntities.
// Check the length with:
//     len(mockedContextBrokerClient.QueryEntitiesCalls())
func (mock *ContextBrokerClientMock) QueryEntitiesCalls() []struct {
	Ctx              context.Context
	EntityTypes      []string
	EntityAttributes []string
	Query            string
	Headers          map[string][]string
} {
	var calls []struct {
		Ctx              context.Context
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
func (mock *ContextBrokerClientMock) RetrieveEntity(ctx context.Context, entityID string, headers map[string][]string) (types.Entity, error) {
	if mock.RetrieveEntityFunc == nil {
		panic("ContextBrokerClientMock.RetrieveEntityFunc: method is nil but ContextBrokerClient.RetrieveEntity was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		EntityID string
		Headers  map[string][]string
	}{
		Ctx:      ctx,
		EntityID: entityID,
		Headers:  headers,
	}
	mock.lockRetrieveEntity.Lock()
	mock.calls.RetrieveEntity = append(mock.calls.RetrieveEntity, callInfo)
	mock.lockRetrieveEntity.Unlock()
	return mock.RetrieveEntityFunc(ctx, entityID, headers)
}

// RetrieveEntityCalls gets all the calls that were made to RetrieveEntity.
// Check the length with:
//     len(mockedContextBrokerClient.RetrieveEntityCalls())
func (mock *ContextBrokerClientMock) RetrieveEntityCalls() []struct {
	Ctx      context.Context
	EntityID string
	Headers  map[string][]string
} {
	var calls []struct {
		Ctx      context.Context
		EntityID string
		Headers  map[string][]string
	}
	mock.lockRetrieveEntity.RLock()
	calls = mock.calls.RetrieveEntity
	mock.lockRetrieveEntity.RUnlock()
	return calls
}

// UpdateEntityAttributes calls UpdateEntityAttributesFunc.
func (mock *ContextBrokerClientMock) UpdateEntityAttributes(ctx context.Context, entityID string, fragment types.EntityFragment, headers map[string][]string) (*ngsild.UpdateEntityAttributesResult, error) {
	if mock.UpdateEntityAttributesFunc == nil {
		panic("ContextBrokerClientMock.UpdateEntityAttributesFunc: method is nil but ContextBrokerClient.UpdateEntityAttributes was just called")
	}
	callInfo := struct {
		Ctx      context.Context
		EntityID string
		Fragment types.EntityFragment
		Headers  map[string][]string
	}{
		Ctx:      ctx,
		EntityID: entityID,
		Fragment: fragment,
		Headers:  headers,
	}
	mock.lockUpdateEntityAttributes.Lock()
	mock.calls.UpdateEntityAttributes = append(mock.calls.UpdateEntityAttributes, callInfo)
	mock.lockUpdateEntityAttributes.Unlock()
	return mock.UpdateEntityAttributesFunc(ctx, entityID, fragment, headers)
}

// UpdateEntityAttributesCalls gets all the calls that were made to UpdateEntityAttributes.
// Check the length with:
//     len(mockedContextBrokerClient.UpdateEntityAttributesCalls())
func (mock *ContextBrokerClientMock) UpdateEntityAttributesCalls() []struct {
	Ctx      context.Context
	EntityID string
	Fragment types.EntityFragment
	Headers  map[string][]string
} {
	var calls []struct {
		Ctx      context.Context
		EntityID string
		Fragment types.EntityFragment
		Headers  map[string][]string
	}
	mock.lockUpdateEntityAttributes.RLock()
	calls = mock.calls.UpdateEntityAttributes
	mock.lockUpdateEntityAttributes.RUnlock()
	return calls
}
