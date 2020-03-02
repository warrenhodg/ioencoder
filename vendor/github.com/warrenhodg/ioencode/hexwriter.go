package ioencode

import (
  "encoding/hex"
  "io"
)

// HexWriter writes a string of bytes as hexadecimal
type HexWriter struct {
  w io.Writer
  b []byte
}

// NewHexWriter creates a new HexWriter
func NewHexWriter(w io.Writer) *HexWriter {
  return &HexWriter{
    w: w,
    b: nil,
  }
}

func (w *HexWriter) Write(p []byte) (n int, err error) {
  if len(p) == 0 {
    return 0, nil
  }

  needSize := hex.EncodedLen(len(p))

  //Grow if necessary
  if cap(w.b) < needSize {
    w.b = make([]byte, needSize)
  }

  _ = hex.Encode(w.b, p)

  n, err = w.w.Write(w.b[:needSize])

  return n/2, err
}
