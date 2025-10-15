## Contributing to go.validator

Thank you for taking the time to contribute! This guide explains how to set up your environment, make changes, run tests, and submit pull requests.

### Prerequisites
- Go 1.18 or newer (see `go.mod`)
- Make (for convenient tasks)
- `golangci-lint` 2.0 or higher installed and available in PATH

### Getting started
1. Fork the repository to your GitHub account.
2. Clone your fork:
   ```bash
   git clone https://github.com/<your-username>/go.validator.git
   cd go.validator
   git remote add upstream https://github.com/behzadsh/go.validator.git
   ```
3. Create a feature branch:
   ```bash
   git checkout -b feature/my-change
   ```

### Development workflow
- Keep your branch up to date:
  ```bash
  git fetch upstream
  git rebase upstream/main
  ```
- Run tests frequently:
  ```bash
  make test
  ```
- Check coverage (and generate an HTML report):
  ```bash
  make coverage
  ```
- Lint locally:
  ```bash
  make lint
  ```
- Clean generated artifacts:
  ```bash
  make clean
  ```

### Coding guidelines
- Write clear, readable, and well-structured Go code.
- Prefer meaningful names; avoid unnecessary abbreviations.
- Keep control flow simple; favor early returns.
- Add comments only where they add non-obvious context or rationale.
- Ensure the code adheres to the repository's `golangci-lint` rules (run `make lint`).
- Keep dependencies minimal; prefer the standard library where possible.

### Tests
- Add or update unit tests for your changes.
- Place tests alongside the code (e.g., `rules/some_rule.go` and `rules/some_rule_test.go`).
- Ensure `make test` passes on your machine.

### Adding a new rule
When implementing a new validation rule, follow these steps:
- Implement the `Rule` interface in `rules/`.
- If your rule accepts parameters (e.g., `between:3,5`), implement `RuleWithParams` (`AddParams([]string)` and `MinRequiredParams() int`).
- For localized messages, embed `translation.BaseTranslatableRule` and use its `Translate` function.
- Register the rule in the core registry by adding it to `registerDefaultRules()` inside `init.go` (file: `init.go`, function: `registerDefaultRules()`). Use the established naming convention, e.g.: `"yourRuleName": &rules.YourRule{}`.
- Add thorough unit tests in `rules/your_rule_test.go`.
- Document the rule briefly in `rules.md` (name, semantics, examples) and don't forget to add it to the index.

### Translations and i18n
- By default, validation messages use translation keys like `validation.required`.
- You can provide or test custom translations using `translation.SetDefaultTranslatorFunc`.
- Ensure your rule’s messages use placeholders like `:field:`, and when using `BaseTranslatableRule`, call `Translate(r.Locale, "validation.yourKey", params)`.

### Commit messages and branches
- Use small, focused commits with clear messages. Conventional style is appreciated:
  - `feat(rules): add minWords rule`
  - `fix(datetime): handle leap years`
  - `docs: update README with translation examples`
  - `test: increase coverage for ip rule`
- Branch naming suggestions: `feat/…`, `fix/…`, `docs/…`, `chore/…`.

### Pull requests
Before opening a PR, please ensure:
- [ ] The code builds and all tests pass (`make test`).
- [ ] Linting passes locally (`make lint`).
- [ ] Coverage is not significantly reduced (`make coverage`).
- [ ] New rules are documented in `rules.md` and thoroughly tested.
- [ ] Public APIs are documented and backward compatibility considered.
- [ ] The PR description clearly explains the change and rationale.

### Reporting issues
When filing an issue, include:
- Expected vs. actual behavior
- Steps to reproduce (minimal code sample if possible)
- Go version and OS
- Any relevant logs or error messages

### Release process
Releases are managed by the maintainer(s). If your change warrants a release note, mention it in the PR description.

---
Thank you for contributing to `go.validator`!


