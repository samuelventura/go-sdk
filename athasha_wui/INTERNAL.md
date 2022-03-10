# Athasha WUI

## Setup

```bash
$ elixir -v
Erlang/OTP 24 [erts-12.2.1] [source] [64-bit] [smp:8:8] [ds:8:8:10] [async-threads:1] [jit] Elixir 1.13.3 (compiled with Erlang/OTP 24)
$ mix local.hex
$ mix archive.install hex phx_new
$ mix phx.new . --module Athasha.Wui --app athasha_wui --database sqlite3 --live
$ mix ecto.create
$ mix phx.server
$ iex -S mix phx.server

https://hexdocs.pm/phoenix/ecto.html#content
mix phx.gen.schema App apps name:string enabled:boolean
https://hexdocs.pm/phoenix_live_view/Phoenix.LiveView.html
```
