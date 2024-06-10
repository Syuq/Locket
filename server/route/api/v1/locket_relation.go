package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Syuq/Locket/internal/util"
	"github.com/Syuq/Locket/store"
)

type LocketRelationType string

const (
	LocketRelationReference LocketRelationType = "REFERENCE"
	LocketRelationComment   LocketRelationType = "COMMENT"
)

func (t LocketRelationType) String() string {
	return string(t)
}

type LocketRelation struct {
	LocketID        int32            `json:"locketId"`
	RelatedLocketID int32            `json:"relatedLocketId"`
	Type          LocketRelationType `json:"type"`
}

type UpsertLocketRelationRequest struct {
	RelatedLocketID int32            `json:"relatedLocketId"`
	Type          LocketRelationType `json:"type"`
}

func (s *APIV1Service) registerLocketRelationRoutes(g *echo.Group) {
	g.GET("/locket/:locketId/relation", s.GetLocketRelationList)
	g.POST("/locket/:locketId/relation", s.CreateLocketRelation)
	g.DELETE("/locket/:locketId/relation/:relatedLocketId/type/:relationType", s.DeleteLocketRelation)
}

// GetLocketRelationList godoc
//
//	@Summary	Get a list of Locket Relations
//	@Tags		locket-relation
//	@Accept		json
//	@Produce	json
//	@Param		locketId	path		int						true	"ID of locket to find relations"
//	@Success	200		{object}	[]store.LocketRelation	"Locket relation information list"
//	@Failure	400		{object}	nil						"ID is not a number: %s"
//	@Failure	500		{object}	nil						"Failed to list locket relations"
//	@Router		/api/v1/locket/{locketId}/relation [GET]
func (s *APIV1Service) GetLocketRelationList(c echo.Context) error {
	ctx := c.Request().Context()
	locketID, err := util.ConvertStringToInt32(c.Param("locketId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("locketId"))).SetInternal(err)
	}

	locketRelationList, err := s.Store.ListLocketRelations(ctx, &store.FindLocketRelation{
		LocketID: &locketID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to list locket relations").SetInternal(err)
	}
	return c.JSON(http.StatusOK, locketRelationList)
}

// CreateLocketRelation godoc
//
//	@Summary		Create Locket Relation
//	@Description	Create a relation between two lockets
//	@Tags			locket-relation
//	@Accept			json
//	@Produce		json
//	@Param			locketId	path		int							true	"ID of locket to relate"
//	@Param			body	body		UpsertLocketRelationRequest	true	"Locket relation object"
//	@Success		200		{object}	store.LocketRelation			"Locket relation information"
//	@Failure		400		{object}	nil							"ID is not a number: %s | Malformatted post locket relation request"
//	@Failure		500		{object}	nil							"Failed to upsert locket relation"
//	@Router			/api/v1/locket/{locketId}/relation [POST]
//
// NOTES:
// - Currently not secured
// - It's possible to create relations to lockets that doesn't exist, which will trigger 404 errors when the frontend tries to load them.
// - It's possible to create multiple relations, though the interface only shows first.
func (s *APIV1Service) CreateLocketRelation(c echo.Context) error {
	ctx := c.Request().Context()
	locketID, err := util.ConvertStringToInt32(c.Param("locketId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("locketId"))).SetInternal(err)
	}

	request := &UpsertLocketRelationRequest{}
	if err := json.NewDecoder(c.Request().Body).Decode(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Malformatted post locket relation request").SetInternal(err)
	}

	locketRelation, err := s.Store.UpsertLocketRelation(ctx, &store.LocketRelation{
		LocketID:        locketID,
		RelatedLocketID: request.RelatedLocketID,
		Type:          store.LocketRelationType(request.Type),
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to upsert locket relation").SetInternal(err)
	}
	return c.JSON(http.StatusOK, locketRelation)
}

// DeleteLocketRelation godoc
//
//	@Summary		Delete a Locket Relation
//	@Description	Removes a relation between two lockets
//	@Tags			locket-relation
//	@Accept			json
//	@Produce		json
//	@Param			locketId			path		int					true	"ID of locket to find relations"
//	@Param			relatedLocketId	path		int					true	"ID of locket to remove relation to"
//	@Param			relationType	path		LocketRelationType	true	"Type of relation to remove"
//	@Success		200				{boolean}	true				"Locket relation deleted"
//	@Failure		400				{object}	nil					"Locket ID is not a number: %s | Related locket ID is not a number: %s"
//	@Failure		500				{object}	nil					"Failed to delete locket relation"
//	@Router			/api/v1/locket/{locketId}/relation/{relatedLocketId}/type/{relationType} [DELETE]
//
// NOTES:
// - Currently not secured.
// - Will always return true, even if the relation doesn't exist.
func (s *APIV1Service) DeleteLocketRelation(c echo.Context) error {
	ctx := c.Request().Context()
	locketID, err := util.ConvertStringToInt32(c.Param("locketId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Locket ID is not a number: %s", c.Param("locketId"))).SetInternal(err)
	}
	relatedLocketID, err := util.ConvertStringToInt32(c.Param("relatedLocketId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Related locket ID is not a number: %s", c.Param("relatedLocketId"))).SetInternal(err)
	}
	relationType := store.LocketRelationType(c.Param("relationType"))

	if err := s.Store.DeleteLocketRelation(ctx, &store.DeleteLocketRelation{
		LocketID:        &locketID,
		RelatedLocketID: &relatedLocketID,
		Type:          &relationType,
	}); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to delete locket relation").SetInternal(err)
	}
	return c.JSON(http.StatusOK, true)
}

func convertLocketRelationFromStore(locketRelation *store.LocketRelation) *LocketRelation {
	return &LocketRelation{
		LocketID:        locketRelation.LocketID,
		RelatedLocketID: locketRelation.RelatedLocketID,
		Type:          LocketRelationType(locketRelation.Type),
	}
}
