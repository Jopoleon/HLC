package app

//{"visits": [{"mark": 3, "visited_at": 1064315078, "place": "\u0424\u043e\u043d\u0430\u0440\u044c"}

//New filters for maps of visits
func (vs *VisIDs) FromDate(fromDate int) VisIDs {
	var res VisIDs
	for _, v := range vs.IDs {
		if VisID_Vis.Map[v].VisitedAt > fromDate {
			res.IDs = append(res.IDs, v)
		}
	}
	return res
}
func (vs *VisIDs) ToDate(toDate int) VisIDs {
	var res VisIDs
	for _, v := range vs.IDs {
		if VisID_Vis.Map[v].VisitedAt < toDate {
			res.IDs = append(res.IDs, v)
		}
	}
	return res
}
func (vs *VisIDs) ToDistance(toDistance int) VisIDs {
	var res VisIDs
	for _, v := range vs.IDs {
		if LocID_Loc.Map[VisID_Vis.Map[v].Location].Distance < toDistance {
			res.IDs = append(res.IDs, v)
		}
	}
	return res
}
func (vs *VisIDs) Country(country string) VisIDs {
	var res VisIDs
	for _, v := range vs.IDs {
		if LocID_Loc.Map[VisID_Vis.Map[v].Location].Country == country {
			res.IDs = append(res.IDs, v)
		}
	}
	return res
}

//`fromAge` - учитывать только путешественников, у которых возраст (считается от текущего timestamp) строго больше этого параметра
//`toAge` -  учитывать только путешественников, у которых возраст (считается от текущего timestamp) строго меньше этого параметра
//`gender` - учитывать оценки только мужчин или женщин
func (vs *VisIDs) FromAge(fromAge int64) VisIDs {
	var res VisIDs
	for _, v := range vs.IDs {
		if int64(UserID_User.Map[VisID_Vis.Map[v].User].BirthDate) < fromAge {
			res.IDs = append(res.IDs, v)
		}
	}
	return res
}
func (vs *VisIDs) ToAge(toAge int64) VisIDs {
	var res VisIDs
	for _, v := range vs.IDs {
		//logrus.Infof("%+v, Age: %v", UserID_User[v.User], toAge)
		if int64(UserID_User.Map[VisID_Vis.Map[v].User].BirthDate) > toAge {
			res.IDs = append(res.IDs, v)
		}
	}
	return res
}
func (vs *VisIDs) Gender(gender string) VisIDs {
	var res VisIDs
	for _, v := range vs.IDs {
		if UserID_User.Map[VisID_Vis.Map[v].User].Gender == gender {
			res.IDs = append(res.IDs, v)
		}
	}
	return res
}

////////(*((((((((((((((&&&&&&&&&&&&&&&&&&----------/////////////////

//

//

//
func (vs *Visits) FromDate(fromDate int) Visits {
	var res Visits
	for _, v := range vs.Visits {
		if v.VisitedAt > fromDate {
			res.Visits = append(res.Visits, v)
		}
	}
	return res
}
func (vs *Visits) ToDate(toDate int) Visits {
	var res Visits
	for _, v := range vs.Visits {
		if v.VisitedAt < toDate {
			res.Visits = append(res.Visits, v)
		}
	}
	return res
}
func (vs *Visits) ToDistance(toDistance int) Visits {
	var res Visits
	for _, v := range vs.Visits {
		if VisID_Loc[v.ID].Distance < toDistance {
			res.Visits = append(res.Visits, v)
		}
	}
	return res
}
func (vs *Visits) Country(country string) Visits {
	var res Visits
	for _, v := range vs.Visits {
		if VisID_Loc[v.ID].Country == country {
			res.Visits = append(res.Visits, v)
		}
	}
	return res
}

