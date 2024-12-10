components {
  id: "platform"
  component: "/main/level/platform.script"
}
embedded_components {
  id: "plank"
  type: "sprite"
  data: "default_animation: \"rock_planks\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/level/level.atlas\"\n"
  "}\n"
  ""
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
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 195.82959\n"
  "  data: 76.21272\n"
  "  data: 27.2\n"
  "}\n"
  ""
}
embedded_components {
  id: "spike"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: -212.0
    y: -4.0
  }
  scale {
    y: 1.164316
  }
}
embedded_components {
  id: "spike1"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 51.0\n"
  "  y: 131.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 207.0
    y: -4.0
  }
  rotation {
    y: 1.0
    w: 6.123234E-17
  }
  scale {
    y: 1.164316
  }
}
embedded_components {
  id: "spike2"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 51.0\n"
  "  y: 131.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 99.0
    y: -95.0
  }
  rotation {
    x: 0.70710677
    y: 0.70710677
    z: 4.3297803E-17
    w: 4.3297803E-17
  }
}
embedded_components {
  id: "spike3"
  type: "sprite"
  data: "default_animation: \"spikes\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 51.0\n"
  "  y: 131.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: -99.0
    y: -95.0
  }
  rotation {
    x: 0.70710677
    y: 0.70710677
    z: 4.3297803E-17
    w: 4.3297803E-17
  }
}
embedded_components {
  id: "danger_edges"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"danger\"\n"
  "mask: \"hero\"\n"
  "mask: \"bullet\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: 190.0\n"
  "      y: -12.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -190.0\n"
  "      y: -12.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 3\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: -78.0\n"
  "    }\n"
  "    rotation {\n"
  "      z: 0.70710677\n"
  "      w: 0.70710677\n"
  "    }\n"
  "    index: 6\n"
  "    count: 3\n"
  "  }\n"
  "  data: 25.33963\n"
  "  data: 69.91381\n"
  "  data: 10.0\n"
  "  data: 25.33963\n"
  "  data: 69.91381\n"
  "  data: 10.0\n"
  "  data: 15.0\n"
  "  data: 200.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "coin_factory"
  type: "factory"
  data: "prototype: \"/main/level/level_objekts/coin/coin.go\"\n"
  ""
}
