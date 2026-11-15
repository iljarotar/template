# Go Template Rendering Tool

This is a helper tool to render a go template with data provided in JSON format.

## Usage

Invoke with

```bash
template -t "<template-file>(,<another-template-file>,...)" -d "<data.json>" -o "<output-file>"
```

## Extra Functions

**includeHTML**

Includes raw HTML from a file path.
