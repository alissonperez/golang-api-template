package encoding

import (
	"encoding/json"
	"io"
)

func JsonDecode(r io.Reader, v interface{}) error {
	decoder := json.NewDecoder(r)

	if err := decoder.Decode(&v); err != nil {
		return err
	}

	return nil
}

func JsonEncode(w io.Writer, v interface{}) {
	json.NewEncoder(w).Encode(v)
}
