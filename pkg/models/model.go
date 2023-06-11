package models

type ModelInterface interface {
	ctx *context.Context
	func Context() *context.Context
}
