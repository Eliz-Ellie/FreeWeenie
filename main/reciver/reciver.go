components {
  id: "reciver"
  component: "/main/reciver/reciver.collisionobject"
}
components {
  id: "reciver1"
  component: "/main/reciver/reciver.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"LaserReciver\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/images/atlas.atlas\"\n"
  "}\n"
  ""
}
