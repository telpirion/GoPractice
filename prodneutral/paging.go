package prodneutral

import (
	"context"
	"log"

	iam "cloud.google.com/go/iam/admin/apiv1"
	adminpb "google.golang.org/genproto/googleapis/iam/admin/v1"
)

func usingPaging() error {
	ctx := context.Background()

	c, err := iam.NewIamClient(ctx)
	if err != nil {
		return err
	}
	defer c.Close()

	req := &adminpb.ListRolesRequest{
		Parent:   "",
		PageSize: 5,
	}

	for {
		resp, err := c.ListRoles(ctx, req)

		if err != nil {
			return err
		}

		for _, role := range resp.Roles {
			log.Println(role)
		}

		t := resp.NextPageToken

		if t == "" {
			break
		}
		req = &adminpb.ListRolesRequest{
			Parent:    "",
			PageSize:  5,
			PageToken: t,
		}
	}

	return nil
}
