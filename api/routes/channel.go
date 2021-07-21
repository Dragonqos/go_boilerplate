package routes

import (
	"github.com/Dragonqos/go_boilerplate/api/model"
	"github.com/Dragonqos/go_boilerplate/api/repository"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetCollectionChannels(w http.ResponseWriter, r *http.Request) {
	results := repository.ChannelRepo.FindAll()

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func GetChannel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var res model.ResponseResult
	result := repository.ChannelRepo.FindById(id)

	if result == nil {
		res.Error = "Not found"
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(res)
		return
	}

	json.NewEncoder(w).Encode(result)
}


func PostChannel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var res model.ResponseResult

	// Validate
	channel, violations := model.ValidatePostChannel(r)

	if len(violations) > 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(violations)
		return
	}

	newID := repository.ChannelRepo.GetNextId()
	channel.ID = *newID

	postId := repository.ChannelRepo.Insert(channel)
	if postId == nil {
		res.Error = "Cannot insert channel"
		json.NewEncoder(w).Encode(res)
		return
	}

	res.Result = "ok"
	json.NewEncoder(w).Encode(res)
	return
}



func PutChannel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//vars := mux.Vars(r)
	//name := vars["id"]
	//
	//result := repository.SagaRepo.DeleteByName(name)
	//
	//if result == false {
	//	fmt.Print("delete saga fail", result)
	//
	//	var res model.ResponseResult
	//	res.Error = "No items removed"
	//	json.NewEncoder(w).Encode(res)
	//	return
	//}

	w.WriteHeader(http.StatusNoContent)
}


func DeleteChannel(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	//vars := mux.Vars(r)
	//name := vars["id"]
	//
	//result := repository.ChannelRepo.DeleteByName(name)
	//
	//if result == false {
	//	fmt.Print("delete saga fail", result)
	//
	//	var res model.ResponseResult
	//	res.Error = "No items removed"
	//	json.NewEncoder(w).Encode(res)
	//	return
	//}

	w.WriteHeader(http.StatusNoContent)
}
