//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializedService(isError bool) (*SimpleService, error) {
	wire.Build(NewSimpleRepository, NewSimpleService)
	return nil, nil
}

func InitializedDatabaseRepository() *DatabaseRepository {
	wire.Build(NewDatabaseMongoDB, NewDatabasePostgresql, NewDatabaseRepository)
	return nil
}

var fooSet = wire.NewSet(NewFooRepository, NewFooService)

var barSet = wire.NewSet(NewBarRepository, NewBarService)

func InitializedFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, NewFooBarService)
	return nil
}

// ini salah
// func InitializedHelloService() *HelloService {
// 	wire.Build(NewSayHelloImpl, NewHelloService)
// 	return nil
// }

var HelloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)),
)

//ini yang benar
func InitializedHelloService() *HelloService {
	wire.Build(HelloSet, NewHelloService)
	return nil
}
