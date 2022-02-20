defmodule Sample2 do
  use Bitwise

  def length_of_longest_substring(s) do
    cal(s, 0, 0, 0, s)
  end

  def cal(<<>>, max, count, _map, _sub) do
    max(count, max)
  end

  def cal(<<c, rest::binary>>, max, count, map, sub) do
    f = 1 <<< c
    d = (f &&& map) > 0

    case d do
      false ->
        map = map ||| f
        cal(rest, max, count + 1, map, sub)

      _ ->
        max = max(count, max)
        <<_, sub::binary>> = sub
        cal(rest, max, count, map, sub)
    end
  end
end
