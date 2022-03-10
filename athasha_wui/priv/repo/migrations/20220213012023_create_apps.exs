defmodule Athasha.Wui.Repo.Migrations.CreateApps do
  use Ecto.Migration

  def change do
    create table(:apps) do
      add :name, :string
      add :enabled, :boolean, default: false, null: false

      timestamps()
    end
  end
end
