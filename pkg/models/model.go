package models

import "context"

type ModelInterface interface {
	GetPath() string
	Ctx() *context.Context
}
