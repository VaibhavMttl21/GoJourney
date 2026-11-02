package main

import (
    "fmt"
    "io"
    "os"

    "ariga.io/atlas-provider-gorm/gormschema"

    "course-api/internal/repository/postgres"
)

func main() {

    stmts, err := gormschema.New("postgres").Load(
        &postgres.CourseEntity{},
    )

    if err != nil {
        fmt.Fprintf(
            os.Stderr,
            "failed to load GORM schema: %v\n",
            err,
        )
        os.Exit(1)
    }

    io.WriteString(os.Stdout, stmts)
}