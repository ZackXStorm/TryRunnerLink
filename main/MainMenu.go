components {
  id: "menuManeger"
  component: "/main/menuManeger.script"
}
embedded_components {
  id: "Title"
  type: "label"
  data: "size {\n"
  "  x: 400.0\n"
  "  y: 32.0\n"
  "}\n"
  "tracking: 0.2\n"
  "line_break: true\n"
  "text: \"Welcome to the Endless runner Marcus edition\"\n"
  "font: \"/main/Fonts/Space Font2.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  ""
  position {
    y: 141.0
  }
  scale {
    x: 1.5
    y: 1.5
    z: 1.5
  }
}
embedded_components {
  id: "Promt"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "pivot: PIVOT_W\n"
  "text: \"Choose a game mode:\"\n"
  "font: \"/main/Fonts/Space Font2.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  ""
  position {
    x: -220.0
    y: 76.0
  }
  scale {
    z: 1.5
  }
}
embedded_components {
  id: "Choice1"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "pivot: PIVOT_W\n"
  "text: \"*[1] Endless\"\n"
  "font: \"/main/Fonts/Space Font2.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  ""
  position {
    x: -188.0
    y: 45.0
  }
  scale {
    z: 1.5
  }
}
embedded_components {
  id: "Choice2"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "pivot: PIVOT_W\n"
  "text: \"*[2] Timed\"\n"
  "font: \"/main/Fonts/Space Font2.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  ""
  position {
    x: -187.0
    y: 13.0
  }
  scale {
    z: 1.5
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"Space button\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "size {\n"
  "  x: 3264.0\n"
  "  y: 2448.0\n"
  "}\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/main/game.atlas\"\n"
  "}\n"
  ""
  position {
    y: 72.0
  }
  scale {
    x: 8.045003
    y: 7.762527
  }
}
