package app

import (
	"github.com/bvisness/SQLJam/node"
	"github.com/bvisness/SQLJam/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func doJoinUpdate(n *node.Node, j *node.Join) {
	n.InputPinHeights = make([]int, len(n.Inputs))

	uiHeight := UIFieldHeight // blank space for first table input
	n.InputPinHeights[0] = 0

	for i := range n.Inputs[1:] {
		uiHeight += UIFieldSpacing
		n.InputPinHeights[i+1] = uiHeight
		uiHeight += UIFieldHeight
	}

	uiHeight += UIFieldSpacing + UIFieldHeight // +/- buttons

	n.UISize = rl.Vector2{300, float32(uiHeight)}
}

func doJoinUI(n *node.Node, j *node.Join) {
	fieldY := n.UIRect.Y + UIFieldHeight + UIFieldSpacing

	uiRight := n.UIRect.X + n.UIRect.Width
	boxWidth := n.UIRect.Width - (UIFieldSpacing+UIFieldHeight)*2

	for _, condition := range j.Conditions {
		boxRect := rl.Rectangle{
			n.UIRect.X,
			float32(fieldY),
			boxWidth,
			UIFieldHeight,
		}
		rl.DrawRectangleRec(boxRect, rl.White)
		condition.Condition = condition.TextBox.Do(boxRect, condition.Condition, 100)
		condition.Left = raygui.Toggle(rl.Rectangle{
			uiRight - (UIFieldHeight + UIFieldSpacing + UIFieldHeight),
			float32(fieldY),
			UIFieldHeight,
			UIFieldHeight,
		}, "L", condition.Left)
		condition.Right = raygui.Toggle(rl.Rectangle{
			uiRight - UIFieldHeight,
			float32(fieldY),
			UIFieldHeight,
			UIFieldHeight,
		}, "R", condition.Right)

		fieldY += UIFieldHeight + UIFieldSpacing
	}

	if raygui.Button(rl.Rectangle{
		n.UIRect.X,
		fieldY,
		n.UIRect.Width/2 - UIFieldSpacing/2,
		UIFieldHeight,
	}, "+") {
		n.Inputs = append(n.Inputs, nil)
		j.Conditions = append(j.Conditions, &node.JoinCondition{})
	}
	if raygui.Button(rl.Rectangle{
		n.UIRect.X + n.UIRect.Width/2 + UIFieldSpacing/2,
		fieldY,
		n.UIRect.Width/2 - UIFieldSpacing/2,
		UIFieldHeight,
	}, "-") {
		if len(j.Conditions) > 1 {
			n.Inputs = n.Inputs[:len(n.Inputs)-1]
			j.Conditions = j.Conditions[:len(j.Conditions)-1]
		}
	}
}