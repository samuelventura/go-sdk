defmodule X86kioskTest do
  use ExUnit.Case
  doctest X86kiosk

  test "greets the world" do
    assert X86kiosk.hello() == :world
  end
end
