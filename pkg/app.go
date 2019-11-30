// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package root

import (
	"github.com/google/wire"
	"github.com/graphlog/heimdall/pkg/handlers"
	"github.com/graphlog/heimdall/pkg/server"
	"github.com/graphlog/heimdall/pkg/services"
)

// AppSet -
var AppSet = wire.NewSet(services.ServiceSet, handlers.RouterSet, server.ServerSet)

// InitializeApp -
func InitializeApp() (*server.AppServer, error) {
	wire.Build(AppSet)
	return nil, nil
}
