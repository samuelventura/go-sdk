defmodule Athasha.Wui.Repo do
  use Ecto.Repo,
    otp_app: :athasha_wui,
    adapter: Ecto.Adapters.SQLite3
end
