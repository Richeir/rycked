package es

import (
	"encoding/json"
	apm "github.com/Richeir/rycked"
	"log"
	"testing"
	"time"
)

func TestGetTracer(t *testing.T) {

	//var tracer apm.Tracer
	//tracer.

	var (
		r map[string]interface{}
	)

	var res = apm.GetTracer("32dd50f4-5d97-11eb-834d-6c626debff99")

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	} else {
		t.Log(res)

		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		}

		log.Printf(
			"[%s] %d hits; took: %dms",
			res.Status(),
			int(r["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
			int(r["took"].(float64)),
		)

		log.Print("hits content:")
		hitContent := r["hits"].(map[string]interface{})["hits"].([]interface{})
		log.Print(hitContent)
		log.Print(hitContent[0])
		log.Print(hitContent[0].(map[string]interface{}))
		log.Print(hitContent[0].(map[string]interface{})["_source"])
		log.Print(hitContent[0].(map[string]interface{})["_source"].(map[string]interface{})["Name"])

		var tracer apm.Tracer
		if len(hitContent) > 0 {
			secondsEastOfUTC := int((8 * time.Hour).Seconds())
			beijing := time.FixedZone("Beijing Time", secondsEastOfUTC)
			layout := "2006-01-02T15:04:05.0000000+08:00"
			var tracerObj = hitContent[0].(map[string]interface{})["_source"]
			tracer.ID = tracerObj.(map[string]interface{})["ID"].(string)
			tracer.DocumentID = tracerObj.(map[string]interface{})["DocumentID"].(string)
			tracer.Name = tracerObj.(map[string]interface{})["Name"].(string)
			tracer.StartAt, _ = time.ParseInLocation(layout, tracerObj.(map[string]interface{})["StartAt"].(string), beijing)
			tracer.FinishAt, _ = time.Parse(layout, tracerObj.(map[string]interface{})["FinishAt"].(string))
			//t.Log(tracerObj.(map[string]interface{})["StartAt"].(string))
			//
			//str := "2021-01-24T00:22:33.9831968+08:00"
			//t, err := time.Parse(layout, str)
			//if err != nil {
			//	log.Print(err)
			//}
			//log.Print(t)
		}

		for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
			log.Printf(" * ID=%s, %s", hit.(map[string]interface{})["_id"], hit.(map[string]interface{})["_source"])
		}
	}
}
