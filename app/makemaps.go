package app

import "sync"

var (
	UserID_User = struct {
		sync.Mutex
		Map map[int]User
	}{} //yes
	VisID_Vis = struct {
		sync.Mutex
		Map map[int]Visit
	}{} //yes

	LocID_Loc = struct {
		sync.Mutex
		Map map[int]Location
	}{}

	UserID_IDVis = make(map[int][]int)
	LocID_IDVis  = make(map[int][]int)

	UserID_Locs     = make(map[int][]Location)           //no
	UserID_Vis_Locs = make(map[int][]map[Visit]Location) //no
	Vis_Loc         = make(map[Visit]Location)           //no
	//no
	LocID_Users = make(map[int][]User) //no

	UserID_Visits = make(map[int][]Visit) //yes
	//yes
	//yes
	VisID_Loc    = make(map[int]Location) //yse
	LocID_Visits = make(map[int][]Visit)  //yse
	//yse
)

//type Locker interface {
//	Lock()
//	Unlock()
//}

//LocID_Loc
//UserID_User
//VisID_Vis
func UpdateUserID_IDVis() {
	UserID_User.Lock()
	for _, u := range UserID_User.Map {
		VisID_Vis.Lock()
		for _, v := range VisID_Vis.Map {
			if u.ID == v.User {
				UserID_IDVis[u.ID] = append(UserID_IDVis[u.ID], v.ID)
			}
		}
		VisID_Vis.Unlock()
	}
	UserID_User.Unlock()
}

func UpdateLocID_IDVis() {
	for _, l := range LocID_Loc.Map {
		for _, v := range VisID_Vis.Map {
			if l.ID == v.Location {
				LocID_IDVis[l.ID] = append(LocID_IDVis[l.ID], v.ID)
			}
		}
	}
}
func InitMaps(users *Users, visits *Visits, locs *Locations) {
	//UserID_User := make(map[int]User)
	UserID_User.Map = make(map[int]User)
	VisID_Vis.Map = make(map[int]Visit)
	LocID_Loc.Map = make(map[int]Location)
	for _, u := range users.Users {
		UserID_User.Map[u.ID] = u
	}
	//UserID_Visits := make(map[int][]Visit)
	for _, u := range users.Users {
		for _, v := range visits.Visits {
			if u.ID == v.User {
				UserID_Visits[u.ID] = append(UserID_Visits[u.ID], v)
				UserID_IDVis[u.ID] = append(UserID_IDVis[u.ID], v.ID)
			}
		}
	}
	//UserID_Locs := make(map[int][]Location)
	for _, u := range users.Users {
		for _, v := range visits.Visits {
			if u.ID == v.User {
				for _, l := range locs.Locations {
					if v.Location == l.ID {
						UserID_Locs[u.ID] = append(UserID_Locs[u.ID], l)
					}
				}
			}
		}
	}
	//UserID_Vis_Locs := make(map[int][]map[Visit]Location)
	for _, u := range users.Users {
		for _, v := range visits.Visits {
			if u.ID == v.User {
				for _, l := range locs.Locations {
					if v.Location == l.ID {
						UserID_Vis_Locs[u.ID] = append(UserID_Vis_Locs[u.ID], map[Visit]Location{v: l})
					}
				}
			}
		}
	}
	//VisID_Vis := make(map[int]Visit)
	for _, v := range visits.Visits {
		VisID_Vis.Map[v.ID] = v
	}
	//VisID_Loc := make(map[int]Visit)
	for _, v := range visits.Visits {
		for _, l := range locs.Locations {
			if l.ID == v.Location {
				VisID_Loc[l.ID] = l
			}
		}
	}
	//Vis_Loc := make(map[Visit]Location)
	for _, v := range visits.Visits {
		for _, l := range locs.Locations {
			if l.ID == v.Location {
				Vis_Loc[v] = l
			}
		}
	}
	//LocID_Loc := make(map[int]Location)
	for _, l := range locs.Locations {
		LocID_Loc.Map[l.ID] = l
	}
	//LocID_Visits := make(map[int][]Visit)
	for _, l := range locs.Locations {
		for _, v := range visits.Visits {
			if l.ID == v.Location {
				LocID_Visits[l.ID] = append(LocID_Visits[l.ID], v)
				LocID_IDVis[l.ID] = append(LocID_IDVis[l.ID], v.ID)
			}
		}
	}
	//LocID_Users := make(map[int][]User)
	for _, l := range locs.Locations {
		for _, v := range visits.Visits {
			if l.ID == v.Location {
				for _, u := range users.Users {
					if v.User == u.ID {
						LocID_Users[l.ID] = append(LocID_Users[l.ID], u)
					}
				}
			}
		}
	}
}

//func InitVisitMaps(users Users, visits Visits, locs Locations) {
//	VisID_Vis := make(map[int]Visit)
//	for _, v := range visits.Visits {
//		VisID_Vis[v.ID] = v
//	}
//}
//func InitLocsMaps(users Users, visits Visits, locs Locations) {
//	LocID_Loc := make(map[int]Location)
//	for _, l := range locs.Locations {
//		LocID_Loc[l.ID] = l
//	}
//}

func mapUsersLocations(u Users, v Visits, l Locations) map[User][]Location {
	res := make(map[User][]Location)

	for _, user := range u.Users {
		var locs []Location
		for _, vis := range v.Visits {
			if user.ID == vis.User {
				for _, loc := range l.Locations {
					if vis.Location == loc.ID {
						locs = append(locs, loc)
					}
				}
			}
		}
		res[user] = locs
		locs = []Location{}
	}
	return res
}
