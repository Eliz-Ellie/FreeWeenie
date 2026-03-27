components {
  id: "portal"
  component: "/game/main/portal/portal.collisionobject"
}
components {
  id: "portal1"
  component: "/game/main/portal/portal.script"
}
components {
  id: "wall"
  component: "/game/main/portal/portalback.collisionobject"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"portal1\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/game/images/atlas.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 0.458643
  }
}
