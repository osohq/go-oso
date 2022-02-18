# Go Client

## Requirements

- Go >=1.14

## Installation

- `go get github.com/osohq/go-oso@client`
  - Note the `@client` suffix, which pulls in the library as it exists on the
    `client` branch of the `github.com/osohq/go-oso` repository.
  - This should add a line to your `go.mod` file similar to:

    ```go
    require github.com/osohq/go-oso v0.0.0-20220210233915-6a23048d4884 // indirect
    ```

- Import the library:

  ```go
  import (
    ...

    oso "github.com/osohq/go-oso"
  )
  ```

- Use it:

  ```go
  oso := oso.NewClient("https://<tenant>.oso.run")
  allowed, e := oso.Authorize(User{id: 1}, "read", Repo{id: 2})
  if e != nil {
  	log.Fatalln(e)
  }
  log.Printf("Authorize: %v\n", allowed)
  ```

## Testing

- Run local `oso-service`: `cargo run -- run`
- `go run cmd/oso.go`
