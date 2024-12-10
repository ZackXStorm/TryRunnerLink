components {
  id: "platform"
  component: "/main/level/platform.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"rock_planks\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 175.0
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
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 369.15002\n"
  "  data: 69.746185\n"
  "  data: 27.2\n"
  "}\n"
  ""
}
embedded_components {
  id: "sprite1"
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
    x: -175.0
  }
}
embedded_components {
  id: "coin_factory"
  type: "factory"
  data: "prototype: \"/main/level/level_objekts/coin/coin.go\"\n"
  ""
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
  "      x: 370.0\n"
  "      y: -21.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      x: -369.0\n"
  "      y: -21.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 3\n"
  "    count: 3\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: -63.0\n"
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
  "  data: 27.3\n"
  "  data: 364.0\n"
  "  data: 18.2\n"
  "}\n"
  ""
}
embedded_components {
  id: "spike"
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
    x: -394.0
    y: -8.0
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
    x: 394.0
    y: -8.0
  }
  rotation {
    z: 1.0
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
    x: 289.0
    y: -84.0
  }
  rotation {
    z: 0.70710677
    w: 0.70710677
  }
  scale {
    y: 1.164316
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
    x: 99.0
    y: -84.0
  }
  rotation {
    z: 0.70710677
    w: 0.70710677
  }
  scale {
    y: 1.164316
  }
}
embedded_components {
  id: "spike4"
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
    x: -102.0
    y: -84.0
  }
  rotation {
    z: 0.70710677
    w: 0.70710677
  }
  scale {
    y: 1.164316
  }
}
embedded_components {
  id: "spike5"
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
    x: -286.0
    y: -84.0
  }
  rotation {
    z: 0.70710677
    w: 0.70710677
  }
  scale {
    y: 1.164316
  }
}
