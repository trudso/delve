package engine

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type NodeFactory func(id string, data map[string]any) Node

var (
	nodeFactories = map[string]NodeFactory{}
)

type Node interface {
	// tree functions
	GetId() string // unique name of the node ( by parent )
	GetParent() Node
	SetParent(node Node)
	AddChild(node Node)
	GetChild(id string) Node
	GetChildren() []Node
	GetTransform() Transform

	// Update functions
	// Input: mostly to support nodes where the input does not depend on deltaTime
	//  or for splitting input and tranformation logic, if that is something that rocks your boat.
	Input()

	// Move: update local transform based on deltaTime
	// 	This is a good place to check for movement based input
	Move(deltaTime float32)

	// Draw: draw the node
	Draw()

	// Initialization and cleanup
	Init()
	Delete()

	GetReplication() ReplicationCollection
}

// --------------- Base node implementation ---------------
type BaseNode struct {
	Id        string
	nodeType  string
	Parent    Node
	Children  []Node
	Transform Transform
}

// Tree functions
func (n BaseNode) GetId() string {
	return n.Id
}

func (n BaseNode) GetParent() Node {
	return n.Parent
}

func (n *BaseNode) SetParent(node Node) {
	n.Parent = node
}

func (n *BaseNode) AddChild(node Node) {
	node.SetParent(n)
	n.Children = append(n.Children, node)
}

func (n BaseNode) GetChild(id string) Node {
	for _, child := range n.Children {
		if child.GetId() == id {
			return child
		}
	}
	return nil
}

func (n BaseNode) GetChildren() []Node {
	return n.Children
}

func (n BaseNode) GetTransform() Transform {
	return n.Transform
}

// Behavioural overridables
func (n *BaseNode) Input() {}

func (n *BaseNode) Move(deltaTime float32) {}

func (n BaseNode) Draw() {}

// Initialization and cleanup
func (n *BaseNode) Init() {
	for _, child := range n.Children {
		InitNode(child)
	}
}

func (n BaseNode) Delete() {
	for _, child := range n.Children {
		DeleteNode(child)
	}
}

// Important: Replication base values determined at creation the replication creation time
//
//	for change-sets. This makes the replication creation point very important when
//	building change-sets.
func (n *BaseNode) GetReplication() ReplicationCollection {
	children := NewReplicationCollection("children", []Replicatable{}, nil, n.addNewChildFromDataSet)
	for _, child := range n.Children {
		children.AddElement(child.GetReplication())
	}

	replication := NewReplicationCollection(n.Id, []Replicatable{
		n.Transform.GetReplication(),
		children,
	}, map[string]string{
		"id":   n.Id,
		"type": n.nodeType,
	}, n.addNewChildFromDataSet)

	return replication
}

func (n *BaseNode) addNewChildFromDataSet(id string, data map[string]any) Replicatable {
	typeName, found := data["type"]
	if !found {
		return nil
	}

	factory, found := nodeFactories[typeName.(string)]
	if !found {
		fmt.Fprintln(os.Stderr, "No factory found for:", typeName)
		return nil
	}

	node := factory(id, data)
	n.AddChild(node)
	return node.GetReplication()
}

// constructors
func NewBaseNode(id string, nodeType reflect.Type, factory NodeFactory ) BaseNode {
	node := BaseNode{
		Id:        id,
		nodeType:  strings.ReplaceAll(nodeType.String(), "*", ""), // remove any pointer annotations
		Parent:    nil,
		Children:  make([]Node, 0),
		Transform: NewTransform(),
	}

	if _, found := nodeFactories[nodeType.String()]; !found && factory != nil {
		nodeFactories[node.nodeType] = factory
	}

	return node
}

// Game loop functionality
func Update(n Node, deltaTime float32) {
	// update movement
	n.Input()
	n.Move(deltaTime)

	// apply transforms
	rl.PushMatrix()
	transform := n.GetTransform()
	rl.Translatef(transform.Position.X, transform.Position.Y, 0)
	rl.Rotatef(transform.Rotation.X, transform.Rotation.Y, 0, 1)
	rl.Scalef(transform.Scale.X, transform.Scale.Y, 1)

	// draw node
	n.Draw()

	// update children
	for _, child := range n.GetChildren() {
		Update(child, deltaTime)
	}

	// pop transforms
	rl.PopMatrix()
}

// --------------- Node functions ---------------
func InitNode(n Node) {
	n.Init()
}

func DeleteNode(n Node) {
	n.Delete()
}
