package project

type Project struct {
	Name        string `json:"name"`
	TogglId     uint64 `json:"togglId"`
	TripletexId int32  `json:"tripletexId"`
}

type Projects struct {
	VersionDigest string            `json:"versionDigest"`
	Projects      map[int32]Project `json:"projects"`
}
