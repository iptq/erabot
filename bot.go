package erabot

type Bot struct {
	config   Config
	commands map[string]Command
}

func New() (bot *Bot) {
	bot = &Bot{}
	return
}

func (bot *Bot) Close() {

}

func (bot *Bot) RegisterCommand() {

}

func (bot *Bot) Run() {

}
