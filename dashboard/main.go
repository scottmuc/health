package main

import (
	"html/template"
	"os"
	"time"
)

func main() {
	dashboardTemplate := `<!DOCTYPE html>
<html>
<head>
<meta charset="UTF-8">
<title>Scott Muc's Health Dashboard</title>
<style type="text/css">

.tile {
    height: 300px;
    width: 300px;
    text-align: center;

    margin: 20px 0 20px 20px;
    padding: 0;

    float: left;
}

.tile span {
    font-size: 96px;
}

.failed {
    background-color: #aa0000;
    color: #eeeeee;
}

.success {
    background-color: #00aa00;
    color: #eeeeee;
}

.warning {
    background-color: #ffaa00;
    color: #eeeeee;  
}

.todo {
    background-color: #0000aa;
    color: #eeeeee;
}

footer {
    display: block;
    position: absolute;
    right: 20px;
    top: 20px;
}

h2 { clear:both; }

</style>
</head>
<body>

<h2>Physical Health</h2>

<div class="tile success">
    <h3>days since last exercise activity</h3>
    <span>2</span>
</div>

<div class="tile success">
    <h3>meals cooked in last week</h3>
    <span>5</span>
</div>

<div class="tile failed">
    <h3>weight</h3>
    <span>91</span>
</div>

<div class="tile todo">
    <h3>alcoholic drinks consumed last week</h3>
    <span></span>
</div>

<h2>Mental Health</h2>

<div class="tile failed">
    <h3>pocket reading queue</h3>
    <span>27</span>
</div>

<div class="tile failed">
    <h3>podcast episodes queued</h3>
    <span>20+ hours</span>
</div>

<div class="tile todo">
    <h3>number of pomodoros spent writing articles/blog posts last week</h3>
    <span></span>
</div>

<div class="tile todo">
    <h3>thinking of more ideas for this</h3>
    <span></span>
</div>

<h2>Financial Health</h2>

<div class="tile warning">
    <h3>growth in networth last month (percentage relative to financial independence target)</h3>
    <span>0.002%</span>
</div>

<div class="tile failed">
    <h3>number of months income that's liquid (and not invested)</h3>
    <span>10</span>
</div>

<div class="tile success">
    <h3>value of liabilities</h3>
    <span>0</span>
</div>

<div class="tile todo">
    <h3>percentage of income saved last month</h3>
    <span></span>
</div>

<footer>
last updated: {{ .GeneratedDate }}
</footer>
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
