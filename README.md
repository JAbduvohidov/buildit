## Build your go project for different OS

**works only in windows 10**

_install the project_

````
go install github.com/JAbduvohidov/buildit
````

use command below to build your project

````
buildit -for [linux|windows|me] -name fileName -v [major|feat|fix]
````

version file content
````go
package version

import "fmt"

const (
	major = 1
	minor = 0
	patch = 0
)

func Current() string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}
````