package build

// Version is the version the binary was built with.
var Version = "development"

// User is the GitHub actor who triggered the build.
var User string

// Time is the UTC timestamp of the build.
var Time string

// GitCommit is the git commit hash the binary was built from.
var GitCommit string

// TargetOS is the operating system the binary was built for.
var TargetOS string
