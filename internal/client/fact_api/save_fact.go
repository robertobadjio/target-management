package fact_api

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	modelBuffer "target-management/internal/buffer/model"
	"time"
)

func (c *Client) SaveFact(ctx context.Context, fact *modelBuffer.Fact) error {
	form := url.Values{}
	form.Add("period_start", fact.PeriodStart.Format(time.DateOnly))
	form.Add("period_end", fact.PeriodEnd.Format(time.DateOnly))
	form.Add("period_key", fact.PeriodKey)
	form.Add("indicator_to_mo_id", strconv.Itoa(fact.IndicatorToMoID))
	form.Add("indicator_to_mo_fact_id", strconv.Itoa(fact.IndicatorToMoFactID))
	form.Add("value", strconv.Itoa(fact.Value))
	form.Add("fact_time", fact.FactTime.Format(time.DateOnly))
	form.Add("is_plan", strconv.Itoa(fact.IsPlan))
	form.Add("auth_user_id", strconv.Itoa(fact.AuthUserID))
	form.Add("comment", fact.Comment)

	req, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/facts/save_fact", c.BaseURL), strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}

	req = req.WithContext(ctx)

	if err = c.sendRequest(req); err != nil {
		return err
	}

	return nil
}
