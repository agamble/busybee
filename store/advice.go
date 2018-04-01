package store

const (
	insertAdviceStmt = `INSERT INTO advice (description, session_id) 
							VALUES 
						(:description, :session_id)`
)

type Advice struct {
	ID          int64  `json:"-"`
	Description string `db:"description"`
	SessionID   string `db:"session_id"`
}

func AddAdvice(a *Advice) (int64, error) {
	r, err := db.NamedExec(insertAdviceStmt, a)
	if err != nil {
		return 0, err
	}

	return r.LastInsertId()
}
