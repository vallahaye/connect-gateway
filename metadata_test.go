package connectgateway

import (
	"net/http"
	"testing"

	"google.golang.org/grpc/metadata"
	"gotest.tools/v3/assert"
)

func TestMetadataFromHeader(t *testing.T) {
	for _, params := range []struct {
		name string
		in   http.Header
		want metadata.MD
	}{
		{
			"empty header",
			make(http.Header),
			metadata.New(nil),
		},
		{
			"non-empty header",
			nonEmptyHeader(),
			nonEmptyMetadata(),
		},
		{
			"non-empty header with duplicates",
			nonEmptyHeaderWithDuplicates(),
			nonEmptyMetadataWithDuplicates(),
		},
	} {
		t.Run(params.name, func(t *testing.T) {
			got := metadataFromHeader(params.in)
			assert.DeepEqual(t, params.want, got)
		})
	}
}

func nonEmptyHeader() http.Header {
	h := make(http.Header)
	h.Set("ETag", "123-a")
	return h
}

func nonEmptyHeaderWithDuplicates() http.Header {
	h := nonEmptyHeader()
	h.Add("ETag", "123-b")
	return h
}

func nonEmptyMetadata() metadata.MD {
	return metadata.New(map[string]string{"ETag": "123-a"})
}

func nonEmptyMetadataWithDuplicates() metadata.MD {
	md := nonEmptyMetadata()
	md.Append("ETag", "123-b")
	return md
}
