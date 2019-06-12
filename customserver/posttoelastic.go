package customserver

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/olivere/elastic"
	"github.com/teris-io/shortid"
)

type ReqBody struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

type Tweet struct {
	User    string `json:"user"`
	Message string `json:"message"`
}

type RaspEvent struct {
	AttackParams struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"attack_params"`

	AttackSource string `json:"attack_source"`
	AttackType   string `json:"attack_type"`
	Body         string `json:"body"`
	ClientIP     string `json:"client_ip"`
	CorID        string `json:"cor_id"`
	EventTime    string `json:"event_time"`
	EventType    string `json:"event_type"`

	Header struct {
		Accept                  string `json:"accept"`
		AcceptEncoding          string `json:"accept-encoding"`
		AcceptLanguage          string `json:"accept-language"`
		CacheControl            string `json:"cache-control"`
		Connection              string `json:"connection"`
		ContentLength           string `json:"content-length"`
		ContentType             string `json:"content-type"`
		Cookie                  string `json:"cookie"`
		Host                    string `json:"host"`
		Origin                  string `json:"origin"`
		Referer                 string `json:"referer"`
		UpgradeInsecureRequests string `json:"upgrade-insecure-requests"`
		UserAgent               string `json:"user-agent"`
	} `json:"header"`

	InterceptState   string `json:"intercept_state"`
	Path             string `json:"path"`
	PluginAlgorithm  string `json:"plugin_algorithm"`
	PluginConfidence int    `json:"plugin_confidence"`
	PluginMessage    string `json:"plugin_message"`
	PluginName       string `json:"plugin_name"`
	RequestID        string `json:"request_id"`
	RequestMethod    string `json:"request_method"`
	ServerHostname   string `json:"server_hostname"`
	ServerIP         string `json:"server_ip"`

	ServerNic []struct {
		IP   string `json:"ip"`
		Name string `json:"name"`
	} `json:"server_nic"`

	ServerType    string `json:"server_type"`
	ServerVersion string `json:"server_version"`
	StackTrace    string `json:"stack_trace"`
	Target        string `json:"target"`
	URL           string `json:"url"`
}

func InsertDataIntoElastic(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}

	// Create a client
	client, err := elastic.NewClient(elastic.SetURL(ElasticHost))
	if err != nil {
		// Handle error
	}

	// Use the IndexExists service to check if a specified index exists.
	exists, err := client.IndexExists("rasp_events").Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}

	if !exists {
		// Create a new index.
		createIndex, err := client.CreateIndex("rasp_events").Do(ctx)
		if err != nil {
			// Handle error
			panic(err)
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}

	//  TODO, create id from agent application registration-
	raspEventData := RaspEvent{}
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &raspEventData)

	// Add a document to the index
	_, err = client.Index().
		Index("rasp_events").
		Id(shortid.MustGenerate()).
		BodyJson(raspEventData).
		Do(ctx)

	if err != nil {
		// Handle error
		panic(err)
	}

	details := make(map[string]string)
	details["success"] = "true"
	Json(w, http.StatusOK, details)
}

// func pingElasticServe(cleint *Client, ctx *context) bool {
// 	info, code, erro := client.Ping(ELASTIC_HOST).Do(ctx)
// }
