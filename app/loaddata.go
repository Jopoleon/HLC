package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"sync"

	"regexp"

	"github.com/sirupsen/logrus"
)

type User struct {
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	BirthDate int    `json:"birth_date,omitempty"`
	Gender    string `json:"gender,omitempty"`
	ID        int    `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
}
type Users struct {
	Users []User `json:"users"`
}
type Location struct {
	Distance int    `json:"distance,omitempty"`
	City     string `json:"city,omitempty"`
	Place    string `json:"place,omitempty"`
	ID       int    `json:"id,omitempty"`
	Country  string `json:"country,omitempty"`
}
type Locations struct {
	Locations []Location `json:"locations"`
}

type Visit struct {
	User      int `json:"user,omitempty"`
	Location  int `json:"location,omitempty"`
	VisitedAt int `json:"visited_at,omitempty"`
	ID        int `json:"id,omitempty"`
	Mark      int `json:"mark,omitempty"`
}

type Visits struct {
	Visits []Visit `json:"visits"`
}

////////

/////

/////

//LocationsAvg is strcut for handling requests for average mark of location
type LocationAvg struct {
	sync.Mutex
	ID       int        `json:"id"`
	Distance int        `json:"distance"`
	City     string     `json:"city"`
	Place    string     `json:"place"`
	Country  string     `json:"country"`
	Visits   []VisitAvg `json:"visitavg"`
}
type VisitAvg struct {
	UserID    int     `json:"userid"`
	VisitedAt int     `json:"visited_at"`
	ID        int     `json:"id"`
	LocID     int     `json:"id"`
	Mark      int     `json:"mark"`
	User      UserAvg `json:"user"`
}
type UserAvg struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate int    `json:"birth_date"`
	Gender    string `json:"gender"`
	ID        int    `json:"id"`
	Email     string `json:"email"`
}

func MakeLocationFull(u *Users, v *Visits, l *Locations) []LocationAvg {
	res := make([]LocationAvg, len(l.Locations))
	for i, loc := range l.Locations {
		res[i] = LocationAvg{
			ID:       loc.ID,
			City:     loc.City,
			Country:  loc.Country,
			Place:    loc.Place,
			Distance: loc.Distance,
			Visits:   []VisitAvg{},
		}
		for _, vis := range v.Visits {
			if res[i].ID == vis.User {
				res[i].Visits = append(res[i].Visits, VisitAvg{
					ID:        vis.ID,
					UserID:    vis.User,
					VisitedAt: vis.VisitedAt,
					LocID:     vis.Location,
					Mark:      vis.Mark,
					User:      UserAvg{},
				})
			}
		}
		for j, vf := range res[i].Visits {
			for _, user := range u.Users {
				if user.ID == vf.UserID {
					//logrus.Println(loc)
					res[i].Visits[j].User = UserAvg{
						ID:        user.ID,
						BirthDate: user.BirthDate,
						Email:     user.Email,
						Gender:    user.Gender,
						LastName:  user.LastName,
					}
				}
			}
		}
	}
	return res
}

//UserFull Struct for getting visits of user
type UserFull struct {
	sync.Mutex
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BirthDate int    `json:"birth_date"`
	Gender    string `json:"gender"`
	ID        int    `json:"id"`
	Email     string `json:"email"`
	Visits    []VisitFull
}
type VisitFull struct {
	User      int          `json:"user"`
	VisitedAt int          `json:"visited_at"`
	ID        int          `json:"id"`
	LocID     int          `json:"id"`
	Mark      int          `json:"mark"`
	Location  LocationFull `json:"location"`
}
type LocationFull struct {
	ID       int    `json:"id"`
	Distance int    `json:"distance"`
	City     string `json:"city"`
	Place    string `json:"place"`
	Country  string `json:"country"`
}

func MakeUserFull(u *Users, v *Visits, l *Locations) []UserFull {
	res := make([]UserFull, len(u.Users))
	for i, user := range u.Users {
		res[i] = UserFull{
			ID:        user.ID,
			Email:     user.Email,
			BirthDate: user.BirthDate,
			FirstName: user.FirstName,
			Gender:    user.Gender,
			LastName:  user.LastName,
			Visits:    []VisitFull{},
		}
		for _, vis := range v.Visits {
			if user.ID == vis.User {
				res[i].Visits = append(res[i].Visits, VisitFull{
					User:      user.ID,
					VisitedAt: vis.VisitedAt,
					Mark:      vis.Mark,
					LocID:     vis.Location,
					Location:  LocationFull{},
				})
			}
		}
		for j, vf := range res[i].Visits {
			for _, loc := range l.Locations {
				if vf.LocID == loc.ID {
					//logrus.Println(loc)
					res[i].Visits[j].Location = LocationFull{ID: loc.ID, City: loc.City, Country: loc.Country, Distance: loc.Distance, Place: loc.Place}
				}
			}
		}
	}
	return res
}

var UserInfo *Users
var VisitInfo *Visits
var LocsInfo *Locations

func LoadUsersToMemory() (*Users, error) {
	var users *Users
	dataDir, err := ioutil.ReadDir("data/TRAIN/data/")
	if err != nil {
		return users, fmt.Errorf("Error while reading file with users: %v", err)
	}
	b, err := ioutil.ReadFile("data/TRAIN/data/users_1.json")
	if err != nil {
		return users, fmt.Errorf("Error while reading file with users: %v", err)
	}
	err = json.Unmarshal(b, &UserInfo)
	if err != nil {
		return users, fmt.Errorf("Error while json.Unmarshal users: %v", err)
	}
	for _, f := range dataDir {
		logrus.Info(f.Name())
		match, err := regexp.MatchString("users_*.json", f.Name())
		if err != nil {
			return users, fmt.Errorf("Error while filepath.Match users file in data folder: %v", err)
		}
		if match {
			um := make(map[string][]User)
			logrus.Info(f.Name(), " matches users_*.json")
			b, err := ioutil.ReadFile("data/TRAIN/data/users_1.json")
			if err != nil {
				return users, fmt.Errorf("Error while reading file with users: %v", err)
			}
			err = json.Unmarshal(b, &um)
			if err != nil {
				return users, fmt.Errorf("Error while json.Unmarshal users: %v", err)
			}
			users.Users = append(users.Users, um["Users"]...)
		}
	}

	logrus.Info("LoadUsersToMemory done")
	return users, nil
}

func LoadVisitsToMemory() (*Visits, error) {
	var visits *Visits
	b, err := ioutil.ReadFile("data/TRAIN/data/visits_1.json")
	if err != nil {
		return visits, fmt.Errorf("Error while reading file with visits: %v", err)
	}
	err = json.Unmarshal(b, &VisitInfo)
	if err != nil {
		return visits, fmt.Errorf("Error while json.Unmarshal visits: %v", err)
	}
	logrus.Info("LoadVisitsToMemory done")
	return visits, nil
}

func LoadLocationsToMemory() (*Locations, error) {
	var locs *Locations
	b, err := ioutil.ReadFile("data/TRAIN/data/locations_1.json")
	if err != nil {
		return locs, fmt.Errorf("Error while reading file with locs: %v", err)
	}
	err = json.Unmarshal(b, &LocsInfo)
	if err != nil {
		return locs, fmt.Errorf("Error while json.Unmarshal locs: %v", err)
	}
	logrus.Info("LoadLocationsToMemory done")
	return locs, nil
}
