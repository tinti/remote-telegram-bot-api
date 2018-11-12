package rbot

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type ConcreteError struct {
	IsNil bool
	Value string
}

func (e *ConcreteError) ToError() error {
	if e.IsNil {
		return nil
	}

	return &ErrorString{e.Value}
}

func NewConcreteError(e error) ConcreteError {
	if e != nil {
		return ConcreteError{false, e.Error()}
	}

	return ConcreteError{true, ""}
}

type ResponseMessage struct {
	Operation     string
	CorrelationId string

	R   tgbotapi.APIResponse
	R2  ConcreteError
	R3  string
	R4  tgbotapi.User
	R5  bool
	R6  tgbotapi.Message
	R7  tgbotapi.UserProfilePhotos
	R8  tgbotapi.File
	R9  []tgbotapi.Update
	R10 tgbotapi.WebhookInfo
	//R11 tgbotapi.UpdatesChannel
	R12 tgbotapi.Chat
	R13 []tgbotapi.ChatMember
	R14 int
	R15 tgbotapi.ChatMember
	R16 []tgbotapi.GameHighScore
}
