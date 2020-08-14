# jsonenv

This package parse JSON-formatted environment value and set environment with key-value.

## Typical usecase

AWS ECS using Secret Manager. On running ECS task with `secrets`, the runtime assignes AWS Secret Manager's name as environment key and stringified string as environment value. This package indends to parse and extract it.

## Usage

```
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
    // do something with extracted envrironments.
    fmt.Println(os.Getenv("FOO")) // outputs BAR
}
```

## License

MIT

## Author

Yoshiaki Sugimoto
