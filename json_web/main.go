package main

import (
    "database/sql"
    "encoding/json"
    "fmt"
    "html/template"
    "log"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
)

type User struct {
    Name string `json:"name"`
    Count int `json:"count"`
}

func main() {
    db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        rows, err := db.Query("SELECT GKLX,COUNT(GKLX) AS item_count FROM `xnsdsjj_12345cbs` GROUP BY GKLX HAVING GKLX LIKE '%纠纷%' ORDER BY item_count DESC")
        if err != nil {
            http.Error(w, "Error querying database", http.StatusInternalServerError)
            return
        }
        defer rows.Close()

        var users []User
        for rows.Next() {
            var user User
            if err := rows.Scan(&user.Name, &user.Count); err != nil {
                http.Error(w, "Error scanning row", http.StatusInternalServerError)
                return
            }
            users = append(users, user)
        }

        if err := rows.Err(); err != nil {
            http.Error(w, "Error iterating rows", http.StatusInternalServerError)
            return
        }

        usersJSON, err := json.Marshal(users)
        if err != nil {
            http.Error(w, "Error marshalling JSON", http.StatusInternalServerError)
            return
        }

        const tmpl = `
        <!DOCTYPE html>
        <html>
        <head>
            <title>表格</title>
            <style>
                table {
                    width: 50%;
                    border-collapse: collapse;
                }
                table, th, td {
                    border: 1px solid black;
                }
                th, td {
                    padding: 8px;
                    text-align: left;
                }
                th {
                    background-color: #f2f2f2;
                }
            </style>
        </head>
        <body>
            <h1>表格</h1>
            <table>
                <tr>
                    <th>名称</th>
                    <th>数量</th>
                </tr>
                {{range .}}
                <tr>
                    <td>{{.Name}}</td>
                    <td>{{.Count}}</td>
                </tr>
                {{end}}
            </table>
        </body>
        </html>`

        t := template.New("users")
        t, err = t.Parse(tmpl)
        if err != nil {
            http.Error(w, "Error parsing template", http.StatusInternalServerError)
            return
        }

        var usersData []User
        if err := json.Unmarshal(usersJSON, &usersData); err != nil {
            http.Error(w, "Error unmarshalling JSON", http.StatusInternalServerError)
            return
        }

        err = t.Execute(w, usersData)
        if err != nil {
            http.Error(w, "Error executing template", http.StatusInternalServerError)
            return
        }
    })

    fmt.Println("浏览器访问：http://127.0.0.1")
    log.Fatal(http.ListenAndServe(":80", nil))
}
