<div id="top"></div>

<!-- PROJECT SHIELDS -->

[<div align="center"> ![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url] [![Stargazers][stars-shield]][stars-url]
[![MIT License][license-shield]][license-url]
[![Issues][issues-shield]][issues-url]
[![Issues Closed][issues-closed-shield]</div>][issues-closed-url]

<!-- ![Visitors](https://estruyf-github.azurewebsites.net/api/VisitorHit?user=wst24365888&repo=ez4o/go-try&countColor=rgb(0,%20126,%20198)) -->

<br />

![go-try](https://socialify.git.ci/ez4o/go-try/image?description=1&font=KoHo&name=1&owner=1&pattern=Circuit%20Board&theme=Light)

<br />
<div align="center">
<p align="center">
    <a href="https://github.com/ez4o/go-try#usage"><strong>Explore Usage »</strong></a>
    <br />
    <br />
    <a href="https://github.com/ez4o/go-try/issues">Report Bug</a>
    ·
    <a href="https://github.com/ez4o/go-try/issues">Request Feature</a>
  </p>
</div>

<details>
  <summary>Table of Contents</summary>
  <ol>
    <li><a href="#about">About</a></li>
    <li><a href="#why-go-try">Why go-try?</a></li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#usage">Usage</a></li>
        <li><a href="#example">Example</a></li>
      </ul>
    </li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

## About

`go-try` is a package that allows you to use `try/catch` block in Go.

This project is still in its early stages, so any thoughts and feedback are very
welcome!

<p align="right">(<a href="#top">back to top</a>)</p>

## Why `go-try`?

Many Go developers get tired of dealing with errors because there are too many
errors to handle one by one, which is intuitive and effective, but really
annoying.

I've been trying to find out if Go has an error handling method like
`try/catch`, I think that would probably be a lot easier, but unfortunately, I
couldn't find any package that's easy to use.

So I tried to make one myself, taking inspiration from the `try/catch` syntax,
then `go-try` was born!

<p align="right">(<a href="#top">back to top</a>)</p>

## Getting Started

### Installation

`go get github.com/ez4o/go-try`

### Usage

```go
Try(func () {
  ...
  ThrowOnError(ce)
  ...
  ThrowOnError(err)
  ...
}).Catch(func (ce CustomError) {
  ...
}).Catch(func (e error, st *StackTrace) {
  ...
  st.Print()
})
```

### Functions

| Name             | Description                                                                                                                                                                                                |
| ---------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Try()`          | Takes `func ()`, wrap your code here!                                                                                                                                                                      |
| `Catch()`        | Takes `func (any)` or `func (any, *StackTrace)`, and it will only accept the error type you have declared. You can accept second parameter, which is the stack trace begin from the last `ThrowOnError()`. |
| `ThrowOnError()` | Takes `any`. **Will only throw an error when the parameter is not** `nil`.                                                                                                                                 |
| `st.Print()`     | If you have declared the second parameter `st *StackTrace`, you can print the stack trace using `st.Print()`                                                                                               |

### Example

Let's say you want to fetch JSON from a url and unmarshal it, you can simply write it like this:

```go
import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "net/http"
    
  . "github.com/ez4o/go-try"
)

func main() {
  Try(func() {
    resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
    ThrowOnError(err)
    defer resp.Body.Close()
    
    b, err := ioutil.ReadAll(resp.Body)
    ThrowOnError(err)

    var data []map[string]interface{}
    err = json.Unmarshal(b, &data)
    ThrowOnError(err)

    fmt.Println(data)
  }).Catch(func(e error, st *StackTrace) {
    fmt.Println(e)
    st.Print()
  })
}
```

<p align="right">(<a href="#top">back to top</a>)</p>

## Roadmap

- [x] Implement catching errors of different types.
- [ ] Tests
- [ ] CI

See the [open issues](https://github.com/ez4o/go-try/issues) for a full list of
proposed features (and known issues).

<p align="right">(<a href="#top">back to top</a>)</p>

## Contributing

Contributions are what make the open source community such an amazing place to
learn, inspire, and create. Any contributions you make are **greatly
appreciated**.

If you have a suggestion that would make this better, please fork the repo and
create a pull request. You can also simply open an issue with the tag
"enhancement". Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feat/amazing-feature`)
3. Commit your Changes with
   [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/)
4. Push to the Branch (`git push origin feat/amazing-feature`)
5. Open a Pull Request

<p align="right">(<a href="#top">back to top</a>)</p>

## License

Distributed under the MIT License. See
[LICENSE](https://github.com/ez4o/go-try/blob/main/LICENSE) for more
information.

<p align="right">(<a href="#top">back to top</a>)</p>

## Contact

### Author

- HSING-HAN, WU (Xyphuz)
  - Mail me: xyphuzwu@gmail.com
  - About me: <https://about.xyphuz.com>
  - GitHub: <https://github.com/wst24365888>

### Project Link

- <https://github.com/ez4o/go-try>

<p align="right">(<a href="#top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/ez4o/go-try.svg?style=for-the-badge
[contributors-url]: https://github.com/ez4o/go-try/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/ez4o/go-try.svg?style=for-the-badge
[forks-url]: https://github.com/ez4o/go-try/network/members
[stars-shield]: https://img.shields.io/github/stars/ez4o/go-try.svg?style=for-the-badge
[stars-url]: https://github.com/ez4o/go-try/stargazers
[issues-shield]: https://img.shields.io/github/issues/ez4o/go-try.svg?style=for-the-badge
[issues-url]: https://github.com/ez4o/go-try/issues
[issues-closed-shield]: https://img.shields.io/github/issues-closed/ez4o/go-try.svg?style=for-the-badge
[issues-closed-url]: https://github.com/ez4o/go-try/issues?q=is%3Aissue+is%3Aclosed
[license-shield]: https://img.shields.io/github/license/ez4o/go-try.svg?style=for-the-badge
[license-url]: https://github.com/ez4o/go-try/blob/main/LICENSE
[product-screenshot]: https://go-try.ez4o.com/?username=wst24365888&img_url=https%3A%2F%2Fimages.unsplash.com%2Fphoto-1506744038136-46273834b3fb%3Fixid%3DMnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8%26ixlib%3Drb-1.2.1%26auto%3Dformat%26fit%3Dcrop%26w%3D1000%26q%3D80&fbclid=IwAR1AUDKHzjzBSjKle6J44dYRSrIbvBu8eTxtrfhpPxhBnBsOizgSq63bYbU
