package nodes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Node interface {
	GetId() string
	GetParent() Node
	SetParent(node Node)
	AddChild(node Node)
	GetChild(id string) Node
	GetChildren() []Node
	GetTransform() Transform
	Move(deltaTime float32)
	Draw()
	Close()
}

type BaseNode struct {
	Id string
	Parent Node
	Children []Node
	Transform Transform
}

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

func (n *BaseNode) Move(deltaTime float32) {}

func (n BaseNode) Draw() {}

func (n BaseNode) Close() {
	for _, child := range n.Children {
		Close(child)
	}
}

func NewBaseNode(id string) BaseNode {
	return BaseNode{
		Id: id,
		Parent: nil,
		Children: make([]Node, 0),
		Transform: NewTransform(),
	}
}

func Update(n Node, deltaTime float32) {
	n.Move(deltaTime)
	rl.PushMatrix()
	transform := n.GetTransform()
	rl.Translatef(transform.Position.X, transform.Position.Y, 0)
	rl.Rotatef(transform.Rotation.X, transform.Rotation.Y, 0, 1)		
	rl.Scalef(transform.Scale.X, transform.Scale.Y, 1)

	n.Draw()

	for _, child := range n.GetChildren() {
		Update(child, deltaTime)
	}

	rl.PopMatrix()
}

func Close(n Node) {
	n.Close()
}
