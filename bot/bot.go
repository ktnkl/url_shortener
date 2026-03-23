package bot

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
	"url-shortener/database"
	"url-shortener/validation"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

func InitTgBot() {
	go func() {
		maxRetries := 5
		retryDelay := 5 * time.Second

		for i := range maxRetries {
			err := startBot()
			if err != nil {
				log.Printf("Cannot run telegram bot. Attempt %d \n", i+1)
				return
			}

			if i < maxRetries-1 {
				time.Sleep(retryDelay)
				retryDelay *= 2
			}
		}

		log.Println("Unable to start telegram bot: max retries: max retries count exceeded")
	}()
}

func startBot() error {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		// bot.WithDefaultHandler(defaultHandler),
	}

	b, err := bot.New(os.Getenv("TG_BOT_API_KEY"), opts...)

	if nil != err {
		return err
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "/start", bot.MatchTypeCommandStartOnly, defaultHandler)
	b.RegisterHandler(bot.HandlerTypeMessageText, "", bot.MatchTypeContains, shortenLinkHandler)

	b.Start(ctx)
	return nil
}

func shortenLinkHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message != nil && update.Message.Text != "/start" {
		var responseText string

		isUrlValid := validation.IsValidURL(update.Message.Text)

		// Можно было бы изменить сигнатуру создания в бд
		// чтобы ожидалась строка и ошибка
		// тогда можно было вызывать цепочкой
		if !isUrlValid {
			responseText = "Присланный текст не является ссылкой"
		} else {
			shortened, err := database.CreateShortenedUrlQuery(update.Message.Text)

			if err != nil {
				responseText = "Что-то пошло не так при сокращении ссылки"
			} else {
				responseText = fmt.Sprintf("Твоя ссылка преобразована \n %s/s/%d", os.Getenv("HOST_NAME"), shortened)
			}
		}

		b.SendMessage(ctx, &bot.SendMessageParams{
			ChatID: update.Message.Chat.ID,
			Text:   responseText,
			// ParseMode: models.ParseModeMarkdown,
		})
	}
}

func defaultHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      "Отправь в сообщении ссылку, которую нужно сократить",
		ParseMode: models.ParseModeMarkdown,
	})
}
