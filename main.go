package main

import (
	tele "gopkg.in/telebot.v4"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	answers = [16]string{
		"Да",
		"Определенно - да!",
		"Бесспорно!",
		"Вне всяких сомнений!",
		"Нет",
		"Точно нет!",
		"Ни за что!",
		"Не бывать этому!",
		"Весьма сомнительно...",
		"Скорее всего - нет...",
		"Вероятность этого крайне мала...",
		"Низкие шансы!",
		"Возможно это 'да'...",
		"В тумане времён, неясно виднеется 'да'",
		"Возможно...",
		"Скорее всего",
	}

	pharses = [11]string{
		"Расплетаю нити судьбы...",
		"Жду ответа от звезд...",
		"Смотрю в магический шар...",
		"Ищу ответа во вселенной...",
		"Гадаю на куриных костях...",
		"Совещаюсь с Сивиллами...",
		"Жду, когда древний дух ответит мне...",
		"Призываю тень отца Гамлета...",
		"Собираю совет десяти...",
		"Нарушаю пространственно временной континуум...",
		"Открываю дверь в будущее...",
	}
)

func main() {
	pref := tele.Settings{
		Token:  os.Getenv("TOK"),
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	b.Handle(tele.OnText, func(c tele.Context) error {
		if strings.HasPrefix(strings.ToLower(c.Message().Text), "скажи") {
			go func() {
				c.Reply(pharses[rand.Intn(len(pharses))])
				time.Sleep(time.Second * 4)
				c.Reply(answers[rand.Intn(len(answers))])
			}()
		}
		return nil
	})

	b.Start()
}
