package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type App struct {
	Rooms []Room
}

func (app *App) handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", app.HomeHandler)

	// Rooms
	router.HandleFunc("/room/{id}", app.GetRoomHandler).Methods("GET")
	router.HandleFunc("/room", app.PostRoomHandler).Methods("POST")
	router.HandleFunc("/room/{id}", app.DeleteRoomHandler).Methods("DELETE")
	router.HandleFunc("/room/{id}", app.UpdateRoomHandler).Methods("PUT")

	// SimpleTasks

	// ComplexTasks

	// Comments

	log.Fatal(http.ListenAndServe(":10000", router))
}

func (app *App) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: HomeHandler")
	fmt.Fprintf(w, "/")
}

func (app *App) GetRoomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: GetRoomHandler")

	id, ok := mux.Vars(r)["id"]
	if !ok {
		fmt.Println("Error: No id parameter found.")
		return
	}

	for _, room := range app.Rooms {
		if room.ID == IDType(id) {
			json.NewEncoder(w).Encode(room)
			return
		}
	}

	fmt.Println("Error: Room with specified id not found.")
}

func (app *App) PostRoomHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: PostRoomHandler")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var room Room

	if err := json.Unmarshal(reqBody, &room); err != nil {
		fmt.Println("Error: Problem with parsing received JSON.")
		return
	}

	app.Rooms = append(app.Rooms, room)
	json.NewEncoder(w).Encode(room)
	fmt.Println("Room has been posted, id: " + room.ID)
}

func (app *App) DeleteRoomHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		fmt.Println("Error: No id parameter found.")
		return
	}

	for idx, room := range app.Rooms {
		if room.ID == IDType(id) {
			app.Rooms = append(app.Rooms[:idx], app.Rooms[idx+1:]...)
			json.NewEncoder(w).Encode(room)
			fmt.Println("Room has been deleted, id: " + id)
			return
		}
	}

	fmt.Println("Error: Room with specified id not found.")
}

func (app *App) UpdateRoomHandler(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["id"]
	if !ok {
		fmt.Println("Error: No id parameter found.")
		return
	}

	var updatedRoom Room
	reqBody, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(reqBody, &updatedRoom)

	for idx, room := range app.Rooms {
		if room.ID == IDType(id) {
			room.Users = updatedRoom.Users
			room.SmallTasks = updatedRoom.SmallTasks
			room.ComplexTasks = updatedRoom.ComplexTasks
			room.Comments = updatedRoom.Comments

			app.Rooms[idx] = room
			json.NewEncoder(w).Encode(room)
			fmt.Println("Room has been updated, id: " + id)
			return
		}
	}

	fmt.Println("Error: Room with specified id not found.")
}

func main() {
	fmt.Println("Hello Hacknarok2022")
	app := App{}
	app.Rooms = []Room{
		{
			ID:           "1000",
			Users:        []string{"szymon", "igor"},
			SmallTasks:   map[string]([]bool){"task0": []bool{false, true}, "task1": []bool{false, true}},
			ComplexTasks: map[string]([]float32){"taskcomplex0": []float32{.75, .9}},
		},
		{ID: "1001"},
		{ID: "1002"},
		{ID: "1003"},
	}
	app.handleRequests()
}
