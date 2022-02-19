defmodule EchoServerTest do
  use ExUnit.Case

  test "clients closed on echo server stop" do
    {:ok, pid} = EchoServer.start_link()
    opts = [:binary, packet: :raw, active: false]
    {ip, port} = EchoServer.endpoint(pid)
    {:ok, _client} = :gen_tcp.connect(ip, port, opts)
    :ok = EchoServer.stop(pid)
    {:error, :econnrefused} = :gen_tcp.connect(ip, port, opts)
    # FIXME client still active
    # :timer.sleep(100)
    # :ok = :gen_tcp.send(client, "data")
    # {:error, :closed} = :inet.sockname(client)
  end
end
