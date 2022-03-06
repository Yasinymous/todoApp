package models

// type User struct {
// 	id            int64
// 	username      string
// 	email         string
// 	password      string
// 	session_token string
// 	created_at    time.Time `pg:"default:now()"`
// 	updated_at    time.Time
// }

// type Board struct {
// 	id      int64
// 	user_id *User `pg:"rel:has-one"`

// 	title      string
// 	created_at time.Time `pg:"default:now()"`
// 	updated_at time.Time
// }

// type BoardMember struct {
// 	id       int64
// 	board_id *Board `pg:"rel:has-many"`
// 	user_id  *User  `pg:"rel:has-many"`

// 	title      string
// 	created_at time.Time `pg:"default:now()"`
// 	updated_at time.Time
// }

// type List struct {
// 	id       int64
// 	board_id *Board `pg:"rel:has-many"`

// 	title      string
// 	ord        int64
// 	created_at time.Time `pg:"default:now()"`
// 	updated_at time.Time
// }

// type Card struct {
// 	id      int64
// 	list_id *List `pg:"rel:has-many"`

// 	title       string
// 	description string
// 	ord         int64
// 	created_at  time.Time `pg:"default:now()"`
// 	updated_at  time.Time
// }

// type Item struct {
// 	id      int64
// 	card_id *Card `pg:"rel:has-many"`

// 	done       int64
// 	created_at time.Time `pg:"default:now()"`
// 	updated_at time.Time
// }

// type CardAssignment struct {
// 	id      int64
// 	card_id *Card `pg:"rel:has-many"`
// 	user_id *User `pg:"rel:has-many"`

// 	created_at time.Time `pg:"default:now()"`
// 	updated_at time.Time
// }
