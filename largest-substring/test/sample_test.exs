defmodule SampleTest do
  use ExUnit.Case
  doctest Sample

  test "greets the world" do
    assert 7 ==
             Sample3.length_of_longest_substring(
               "abcab|asd,k|fasdf alskfasdfasdfasdf adfas fasdfcbb"
             )

    assert 1 == Sample3.length_of_longest_substring(" ")
  end
end
