package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"rpg-game/database/mongo"
	"rpg-game/entities"
	"rpg-game/entities/contracts"
	"rpg-game/entities/enums"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx)
	if err != nil {
		fmt.Errorf(err.Error())
	}

	database := client.Database("rpg-game")
	collection := database.Collection("players")

	inventory := []contracts.Object{
		entities.Weapon{
			Name:   "Sword",
			Damage: 10,
		},
		entities.Armor{
			Name:       "Chestplate",
			Resistance: 5,
		},
	}

	player := entities.Datasheet{
		Name:      "Baldurs",
		Age:       25,
		Inventory: inventory,
		Class:     enums.Barbarian,
		Race:      enums.Alabrastos,
		Element:   enums.Aqua,
		Attributes: entities.Attributes{
			Health:    20,
			Energy:    10,
			Attack:    5,
			Defense:   0,
			Evasion:   0,
			Precision: 1,
			Dexterity: 0,
		},
	}
	collection.InsertOne(ctx, player)

	singleObjet := collection.FindOne(ctx, bson.M{"name": "Baldurs"})
	var result entities.Datasheet
	singleObjet.Decode(&result)
	fmt.Println(player)

	// Disconnect from MongoDB when done
	err = client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Disconnected from MongoDB!")
}
