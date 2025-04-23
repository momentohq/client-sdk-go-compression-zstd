package test_helpers

import (
	"fmt"

	"github.com/momentohq/client-sdk-go-compression-zstd/zstd_compression"
	"github.com/momentohq/client-sdk-go/config/compression"
	"github.com/momentohq/client-sdk-go/config/logger"
	"github.com/momentohq/client-sdk-go/config/middleware"
	impl "github.com/momentohq/client-sdk-go/config/middleware/impl"
)

// ZstdTestCompressorFactory is a wrapper around the ZstdCompressorFactory that allows us to test the
// zstd compression middleware more thoroughly to confirm compression is working as expected.
type ZstdTestCompressorFactory struct {
	CompressedDataChannel   chan int // receive data size in bytes
	DecompressedDataChannel chan int // receive data size in bytes
}

func (f ZstdTestCompressorFactory) NewCompressionStrategy(props compression.CompressionStrategyProps) compression.CompressionStrategy {
	compressionStrategy := zstdTestCompressor{
		compressor:              zstd_compression.ZstdCompressorFactory{}.NewCompressionStrategy(props),
		logger:                  props.Logger,
		CompressedDataChannel:   f.CompressedDataChannel,
		DecompressedDataChannel: f.DecompressedDataChannel,
	}
	return compressionStrategy
}

type zstdTestCompressor struct {
	compressor              compression.CompressionStrategy
	logger                  logger.MomentoLogger
	CompressedDataChannel   chan int // receive data size in bytes
	DecompressedDataChannel chan int // receive data size in bytes
}

func (h zstdTestCompressor) Compress(data []byte) ([]byte, error) {
	compressed, err := h.compressor.Compress(data)
	if err != nil {
		return nil, fmt.Errorf("failed to compress data: %v", err)
	}

	h.logger.Trace("Compressed data: %d bytes -> %d bytes", len(data), len(compressed))
	if h.CompressedDataChannel != nil {
		h.CompressedDataChannel <- len(compressed)
	}
	return compressed, nil
}

func (h zstdTestCompressor) Decompress(data []byte) ([]byte, error) {
	decompressed, err := h.compressor.Decompress(data)
	if err != nil {
		return nil, fmt.Errorf("failed to decompress data: %v", err)
	}

	h.logger.Trace("Decompressed data: %d bytes -> %d bytes", len(data), len(decompressed))
	if h.DecompressedDataChannel != nil {
		h.DecompressedDataChannel <- len(decompressed)
	}
	return decompressed, nil
}

type ZstdCompressionTestMiddlewareProps struct {
	IncludeTypes            []interface{}
	CompressionLevel        compression.CompressionLevel
	Logger                  logger.MomentoLogger
	CompressedDataChannel   chan int // receive data size in bytes
	DecompressedDataChannel chan int // receive data size in bytes
}

// NewZstdCompressionTestMiddleware creates a new compression middleware as a test helper
// for verifying compression is working as expected.
func NewZstdCompressionTestMiddleware(props ZstdCompressionTestMiddlewareProps) middleware.Middleware {
	compressionMiddlewareProps := impl.CompressionMiddlewareProps{
		CompressorFactory: ZstdTestCompressorFactory{
			CompressedDataChannel:   props.CompressedDataChannel,
			DecompressedDataChannel: props.DecompressedDataChannel,
		},
		CompressionStrategyProps: compression.CompressionStrategyProps{
			CompressionLevel: props.CompressionLevel,
			Logger:           props.Logger,
		},
		IncludeTypes: props.IncludeTypes,
	}
	return impl.NewCompressionMiddleware(compressionMiddlewareProps)
}
