package main

import (
	"io"
	"log/slog"
	"net/http"
)

func HTTPRequestAttr(key string, req *http.Request) slog.Attr {
	var body string

	if req.GetBody != nil {
		rc, err := req.GetBody()
		if err != nil {
			body = err.Error()
		}
		bytes, err := io.ReadAll(rc)
		if err != nil {
			body = err.Error()
		}
		body = string(bytes)
	}

	return slog.Group(key,
		slog.String("method", req.Method),
		slog.String("url", req.URL.String()),
		HTTPHeaderAttr("headers", req.Header),
		slog.String("body", body),
	)
}

func HTTPResponseAttr(key string, respWriter http.ResponseWriter, body string) slog.Attr {
	return slog.Group(key,
		slog.Int("status code", http.StatusOK),
		slog.String("status", http.StatusText(http.StatusOK)),
		HTTPHeaderAttr("headers", respWriter.Header()),
		slog.String("body", body),
	)
}

func HTTPHeaderAttr(key string, value http.Header) slog.Attr {
	args := make([]any, 0, len(value))

	for k, v := range value {
		args = append(args, slog.Any(k, slog.AnyValue(v)))
	}

	return slog.Group(
		key,
		args...,
	)
}
