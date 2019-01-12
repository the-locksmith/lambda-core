package renderer

import (
	"github.com/galaco/Lambda-Core/client/scene/world"
	"github.com/galaco/Lambda-Core/core/entity"
	"github.com/galaco/Lambda-Core/core/model"
	"github.com/go-gl/mathgl/mgl32"
)

type IRenderer interface {
	StartFrame(*entity.Camera)
	LoadShaders()
	DrawBsp(*world.World)
	DrawSkybox(*world.Sky)
	DrawModel(*model.Model, mgl32.Mat4)
	DrawSkyMaterial(*model.Model)
	SetWireframeMode(bool)
	EndFrame()
	Unregister()
}
