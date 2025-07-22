package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {

	//менять вот это
	projecType := "kastrulivideo"
	countScen := 3
	//var chatId int64 = -861820468 // хуеверты
	var chatId int64 = 400809044 // личка
	//var chatId int64 = -853687329 // недвижимость трясунова
	//var chatId int64 = -614851423 // юрий мурадян
	//var chatId int64 = -947775305 // добрая ведьма
	//var chatId int64 = -913740018 // байеры
	//var chatId int64 = -897655966 // тейпирование

	var imageType = "mp4"
	// после не надо

	bot, err := tgbotapi.NewBotAPI("6252362203:AAEDfzcgEShaJCDs6rZmoMVv6LipZFXNgbE")
	if err != nil {
		log.Panic(err)
	}

	postToLinks, err := ioutil.ReadFile("./links/postTo/inp.html")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	inPostLinks, err := ioutil.ReadFile("./links/inPost/inp.html")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)

	//switch projecType {
	//case "tomaexp2":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "tomavyaz":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "tomaamig":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "tyurina2":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//
	//case "tomaexp":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "tomamas":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "tape":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "tape3":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "5prism":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "dkouch":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "tapenew":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "tapenew2":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//
	//case "muradyan":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//
	//case "tryasunova":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//
	//case "bayer":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//
	//case "dobvedma":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "kimp":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "ispzhel":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "bebra":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "shoporisto":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "schoolb":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//
	//case "murobr":
	//	funcPidarasUraObr(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "murzhen":
	//	funcPidarasUraZhen(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "murbiz":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "murotn":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "murpsih":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "ulyana":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "toni":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "tomapsih":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "tomaruk":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "tomasmm":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "glaza":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "smmsite":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "smmchan":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "mirprog":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "mirbuz":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "huistka":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "justceo":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "tomaster":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "muradzh":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "muradbus":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "muradotn":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "muradps":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "tapezhiv":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "yuranew":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "yuranew2":
	//	funcPidarasUraObr(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "logoped":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "orlova":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "orlovabot":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "murnewstart":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "kimpsite":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "turkan":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "kimp1":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "kimpmaster":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "gorbun":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "yuranew3":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "tyurina":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "tyurina1":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//case "ribina":
	//	funcName(err, bot, projecType, string(postToLinks), string(inPostLinks), countScen, chatId, imageType)
	//
	//}

}

func funcName(
	err error, bot *tgbotapi.BotAPI, prjType string, postToLinks string, inPostLinks string, countScen int,
	chatId int64, imageType string,
) {
	scanner := bufio.NewScanner(strings.NewReader(postToLinks))
	var linesPostToLinks []string
	for scanner.Scan() {
		linesPostToLinks = append(linesPostToLinks, scanner.Text())
	}

	scanner = bufio.NewScanner(strings.NewReader(inPostLinks))
	var linesInPostLinks []string
	for scanner.Scan() {
		linesInPostLinks = append(linesInPostLinks, scanner.Text())
	}

	if len(linesPostToLinks) == len(linesInPostLinks) {
		for i := 0; i < len(linesPostToLinks); i += countScen {
			//if i%9 == 0 {
			//	time.Sleep(5 * time.Second)
			//}
			for j := 1; j <= countScen && j+i-1 < len(linesPostToLinks); j++ {
				temp := fmt.Sprintf("./scenarios/%s%d.html", prjType, j)
				//fileContent, err := ioutil.ReadFile("./scenarios/" + prjType + "1.html")
				fileContent, err := ioutil.ReadFile(temp)
				if err != nil {
					fmt.Println("Error reading file:", err)
					return
				}
				formattedText := string(fileContent)
				formattedText = strings.ReplaceAll(formattedText, "#", linesInPostLinks[i+j-1])
				var typePhoto int
				if prjType == "bayer" {
					typePhoto = rand.Intn(4)
				} else {
					typePhoto = j
				}

				msg1 := tgbotapi.NewMessage(chatId, "👇 пост для "+linesPostToLinks[i+j-1])
				// Set up a message configuration with rich text formatting and an image
				temp = fmt.Sprintf("./images/%s/%s%d.%s", prjType, prjType, typePhoto, imageType)
				//msg := tgbotapi.NewPhotoUpload(505853752, "./images/"+prjType+"/"+prjType+"1.jpg")
				_, err = bot.Send(msg1)
				if err != nil {
					log.Panic(err)
				}
				if prjType == "ispzhel" || prjType == "justceo" {
					msg := tgbotapi.NewMessage(chatId, formattedText)
					msg.ParseMode = tgbotapi.ModeHTML
					//msg.Caption = formattedText
					_, err = bot.Send(msg)
					if err != nil {
						log.Panic(err)
					}
				} else {
					if imageType == "mp4" {
						//msg := tgbotapi.NewPhotoUpload(chatId, temp)
						msg := tgbotapi.NewVideoUpload(chatId, temp)
						msg.ParseMode = tgbotapi.ModeHTML
						msg.Caption = formattedText
						_, err = bot.Send(msg)
						if err != nil {
							log.Panic(err)
						}
					} else {
						msg := tgbotapi.NewPhotoUpload(chatId, temp)
						//msg = tgbotapi.NewV
						msg.ParseMode = tgbotapi.ModeHTML
						msg.Caption = formattedText
						_, err = bot.Send(msg)
						if err != nil {
							log.Panic(err)
						}
					}

				}

				// Send the message to the user

			}

		}
	}

}

