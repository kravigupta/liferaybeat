package beater

import (
	"fmt"
	"time"
    "io/ioutil"
	"net/http"
	"strconv"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/packt/liferaybeat/config"
)

type Liferaybeat struct {
	done   chan struct{}
	config config.Config
	client publisher.Client
}

// Creates beater
func New(b *beat.Beat, cfg *common.Config) (beat.Beater, error) {
    
	config := config.DefaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	bt := &Liferaybeat{
		done: make(chan struct{}),
		config: config,
	}
	return bt, nil
}

func (bt *Liferaybeat) Run(b *beat.Beat) error {
    fmt.Println("In run")
	logp.Info("liferaybeat is running! Hit CTRL-C to stop it.")

	bt.client = b.Publisher.Connect()
	ticker := time.NewTicker(bt.config.Period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

        resp, err := http.Get("http://localhost:8080/api/jsonws/liferay-status-portlet.liferaystatus/get-used-memory")
        if err != nil {
            fmt.Println("Something went wrong")
        }
        defer resp.Body.Close()
        body, err := ioutil.ReadAll(resp.Body)

        fmt.Printf("Body: %s\n", body)
        fmt.Printf("Error: %v\n", err)

        jsonStr := fmt.Sprintf("%s", body)
        fmt.Println(jsonStr)
        
        //var numbers float64
        numbers, err := strconv.ParseInt(jsonStr,10, 64)

        //numbers = append(numbers, i)
        fmt.Println(numbers)
        
		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
			"counter":    counter,
            "memoryUsage":numbers, 
		}
		bt.client.PublishEvent(event)
		logp.Info("Event sent")
		counter++
	}
}

func (bt *Liferaybeat) Stop() {
	bt.client.Close()
	close(bt.done)
}
