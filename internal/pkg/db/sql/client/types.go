package client

import "context"

type NamedQueryExecutionWithContextProps struct {
	Ctx   context.Context
	CorId string
	Query string
	Args  interface{}
}

type GetQueryWithContextProps struct {
	Ctx   context.Context
	CorId string
	Query string
	Dest  interface{}
	Args  interface{}
}

type UpdateQueryExecutionTxWithContextProps struct {
	Ctx   context.Context
	CorId string
	Query string
	Args  interface{}
}

type PrepareNamedTxWithContextProps struct {
	Ctx   context.Context
	CorId string
	Query string
	Dest  interface{}
	Args  interface{}
}
