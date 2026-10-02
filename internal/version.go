// Package internal holds values shared by the propraven package and its tools.
package internal

// PackageVersion is the SDK version. The release workflow refuses a tag that
// does not match it, and it is sent in the User-Agent header.
const PackageVersion = "0.3.0"
