# jsonenv

This package parse JSON-formatted environment value and set environment with key-value.

Typical usecase is AWS ECS using Secret Manager. On running ECS task with `secrets` task definition,
the runtime assignes AWS Secret Manager's name as environment key and JSON-Stringified string as environment value.
This package intends to parse and extract, and set to environment.

## Usage

```go
import (
    "fmt"
    "log"
    "os"

    "github.com/ysugimoto/jsonenv"
)

// Case: json env exists like `example-secret={"FOO":"BAR"}` via AWS Secrets Manager
func main() {
    if err := jsonenv.Extract("example-secret"); err != nil {
        log.Fatalln(err)
    }
    // do something with extracted environment variables.
    fmt.Println(os.Getenv("FOO")) // outputs BAR
}
```

## Test

```shell
go test
```

## License

MIT

## Author

Yoshiaki Sugimoto
