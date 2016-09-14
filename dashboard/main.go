package main

import (
	"html/template"
	"os"
	"time"
)

func main() {
	dashboardTemplate := `<html>
<body>
{{ .GeneratedDate }}
</body>
</html>
`

	t, err := template.New("dashboard").Parse(dashboardTemplate)

	if err != nil {
		panic(err)
	}

	values := struct {
		GeneratedDate string
	}{
		GeneratedDate: time.Now().Format("Jan 2 2006"),
	}

	t.Execute(os.Stdout, values)
}
