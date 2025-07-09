package engine

import (
	"strings"
)

type NodeTree interface {
	GetNode(path string) Node
	GetRootNode() Node
	SetRootNode(node Node)
}

// singleton
var gameContext *GameContext = nil

type GameContext struct {
	nodeTree NodeTree
}

func (g GameContext) GetNodeTree() NodeTree {
	return g.nodeTree
}

func GetGameContext() GameContext {
	return *gameContext
}

func NewGameContext(nodeTree NodeTree) {
	if gameContext != nil {
		panic("Game context already initialized")
	}

	gameContext = &GameContext{
		nodeTree: nodeTree,
	}
}

// base node tree
type BaseNodeTree struct {
	rootNode Node
}

func (t BaseNodeTree) GetRootNode() Node {
	return t.rootNode
}

func (t BaseNodeTree) GetNode(path string) Node {
	elements := strings.FieldsFunc(path, func(r rune) bool { return r == '/' })

	currentNode := t.rootNode
	for _, element := range elements {
		currentNode = currentNode.GetChild(element)
		if currentNode == nil {
			panic("No node found at path: " + path)
		}
	}

	return currentNode
}

func (t *BaseNodeTree) SetRootNode(node Node) {
	t.rootNode = node
}

func NewBaseNodeTree() BaseNodeTree {
	return BaseNodeTree{}
}
