package users_transport_http

import (
	"net/http"

	core_logger "github.com/SlavaculoN/golang-todoapp/internal/core/logger"
	core_http_response "github.com/SlavaculoN/golang-todoapp/internal/core/transport/http/response"
	core_http_utils "github.com/SlavaculoN/golang-todoapp/internal/core/transport/http/utils"
)

type GetUserPesponse UserDTOResponse

func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHttpResponseHandler(log, rw)

	userID, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed tp get userID path value",
		)

		return
	}

	user, err := h.usersService.GetUser(ctx, userID)
	if err != nil {
		responseHandler.ErrorResponse(
			err,
			"failed to get user",
		)

		return
	}

	response := GetUserPesponse(userDTOFromDomain(user))

	responseHandler.JSONResponse(response, http.StatusOK)
}
