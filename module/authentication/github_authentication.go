package authentication

import (
	"github.com/google/go-github/v71/github"
	"go.uber.org/fx"
)

var githubModule = fx.Module("github", fx.Provide(newGithubClient))

func newGithubClient() *github.Client {
	return github.NewClientWithEnvProxy()
}
