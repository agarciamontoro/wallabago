package wallabago

// Version returns the version of the configured wallabag instance
func Version(bodyGetterFunc BodyGetter) string {
	v := bodyGetterFunc(Config.WallabagURL + "/api/version")
	// strip of the quotation marks from the version string being return from the API
	return v[1 : len(v)-1]
}
