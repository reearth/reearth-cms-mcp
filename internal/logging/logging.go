package logging

import (
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

// Logger returns a slog.Logger that sends logs to the MCP client.
// If the session is nil or DEBUG environment variable is not set,
// it returns a stderr logger or no-op logger.
func Logger(session *mcp.ServerSession) *slog.Logger {
	if session != nil {
		return slog.New(mcp.NewLoggingHandler(session, &mcp.LoggingHandlerOptions{
			LoggerName:  "reearth-cms",
			MinInterval: 50 * time.Millisecond,
		}))
	}
	return NewStderrLogger()
}

// NewStderrLogger creates a logger that writes to stderr.
// This is useful for debugging when no session is available.
func NewStderrLogger() *slog.Logger {
	if os.Getenv("DEBUG") == "" {
		return slog.New(slog.NewTextHandler(io.Discard, nil))
	}
	return slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	}))
}

// NewLoggingTransport wraps a transport with logging capabilities.
// Logs are written to the provided writer (typically os.Stderr for MCP servers).
func NewLoggingTransport(transport mcp.Transport, w io.Writer) *mcp.LoggingTransport {
	return &mcp.LoggingTransport{
		Transport: transport,
		Writer:    w,
	}
}

// IsDebugEnabled returns true if DEBUG environment variable is set.
func IsDebugEnabled() bool {
	return os.Getenv("DEBUG") != ""
}
