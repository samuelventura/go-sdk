defmodule LinkKilledTest do
  use ExUnit.Case

  test "Process killed on crash process exit" do
    self = self()

    pid =
      spawn(fn ->
        lid =
          spawn_link(fn ->
            receive do
              any -> any
            end
          end)

        send(self, lid)
        Process.exit(self(), :crash)
      end)

    ref = :erlang.monitor(:process, pid)
    assert_receive lid, 400
    assert_receive {:DOWN, ^ref, :process, ^pid, :crash}, 400
    assert false == Process.alive?(lid)
  end

  test "Process killed on GenServer stopped exit" do
    self = self()

    {:error, :stopped} =
      GenServer.start(Server, fn ->
        lid =
          spawn_link(fn ->
            receive do
              any -> any
            end
          end)

        send(self, {self(), lid})
        {:stop, :stopped}
      end)

    assert_receive {pid, lid}, 400
    ref = :erlang.monitor(:process, pid)
    assert_receive {:DOWN, ^ref, :process, ^pid, :stopped}, 400
    assert false == Process.alive?(lid)
  end

  test "Process killed on GenServer tuple exit" do
    self = self()

    {:error, {:stopped, :now}} =
      GenServer.start(Server, fn ->
        lid =
          spawn_link(fn ->
            receive do
              any -> any
            end
          end)

        send(self, {self(), lid})
        {:stop, {:stopped, :now}}
      end)

    assert_receive {pid, lid}, 400
    ref = :erlang.monitor(:process, pid)
    assert_receive {:DOWN, ^ref, :process, ^pid, {:stopped, :now}}, 400
    assert false == Process.alive?(lid)
  end
end
