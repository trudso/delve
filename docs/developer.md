# Delve development

## State centric development

The state of the game should always be in the state that can be persisted and restored. This means that no state is outside the central state model.

## Fast iterations

To ensure fast iterations on the code base, we use air for live reloading on save.
This means that in order to keep the current place when live reloading the client needs to be able to load a snapshot of the state.
The best solution is probably the ability to save/load snapshots directly in the client. This will also be a good solution for replicating game states in bug reports.

## Architecture

* State based
* Scene based

type Transform struct {
    Position Vector2
    Scale Vector2
}

type Scene struct {
    SceneId string
    Tranform Transform
    Children []Scene 
}

type GameState struct {
    CurrentScene *Scene
    CurrentSceneId string
    Scenes []Scene
}

// elements that should be run on each frame
type Updatable interface {
    Update()
}

## Folder layout

src/scenes/ -> all individual scenes: players, bullets, monsters, items ...
src/levels/ -> ie. town, dungeon_level_1, ... 
