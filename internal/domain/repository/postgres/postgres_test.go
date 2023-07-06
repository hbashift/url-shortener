package postgres

import (
	"gorm.io/gorm"
	"testing"
)

func Test_postgresDb_GetUrl(t *testing.T) {
	type fields struct {
		db *gorm.DB
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
			p := &postgresDb{
				db: tt.fields.db,
			}
			got, err := p.GetUrl(tt.args.shortUrl)
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

func Test_postgresDb_PostUrl(t *testing.T) {
	type fields struct {
		db *gorm.DB
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
			p := &postgresDb{
				db: tt.fields.db,
			}
			got, err := p.PostUrl(tt.args.longUrl)
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
