components {
  id: "horizontal_moving"
  component: "/main/horizontal_moving.script"
}
embedded_components {
  id: "Sun"
  type: "sprite"
  data: "default_animation: \"Sun\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/level/level.atlas\"\n"
  "}\n"
  ""
}
