# Markdown to Safe HTML Converter

This Go project converts Markdown files into safe, sanitized HTML using the `blackfriday` and `bluemonday` libraries.

---

## 🚀 What does this program do?

It reads a `.md` (Markdown) file, converts it to HTML, and then sanitizes the HTML content to remove any potentially harmful code (such as JavaScript injections), ensuring it’s safe to display in a browser.

---

## 🧩 Technologies Used

- [`blackfriday/v2`](https://github.com/russross/blackfriday) – Parses Markdown content into raw HTML.
- [`bluemonday`](https://github.com/microcosm-cc/bluemonday) – Sanitizes the HTML output to prevent XSS and other security vulnerabilities.
- Standard Go packages: `os`, `fmt`, `io/ioutil`.

---

## 📌 Core Function: `parseContent`

```go
func parseContent(input []byte) []byte {
    unsafe := blackfriday.Run(input)
    safe := bluemonday.UGCPolicy().SanitizeBytes(unsafe)
    return safe
}