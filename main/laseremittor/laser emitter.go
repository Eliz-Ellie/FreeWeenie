components {
  id: "laser emitter"
  component: "/main/laseremittor/laseremitter.script"
}
components {
  id: "laser"
  component: "/main/laseremittor/laser.collisionobject"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"laserEmitter\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/images/atlas.atlas\"\n"
  "}\n"
  ""
}
