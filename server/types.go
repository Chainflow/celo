package server

type ProposalItems struct {
	Status   string `json:"status"`
	Timespan int    `json:"timespan"`
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Proposer struct {
		Address   string `json:"address"`
		Deposit   int    `json:"deposit"`
		Timestamp int    `json:"timestamp"`
	} `json:"proposer"`
	Upvoted struct {
		Peoples int `json:"peoples"`
		Upvotes int `json:"upvotes"`
	} `json:"upvoted"`
	Dequeue struct {
		Address   string `json:"address"`
		Timestamp int    `json:"timestamp"`
	} `json:"dequeue"`
	Approval struct {
		Address   string `json:"address"`
		Timestamp string `json:"timestamp"`
	} `json:"approval"`
	Voted struct {
		Peoples int `json:"peoples"`
		Weight  int `json:"weight"`
	} `json:"voted"`
	Executed struct {
		From            string `json:"from"`
		Timestamp       string `json:"timestamp"`
		BlockNumber     string `json:"blockNumber"`
		TransactionHash string `json:"transactionHash"`
	} `json:"executed"`
}

type Proposals struct {
	Items     interface{}   `json:"items"`
	ItemsVote []interface{} `json:"items_vote"`
}

type (
	QueryParams map[string]string

	// HTTPOptions of a target
	HTTPOptions struct {
		Endpoint    string
		QueryParams QueryParams
		Body        []byte
		Method      string
	}

	// PingResp struct
	PingResp struct {
		StatusCode int
		Body       []byte
	}
)
