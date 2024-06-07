package main

import (
    "database/sql"
    "encoding/json"
    // "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"

    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    ID int `json:"id"`
    Name string `json:"name"`
    Gender string `json:"gender"`
    Sfz string `json:"sfz"`
}

func main() {
    http.HandleFunc("/user", userHandler)
    log.Println("Server starting on port 80")
    log.Fatal(http.ListenAndServe(":80", nil))
}

func userHandler(w http.ResponseWriter, r *http.Request) {
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer db.Close()

    // rows, err := db.Query("SELECT ID, name, gender, sfz FROM web")
    // if err != nil {
    //     http.Error(w, err.Error(), http.StatusInternalServerError)
    //     return
    // }
    // defer rows.Close()

    idsParam := r.URL.Query()["id"]
    namesParam := r.URL.Query()["name"]

    var ids []int
    for _, idStr := range idsParam {
        id, err := strconv.Atoi(idStr)
        if err != nil {
            http.Error(w, "Invalid 'id' parameter", http.StatusBadRequest)
            return
        }
        ids = append(ids, id)
    }

    query := "select id, name, gender, sfz from user where 1=1"
    var args []interface{}

    if len(ids) > 0 {
        query += " AND id IN (?" + strings.Repeat(",?", len(ids)-1) + ")"
        for _, id := range ids {
            args = append(args, id)
        }
    }

    if len(namesParam) > 0 {
        query += " and name in (?" + strings.Repeat(",?", len(namesParam)-1) + ")"
        for _, name := range namesParam {
            args = append(args, name)
        }
    }

    rows, err := db.Query(query, args...)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var user User
        if err := rows.Scan(&user.ID, &user.Name, &user.Gender, &user.Sfz); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        users = append(users, user)
    }

    if err := rows.Err(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(users); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
