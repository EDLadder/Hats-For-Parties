package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/EDLadder/Hats-For-Parties/config"
	"github.com/EDLadder/Hats-For-Parties/models"
	"github.com/EDLadder/Hats-For-Parties/response"
	"github.com/EDLadder/Hats-For-Parties/validators"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

func GetParties(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var parties []*models.Party

		collectionParty := client.Database("party-hats").Collection("party")
		cursor, err := collectionParty.Find(context.TODO(), bson.D{{}})
		if err != nil {
			response.ServerErrResponse(err.Error(), w)
			return
		}
		if err = cursor.All(context.TODO(), &parties); err != nil {
			println("Here")
			response.ServerErrResponse(err.Error(), w)
			return
		}
		responseValue := map[string]interface{}{
			"parties": parties,
		}
		response.SuccessResponse(responseValue, w)
	}
}

func GetHats(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var hats []*models.Hat

		collectionHat := client.Database("party-hats").Collection("hat")
		cursor, err := collectionHat.Find(context.TODO(), bson.D{{}})
		if err != nil {
			response.ServerErrResponse(err.Error(), w)
			return
		}
		if err = cursor.All(context.TODO(), &hats); err != nil {
			response.ServerErrResponse(err.Error(), w)
			return
		}
		responseValue := map[string]interface{}{
			"hats": hats,
		}
		response.SuccessResponse(responseValue, w)
	}
}

// func CreateParty(client *mongo.Client) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		envMaxHats, err := config.GetEnvVariable("MAX_PARTY_HATS_COUNT")
// 		if err != nil {
// 			response.ErrorResponse(err.Error(), w)
// 			return
// 		}
// 		partyHatsLimit, err := strconv.Atoi(envMaxHats)
// 		if err != nil {
// 			response.ErrorResponse(err.Error(), w)
// 			return
// 		}
// 		var party models.Party
// 		err = json.NewDecoder(r.Body).Decode(&party)
// 		if err != nil {
// 			response.ServerErrResponse(err.Error(), w)
// 			return
// 		}
// 		if ok, err := validators.ValidateInputs(party); !ok {
// 			response.ValidationResponse(err, w)
// 			return
// 		}
// 		if party.Hats > partyHatsLimit {
// 			response.ErrorResponse("Limit of renting hats per party is "+strconv.Itoa(partyHatsLimit), w)
// 			return
// 		}
// 		party.Status = "Started"
// 		party.UpdatedAt = time.Now()
// 		responseValue := map[string]interface{}{
// 			"id": "",
// 		}
// 		collectionParty := client.Database("party-hats").Collection("party")
// 		collectionHats := client.Database("party-hats").Collection("hat")

// 		err = client.UseSession(context.TODO(), func(sctx mongo.SessionContext) error {
// 			// Create party
// 			result, err := collectionParty.InsertOne(sctx, party)
// 			if err != nil {
// 				return err
// 			}
// 			responseValue["id"] = result.InsertedID
// 			// Get free hats
// 			hatsFilter := bson.D{
// 				{Key: "partyId", Value: bson.D{{Key: "$eq", Value: nil}}},
// 				{Key: "canBeUseAfter", Value: bson.D{{Key: "$lt", Value: primitive.NewDateTimeFromTime(time.Now())}}},
// 			}
// 			hatsOpts := options.Find().SetSort(bson.D{
// 				{Key: "firstUse", Value: 1},
// 			}).SetLimit(int64(party.Hats))

// 			hatsCursor, err := collectionHats.Find(sctx, hatsFilter, hatsOpts)
// 			if err != nil {
// 				return err
// 			}
// 			var freeHats []bson.M
// 			err = hatsCursor.All(sctx, &freeHats)
// 			if err != nil {
// 				return err
// 			}
// 			if party.Hats > len(freeHats) {
// 				return errors.New("They are only available " + strconv.Itoa(len(freeHats)) + " hats")
// 			}

// 			// For test
// 			time.Sleep(4 * time.Second)
// 			// Update hats
// 			for _, hat := range freeHats {
// 				updateFilter := bson.D{primitive.E{Key: "_id", Value: hat["_id"]}}
// 				updateData := bson.D{
// 					primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "partyId", Value: result.InsertedID}}},
// 				}

// 				if hat["firstUse"] == nil {
// 					updateData = append(updateData, primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "firstUse", Value: time.Now()}}})
// 				}

// 				_, err = collectionHats.UpdateOne(sctx, updateFilter, updateData)
// 				if err != nil {
// 					return err
// 				}
// 			}
// 			return nil
// 		})

// 		if err != nil {
// 			response.ServerErrResponse(err.Error(), w)
// 			return
// 		}
// 		response.SuccessResponse(responseValue, w)
// 	}
// }

