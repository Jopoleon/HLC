package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"io/ioutil"

	"github.com/sirupsen/logrus"
)

func GetUsersVisits_old(w http.ResponseWriter, req *http.Request) {
	// FromDate - посещения с visited_at > FromDate
	// toDate - посещения до visited_at < toDate
	// country - название страны, в которой находятся интересующие достопримечательности
	// toDistance - возвращать только те места, у которых расстояние от города меньше этого параметра
	query := req.URL.Query()
	urlId := query.Get(":id")
	logrus.Info("URL query: ", query)
	userID, err := strconv.Atoi(urlId)
	if err != nil {
		http.Error(w, "Bad_URL_ID", 404)
		return
	}
	var fromDate, toDate, toDistance int
	var country string
	fd, okfd := query["FromDate"]
	if okfd {
		fromDate, err = strconv.Atoi(fd[0])
		if err != nil {
			http.Error(w, "Bad_URL_ID", 404)
			return
		}
	}
	td, oktd := query["toDate"]
	if oktd {
		toDate, err = strconv.Atoi(td[0])
		if err != nil {
			http.Error(w, "Bad_URL_ID", 404)
			return
		}
	}
	tdi, oktdi := query["toDistance"]
	if oktdi {
		toDistance, err = strconv.Atoi(tdi[0])
		if err != nil {
			http.Error(w, "Bad_URL_ID", 404)
			return
		}
	}
	c, okc := query["country"]
	if okc {
		country = c[0]
	}
	//userID && u.VisitedAt > FromDate && u.VisitedAt < toDate && u.toDate

	// FromDate - посещения с visited_at > FromDate
	// toDate - посещения до visited_at < toDate
	// country - название страны, в которой находятся интересующие достопримечательности
	// toDistance - возвращать только те места, у которых расстояние от города меньше этого параметра
	var res []VisitFull
	var answer RespVisits
	for _, u := range UsersFull {
		if u.ID == userID {
			res = u.Visits
			if okfd {
				res = byFromDate(res, fromDate)
			}
			if oktd {
				res = byToDate(res, toDate)
			}
			if oktdi {
				res = byToDistance(res, toDistance)
			}
			if okc {
				res = byCountry(res, country)
			}
			for _, v := range res {
				answer.RespVisits = append(answer.RespVisits, RespVisit{Place: v.Location.Place, Mark: v.Mark, Visited_at: v.VisitedAt})
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
	}
	//return
	//toDistance := req.URL.Query().Get("toDistance")
	logrus.Println(userID, fromDate, toDate, country, toDistance)
	fmt.Fprintf(w, "%v %v %v %v %v", userID, fromDate, toDate, country, toDistance)

}

func GetLocationAvg_old(w http.ResponseWriter, req *http.Request) {
	//`FromDate` - учитывать оценки только с `visited_at` > `FromDate`
	//`toDate` - учитывать оценки только до `visited_at` < `toDate`
	//`fromAge` - учитывать только путешественников, у которых возраст (считается от текущего timestamp) строго больше этого параметра
	//`toAge` -  учитывать только путешественников, у которых возраст (считается от текущего timestamp) строго меньше этого параметра
	//`gender` - учитывать оценки только мужчин или женщин
	query := req.URL.Query()
	urlId := query.Get(":id")
	logrus.Info("URL query: ", query)
	locID, err := strconv.Atoi(urlId)
	if err != nil {
		http.Error(w, "", 400)
		return
	}
	var fromDate, toDate, fromAge, toAge int
	var gender string
	fd, okfd := query["FromDate"]
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
	tdi, okfa := query["fromAge"]
	if okfa {
		fromAge, err = strconv.Atoi(tdi[0])
		if err != nil {
			http.Error(w, "", 400)
			return
		}
	}
	tdi, okta := query["toAge"]
	if okta {
		toAge, err = strconv.Atoi(tdi[0])
		if err != nil {
			http.Error(w, "", 400)
			return
		}
	}
	c, okg := query["gender"]
	if okg {
		gender = c[0]

	}
	var res []VisitAvg
	var sum int

	for _, l := range LocationsAvg {
		if l.ID == locID {
			res = l.Visits
			if okfd {
				res = byFromDateL(res, fromDate)
			}
			if oktd {
				res = byToDateL(res, toDate)
			}
			if okfa {
				res = byFromAgeL(res, fromAge)
			}
			if okta {
				res = byToAgeL(res, toAge)
			}
			if okg {
				res = byGenderL(res, gender)
			}
			for _, v := range res {
				sum = sum + v.Mark
			}
			//logrus.Infof("result: %+v", res)
			//logrus.Infof("Mark sum: %+v", sum)
			//logrus.Infof("Length of res: %+v", len(res))
			//logrus.Infof("Avg mark: : %+v", sum/len(res))
			//logrus.Infof("After round by 5 position: %+v", Round(float64(sum/len(res)), .5, 5))
			b, err := json.Marshal(AvgMark{
			//Avg: Round(float64(sum/len(res)), .5, 5),
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write(b)
			return
		}
	}
	fmt.Fprintf(w, "%v %v %v %v %v", fromDate, toDate, fromAge, toAge, gender)
}
func PostUpdateUser_old(w http.ResponseWriter, req *http.Request) {
	query := req.URL.Query()
	urlId := query.Get(":id")
	logrus.Info("URL query: ", query)
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
	var user User
	err = json.Unmarshal(b, &user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logrus.Infof("PostUpdateUser used with user: %+v ", user)
	if user.BirthDate != 0 && user.Email != "" &&
		user.FirstName != "" && user.LastName != "" {

	} else {
		http.Error(w, "", 400)
		return
	}
	for i, u := range UsersFull {
		if u.ID == userID {
			logrus.Infof("User %v found:\n %+v", u.ID, u)
			UsersFull[i].Lock()
			UsersFull[i] = UserFull{
				Email:     user.Email,
				LastName:  user.LastName,
				FirstName: user.FirstName,
				BirthDate: user.BirthDate,
			}
			UsersFull[i].Unlock()

		} else {
			http.Error(w, "", 404)
			return
		}
	}
	for i, loc := range LocationsAvg {
		for j, v := range loc.Visits {
			if v.UserID == userID {
				LocationsAvg[i].Lock()
				LocationsAvg[i].Visits[j].User = UserAvg{
					Email:     user.Email,
					LastName:  user.LastName,
					FirstName: user.FirstName,
					BirthDate: user.BirthDate,
				}
				LocationsAvg[i].Unlock()
			} else {
				http.Error(w, "", 404)
				return
			}
		}

	}
}
