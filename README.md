# Re:Earth CMS MCP Server

> MCP server for [Re:Earth CMS](https://reearth.io/product/cms).

## Usecases

Here are some other useful prompts to get started:

| Usecase | Prompt |
|---------|--------|
| Add comment | `Add a comment "Needs review" to item XXX` |
| Bulk update | `Update all items in "locations" model to set status to "published"` |
| Copy structure | `Copy the "template" model and name it "new-project"` |
| Create item | `Create a new item in the "events" model with title "Summer Festival" and date "2025-08-15"` |
| Data summary | `Summarize the data in the "sensors" model` |
| Explore models | `What models are available in project XXX?` |
| Export data | `Export all items from "buildings" model as GeoJSON` |
| List all projects | `Show me all projects in my workspace` |
| Publish asset | `Publish asset 123 in project abc` |
| Publish item | `Publish item 456 in model "articles"` |
| Search items | `Find all items in the "products" model that contain "Tokyo"` |
| Unpublish asset | `Unpublish asset 789 in project abc` |
| View schema | `Show me the schema for the "articles" model` |

## Features

- **Assets**: Manage assets and publish/unpublish them
- **Comments**: Add, update, and delete comments on items and assets
- **Export**: Export items as CSV or GeoJSON
- **Groups**: Create, update, and delete groups
- **Items**: Create, read, update, delete, and publish items
- **Models**: List models and retrieve schema definitions
- **Projects**: List and manage projects in your workspace

Note that this MCP doesn't have capabilities to manage workspaces, users and model's schemas.

## Installation

### Docker (Recommended)

```bash
docker pull reearth/reearth-cms-mcp
```

### Go

```bash
go install github.com/reearth/reearth-cms-mcp@latest
```

## Configuration

### Prerequisites

This MCP server uses the Re:Earth CMS integration API, so you need the following information:

**How to obtain:**

- **Integration Token**: Create an integration in your Re:Earth CMS workspace settings. (Documentation is not yet available)
- **Workspace ID**: You can find your workspace ID in the URL when using Re:Earth CMS editor: `https://cms.reearth.io/workspaces/{workspace-id}/...`

<details>
<summary>Claude Code</summary>

<table>
<tr><th>Docker (Recommended)</th><th>Go binary</th></tr>
<tr valign=top>
<td>

```bash
claude mcp add reearth-cms -e REEARTH_CMS_TOKEN=your-api-token -e REEARTH_CMS_WORKSPACE_ID=your-workspace-id -- docker run -i --rm -e REEARTH_CMS_TOKEN -e REEARTH_CMS_WORKSPACE_ID reearth/reearth-cms-mcp
```

</td>
<td>

```bash
claude mcp add reearth-cms -e REEARTH_CMS_TOKEN=your-api-token -e REEARTH_CMS_WORKSPACE_ID=your-workspace-id -- reearth-cms-mcp
```

</td>
</tr>
</table>

</details>

<details>
<summary>Claude Desktop</summary>

Add the following to your Claude Desktop config file:

- macOS: `~/Library/Application Support/Claude/claude_desktop_config.json`
- Windows: `%APPDATA%\Claude\claude_desktop_config.json`

<table>
<tr><th>Docker (Recommended)</th><th>Go binary</th></tr>
<tr valign=top>
<td>

```json
{
  "mcpServers": {
    "reearth-cms": {
      "command": "docker",
      "args": ["run", "-i", "--rm", "-e", "REEARTH_CMS_TOKEN", "-e", "REEARTH_CMS_WORKSPACE_ID", "reearth/reearth-cms-mcp"],
      "env": {
        "REEARTH_CMS_TOKEN": "your-api-token",
        "REEARTH_CMS_WORKSPACE_ID": "your-workspace-id"
      }
    }
  }
}
```

</td>
<td>

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

</td>
</tr>
</table>

</details>

<details>
<summary>Codex</summary>

<table>
<tr><th>Docker (Recommended)</th><th>Go binary</th></tr>
<tr valign=top>
<td>

```toml
[mcp_servers.reearth-cms]
args = ["run", "-i", "--rm", "reearth/reearth-cms-mcp"]
command = "docker"

[mcp_servers.reearth-cms.env]
REEARTH_CMS_TOKEN = "you-api-token"
REEARTH_CMS_WORKSPACE_ID = "your-workspace-id"
```

</td>
<td>

Add the following to your Codex MCP configuration:

```toml
[mcp_servers.reearth-cms]
command = "reearth-cms-mcp"

[mcp_servers.reearth-cms.env]
REEARTH_CMS_TOKEN = "you-api-token"
REEARTH_CMS_WORKSPACE_ID = "your-workspace-id"
```

</td>
</tr>
</table>

</details>

<details>
<summary>Cursor</summary>

<table>
<tr><th>Docker (Recommended)</th><th>Go binary</th></tr>
<tr valign=top>
<td>

```json
{
  "mcpServers": {
    "reearth-cms": {
      "command": "docker",
      "args": ["run", "-i", "--rm", "-e", "REEARTH_CMS_TOKEN", "-e", "REEARTH_CMS_WORKSPACE_ID", "reearth/reearth-cms-mcp"],
      "env": {
        "REEARTH_CMS_TOKEN": "your-api-token",
        "REEARTH_CMS_WORKSPACE_ID": "your-workspace-id"
      }
    }
  }
}
```

</td>
<td>

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

</td>
</tr>
</table>

</details>

<details>
<summary>Raycast</summary>

<table>
<tr><th>Docker (Recommended)</th><th>Go binary</th></tr>
<tr><th align=left colspan=2>Raycast</th></tr>
<tr valign=top>
<td>

| Field | Value |
|-------|-------|
| Name | `reearth-cms` |
| Command | `docker` |
| Arguments | `run -i --rm -e REEARTH_CMS_TOKEN -e REEARTH_CMS_WORKSPACE_ID reearth/reearth-cms-mcp` |

Or via config file (`~/.config/raycast/mcp.json`):

```json
{
  "mcpServers": {
    "reearth-cms": {
      "command": "docker",
      "args": ["run", "-i", "--rm", "-e", "REEARTH_CMS_TOKEN", "-e", "REEARTH_CMS_WORKSPACE_ID", "reearth/reearth-cms-mcp"],
      "env": {
        "REEARTH_CMS_TOKEN": "your-api-token",
        "REEARTH_CMS_WORKSPACE_ID": "your-workspace-id"
      }
    }
  }
}
```

</td>
<td>

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

</td>
</tr>
</table>

</details>

<details>
<summary>VS Code</summary>

<table>
<tr><th>Docker (Recommended)</th><th>Go binary</th></tr>
<tr><th align=left colspan=2>VS Code (version 1.101 or greater)</th></tr>
<tr valign=top>
<td>

```json
{
  "mcp": {
    "servers": {
      "reearth-cms": {
        "command": "docker",
        "args": ["run", "-i", "--rm", "-e", "REEARTH_CMS_TOKEN", "-e", "REEARTH_CMS_WORKSPACE_ID", "reearth/reearth-cms-mcp"],
        "env": {
          "REEARTH_CMS_TOKEN": "your-api-token",
          "REEARTH_CMS_WORKSPACE_ID": "your-workspace-id"
        }
      }
    }
  }
}
```

</td>
<td>

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

</td>
</tr>
</table>

</details>

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](./LICENSE) file for details.
