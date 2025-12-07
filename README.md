# Re:Earth CMS MCP Server

> MCP server for interacting with Re:Earth CMS API.

## Installation

```bash
go install github.com/KeisukeYamashita/reearth-cms-mcp@latest
```

## Configuration

<details>
<summary>Claude Code</summary>

Run the following command:

```bash
claude mcp add reearth-cms -- reearth-cms-mcp
```

Then set the environment variables in your shell or add them to the configuration:

```bash
export REEARTH_CMS_TOKEN="your-api-token"
export REEARTH_CMS_WORKSPACE_ID="your-workspace-id"
```

Or add to `~/.claude/claude_desktop_config.json`:

```json
{
  "mcpServers": {
    "reearth-cms": {
      "command": "reearth-cms-mcp",
      "env": {
        "REEARTH_CMS_TOKEN": "your-api-token",
        "REEARTH_CMS_WORKSPACE_ID": "your-workspace-id"
      }
    }
  }
}
```

</details>

<details>
<summary>Codex</summary>

Add the following to your Codex MCP configuration:

```json
{
  "mcpServers": {
    "reearth-cms": {
      "command": "reearth-cms-mcp",
      "env": {
        "REEARTH_CMS_TOKEN": "your-api-token",
        "REEARTH_CMS_WORKSPACE_ID": "your-workspace-id"
      }
    }
  }
}
```

</details>

<details>
<summary>Cursor</summary>

Add the following to your Cursor MCP configuration (`~/.cursor/mcp.json`):

```json
{
  "mcpServers": {
    "reearth-cms": {
      "command": "reearth-cms-mcp",
      "env": {
        "REEARTH_CMS_TOKEN": "your-api-token",
        "REEARTH_CMS_WORKSPACE_ID": "your-workspace-id"
      }
    }
  }
}
```

</details>

<details>
<summary>Raycast</summary>

1. Open Raycast Settings
2. Go to **Extensions** > **AI Commands** > **MCP Servers**
3. Click **Add Server** and configure:

| Field | Value |
|-------|-------|
| Name | `reearth-cms` |
| Command | `reearth-cms-mcp` |

4. Add environment variables:

| Variable | Value |
|----------|-------|
| `REEARTH_CMS_TOKEN` | `your-api-token` |
| `REEARTH_CMS_WORKSPACE_ID` | `your-workspace-id` |

Alternatively, you can add the server via Raycast's config file (`~/.config/raycast/mcp.json`):

```json
{
  "mcpServers": {
    "reearth-cms": {
      "command": "reearth-cms-mcp",
      "env": {
        "REEARTH_CMS_TOKEN": "your-api-token",
        "REEARTH_CMS_WORKSPACE_ID": "your-workspace-id"
      }
    }
  }
}
```

</details>

<details>
<summary>VS Code</summary>

Add the following to your VS Code `settings.json`:

```json
{
  "mcp": {
    "servers": {
      "reearth-cms": {
        "command": "reearth-cms-mcp",
        "env": {
          "REEARTH_CMS_TOKEN": "your-api-token",
          "REEARTH_CMS_WORKSPACE_ID": "your-workspace-id"
        }
      }
    }
  }
}
```

</details>

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](./LICENSE) file for details.
