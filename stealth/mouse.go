package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// MoveMouseHuman simulates human-like mouse movement using JS function eval
func MoveMouseHuman(page *rod.Page, targetX, targetY int) {
	rand.Seed(time.Now().UnixNano())

	startX := rand.Intn(200)
	startY := rand.Intn(200)
	steps := rand.Intn(15) + 20 // 20–35 steps

	for i := 0; i <= steps; i++ {
		x := startX + (targetX-startX)*i/steps
		y := startY + (targetY-startY)*i/steps

		page.MustEval(`() => {
			const evt = new MouseEvent('mousemove', {
				clientX: arguments[0],
				clientY: arguments[1],
				bubbles: true
			});
			document.body.dispatchEvent(evt);
		}`, x, y)

		time.Sleep(time.Duration(rand.Intn(40)+20) * time.Millisecond)
	}
}
