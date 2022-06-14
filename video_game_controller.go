package main

func createVideoGame(videoGame VideoGame) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("INSERT INTO juegos (idjuegos,Nombre,Genero,Año)VALUES(?,?,?,?)", videoGame.Id, videoGame.Name, videoGame.Genero, videoGame.Year)
	return err
}

func deleteVideoGame(id int64) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("DELETE FROM juegos WHERE idjuegos=?", id)
	return err
}

func updateVideoGame(videoGame VideoGame) error {
	bd, err := getDB()
	if err != nil {
		return err
	}
	_, err = bd.Exec("UPDATE juegos SET Nombre=?, Genero=?,Año=? WHERE idjuegos=?", videoGame.Name, videoGame.Genero, videoGame.Year, videoGame.Id)
	return err
}

func getVideoGames() ([]VideoGame, error) {
	//Se devolvera un arreglo por si hay error
	videoGames := []VideoGame{}
	bd, err := getDB()
	if err != nil {
		return videoGames, err
	}
	//Se ontienen renglones con los que se pueden interactuar
	rows, err := bd.Query("SELECT idjuegos, Nombre, Genero, Año FROM juegos")

	if err != nil {
		return videoGames, err
	}
	//Iteraccion con renglones
	for rows.Next() {
		var videoGame VideoGame
		err = rows.Scan(&videoGame.Id, &videoGame.Name, &videoGame.Genero, &videoGame.Year)
		if err != nil {
			return videoGames, err
		}
		videoGames = append(videoGames, videoGame)
	}
	return videoGames, nil
}

func getVideoGameById(id int64) (VideoGame, error) {
	var videoGame VideoGame
	bd, err := getDB()
	if err != nil {
		return videoGame, err
	}
	row := bd.QueryRow("SELECT idjuegos, Nombre, Genero, Año FROM juegos WHERE idjuegos=?", id)
	err = row.Scan(&videoGame.Id, &videoGame.Name, &videoGame.Genero, &videoGame.Year)
	if err != nil {
		return videoGame, err
	}
	return videoGame, nil
}
