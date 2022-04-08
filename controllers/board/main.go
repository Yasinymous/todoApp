package BoardController

import (
	h "TaskManagement/helpers"
	db "TaskManagement/models"
	"encoding/json"

	b "TaskManagement/models/board"
	u "TaskManagement/models/user"
	"fmt"
	"strconv"

	// "encoding/json"
	"net/http"

	"github.com/gorilla/context"
)

func GetAll(res http.ResponseWriter, req *http.Request) {
	user := &u.User{}
	User := context.Get(req, "User")
	if u, ok := User.(u.User); ok {
		user = &u
	}

	var boards []b.Board
	result := db.GetDB().Table("boards").Where("user_id = ?", user.ID).Find(&boards)

	response := h.Message(true, "Panolar Listelendi!")
	response["data"] = result.Value

	h.Respond(res, response)
}

func GetBoard(res http.ResponseWriter, req *http.Request) {
	user := &u.User{}
	User := context.Get(req, "User")
	if u, ok := User.(u.User); ok {
		user = &u
	}
	boardId, ok := req.URL.Query()["boardId"]
	if !ok || len(boardId[0]) < 1 {
		h.Respond(res, h.Message(false, "Url Param 'key' is missing!"))
		return
	}
	bId, err := strconv.Atoi(boardId[0])
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	board := &b.Board{}
	db.GetDB().Table("boards").Where("id = ?", bId).Where("user_id = ?", user.ID).First(board)
	if board.Title == "" {
		h.Respond(res, h.Message(false, "Pano Bulunamadi!"))
		return
	}

	response := h.Message(true, "Pano Bulundu")
	response["data"] = board

	h.Respond(res, response)
}

func AddBoard(res http.ResponseWriter, req *http.Request) {
	board := &b.Board{}
	user := &u.User{}
	User := context.Get(req, "User")
	if u, ok := User.(u.User); ok {
		user = &u
	}
	errbody := json.NewDecoder(req.Body).Decode(board)
	if errbody != nil {
		h.Respond(res, h.Message(false, "Geçersiz istek. Lütfen kontrol ediniz!"))
		return
	}
	board.UserId = int(user.ID)
	db.GetDB().Create(board)

	if user.ID <= 0 {
		h.Respond(res, h.Message(false, "Bağlantı hatası oluştu. Pano yaratılamadı!"))
		return
	}

	response := h.Message(true, "Pano başarıyla yaratıldı!")
	response["data"] = board

	h.Respond(res, response)
}

func SetBoard(res http.ResponseWriter, req *http.Request) {
	user := &u.User{}
	User := context.Get(req, "User")
	if u, ok := User.(u.User); ok {
		user = &u
	}
	boardId, ok := req.URL.Query()["boardId"]
	if !ok || len(boardId[0]) < 1 {
		h.Respond(res, h.Message(false, "Url Param 'key' is missing!"))
		return
	}

	updateBoard := &b.Board{}
	board := &b.Board{}
	errbody := json.NewDecoder(req.Body).Decode(updateBoard)
	if errbody != nil {
		h.Respond(res, h.Message(false, "Geçersiz istek. Lütfen kontrol ediniz!"))
		return
	}

	bId, err := strconv.Atoi(boardId[0])
	if err != nil {
		// handle error
		fmt.Println(err)
	}

	db.GetDB().First(&board, bId)
	if board.UserId != int(user.ID) {
		h.Respond(res, h.Message(false, "Yetkiniz Bulunamadi!"))
		return
	}
	if board.Title == "" {
		h.Respond(res, h.Message(false, "Pano Bulunamadi!"))
		return
	}

	board.Title = updateBoard.Title
	board.Description = updateBoard.Description

	db.GetDB().Save(&board)

	response := h.Message(true, "Board UPDATE SUCCESS!")
	response["data"] = board
	h.Respond(res, response)
}

func DeleteBoard(res http.ResponseWriter, req *http.Request) {

	user := &u.User{}
	User := context.Get(req, "User")
	if u, ok := User.(u.User); ok {
		user = &u
	}
	boardId, ok := req.URL.Query()["boardId"]
	if !ok || len(boardId[0]) < 1 {
		h.Respond(res, h.Message(false, "Url Param 'key' is missing!"))
		return
	}

	bId, err := strconv.Atoi(boardId[0])
	if err != nil {
		// handle error
		fmt.Println(err)
	}
	board := &b.Board{}
	db.GetDB().First(&board, bId)
	if board.UserId != int(user.ID) {
		h.Respond(res, h.Message(false, "Yetkiniz Bulunamadi!"))
		return
	}
	if board.Title == "" {
		h.Respond(res, h.Message(false, "Pano Bulunamadi!"))
		return
	}

	db.GetDB().Delete(&board)

	response := h.Message(true, "Board DELETE SUCCESS!")
	response["data"] = board
	h.Respond(res, response)

}
