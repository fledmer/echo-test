data "composite_schema" "app" {
  schema "public" {
    url = "ent://ent/schema"
  }
}

env "local" {
  src = data.composite_schema.app.url
  dev = "docker://postgres/16/dev?search_path=public"
  url = getenv("DATABASE_URL")

  migration {
    dir = "file://migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}

env "ci" {
  src = data.composite_schema.app.url
  dev = "docker://postgres/16/dev?search_path=public"

  migration {
    dir = "file://migrations"
  }

  lint {
    git {
      base = "main"
    }
  }
}
