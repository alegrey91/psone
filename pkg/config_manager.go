import (
    "gopkg.in/yaml.v2"
)

type Psone struct {
    Name string
    Value string
}

type Configs struct {
    Psones []Psone `psones`
}
