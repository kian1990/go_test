package main

import (
    "log"
    "context"
    "github.com/beltran/gohive"
)

func main() {
    configuration := gohive.NewConnectConfiguration()
    connection, errConn := gohive.Connect("172.28.166.223", 10000, "NONE", configuration)
    if errConn != nil {
        log.Fatal(errConn)
    }
    cursor := connection.Cursor()
    ctx := context.Background()

    cursor.Exec(ctx, "SELECT web_ranking, web_id, web_url, web_type FROM web")
    if cursor.Err != nil {
        log.Fatal(cursor.Err)
    }

    var web_ranking int32
    var web_id int64
    var web_url string
    var web_type string
    for cursor.HasMore(ctx) {
        cursor.FetchOne(ctx, &web_ranking, &web_id, &web_url, &web_type)
        if cursor.Err != nil {
            log.Fatal(cursor.Err)
        }
        log.Println(web_ranking, web_id, web_url, web_type)
    }
    cursor.Close()
    connection.Close()
}