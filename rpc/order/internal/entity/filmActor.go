package entity

type FilmActor struct {
	FaId      int64  `json:"fa_id" db:"fa_id"`
	FilmId    string `json:"film_id" db:"film_id"`
	FilmName  string `json:"film_name" db:"film_name"`
	ActorId   string `json:"actor_id" db:"actor_id"`
	ActorName string `json:"actor_name" db:"actor_name"`
}
