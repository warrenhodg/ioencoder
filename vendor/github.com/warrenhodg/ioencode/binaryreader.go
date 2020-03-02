package ioencode

import (
  "io"
)

// BinaryReader writes a string of bytes as Binaryadecimal
type BinaryReader struct {
  r io.Reader
  b []byte
  littleEndian bool
}

// NewBinaryReader creates a new BinaryReader
func NewBinaryReader(r io.Reader) *BinaryReader {
  return &BinaryReader{
    r: r,
    b: nil,
  }
}

func (r *BinaryReader) Read(p []byte) (n int, err error) {
  if len(p) == 0 {
    return 0, nil
  }

  needSize := len(p) * 8

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

  n = n / 8
  o := 0
  for i := 0; i < n; i++ {
    b := byte(0)
    for j := 0; j < 8; j++ {
      v := decodeBuffer[o] 

      switch v {
      case '0', '1':
        v -= '0'

      default:
        return i, ErrDecode(v)
      }

      if r.littleEndian {
        b += v << (7 - j)
      } else {
        b = b << 1 + v
      }

      o++
    }

    p[i] = b
  }

  if errRead != nil {
    return n, errRead
  }

  return n, nil
}
