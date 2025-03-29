# 🚀 GodePack

> A simple Go CLI tool to bundle your Git repository into a single, AI-friendly file.

## 🎯 Features

- Quickly bundles Git repository contents into a single output file.
- Automatically excludes irrelevant directories and file types (e.g., `.git`, `node_modules`, images, binaries).
- Easy-to-use CLI written in Go.

## ⚡ Installation

Install via `go install`:

```bash
go install github.com/instantraaamen/godepack/cmd/godepack@latest
```

## 🛠 Usage

Navigate to your repository's root and run:

```bash
godepack
```

This generates a file named `repomix-output.txt`, containing your repository content optimized for AI processing.

## 🌳 Project Structure

```
godepack/
├── cmd/
│   └── godepack/
│       └── main.go          # Entry point for the CLI application
├── internal/
│   └── bundler.go           # Core bundling logic
├── testdata/                # Test data for unit tests
├── main_test.go             # Unit tests
├── repomix-output.txt       # Generated bundled output
├── go.mod                   # Go module definition
├── go.sum                   # Go module checksum
└── README.md                # This document
```

## 🤝 Contributing

Contributions are welcome! Please submit an issue or pull request.

## 📄 License

## License

This project is planned to be released under the [MIT License](https://opensource.org/licenses/MIT) in the future.  
The LICENSE file will be updated upon the official release.


