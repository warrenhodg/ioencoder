package ioencode

import (
  "io"
)

// BinaryWriter writes a string of bytes as Binaryadecimal
type BinaryWriter struct {
  w io.Writer
  b []byte
  littleEndian bool
}

// NewBinaryWriter creates a new BinaryWriter
func NewBinaryWriter(w io.Writer) *BinaryWriter {
  return &BinaryWriter{
    w: w,
    b: nil,
  }
}

func (w *BinaryWriter) Write(p []byte) (n int, err error) {
  var v byte

  if len(p) == 0 {
    return 0, nil
  }

  needSize := len(p) * 8

  //Grow if necessary
  if len(w.b) < needSize {
    w.b = make([]byte, needSize)
  }

  o := 0
  for i := 0; i < len(p); i++ {
    b := p[i]

    for j := 0; j < 8; j++ {
      if w.littleEndian {
        v = b & 1 + '0'
        b = b >> 1
      } else {
        v = b >> 7 + '0'
        b = b << 1
      }
      w.b[o] = v

      o++
    }
  }

  n, err = w.w.Write(w.b[:needSize])

  return n/8, err
}
