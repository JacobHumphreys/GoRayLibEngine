package nodes2d

import (
	"github.com/JacobHumphreys/GoRayLibEngine/scenes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Node2d struct {
	LocalPosition rl.Vector2
	LocalScale    rl.Vector2

	GlobalPosition rl.Vector2
	GlobalScale    rl.Vector2
}

func (Node2d) Init(position rl.Vector2) *Node2d {
	return &Node2d{
		LocalPosition:  position,
		GlobalPosition: rl.NewVector2(0, 0),
		LocalScale:     rl.NewVector2(1, 1),
		GlobalScale:    rl.NewVector2(1, 1),
	}
}

func (Node2d) InitWithScale(position, scale rl.Vector2) *Node2d {
	n2d := Node2d{}.Init(position)
	n2d.SetScale(scale)
	return n2d
}

func (n *Node2d) SetScale(newScale rl.Vector2) {
	n.LocalScale = newScale
}

func (n Node2d) Process(delta float32) {}
func (n Node2d) Input()                {}
func (n Node2d) Draw()                 {}
func (n Node2d) Destroy()              {}

func UpdateScenePositions(currentScene *scenes.Scene) {
	startPosition := rl.NewVector2(0, 0)

    if currentScene == nil{
        return
    }

	for _, child := range (*currentScene).GetChildrenTree().Children {
		if child == nil {
			continue
		}

		if node2d, ok := (*child.Value).(*Node2d); ok {
			node2d.GlobalPosition = rl.Vector2Add(node2d.LocalPosition, startPosition)
			startPosition = node2d.GlobalPosition
			break
		}
	}

	for _, childTree := range (*currentScene).GetChildrenTree().Children {
		if childTree == nil {
			continue
		}
		if _, ok := (*childTree.Value).(*Node2d); !ok {
			UpdateTreePositions(childTree, startPosition)
		}
	}
}

func UpdateTreePositions(tree *scenes.Tree, startPosition rl.Vector2) {
    if tree == nil{
        return
    }

	for _, childTree := range tree.Children {
		if childTree == nil {
			continue
		}
		if node2d, ok := (*childTree.Value).(*Node2d); ok {
			node2d.GlobalPosition = rl.Vector2Add(startPosition, node2d.LocalPosition)
			startPosition = node2d.GlobalPosition
		}
	}
	for _, childTree := range tree.Children {
		if childTree == nil {
			continue
		}
		if _, ok := (*childTree.Value).(*Node2d); !ok {
			UpdateTreePositions(childTree, startPosition)
		}
	}
}

func UpdateSceneScale(currentScene *scenes.Scene) {
	startScale := rl.NewVector2(1, 1)

	for _, child := range (*currentScene).GetChildrenTree().Children {
		if child == nil {
			continue
		}

		if node2d, ok := (*child.Value).(*Node2d); ok {
			node2d.GlobalScale = node2d.LocalScale
			startScale = node2d.GlobalScale
            break
		}
	}

	for _, child := range (*currentScene).GetChildrenTree().Children {
		if child == nil {
			continue
		}
		if _, ok := (*child.Value).(*Node2d); !ok {
			UpdateTreeScale(child, startScale)
		}
	}
}

func UpdateTreeScale(tree *scenes.Tree, startScale rl.Vector2) {
	for _, childTree := range tree.Children {
		if childTree == nil {
			continue
		}

		if node2d, ok := (*childTree.Value).(*Node2d); ok {
			node2d.GlobalScale = rl.Vector2Multiply(startScale, node2d.LocalScale)
			startScale = node2d.GlobalScale
		}
	}
	for _, childTree := range tree.Children {
		if childTree == nil {
			continue
		}
		if _, ok := (*childTree.Value).(*Node2d); !ok {
			UpdateTreeScale(childTree, startScale)
		}
	}
}
