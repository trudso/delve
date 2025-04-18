package engine

import (
	"strings"
)

type NodeCreator interface {
	CreateNode(typeName string, data map[string]any) Node
}

type NodeTree interface {
	GetNode(path string) Node
	GetRootNode() Node
	SetRootNode(node Node)
}

// singleton
var gameContext *GameContext = nil

type GameContext struct {
	nodeCreator NodeCreator
	nodeTree    NodeTree
}

func (g GameContext) GetNodeCreator() NodeCreator {
	return g.nodeCreator
}

func (g GameContext) GetNodeTree() NodeTree {
	return g.nodeTree
}

func GetGameContext() GameContext {
	return *gameContext
}

func NewGameContext(nodeCreator NodeCreator, nodeTree NodeTree) {
	if gameContext != nil {
		panic("Game context already initialized")
	}

	gameContext = &GameContext{
		nodeCreator: nodeCreator,
		nodeTree:    nodeTree,
	}
}

// base node creator
type NodeInstantiator func(data map[string]any) Node

type BaseNodeCreator struct {
	nodeInstantiators map[string]NodeInstantiator
}

func (c BaseNodeCreator) CreateNode(typeName string, data map[string]any) Node {
	return c.nodeInstantiators[typeName](data)
}

func (c *BaseNodeCreator) Register(typeName string, instantiator NodeInstantiator) {
	c.nodeInstantiators[typeName] = instantiator
}

func NewBaseNodeCreator() BaseNodeCreator {
	return BaseNodeCreator{
		nodeInstantiators: map[string]NodeInstantiator{
			SPRITE_NODE: NewSpriteFromDataSet,
		},
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
