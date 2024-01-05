package service

import "local-build/internal/store/model"

// export init tools
func InitTools() {
	initTools()
}

// export tool list
func ToolList() []model.Tool {
	return toolList()
}