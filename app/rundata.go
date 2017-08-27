package app

import (
	"time"

	"github.com/sirupsen/logrus"
)

var UserLocs map[User][]Location
var UsersFull []UserFull
var LocationsAvg []LocationAvg

func Run() error {
	start := time.Now()
	var err error
	_, err = LoadUsersToMemory()
	if err != nil {
		return err
	}
	utm := time.Now()
	_, err = LoadVisitsToMemory()
	if err != nil {
		return err
	}
	vtm := time.Now()
	_, err = LoadLocationsToMemory()
	if err != nil {
		return err
	}
	ltm := time.Now()
	//UserLocs = mapUsersLocations(UserInfo, VisitInfo, LocsInfo)
	InitMaps(UserInfo, VisitInfo, LocsInfo)
	iintm := time.Now()
	UsersFull = MakeUserFull(UserInfo, VisitInfo, LocsInfo)
	LocationsAvg = MakeLocationFull(UserInfo, VisitInfo, LocsInfo)

	logrus.Infof("\n LoadUsersToMemory time: %v	\n LoadVisitsToMemory time: %v \n LoadLocationsToMemory time %v \n InitMaps time: %v", start.Sub(utm), utm.Sub(vtm), vtm.Sub(ltm), ltm.Sub(iintm))
	return nil
}
