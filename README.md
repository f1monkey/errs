# Errs

A simple wrapper over the standard `errors` package to get error stack traces.

## Install

```
$ go get -v github.com/f1monkey/errs
```

## Usage

```go
package main

import (
    "errors"

    "github.com/f1monkey/errs"
)


func doSomething() error {
    // ...

    result, err := service.Do(ctx)
    if err != nil {
        return errs.Errorf("service err: %w", err) // stacktrace will be captured here
    }

    // ...
}

type stacktracer interface {
	StackTrace() []byte
}

func DoSomething() error {
    // ...

    if err := doSomething(); err != nil {
		var stacktraceErr *errs.Error
		// get stack trace, option 1: use errors.As()
		if errors.As(err, &stacktraceErr) {
			fmt.Println(stacktraceErr.StackTrace())
		}

		// get stack trace, option 2: use type assertion
		if e, ok := err.(stacktracer); ok {
			fmt.Println(e.StackTrace())
		}
	}

	// ...
}

```