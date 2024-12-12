components {
  id: "hero"
  component: "/main/hero/hero.script"
}
components {
  id: "CoinCollect"
  component: "/main/CoinCollect.particlefx"
  position {
    y: 300.0
    z: 1.0
  }
}
embedded_components {
  id: "spinemodel"
  type: "spinemodel"
  data: "spine_scene: \"/main/hero/hero.spinescene\"\n"
  "default_animation: \"run\"\n"
  "skin: \"\"\n"
  "material: \"/defold-spine/assets/spine.material\"\n"
  ""
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"hero\"\n"
  "mask: \"geometry\"\n"
  "mask: \"danger\"\n"
  "mask: \"pickup\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      y: 500.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  shapes {\n"
  "    shape_type: TYPE_BOX\n"
  "    position {\n"
  "      y: 200.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 1\n"
  "    count: 3\n"
  "  }\n"
  "  data: 100.0\n"
  "  data: 72.91545\n"
  "  data: 200.0\n"
  "  data: 10.0\n"
  "}\n"
  ""
}
embedded_components {
  id: "bullet_factory"
  type: "factory"
  data: "prototype: \"/main/bullet/bullet.go\"\n"
  ""
}
