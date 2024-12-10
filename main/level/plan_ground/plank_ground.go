components {
  id: "platform"
  component: "/main/level/platform.script"
}
embedded_components {
  id: "plank"
  type: "sprite"
  data: "default_animation: \"rock_planks\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 391.0\n"
  "  y: 153.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 196.0
    y: -48.0
  }
  scale {
    y: 0.625657
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"geometry\"\n"
  "mask: \"hero\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 196.0\n"
  "      y: -49.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 186.42224\n"
  "  data: 37.677197\n"
  "  data: 27.2\n"
  "}\n"
  ""
}
embedded_components {
  id: "coin_factory"
  type: "factory"
  data: "prototype: \"/main/level/level_objekts/coin/coin.go\"\n"
  ""
}
