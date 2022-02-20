defmodule Sample3 do
  def length_of_longest_substring(s) do
    cal(s, 0, 0, %{})
  end

  def cal(<<>>, max, count, _map) do
    max(count, max)
  end

  def cal(<<c, rest::binary>>, max, count, map) do
    prest = map[c]

    case prest do
      nil ->
        map = Map.put(map, c, rest)
        cal(rest, max, count + 1, map)

      _ ->
        max = max(count, max)
        cal(prest, max, 0, %{})
    end
  end
end
