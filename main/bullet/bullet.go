components {
  id: "bullet"
  component: "/main/bullet/bullet.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"bullet\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 16.0\n"
  "  y: 7.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/bullet/bullet.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 10.0
    y: 10.0
    z: 10.0
  }
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"bullet\"\n"
  "mask: \"geometry\"\n"
  "mask: \"danger\"\n"
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
  "  data: 79.2659\n"
  "  data: 34.65116\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
