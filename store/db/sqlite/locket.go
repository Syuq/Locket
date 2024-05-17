package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/syuq/locket/store"
)

func (d *DB) CreateLocket(ctx context.Context, create *store.Locket) (*store.Locket, error) {
	fields := []string{"`uid`", "`creator_id`", "`content`", "`visibility`"}
	placeholder := []string{"?", "?", "?", "?"}
	args := []any{create.UID, create.CreatorID, create.Content, create.Visibility}

	stmt := "INSERT INTO `locket` (" + strings.Join(fields, ", ") + ") VALUES (" + strings.Join(placeholder, ", ") + ") RETURNING `id`, `created_ts`, `updated_ts`, `row_status`"
	if err := d.db.QueryRowContext(ctx, stmt, args...).Scan(
		&create.ID,
		&create.CreatedTs,
		&create.UpdatedTs,
		&create.RowStatus,
	); err != nil {
		return nil, err
	}

	return create, nil
}

func (d *DB) ListLockets(ctx context.Context, find *store.FindLocket) ([]*store.Locket, error) {
	where, args := []string{"1 = 1"}, []any{}

	if v := find.ID; v != nil {
		where, args = append(where, "`locket`.`id` = ?"), append(args, *v)
	}
	if v := find.UID; v != nil {
		where, args = append(where, "`locket`.`uid` = ?"), append(args, *v)
	}
	if v := find.CreatorID; v != nil {
		where, args = append(where, "`locket`.`creator_id` = ?"), append(args, *v)
	}
	if v := find.RowStatus; v != nil {
		where, args = append(where, "`locket`.`row_status` = ?"), append(args, *v)
	}
	if v := find.CreatedTsBefore; v != nil {
		where, args = append(where, "`locket`.`created_ts` < ?"), append(args, *v)
	}
	if v := find.CreatedTsAfter; v != nil {
		where, args = append(where, "`locket`.`created_ts` > ?"), append(args, *v)
	}
	if v := find.UpdatedTsBefore; v != nil {
		where, args = append(where, "`locket`.`updated_ts` < ?"), append(args, *v)
	}
	if v := find.UpdatedTsAfter; v != nil {
		where, args = append(where, "`locket`.`updated_ts` > ?"), append(args, *v)
	}
	if v := find.ContentSearch; len(v) != 0 {
		for _, s := range v {
			where, args = append(where, "`locket`.`content` LIKE ?"), append(args, fmt.Sprintf("%%%s%%", s))
		}
	}
	if v := find.VisibilityList; len(v) != 0 {
		placeholder := []string{}
		for _, visibility := range v {
			placeholder = append(placeholder, "?")
			args = append(args, visibility.String())
		}
		where = append(where, fmt.Sprintf("`locket`.`visibility` IN (%s)", strings.Join(placeholder, ",")))
	}
	if find.ExcludeComments {
		where = append(where, "`parent_id` IS NULL")
	}

	orderBy := []string{}
	if find.OrderByPinned {
		orderBy = append(orderBy, "`pinned` DESC")
	}
	if find.OrderByUpdatedTs {
		orderBy = append(orderBy, "`updated_ts` DESC")
	} else {
		orderBy = append(orderBy, "`created_ts` DESC")
	}
	orderBy = append(orderBy, "`id` DESC")
	if find.Random {
		orderBy = []string{"RANDOM()"}
	}

	fields := []string{
		"`locket`.`id` AS `id`",
		"`locket`.`uid` AS `uid`",
		"`locket`.`creator_id` AS `creator_id`",
		"`locket`.`created_ts` AS `created_ts`",
		"`locket`.`updated_ts` AS `updated_ts`",
		"`locket`.`row_status` AS `row_status`",
		"`locket`.`visibility` AS `visibility`",
		"IFNULL(`locket_organizer`.`pinned`, 0) AS `pinned`",
		"`locket_relation`.`related_locket_id` AS `parent_id`",
	}
	if !find.ExcludeContent {
		fields = append(fields, "`locket`.`content` AS `content`")
	}

	query := "SELECT " + strings.Join(fields, ", ") + "FROM `locket` " +
		"LEFT JOIN `locket_organizer` ON `locket`.`id` = `locket_organizer`.`locket_id` AND `locket`.`creator_id` = `locket_organizer`.`user_id` " +
		"LEFT JOIN `locket_relation` ON `locket`.`id` = `locket_relation`.`locket_id` AND `locket_relation`.`type` = \"COMMENT\" " +
		"WHERE " + strings.Join(where, " AND ") + " " +
		"ORDER BY " + strings.Join(orderBy, ", ")
	if find.Limit != nil {
		query = fmt.Sprintf("%s LIMIT %d", query, *find.Limit)
		if find.Offset != nil {
			query = fmt.Sprintf("%s OFFSET %d", query, *find.Offset)
		}
	}

	rows, err := d.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	list := make([]*store.Locket, 0)
	for rows.Next() {
		var locket store.Locket
		dests := []any{
			&locket.ID,
			&locket.UID,
			&locket.CreatorID,
			&locket.CreatedTs,
			&locket.UpdatedTs,
			&locket.RowStatus,
			&locket.Visibility,
			&locket.Pinned,
			&locket.ParentID,
		}
		if !find.ExcludeContent {
			dests = append(dests, &locket.Content)
		}
		if err := rows.Scan(dests...); err != nil {
			return nil, err
		}
		list = append(list, &locket)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (d *DB) UpdateLocket(ctx context.Context, update *store.UpdateLocket) error {
	set, args := []string{}, []any{}
	if v := update.UID; v != nil {
		set, args = append(set, "`uid` = ?"), append(args, *v)
	}
	if v := update.CreatedTs; v != nil {
		set, args = append(set, "`created_ts` = ?"), append(args, *v)
	}
	if v := update.UpdatedTs; v != nil {
		set, args = append(set, "`updated_ts` = ?"), append(args, *v)
	}
	if v := update.RowStatus; v != nil {
		set, args = append(set, "`row_status` = ?"), append(args, *v)
	}
	if v := update.Content; v != nil {
		set, args = append(set, "`content` = ?"), append(args, *v)
	}
	if v := update.Visibility; v != nil {
		set, args = append(set, "`visibility` = ?"), append(args, *v)
	}
	args = append(args, update.ID)

	stmt := "UPDATE `locket` SET " + strings.Join(set, ", ") + " WHERE `id` = ?"
	if _, err := d.db.ExecContext(ctx, stmt, args...); err != nil {
		return err
	}
	return nil
}

func (d *DB) DeleteLocket(ctx context.Context, delete *store.DeleteLocket) error {
	where, args := []string{"`id` = ?"}, []any{delete.ID}
	stmt := "DELETE FROM `locket` WHERE " + strings.Join(where, " AND ")
	result, err := d.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}
	if _, err := result.RowsAffected(); err != nil {
		return err
	}

	if err := d.Vacuum(ctx); err != nil {
		// Prevent linter warning.
		return err
	}
	return nil
}

func vacuumLocket(ctx context.Context, tx *sql.Tx) error {
	stmt := "DELETE FROM `locket` WHERE `creator_id` NOT IN (SELECT `id` FROM `user`)"
	_, err := tx.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	return nil
}
