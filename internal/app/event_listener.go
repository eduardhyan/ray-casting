package app

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type EventListener struct {
	Listeners []func()
}

func (l *EventListener) Add(key ebiten.Key, f func()) {
	l.Listeners = append(l.Listeners, CreateListener(key, f))
}

func (l *EventListener) Handle() {
	for _, listener := range l.Listeners {
		listener()
	}
}

func CreateListener(key ebiten.Key, handler func()) func() {
	hold := false

	return func() {
		if inpututil.IsKeyJustPressed(key) {
			hold = true
		}

		if inpututil.IsKeyJustReleased(key) {
			hold = false
		}

		if hold {
			handler()
		}
	}
}
