package logic

import (
	"github.com/wirekang/mouseable/internal/typ"
)

type cmdLogic struct {
	onBegin func(s *logicState)
	onStep  func(s *logicState)
	onEnd   func(s *logicState)
}

var nop = func(s *logicState) {}

var cmdLogicMap = map[typ.CommandName]cmdLogic{
	"activate": {
		onBegin: func(s *logicState) {
			s.when = typ.Activated
		},
	},
	"deactivate": {
		onBegin: func(s *logicState) {
			s.when = typ.Deactivated
		},
	},
	"click-left": {
		onBegin: func(s *logicState) {
			go s.hookManager.MouseDown(typ.Left)
		},
		onEnd: func(s *logicState) {
			go s.hookManager.MouseUp(typ.Left)
		},
	},
}
