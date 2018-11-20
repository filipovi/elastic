package elastic

import (
	"github.com/filipovi/elastic/config"
	"github.com/olivere/elastic"
)

// Client is the ElasticSearch Client structure
type Client struct {
	*elastic.Client
}

// Connect returns an Elastic Client
func Connect(url string) (*Client, error) {
	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}

// New returns an Elastic Connection based on a config file
func New(path string) (*Client, error) {
	cfg, err := config.New(path)
	if err != nil {
		return nil, err
	}
	conn, err := Connect(cfg.Elastic.URL)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
