package ioencode

import (
  "encoding/hex"
  "io"
)

// HexReader writes a string of bytes as hexadecimal
type HexReader struct {
  r io.Reader
  b []byte
}

// NewHexReader creates a new HexReader
func NewHexReader(r io.Reader) *HexReader {
  return &HexReader{
    r: r,
    b: nil,
  }
}

func (r *HexReader) Read(p []byte) (n int, err error) {
  if len(p) == 0 {
    return 0, nil
  }

  needSize := hex.EncodedLen(len(p))

  //Grow if necessary
  if len(r.b) < needSize {
    r.b = make([]byte, needSize)
  }

  decodeBuffer := r.b[:needSize]
  n, errRead := r.r.Read(decodeBuffer)
  if n == 0 {
    return 0, errRead
  }

  decodeBuffer = decodeBuffer[:n]
  n, errDecode := hex.Decode(p, decodeBuffer)

  if errRead != nil {
    return n, errRead
  }

  return n, errDecode
}
