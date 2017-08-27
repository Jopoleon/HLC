package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
)

func PostNewEntity(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	entity := query.Get(":entity")
	switch entity {
	case "users":
		PostNewUser(w, req)
	case "locations":
		PostNewLocation(w, req)
	case "visits":
		PostNewVisit(w, req)
	default:
		http.Error(w, "", 400)
	}
}

func PostNewUser(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if strings.Contains(string(b), "null") || len(b) < 5 {
		logrus.Error("PostNewUser request body contains 'null' or empty", string(b))
		w.WriteHeader(400)
		return
	}
	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		logrus.Errorf("PostNewUser json.Unmarshal request body error: %v, body: %v", err, string(b))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if user.Gender != "" && user.BirthDate != 0 && user.FirstName != "" && user.LastName != "" && user.Email != "" && user.ID != 0 {
		//logrus.Infof("PostNewUser used with user: %+v ", user)
		existUser, uok := UserID_User.Map[user.ID]
		if uok {
			logrus.Errorf("PostNewVisit locations with such id already exists: %+v", existUser)
			w.WriteHeader(400)
			return
		}
		UserID_User.Lock()
		UserID_User.Map[user.ID] = user
		UserID_User.Unlock()
		//UserID_IDVis
		w.WriteHeader(200)
		w.Write([]byte("{}"))
		return
	}
	logrus.Errorf("PostNewUser some fields are empty: %+v", user)
	w.WriteHeader(400)
	return
}
func PostNewVisit(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if strings.Contains(string(b), "null") || len(b) < 5 {
		logrus.Error("PostNewVisit request body contains 'null' or empty", string(b))
		w.WriteHeader(400)
		return
	}
	var vis Visit
	err = json.Unmarshal(b, &vis)
	if err != nil {
		logrus.Errorf("PostNewVisit json.Unmarshal request body error: %v, body: %v", err, string(b))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if vis.ID != 0 && vis.Location != 0 && vis.VisitedAt != 0 && vis.User != 0 && vis.Mark != 0 {
		//logrus.Infof("PostNewVisit used with user: %+v ", vis)
		existVis, uok := VisID_Vis.Map[vis.ID]
		if uok {
			logrus.Errorf("PostNewVisit locations with such id already exists: %+v", existVis)
			w.WriteHeader(400)
			return
		}
		VisID_Vis.Lock()
		VisID_Vis.Map[vis.ID] = vis
		UserID_IDVis[vis.User] = append(UserID_IDVis[vis.User], vis.ID)
		LocID_IDVis[vis.Location] = append(LocID_IDVis[vis.Location], vis.ID)
		VisID_Vis.Unlock()
		//UserID_IDVis
		w.WriteHeader(200)
		w.Write([]byte("{}"))
		return
	}
	logrus.Errorf("PostNewVisit some fields are empty: %+v", vis)
	w.WriteHeader(400)
	return
}
func PostNewLocation(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if strings.Contains(string(b), "null") || len(b) < 5 {
		logrus.Error("PostNewLocation request body contains 'null' or empty", string(b))
		w.WriteHeader(400)
		return
	}
	var loc Location
	err = json.Unmarshal(b, &loc)
	if err != nil {
		logrus.Errorf("PostNewLocation json.Unmarshal request body error: %v, body: %v", err, string(b))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if loc.ID != 0 && loc.City != "" && loc.Place != "" && loc.Country != "" && loc.Distance != 0 {
		//logrus.Infof("PostNewLocation used with user: %+v ", loc)
		existLoc, uok := LocID_Loc.Map[loc.ID]
		if uok {
			logrus.Errorf("PostNewLocation locations with such id already exists: %+v", existLoc)
			w.WriteHeader(400)
			return
		}
		LocID_Loc.Lock()
		LocID_Loc.Map[loc.ID] = loc
		LocID_Loc.Unlock()
		//UserID_IDVis
		w.WriteHeader(200)
		w.Write([]byte("{}"))
		return
	}
	logrus.Errorf("PostNewLocation some fields are empty: %+v", loc)
	w.WriteHeader(400)
	return
}

//PostUpdateEntity
func PostUpdateEntity(w http.ResponseWriter, req *http.Request) {

	query := req.URL.Query()
	entity := query.Get(":entity")
	switch entity {
	case "users":
		PostUpdateUser(w, req)
	case "locations":
		PostUpdateLocation(w, req)
	case "visits":
		PostUpdateVisit(w, req)
	default:
		http.Error(w, "", 400)
	}
}
func PostUpdateUser(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	urlId := query.Get(":id")
	//logrus.Info("URL query: ", query)
	userID, err := strconv.Atoi(urlId)
	if err != nil {
		http.Error(w, "", 400)
		return
	}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if strings.Contains(string(b), "null") || len(b) < 5 {
		logrus.Error("PostUpdateUser request body contains 'null' or empty", string(b))
		w.WriteHeader(400)
		return
	}
	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		logrus.Errorf("PostUpdateUser json.Unmarshal request body error: %v, body: %v", err, string(b))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//logrus.Infof("PostUpdateUser used with user: %+v ", user)
	oldUser, uok := UserID_User.Map[userID]
	if !uok {
		logrus.Errorf("PostUpdateUser user with id %v not found.", userID)
		w.WriteHeader(404)
		return
	}
	if user.Gender != "" && user.Gender != "null" {
		oldUser.Gender = user.Gender
	}
	if user.BirthDate != 0 {
		oldUser.BirthDate = user.BirthDate
	}
	if user.FirstName != "" && user.FirstName != "null" {
		oldUser.FirstName = user.FirstName
	}
	if user.LastName != "" && user.LastName != "null" {
		oldUser.LastName = user.LastName
	}
	if user.Email != "" && user.Email != "null" {
		oldUser.Email = user.Email
	}
	UserID_User.Lock()
	UserID_User.Map[userID] = oldUser
	UserID_User.Unlock()
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}
func PostUpdateVisit(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	urlId := query.Get(":id")
	//logrus.Info("PostUpdateVisit URL query: ", query)
	visID, err := strconv.Atoi(urlId)
	if err != nil {
		http.Error(w, "", 400)
		return
	}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//logrus.Info("PostUpdateVisit URL query: ", query)
	if strings.Contains(string(b), "null") || len(b) < 5 {
		logrus.Error("PostUpdateVisit body contains 'null'", string(b))
		w.WriteHeader(400)
		return
	}
	var vis Visit
	err = json.Unmarshal(b, &vis)
	if err != nil {
		logrus.Errorf("PostUpdateVisit json.Unmarshal request body error: %v, body: %v", err, string(b))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//logrus.Infof("PostUpdateVisit used with visit: %+v ", vis)
	oldVis, uok := VisID_Vis.Map[visID]
	if !uok {
		logrus.Errorf("PostUpdateVisit visit with id %v not found.", visID)
		w.WriteHeader(404)
		return
	}
	if vis.Mark != 0 {
		oldVis.Mark = vis.Mark
	}
	if vis.Location != 0 {
		oldVis.Location = vis.Location
	}
	if vis.User != 0 {
		oldVis.User = vis.User
	}
	if vis.VisitedAt != 0 {
		oldVis.VisitedAt = vis.VisitedAt
	}
	VisID_Vis.Lock()
	VisID_Vis.Map[visID] = oldVis
	VisID_Vis.Unlock()
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}
func PostUpdateLocation(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	urlId := query.Get(":id")
	//logrus.Info("PostUpdateLocation URL query: ", query)
	locID, err := strconv.Atoi(urlId)
	if err != nil {
		http.Error(w, "", 400)
		return
	}

	b, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if strings.Contains(string(b), "null") || len(b) < 5 {
		logrus.Error("PostUpdateLocation body contains 'null' or empty", string(b))
		w.WriteHeader(400)
		return
	}
	//logrus.Info("PostUpdateLocation body: ", string(b))

	var loc Location
	err = json.Unmarshal(b, &loc)
	if err != nil {
		logrus.Errorf("PostUpdateLocation json.Unmarshal request body error: %v, body: %v", err, string(b))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//logrus.Infof("PostUpdateLocation used with loc: %+v ", loc)
	oldLoc, uok := LocID_Loc.Map[locID]
	if !uok {
		logrus.Errorf("PostUpdateLocation locations with id %v not found.", locID)
		w.WriteHeader(404)
		return
	}
	if loc.Distance != 0 {
		oldLoc.Distance = loc.Distance
	}
	if loc.Country != "" && loc.Country != "null" {
		oldLoc.Country = loc.Country
	}
	if loc.Place != "" && loc.Place != "null" {
		oldLoc.Place = loc.Place
	}
	if loc.City != "" && loc.City != "null" {
		oldLoc.City = loc.City
	}
	LocID_Loc.Lock()
	LocID_Loc.Map[locID] = oldLoc
	LocID_Loc.Unlock()
	w.WriteHeader(200)
	w.Write([]byte("{}"))
}
