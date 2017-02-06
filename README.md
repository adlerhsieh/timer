# Timer

Timer is a package that prints out current time.

## Installtion

```
go get github.com/adlerhsieh/timer
```

## Example

```
package main

import(
  "fmt"
  "github.com/adlerhsieh/timer"
)

func main() {
  t := timer.Now()
  fmt.Println(t) // => "2017-01-01 10:12:25"
}
```
