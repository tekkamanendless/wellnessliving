package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/sirupsen/logrus"
	"github.com/tekkamanendless/wellnessliving"
)

func main() {
	ctx := context.Background()

	logrus.SetLevel(logrus.DebugLevel)

	businessID := os.Getenv("BUSINESS_ID")

	client := wellnessliving.Client{}

	logrus.WithContext(ctx).Debugf("Events:\n")
	values := url.Values{}
	values.Set("id_flag", "3")
	values.Set("is_ignore_requirement", "")
	values.Set("is_tab_all", "1")
	values.Set("k_business", businessID)
	values.Set("text_search", "")

	contents, err := client.Raw(ctx, http.MethodGet, "/Wl/Event/EventList.json", values)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Error: [%T] %v\n", err, err)
	}
	fmt.Printf("%s\n", contents)

	var eventListResponse wellnessliving.EventListResponse
	err = client.Request(ctx, http.MethodGet, "/Wl/Event/EventList.json", values, &eventListResponse)
	if err != nil {
		logrus.WithContext(ctx).Errorf("Error: [%T] %v\n", err, err)
	}
	spew.Dump(eventListResponse)
}
