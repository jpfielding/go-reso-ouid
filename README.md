# go-reso-ouid
RESO's OUID api (go)
```
	ctx := context.Background()
	city := ouid.ByCity("dallas")
	active := ouid.ByActive(true)
	scope := ouid.And(all, city, active)
	orgs := ouid.Organizations{}
	err := ouid.Process(ctx, cfg.Request(scope), func(org ouid.Organization, err error) error {
		orgs.Organization = append(orgs.Organization, org)
		return nil
	})

```
