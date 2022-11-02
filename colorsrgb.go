package color

func ColorRGB(r, g, b string) string {
	var color = "\x1b[38;2;" + r + ";" + g + ";" + b + "m"
	return color
}
