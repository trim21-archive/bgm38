package docs

import "fmt"

func OpenAPI() string {
	return (&s{}).ReadDoc()
}

func SwagHTML(path string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
<link type="text/css" rel="stylesheet" href="https://cdn.jsdelivr.net/npm/swagger-ui-dist@3/swagger-ui.css">
<link rel="shortcut icon" href="https://fastapi.tiangolo.com/img/favicon.png">
<title>Pol server - Swagger UI</title>
</head>
<body>
<div id="swagger-ui">
</div>
<script src="https://cdn.jsdelivr.net/npm/swagger-ui-dist@3/swagger-ui-bundle.js"></script>
<script>
const ui = SwaggerUIBundle({
	url: '%s',

	dom_id: '#swagger-ui',
	presets: [
	SwaggerUIBundle.presets.apis,
	SwaggerUIBundle.SwaggerUIStandalonePreset
	],
	layout: "BaseLayout",
	deepLinking: true
})
</script>
</body>
</html>
`, path)
}
