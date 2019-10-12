package main

func luminosidade(sol int) string {

	var solmodo string = " - "
	if sol >= 0 && sol < 20 {
		solmodo = "Muito Nublado"
	}

	if sol >= 20 && sol < 40 {
		solmodo = "Nublado"
	}

	if sol >= 40 && sol < 60 {
		solmodo = "Normal"
	}

	if sol >= 60 && sol < 80 {
		solmodo = "Ensolado"
	}

	if sol >= 80 {
		solmodo = "Muito Ensolado"
	}

	return solmodo

}
