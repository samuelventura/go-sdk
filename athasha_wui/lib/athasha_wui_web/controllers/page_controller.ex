defmodule Athasha.WuiWeb.PageController do
  use Athasha.WuiWeb, :controller

  def index(conn, _params) do
    render(conn, "index.html")
  end
end
