# Golang Solution Advent of Code

Folders have already solved golang solutions. They also contain the Testdata which i got from the site.

## Generator for new Tasks

```
./bin/aoc -y year -d day
```

Creates a new Folder for the year and day. It also downloads the data for the task and creates a new main.go file from the template which also loads the data from the cache. 

Important. You must fill the session_id file in the root. There is already a dist file which you copy and fill with you session_id from the cookie of the site.
