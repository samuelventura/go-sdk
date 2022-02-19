defmodule LinkLeaksTest do
  use ExUnit.Case

  test "Process leaks on normal process exit" do
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
      end)

    ref = :erlang.monitor(:process, pid)
    assert_receive lid, 400
    assert_receive {:DOWN, ^ref, :process, ^pid, :normal}, 400
    assert true == Process.alive?(lid)
  end

  test "Process leak on GenServer stop call" do
    self = self()

    {:ok, pid} =
      GenServer.start(Server, fn ->
        lid =
          spawn_link(fn ->
            receive do
              any -> any
            end
          end)

        send(self, lid)
        {:ok, lid}
      end)

    assert_receive lid, 400
    ref = :erlang.monitor(:process, pid)
    GenServer.stop(pid)
    assert_receive {:DOWN, ^ref, :process, ^pid, :normal}, 400
    assert true == Process.alive?(lid)
  end

  test "Process leaks on GenServer normal stop" do
    self = self()

    {:error, :normal} =
      GenServer.start(Server, fn ->
        lid =
          spawn_link(fn ->
            receive do
              any -> any
            end
          end)

        send(self, {self(), lid})
        {:stop, :normal}
      end)

    assert_receive {pid, lid}, 400
    ref = :erlang.monitor(:process, pid)
    assert_receive {:DOWN, ^ref, :process, ^pid, :noproc}, 400
    assert true == Process.alive?(lid)
  end

  test "Agent leaks on normal process exit" do
    self = self()

    pid =
      spawn(fn ->
        {:ok, lid} = Agent.start_link(fn -> 0 end)
        send(self, lid)
      end)

    ref = :erlang.monitor(:process, pid)
    assert_receive lid, 400
    assert_receive {:DOWN, ^ref, :process, ^pid, :normal}, 400
    assert true == Process.alive?(lid)
  end

  test "Supervisor leaks on normal process exit" do
    self = self()

    pid =
      spawn(fn ->
        {:ok, lid} = Supervisor.start_link([], strategy: :one_for_one)
        send(self, lid)
      end)

    ref = :erlang.monitor(:process, pid)
    assert_receive lid, 400
    assert_receive {:DOWN, ^ref, :process, ^pid, :normal}, 400
    assert true == Process.alive?(lid)
  end
end
