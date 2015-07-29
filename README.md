# trash [![GoDoc](https://godoc.org/github.com/go-zoo/trash?status.svg)](https://godoc.org/github.com/go-zoo/trash) [![Build Status](https://travis-ci.org/go-zoo/trash.svg?branch=master)](https://travis-ci.org/go-zoo/trash)

### What is trash ?
trash deal with logging, sending and stacking errors.
Currently supporting json and xml format.
And also implementing the standard `error` interface.

![alt tag](http://lifescapesdesignblog.weebly.com/uploads/1/2/8/0/12806369/__6190365.jpg)

### Example
Creating a new err.

``` go
  import "log"
  import "errors"

  func main() {
    // Takes a logger and data format (json, xml)
    t := trash.New(logger, "json")

    if 1 != 2 {
      // Default Err
      // Arguments (Error Type, Error Message) -> (string, string or error)
      t.NewErr(trash.GenericErr, errors.New("example err")).Send(rw).Log()

      // HTTP Err
      t.NewErr(trash.GenericErr, "useless error").SendHTTP(rw, 404).LogHTTP(req)

      // Standalone inline HTTP error declaring
      trash.NewJSONErr(trash.InvalidDataErr, "1 not equal 2 ...").LogHTTP(req).SendHTTP(rw, 406)
    }
  }
```
You can also create a dump. Which is an error stack you want to backup.
A dump need a `io.Writer` to save the errors stack.

``` go
  func main() {
    t := trash.New(logger, "json")
    t2 := trash.New(logger, "xml")
    // The dump will write all errors from the provided trash in a log file.
    trash.NewDump(file, t, t2)

    if 1 != 2 {
      t.NewErr(trash.InvalidDataErr, "1 not equal 2")
    }
}

```

### TODO

- DOC
- More Testing
- Debugging
- Optimisation

### Contributing

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Write Tests!
4. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request

### License
MIT
