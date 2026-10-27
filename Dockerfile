FROM gcr.io/distroless/static@sha256:f7f8f729987ad0fdf6b05eeeae94b26e6a0f613bdf46feea7fc40f7bd72953e6 # nonroot

COPY reearth-cms-mcp /reearth-cms-mcp

ENTRYPOINT ["/reearth-cms-mcp"]
