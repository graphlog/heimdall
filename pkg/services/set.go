package services

import "github.com/google/wire"

var ServiceSet = wire.NewSet(NewMessageService, NewApplicationService, NewAMQPConnection, NewDBConnection)
