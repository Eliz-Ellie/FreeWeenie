components {
  id: "portal1"
  component: "/main/portal/portal.script"
}
components {
  id: "portalback_copy"
  component: "/main/portal/portalback_copy.collisionobject"
}
components {
  id: "portal_copy"
  component: "/main/portal/portal_copy.collisionobject"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"portal2\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/images/atlas.atlas\"\n"
  "}\n"
  ""
}
