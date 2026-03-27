components {
  id: "portal1"
  component: "/game/main/portal/portal.script"
}
components {
  id: "portalback_copy"
  component: "/game/main/portal/portalback_copy.collisionobject"
}
components {
  id: "portal_copy"
  component: "/game/main/portal/portal_copy.collisionobject"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"portal2\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/game/images/atlas.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.557181
  }
}
