package repositories

import (
	"context"
	"fmt"
	"testing"
)

func TestGetRankList(t *testing.T) {
	fmt.Println(GetRankList(context.Background()))
}