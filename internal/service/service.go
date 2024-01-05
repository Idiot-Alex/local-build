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

// export save tool
func SaveTool() bool {
	return true
}

// export del tool
func DelTool(ids []string) bool {
	return delTool(ids)
}