# Tabletizer

## Installation

```shell
go get github.com/nyan2d/tabletizer
```

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/nyan2d/tabletizer"
)

type User struct {
	ID        int    `sqlarg:"PRIMARY KEY"`
	Name      string `sqlarg:"NOT NULL"`
	Age       uint
	IsAdmin   bool
	CreatedAt time.Time `sqltype:"INTEGER"`
}

func main() {
	x := tabletizer.DoMagic(&User{}, "users")
	fmt.Println(x)
}
```

The code above will produce an output similar to the following:
```sql
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    name TEXT NOT NULL,
    age INTEGER,
    is_admin INTEGER,
    created_at INTEGER
)
```