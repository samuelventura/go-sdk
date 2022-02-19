defmodule SocketKilledTest do
  use ExUnit.Case

  test "Listener closed on normal process exit" do
    self = self()

    pid =
      spawn(fn ->
        ip = {127, 0, 0, 1}
        opts = [:binary, ip: ip, packet: :raw, active: false]
        {:ok, listener} = :gen_tcp.listen(0, opts)
        {:ok, {^ip, port}} = :inet.sockname(listener)
        send(self, {listener, ip, port})

        receive do
          any -> any
        end
      end)

    ref = :erlang.monitor(:process, pid)
    assert_receive {listener, ip, port}, 400
    {:ok, {^ip, ^port}} = :inet.sockname(listener)
    send(pid, :continue)
    assert_receive {:DOWN, ^ref, :process, ^pid, :normal}, 400
    assert {:error, :einval} == :inet.sockname(listener)
  end

  test "Client closed on normal process exit" do
    self = self()
    ip = {127, 0, 0, 1}
    opts = [:binary, ip: ip, packet: :raw, active: false]
    {:ok, listener} = :gen_tcp.listen(0, opts)
    {:ok, {^ip, port}} = :inet.sockname(listener)

    pid =
      spawn(fn ->
        opts = [:binary, packet: :raw, active: false]
        {:ok, client} = :gen_tcp.connect(ip, port, opts)
        {:ok, {^ip, port}} = :inet.sockname(client)
        send(self, {client, ip, port})

        receive do
          any -> any
        end
      end)

    ref = :erlang.monitor(:process, pid)
    assert_receive {client, ip, port}, 400
    {:ok, {^ip, ^port}} = :inet.sockname(client)
    send(pid, :continue)
    assert_receive {:DOWN, ^ref, :process, ^pid, :normal}, 400
    assert {:error, :einval} == :inet.sockname(client)
  end
end
