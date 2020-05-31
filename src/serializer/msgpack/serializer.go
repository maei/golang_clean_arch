package msgpack

import (
	"github.com/maei/golang_clean_arch/src/shortener"
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

type Serial struct{}

func (r *Serial) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Serial.Decode")
	}
	return redirect, nil
}

func (r *Serial) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Serial.Encode")
	}
	return rawMsg, nil
}
