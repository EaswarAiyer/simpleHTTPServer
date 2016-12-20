/* Before you execute the program, Launch `cqlsh` and execute:
create keyspace example with replication = { 'class' : 'SimpleStrategy', 'replication_factor' : 1 };
create table example.tweet(timeline text, id UUID, text text, PRIMARY KEY(id));
create index on example.tweet(timeline);
*/
package main

import (
	"log"
	"net/http"
)

func init() {

}

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}