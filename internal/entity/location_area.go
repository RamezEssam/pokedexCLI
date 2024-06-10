package entity

type PokemonEncounter struct {
    Pokemon struct {
        Name string `json:"name"`
    } `json:"pokemon"`
}

type LocationArea struct {
    ID                int               `json:"id"`
    Name              string            `json:"name"`
    GameIndex         int               `json:"game_index"`
    EncounterMethodRates []struct {
        EncounterMethod struct {
            Name string `json:"name"`
            URL  string `json:"url"`
        } `json:"encounter_method"`
        VersionDetails []struct {
            Rate    int `json:"rate"`
            Version struct {
                Name string `json:"name"`
                URL  string `json:"url"`
            } `json:"version"`
        } `json:"version_details"`
    } `json:"encounter_method_rates"`
    Location struct {
        Name string `json:"name"`
        URL  string `json:"url"`
    } `json:"location"`
    Names []struct {
        Name     string `json:"name"`
        Language struct {
            Name string `json:"name"`
            URL  string `json:"url"`
        } `json:"language"`
    } `json:"names"`
    PokemonEncounters []PokemonEncounter `json:"pokemon_encounters"`
} 