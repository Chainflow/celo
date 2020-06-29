package server

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	client "github.com/influxdata/influxdb1-client/v2"

	"github.com/chainflow/celo/config"
)

// NewProposalAlert alerts for a new proposal
func NewProposalAlert(cfg *config.Config, c client.Client) error {
	log.Println("Coming inside new proposals")

	bp, err := createBatchPoints(cfg.InfluxDB.Database)
	if err != nil {
		return err
	}

	ops := HTTPOptions{
		Endpoint: "http://thecelo.com/api/v0.1?method=proposalList",
		Method:   "GET",
	}

	res, err := HitHTTPTarget(ops)
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	var proposal Proposals
	err = json.Unmarshal(res.Body, &proposal)
	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}

	length := len(proposal.ItemsVote)
	p := (proposal.ItemsVote[length-1])
	value := p.([]interface{})[0]
	id := value.([]interface{})[1]

	var lengthOfProposals string

	q := client.NewQuery(fmt.Sprintf("SELECT * FROM celo_proposals"), cfg.InfluxDB.Database, "")
	if response, err := c.Query(q); err == nil && response.Error() == nil {
		for _, r := range response.Results {
			if len(r.Series) != 0 {
				n := len(r.Series[0].Values)

				proposalsLen := fmt.Sprintf("%v", r.Series[0].Values[n-1][1])
				lengthOfProposals = proposalsLen
			}
		}
	}

	l, _ := strconv.Atoi(lengthOfProposals)

	if l != length {
		_ = writeToInfluxDb(c, bp, "celo_proposals", map[string]string{}, map[string]interface{}{"proposal_id": id, "length": length})
		_ = SendTelegramAlert(fmt.Sprintf("A new roposal with proposal id = %s has been added", id), cfg)
		_ = SendEmailAlert(fmt.Sprintf("A new proposal with proposal id = %s has been added", id), cfg)
		log.Println("Sent new proposal alerting")
	}
	return nil
}
