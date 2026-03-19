components {
  id: "wall"
  component: "/main/wall/wall.collisionobject"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"laser_segment\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/images/atlas.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 4.66307
  }
}
