defmodule ElixirTestCacheTest do
  use ExUnit.Case
  # check for documentation examples
  # doctest ElixirTestCache

  test "greets the world" do
    IO.inspect(ElixirTestCache.hello())
    assert ElixirTestCache.hello() == :world2
  end
end
