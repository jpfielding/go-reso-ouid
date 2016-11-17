# go-reso-ouid
RESO's OUID api (go)

https://github.com/jpfielding/reso-ouid/blob/master/organizations.json
```
	ctx := context.Background()
	city := ouid.ByCity("dallas")
	active := ouid.ByActive(true)
	scope := ouid.And(city, active)
	orgs := ouid.Organizations{}
	err := ouid.Process(ctx, cfg.Request(scope), func(org ouid.Organization, err error) error {
		orgs.Organization = append(orgs.Organization, org)
		return nil
	})

```
