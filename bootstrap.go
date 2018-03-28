package shell

import (
	"context"
	"encoding/json"
)

type respBody struct {
	Peers []string
}

// Add bootstrap node
func (s *Shell) AddBootstrapNode(addr string) ([]string, error) {
	resp, err := s.newRequest(context.Background(), "bootstrap/add", addr).Send(s.httpcli)
	if err != nil {
		return nil, err
	}
	if resp.Error != nil {
		return nil, resp.Error
	}
	body := respBody{}
	err = json.NewDecoder(resp.Output).Decode(&body)
	if err != nil {
		return nil, err
	}
	return body.Peers, nil
}

// Clear bootstrap node list
func (s *Shell) ClearBootstrap() ([]string, error) {
	resp, err := s.newRequest(context.Background(), "bootstrap/rm/all").Send(s.httpcli)
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, resp.Error
	}
	body := respBody{}
	err = json.NewDecoder(resp.Output).Decode(&body)
	if err != nil {
		return nil, err
	}
	return body.Peers, nil
}

// Show all bootstrap nodes
func (s *Shell) ShowBootstrapList() ([]string, error) {
	resp, err := s.newRequest(context.Background(), "bootstrap/list").Send(s.httpcli)
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, resp.Error
	}
	body := respBody{}
	err = json.NewDecoder(resp.Output).Decode(&body)
	if err != nil {
		return nil, err
	}
	return body.Peers, nil
}

// Add default bootstrap nodes
func (s *Shell) AddDefaultBootstrap() ([]string, error) {
	resp, err := s.newRequest(context.Background(), "bootstrap/add/default").Send(s.httpcli)
	if err != nil {
		return nil, err
	}

	if resp.Error != nil {
		return nil, resp.Error
	}
	body := respBody{}
	err = json.NewDecoder(resp.Output).Decode(&body)
	if err != nil {
		return nil, err
	}
	return body.Peers, nil
}
