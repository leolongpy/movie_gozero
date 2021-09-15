package db

import (
	"database/sql"
	"movie_gozero/rpc/film/internal/entity"
)

func SelectFilmActorByMid(mid int64) ([]*entity.FilmActor, error) {

	filmActors := []*entity.FilmActor{}
	err := db.Select(&filmActors, "SELECT `film_id`,`film_name`,`actor_id`,`actor_name` FROM `film_actor` WHERE `film_id` = ?", mid)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	return filmActors, err
}
