package utils

import (
	"encoding/json"
	"github.com/Rican7/conjson"
	"github.com/Rican7/conjson/transform"
	"io"
)

const (
	marshalPrefix = ""
	marshalIndent = "    "
)

func JsonEncoder(w io.Writer) conjson.Encoder {
	enc := json.NewEncoder(w)
	enc.SetIndent(marshalPrefix, marshalIndent)
	return conjson.NewEncoder(enc, transform.ConventionalKeys())
}

func JsonDecoder(r io.Reader) conjson.Decoder {
	return conjson.NewDecoder(
		json.NewDecoder(r),
		transform.ConventionalKeys(),
		transform.ValidIdentifierKeys(),
	)
}
