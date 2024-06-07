package http

import (
	"encoding/json"
	"log"
	"strings"

	"source.golabs.io/engineering-platforms/productivity/dbaas/dspro/httpUtils"
)

var lmURLTemplate = "https://lambda.golabs.io/api/v1/dms/app-info/app-group-id/{appNamespace}/app-name/{appName}"

func (lm *LambdaManager) LoadObject(appNamespace string, appName string) {

	lmURLNamespace := strings.Replace(lmURLTemplate, "{appNamespace}", appNamespace, 1)
	lmURL := strings.Replace(lmURLNamespace, "{appName}", appName, 1)

	lambdaResponse := httpUtils.GetRequest(lmURL)
	err := json.Unmarshal(lambdaResponse, lm)
	if err != nil {
		log.Fatalln("Error while unmarshling Lambda response: ", err)
	}

	log.Printf("Lambda Manager response: %+v\n\n", lm)
}
