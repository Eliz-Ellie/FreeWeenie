components {
  id: "portal"
  component: "/main/portal/portal.collisionobject"
}
components {
  id: "portal1"
  component: "/main/portal/portal.script"
}
components {
  id: "wall"
  component: "/main/portal/portalback.collisionobject"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"portal1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/images/atlas.atlas\"\n"
  "}\n"
  ""
}
