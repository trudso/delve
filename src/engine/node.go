package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
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

	// Serialization functions
	// GetDataSet(onlyChangedFields bool) map[string]any
	// ApplyDataSet(map[string]any)

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
//  for change-sets. This makes the replication creation point very important when
//  building change-sets.
func (n *BaseNode) GetReplication() ReplicationCollection {
	children := NewReplicationCollection( "children", []Replicatable {})
	for _, child := range n.Children {
		children.AddElement(child.GetReplication())
	}

	replication := NewReplicationCollection( n.Id, []Replicatable {
		NewReplicationPrimitive( "id", &n.Id, true, nil),	
		NewReplicationPrimitive( "type", &n.nodeType, true, nil),
		n.Transform.GetReplication(),
		children,	
	})

	return replication
}  

// TODO[mt]: rework to use replication instead
//func (n BaseNode) GetDataSet(onlyChangedFields bool) map[string]any {
//	children := make(map[string]any)
//	for _, child := range n.Children {
//		children[child.GetId()] = NodeToDataSet(child, onlyChangedFields)
//	}
//
//	res := map[string]any{
//		"id":        n.Id,
//		"type":      n.nodeType,
//		"transform": n.Transform.GetDataSet(onlyChangedFields),
//		"children":  children,
//	}
//
//	return res
//}

// TODO[mt]: rework to use replication instead
//func (n *BaseNode) ApplyDataSet(data map[string]any) {
//	if data["id"] != nil {
//		n.Id = data["id"].(string)
//	}
//
//	if d, found := data["transform"]; found {
//		n.Transform.ApplyDataSet(d.(map[string]any))
//	}
//
//	if d, found := data["children"]; found {
//		children := d.(map[string]any)
//		for key, childData := range children {
//			// TODO[mt]: add support for deletion and for node tree mutation
//			// modify existing child
//			existingChild := n.GetChild(key)
//			if existingChild != nil {
//				existingChild.ApplyDataSet(childData.(map[string]any))
//			} else {
//				newChild := CreateNodeFromDataSet(childData.(map[string]any))
//				n.AddChild(newChild)
//			}
//		}
//	}
//}

// constructors
func NewBaseNode(nodeType string, id string) BaseNode {
	return BaseNode{
		Id:        id,
		nodeType:  nodeType,
		Parent:    nil,
		Children:  make([]Node, 0),
		Transform: NewTransform(),
	}
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

//func NodeToDataSet(n Node, onlyChangedFields bool) map[string]any {
//	return n.GetDataSet(onlyChangedFields)
//}
//
//func DataSetToNode(data map[string]any) Node {
//	return CreateNodeFromDataSet(data)
//}
//
//func ApplyDataSet(n Node, data map[string]any) Node {
//	n.ApplyDataSet(data)
//	return n
//}
//
//func CreateNodeFromDataSet(data map[string]any) Node {
//	return GetGameContext().GetNodeCreator().CreateNode(data["type"].(string), data)
//}

