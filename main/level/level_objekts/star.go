components {
  id: "horizontal_moving"
  component: "/main/horizontal_moving.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"star\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/level/level.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.5
    y: 0.5
    z: 0.5
  }
}
