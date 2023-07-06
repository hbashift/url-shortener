package encoder

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDecryptUrl(t *testing.T) {
	type args struct {
		shortUrl string
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "test_1",
			args: args{shortUrl: "CdFksa_sdf"},
			want: 438614015246978198,
		},
		{
			name: "test_2",
			args: args{shortUrl: "aaaaaaaaaa"},
			want: 0,
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, DecryptUrl(tt.args.shortUrl),
			fmt.Sprintf("DecryptUrl() = %v, want %v", DecryptUrl(tt.args.shortUrl), tt.want))
	}
}

func TestEncodeUrl(t *testing.T) {
	type args struct {
		id uint64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test_1",
			args: args{id: 1},
			want: "aaaaaaaaab",
		},
		{
			name: "test_2",
			args: args{id: 100},
			want: "aaaaaaaab" + string([]byte(alphabet)[37]),
		},
		{
			name: "test_3",
			args: args{id: 0},
			want: "aaaaaaaaaa",
		},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, EncodeUrl(tt.args.id),
			fmt.Sprintf("EncodeUrl() = %v, want %v", EncodeUrl(tt.args.id), tt.want))
	}
}
