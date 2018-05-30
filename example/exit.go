package main

import (
    "os"
    "fmt"
    "github.com/sugoiuguu/go-exit"
)

type ExitStatus int

func (s ExitStatus) ExitCode() int {
    fmt.Fprintln(os.Stderr, "OMG IT ALL WENT DOWN THE SHITTER")
    return int(s)
}

func main() {
    defer exit.Handler()
    defer fmt.Println("it worked I guess")
    exit.WithStatus(ExitStatus(2))
}
