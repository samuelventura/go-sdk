defmodule ElixirTryoutTest do
  use ExUnit.Case
  doctest ElixirTryout

  test "greets the world" do
    assert ElixirTryout.hello() == :world
  end
end
