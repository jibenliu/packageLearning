package gen

import (
	"bytes"
	"fmt"
	"io"

	"mvdan.cc/gofumpt/format"
)

func Format(in io.Reader, out io.Writer) error {
	b := &bytes.Buffer{}
	_, err := io.Copy(b, in)
	if err != nil {
		return err
	}

	res, err := format.Source(b.Bytes(), format.Options{
		LangVersion: "1.16",
		ExtraRules:  true,
	})
	if err != nil {
		return err
	}

	fmt.Fprint(out, string(res))

	return err
}
