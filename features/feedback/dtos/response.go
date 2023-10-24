package dtos

type ResFeedback struct {
	Member			string `json:"member"`
	Comment 		string `json:"comment"`
	PriorityStatus 	string `json:"priority-status"`
}

type FeedbackJoinReply struct {
	Member			string `json:"member"`
	Comment 		string `json:"comment"`
	PriorityStatus 	string `json:"priority-status"`
	Staff			string `json:"staff"`	
	Reply 			string `json:"reply"`
}

type FeedbackWithReply struct {
	Member			string 		`json:"member"`
	Comment 		string 		`json:"comment"`
	PriorityStatus 	string 		`json:"priority-status"`
	Reply 			StaffReply 	`json:"reply"`
}

type StaffReply struct {
	Staff	string `json:"staff"`	
	Comment string `json:"comment"`
}

