package logger

import "github.com/fatih/color"

type LogItem struct {
	Level   string `json:"level"`
	Message string `json:"message"`
}

func NewLogItem(level string, message string) *LogItem {
	return &LogItem{
		Level:   level,
		Message: message,
	}
}

func Log(logItem *LogItem) {
	if logItem.Level == "ERROR" {
		color.Red(logItem.Level + ":" + logItem.Message)
		return
	}

	if logItem.Level == "INFO" {
		color.Blue(logItem.Level + ":" + logItem.Message)
		return
	}

	if logItem.Level == "DEBUG" {
		color.Yellow(logItem.Level + ":" + logItem.Message)
		return
	}

	if logItem.Level == "WARN" {
		color.Yellow(logItem.Level + ":" + logItem.Message)
		return
	}

	if logItem.Level == "SUCCESS" {
		color.Green(logItem.Level + ":" + logItem.Message)
		return
	}
}
