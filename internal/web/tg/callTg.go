package tg

import (
	"fmt"
	"github.com/elliotchance/pie/pie"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
	"strings"
)

const (
	mainMenu   = "main menu"
	blacklist  = "open blacklist"
	request    = "new request"
	parameters = "request parameters"
	incorrect  = "invalid command"
	end        = "Nothing found. Make a new request and try entering other request parameters"
)

var (
	menu = map[string]string{
		mainMenu:  "menu",
		blacklist: "blacklist",
	}
)

func (tA *tgApi) Connection() error {

	for update := range tA.updatesChan {
		if update.Message == nil {
			continue
		}

		tA.initSources(update.Message.Chat.ID)

		tA.logger.Infof("tg: [%s] %s", update.Message.From.UserName, update.Message.Text)

		if _, ok := menu[update.Message.Text]; ok {
			update.Message.Text = menu[update.Message.Text]
		}

		tA.handlerMsg(update.Message.Text, update, tgbotapi.NewMessage(update.Message.Chat.ID, ""))

	}
	return nil
}

func (tA *tgApi) generateAuthUrl() string {
	return tA.authURL
}

func createMenu(str string) tgbotapi.ReplyKeyboardMarkup {

	selectMenu := []string{}

	switch str {
	case mainMenu:
		selectMenu = append(selectMenu, request, "next", "blacklist")
	case blacklist:
		selectMenu = append(selectMenu, blacklist, "remove from blacklist", mainMenu)
	case parameters:
		selectMenu = append(selectMenu, "send request", mainMenu)
	}

	menu := tgbotapi.NewReplyKeyboard()
	button := tgbotapi.NewKeyboardButtonRow()

	for _, menuButton := range selectMenu {
		button = append(button, tgbotapi.NewKeyboardButton(menuButton))
		menu = tgbotapi.NewReplyKeyboard(button)
	}

	return menu
}

func (tA *tgApi) handlerMsg(updMsg string, update tgbotapi.Update, msg tgbotapi.MessageConfig) {

	a, ok := tA.checkIsBoolOn(updMsg)
	if !ok {
		msg.Text = a
	} else {
		updMsg = a
		switch updMsg {
		case "/start":
			if !strings.Contains(tA.DataCollect.CheckStatus(), "already") {
				msg.Text = fmt.Sprintf("follow the link to use the bot's functions: \n%s", tA.generateAuthUrl())
			} else {
				msg.Text = fmt.Sprintf("%s\nclick on \"new request\" to set up the search and see the photos", tA.DataCollect.CheckStatus())
				msg.ReplyMarkup = createMenu(mainMenu)
			}
		case "menu":
			msg.ReplyMarkup = createMenu(mainMenu)
			msg.Text = mainMenu
		case request:
			msg.ReplyMarkup = createMenu(parameters)
			msg.Text, ok = tA.requestToDataSource(updMsg)
			if ok {
				updMsg = msg.Text
			}
		case "next":
			msg.ReplyMarkup = createMenu(mainMenu)
			url, id := tA.formattingMsgResponse()
			if strings.Contains(id, end) {
				tA.requestToDataSource(" ")
				url, id = tA.formattingMsgResponse()
			}
			msgMedia := tgbotapi.NewMediaGroup(update.Message.Chat.ID, url)
			msg.Text = fmt.Sprintf("%s\nclick \"next\" to continue viewing or make a new request", id)
			tA.bot.Send(msgMedia)
		case "blacklist":
			msg.ReplyMarkup = createMenu(blacklist)
			msg.Text = tA.handleBlacklist("get")
		case "remove from blacklist":
			msg.Text = tA.handleBlacklist("remove")
		default:
			msg.Text = tA.inputBlacklist(update.Message.Text,strconv.FormatInt(update.Message.Chat.ID, 10))
		}
	}
	if _, err := tA.bot.Send(msg); err != nil {
		msg.Text = "error"
		tA.bot.Send(msg)
	}
}

func (tA *tgApi) checkIsBoolOn(updMsg string) (string, bool) {
	if updMsg == "menu" || updMsg == "/start" {
		tA.sC.removeIsOn = false
		tA.sC.requestIsOn = false
	}

	if tA.sC.removeIsOn == true {
		return tA.handleBlacklist(updMsg), false
	}
	if tA.sC.requestIsOn == true {
		return tA.requestToDataSource(updMsg)
	}
	return updMsg, true
}

