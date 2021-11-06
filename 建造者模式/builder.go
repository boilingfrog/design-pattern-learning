package 建造者模式

import (
	"errors"
	"fmt"
)

const (
	defaultMaxTotal = 10
	defaultMaxIdle  = 10
	defaultMinIdle  = 1
)

// ResourcePoolConfig ...
type ResourcePoolConfig struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

func NewResourcePoolConfig(name string, maxTotal, maxIdle, minIdle *int) (*ResourcePoolConfig, error) {
	rc := &ResourcePoolConfig{
		maxTotal: defaultMaxTotal,
		maxIdle:  defaultMaxIdle,
		minIdle:  defaultMinIdle,
	}
	if name == "" {
		return nil, errors.New("name is empty")
	}
	rc.name = name

	if maxTotal != nil {
		if *maxTotal <= 0 {
			return nil, errors.New("maxTotal should be positive")
		}
		rc.maxTotal = *maxTotal
	}

	if maxIdle != nil {
		if *maxIdle <= 0 {
			return nil, errors.New("maxIdle should not be negative")
		}
		rc.maxIdle = *maxIdle
	}

	if minIdle != nil {
		if *minIdle <= 0 {
			return nil, errors.New("minIdle should not be negative")
		}
		rc.minIdle = *minIdle
	}

	return rc, nil
}

// ---------------------------------------------------------------------------------------------------------------------

func NewResourcePoolConfigSet(name string) (*ResourcePoolConfig, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}

	return &ResourcePoolConfig{
		maxTotal: defaultMaxTotal,
		maxIdle:  defaultMaxIdle,
		minIdle:  defaultMinIdle,
		name:     name,
	}, nil
}

// SetMinIdle ...
func (rc *ResourcePoolConfig) SetMinIdle(minIdle int) error {
	if minIdle < 0 {
		return fmt.Errorf("min idle cannot < 0, input: %d", minIdle)
	}
	rc.minIdle = minIdle
	return nil
}

// SetMaxIdle ...
func (rc *ResourcePoolConfig) SetMaxIdle(maxIdle int) error {
	if maxIdle < 0 {
		return fmt.Errorf("max idle cannot < 0, input: %d", maxIdle)
	}
	rc.maxIdle = maxIdle
	return nil
}

// SetMaxTotal ...
func (rc *ResourcePoolConfig) SetMaxTotal(maxTotal int) error {
	if maxTotal <= 0 {
		return fmt.Errorf("max total cannot <= 0, input: %d", maxTotal)
	}
	rc.maxTotal = maxTotal
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

// ResourcePoolConfigBuilder ...
type ResourcePoolConfigBuilder struct {
	name     string
	maxTotal int
	maxIdle  int
	minIdle  int
}

// SetName ...
func (rb *ResourcePoolConfigBuilder) SetName(name string) error {
	if name == "" {
		return fmt.Errorf("name can not be empty")
	}
	rb.name = name
	return nil
}

// SetMinIdle ...
func (rb *ResourcePoolConfigBuilder) SetMinIdle(minIdle int) error {
	if minIdle < 0 {
		return fmt.Errorf("max total cannot < 0, input: %d", minIdle)
	}
	rb.minIdle = minIdle
	return nil
}

// SetMaxIdle ...
func (rb *ResourcePoolConfigBuilder) SetMaxIdle(maxIdle int) error {
	if maxIdle < 0 {
		return fmt.Errorf("max total cannot < 0, input: %d", maxIdle)
	}
	rb.maxIdle = maxIdle
	return nil
}

// SetMaxTotal ...
func (rb *ResourcePoolConfigBuilder) SetMaxTotal(maxTotal int) error {
	if maxTotal <= 0 {
		return fmt.Errorf("max total cannot <= 0, input: %d", maxTotal)
	}
	rb.maxTotal = maxTotal
	return nil
}

// Build ...
func (rb *ResourcePoolConfigBuilder) Build() (*ResourcePoolConfig, error) {
	if rb.name == "" {
		return nil, errors.New("name can not be empty")
	}

	// 设置默认值
	if rb.minIdle == 0 {
		rb.minIdle = defaultMinIdle
	}

	if rb.maxIdle == 0 {
		rb.maxIdle = defaultMaxIdle
	}

	if rb.maxTotal == 0 {
		rb.maxTotal = defaultMaxTotal
	}

	if rb.maxTotal < rb.maxIdle {
		return nil, fmt.Errorf("max total(%d) cannot < max idle(%d)", rb.maxTotal, rb.maxIdle)
	}

	if rb.minIdle > rb.maxIdle {
		return nil, fmt.Errorf("max idle(%d) cannot < min idle(%d)", rb.maxIdle, rb.minIdle)
	}

	return &ResourcePoolConfig{
		name:     rb.name,
		maxTotal: rb.maxTotal,
		maxIdle:  rb.maxIdle,
		minIdle:  rb.minIdle,
	}, nil
}

// ---------------------------------------------------------------------------------------------------------------------

type Param func(*ResourcePoolConfig)

func NewResourcePoolConfigOption(name string, param ...Param) (*ResourcePoolConfig, error) {
	if name == "" {
		return nil, errors.New("name is empty")
	}
	ps := &ResourcePoolConfig{
		maxIdle:  defaultMinIdle,
		minIdle:  defaultMinIdle,
		maxTotal: defaultMaxTotal,
		name:     name,
	}

	for _, p := range param {
		p(ps)
		fmt.Println(ps)
	}

	if ps.maxTotal < 0 || ps.maxIdle < 0 || ps.minIdle < 0 {
		return nil, fmt.Errorf("args err, option: %v", ps)
	}

	if ps.maxTotal < ps.maxIdle || ps.minIdle > ps.maxIdle {
		return nil, fmt.Errorf("args err, option: %v", ps)
	}

	return ps, nil
}

func MaxTotal(maxTotal int) Param {
	return func(o *ResourcePoolConfig) {
		o.maxTotal = maxTotal
	}
}

func MaxIdle(maxIdle int) Param {
	return func(o *ResourcePoolConfig) {
		o.maxIdle = maxIdle
	}
}

func MinIdle(minIdle int) Param {
	return func(o *ResourcePoolConfig) {
		o.minIdle = minIdle
	}
}
