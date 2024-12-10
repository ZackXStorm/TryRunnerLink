embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"download-removebg-preview\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/level/level.atlas\"\n"
  "}\n"
  ""
  position {
    x: 1324.0
    y: -35.5
  }
  scale {
    x: 9.0
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
  "      x: 1327.0\n"
  "      y: -38.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 3\n"
  "  }\n"
  "  data: 1323.256\n"
  "  data: 33.82653\n"
  "  data: 27.2\n"
  "}\n"
  ""
}
