# Effective Go – Cheat Sheet

A concise reference of Go idioms and best practices from the Effective Go guide.

---

## ✅ Formatting
- Always use `gofmt` or `go fmt` to format your code.
- Use tabs (not spaces) for indentation.
- No parentheses around control structure conditions (e.g. `if`, `for`, `switch`).
- Opening braces go on the same line.

---

## 🧾 Comments
- Use `//` for single-line comments.
- Doc comments (above exported declarations) should be full sentences that start with the name being documented.
  - Example: `// Package math provides...`

---

## 📦 Package Naming
- Short, lowercase, no underscores or mixedCaps.
  - Example: `math`, `strings`

---

## ✨ Names and Exporting
- Capitalized identifiers are **exported** (public).
- Unexported names start with a lowercase letter.
- Avoid stutter: if the package is `circle`, avoid names like `circle.Circle`.

---

## 🔁 Control Structures
- `if`, `for`, `switch`, `select` do not use parentheses.
- Use `for` for loops and `for range` to iterate over collections.
- Prefer `switch` over long `if-else` chains.
- Avoid unnecessary `else` after `return`.

---

## 🧪 Functions
- Prefer multiple return values for error handling: `value, err := SomeFunc()`.
- Keep functions short and focused.
- Use named return values for clarity in short functions, but avoid "naked" returns in longer ones.

---

## 📦 Imports
- Group standard and external packages.
- Omit unused imports and variables—they are compilation errors.

---

## 🧱 Defer
- Use `defer` for cleanup operations (e.g., closing files).
- Deferred functions are executed in LIFO order when the function returns.

---

## ⚙️ Error Handling
- Return errors instead of using `panic` for normal error cases.
- Always check returned errors.
- Error messages should be lowercase and not end in punctuation.

---

## 👤 Methods and Interfaces
- Use value receivers if the method does not modify the receiver.
- Use pointer receivers to modify or avoid copying the receiver.
- Small, focused interfaces are better (e.g., `io.Reader`).
- Prefer defining interfaces where they're used, not at the implementation.

---

## 🧩 Composition Over Inheritance
- Use struct embedding to extend functionality instead of inheritance.

---

## 📌 Slices and Arrays
- `nil` slices are valid and have zero length.
- Use `append` to grow slices.
- Avoid creating slices with zero capacity unless necessary.

---

## 🗺️ Maps
- Reading a missing key returns the zero value.
- Use `value, ok := map[key]` to test for presence.

---

## 🧵 Concurrency
- Use goroutines for lightweight concurrent functions: `go func() {...}()`
- Use channels to safely communicate between goroutines.
- Don’t share memory—communicate instead.

---

## 🧪 Testing
- Use the `testing` package with functions like `TestXxx(t *testing.T)`.
- Use `t.Error`, `t.Fail`, `t.Fatal`, etc. for test results.
- Table-driven tests (using slices of input/output structs) are idiomatic.

---

## 📚 Examples
- Example functions should be named `ExampleXxx` and can be used by `go test` to run documentation examples.

---

## 💡 Misc Tips
- Avoid overdesigning with too many layers of abstraction.
- Use interfaces to make code more flexible and testable.
- Avoid exporting types or functions prematurely.

---

## 📄 Useful Tools
- `go fmt` – auto-format code
- `go vet` – static analysis
- `golint` – style suggestions
- `godoc` – documentation

---
