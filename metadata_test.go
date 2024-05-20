package connectgateway

import (
	"maps"
	"net/http"
	"slices"
	"testing"

	"google.golang.org/grpc/metadata"
)

func TestMetadataFromHeader(t *testing.T) {
	for _, params := range []struct {
		name string
		in   http.Header
		want metadata.MD
	}{
		{
			name: "empty header",
			in:   makeEmptyHeader(),
			want: makeEmptyMetadata(),
		},
		{
			name: "non-empty header",
			in:   makeNonEmptyHeader(),
			want: makeNonEmptyMetadata(),
		},
		{
			name: "non-empty header with duplicates",
			in:   makeNonEmptyHeaderWithDuplicates(),
			want: makeNonEmptyMetadataWithDuplicates(),
		},
	} {
		t.Run(params.name, func(t *testing.T) {
			got := metadataFromHeader(params.in)
			if !maps.EqualFunc(got, params.want, slices.Equal) {
				t.Errorf("unexpected metadata: got %v, want %v", got, params.want)
			}
		})
	}
}

func makeEmptyHeader() http.Header {
	return make(http.Header)
}

func makeNonEmptyHeader() http.Header {
	h := makeEmptyHeader()
	h.Set("ETag", "123-a")
	return h
}

func makeNonEmptyHeaderWithDuplicates() http.Header {
	h := makeNonEmptyHeader()
	h.Add("ETag", "123-b")
	return h
}

func makeEmptyMetadata() metadata.MD {
	return metadata.New(nil)
}

func makeNonEmptyMetadata() metadata.MD {
	md := makeEmptyMetadata()
	md.Set("ETag", "123-a")
	return md
}

func makeNonEmptyMetadataWithDuplicates() metadata.MD {
	md := makeNonEmptyMetadata()
	md.Append("ETag", "123-b")
	return md
}
