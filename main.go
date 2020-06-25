package main

import (
	"log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL")
		token     = os.Getenv("TOKEN")
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatal(err)
	}

	b.Handle("/hello", func(m *tb.Message) {
		b.Send(m.Sender, "Hi!")
	})

	b.Handle("/anny", func(m *tb.Message) {
		b.Send(m.Sender, "Hola mi amor. Te amo mucho")
	})

	b.Handle("/edgar", func(m *tb.Message) {
		b.Send(m.Sender, "Hola, no eres el lider de pelotón")
	})

	b.Handle("/dilia", func(m *tb.Message) {
		b.Send(m.Sender, "Hola madre, la comida de hoy estaba rica")
	})

	b.Handle("/david", func(m *tb.Message) {
		b.Send(m.Sender, "Hola Davidson, te saluda el primer bot de Edwin!")
	})

	b.Start()
}
