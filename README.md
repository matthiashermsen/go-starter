# Go Starter

Quickly get started with a minimal Golang starter template!

*This template is strongly opinionated*

## Features

- Bare Golang application
- Release management using [GoReleaser](https://goreleaser.com/)
- Github container registry for [Docker](https://www.docker.com/)
- Updates using [Dependabot](https://github.com/dependabot)
- CI workflows using [Github Actions](https://github.com/features/actions)
    - Lint Pull Request titles following the [Conventional Commits specification](https://www.conventionalcommits.org/)
    - Run QA for Pull Requests ( analyze, Tests, build and release in drymode )
    - Release on push on `main` branch ( Tagging, Publish binaries and Docker image )

## Getting started

[Create a new repository from this template on GitHub](https://github.com/matthiashermsen/go-starter/generate).

### Checklist

When you use this template, try follow the checklist to update your info properly

- [ ] Change the module path in [go.mod](./go.mod).
- [ ] Fix the code in [main.go](./main.go).
- [ ] Change the Github container registry owner in [.goreleaser.yaml](./.goreleaser.yaml).
- [ ] Change the binary name in [Dockerfile](./Dockerfile).
- [ ] Change the `assignees` in [dependabot.yml](./.github/dependabot.yml).
- [ ] Change the author name in [LICENSE](./LICENSE).
- [ ] Clean up the [README](./README.md).
