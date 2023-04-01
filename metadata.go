package connectgateway

import (
	"net/http"

	"google.golang.org/grpc/metadata"
)

func metadataFromHeader(in http.Header) (out metadata.MD) {
	for k, v := range in {
		out.Set(k, v...)
	}
	return
}
