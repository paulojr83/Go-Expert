## Dependency injection


### Fx is a dependency injection system for Go.

### Benefits

* Eliminate globals: Fx helps you remove global state from your application. No more init() or global variables. Use Fx-managed singletons.
* Code reuse: Fx lets teams within your organization build loosely-coupled and well-integrated shareable components.
* Battle tested: Fx is the backbone of nearly all Go services at Uber.

See our [docs](https://uber-go.github.io/fx/get-started/) to get started and/or learn more about Fx.

 
### Wire: Automated Initialization in Go

Wire is a code generation tool that automates connecting components using dependency injection. Dependencies between components are represented in Wire as function parameters, encouraging explicit initialization instead of global variables. Because Wire operates without runtime state or reflection, code written to be used with Wire is useful even for hand-written initialization.

For an overview, see the [introductory blog post](https://go.dev/blog/wire).

#### Documentation
* [Tutorial](https://github.com/google/wire/blob/main/_tutorial/README.md)
* [User Guide](https://github.com/google/wire/blob/main/docs/guide.md)
* [Best Practices](https://github.com/google/wire/blob/main/docs/best-practices.md)
* [FAQ](https://github.com/google/wire/blob/main/docs/faq.md)