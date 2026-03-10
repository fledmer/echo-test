data "composite_schema" "app" {
  schema "public" {
    url = "ent://ent/schema"
  }
}

env "local" {
  src = data.composite_schema.app.url
  dev = "docker://postgres/16/dev?search_path=public"

  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "docker" {
  src = data.composite_schema.app.url
  url = "postgres://postgres:postgres@db:5432/echodb?sslmode=disable"

  migration {
    dir = "file://migrations"
  }
}
