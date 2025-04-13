package repository

import (
	"dolott_profile/internal/types"
	"dolott_profile/internal/utils"
	"fmt"
)

func (c *profileRepository) ChangeImpressionAndDCoin(userId int32, impression int32, dCredit int32) *types.Error {
	var query string
	computedRank := utils.ComputeRank(uint32(impression))
	query = `UPDATE profiles 
SET 
    impression = impression + $1, 
    d_coin = d_coin + $2, 
    rank = GREATEST(1, rank - $3) 
WHERE user_id = $4;
`
	_, err := c.db.Exec(query, impression, dCredit, computedRank, userId)
	if err != nil {
		fmt.Println(err)
		return types.NewInternalError("failed to increment ranks #3008")
	}
	return nil
}
