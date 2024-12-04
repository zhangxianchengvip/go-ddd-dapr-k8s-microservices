package infrastructure

import "go.uber.org/fx"

func DependencyInjection() []fx.Option {

	return []fx.Option{
		fx.Provide(NewRespository),
	}

}
