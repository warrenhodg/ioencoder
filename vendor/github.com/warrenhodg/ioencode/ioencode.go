package ioencode

import (
  "fmt"
)

// ErrDecode is returned when decoding receives a bad input
type ErrDecode byte

func (e ErrDecode) Error() string {
  return fmt.Sprintf("error decoding byte %d", byte(e))
}
