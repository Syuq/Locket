package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Syuq/Locket/store"
)

func (d *DB) UpsertLocketOrganizer(ctx context.Context, upsert *store.LocketOrganizer) (*store.LocketOrganizer, error) {
	pinned := 0
	if upsert.Pinned {
		pinned = 1
	}
	stmt := `
		INSERT INTO locket_organizer (
			locket_id,
			user_id,
			pinned
		)
		VALUES (` + placeholders(3) + `)
		ON CONFLICT(locket_id, user_id) DO UPDATE 
		SET pinned = EXCLUDED.pinned`
	if _, err := d.db.ExecContext(ctx, stmt, upsert.LocketID, upsert.UserID, pinned); err != nil {
		return nil, err
	}

	return upsert, nil
}

func (d *DB) ListLocketOrganizer(ctx context.Context, find *store.FindLocketOrganizer) ([]*store.LocketOrganizer, error) {
	where, args := []string{"1 = 1"}, []any{}
	if find.LocketID != 0 {
		where, args = append(where, "locket_id = "+placeholder(len(args)+1)), append(args, find.LocketID)
	}
	if find.UserID != 0 {
		where, args = append(where, "user_id = "+placeholder(len(args)+1)), append(args, find.UserID)
	}

	query := fmt.Sprintf(`
		SELECT
			locket_id,
			user_id,
			pinned
		FROM locket_organizer
		WHERE %s
	`, strings.Join(where, " AND "))
	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := []*store.LocketOrganizer{}
	for rows.Next() {
		locketOrganizer := &store.LocketOrganizer{}
		pinned := 0
		if err := rows.Scan(
			&locketOrganizer.LocketID,
			&locketOrganizer.UserID,
			&pinned,
		); err != nil {
			return nil, err
		}

		locketOrganizer.Pinned = pinned == 1
		list = append(list, locketOrganizer)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (d *DB) DeleteLocketOrganizer(ctx context.Context, delete *store.DeleteLocketOrganizer) error {
	where, args := []string{}, []any{}
	if v := delete.LocketID; v != nil {
		where, args = append(where, "locket_id = "+placeholder(len(args)+1)), append(args, *v)
	}
	if v := delete.UserID; v != nil {
		where, args = append(where, "user_id = "+placeholder(len(args)+1)), append(args, *v)
	}
	stmt := `DELETE FROM locket_organizer WHERE ` + strings.Join(where, " AND ")
	if _, err := d.db.ExecContext(ctx, stmt, args...); err != nil {
		return err
	}
	return nil
}

func vacuumLocketOrganizer(ctx context.Context, tx *sql.Tx) error {
	stmt := `
	DELETE FROM 
		locket_organizer 
	WHERE 
		locket_id NOT IN (SELECT id FROM locket)
		OR user_id NOT IN (SELECT id FROM "user")`
	_, err := tx.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	return nil
}
