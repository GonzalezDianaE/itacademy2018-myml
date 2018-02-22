package addresses

import (
    "github.com/gin-gonic/gin"
    "strconv"
    "net/http"
    "github.com/mercadolibre/itacademy-myml/src/api/util/apierrors"
    addressesService "github.com/mercadolibre/itacademy-myml/src/api/services/addresses"
)

func GetAddressByID(c *gin.Context) {
    addressID, err := strconv.ParseInt(c.Param("addressID"), 10, 64)
    if err != nil {
        c.JSON(http.StatusBadRequest, apierrors.ApiError{
            StatusCode: http.StatusBadRequest,
            Error: err.Error(),
        })
        return
    }

    address, apiErr := addressesService.GetAddressByID(addressID)
    if apiErr != nil {
        c.JSON(apiErr.StatusCode, apiErr)
        return
    }

    c.JSON(http.StatusOK, address)
}