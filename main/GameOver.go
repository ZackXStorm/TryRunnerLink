components {
  id: "GameOver"
  component: "/main/GameOver.script"
}
embedded_components {
  id: "GameLabel"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "text: \"That\\\'s Game\"\n"
  "font: \"/builtins/fonts/default.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  ""
  scale {
    x: 2.0
    y: 2.0
    z: 2.0
  }
}
embedded_components {
  id: "GameStats"
  type: "label"
  data: "size {\n"
  "  x: 128.0\n"
  "  y: 32.0\n"
  "}\n"
  "text: \"---------\"\n"
  "font: \"/builtins/fonts/default.font\"\n"
  "material: \"/builtins/fonts/label-df.material\"\n"
  ""
  position {
    y: -109.0
  }
  scale {
    x: 1.5
    y: 1.5
    z: 1.5
  }
}
