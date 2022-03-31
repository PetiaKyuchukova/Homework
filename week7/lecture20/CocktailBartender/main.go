package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Cocktail struct {
	IdDrink                     string `json:"idDrink"`
	StrDrink                    string `json:"strDrink"`
	StrDrinkAlternate           string `json:"strDrinkAlternate"`
	StrVideo                    string `json:"strVideo"`
	StrCategory                 string `json:"strCategory"`
	StrIBA                      string `json:"strIBA"`
	StrAlcoholic                string `json:"strAlcoholic"`
	StrGlass                    string `json:"strGlass"`
	StrInstructions             string `json:"strInstructions"`
	StrInstructionsES           string `json:"strInstructionsES"`
	StrInstructionsDE           string `json:"strInstructionsDE"`
	StrInstructionsFR           string `json:"strInstructionsFR"`
	StrInstructionsIT           string `json:"strInstructionsIT"`
	StrInstructionsZH_HANS      string `json:"strInstructionsZH-HANS"`
	StrInstructionsZH_HANT      string `json:"strInstructionsZH-HANT"`
	StrIngredient1              string `json:"strIngredient1"`
	StrIngredient2              string `json:"strIngredient2"`
	StrIngredient3              string `json:"strIngredient3"`
	StrIngredient4              string `json:"strIngredient4"`
	StrIngredient6              string `json:"strIngredient6"`
	StrIngredient7              string `json:"strIngredient7"`
	StrIngredient8              string `json:"strIngredient8"`
	StrIngredient9              string `json:"strIngredient9"`
	StrIngredient10             string `json:"strIngredient10"`
	StrIngredient11             string `json:"strIngredient11"`
	StrIngredient12             string `json:"strIngredient12"`
	StrIngredient13             string `json:"strIngredient13"`
	StrIngredient14             string `json:"strIngredient14"`
	StrIngredient15             string `json:"strIngredient15"`
	StrMeasure1                 string `json:"strMeasure1"`
	StrMeasure2                 string `json:"strMeasure2"`
	StrMeasure3                 string `json:"strMeasure3"`
	StrMeasure4                 string `json:"strMeasure4"`
	StrMeasure5                 string `json:"strMeasure5"`
	StrMeasure6                 string `json:"strMeasure6"`
	StrMeasure7                 string `json:"strMeasure7"`
	StrMeasure8                 string `json:"strMeasure8"`
	StrMeasure9                 string `json:"strMeasure9"`
	StrMeasure10                string `json:"strMeasure10"`
	StrMeasure11                string `json:"strMeasure11"`
	StrMeasure12                string `json:"strMeasure12"`
	StrMeasure13                string `json:"strMeasure13"`
	StrMeasure14                string `json:"strMeasure14"`
	StrMeasure15                string `json:"strMeasure15"`
	StrImageSource              string `json:"strImageSource"`
	StrImageAttribution         string `json:"strImageAttribution"`
	StrCreativeCommonsConfirmed string `json:"strCreativeCommonsConfirmed"`
	DateModified                string `json:"dateModified"`
}
type Drinks struct {
	Drinks []Cocktail `json:"drinks"`
}

func Start() {
	var cocktailName string
	var drinks Drinks

	fmt.Print("What would you want to drink? ")
	fmt.Scanln(&cocktailName)

	if cocktailName == "nothing" {
		fmt.Print("byeeeeeeeeeeeeeeeee")
		return
	} else {
		urlA, err := url.Parse("https://thecocktaildb.com/api/json/v1/1/search.php?s=gg")
		if err != nil {
			log.Fatal(err)
		}
		values := urlA.Query()
		values.Set("s", cocktailName)
		urlA.RawQuery = values.Encode()

		resp, err := http.Get(urlA.String())
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		err1 := json.Unmarshal(body, &drinks)
		if err1 != nil {
			fmt.Println(err1)
		}

		if len(drinks.Drinks) == 0 {
			fmt.Printf("We don`t have cocktail %s!\n", cocktailName)
		} else {
			for i, cocktail := range drinks.Drinks {
				if i == 0 {
					recipe := strings.Split(cocktail.StrInstructions, ". ")
					for _, s := range recipe {
						fmt.Printf("%s.\n", s)
					}
					break
				}
			}
		}

		Start()

	}

}
func main() {

	Start()
}
