<a name="readme-top"></a>

<!-- PROJECT SHIELDS -->
[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/rcebrian/go-api-template">
    <img src="https://go.dev/images/go-logo-white.svg" alt="Logo">
  </a>

<h3 align="center">Go API template</h3>

  <p align="center">
    Go API template based on OpenAPI initiative
    <br />
    <a href="https://github.com/rcebrian/go-api-template"><strong>Explore the docs »</strong></a>
    <br />
    <br />
    <a href="https://github.com/rcebrian/go-api-template/issues">Report Bug</a>
    ·
    <a href="https://github.com/rcebrian/go-api-template/issues">Request Feature</a>
  </p>
</div>



<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#roadmap">Roadmap</a></li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#license">License</a></li>
    <li><a href="#acknowledgments">Acknowledgments</a></li>
    <li><a href="#notes">Notes</a></li>
  </ol>
</details>

## About The Project

[//]: # (todo: why)

[//]: # (todo: internal routing)

[//]: # (todo: openapi templating)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

* [![golang][golang]][golang-url]
* [![openapi][openapi]][openapi-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- GETTING STARTED -->

## Getting Started

This is an example of how you may give instructions on setting up your project locally.
To get a local copy up and running follow these simple example steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.

* npm
  ```sh
  npm install npm@latest -g
  ```

### Installation

1. Get a free API Key at [https://example.com](https://example.com)
2. Clone the repo
   ```sh
   git clone https://github.com/rcebrian/go-api-template.git
   ```
3. Install NPM packages
   ```sh
   npm install
   ```
4. Enter your API in `config.js`
   ```js
   const API_KEY = 'ENTER YOUR API';
   ```

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- USAGE EXAMPLES -->

## Usage

Use this space to show useful examples of how a project can be used. Additional screenshots, code examples and demos
work well in this space. You may also link to more resources.

_For more examples, please refer to the [Documentation](https://example.com)_

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ROADMAP -->

## Roadmap

- [x] Openapi generator
- [x] Graceful shutdown
- [x] Makefile based
- [x] Git hooks
- [ ] Internal endpoints
    - [x] Healthcheck
    - [ ] ReDoc
- [ ] Prometheus metrics

See the [open issues](https://github.com/rcebrian/go-api-template/issues) for a full list of proposed features (and
known issues).

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTRIBUTING -->

## Contributing

Contributions are what make the open source community such an amazing place to learn, inspire, and create. Any
contributions you make are **greatly appreciated**.

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also
simply open an issue with the tag "enhancement".
Don't forget to give the project a star! Thanks again!

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- LICENSE -->

## License

Distributed under the MIT License. See `LICENSE.txt` for more information.

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- ACKNOWLEDGMENTS -->

## Acknowledgments

* [OpenAPI Generator](https://openapi-generator.tech)
* [Choose an Open Source License](https://choosealicense.com)
* [README template](https://github.com/othneildrew/Best-README-Template)

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- NOTES -->

## Notes

* Makefile **MUST NOT** change well-defined command semantics, see Makefile for details.
* `cmd` and `build` **MUST** have the same set of subdirectories for main targets
    * For example, `cmd/admin,cmd/controller` and `build/admin,build/controller`.
    * Dockerfile **MUST** be put under `build` directory even if you have only one Dockerfile.

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/rcebrian/go-api-template.svg?style=for-the-badge

[contributors-url]: https://github.com/rcebrian/go-api-template/graphs/contributors

[forks-shield]: https://img.shields.io/github/forks/rcebrian/go-api-template.svg?style=for-the-badge

[forks-url]: https://github.com/rcebrian/go-api-template/network/members

[stars-shield]: https://img.shields.io/github/stars/rcebrian/go-api-template.svg?style=for-the-badge

[stars-url]: https://github.com/rcebrian/go-api-template/stargazers

[issues-shield]: https://img.shields.io/github/issues/rcebrian/go-api-template.svg?style=for-the-badge

[issues-url]: https://github.com/rcebrian/go-api-template/issues

[license-shield]: https://img.shields.io/github/license/rcebrian/go-api-template.svg?style=for-the-badge

[license-url]: https://github.com/rcebrian/go-api-template/blob/master/LICENSE.txt

[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555

[linkedin-url]: https://linkedin.com/in/rcebrian

[golang]: https://img.shields.io/badge/-golang-black.svg?style=for-the-badge&logo=go&colorB=007F9f

[golang-url]: https://go.dev

[openapi]: https://img.shields.io/static/v1?style=for-the-badge&message=OpenAPI&color=6BA539&logo=OpenAPI+Initiative&logoColor=FFFFFF&label=

[openapi-url]: https://www.openapis.org/


