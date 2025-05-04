package authentication

import "go.uber.org/fx"

var Module = fx.Module("authentication", githubModule)
