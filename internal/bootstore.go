package internal

import (
	"github.com/zapkub/bootstore-angular-go-grpc/pkg/apis"
)

type BookStoreEndpointHandler struct {
	apis.UnimplementedBookstoreServiceServer
}
