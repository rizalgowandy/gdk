package cronx

const statusPageTemplate = `
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>Cron Status</title>
	<style>
        body {
            margin: 30px 0 0 0;
            font-size: 16px;
            font-family: sans-serif;
            color: #345;
        }

        h1 {
            font-size: 36px;
            text-align: center;
            padding: 10px 0 30px 0;
        }

        table {
            margin: 0 auto;
            border-collapse: collapse;
            border: none;
        }

        table td, table th {
            min-width: 25px;
            width: auto;
            padding: 15px 20px;
            border: none;
        }

        table tr:nth-child(odd) {
            background-color: #F0F0F0;
        }

        table tr:nth-child(1) {
            background-color: #345;
            color: white;
        }

        th {
            text-align: left;
        }
	</style>
	<title>Cron Status</title>
</head>
<body>
<h1>Cron Status</h1>
<table>
	<tr>
		<th>ID</th>
		<th>Name</th>
		<th>Status</th>
		<th>Last run</th>
		<th>Next run</th>
		<th>Latency</th>
	</tr>
	{{range .}}
	<tr>
		<td>{{.ID}}</td>
		<td>{{.Job.Name}}</td>
		<td>{{.Job.Status}}</td>
		<td>{{if not .Prev.IsZero}}{{.Prev.Format "2006-01-02 15:04:05"}}{{end}}</td>
		<td>{{if not .Next.IsZero}}{{.Next.Format "2006-01-02 15:04:05"}}{{end}}</td>
		<td>{{.Job.Latency}}</td>
	</tr>
	{{end}}
</table>
</body>
</html>
`
