# Advision

## Introduction

ADVision is a REST service capable of converting a given image into detailed marketing suggestions. This tool leverages advanced image recognition powered by the OpenAI API to analyze visual content and provide actionable insights for marketing strategies.
As of now, it serves as the backend of [the Deployed Application]().

## How it works

It uses the OpenAI API to analyze the content of Query parameter and generates a detailed report with marketing suggestions. The service is capable of identifying objects, people, and other elements in the image, and provides recommendations on how to leverage this information for marketing purposes.

## Example

(Insert video here)

## Usage

The root endpoint `/` redirects to a swagger documentation page where you can test the API. Try it directly from [here]()

## Installation

You can run the api locally the way you'd do it with any go project.

```bash
go mod download

```

Create a .env file in the root directory and add the following variables:

```bash
OPENAI_API_KEY= # your_openai_api_key
ENVIRONMENT= # 'prod' or 'dev'
```

Then you can run the server with:

```bash
# install swaggo globally to generate swagger documentation on build
go install github.com/swaggo/swag/cmd/swag

# If building to prod(strips debug symbols)
make build-prod
# else you can use
make build

# run the server
make run
```

### Tools used

- Go 1.23.0
- [Gin](github.com/gin-gonic/gin) (Routing)
- [Swaggo](github.com/swaggo/swag) (Swagger Documentation)
- [go-openai](github.com/sashabaranov/go-openai) (Wrapper over the OpenAI API)

### Implementation Status

- [ ] Regenerate specific field suggestions
- [ ] Persist suggestions generation in a database
- [ ] Implement limiter
- [ ] Improve prompting
- [ ] Tune Model with datasets to generate accurate suggestions
- [ ] Improve Logging system
