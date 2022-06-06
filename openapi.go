package libopenapi

import (
	"github.com/edsonmichaque/libopenapi/types"
	"github.com/edsonmichaque/libopenapi/validator"
)

func New(dec Decoder, options ...Option) (*types.Spec, error) {
	validators := []Option{
		WithValidator(validator.New()),
	}

	return newBuilder(dec, append(options, validators...)...)
}

func newBuilder(dec Decoder, options ...Option) (*types.Spec, error) {
	spec, err := dec.Decode()
	if err != nil {
		return nil, err
	}

	doc := doc{
		s: spec,
	}

	for _, option := range options {
		option.Apply(&doc)
	}

	return spec, nil
}

type doc struct {
	s *types.Spec
	v []Validator
}
