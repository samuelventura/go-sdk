defmodule NervesRpi3Test do
  use ExUnit.Case
  doctest NervesRpi3

  test "greets the world" do
    assert NervesRpi3.hello() == :world
  end
end
