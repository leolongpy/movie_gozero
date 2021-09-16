package db

import (
	"database/sql"
	"movie_gozero/rpc/order/internal/entity"
)

func SelectFilmActorByMid(mid int64) ([]*entity.FilmActor, error) {
	filmActor := []*entity.FilmActor{}
	err := db.Select(&filmActor, "SELECT `actor_name` FROM `film_actor` WHERE `film_id`=?", mid)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return filmActor, err
}
