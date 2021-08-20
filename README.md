# Setup
## jk:note$pwd
## /home/jk/dev/go/jk.com/note
## $execute db.sql on psql
## $go build
## $go run server.go
#

# Unit Test
## jk:note$go test -v
#### === RUN   TestNoteList
#### --- PASS: TestNoteList (0.00s)
#### === RUN   TestNoteGet
#### --- PASS: TestNoteGet (0.00s)
#### === RUN   TestNoteCreate
#### --- PASS: TestNoteCreate (0.00s)
#### === RUN   TestNoteUpdate
#### --- PASS: TestNoteUpdate (0.00s)
#### === RUN   TestNoteDelete
#### --- PASS: TestNoteDelete (0.00s)
#### PASS
#### ok      jk.com/note     0.003s
### jk:note$
#
# Tests Using Curl
## $curl -X POST -i localhost:8080/notes --data '{"content":"how to merge video with text", "author":"John Doe"}'
### {"id":1,"content":"how to merge video with text","author":"John Doe"}

## $curl -X GET -i localhost:8080/notes 
### [{"id":1,"content":"how to merge video with text","author":"John Doe"}]

## $curl -X GET -i localhost:8080/notes/1
#### {"id":1,"content":"how to merge video with text","author":"John Doe"}

## $curl -X PATCH -i localhost:8080/notes/1 --data '{"content":"how to fly with video", "author":"Paul Young"}'
### {"id":1,"content":"how to fly with video","author":"Paul Young"}

## $curl -X GET -i localhost:8080/notes 
### [{"id":1,"content":"how to fly with video","author":"Paul Young"}]

## $curl -X DELETE -i localhost:8080/notes/1

## $curl -X GET -i localhost:8080/notes 
### null
#