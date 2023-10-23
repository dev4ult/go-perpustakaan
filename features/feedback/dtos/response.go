package dtos

type ResFeedback struct {
	Member			string `json:"member"`
	Comment 		string `json:"comment"`
	PriorityStatus 	string `json:"priority-status"`
}

type FeedbackWithReply struct {
	Member			string `json:"member"`
	Comment 		string `json:"comment"`
	PriorityStatus 	string `json:"priority-status"`
	Reply 			StaffReply
}

type StaffReply struct {
	Staff	string `json:"staff"`	
	Comment string `json:"comment"`
}

