<h1 align="center">Backend Test Majoo</h1>

<p align="center">
  <img alt="Github top language" src="https://img.shields.io/github/languages/top/bangadam/backend-test-majoo?color=56BEB8">

  <img alt="Github language count" src="https://img.shields.io/github/languages/count/bangadam/backend-test-majoo?color=56BEB8">

  <img alt="Repository size" src="https://img.shields.io/github/repo-size/bangadam/backend-test-majoo?color=56BEB8">

  <img alt="License" src="https://img.shields.io/github/license/bangadam/backend-test-majoo?color=56BEB8">

  <!-- <img alt="Github issues" src="https://img.shields.io/github/issues/bangadam/backend-test-majoo?color=56BEB8" /> -->

  <!-- <img alt="Github forks" src="https://img.shields.io/github/forks/bangadam/backend-test-majoo?color=56BEB8" /> -->

  <!-- <img alt="Github stars" src="https://img.shields.io/github/stars/bangadam/backend-test-majoo?color=56BEB8" /> -->
</p>

<!-- Status -->

<!-- <h4 align="center"> 
	🚧  Backend Test Majoo 🚀 Under construction...  🚧
</h4> 

<hr> -->

<p align="center">
  <a href="#dart-about">About</a> &#xa0; | &#xa0; 
  <a href="#sparkles-features">List API</a> &#xa0; | &#xa0;
  <a href="#rocket-technologies">Technologies</a> &#xa0; | &#xa0;
  <a href="#white_check_mark-requirements">Requirements</a> &#xa0; | &#xa0;
  <a href="#checkered_flag-starting">Starting</a> &#xa0; | &#xa0;
  <a href="#memo-license">License</a> &#xa0; | &#xa0;
  <a href="https://github.com/bangadam" target="_blank">Author</a>
</p>

<br>

## :dart: About ##

Describe your project

## :sparkles: Features ##

:heavy_check_mark: API user login;\
:heavy_check_mark: API report transaction of merchant;\
:heavy_check_mark: API report transaction of outlets;

## :rocket: Technologies ##

The following tools were used in this project:

- [Go](https://go.dev) Programming language;
- [Fiber](https://docs.gofiber.io/) For Framework;
- [Swaggo](https://github.com/swaggo/swag) For API Documentation;
- [Gorm](http://gorm.io/index.html) For ORM;
- [Ozzo Validation](https://github.com/go-ozzo/ozzo-validation) For validation request;
- [Zap Logger](https://pkg.go.dev/go.uber.org/zap) For logging;
- [Viper](https://github.com/spf13/viper) For configuration;

## :white_check_mark: Requirements ##

Before starting :checkered_flag:, you need to have [Go](https://go.dev) installed.

## :checkered_flag: Starting ##

```bash
# Clone this project
$ git clone https://github.com/bangadam/backend-test-majoo

# Access
$ cd backend-test-majoo

# Migrate Database with file `database.sql`

# Set Environment Variables with your own data, check on file config.yaml
server:
  host: 127.0.0.1
  port: 8080
  secret_key: "secret"
mysql:
  host: localhost
  port: 3306
  password: 
  user: 
  database: 


# Install dependencies
$ go get ./...

# Run the project
$ go run cmd/api/main.go
# The server will initialize in the <http://localhost:${server.port}>

# Go to API Documentation in the <http://localhost:${server.port}/swagger/index.html>
```

## :memo: License ##

This project is under license from MIT. For more details, see the [LICENSE](LICENSE.md) file.


Made with :heart: by <a href="https://github.com/bangadam" target="_blank">Muhammad Adam</a>

&#xa0;

<a href="#top">Back to top</a>
