# trash

### What is trash ?
trash deal with logging, sending and stacking errors.
Currently supporting json and xml format.

![alt tag](http://lifescapesdesignblog.weebly.com/uploads/1/2/8/0/12806369/__6190365.jpg)

### Example
Creating a new err.

``` go
  import "log"

  func main() {
    // Takes a logger and data format (json, xml)
    t := trash.New(logger, "json")

    if 1 != 2 {
      // Default Err
      t.NewErr(trash.GenericErr, "example err").Send(rw).Log()

      // HTTP Err
      t.NewErr(trash.GenericErr, "useless error").SendHTTP(rw, 404).LogHTTP(req)
      
      // Standalone inline HTTP error declaring
      trash.NewJSONErr(trash.InvalidDataErr, "1 not equal 2 ...").LogHTTP(req).SendHTTP(rw, 406)
    }
  }
```
You can also create a dump. Which is an error stack you want to process in the same way.
A dump need a `io.Writer` to save the errors stack.

``` go
  func main() {
    d := trash.NewDump(os.Stdout, "json")
    if 1 != 2 {
      d.NewErr(trash.InvalidDataErr, "1 not equal 2")
    }
    // Write the dump to the io.Writer
    d.Dump()
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
