package contract

// EntityInterface is a set of functions that should be implemented by entity
type EntityInterface interface {
	TableName() string
	FilterableFields() []interface{}
	TimeFields() []interface{}
}
