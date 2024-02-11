package models

type RecentActions struct {
	Time     int       `json:"timeSeconds"`
	Blog     BlogEntry `json:"blogEntry"`
	Comments Comment   `json:"comment"`
}

// type Actions []RecentActions
type BlogEntry struct {
	ViewHistory      bool   `json:"allowViewHistory"`
	Creationtime     int    `json:"creationTimeSeconds"`
	Rating           int    `json:"rating"`
	AuthorHandle     string `json:"authorHandle"`
	Modificationtime int    `json:"modificationTimeSeconds"`
	Id               int    `json:"id"`
	Title            string `json:"title"`
}

type Comment struct {
	Id                int    `json:"id"`
	CreationTime      int    `json:"creationTimeSeconds"`
	CommentatorHandle string `json:"commentatorHandle"`
	Comment           string `json:"text"`
	ParentCommentId   int    `json:"parentCommentId"`
	Rating            int    `json:"rating"`
}

type User struct {
	Cfhandle string `json:"cfhandle"`
	Username string `json:"username"`
	Password string `json:"password"`
}
