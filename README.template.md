{{ ossHeader }}

## Packages

The zstd compression middleware for the Momento Golang SDK is available here on github: [momentohq/client-sdk-go-compression-zstd](https://github.com/momentohq/client-sdk-go-compression-zstd).

```bash
go get github.com/momentohq/client-sdk-go-compression-zstd
```

## Usage

To use the middleware, create a new instance of the zstd compression middleware and add it to the cache configuration from the [Momento Golang SDK](https://github.com/momentohq/client-sdk-go).

```go
config := config.LaptopLatestWithLogger(loggerFactory).WithMiddleware(
  []middleware.Middleware{
    zstd_compression.NewZstdCompressionMiddleware(zstd_compression.ZstdCompressionMiddlewareProps{
      CompressionStrategyProps: compression.CompressionStrategyProps{
        CompressionLevel: compression.CompressionLevelDefault,
        Logger:           loggerFactory.GetLogger("compression-middleware"),
      },
    }),
  },
)
```

{{ ossFooter }}