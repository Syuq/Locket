package sqlite

import (
	"context"
	"database/sql"
	"strings"

	"github.com/syuq/locket/store"
)

func (d *DB) UpsertLocketRelation(ctx context.Context, create *store.LocketRelation) (*store.LocketRelation, error) {
	stmt := `
		INSERT INTO locket_relation (
			locket_id,
			related_locket_id,
			type
		)
		VALUES (?, ?, ?)
		RETURNING locket_id, related_locket_id, type
	`
	locketRelation := &store.LocketRelation{}
	if err := d.db.QueryRowContext(
		ctx,
		stmt,
		create.LocketID,
		create.RelatedLocketID,
		create.Type,
	).Scan(
		&locketRelation.LocketID,
		&locketRelation.RelatedLocketID,
		&locketRelation.Type,
	); err != nil {
		return nil, err
	}

	return locketRelation, nil
}

func (d *DB) ListLocketRelations(ctx context.Context, find *store.FindLocketRelation) ([]*store.LocketRelation, error) {
	where, args := []string{"TRUE"}, []any{}
	if find.LocketID != nil {
		where, args = append(where, "locket_id = ?"), append(args, find.LocketID)
	}
	if find.RelatedLocketID != nil {
		where, args = append(where, "related_locket_id = ?"), append(args, find.RelatedLocketID)
	}
	if find.Type != nil {
		where, args = append(where, "type = ?"), append(args, find.Type)
	}

	rows, err := d.db.QueryContext(ctx, `
		SELECT
			locket_id,
			related_locket_id,
			type
		FROM locket_relation
		WHERE `+strings.Join(where, " AND "), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := []*store.LocketRelation{}
	for rows.Next() {
		locketRelation := &store.LocketRelation{}
		if err := rows.Scan(
			&locketRelation.LocketID,
			&locketRelation.RelatedLocketID,
			&locketRelation.Type,
		); err != nil {
			return nil, err
		}
		list = append(list, locketRelation)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (d *DB) DeleteLocketRelation(ctx context.Context, delete *store.DeleteLocketRelation) error {
	where, args := []string{"TRUE"}, []any{}
	if delete.LocketID != nil {
		where, args = append(where, "locket_id = ?"), append(args, delete.LocketID)
	}
	if delete.RelatedLocketID != nil {
		where, args = append(where, "related_locket_id = ?"), append(args, delete.RelatedLocketID)
	}
	if delete.Type != nil {
		where, args = append(where, "type = ?"), append(args, delete.Type)
	}
	stmt := `
		DELETE FROM locket_relation
		WHERE ` + strings.Join(where, " AND ")
	result, err := d.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}
	if _, err = result.RowsAffected(); err != nil {
		return err
	}
	return nil
}

func vacuumLocketRelations(ctx context.Context, tx *sql.Tx) error {
	if _, err := tx.ExecContext(ctx, `
		DELETE FROM locket_relation
		WHERE locket_id NOT IN (SELECT id FROM locket) OR related_locket_id NOT IN (SELECT id FROM locket)
	`); err != nil {
		return err
	}
	return nil
}
