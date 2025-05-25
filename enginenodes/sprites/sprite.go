package sprites

import (
	"github.com/Tigy01/GoRayLibEngine/enginenodes/nodes2d"
	"github.com/Tigy01/GoRayLibEngine/nodes"
	"github.com/Tigy01/GoRayLibEngine/scenes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	*nodes2d.Node2d
	Texture rl.Texture2D
	Hidden  bool
	FlipH   bool
}

func (Sprite) Init(path string, node2d *nodes2d.Node2d) *Sprite {
	return &Sprite{
		Node2d:  node2d,
		Texture: rl.LoadTexture(path),
		Hidden:  false,
		FlipH:   false,
	}
}

func (s *Sprite) Destroy() {
	rl.UnloadTexture(s.Texture)
}

func (s Sprite) Process(delta float32) {}
func (s Sprite) Input()                {}

func (s *Sprite) GetChildrenTree() scenes.Hierarchy {
	return scenes.Hierarchy{
		Scene: scenes.AsScenePtr(s),
		Children: []*scenes.Tree{
			{
				Value: nodes.AsNodePtr(s.Node2d),
			},
		},
	}
}

func (s Sprite) Draw() {
	if s.Hidden {
		return
	}

	rl.DrawTexturePro(
		s.Texture,
		s.getSourceRec(),
		s.getDestRec(),
		rl.NewVector2(0, 0),
		0,
		rl.White,
	)
}

func (s Sprite) getDestRec() rl.Rectangle {
	return rl.NewRectangle(
		s.Node2d.GlobalPosition.X,
		s.Node2d.GlobalPosition.Y,
		float32(s.Texture.Width)*s.LocalScale.X,
		float32(s.Texture.Height)*s.LocalScale.Y,
	)
}

func (s Sprite) getSourceRec() rl.Rectangle {
	return rl.NewRectangle(0, 0, float32(s.Texture.Width), float32(s.Texture.Height))
}

func (s *Sprite) Center() {
	s.Offset(
		rl.NewVector2(
			(0 - float32(s.Texture.Width)/2),
			(0 - float32(s.Texture.Height)/2),
		),
	)
}

func (s *Sprite) Offset(offset rl.Vector2) {
	s.Node2d.LocalPosition = rl.Vector2Add(s.Node2d.LocalPosition, offset)
}
