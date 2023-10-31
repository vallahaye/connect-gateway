package connectgateway

import (
	"net/http"

	"google.golang.org/grpc/metadata"
)

func metadataFromHeader(in http.Header) metadata.MD {
	out := metadata.New(nil)
	for k, vals := range in {
		out.Set(k, vals...)
	}
	return out
}
