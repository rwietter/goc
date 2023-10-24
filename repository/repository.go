package repository

type Repository struct {
	Path   string `json:"path"`
	GitHub string `json:"github"`
	Name   string `json:"name"`
}

var repositories []Repository

func GetRepositories() []Repository {
	return repositories
}

func AddRepository(repository Repository) {
	repositories := GetRepositories()
	repositories = append(repositories, repository)
}

func RemoveRepository(repository Repository) {
	repositories := GetRepositories()
	for i, r := range repositories {
		if r.Path == repository.Path {
			repositories = append(repositories[:i], repositories[i+1:]...)
		}
	}
}

func GetRepositoryByPath(path string) Repository {
	repositories := GetRepositories()
	for _, r := range repositories {
		if r.Path == path {
			return r
		}
	}
	return Repository{}
}
