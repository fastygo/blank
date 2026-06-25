package hero

// DefaultProps returns English wireframe demo data for isolated previews.
func DefaultProps() Props {
	return Props{
		Welcome:      "Welcome",
		WelcomeBrand: "to FastyGo",
		Description:  "Minimal Go + templ starter with mobile sheet, dark theme, and locale switching.",
	}
}
