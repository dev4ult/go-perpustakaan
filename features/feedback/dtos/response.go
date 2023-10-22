package dtos

type ResFeedback struct {
	User			string `json:"user"`
	Comment 		string `json:"comment"`
	PriorityStatus 	string `json:"priority-status"`
	ReplyComment 	string `json:"reply-comment"`
}

