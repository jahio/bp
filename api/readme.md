# An app to track blood pressure

This is just a hobby app that can be used locally to track your blood
pressure. Just start the app (recommended: start it on boot) and go
to the URL in your browser to log your blood pressure and view history.

Work in progress.

## Dev Notes (delete at some point)

`curl -i -X POST -d @./entry.json -H "Content-Type: application/json" http://localhost:9000/entries/new`
