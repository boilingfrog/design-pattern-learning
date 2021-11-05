package 建造者模式

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tj/assert"
)

// IntP return pointer
func IntP(v int) *int { return &v }

func TestResourcePoolConfig(t *testing.T) {
	tests := []struct {
		name    string
		conf    *ResourcePoolConfig
		want    *ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "name empty",
			conf: &ResourcePoolConfig{
				name:     "",
				maxTotal: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test normal",
			conf: &ResourcePoolConfig{
				name:     "小明",
				maxTotal: 10,
				maxIdle:  10,
				minIdle:  2,
			},
			want: &ResourcePoolConfig{
				name:     "小明",
				maxTotal: 10,
				maxIdle:  10,
				minIdle:  2,
			},
			wantErr: false,
		},
		{
			name: "test unNormal",
			conf: &ResourcePoolConfig{
				name:    "小明",
				maxIdle: 10,
				minIdle: 2,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewResourcePoolConfig(tt.conf.name, IntP(tt.conf.maxTotal), IntP(tt.conf.maxIdle), IntP(tt.conf.minIdle))
			require.Equalf(t, tt.wantErr, err != nil, "NewResourcePoolConfig() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestResourcePoolConfigSet(t *testing.T) {
	rc, err := NewResourcePoolConfigSet("小红")
	if err != nil {
		t.Fatal(err)
	}
	if err := rc.SetMaxIdle(12); err != nil {
		t.Fatal(err)
	}
	if err := rc.SetMinIdle(2); err != nil {
		t.Fatal(err)
	}

	t.Log(rc)
}

func TestResourcePoolConfigBuilder(t *testing.T) {
	tests := []struct {
		name    string
		builder *ResourcePoolConfigBuilder
		want    *ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "name empty",
			builder: &ResourcePoolConfigBuilder{
				name:     "",
				maxTotal: 0,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "maxIdle < minIdle",
			builder: &ResourcePoolConfigBuilder{
				name:     "test",
				maxTotal: 0,
				maxIdle:  10,
				minIdle:  20,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			builder: &ResourcePoolConfigBuilder{
				name: "test",
			},
			want: &ResourcePoolConfig{
				name:     "test",
				maxTotal: defaultMaxTotal,
				maxIdle:  defaultMaxIdle,
				minIdle:  defaultMinIdle,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.builder.Build()
			require.Equalf(t, tt.wantErr, err != nil, "Build() error = %v, wantErr %v", err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}
