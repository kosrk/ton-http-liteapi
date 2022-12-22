package utils

import (
	"bytes"
	"fmt"
	"github.com/startfellows/tongo/liteclient"
	"github.com/startfellows/tongo/tl"
	"reflect"
	"testing"
)

func TestCamelCase(t *testing.T) {
	block := liteclient.TonNodeBlockIdExt{
		Workchain: -1,
		Shard:     -90000000,
		Seqno:     123,
		RootHash:  tl.Int256{1, 2, 3, 4, 5},
		FileHash:  tl.Int256{5, 4, 3, 2, 1},
	}
	x := liteclient.LiteServerGetBlockRequest{
		Id: block,
	}

	buf := new(bytes.Buffer)

	err := JsonEncoder(buf).Encode(x)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%s\n", string(buf.Bytes()))

	var x1 liteclient.LiteServerGetBlockRequest
	err = JsonDecoder(buf).Decode(&x1)
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(x, x1) {
		t.Fatal("not equal")
	}
}
