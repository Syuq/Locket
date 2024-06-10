package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Syuq/Locket/store"
)

func (d *DB) CreateResource(ctx context.Context, create *store.Resource) (*store.Resource, error) {
	fields := []string{"`uid`", "`filename`", "`blob`", "`external_link`", "`type`", "`size`", "`creator_id`", "`internal_path`", "`locket_id`"}
	placeholder := []string{"?", "?", "?", "?", "?", "?", "?", "?", "?"}
	args := []any{create.UID, create.Filename, create.Blob, create.ExternalLink, create.Type, create.Size, create.CreatorID, create.InternalPath, create.LocketID}

	stmt := "INSERT INTO `resource` (" + strings.Join(fields, ", ") + ") VALUES (" + strings.Join(placeholder, ", ") + ")"
	result, err := d.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	id32 := int32(id)
	return d.GetResource(ctx, &store.FindResource{ID: &id32})
}

func (d *DB) ListResources(ctx context.Context, find *store.FindResource) ([]*store.Resource, error) {
	where, args := []string{"1 = 1"}, []any{}

	if v := find.ID; v != nil {
		where, args = append(where, "`id` = ?"), append(args, *v)
	}
	if v := find.UID; v != nil {
		where, args = append(where, "`uid` = ?"), append(args, *v)
	}
	if v := find.CreatorID; v != nil {
		where, args = append(where, "`creator_id` = ?"), append(args, *v)
	}
	if v := find.Filename; v != nil {
		where, args = append(where, "`filename` = ?"), append(args, *v)
	}
	if v := find.LocketID; v != nil {
		where, args = append(where, "`locket_id` = ?"), append(args, *v)
	}
	if find.HasRelatedLocket {
		where = append(where, "`locket_id` IS NOT NULL")
	}

	fields := []string{"`id`", "`uid`", "`filename`", "`external_link`", "`type`", "`size`", "`creator_id`", "UNIX_TIMESTAMP(`created_ts`)", "UNIX_TIMESTAMP(`updated_ts`)", "`internal_path`", "`locket_id`"}
	if find.GetBlob {
		fields = append(fields, "`blob`")
	}

	query := fmt.Sprintf("SELECT %s FROM `resource` WHERE %s ORDER BY `updated_ts` DESC, `created_ts` DESC", strings.Join(fields, ", "), strings.Join(where, " AND "))
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

	list := make([]*store.Resource, 0)
	for rows.Next() {
		resource := store.Resource{}
		var locketID sql.NullInt32
		dests := []any{
			&resource.ID,
			&resource.UID,
			&resource.Filename,
			&resource.ExternalLink,
			&resource.Type,
			&resource.Size,
			&resource.CreatorID,
			&resource.CreatedTs,
			&resource.UpdatedTs,
			&resource.InternalPath,
			&locketID,
		}
		if find.GetBlob {
			dests = append(dests, &resource.Blob)
		}
		if err := rows.Scan(dests...); err != nil {
			return nil, err
		}
		if locketID.Valid {
			resource.LocketID = &locketID.Int32
		}
		list = append(list, &resource)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (d *DB) GetResource(ctx context.Context, find *store.FindResource) (*store.Resource, error) {
	list, err := d.ListResources(ctx, find)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		return nil, nil
	}

	return list[0], nil
}

func (d *DB) UpdateResource(ctx context.Context, update *store.UpdateResource) (*store.Resource, error) {
	set, args := []string{}, []any{}

	if v := update.UID; v != nil {
		set, args = append(set, "`uid` = ?"), append(args, *v)
	}
	if v := update.UpdatedTs; v != nil {
		set, args = append(set, "`updated_ts` = FROM_UNIXTIME(?)"), append(args, *v)
	}
	if v := update.Filename; v != nil {
		set, args = append(set, "`filename` = ?"), append(args, *v)
	}
	if v := update.InternalPath; v != nil {
		set, args = append(set, "`internal_path` = ?"), append(args, *v)
	}
	if v := update.ExternalLink; v != nil {
		set, args = append(set, "`external_link` = ?"), append(args, *v)
	}
	if v := update.LocketID; v != nil {
		set, args = append(set, "`locket_id` = ?"), append(args, *v)
	}
	if v := update.Blob; v != nil {
		set, args = append(set, "`blob` = ?"), append(args, v)
	}

	args = append(args, update.ID)
	stmt := "UPDATE `resource` SET " + strings.Join(set, ", ") + " WHERE `id` = ?"
	if _, err := d.db.ExecContext(ctx, stmt, args...); err != nil {
		return nil, err
	}

	return d.GetResource(ctx, &store.FindResource{ID: &update.ID})
}

func (d *DB) DeleteResource(ctx context.Context, delete *store.DeleteResource) error {
	stmt := "DELETE FROM `resource` WHERE `id` = ?"
	result, err := d.db.ExecContext(ctx, stmt, delete.ID)
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

func vacuumResource(ctx context.Context, tx *sql.Tx) error {
	stmt := "DELETE FROM `resource` WHERE `creator_id` NOT IN (SELECT `id` FROM `user`)"
	_, err := tx.ExecContext(ctx, stmt)
	if err != nil {
		return err
	}

	return nil
}
