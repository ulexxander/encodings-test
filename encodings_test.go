package encodings

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"testing"

	"github.com/gbrlsnchs/jwt/v3"
)

var key = []byte("this is my key")

type payload struct {
	UID int64
}

func SignVerifyJWTGbrlsnchs(alg jwt.Algorithm) error {
	token, err := jwt.Sign(payload{UID: 10}, alg)
	if err != nil {
		return fmt.Errorf("failed to sign: %s", err)
	}

	// fmt.Printf("jwt len=%d string=%s", len(token), string(token))

	var tokenPayload payload
	_, err = jwt.Verify(token, alg, &tokenPayload)
	if err != nil {
		return fmt.Errorf("failed to verify: %s", err)
	}

	return nil
}

func EncodeDecodeGob() error {
	var buf bytes.Buffer
	if err := gob.NewEncoder(&buf).Encode(payload{
		UID: 10,
	}); err != nil {
		return fmt.Errorf("failed to encode: %s", err)
	}

	// fmt.Printf("gob len=%d string=%s base64=%s\n", buf.Len(), buf.String(), base64.StdEncoding.EncodeToString(buf.Bytes()))

	var tokenPayload payload
	if err := gob.NewDecoder(&buf).Decode(&tokenPayload); err != nil {
		return fmt.Errorf("failed to decode: %s", err)
	}

	return nil
}

func TestJWTGbrlsnchs(t *testing.T) {
	alg := jwt.NewHS256(key)
	if err := SignVerifyJWTGbrlsnchs(alg); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func TestGob(t *testing.T) {
	if err := EncodeDecodeGob(); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
}

func BenchmarkJWTGbrlsnchs(b *testing.B) {
	alg := jwt.NewHS256(key)
	for i := 0; i < b.N; i++ {
		if err := SignVerifyJWTGbrlsnchs(alg); err != nil {
			b.Fatalf("unexpected error: %s", err)
		}
	}
}

func BenchmarkGob(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if err := EncodeDecodeGob(); err != nil {
			b.Fatalf("unexpected error: %s", err)
		}
	}
}
