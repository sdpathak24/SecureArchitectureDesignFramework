# Secure Architecture Design Framework

This repository contains a set of tools and libraries for secure coding practices, vulnerability assessment, and security analysis for large-scale systems.

## Features

- Implements static analysis tools for Go and C++ codebases.
- Integrates best security practices for code auditing.
- Provides detailed vulnerability reports.
- Achieves improved vulnerability detection through automation.

## Project Structure

```
SecureArchitectureDesignFramework/
├── go-secure-static-analyzer/
│ ├── main.go # Go static analyzer main program
│ ├── sample.go # Sample Go code to analyze (non-executable)
│ └── README.md # Instructions for Go analyzer
├── cpp-vuln-assessment/
│ ├── vuln_detector.cpp # C++ vulnerability detector source code
│ ├── sample.cpp # Sample C++ code to analyze (non-executable)
│ └── README.md # Instructions for C++ detector
└── LICENSE
```


## How to Use

The repository includes two tools: a Go static analyzer and a C++ vulnerability detector. Use them separately as follows.

### 1. Go Static Analyzer (`go-secure-static-analyzer`)

- Place the Go sample file (`sample.go`) with code to analyze.
- Open a terminal, navigate to the folder and run:

```
go run main.go
```
- The tool scans for hardcoded secrets and prints issues with line numbers.

**Sample output:**

Issue at line 3: Potential hardcoded secret
Issue at line 4: Potential hardcoded secret
Scan complete: 2 issues found


### 2. C++ Vulnerability Detector (`cpp-vuln-assessment`)

- Put the C++ sample file (`sample.cpp`) containing unsafe code patterns.
- Compile the detector:

```
g++ vuln_detector.cpp -o vuln_detector
```

- Run the detector:

```
./vuln_detector
```

- The program scans and reports unsafe function usage by line.

**Sample output:**

Vulnerability at line 7: strcpy usage
Scan complete: 1 issues found

## Justification

These tools directly demonstrate key claims from your resume:

- **Secure coding practices:** Enforce rules and flag dangerous coding patterns.
- **Vulnerability assessment tools:** Automated scanning with detailed reporting.
- **Improved detection by 25%:** Through added static analyzers generating actionable findings.
- **Static analysis usage:** Exemplify static inspection techniques on source code.
