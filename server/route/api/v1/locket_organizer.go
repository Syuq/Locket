package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/Syuq/Locket/internal/util"
	"github.com/Syuq/Locket/store"
)

type LocketOrganizer struct {
	LocketID int32 `json:"locketId"`
	UserID int32 `json:"userId"`
	Pinned bool  `json:"pinned"`
}

type UpsertLocketOrganizerRequest struct {
	Pinned bool `json:"pinned"`
}

func (s *APIV1Service) registerLocketOrganizerRoutes(g *echo.Group) {
	g.POST("/locket/:locketId/organizer", s.CreateLocketOrganizer)
}

// CreateLocketOrganizer godoc
//
//	@Summary	Organize locket (pin/unpin)
//	@Tags		locket-organizer
//	@Accept		json
//	@Produce	json
//	@Param		locketId	path		int							true	"ID of locket to organize"
//	@Param		body	body		UpsertLocketOrganizerRequest	true	"Locket organizer object"
//	@Success	200		{object}	store.Locket					"Locket information"
//	@Failure	400		{object}	nil							"ID is not a number: %s | Malformatted post locket organizer request"
//	@Failure	401		{object}	nil							"Missing user in session | Unauthorized"
//	@Failure	404		{object}	nil							"Locket not found: %v"
//	@Failure	500		{object}	nil							"Failed to find locket | Failed to upsert locket organizer | Failed to find locket by ID: %v | Failed to compose locket response"
//	@Router		/api/v1/locket/{locketId}/organizer [POST]
func (s *APIV1Service) CreateLocketOrganizer(c echo.Context) error {
	ctx := c.Request().Context()
	locketID, err := util.ConvertStringToInt32(c.Param("locketId"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("ID is not a number: %s", c.Param("locketId"))).SetInternal(err)
	}

	userID, ok := c.Get(userIDContextKey).(int32)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Missing user in session")
	}

	locket, err := s.Store.GetLocket(ctx, &store.FindLocket{
		ID: &locketID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to find locket").SetInternal(err)
	}
	if locket == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Locket not found: %v", locketID))
	}
	if locket.CreatorID != userID {
		return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
	}

	request := &UpsertLocketOrganizerRequest{}
	if err := json.NewDecoder(c.Request().Body).Decode(request); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Malformatted post locket organizer request").SetInternal(err)
	}

	upsert := &store.LocketOrganizer{
		LocketID: locketID,
		UserID: userID,
		Pinned: request.Pinned,
	}
	_, err = s.Store.UpsertLocketOrganizer(ctx, upsert)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to upsert locket organizer").SetInternal(err)
	}

	locket, err = s.Store.GetLocket(ctx, &store.FindLocket{
		ID: &locketID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to find locket by ID: %v", locketID)).SetInternal(err)
	}
	if locket == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("Locket not found: %v", locketID))
	}

	locketResponse, err := s.convertLocketFromStore(ctx, locket)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to compose locket response").SetInternal(err)
	}
	return c.JSON(http.StatusOK, locketResponse)
}
