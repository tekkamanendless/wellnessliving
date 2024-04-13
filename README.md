# `wellnessliving`

![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/tekkamanendless/wellnessliving?label=version&logo=version&sort=semver)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/tekkamanendless/wellnessliving)](https://pkg.go.dev/github.com/tekkamanendless/wellnessliving)

This is a WellnessLiving client for Go.

Its primary purpose is to provide a simple client that you can use to make authenticated WellnessLiving API requests.
(The signature-verification part is particularly annoying to do without such a client.)

At this time, individual response structures are generally not present; there are a few based on the use cases that I have had a personal need for.
You can make your own with a `struct` and `json` tags on its fields.
