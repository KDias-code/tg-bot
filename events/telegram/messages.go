package telegram

const (
	msgHelp = `I can save your pages and links. Also I can offer you them to read.
	
				In order, to save the page, you need to send me link(with prefix, for example http//:google.com).
	
				To get random page from your list, write command /rnd
				CAUTION!!! After that this page will be removed from your list.`

	msgHello = "Hi, I am bot by @dkubay! \n\n" + msgHelp

	msgUnknownCmd    = "I don't understand you. Write /help to see, what I can do"
	msgNoPages       = "You have no saved pages :("
	msgSaved         = "Saved!"
	msgAlreadyExists = "You already saved this link)"
)
