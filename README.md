# trash

### What is trash ?
trash deal with logging, sending and stacking errors.

![alt tag](http://www.elementaryos-fr.org/wp-content/uploads/2013/12/FreeGreatPicture.com-23409-trash.jpg)

### Example
Creating a new err.

``` go
  func main() {
    if 1 != 2 {
      // inline error declaring
      trash.NewErr(trash.INVALID_DATA_ERR, "1 not equal 2 ...", "json").Log().Send(rw)
    }
  }
```
You can also create a dump. Which is an error stack you want to process in the same way.
A dump need a `io.Writer` to save the errors stack. 

``` go 
  func main() {
    d := trash.NewDump(os.Stdout, "json")
    if 1 != 2 {
      d.NewErr(trash.INVALID_DATA_ERR, "1 not equal 2")
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