func (tA *tgApi) requestToDataSource(updMsg string) (string, bool) {
	tA.sC.requestIsOn = true
	statusFromDataSource, ok := tA.DataCollect.SetRequestParametrs(updMsg)
	if !ok {
		return statusFromDataSource, ok
	}
	tA.sC.requestIsOn = false
	tA.getDataFromSource()
	return "next", true
}

func (tA *tgApi) getDataFromSource() {
	tA.result = nil
	tA.sC.indexCounter = 0
	tA.logger.Info("tg: prepare to send request from TG")
	tA.result = tA.DataCollect.GetData(request)
}

func (tA *tgApi) formattingMsgResponse() ([]interface{}, string) {
	var (
		resPhoto   []interface{}
		itemNumber = 1
		listID     string
	)

	tA.blacklist = tA.Storage.Get(tA.chatId)
	sieveIndex := filter(tA.blacklist, tA.result["id"], tA.result["url"]) //indexes of the elements that will be shown

	for i, _ := range sieveIndex {
		if sieveIndex[tA.sC.indexCounter] < len(tA.result["url"]) {
			a := tgbotapi.FileURL(fmt.Sprintf("%s", tA.result["url"][sieveIndex[tA.sC.indexCounter]]))
			temp := tgbotapi.NewInputMediaPhoto(a)
			resPhoto = append(resPhoto, temp)
			listID += fmt.Sprintf("%d. %s date: %s enter ' %d ' for ignore\n",
				itemNumber,
				tA.result["id"][sieveIndex[tA.sC.indexCounter]],
				tA.result["date"][sieveIndex[tA.sC.indexCounter]],
				sieveIndex[tA.sC.indexCounter])
			tA.sC.indexCounter++
			itemNumber++
			i++
			if i == 9 {
				break
			}
		} else {
			listID += fmt.Sprint(end)
			break
		}
	}
	return resPhoto, listID
}

func (tA *tgApi) handleBlacklist(updMsg string) string {
	tA.blacklist = tA.Storage.Get(tA.chatId)
	var response string

	if tA.sC.removeIsOn == true {
		msg, err := strconv.Atoi(updMsg)
		if err != nil {
			return fmt.Sprintf("%s \nplease enter the ID's number before the link", incorrect)
		}
		if msg > len(tA.blacklist) || msg < 1 {
			return fmt.Sprintf("invalid number")
		} else {
			tA.Storage.Delete(tA.blacklist[msg-1],tA.chatId)
			tA.sC.removeIsOn = false
			return fmt.Sprintf("%s has been removed from the blacklist", tA.blacklist[msg-1])
		}
	}

	for i, v := range tA.blacklist {
		response += fmt.Sprintf("%d. %s", i+1, v)
	}

	if updMsg == "remove" {
		tA.sC.removeIsOn = true
		return fmt.Sprintf("who should be removed from the blacklist? \n%s \n enter number", response)
	}

	return fmt.Sprintf("currently on the blacklist:\n%s", response)
}

func (tA *tgApi) inputBlacklist(userId string,chatId string) string {
	id, err := strconv.Atoi(userId)
	if err != nil {
		return "invalid command"
	}

	sieveIndex := filter(tA.blacklist, tA.result["id"], tA.result["url"])

	for i := 0; i <= len(tA.result["id"]); i++ {
		if id == i {
			tA.Storage.Put(tA.result["id"][sieveIndex[i]],chatId)
			return fmt.Sprintf("%s added to blacklist", tA.result["id"][sieveIndex[i]])
		}
	}
	tA.blacklist = tA.Storage.Get(tA.chatId)

	return incorrect
}

func filter(blacklist, id, url []interface{}) []int {
	index := pie.Ints{}

	for i, v := range id {
		for _, q := range blacklist {
			if v == q {
				index = append(index, i)
			}
		}
	}

	tempArray := pie.Ints{}
	for i := 0; i <= len(url); i++ {
		tempArray = append(tempArray, i)
	}

	_, sieve := tempArray.Diff(index)

	return sieve
}
