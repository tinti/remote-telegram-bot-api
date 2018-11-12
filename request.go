package rbot

import (
	"fmt"
	"net/url"
	"reflect"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type ConcreteChattable struct {
	Type string

	ValueMessageConfig                tgbotapi.MessageConfig
	ValueForwardConfig                tgbotapi.ForwardConfig
	ValuePhotoConfig                  tgbotapi.PhotoConfig
	ValueAudioConfig                  tgbotapi.AudioConfig
	ValueDocumentConfig               tgbotapi.DocumentConfig
	ValueStickerConfig                tgbotapi.StickerConfig
	ValueVideoConfig                  tgbotapi.VideoConfig
	ValueAnimationConfig              tgbotapi.AnimationConfig
	ValueVideoNoteConfig              tgbotapi.VideoNoteConfig
	ValueVoiceConfig                  tgbotapi.VoiceConfig
	ValueMediaGroupConfig             tgbotapi.MediaGroupConfig
	ValueLocationConfig               tgbotapi.LocationConfig
	ValueVenueConfig                  tgbotapi.VenueConfig
	ValueContactConfig                tgbotapi.ContactConfig
	ValueGameConfig                   tgbotapi.GameConfig
	ValueSetGameScoreConfig           tgbotapi.SetGameScoreConfig
	ValueGetGameHighScoresConfig      tgbotapi.GetGameHighScoresConfig
	ValueChatActionConfig             tgbotapi.ChatActionConfig
	ValueEditMessageTextConfig        tgbotapi.EditMessageTextConfig
	ValueEditMessageCaptionConfig     tgbotapi.EditMessageCaptionConfig
	ValueEditMessageReplyMarkupConfig tgbotapi.EditMessageReplyMarkupConfig
	ValueInvoiceConfig                tgbotapi.InvoiceConfig
	ValueDeleteMessageConfig          tgbotapi.DeleteMessageConfig
	ValuePinChatMessageConfig         tgbotapi.PinChatMessageConfig
	ValueUnpinChatMessageConfig       tgbotapi.UnpinChatMessageConfig
	ValueSetChatTitleConfig           tgbotapi.SetChatTitleConfig
	ValueSetChatDescriptionConfig     tgbotapi.SetChatDescriptionConfig
	ValueDeleteChatPhotoConfig        tgbotapi.DeleteChatPhotoConfig
	//ValueSetChatPhotoConfig         tgbotapi.SetChatPhotoConfig
}

func NewConcreteChattable(c tgbotapi.Chattable) ConcreteChattable {
	cC := ConcreteChattable{
		Type: reflect.TypeOf(c).String(),
	}

	switch cC.Type {
	case reflect.TypeOf(tgbotapi.MessageConfig{}).String():
		cC.ValueMessageConfig = c.(tgbotapi.MessageConfig)
	case reflect.TypeOf(tgbotapi.ForwardConfig{}).String():
		cC.ValueForwardConfig = c.(tgbotapi.ForwardConfig)
	case reflect.TypeOf(tgbotapi.PhotoConfig{}).String():
		cC.ValuePhotoConfig = c.(tgbotapi.PhotoConfig)
	case reflect.TypeOf(tgbotapi.AudioConfig{}).String():
		cC.ValueAudioConfig = c.(tgbotapi.AudioConfig)
	case reflect.TypeOf(tgbotapi.DocumentConfig{}).String():
		cC.ValueDocumentConfig = c.(tgbotapi.DocumentConfig)
	case reflect.TypeOf(tgbotapi.StickerConfig{}).String():
		cC.ValueStickerConfig = c.(tgbotapi.StickerConfig)
	case reflect.TypeOf(tgbotapi.VideoConfig{}).String():
		cC.ValueVideoConfig = c.(tgbotapi.VideoConfig)
	case reflect.TypeOf(tgbotapi.AnimationConfig{}).String():
		cC.ValueAnimationConfig = c.(tgbotapi.AnimationConfig)
	case reflect.TypeOf(tgbotapi.VideoNoteConfig{}).String():
		cC.ValueVideoNoteConfig = c.(tgbotapi.VideoNoteConfig)
	case reflect.TypeOf(tgbotapi.VoiceConfig{}).String():
		cC.ValueVoiceConfig = c.(tgbotapi.VoiceConfig)
	case reflect.TypeOf(tgbotapi.MediaGroupConfig{}).String():
		cC.ValueMediaGroupConfig = c.(tgbotapi.MediaGroupConfig)
	case reflect.TypeOf(tgbotapi.LocationConfig{}).String():
		cC.ValueLocationConfig = c.(tgbotapi.LocationConfig)
	case reflect.TypeOf(tgbotapi.VenueConfig{}).String():
		cC.ValueVenueConfig = c.(tgbotapi.VenueConfig)
	case reflect.TypeOf(tgbotapi.ContactConfig{}).String():
		cC.ValueContactConfig = c.(tgbotapi.ContactConfig)
	case reflect.TypeOf(tgbotapi.GameConfig{}).String():
		cC.ValueGameConfig = c.(tgbotapi.GameConfig)
	case reflect.TypeOf(tgbotapi.SetGameScoreConfig{}).String():
		cC.ValueSetGameScoreConfig = c.(tgbotapi.SetGameScoreConfig)
	case reflect.TypeOf(tgbotapi.GetGameHighScoresConfig{}).String():
		cC.ValueGetGameHighScoresConfig = c.(tgbotapi.GetGameHighScoresConfig)
	case reflect.TypeOf(tgbotapi.ChatActionConfig{}).String():
		cC.ValueChatActionConfig = c.(tgbotapi.ChatActionConfig)
	case reflect.TypeOf(tgbotapi.EditMessageTextConfig{}).String():
		cC.ValueEditMessageTextConfig = c.(tgbotapi.EditMessageTextConfig)
	case reflect.TypeOf(tgbotapi.EditMessageCaptionConfig{}).String():
		cC.ValueEditMessageCaptionConfig = c.(tgbotapi.EditMessageCaptionConfig)
	case reflect.TypeOf(tgbotapi.EditMessageReplyMarkupConfig{}).String():
		cC.ValueEditMessageReplyMarkupConfig = c.(tgbotapi.EditMessageReplyMarkupConfig)
	case reflect.TypeOf(tgbotapi.InvoiceConfig{}).String():
		cC.ValueInvoiceConfig = c.(tgbotapi.InvoiceConfig)
	case reflect.TypeOf(tgbotapi.DeleteMessageConfig{}).String():
		cC.ValueDeleteMessageConfig = c.(tgbotapi.DeleteMessageConfig)
	case reflect.TypeOf(tgbotapi.PinChatMessageConfig{}).String():
		cC.ValuePinChatMessageConfig = c.(tgbotapi.PinChatMessageConfig)
	case reflect.TypeOf(tgbotapi.UnpinChatMessageConfig{}).String():
		cC.ValueUnpinChatMessageConfig = c.(tgbotapi.UnpinChatMessageConfig)
	case reflect.TypeOf(tgbotapi.SetChatTitleConfig{}).String():
		cC.ValueSetChatTitleConfig = c.(tgbotapi.SetChatTitleConfig)
	case reflect.TypeOf(tgbotapi.SetChatDescriptionConfig{}).String():
		cC.ValueSetChatDescriptionConfig = c.(tgbotapi.SetChatDescriptionConfig)
	case reflect.TypeOf(tgbotapi.DeleteChatPhotoConfig{}).String():
		cC.ValueDeleteChatPhotoConfig = c.(tgbotapi.DeleteChatPhotoConfig)
	//case reflect.TypeOf(tgbotapi.SetChatPhotoConfig{}).String():
	//	cC.ValueSetChatPhotoConfig = c.(tgbotapi.SetChatPhotoConfig)
	default:
		panic(fmt.Errorf("can't create ConcreteChattable from Chattable"))
	}

	return cC
}

func (c *ConcreteChattable) ToChattable() tgbotapi.Chattable {
	var m tgbotapi.Chattable

	switch c.Type {
	case reflect.TypeOf(tgbotapi.MessageConfig{}).String():
		m = c.ValueMessageConfig
	case reflect.TypeOf(tgbotapi.ForwardConfig{}).String():
		m = c.ValueForwardConfig
	case reflect.TypeOf(tgbotapi.PhotoConfig{}).String():
		m = c.ValuePhotoConfig
	case reflect.TypeOf(tgbotapi.AudioConfig{}).String():
		m = c.ValueAudioConfig
	case reflect.TypeOf(tgbotapi.DocumentConfig{}).String():
		m = c.ValueDocumentConfig
	case reflect.TypeOf(tgbotapi.StickerConfig{}).String():
		m = c.ValueStickerConfig
	case reflect.TypeOf(tgbotapi.VideoConfig{}).String():
		m = c.ValueVideoConfig
	case reflect.TypeOf(tgbotapi.AnimationConfig{}).String():
		m = c.ValueAnimationConfig
	case reflect.TypeOf(tgbotapi.VideoNoteConfig{}).String():
		m = c.ValueVideoNoteConfig
	case reflect.TypeOf(tgbotapi.VoiceConfig{}).String():
		m = c.ValueVoiceConfig
	case reflect.TypeOf(tgbotapi.MediaGroupConfig{}).String():
		m = c.ValueMediaGroupConfig
	case reflect.TypeOf(tgbotapi.LocationConfig{}).String():
		m = c.ValueLocationConfig
	case reflect.TypeOf(tgbotapi.VenueConfig{}).String():
		m = c.ValueVenueConfig
	case reflect.TypeOf(tgbotapi.ContactConfig{}).String():
		m = c.ValueContactConfig
	case reflect.TypeOf(tgbotapi.GameConfig{}).String():
		m = c.ValueGameConfig
	case reflect.TypeOf(tgbotapi.SetGameScoreConfig{}).String():
		m = c.ValueSetGameScoreConfig
	case reflect.TypeOf(tgbotapi.GetGameHighScoresConfig{}).String():
		m = c.ValueGetGameHighScoresConfig
	case reflect.TypeOf(tgbotapi.ChatActionConfig{}).String():
		m = c.ValueChatActionConfig
	case reflect.TypeOf(tgbotapi.EditMessageTextConfig{}).String():
		m = c.ValueEditMessageTextConfig
	case reflect.TypeOf(tgbotapi.EditMessageCaptionConfig{}).String():
		m = c.ValueEditMessageCaptionConfig
	case reflect.TypeOf(tgbotapi.EditMessageReplyMarkupConfig{}).String():
		m = c.ValueEditMessageReplyMarkupConfig
	case reflect.TypeOf(tgbotapi.InvoiceConfig{}).String():
		m = c.ValueInvoiceConfig
	case reflect.TypeOf(tgbotapi.DeleteMessageConfig{}).String():
		m = c.ValueDeleteMessageConfig
	case reflect.TypeOf(tgbotapi.PinChatMessageConfig{}).String():
		m = c.ValuePinChatMessageConfig
	case reflect.TypeOf(tgbotapi.UnpinChatMessageConfig{}).String():
		m = c.ValueUnpinChatMessageConfig
	case reflect.TypeOf(tgbotapi.SetChatTitleConfig{}).String():
		m = c.ValueSetChatTitleConfig
	case reflect.TypeOf(tgbotapi.SetChatDescriptionConfig{}).String():
		m = c.ValueSetChatDescriptionConfig
	case reflect.TypeOf(tgbotapi.DeleteChatPhotoConfig{}).String():
		m = c.ValueDeleteChatPhotoConfig
	//case reflect.TypeOf(tgbotapi.SetChatPhotoConfig{}).String():
	//	m = c.ValueSetChatPhotoConfig
	default:
		panic(fmt.Errorf("can't create Chattable from ConcreteChattable"))
	}

	return m
}

type RequestMessage struct {
	Operation     string
	CorrelationId string

	C         ConcreteChattable
	Config    tgbotapi.UserProfilePhotosConfig
	Config2   tgbotapi.FileConfig
	Config3   tgbotapi.UpdateConfig
	Config4   tgbotapi.WebhookConfig
	Config5   tgbotapi.InlineConfig
	Config6   tgbotapi.CallbackConfig
	Config7   tgbotapi.KickChatMemberConfig
	Config8   tgbotapi.ChatConfig
	Config9   tgbotapi.ChatConfigWithUser
	Config10  tgbotapi.ChatMemberConfig
	Config11  tgbotapi.RestrictChatMemberConfig
	Config12  tgbotapi.PromoteChatMemberConfig
	Config13  tgbotapi.GetGameHighScoresConfig
	Config14  tgbotapi.ShippingConfig
	Config15  tgbotapi.PreCheckoutConfig
	Config16  tgbotapi.DeleteMessageConfig
	Config17  tgbotapi.PinChatMessageConfig
	Config18  tgbotapi.UnpinChatMessageConfig
	Config19  tgbotapi.SetChatTitleConfig
	Config20  tgbotapi.SetChatDescriptionConfig
	Config21  tgbotapi.SetChatPhotoConfig
	Config22  tgbotapi.DeleteChatPhotoConfig
	Endpoint  string
	Fieldname string
	File      interface{}
	FileID    string
	Message   tgbotapi.Message
	Params    url.Values
	Params2   map[string]string
	Pattern   string
}
