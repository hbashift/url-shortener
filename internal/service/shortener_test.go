package service

import (
	"github.com/hbashift/url-shortener/internal/domain/repository"
	shortener "github.com/hbashift/url-shortener/pb"
	"reflect"
	"testing"
)

func TestNewShortenerService(t *testing.T) {
	type args struct {
		db repository.Repository
	}
	tests := []struct {
		name string
		args args
		want *ShortenerService
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewShortenerService(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewShortenerService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestShortenerService_GetUrl(t *testing.T) {
	type fields struct {
		db repository.Repository
	}
	type args struct {
		shortUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *shortener.LongUrl
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ShortenerService{
				db: tt.fields.db,
			}
			got, err := s.GetUrl(tt.args.shortUrl)
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

func TestShortenerService_PostUrl(t *testing.T) {
	type fields struct {
		db repository.Repository
	}
	type args struct {
		longUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *shortener.ShortUrl
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &ShortenerService{
				db: tt.fields.db,
			}
			got, err := s.PostUrl(tt.args.longUrl)
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
