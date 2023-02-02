# How to use ./internal/ package path for internal project packages

Make sure import path is of the following syntax:
```go
import "module-name/internal/package-folder"
```

Make sure to use the package folder, not any specific file.
Folder internal should be protected by go from external import.
Package-folder should have the same name as package, but in the case they do not, it is the package-folder that needs to be in the import path.
