package client

import (
	"context"
	pb "github.com/hbashift/url-shortener/pb"
	"google.golang.org/grpc"
	"reflect"
	"testing"
)

func TestRunClient(t *testing.T) {
	type args struct {
		URL string
	}
	tests := []struct {
		name string
		args args
		want pb.ShortenerClient
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RunClient(tt.args.URL); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortener_GetUrl(t *testing.T) {
	type fields struct {
		client pb.ShortenerClient
	}
	type args struct {
		ctx      context.Context
		shortUrl *pb.ShortUrl
		opts     []grpc.CallOption
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
			s := &Shortener{
				client: tt.fields.client,
			}
			got, err := s.GetUrl(tt.args.ctx, tt.args.shortUrl, tt.args.opts...)
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

func TestShortener_PostUrl(t *testing.T) {
	type fields struct {
		client pb.ShortenerClient
	}
	type args struct {
		ctx     context.Context
		longUrl *pb.LongUrl
		opts    []grpc.CallOption
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
			s := &Shortener{
				client: tt.fields.client,
			}
			got, err := s.PostUrl(tt.args.ctx, tt.args.longUrl, tt.args.opts...)
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
