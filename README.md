# LATWD

It's links.

**L**inks **A**ll **T**he **W**ay **D**own.

## Installation

**1. Using `go get` (recommended):**

```bash
go get -u github.com/tealwp/latwd
```

**2. Using `curl`:**

```bash
curl -L -O https://github.com/tealwp/latwd/archive/refs/heads/main.zip
tar -zxvf
go install <path_to_downloaded_files>
```

**3. Using `git`:**

```bash
git clone github.come/tealwp/latwd
cd latwd
go install .
```

## Usage

latwd is used via cli:

```bash
latwd <url> <max_depth> <max_breadth>
```

URL: [required] [string] - the base url to begin at
MAX_DEPTH: [not required] [integer] - the max depth to recursively call down to. Default is 5.
MAX_BREADTH: [not required] [integer] - the max breadth to recursively call from a single page. Default is 5.
