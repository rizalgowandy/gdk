# Stack

Stack is package to do operation on debug stack. This package should be used to reduce unimportant stack trace. Smaller stack trace helps to debug trace the error faster.

# Warning

**This package doesn't work when imported as vendor!**

You should copy the code to each repository and modify:

```go
var (
    // replace with the path where you place the code
    functionPackageKeyword = []byte("/github.com/rizalgowandy/gdk/pkg/stack/")
)
```