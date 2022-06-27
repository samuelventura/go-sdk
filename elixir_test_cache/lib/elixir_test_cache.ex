defmodule ElixirTestCache do
  @moduledoc """
  Documentation for `ElixirTestCache`.
  """

  @doc """
  Hello world.

  ## Examples

      iex> ElixirTestCache.hello()
      :world

  """
  def hello do
    IO.inspect(System.get_env("MIX_ENV"))
    :world2
  end
end
