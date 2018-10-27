package main

import (
	"github.com/galaco/Gource-Engine/engine"
	"github.com/galaco/Gource-Engine/engine/config"
	"github.com/galaco/Gource-Engine/engine/core/event"
	"github.com/galaco/Gource-Engine/engine/core/event/message"
	"github.com/galaco/Gource-Engine/engine/core/event/message/messages"
	"github.com/galaco/Gource-Engine/engine/core/event/message/messagetype"
	"github.com/galaco/Gource-Engine/engine/loader/entity/classmap"
	"github.com/galaco/Gource-Engine/engine/renderer"
	"github.com/galaco/Gource-Engine/engine/scene"
	"github.com/galaco/Gource-Engine/engine/window"
	"github.com/galaco/Gource-Engine/entity/common"
	"github.com/go-gl/glfw/v3.2/glfw"
)

func main() {
	// Build our engine setup
	Application := engine.NewEngine()

	// Initialise current setup. Note this doesn't start any loop, but
	// allows for configuration of systems by the engine
	Application.Initialise()

	Application.AddManager(&window.Manager{})
	Application.AddManager(&renderer.Manager{})

	RegisterEntityList()

	// Load a map!
	scene.LoadFromFile(config.Get().GameDirectory + "/maps/de_dust2.bsp")

	// Register behaviour that needs to exist outside of game simulation & control
	RegisterShutdownMethod(Application)

	//Application.SetSimulationSpeed(2.5)


	// Run the engine
	Application.Run()
}

// Simple object to control engine shutdown utilising the internal event manager
type Closeable struct {
	target *engine.Engine
}

func (closer Closeable) ReceiveMessage(message message.IMessage) {
	if message.GetType() == messagetype.KeyDown {
		if message.(*messages.KeyDown).Key == glfw.KeyEscape {
			// Will shutdown the engine at the end of the current loop
			closer.target.Close()
		}
	}
}

//Implement a way of shutting down the engine
func RegisterShutdownMethod(app *engine.Engine) {
	event.GetEventManager().Listen(messagetype.KeyDown, Closeable{app})
}

func RegisterEntityList() {
	loader.RegisterClass(&common.PropDynamic{})
	loader.RegisterClass(&common.PropDynamicOverride{})
	loader.RegisterClass(&common.PropPhysics{})
	loader.RegisterClass(&common.PropPhysicsMultiplayer{})
}