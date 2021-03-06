package mouse

import "github.com/go-vgo/robotgo"

// Handler function to handle all mouse actions
func Handler(action Action) {

	switch action.Action {
	case "move":
		moveMouse(action.X, action.Y)
	case "click":
		clickMouse()
	}
}

// clickMouse register left click
// TODO: extend for right as well?
func clickMouse() {
	robotgo.MouseClick()
}

// moveMouse
func moveMouse(x int, y int) {
	posX, posY := robotgo.GetMousePos()
	moveX := x / 10
	moveY := y / 10

	robotgo.MoveMouseSmooth(posX+moveX, posY-moveY)

}
