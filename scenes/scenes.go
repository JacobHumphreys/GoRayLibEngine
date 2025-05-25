package scenes

import (
	"slices"

	"github.com/JacobHumphreys/GoRayLibEngine/nodes"
)

type Scene interface {
	nodes.Node
	GetChildrenTree() Hierarchy
}

func AsScenePtr[T interface {
	*U
	Scene
}, U any](n T) *Scene {
	var scene Scene = n
	return &scene
}

type Hierarchy struct {
	Scene    *Scene
	Children []*Tree
}

type Tree struct {
	Value    *nodes.Node
	Children []*Tree
}

func (h *Hierarchy) RemoveNode(n *nodes.Node) bool {
	for i, c := range h.Children {
		if c.Value == n {
			h.Children = slices.Delete(h.Children, i, i+1)
			return true
		}
		for _, t := range c.Children {
			if t.RemoveNode(n) {
				return true
			}
		}
	}
	return false
}

func (t *Tree) RemoveNode(n *nodes.Node) bool {
	for i, c := range t.Children {
		if c.Value == n {
			c.Children = slices.Delete(c.Children, i, i+1)
			return true
		}
		if c.RemoveNode(n) {
			return true
		}
	}
	return false
}

func (h Hierarchy) DoOnEveryScene(function func(*Scene)) {
	function(h.Scene)
	for _, c := range h.Children {
		if c == nil {
			continue
		}

		c.DoOnEveryScene(function)
	}

}

func (t Tree) DoOnEveryScene(function func(*Scene)) {
	if s, ok := (*t.Value).(Scene); ok {
		function(&s)
	}
	for _, c := range t.Children {
		if c == nil {
			continue
		}

		c.DoOnEveryScene(function)
	}
}

func (h Hierarchy) DoOnEveryNode(function func(*nodes.Node)) {
	var node nodes.Node = *h.Scene
	function(&node)
	for _, c := range h.Children {
		if c == nil {
			continue
		}

		c.DoOnEveryNode(function)
	}

}

func (t Tree) DoOnEveryNode(function func(*nodes.Node)) {
	function(t.Value)
	for _, c := range t.Children {
		if c == nil {
			continue
		}

		c.DoOnEveryNode(function)
	}
}
