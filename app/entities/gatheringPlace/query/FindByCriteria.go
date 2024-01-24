package query

import (
	"cmd/app/entities/gatheringPlace"
	"cmd/app/models"
	"context"
	"database/sql"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"
)

func FindByCriteria(ctx context.Context, criteria FindCriteria, db *sql.DB) (*sql.Rows, error) {
	psql := sq.StatementBuilder.PlaceholderFormat(sq.Dollar)

	sqlStatement := psql.Select("*").From("gathering_places").RunWith(db)
	if criteria.ID.Valid != false {
		sqlStatement = findByPlace(sqlStatement, criteria.ID.UUID)
	}
	if criteria.Address != nil {
		sqlStatement = findByAddress(sqlStatement, criteria.Address)
	}
	if criteria.CuisineType != nil {
		sqlStatement = findByCuisineType(sqlStatement, criteria.CuisineType)
	}
	var rows, err = sqlStatement.Query()
	if err != nil {
		return nil, fmt.Errorf("problem with quering to database %w", err)
	}
	return rows, nil
}

func findByPlace(sql sq.SelectBuilder, id uuid.UUID) sq.SelectBuilder {
	return sql.Where(sq.Eq{"id": id})
}
func findByAddress(sql sq.SelectBuilder, address *models.Address) sq.SelectBuilder {
	return sql.Where(sq.Eq{"country": address.Country, "city": address.City,
		"street_name": address.StreetName, "house_number": address.HouseNumber,
		"building_number": address.BuildingNumber})
}
func findByCuisineType(sql sq.SelectBuilder, cuisineType *gatheringPlace.CuisineType) sq.SelectBuilder {
	return sql.Where(sq.Eq{"cuisine_type": cuisineType})
}
