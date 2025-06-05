package pokeapi

type RespShallowLocationExplore struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Encounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	}
}