func CreateParty(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		envMaxHats, err := config.GetEnvVariable("MAX_PARTY_HATS_COUNT")
		if err != nil {
			response.ErrorResponse(err.Error(), w)
			return
		}
		partyHatsLimit, err := strconv.Atoi(envMaxHats)
		if err != nil {
			response.ErrorResponse(err.Error(), w)
			return
		}
		var party models.Party
		err = json.NewDecoder(r.Body).Decode(&party)
		if err != nil {
			response.ServerErrResponse(err.Error(), w)
			return
		}
		if ok, err := validators.ValidateInputs(party); !ok {
			response.ValidationResponse(err, w)
			return
		}
		if party.Hats > partyHatsLimit {
			response.ErrorResponse("Limit of renting hats per party is "+strconv.Itoa(partyHatsLimit), w)
			return
		}
		party.Status = "Started"
		party.UpdatedAt = time.Now()
		responseValue := map[string]interface{}{
			"id": "",
		}
		collectionParty := client.Database("party-hats").Collection("party")
		collectionHats := client.Database("party-hats").Collection("hat")

		err = client.UseSession(context.TODO(), func(sctx mongo.SessionContext) error {
			err := sctx.StartTransaction(options.Transaction().
				SetReadConcern(readconcern.Snapshot()).
				SetWriteConcern(writeconcern.New(writeconcern.WMajority())),
			)
			if err != nil {
				return err
			}
			// Create party
			result, err := collectionParty.InsertOne(sctx, party)
			if err != nil {
				sctx.AbortTransaction(sctx)
				return err
			}
			// Get free hats
			hatsFilter := bson.D{
				{Key: "partyId", Value: bson.D{{Key: "$eq", Value: nil}}},
				{Key: "canBeUseAfter", Value: bson.D{{Key: "$lt", Value: primitive.NewDateTimeFromTime(time.Now())}}},
			}
			hatsOpts := options.Find().SetSort(bson.D{
				{Key: "firstUse", Value: 1},
			}).SetLimit(int64(party.Hats))

			hatsCursor, err := collectionHats.Find(sctx, hatsFilter, hatsOpts)
			if err != nil {
				sctx.AbortTransaction(sctx)
				return err
			}
			var freeHats []bson.M
			err = hatsCursor.All(sctx, &freeHats)
			if err != nil {
				sctx.AbortTransaction(sctx)
				return err
			}
			if party.Hats > len(freeHats) {
				return errors.New("They are only available " + strconv.Itoa(len(freeHats)) + " hats")
			}
			// Update hats
			for _, hat := range freeHats {
				updateFilter := bson.D{primitive.E{Key: "_id", Value: hat["_id"]}}
				updateData := bson.D{
					primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "partyId", Value: result.InsertedID}}},
				}

				if hat["firstUse"] == nil {
					updateData = append(updateData, primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "firstUse", Value: time.Now()}}})
				}

				_, err = collectionHats.UpdateOne(sctx, updateFilter, updateData)
				if err != nil {
					sctx.AbortTransaction(sctx)
					return err
				}
			}

			responseValue["id"] = result.InsertedID
			for {
				err = sctx.CommitTransaction(sctx)
				switch e := err.(type) {
				case nil:
					return nil
				case mongo.CommandError:
					if e.HasErrorLabel("UnknownTransactionCommitResult") {
						continue
					}
					return e
				default:
					return e
				}
			}
		})

		if err != nil {
			response.ServerErrResponse(err.Error(), w)
			return
		}
		response.SuccessResponse(responseValue, w)
	}
}

func StopParty(client *mongo.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		id, _ := primitive.ObjectIDFromHex(params["id"])

		envCleaningTime, err := config.GetEnvVariable("CLEANING_TIME")
		if err != nil {
			response.ServerErrResponse(err.Error(), w)
			return
		}
		cleaningTime, _ := strconv.Atoi(envCleaningTime)
		newTime := time.Now().Add(time.Duration(cleaningTime) * time.Second)

		collectionHats := client.Database("party-hats").Collection("hat")

		filter := bson.D{primitive.E{Key: "partyId", Value: id}}
		update := bson.D{
			primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "partyId", Value: nil}}},
			primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "canBeUseAfter", Value: newTime}}},
		}

		_, updateErr := collectionHats.UpdateMany(context.TODO(), filter, update)

		if updateErr != nil {
			response.ServerErrResponse(err.Error(), w)
			return
		}

		collectionParty := client.Database("party-hats").Collection("party")
		res, updateErr := collectionParty.UpdateOne(context.TODO(), bson.D{
			primitive.E{Key: "_id", Value: id}},
			bson.D{
				primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "status", Value: "Stopped"}}},
				primitive.E{Key: "$set", Value: bson.D{primitive.E{Key: "updatedAt", Value: time.Now()}}},
			})
		if updateErr != nil {
			response.ServerErrResponse(err.Error(), w)
			return
		}
		if res.MatchedCount == 0 {
			response.ErrorResponse("Party does not exist", w)
			return
		}

		responseValue := map[string]interface{}{
			"message": "Party stopped",
		}
		response.SuccessResponse(responseValue, w)
	}
}
