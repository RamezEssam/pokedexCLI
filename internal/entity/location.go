package entity


type LocationResponse struct {
    Count    int         `json:"count"`
    Next     string      `json:"next"`
    Previous interface{} `json:"previous"`
    Results  []Location  `json:"results"`
}

type Location struct {
    Name string `json:"name"`
    URL  string `json:"url"`
}