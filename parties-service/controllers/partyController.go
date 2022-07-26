package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/EDLadder/Hats-For-Parties/config"
	"github.com/EDLadder/Hats-For-Parties/db"
	"github.com/EDLadder/Hats-For-Parties/models"
	"github.com/EDLadder/Hats-For-Parties/response"
	"github.com/EDLadder/Hats-For-Parties/validators"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client = db.Dbconnect()

var GetParties = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var parties []*models.Party

	collectionParty := client.Database("party-hats").Collection("party")
	cursor, err := collectionParty.Find(context.TODO(), bson.D{{}})
	if err != nil {
		response.ServerErrResponse(err.Error(), w)
		return
	}
	for cursor.Next(context.TODO()) {
		var party models.Party
		err := cursor.Decode(&party)
		if err != nil {
			log.Fatal(err)
		}

		parties = append(parties, &party)
	}
	if err := cursor.Err(); err != nil {
		response.ServerErrResponse(err.Error(), w)
		return
	}
	partiesString, _ := json.Marshal(parties)
	responseValue := "{\"parties\":" + string(partiesString) + "}"
	response.SuccessResponse(responseValue, w)
})

var GetHats = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var hats []*models.Hat

	collectionHat := client.Database("party-hats").Collection("hat")
	cursor, err := collectionHat.Find(context.TODO(), bson.D{{}})
	if err != nil {
		response.ServerErrResponse(err.Error(), w)
		return
	}
	for cursor.Next(context.TODO()) {
		var hat models.Hat
		err := cursor.Decode(&hat)
		if err != nil {
			log.Fatal(err)
		}

		hats = append(hats, &hat)
	}
	if err := cursor.Err(); err != nil {
		response.ServerErrResponse(err.Error(), w)
		return
	}
	hatsString, _ := json.Marshal(hats)
	responseValue := "{\"hats\":" + string(hatsString) + "}"
	response.SuccessResponse(responseValue, w)
})

var CreateParty = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	partyHatsLimit, _ := strconv.Atoi(config.GetEnvVariable("MAX_PARTY_HATS_COUNT"))

	var party models.Party
	decodeErr := json.NewDecoder(r.Body).Decode(&party)

	if decodeErr != nil {
		response.ServerErrResponse(decodeErr.Error(), w)
		return
	}
	if ok, validateErrors := validators.ValidateInputs(party); !ok {
		response.ValidationResponse(validateErrors, w)
		return
	}
	if party.Hats > partyHatsLimit {
		response.ErrorResponse("Limit of renting hats per party is "+strconv.Itoa(partyHatsLimit), w)
		return
	}

	collectionParty := client.Database("party-hats").Collection("party")
	collectionHats := client.Database("party-hats").Collection("hat")

	// Get free hats
	hatsOpts := options.Find().SetSort(bson.D{
		{Key: "firstUse", Value: 1},
	}).SetLimit(int64(party.Hats))

	hatsFilter := bson.D{
		{Key: "partyId", Value: bson.D{{Key: "$eq", Value: nil}}},
		{Key: "canBeUseAfter", Value: bson.D{{Key: "$lt", Value: primitive.NewDateTimeFromTime(time.Now())}}},
	}

	hatsCursor, _ := collectionHats.Find(context.TODO(), hatsFilter, hatsOpts)

	var freeHats []bson.M
	hatsError := hatsCursor.All(context.TODO(), &freeHats)

	if hatsError != nil {
		response.ServerErrResponse(hatsError.Error(), w)
		return
	}
	if party.Hats > len(freeHats) {
		response.ErrorResponse("They are only available "+strconv.Itoa(len(freeHats))+" hats", w)
		return
	}

	// Start party
	party.Status = "Started"
	party.UpdatedAt = time.Now()

	result, err := collectionParty.InsertOne(context.TODO(), party)

	if err != nil {
		response.ServerErrResponse(err.Error(), w)
		return
	}

	// Update hats party id
	for _, hat := range freeHats {
		updateFilter := bson.D{primitive.E{Key: "_id", Value: hat["_id"]}}
		updateData := bson.D{
			primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "partyId", Value: result.InsertedID}}},
		}

		if hat["firstUse"] == nil {
			updateData = append(updateData, primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "firstUse", Value: time.Now()}}})
		}

		collectionHats.UpdateOne(context.TODO(), updateFilter, updateData)
	}

	// Return party ID
	res, _ := json.Marshal(result.InsertedID)
	response.SuccessResponse("{\"ID\": \""+strings.Replace(string(res), `"`, ``, 2)+"\"}", w)
})

var StopParty = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])

	cleaningTime, _ := strconv.Atoi(config.GetEnvVariable("CLEANING_TIME"))
	newTime := time.Now().Add(time.Duration(cleaningTime) * time.Second)

	collectionHats := client.Database("party-hats").Collection("hat")

	filter := bson.D{primitive.E{Key: "partyId", Value: id}}
	update := bson.D{
		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "partyId", Value: nil}}},
		primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "canBeUseAfter", Value: newTime}}},
	}

	_, err := collectionHats.UpdateMany(context.TODO(), filter, update)

	if err != nil {
		response.ServerErrResponse(err.Error(), w)
		return
	}

	collectionParty := client.Database("party-hats").Collection("party")
	res, err := collectionParty.UpdateOne(context.TODO(), bson.D{
		primitive.E{Key: "_id", Value: id}},
		bson.D{
			primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "status", Value: "Stopped"}}},
			primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "updatedAt", Value: time.Now()}}},
		})
	if err != nil {
		response.ServerErrResponse(err.Error(), w)
		return
	}
	if res.MatchedCount == 0 {
		response.ErrorResponse("Party does not exist", w)
		return
	}
	result := "{\"message\": \"Party stopped\"}"
	response.SuccessResponse(result, w)
})
