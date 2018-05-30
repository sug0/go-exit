package exit

import (
    "os"
    "io"
    "fmt"
)

// Represents the exit status of a program
type Status interface {
    ExitCode() int
}

type exitCode int

func (e exitCode) ExitCode() int {
    return int(e)
}

type msg struct {
    code int
    msg  string
    out  io.Writer
}

func (m msg) ExitCode() int {
    if m.out != nil {
        fmt.Fprintln(m.out, m.msg)
    } else {
        fmt.Fprintln(os.Stderr, m.msg)
    }
    return m.code
}

// This should be the first deferred function in
// your code; it will respect all your deffered
// calls after it
func Handler() {
    if e := recover(); e != nil {
        if s, ok := e.(Status); ok {
            os.Exit(s.ExitCode())
        }
        // not something that we can handle...
        // time to panic lol
        panic(e)
    }
}

// Stops the program execution with the exit code 's.ExitCode()',
// honoring all deferred calls
func WithStatus(s Status) {
    panic(s)
}

// Stops the program execution with the exit code 'code',
// honoring all deferred calls
func WithCode(code int) {
    panic(exitCode(code))
}

// Stops the program execution with the exit code 'code',
// printing a message to 'w', and honoring all deferred calls
func WithMsg(w io.Writer, code int, format string, a ...interface{}) {
    panic(msg{
        out: w,
        code: code,
        msg: fmt.Sprintf(format, a...),
    })
}
