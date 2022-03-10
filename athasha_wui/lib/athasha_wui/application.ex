defmodule Athasha.Wui.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    children = [
      # Start the Ecto repository
      Athasha.Wui.Repo,
      # Start the Telemetry supervisor
      Athasha.WuiWeb.Telemetry,
      # Start the PubSub system
      {Phoenix.PubSub, name: Athasha.Wui.PubSub},
      # Start the Endpoint (http/https)
      Athasha.WuiWeb.Endpoint
      # Start a worker by calling: Athasha.Wui.Worker.start_link(arg)
      # {Athasha.Wui.Worker, arg}
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: Athasha.Wui.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  @impl true
  def config_change(changed, _new, removed) do
    Athasha.WuiWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end
