package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
)

func Test_redisDb_GetUrl(t *testing.T) {
	type fields struct {
		ctx      context.Context
		mainDB   *redis.Client
		uniqueDB *redis.Client
		id       uint64
	}
	type args struct {
		shortUrl uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisDb{
				ctx:      tt.fields.ctx,
				mainDB:   tt.fields.mainDB,
				uniqueDB: tt.fields.uniqueDB,
				id:       tt.fields.id,
			}
			got, err := r.GetUrl(tt.args.shortUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_redisDb_PostUrl(t *testing.T) {
	type fields struct {
		ctx      context.Context
		mainDB   *redis.Client
		uniqueDB *redis.Client
		id       uint64
	}
	type args struct {
		longUrl string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &redisDb{
				ctx:      tt.fields.ctx,
				mainDB:   tt.fields.mainDB,
				uniqueDB: tt.fields.uniqueDB,
				id:       tt.fields.id,
			}
			got, err := r.PostUrl(tt.args.longUrl)
			if (err != nil) != tt.wantErr {
				t.Errorf("PostUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PostUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}
