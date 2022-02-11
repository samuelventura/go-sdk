defmodule Sample do
  def length_of_longest_substring(s) do
    cal(s, 0, 0, %{}, 0)
  end

  def cal(<<>>, max, pos, _map, rpos) do
    max(pos - rpos, max)
  end

  def cal(<<c, rest::binary>>, max, pos, map, rpos) do
    # IO.inspect {pos, c, max, map}
    m = map[c]

    case m do
      nil ->
        map = Map.put(map, c, {pos, rest})
        cal(rest, max, pos + 1, map, rpos)

      _ ->
        {ppos, prest} = m
        # IO.inspect {pos, rpos, pos - rpos, max}
        nmax = max(pos - rpos, max)

        cal(prest, nmax, ppos + 1, %{}, ppos + 1)
    end
  end
end