//`fromAge` - учитывать только путешественников, у которых возраст (считается от текущего timestamp) строго больше этого параметра
//`toAge` -  учитывать только путешественников, у которых возраст (считается от текущего timestamp) строго меньше этого параметра
//`gender` - учитывать оценки только мужчин или женщин
//func (vs *Visits) FromAge(fromAge int64) Visits {
//	var res Visits
//	for _, v := range vs.Visits {
//		if int64(UserID_User[v.User].BirthDate) < fromAge {
//			res.Visits = append(res.Visits, v)
//		}
//	}
//	return res
//}
//func (vs *Visits) ToAge(toAge int64) Visits {
//	var res Visits
//	for _, v := range vs.Visits {
//		logrus.Infof("%+v, Age: %v", UserID_User[v.User], toAge)
//		if int64(UserID_User[v.User].BirthDate) > toAge {
//			res.Visits = append(res.Visits, v)
//		}
//	}
//	return res
//}
//func (vs *Visits) Gender(gender string) Visits {
//	var res Visits
//	for _, v := range vs.Visits {
//		if UserID_User[v.User].Gender == gender {
//			res.Visits = append(res.Visits, v)
//		}
//	}
//	return res
//}

//func bFromDate(vs []Visit, FromDate int) []Visit {
//	var res []Visit
//	for _, v := range vs {
//		if v.VisitedAt > FromDate {
//			res = append(res, v)
//		}
//	}
//	return res
//}
//func bToDate(vs []Visit, toDate int) []Visit {
//	var res []Visit
//	for _, v := range vs {
//		if v.VisitedAt < toDate {
//			res = append(res, v)
//		}
//	}
//	return res
//}
//func bToDistance(vs []Visit, toDistance int) []Visit {
//	var res []Visit
//	for _, v := range vs {
//		if v.Distance < toDistance {
//			res = append(res, v)
//		}
//	}
//	return res
//}
//func bCountry(vs []Visit, country string) []Visit {
//	var res []Visit
//	for _, v := range vs {
//		if v.Location.Country == country {
//			res = append(res, v)
//		}
//	}
//	return res
//}

//Filters for visit endpoint
func byFromDate(vs []VisitFull, fromDate int) []VisitFull {
	var res []VisitFull
	for _, v := range vs {
		if v.VisitedAt > fromDate {
			res = append(res, v)
		}
	}
	return res
}
func byToDate(vs []VisitFull, toDate int) []VisitFull {
	var res []VisitFull
	for _, v := range vs {
		if v.VisitedAt < toDate {
			res = append(res, v)
		}
	}
	return res
}
func byToDistance(vs []VisitFull, toDistance int) []VisitFull {
	var res []VisitFull
	for _, v := range vs {
		if v.Location.Distance < toDistance {
			res = append(res, v)
		}
	}
	return res
}
func byCountry(vs []VisitFull, country string) []VisitFull {
	var res []VisitFull
	for _, v := range vs {
		if v.Location.Country == country {
			res = append(res, v)
		}
	}
	return res
}

//Filters for location avg endpoint
//FromDate, toDate, fromAge, toAge, gender
func byFromDateL(vs []VisitAvg, fromDate int) []VisitAvg {
	var res []VisitAvg
	for _, v := range vs {
		if v.VisitedAt > fromDate {
			res = append(res, v)
		}
	}
	return res
}
func byToDateL(vs []VisitAvg, toDate int) []VisitAvg {
	var res []VisitAvg
	for _, v := range vs {
		if v.VisitedAt < toDate {
			res = append(res, v)
		}
	}
	return res
}
func byFromAgeL(vs []VisitAvg, fromAge int) []VisitAvg {
	var res []VisitAvg
	for _, v := range vs {
		if v.User.BirthDate >= fromAge {
			res = append(res, v)
		}
	}
	return res
}
func byToAgeL(vs []VisitAvg, toAge int) []VisitAvg {
	var res []VisitAvg
	for _, v := range vs {
		if v.User.BirthDate <= toAge {
			res = append(res, v)
		}
	}
	return res
}
func byGenderL(vs []VisitAvg, gender string) []VisitAvg {
	var res []VisitAvg
	for _, v := range vs {
		if v.User.Gender == gender {
			res = append(res, v)
		}
	}
	return res
}
