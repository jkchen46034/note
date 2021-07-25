jk:note$pwd
/home/jk/dev/go/jk.com/note
$go build
$go run server.go

$curl -X POST -i localhost:8080/notes --data '{"content":"how to merge video with text", "author":"John Doe"}'
{"id":1,"content":"how to merge video with text","author":"John Doe"}

$curl -X GET -i localhost:8080/notes 
[{"id":1,"content":"how to merge video with text","author":"John Doe"}]

$curl -X GET -i localhost:8080/notes/1
{"id":1,"content":"how to merge video with text","author":"John Doe"}

$curl -X PATCH -i localhost:8080/notes/1 --data '{"content":"how to fly with video", "author":"Paul Young"}'
{"id":1,"content":"how to fly with video","author":"Paul Young"}

$curl -X GET -i localhost:8080/notes 
[{"id":1,"content":"how to fly with video","author":"Paul Young"}]

$curl -X DELETE -i localhost:8080/notes/1

$curl -X GET -i localhost:8080/notes 
null

 



