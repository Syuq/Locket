package sqlite

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/Syuq/Locket/store"
)

func (d *DB) CreateIdentityProvider(ctx context.Context, create *store.IdentityProvider) (*store.IdentityProvider, error) {
	var configBytes []byte
	if create.Type == store.IdentityProviderOAuth2Type {
		bytes, err := json.Marshal(create.Config.OAuth2Config)
		if err != nil {
			return nil, err
		}
		configBytes = bytes
	} else {
		return nil, errors.Errorf("unsupported idp type %s", string(create.Type))
	}

	placeholders := []string{"?", "?", "?", "?"}
	fields := []string{"`name`", "`type`", "`identifier_filter`", "`config`"}
	args := []any{create.Name, create.Type, create.IdentifierFilter, string(configBytes)}

	stmt := "INSERT INTO `idp` (" + strings.Join(fields, ", ") + ") VALUES (" + strings.Join(placeholders, ", ") + ") RETURNING `id`"
	if err := d.db.QueryRowContext(ctx, stmt, args...).Scan(&create.ID); err != nil {
		return nil, err
	}

	identityProvider := create
	return identityProvider, nil
}

func (d *DB) ListIdentityProviders(ctx context.Context, find *store.FindIdentityProvider) ([]*store.IdentityProvider, error) {
	where, args := []string{"1 = 1"}, []any{}
	if v := find.ID; v != nil {
		where, args = append(where, fmt.Sprintf("id = $%d", len(args)+1)), append(args, *v)
	}

	rows, err := d.db.QueryContext(ctx, `
		SELECT
			id,
			name,
			type,
			identifier_filter,
			config
		FROM idp
		WHERE `+strings.Join(where, " AND ")+` ORDER BY id ASC`,
		args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var identityProviders []*store.IdentityProvider
	for rows.Next() {
		var identityProvider store.IdentityProvider
		var identityProviderConfig string
		if err := rows.Scan(
			&identityProvider.ID,
			&identityProvider.Name,
			&identityProvider.Type,
			&identityProvider.IdentifierFilter,
			&identityProviderConfig,
		); err != nil {
			return nil, err
		}

		if identityProvider.Type == store.IdentityProviderOAuth2Type {
			oauth2Config := &store.IdentityProviderOAuth2Config{}
			if err := json.Unmarshal([]byte(identityProviderConfig), oauth2Config); err != nil {
				return nil, err
			}
			identityProvider.Config = &store.IdentityProviderConfig{
				OAuth2Config: oauth2Config,
			}
		} else {
			return nil, errors.Errorf("unsupported idp type %s", string(identityProvider.Type))
		}
		identityProviders = append(identityProviders, &identityProvider)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return identityProviders, nil
}

func (d *DB) UpdateIdentityProvider(ctx context.Context, update *store.UpdateIdentityProvider) (*store.IdentityProvider, error) {
	set, args := []string{}, []any{}
	if v := update.Name; v != nil {
		set, args = append(set, "name = ?"), append(args, *v)
	}
	if v := update.IdentifierFilter; v != nil {
		set, args = append(set, "identifier_filter = ?"), append(args, *v)
	}
	if v := update.Config; v != nil {
		var configBytes []byte
		if update.Type == store.IdentityProviderOAuth2Type {
			bytes, err := json.Marshal(update.Config.OAuth2Config)
			if err != nil {
				return nil, err
			}
			configBytes = bytes
		} else {
			return nil, errors.Errorf("unsupported idp type %s", string(update.Type))
		}
		set, args = append(set, "config = ?"), append(args, string(configBytes))
	}
	args = append(args, update.ID)

	stmt := `
		UPDATE idp
		SET ` + strings.Join(set, ", ") + `
		WHERE id = ?
		RETURNING id, name, type, identifier_filter, config
	`
	var identityProvider store.IdentityProvider
	var identityProviderConfig string
	if err := d.db.QueryRowContext(ctx, stmt, args...).Scan(
		&identityProvider.ID,
		&identityProvider.Name,
		&identityProvider.Type,
		&identityProvider.IdentifierFilter,
		&identityProviderConfig,
	); err != nil {
		return nil, err
	}

	if identityProvider.Type == store.IdentityProviderOAuth2Type {
		oauth2Config := &store.IdentityProviderOAuth2Config{}
		if err := json.Unmarshal([]byte(identityProviderConfig), oauth2Config); err != nil {
			return nil, err
		}
		identityProvider.Config = &store.IdentityProviderConfig{
			OAuth2Config: oauth2Config,
		}
	} else {
		return nil, errors.Errorf("unsupported idp type %s", string(identityProvider.Type))
	}

	return &identityProvider, nil
}

func (d *DB) DeleteIdentityProvider(ctx context.Context, delete *store.DeleteIdentityProvider) error {
	where, args := []string{"id = ?"}, []any{delete.ID}
	stmt := `DELETE FROM idp WHERE ` + strings.Join(where, " AND ")
	result, err := d.db.ExecContext(ctx, stmt, args...)
	if err != nil {
		return err
	}
	if _, err = result.RowsAffected(); err != nil {
		return err
	}
	return nil
}
