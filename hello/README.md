# [hello](https://blog.golang.org/using-go-modules)

A module is a collection of go packages with `go.mod` file at it's root.

`go.mod` file:

1. Defines the "main path" for a module. i.e. inside that path it is going to be the incharge.
2. For the dependency management for all packages in the root directory.
3. (Not in my most of the cases)If go packages are present outside `$GOPATH/src` path, go commands considers `go.mod` file in the current or any parent directory.
4. If go packages are present inside `$GOPATH/src` path, it always creates "main path" from `$GOPATH/src`.

`go.mod` dependencies:

1. go command resolves dependencies by looking into the `go.mod` file.
2. If an import in a go package file not present in `go.mod` file, then go commands downloads in and add it to the `go.mod` file with latest version.
    Note: latest version is either release, prerelease or if prerelease is not present then latest (There is no priority order when go command automatically downloads). 
3. Only direct dependencies are present in the `go.mod` file. For indirect dependencies that are linked with direct dependencies, they are loaded in the `go.sum` file (More on `go.sum` later).
4. Once `go.mod` has installed the dependencies, their cache is created locally at `$GOPATH/pkg/mod` and it is used further instead of reinstalling every time.

Since adding direct dependencies also brings indirect dependencies. To view all the dependencies:

```zsh
go list -m all
```

`go.sum` file:
1. Stores all the dependencies(direct+indirect) with their cryptographic hashes.
2. go command uses the `go.sum` file to retrieve the same bits as that of the first download. This is needed, so module's versions are not changed unexpecetly.
3. Both `go.mod` and `go.sum` should go into the version control.

If there are any untagged version created automatically by the go command while running some package (Always run `go list -m all` to check).
Then, upgrade it to tagged version using `go get <module>`.

To list the versions for a module, use `go list -m -versions <module>`.

To install a specific version, `go get <module>@<version>`.

If a module has two or more different packages in different versions.
Then both can be imported by using the complete path of the import and naming the import.

For eg.:

```go
import (
    "rsc.io/quote"
    quoteV3 "rsc.io/quote/v3"
)
```

Quick look into the fuctions of a package:

```zsh
go doc rsc.io/quote/v3
```

To remove unused dependencies:
```
go mod tidy
```
