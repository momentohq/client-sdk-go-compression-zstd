<img src="https://docs.momentohq.com/img/momento-logo-forest.svg" alt="logo" width="400"/>

[![project status](https://momentohq.github.io/standards-and-practices/badges/project-status-official.svg)](https://github.com/momentohq/standards-and-practices/blob/main/docs/momento-on-github.md)
[![project stability](https://momentohq.github.io/standards-and-practices/badges/project-stability-alpha.svg)](https://github.com/momentohq/standards-and-practices/blob/main/docs/momento-on-github.md)


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

----------------------------------------------------------------------------------------
For more info, visit our website at [https://gomomento.com](https://gomomento.com)!
