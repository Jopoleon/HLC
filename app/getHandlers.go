package app

import (
	"encoding/json"
	"math"
	"net/http"
	"sort"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

func GetUsersVisits(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	urlId := query.Get(":id")
	//logrus.Info("URL query: ", query)
	userID, err := strconv.Atoi(urlId)
	if err != nil {
		http.Error(w, "", 400)
		return
	}
	var fromDate, toDate, toDistance int
	var country string
	fd, okfd := query["fromDate"]
	if okfd {
		fromDate, err = strconv.Atoi(fd[0])
		if err != nil {
			http.Error(w, "", 400)
			return
		}
	}
	td, oktd := query["toDate"]
	if oktd {
		toDate, err = strconv.Atoi(td[0])
		if err != nil {
			http.Error(w, "", 400)
			return
		}
	}
	tdi, oktdi := query["toDistance"]
	if oktdi {
		toDistance, err = strconv.Atoi(tdi[0])
		if err != nil {
			http.Error(w, "", 400)
			return
		}
	}
	c, okc := query["country"]
	if okc {
		country = c[0]

	}
	_, idOk := UserID_User.Map[userID]
	if !idOk {
		http.Error(w, "", 404)
	}
	var res VisIDs
	res.IDs = make([]int, 0)
	var answer RespVisits
	visIds, vsok := UserID_IDVis[userID]
	if !vsok {
		http.Error(w, "", 404)
	}
	res.IDs = visIds
	if okfd {
		res = res.FromDate(fromDate)
	}
	if oktd {
		res = res.ToDate(toDate)
	}
	if oktdi {
		res = res.ToDistance(toDistance)
	}
	if okc {
		res = res.Country(country)
	}
	logrus.Infof("GetUsersVisits Ids after all filters %+v", res)
	for _, v := range res.IDs {
		answer.RespVisits = append(answer.RespVisits, RespVisit{
			Mark:       VisID_Vis.Map[v].Mark,
			Visited_at: VisID_Vis.Map[v].VisitedAt,
			Place:      LocID_Loc.Map[VisID_Vis.Map[v].Location].Place,
		})
	}
	if len(res.IDs) == 0 {
		var emptyAnswer RespVisits
		emptyAnswer.RespVisits = make([]RespVisit, 0)
		b, err := json.Marshal(emptyAnswer)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(200)
		w.Write(b)
		return
	}
	sort.Sort(answer)
	b, err := json.Marshal(answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
	return
}

type AvgMark struct {
	Avg float64 `json:"avg"`
}

func GetLocationAvg(w http.ResponseWriter, req *http.Request) {
	//`FromDate` - учитывать оценки только с `visited_at` > `FromDate`
	//`toDate` - учитывать оценки только до `visited_at` < `toDate`
	//`fromAge` - учитывать только путешественников, у которых возраст (считается от текущего timestamp) строго больше этого параметра
	//`toAge` -  учитывать только путешественников, у которых возраст (считается от текущего timestamp) строго меньше этого параметра
	//`gender` - учитывать оценки только мужчин или женщин
	query := req.URL.Query()
	Idstr := query.Get(":id")
	logrus.Info("GetLocationAvg URL query: ", query)
	locID, err := strconv.Atoi(Idstr)
	if err != nil {
		http.Error(w, "", 400)
		return
	}
	var fromDate, toDate int
	var fromAge, toAge int64
	var gender string
	fd, okfd := query["fromDate"]
	if okfd {
		fromDate, err = strconv.Atoi(fd[0])
		if err != nil {
			http.Error(w, "", 400)
			return
		}
		logrus.Info("GetLocationAvg fromDate query: ", fromDate)
	}
	td, oktd := query["toDate"]
	if oktd {
		toDate, err = strconv.Atoi(td[0])
		if err != nil {
			http.Error(w, "", 400)
			return
		}
	}
	fa, okfa := query["fromAge"]
	if okfa {
		af, err := strconv.Atoi(fa[0])
		if err != nil {
			//logrus.Error("GetLocationAvg fromAge error:", query["fromAge"])
			http.Error(w, "", 400)
			return
		}
		fromAge = time.Now().UTC().AddDate(af*(-1), 0, 0).Unix()
	}
	tag, okta := query["toAge"]
	if okta {
		ta, err := strconv.Atoi(tag[0])
		if err != nil {
			//logrus.Error("GetLocationAvg toAge error:", query["toAge"])
			http.Error(w, "", 400)
			return
		}

		toAge = time.Now().UTC().AddDate(ta*(-1), 0, 0).Unix()
	}
	ge, okg := query["gender"]
	if okg {
		gender = ge[0]
		if len(gender) > 2 {
			//logrus.Error("GetLocationAvg gender error:", query["gender"], gender)
			http.Error(w, "", 400)
			return
		}
	}

	//var res Visits
	var sum int
	//lVis, lok := LocID_Visits[locID]
	//if !lok {
	//	http.Error(w, "", 404)
	//}
	var res VisIDs
	visIds, vsok := LocID_IDVis[locID]
	if !vsok {
		http.Error(w, "", 404)
		return
	}
	// дата рождения дожна быть меньше, если считать ОтДата
	// от даты должно быть меньше
	res.IDs = visIds
	//logrus.Infof("GetLocationAvg res: %+v", res)
	if okfd {
		res = res.FromDate(fromDate)
		//logrus.Infof("GetLocationAvg res AFTER FromDate: %+v", res)
	}
	if oktd {
		res = res.ToDate(toDate)
		//logrus.Infof("GetLocationAvg res AFTER ToDate: %+v", res)
	}
	if okfa {
		res = res.FromAge(fromAge)
		//logrus.Infof("GetLocationAvg res AFTER FromAge: %+v", res)
	}
	if okta {
		res = res.ToAge(toAge)
		//logrus.Infof("GetLocationAvg res AFTER ToAge: %+v", res)
	}
	if okg {
		res = res.Gender(gender)
		//logrus.Infof("GetLocationAvg res AFTER Gender: %+v", res)
	}
	//logrus.Infof("GetLocationAvg res AFTER filters: %+v", res)
	if len(res.IDs) == 0 {
		b, err := json.Marshal(AvgMark{
			Avg: 0.0,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(200)
		w.Write(b)
		return
	}

	for _, v := range res.IDs {
		sum = sum + VisID_Vis.Map[v].Mark
	}
	logrus.Infof("result: %+v", res)
	logrus.Infof("Mark sum: %+v", sum)
	logrus.Infof("Length of res: %+v", len(res.IDs))
	logrus.Infof("Avg mark: : %+v", float64(sum)/float64(len(res.IDs)))
	logrus.Infof("After round by 5 position: %+v", Round(float64(sum)/float64(len(res.IDs)), .1, 5))
	logrus.Infof("After round by New function 5: %+v", toFixed(float64(sum)/float64(len(res.IDs)), 5))
	b, err := json.Marshal(AvgMark{
		Avg: toFixed(float64(sum)/float64(len(res.IDs)), 5),
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(b)
	return
}
func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

func GetUser(w http.ResponseWriter, req *http.Request) {
	urlId := req.URL.Query().Get(":id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		http.Error(w, "", 404)
		return
	}
	//logrus.Info("GetUser____:\n", req)
	logrus.Info("GetUser used with query : ", req.URL.Query())
	for _, u := range UserInfo.Users {
		if u.ID == id {
			b, err := json.Marshal(u)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(b)
			return
		}
	}
	http.Error(w, "", 404)
	return
}

func GetLocation(w http.ResponseWriter, req *http.Request) {
	urlId := req.URL.Query().Get(":id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		http.Error(w, "", 404)
		return
	}
	logrus.Info("GetLocation used with query : ", req.URL.Query())
	for _, u := range LocsInfo.Locations {
		if u.ID == id {
			b, err := json.Marshal(u)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(b)
			return
		}

	}
	http.Error(w, "", 404)
	return
}

func GetVisit(w http.ResponseWriter, req *http.Request) {
	urlId := req.URL.Query().Get(":id")
	id, err := strconv.Atoi(urlId)
	if err != nil {
		http.Error(w, "", 404)
		return
	}
	logrus.Info("GetVisit used with query : ", req.URL.Query())
	for _, u := range VisitInfo.Visits {
		if u.ID == id {
			b, err := json.Marshal(u)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(b)
			return
		}

	}
	http.Error(w, "", 404)
	return
}
