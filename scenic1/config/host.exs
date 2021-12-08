use Mix.Config

config :scenic1, :viewport, %{
  name: :main_viewport,
  # default_scene: {Scenic1.Scene.Crosshair, nil},
  default_scene: {Scenic1.Scene.SysInfo, nil},
  size: {800, 480},
  opts: [scale: 1.0],
  drivers: [
    %{
      module: Scenic.Driver.Glfw,
      opts: [title: "MIX_TARGET=host, app = :scenic1"]
    }
  ]
}
