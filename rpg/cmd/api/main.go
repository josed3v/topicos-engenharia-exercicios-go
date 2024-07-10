package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Player struct {
	Nickname string `json:"nickname"`
	Life     int    `json:"life"`
	Attack   int    `json:"attack"`
}

type Enemy struct {
	Nickname string `json:"nickname"`
	Life     int    `json:"life"`
	Attack   int    `json:"attack"`
}

type Battle struct {
	ID         string `json:"id"`
	Enemy      string `json:"enemy"`
	Player     string `json:"player"`
	DiceThrown int    `json:"diceThrown"`
}

type PlayerResponse struct {
	Message string `json:"message"`
}

type EnemyResponse struct {
	Message string `json:"message"`
}

type BattleRequest struct {
	Enemy  string `json:"enemy"`
	Player string `json:"player"`
}

type BattleResponse struct {
	ID         string `json:"id"`
	DiceThrown int    `json:"diceThrown"`
	Player     Player `json:"player"`
	Enemy      Enemy  `json:"enemy"`
	Winner     string `json:"winner"`
}

type Response struct {
	Message string `json:"message"`
}

var (
	players = make(map[string]Player)
	enemies = make(map[string]Enemy)
	battles = make(map[string]Battle)
	mu      sync.Mutex
)

func main() {
	rand.Seed(time.Now().UnixNano())
	router := mux.NewRouter()

	router.HandleFunc("/player", AddPlayer).Methods("POST")
	router.HandleFunc("/player", LoadPlayers).Methods("GET")
	router.HandleFunc("/player/{nickname}", DeletePlayer).Methods("DELETE")
	router.HandleFunc("/player/{nickname}", LoadPlayerByNickname).Methods("GET")
	router.HandleFunc("/player/{nickname}", SavePlayer).Methods("PUT")

	router.HandleFunc("/enemy", AddEnemy).Methods("POST")
	router.HandleFunc("/enemy", LoadEnemies).Methods("GET")
	router.HandleFunc("/enemy/{nickname}", LoadEnemyByNickname).Methods("GET")
	router.HandleFunc("/enemy/{nickname}", UpdateEnemy).Methods("PUT")
	router.HandleFunc("/enemy/{nickname}", DeleteEnemy).Methods("DELETE")

	router.HandleFunc("/battle", AddBattle).Methods("POST")
	router.HandleFunc("/battle", LoadBattles).Methods("GET")

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println(err)
	}
}

// Enemy Handlers

func AddEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var enemy Enemy
	if err := json.NewDecoder(r.Body).Decode(&enemy); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(EnemyResponse{Message: "Internal Server Error"})
		return
	}

	if enemy.Nickname == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(EnemyResponse{Message: "Enemy nickname is required"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := enemies[enemy.Nickname]; exists {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(EnemyResponse{Message: "Enemy nickname already exists"})
		return
	}

	enemy.Life = rand.Intn(10) + 1
	enemy.Attack = rand.Intn(10) + 1

	enemies[enemy.Nickname] = enemy

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(enemy)
}

func LoadEnemies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mu.Lock()
	defer mu.Unlock()

	enemyList := make([]Enemy, 0, len(enemies))
	for _, enemy := range enemies {
		enemyList = append(enemyList, enemy)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(enemyList)
}

func LoadEnemyByNickname(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nickname := mux.Vars(r)["nickname"]

	mu.Lock()
	defer mu.Unlock()

	if enemy, exists := enemies[nickname]; exists {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(enemy)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(EnemyResponse{Message: "Enemy nickname not found"})
}

func UpdateEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nickname := mux.Vars(r)["nickname"]

	var updateRequest struct {
		NewNickname string `json:"newNickname"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(EnemyResponse{Message: "Internal Server Error"})
		return
	}

	if updateRequest.NewNickname == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(EnemyResponse{Message: "New enemy nickname is required"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if enemy, exists := enemies[nickname]; exists {
		if _, exists := enemies[updateRequest.NewNickname]; exists && updateRequest.NewNickname != nickname {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(EnemyResponse{Message: "Enemy nickname already exists"})
			return
		}

		enemy.Nickname = updateRequest.NewNickname
		enemies[updateRequest.NewNickname] = enemy
		delete(enemies, nickname)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(enemy)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(EnemyResponse{Message: "Enemy nickname not found"})
}

func DeleteEnemy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nickname := mux.Vars(r)["nickname"]

	mu.Lock()
	defer mu.Unlock()

	if _, exists := enemies[nickname]; exists {
		delete(enemies, nickname)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(EnemyResponse{Message: "Enemy nickname not found"})
}

// Battle Handlers

func AddBattle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var battleRequest BattleRequest
	if err := json.NewDecoder(r.Body).Decode(&battleRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(Response{Message: "Internal Server Error"})
		return
	}

	if battleRequest.Enemy == "" || battleRequest.Player == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Enemy and player nicknames are required"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	player, playerExists := players[battleRequest.Player]
	enemy, enemyExists := enemies[battleRequest.Enemy]

	if !playerExists || !enemyExists {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Player or enemy not found"})
		return
	}

	if player.Life <= 0 || enemy.Life <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Message: "Player or enemy is dead"})
		return
	}

	diceThrown := rand.Intn(6) + 1
	battleID := uuid.New().String()

	winner := ""
	if diceThrown >= 1 && diceThrown <= 3 {
		player.Life -= enemy.Attack
		winner = "Enemy"
	} else if diceThrown >= 4 && diceThrown <= 6 {
		enemy.Life -= player.Attack
		winner = "Player"
	}

	battles[battleID] = Battle{
		ID:         battleID,
		Enemy:      enemy.Nickname,
		Player:     player.Nickname,
		DiceThrown: diceThrown,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(BattleResponse{
		ID:         battleID,
		DiceThrown: diceThrown,
		Player:     player,
		Enemy:      enemy,
		Winner:     winner,
	})
}

func LoadBattles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mu.Lock()
	defer mu.Unlock()

	battleList := make([]BattleResponse, 0, len(battles))
	for _, battle := range battles {
		player, _ := players[battle.Player]
		enemy, _ := enemies[battle.Enemy]
		battleList = append(battleList, BattleResponse{
			ID:         battle.ID,
			DiceThrown: battle.DiceThrown,
			Player:     player,
			Enemy:      enemy,
			Winner:     determineWinner(battle.DiceThrown, player, enemy),
		})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(battleList)
}

func determineWinner(diceThrown int, player Player, enemy Enemy) string {
	if diceThrown >= 1 && diceThrown <= 3 {
		return "Enemy"
	} else if diceThrown >= 4 && diceThrown <= 6 {
		return "Player"
	}
	return "None"
}

// Player Handlers

func AddPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var player Player
	if err := json.NewDecoder(r.Body).Decode(&player); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(PlayerResponse{Message: "Internal Server Error"})
		return
	}

	if player.Nickname == "" || player.Life == 0 || player.Attack == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(PlayerResponse{Message: "Player nickname, life and attack are required"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, exists := players[player.Nickname]; exists {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(PlayerResponse{Message: "Player nickname already exists"})
		return
	}

	players[player.Nickname] = player

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(player)
}

func LoadPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	mu.Lock()
	defer mu.Unlock()

	playerList := make([]Player, 0, len(players))
	for _, player := range players {
		playerList = append(playerList, player)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(playerList)
}

func DeletePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nickname := mux.Vars(r)["nickname"]

	mu.Lock()
	defer mu.Unlock()

	if _, exists := players[nickname]; exists {
		delete(players, nickname)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(PlayerResponse{Message: "Player nickname not found"})
}

func LoadPlayerByNickname(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nickname := mux.Vars(r)["nickname"]

	mu.Lock()
	defer mu.Unlock()

	if player, exists := players[nickname]; exists {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(player)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(PlayerResponse{Message: "Player nickname not found"})
}

func SavePlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nickname := mux.Vars(r)["nickname"]

	var updateRequest struct {
		NewNickname string `json:"newNickname"`
	}

	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(PlayerResponse{Message: "Internal Server Error"})
		return
	}

	if updateRequest.NewNickname == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(PlayerResponse{Message: "New player nickname is required"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if player, exists := players[nickname]; exists {
		if _, exists := players[updateRequest.NewNickname]; exists && updateRequest.NewNickname != nickname {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(PlayerResponse{Message: "Player nickname already exists"})
			return
		}

		player.Nickname = updateRequest.NewNickname
		players[updateRequest.NewNickname] = player
		delete(players, nickname)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(player)
		return
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(PlayerResponse{Message: "Player nickname not found"})
}
