# [gowiki](https://golang.org/doc/articles/wiki/)

## Go methods

Since go doesn't have class like other languages. It uses methods to access specific functions.

Example:

```go
// `func_name` is a method with reciver of type Struct_name
func (object Struct_name) func_name() error {
    // Use `object` of the structure Struct_name as call by value.
}

// `func_name` is a method with reciver of type Struct_name pointer
func (object *Struct_name) func_name() error {
    // Use `object` of the structure Struct_name as call by reference.
}
```

Note:
    - First function will always be called as call by value irrective of whether it is called by an object or object pointer.
    - Second function will always be called as call by reference irrective of whether it is called by an object or object pointer.

## Go extras

- Function name is used as camel case