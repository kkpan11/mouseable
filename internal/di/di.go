package di

import (
	"github.com/wirekang/mouseable/internal/data"
	"github.com/wirekang/mouseable/internal/def"
	"github.com/wirekang/mouseable/internal/hook"
	"github.com/wirekang/mouseable/internal/logic"
	"github.com/wirekang/mouseable/internal/view"
)

func Init() {
	logic.DI.SetCursorPos = hook.SetCursorPos
	logic.DI.GetCursorPos = hook.GetCursorPos
	logic.DI.AddCursorPos = hook.AddCursorPos
	logic.DI.MouseDown = hook.MouseDown
	logic.DI.MouseUp = hook.MouseUp
	logic.DI.Wheel = hook.Wheel
	hook.DI.OnKey = logic.OnKey
	hook.DI.AlertError = view.AlertError
	hook.DI.OnUnhook = logic.StopAllFunction
	view.DI.LoadConfig = data.LoadConfig
	view.DI.SaveConfig = data.SaveConfig
	data.DI.SetConfig = func(config def.Config) {
		hook.SetKey(config.ActivateKey, config.DeactivateKey)
		logic.SetConfig(config)
	}
	data.DI.AlertError = view.AlertError
}
