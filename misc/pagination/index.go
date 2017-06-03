package pagination

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/twalkapp/server/models/misc"
	"github.com/twalkapp/server/misc/config"
)

func GetPagination(c *gin.Context) (misc.Pagination, error) {
	result := misc.Pagination{Current: -1, Size: -1, Next: -1, Prev: -1, Last: -1}
	current, err := strconv.Atoi(c.DefaultQuery("page", "-1"))
	if err != nil {
		return result, err
	}
	size, err := strconv.Atoi(c.DefaultQuery("per_page", strconv.Itoa(config.Conf.Pagination.Size)))
	if err != nil {
		return result, err
	}
	result.Current = current
	result.Size = size
	return result, nil
}
