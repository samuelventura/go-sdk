defmodule Athasha.Wui.App do
  use Ecto.Schema
  import Ecto.Changeset

  schema "apps" do
    field :enabled, :boolean, default: false
    field :name, :string

    timestamps()
  end

  @doc false
  def changeset(app, attrs) do
    app
    |> cast(attrs, [:name, :enabled])
    |> validate_required([:name, :enabled])
  end
end
