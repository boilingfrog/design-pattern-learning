package 建造者模式

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			got, err := NewResourcePoolConfig(item.conf.name, IntP(item.conf.maxTotal), IntP(item.conf.maxIdle), IntP(item.conf.minIdle))
			require.Equalf(t, item.wantErr, err != nil, "NewResourcePoolConfig() error = %v, wantErr %v", err, item.wantErr)
			assert.Equal(t, item.want, got)
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
	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			got, err := item.builder.Build()
			require.Equalf(t, item.wantErr, err != nil, "Build() error = %v, wantErr %v", err, item.wantErr)
			assert.Equal(t, item.want, got)
		})
	}
}

func TestRateLimiter(t *testing.T) {
	type arg struct {
		name string
		opt  []Param
	}
	tests := []struct {
		name    string
		arg     *arg
		want    *ResourcePoolConfig
		wantErr bool
	}{
		{
			name: "name empty",
			arg: &arg{
				name: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "maxIdle < minIdle",
			arg: &arg{
				name: "test",
				opt: []Param{
					MaxTotal(0),
					MaxIdle(10),
					MinIdle(10),
				},
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "success",
			arg: &arg{
				name: "test",
				opt: []Param{
					MaxTotal(defaultMaxTotal),
					MaxIdle(defaultMaxIdle),
					MinIdle(defaultMinIdle),
				},
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
	for _, item := range tests {
		t.Run(item.name, func(t *testing.T) {
			got, err := NewResourcePoolConfigOption(item.arg.name, item.arg.opt...)
			require.Equalf(t, item.wantErr, err != nil, "Build() error = %v, wantErr %v", err, item.wantErr)
			assert.Equal(t, item.want, got)
		})
	}

}
