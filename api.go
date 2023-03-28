package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "strconv"
    "strings"

    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
)

type Patient struct {
    FullName string `json:"fullName"`
    CPF      string `json:"cpf"`
}

func main() {
    // Inicia o servidor Gin
    r := gin.Default()

    // Define o endpoint para cadastrar um paciente
    r.POST("/cadastra-paciente", cadastraPaciente)

    // Inicia o banco de dados
    db, err := sql.Open("sqlite3", "pacientes.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // Cria a tabela 'pacientes', se não existir
    _, err = db.Exec("CREATE TABLE IF NOT EXISTS pacientes (id INTEGER PRIMARY KEY, full_name TEXT, cpf TEXT)")
    if err != nil {
        log.Fatal(err)
    }

    // Inicia o servidor
    err = r.Run(":8080")
    if err != nil {
        log.Fatal(err)
    }
}

func cadastraPaciente(c *gin.Context) {
    // Lê o JSON enviado no corpo da requisição
    var patient Patient
    err := c.BindJSON(&patient)
    if err != nil {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    // Valida o CPF
    cpf := strings.ReplaceAll(patient.CPF, ".", "")
    cpf = strings.ReplaceAll(cpf, "-", "")
    if len(cpf) != 11 {
        c.AbortWithStatus(http.StatusBadRequest)
        return
    }

    // Insere o paciente no banco de dados
    db, err := sql.Open("sqlite3", "pacientes.db")
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }
    defer db.Close()

    stmt, err := db.Prepare("INSERT INTO pacientes (full_name, cpf) VALUES (?, ?)")
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }
    defer stmt.Close()

    _, err = stmt.Exec(patient.FullName, cpf)
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }

    // Retorna o ID do paciente cadastrado
    var id int64
    err = db.QueryRow("SELECT last_insert_rowid()").Scan(&id)
    if err != nil {
        c.AbortWithError(http.StatusInternalServerError, err)
        return
    }

    c.JSON(http.StatusOK, gin.H{"id": strconv.FormatInt(id, 10)})
}