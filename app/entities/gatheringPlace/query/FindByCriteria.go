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
	if criteria.Address != nil {
		sqlStatement = findByAddress(sqlStatement, criteria.Address)
	}
	if criteria.InitiatorID.Valid != false {
		sqlStatement = findByInitiatorID(sqlStatement, criteria.InitiatorID.UUID)
	}
	if criteria.Rating != 0 {
		sqlStatement = findByRating(sqlStatement, criteria.Rating)
	}
	if criteria.CuisineType != 0 {
		sqlStatement = findByCuisineType(sqlStatement, &criteria.CuisineType)
	}
	var rows, err = sqlStatement.Query()
	if err != nil {
		return nil, fmt.Errorf("problem with quering to database %w", err)
	}
	return rows, nil
}

func findByInitiatorID(sql sq.SelectBuilder, id uuid.UUID) sq.SelectBuilder {
	return sql.Where(sq.Eq{"initiators_id": id})
}
func findByRating(sql sq.SelectBuilder, rating int) sq.SelectBuilder {
	return sql.Where(sq.Eq{"rating": rating})
}
func findByAddress(sql sq.SelectBuilder, address *models.Address) sq.SelectBuilder {
	return sql.Where(sq.Eq{"country": address.Country, "city": address.City,
		"street_name": address.StreetName, "house_number": address.HouseNumber,
		"building_number": address.BuildingNumber})
}
func findByCuisineType(sql sq.SelectBuilder, cuisineType *gatheringPlace.CuisineType) sq.SelectBuilder {
	return sql.Where(sq.Eq{"cuisine_type": cuisineType})
}
