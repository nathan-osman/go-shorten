package main

// Template for the admin page.
const adminTemplate = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <link href="https://fonts.googleapis.com/css?family=Roboto:300|Roboto+Mono:300" rel="stylesheet" type="text/css">
        <style>
            body { font-family: "Roboto", sans; }
            input[type=text] { font-family: "Roboto Mono" monospace; font-size: 12pt; margin: 8px 0; width: 256px; }
            input[type=text]:focus { outline: none; }
            label { display: block; margin: 16px 0; }
            label input { display: block; border: none; border-bottom: 1px solid #ccc; }
            table { border: 1px solid #ccc; border-collapse: collapse; width: 100%; }
            tr:nth-child(odd) { background-color: #eee; }
            th { background-color: #ddd; text-align: left; }
            th, td { padding: 4px 8px; }
            td.break { word-break: break-word; }
            td.right { text-align: right; }
            .message { padding: 4px 8px; }
            .message.info { background-color: #def; border: 1px solid #5cf; }
            .message.error { background-color: #fdd; border: 1px solid #f55; }
            @media (min-width:768px) {
                .container { margin: auto; width: 640px; }
            }
        </style>
    </head>
    <body>
        <div class="container">
            {{ range $m := .messages }}
                <div class="message {{ $m.Type }}">{{ $m.Body }}</div>
            {{ end }}
            <h1>go-shorten</h1>
            <p>The redirects are displayed in the table below.</p>
            <table>
                <tr>
                    <th>Path</th>
                    <th>Destination</th>
                    <th></th>
                </tr>
                {{ range $k, $v := .database.Paths }}
                    <tr>
                        <td class="break">{{ $k }}</td>
                        <td class="break">
                            <a href="{{ $v }}" target="_blank">{{ $v }}</a>
                        </td>
                        <td class="right">
                            <form method="post">
                                <input type="hidden" name="action" value="delete">
                                <input type="hidden" name="path" value="{{ $k }}">
                                <button type="submit">Delete</button>
                            </form>
                        </td>
                    </tr>
                {{ end }}
            </table>
            <h2>Add New</h2>
            <form method="post">
                <input type="hidden" name="action" value="new">
                <label>
                    Path: <input type="text" name="path">
                </label>
                <label>
                    Destination: <input type="text" name="destination">
                </label>
                <button type="submit">Add</button>
            </form>
        </div>
    </body>
</html>
`
