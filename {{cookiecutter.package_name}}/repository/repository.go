package repository

import (
	resources "github.com/{{cookiecutter.org_name}}/{{cookiecutter.package_name}}/resources"
)

type ClientRepository interface {
	GetById(id int) (resources.Client, bool)
}
