package main

import (
        "fmt"
	"log"
        "gopkg.in/mgo.v2"
        "gopkg.in/mgo.v2/bson"
)

type Movie struct {
        Name string `bson:"name"`
        Year string `bson:"year"`
        Directors []string `bson:"directors"`
        Writers []string  `bson:"writers"`
        BoxOffice `bson:"boxOffice"`
}

type BoxOffice struct {
        Budget uint64 `bson:"budget"`
        Gross uint64  `bson:"gross"`
}

func main() {
        session, err := mgo.Dial("127.0.0.1")
        if err != nil {
                panic(err)
        }
        defer session.Close()

        c := session.DB("appdb").C("movies")

        darkNight := &Movie{
                Name: "The Dark Knight",
                Year: "2008",
                Directors: []string{"Christopher Nolan"},
                Writers: []string{"Jonathan Nolan", "Christopher Nolan"},
                BoxOffice: BoxOffice{
                        Budget: 185000000,
                        Gross: 533316061,
                },
        }
        err = c.Insert(darkNight)
        if err != nil {
                log.Fatal(err)
        }

        result := Movie{}
        err = c.Find(bson.M{"boxOffice.budget": bson.M{"$gt": 150000000}}).One(&result)
        if err != nil {
                log.Fatal(err)
        }

        fmt.Println("Movie:", result.Name)
}