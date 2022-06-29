defmodule X86nervesTest do
  use ExUnit.Case
  doctest X86nerves

  test "greets the world" do
    assert X86nerves.hello() == :world
  end
end
