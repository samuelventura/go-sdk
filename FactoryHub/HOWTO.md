# HOWTO

```bash
dotnet new blazorserver -f net6.0 --no-https
dotnet watch

dotnet add package LiteDB
```

## Architecture

- IO Server
  - Grouped by device
  - A custom IO server (OPC)
  - API (change and read event)
- Data Loggers, Web Views, DB Links
  - Connect to io-server API

## Use Cases

- Laurel Monitor
  - No buttons
  - Web viewer
  - Data recording
  - Alarms
- Modbus Monitor
  - Digital or analog
  - Widget

## Research

- Installer
- Auth2
- License
- Setup (TCP ports, ...)
- HTTPS
