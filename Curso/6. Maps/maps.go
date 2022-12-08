package main

// MAPS têm a função de criar tipos (user, job, company) que usem os mesmos
//tipos de chaves(string / int)
import "fmt"

func main() {
	novoUsuario := map[string]string{ // map + [tipo da chave] + tipo do valor
		"Nome":      "Juliano", // maps precisam de vírgula após a inserção de valor
		"Sobrenome": "da Silva",
	}
	fmt.Println(novoUsuario)
	fmt.Println(novoUsuario["Sobrenome"])

	usuario2 := map[string]map[string]string{
		"Name": {
			"First name": "Joao",
			"Last name":  "da silva",
		},
		"Secondary data": {
			"Nationality": "Brazilian",
			"Language":    "Portuguese",
		},
	}
	fmt.Println(usuario2)
	delete(usuario2, "Secondary data") // usamos delete para apagar chaves existentes e seus respectivos valores
	fmt.Println(usuario2)
	usuario2["Food allergies"] = map[string]string{ //criamos novas chaves com essa sintaxe
		"Allergic to something?": "No",
	}
	fmt.Println(usuario2)
}
