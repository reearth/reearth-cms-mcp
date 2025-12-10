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
| Search items | `Find all items in the "products" model that contain "Tokyo"` |
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

#### Using Go binary

```bash
claude mcp add reearth-cms -e REEARTH_CMS_TOKEN=your-api-token -e REEARTH_CMS_WORKSPACE_ID=your-workspace-id -- reearth-cms-mcp
```

Or set the environment variables in your shell and run:

```bash
export REEARTH_CMS_TOKEN="your-api-token"
export REEARTH_CMS_WORKSPACE_ID="your-workspace-id"
claude mcp add reearth-cms -- reearth-cms-mcp
```

#### Using Docker

```bash
claude mcp add reearth-cms -e REEARTH_CMS_TOKEN=your-api-token -e REEARTH_CMS_WORKSPACE_ID=your-workspace-id -- docker run -i --rm -e REEARTH_CMS_TOKEN -e REEARTH_CMS_WORKSPACE_ID reearth/reearth-cms-mcp
```

</details>

<details>
<summary>Claude Desktop</summary>

Add the following to your Claude Desktop config file:

- macOS: `~/Library/Application Support/Claude/claude_desktop_config.json`
- Windows: `%APPDATA%\Claude\claude_desktop_config.json`

#### Using Go binary

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

#### Using Docker

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

</details>

<details>
<summary>Codex</summary>

#### Using Go binary

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

#### Using Docker

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

</details>

<details>
<summary>Cursor</summary>

#### Using Go binary

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

#### Using Docker

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

</details>

<details>
<summary>Raycast</summary>

#### Using Go binary

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

#### Using Docker

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

</details>

<details>
<summary>VS Code</summary>

#### Using Go binary

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

#### Using Docker

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

</details>

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](./LICENSE) file for details.
