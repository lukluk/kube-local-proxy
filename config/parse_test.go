package config

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestExplode(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want Konfig
	}{
		{"TestShouldSuccesBuildKonfig",
			args{
				"vhost=context/service:8080:com.gopay.withdrawal",
			},
			Konfig{
				"context",
				"vhost",
				"service",
				8080,
				"com.gopay.withdrawal",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Explode(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Explode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetConfig(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []Konfig
	}{
		{
			"TestShouldGetConfig",
			args{
				"/tmp/test.cfg",
			},
			[]Konfig{
				{
					"context",
					"vhost",
					"serviceName",
					8080,
					"",
				},
				{
					"context",
					"vhost2",
					"fooService",
					8080,
					"com.gopay.withdrawal",
				},
			},
		},
	}
	d1 := []byte("vhost=context/serviceName:8080\nvhost2=context/fooService:8080:com.gopay.withdrawal")
	ioutil.WriteFile("/tmp/test.cfg", d1, 0644)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetConfig(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
