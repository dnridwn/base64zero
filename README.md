# base64zero  

**base64zero** is a Go library that provides functions for encoding and decoding data using the Base64 scheme. This library is designed to simplify the conversion of binary data into a text format that is safe for transmission and storage.  

## Features  

- **Encoding**: Converts binary data into a Base64 text representation.  
- **Decoding**: Restores encoded data back gito its original form.  

## Installation  

To add **base64zero** to your project, run the following command:  

```bash
go get github.com/dnridwn/base64zero
```

## Usage
Here is a simple example of how to use **base64zero**:
```go
package main

import (
    "fmt"
    "github.com/dnridwn/base64zero"
)

func main() {
    data := []byte("Example data")
    
    encoded := base64zero.Encode(data)
    fmt.Println("Encoded:", string(encoded))

    decoded, err := base64zero.Decode(encoded)
    fmt.Println("Decoded:", string(decoded))
    fmt.Println("Decoded Err:", err)
}
```

## Contribution
Contributions are welcome! Feel free to fork this repository and submit a pull request with improvements, features, or fixes.

## License
This project is licensed under the [MIT License](LICENSE).