func funcPidarasUraObr(
	err error, bot *tgbotapi.BotAPI, prjType string, postToLinks string, inPostLinks string, countScen int,
	chatId int64, imageType string,
) {
	scanner := bufio.NewScanner(strings.NewReader(postToLinks))
	var linesPostToLinks []string
	for scanner.Scan() {
		linesPostToLinks = append(linesPostToLinks, scanner.Text())
	}

	scanner = bufio.NewScanner(strings.NewReader(inPostLinks))
	var linesInPostLinks []string
	for scanner.Scan() {
		linesInPostLinks = append(linesInPostLinks, scanner.Text())
	}

	if len(linesPostToLinks) == len(linesInPostLinks) {
		for i := 0; i < len(linesPostToLinks); i += countScen {
			//if i%9 == 0 {
			//	time.Sleep(5 * time.Second)
			//}
			for j := 1; j <= countScen && j+i-1 < len(linesPostToLinks); j++ {
				temp := fmt.Sprintf("./scenarios/%s%d.html", prjType, j)
				//fileContent, err := ioutil.ReadFile("./scenarios/" + prjType + "1.html")
				fileContent, err := ioutil.ReadFile(temp)
				if err != nil {
					fmt.Println("Error reading file:", err)
					return
				}
				formattedText := string(fileContent)
				formattedText = strings.ReplaceAll(formattedText, "#", linesInPostLinks[i+j-1])
				var typePhoto int
				if prjType == "bayer" {
					typePhoto = rand.Intn(4)
				} else {
					typePhoto = j
				}
				if typePhoto == 1 || typePhoto == 2 {
					imageType = "mp4"
				} else {
					imageType = "png"
				}
				msg1 := tgbotapi.NewMessage(chatId, "👇 пост для "+linesPostToLinks[i+j-1])
				// Set up a message configuration with rich text formatting and an image
				temp = fmt.Sprintf("./images/%s/%s%d.%s", prjType, prjType, typePhoto, imageType)
				//msg := tgbotapi.NewPhotoUpload(505853752, "./images/"+prjType+"/"+prjType+"1.jpg")
				_, err = bot.Send(msg1)
				if err != nil {
					log.Panic(err)
				}
				if typePhoto == 1 || typePhoto == 2 {
					//msg := tgbotapi.NewPhotoUpload(chatId, temp)
					msg := tgbotapi.NewVideoUpload(chatId, temp)
					msg.ParseMode = tgbotapi.ModeHTML
					msg.Caption = formattedText
					_, err = bot.Send(msg)
					if err != nil {
						log.Panic(err)
					}
				} else {
					msg := tgbotapi.NewPhotoUpload(chatId, temp)
					//msg = tgbotapi.NewV
					msg.ParseMode = tgbotapi.ModeHTML
					msg.Caption = formattedText
					_, err = bot.Send(msg)
					if err != nil {
						log.Panic(err)
					}
				}

				// Send the message to the user

			}
		}
	}
}

func funcPidarasUraZhen(
	err error, bot *tgbotapi.BotAPI, prjType string, postToLinks string, inPostLinks string, countScen int,
	chatId int64, imageType string,
) {
	scanner := bufio.NewScanner(strings.NewReader(postToLinks))
	var linesPostToLinks []string
	for scanner.Scan() {
		linesPostToLinks = append(linesPostToLinks, scanner.Text())
	}

	scanner = bufio.NewScanner(strings.NewReader(inPostLinks))
	var linesInPostLinks []string
	for scanner.Scan() {
		linesInPostLinks = append(linesInPostLinks, scanner.Text())
	}

	if len(linesPostToLinks) == len(linesInPostLinks) {
		for i := 0; i < len(linesPostToLinks); i += countScen {
			//if i%9 == 0 {
			//	time.Sleep(5 * time.Second)
			//}
			for j := 1; j <= countScen && j+i-1 < len(linesPostToLinks); j++ {
				temp := fmt.Sprintf("./scenarios/%s%d.html", prjType, j)
				//fileContent, err := ioutil.ReadFile("./scenarios/" + prjType + "1.html")
				fileContent, err := ioutil.ReadFile(temp)
				if err != nil {
					fmt.Println("Error reading file:", err)
					return
				}
				formattedText := string(fileContent)
				formattedText = strings.ReplaceAll(formattedText, "#", linesInPostLinks[i+j-1])
				var typePhoto int
				if prjType == "bayer" {
					typePhoto = rand.Intn(4)
				} else {
					typePhoto = j
				}
				if typePhoto == 2 {
					imageType = "mp4"
				} else {
					imageType = "jpg"
				}
				msg1 := tgbotapi.NewMessage(chatId, "👇 пост для "+linesPostToLinks[i+j-1])
				// Set up a message configuration with rich text formatting and an image
				temp = fmt.Sprintf("./images/%s/%s%d.%s", prjType, prjType, typePhoto, imageType)
				//msg := tgbotapi.NewPhotoUpload(505853752, "./images/"+prjType+"/"+prjType+"1.jpg")
				_, err = bot.Send(msg1)
				if err != nil {
					log.Panic(err)
				}
				if typePhoto == 2 {
					//msg := tgbotapi.NewPhotoUpload(chatId, temp)
					msg := tgbotapi.NewVideoUpload(chatId, temp)
					msg.ParseMode = tgbotapi.ModeHTML
					msg.Caption = formattedText
					_, err = bot.Send(msg)
					if err != nil {
						log.Panic(err)
					}
				} else {
					msg := tgbotapi.NewPhotoUpload(chatId, temp)
					//msg = tgbotapi.NewV
					msg.ParseMode = tgbotapi.ModeHTML
					msg.Caption = formattedText
					_, err = bot.Send(msg)
					if err != nil {
						log.Panic(err)
					}
				}
			}
		}
	}
}
