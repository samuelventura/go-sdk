defmodule x86nervesTest do
  use ExUnit.Case
  doctest x86nerves

  test "greets the world" do
    assert x86nerves.hello() == :world
  end
end
