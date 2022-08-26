defmodule ElixirTryout do
  @moduledoc """
  Documentation for `ElixirTryout`.
  """

  @doc """
  Hello world.

  ## Examples

      iex> ElixirTryout.hello()
      :world

  """
  def hello do
    :world
  end

  def hello2(name) do
    IO.inspect("Hello #{name}")
  end

  def sum(list) do
    _sum(list, 0)
  end

  defp _sum([], acc) do
    acc
  end

  defp _sum([c | t], acc) do
    c + _sum(t, acc)
  end

  def double(list) do
    _double(list, [])
  end

  defp _double([], acc) do
    acc
  end

  defp _double([c | t], acc) do
    [2 * c | _double(t, acc)]
  end

  def sum2(list, acc \\ 0) do
    case list do
      [] -> acc
      [c | t] -> c + sum2(t, acc)
    end
  end

  def sum3(list, acc \\ 0) do
    if length(list) == 0 do
      acc
    else
      [c | t] = list
      c + sum3(t, acc)
    end
  end

  def sum4(list) do
    Enum.reduce(list, 0, fn c, acc -> c + acc end)
  end

  def double2(list) do
    Enum.map(list, fn c -> 2 * c end)
  end

  def double3(list) do
    Enum.map(list, &(&1 * 2))
  end
end
