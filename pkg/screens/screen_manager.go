package screens

type ScreenInterface interface {
	Process()
}

type ScreenManager struct {
	CurrentScreen ScreenInterface
	screens       []ScreenInterface
}

func NewScreenManager() *ScreenManager {
	screenManager := &ScreenManager{}

	screenManager.screens = make([]ScreenInterface, 2)
	screenManager.screens = append(screenManager.screens, NewStartScreen(screenManager))
	screenManager.CurrentScreen = screenManager.screens[0]

	return screenManager
}
