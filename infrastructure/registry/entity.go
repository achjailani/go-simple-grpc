package registry

import (
	"github/achjailani/go-simple-grpc/domain/contract"
	"github/achjailani/go-simple-grpc/domain/entity"
)

// CollectEntities is function collects entities
func CollectEntities() []contract.Entity {
	return []contract.Entity{
		{Entity: entity.User{}},
		{Entity: entity.AuthToken{}},
		{Entity: entity.HttpLog{}},
	}
}

// CollectTables is function collects entity names
func CollectTables() []contract.Table {
	var user entity.User
	var authToken entity.AuthToken
	var httpLog entity.HttpLog

	return []contract.Table{
		{Name: user.TableName()},
		{Name: authToken.TableName()},
		{Name: httpLog.TableName()},
	}
}

// NewEntityRegistry is constructor of Registry
func NewEntityRegistry() *entity.Registry {
	var entityRegistry []contract.Entity
	var tableRegistry []contract.Table

	entityRegistry = append(entityRegistry, CollectEntities()...)
	tableRegistry = append(tableRegistry, CollectTables()...)

	return &entity.Registry{
		Entities: entityRegistry,
		Tables:   tableRegistry,
	}
}
