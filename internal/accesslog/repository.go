package accesslog

import (
	"context"
	"fmt"
	"strings"

	"github.com/MedHubUz/access-log/internal/database"
	"github.com/MedHubUz/access-log/internal/entity"
	"github.com/MedHubUz/access-log/internal/search"
	"github.com/MedHubUz/access-log/pkg/slice"
)

type accessLogRepo struct {
	db               *database.PostgresDB
	tableName        string
	searchParameters []string
}

func NewRepository(db *database.PostgresDB) entity.AccessLogRepository {
	return &accessLogRepo{db, "access_log", []string{
		"id",
		"user_id",
		"method",
		"url",
		"date",
		"created_at",
	}}
}

func (r *accessLogRepo) Create(ctx context.Context, accessLog *entity.AccessLog) error {
	clauses := map[string]interface{}{
		"id":         accessLog.ID,
		"user_id":    accessLog.UserID,
		"method":     accessLog.Method,
		"url":        accessLog.URL,
		"data":       accessLog.Data,
		"created_at": accessLog.CreatedAt,
	}

	sqlStr, args, err := r.db.Sq.Builder.
		Insert(r.tableName).
		SetMap(clauses).
		ToSql()

	if err != nil {
		return fmt.Errorf("error during sql build, access log create: %w", err)
	}

	if _, err = r.db.Exec(ctx, sqlStr, args...); err != nil {
		return fmt.Errorf("error during create access log: %w", err)
	}

	return nil
}

func (r *accessLogRepo) ReadAll(ctx context.Context, parameters map[string][]string) ([]*entity.AccessLog, error) {
	var items []*entity.AccessLog

	queryParameters := search.NewQueryParameters(parameters)

	query := r.db.Sq.Builder.
		Select("id, user_id, method, url, data, created_at").
		From(r.tableName).
		Limit(queryParameters.GetLimit()).
		Offset(queryParameters.GetOffset())

	for k, v := range queryParameters.GetParameters() {
		if !slice.Contains(r.searchParameters, k) {
			continue
		}

		if k == "url" {
			query = query.Where("url ilike '%" + v.(string) + "%' ")
			continue
		}

		if k == "date" {
			if date := strings.Split(v.(string), "|"); len(date) == 2 {
				query = query.Where("created_at >= ? AND  created_at <= ?", date[0], date[1])
			}
			continue
		}

		query = query.Where(r.db.Sq.Equal(k, v))
	}

	sqlStr, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("error during sql build, access log read all: %w", err)
	}

	rows, err := r.db.Query(ctx, sqlStr, args...)
	if err != nil {
		return nil, fmt.Errorf("error during access log read all sql query: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var accessLog entity.AccessLog
		err := rows.Scan(
			&accessLog.ID,
			&accessLog.UserID,
			&accessLog.Method,
			&accessLog.URL,
			&accessLog.Data,
			&accessLog.CreatedAt,
		)
		if err != nil {
			return items, fmt.Errorf("error during scan access log read all %w", err)
		}
		items = append(items, &accessLog)
	}
	return items, nil
}
