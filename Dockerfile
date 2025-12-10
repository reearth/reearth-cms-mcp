FROM gcr.io/distroless/static:nonroot

COPY reearth-cms-mcp /reearth-cms-mcp

ENTRYPOINT ["/reearth-cms-mcp"]
