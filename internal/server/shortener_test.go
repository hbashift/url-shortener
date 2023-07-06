package server

import (
	"context"
	"github.com/hbashift/url-shortener/internal/service"
	pb "github.com/hbashift/url-shortener/pb"
	"reflect"
	"testing"
)

func TestNewShortenerServer(t *testing.T) {
	type args struct {
		s *service.ShortenerService
	}
	tests := []struct {
		name string
		args args
		want *ShortenerServer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewShortenerServer(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShortenerServer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortenerServer_GetUrl(t *testing.T) {
	type fields struct {
		shortener                    *service.ShortenerService
		UnimplementedShortenerServer pb.UnimplementedShortenerServer
	}
	type args struct {
		ctx context.Context
		url *pb.ShortUrl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.LongUrl
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ShortenerServer{
				shortener:                    tt.fields.shortener,
				UnimplementedShortenerServer: tt.fields.UnimplementedShortenerServer,
			}
			got, err := s.GetUrl(tt.args.ctx, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortenerServer_PostUrl(t *testing.T) {
	type fields struct {
		shortener                    *service.ShortenerService
		UnimplementedShortenerServer pb.UnimplementedShortenerServer
	}
	type args struct {
		ctx context.Context
		url *pb.LongUrl
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.ShortUrl
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ShortenerServer{
				shortener:                    tt.fields.shortener,
				UnimplementedShortenerServer: tt.fields.UnimplementedShortenerServer,
			}
			got, err := s.PostUrl(tt.args.ctx, tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}
