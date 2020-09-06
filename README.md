## Go-Helpers package

The package includes functions, types and tools which may help to develop GoLang based software.

#### Current structure
 * `helpers.go` - a list of functions and tools like `GenerateRandomString()`, `StringInSlice()` etc;
 * `struct_mathcer.go` - an implementation of _GoMock_ (see more [here](https://godoc.org/github.com/golang/mock/gomock)) matcher which provide you structs matching with difficult struct fields values e.g. `time.Now()`, randomized strings, etc. It may used for unit testing when you need to check certian mocked function call, eg:
 
```
timeMatcher := helpers.TimeMatcher{time.Now()}

createAppCall := repoMock.EXPECT().CreateApplication(helpers.StructMatcher{
    Matching:      matchedValue,
    MatcherFields: map[string]gomock.Matcher{"CreatedAt": timeMatcher, "UpdatedAt": timeMatcher},
    SkipFields:    []string{"tableName", "Key", "PasswordDigest", "Secret"},
}).Do(func(app *structs.Application) {
    //...
}) 		
```

#### Testing

The package has unit test coverage, to run tests just call a following command:
```sh
go test -v --race ./...
```