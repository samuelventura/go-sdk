defmodule Server do
  use GenServer

  def init(exec) do
    exec.()
  end
end
