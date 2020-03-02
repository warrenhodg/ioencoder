package main

import (
  "os"
  "io"
  "fmt"

  "github.com/warrenhodg/ioencode"
	"gopkg.in/alecthomas/kingpin.v2"
)

type config struct {
  inputEncoding string
  outputEncoding string

  inputFile string
  outputFile string
}

func getConfig() (*config, error) {
  var err error

  cfg := &config{}

  app := kingpin.New("ioencoder", "Too to encode and decode io streams")

  app.Flag("input-file", "File from which to input (- for stdin)").Short('i').Default("").StringVar(&cfg.inputFile)

  app.Flag("output-file", "File to which to output (- for stdout)").Short('o').Default("").StringVar(&cfg.outputFile)

  app.Flag("input-encoding", "Type of input encoding (raw=r, hex=x)").Short('I').Default("raw").StringVar(&cfg.inputEncoding)

  app.Flag("output-encoding", "Type of output encoding (raw=r, hex=x, binary=b)").Short('O').Default("raw").StringVar(&cfg.outputEncoding)

  _, err = app.Parse(os.Args[1:])

  if cfg.inputFile == "-" {
    cfg.inputFile = ""
  }

  if cfg.outputFile == "-" {
    cfg.outputFile = ""
  }

  return cfg, err
}

func run() error {
  var (
    cfg *config
    err error
    input io.Reader
    output io.Writer
  )

  cfg, err = getConfig()
  if err != nil {
    return err
  }

  input = os.Stdin
  if cfg.inputFile != "" {
    inputFile, err := os.Open(cfg.inputFile)
    if err != nil {
      return err
    }

    input = inputFile
    defer inputFile.Close()
  }

  output = os.Stdout
  if cfg.outputFile != "" {
    outputFile, err := os.OpenFile(cfg.outputFile, os.O_CREATE & os.O_RDWR, 0755)
    if err != nil {
      return err
    }

    output = outputFile
    defer outputFile.Close()
  }

  switch cfg.inputEncoding {
    case "hex", "x":
      input = ioencode.NewHexReader(input)
    case "binary", "b":
      input = ioencode.NewBinaryReader(input)
    case "raw", "r":
    default:
      return fmt.Errorf("unknown input encoding")
  }

  switch cfg.outputEncoding {
    case "hex", "x":
      output = ioencode.NewHexWriter(output)
    case "binary", "b":
      output = ioencode.NewBinaryWriter(output)
    case "raw", "r":
    default:
      return fmt.Errorf("unknown output encoding")
  }

  _, err = io.Copy(output, input)
  return err
}

func main() {
  if err := run(); err != nil {
    fmt.Fprintf(os.Stderr, "%v", err)
    os.Exit(1)
  }
}
