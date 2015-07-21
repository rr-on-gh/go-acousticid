package acousticid
import (
	"net/http"
	"io/ioutil"
	"net/url"
	"strconv"
	"encoding/json"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

type AcousticidRequest struct {
	Fingerprint string `json:fingerprint`
	Duration    int `json:duration`
	Apikey      string `json:"client"`
	Metadata    string `json:meta`
}

//Generated using http://json2struct.mervine.net/
type AcousticidResponse struct {
	Results []struct {
		ID         string `json:"id"`
		Recordings []struct {
			Artists  []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"artists"`
			Duration float64    `json:"duration"`
			ID       string `json:"id"`
			Title    string `json:"title"`
		} `json:"recordings"`
		Score      float64 `json:"score"`
	} `json:"results"`
	Status  string `json:"status"`
}

func (a *AcousticidRequest) Request() AcousticidResponse {
	client := http.Client{}
	response, err := client.PostForm("http://api.acoustid.org/v2/lookup", a.PostValues())
	check(err)
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	aidresponse := AcousticidResponse{}
	//fmt.Println("-------------->", string(body))
	err = json.Unmarshal(body, &aidresponse)
	check(err)
	return aidresponse
}

func (a *AcousticidRequest) PostValues() url.Values {
	values, err := url.ParseQuery("client=" + a.Apikey + "&duration=" + strconv.Itoa(a.Duration) + "&meta=" + a.Metadata + "&fingerprint=" + a.Fingerprint)
	check(err)
	return values
}

