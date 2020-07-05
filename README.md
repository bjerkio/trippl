![Tripletex Golang Client][banner]

![Build & Deploy][build-badge]
[![Total alerts][lgtm-badge]][lgtm-alerts]
[![Maintainability][codeclimate-badge]][codeclimate]

## About

**IMPORTANT**: This product is at an early stage and not yet used in production.
We do not offer support for this version, but will release supported versions
in the future. Feel free to play around, but for full functionality and support,
please wait for the supported, stable release.

**Want to help out? We have a [v1 milestone](https://github.com/bjerkio/trippl/milestone/1)! :tada:**

Trippl helps synchronize projects and time tracking between Tripletex and Toggl.
Tripletex is a Norwegian accounting software, that features both time tracking,
project management and CRM. Toggl is a time tracking software. The motivation behind
this product in most cases, Tripletex is has a good time tracking system, but integrating
it to everything is a big hurdle. Bjerk, the company behind Trippl, uses tracks projects
in Atlassian Jira, and uses a module to track hours from Jira into Toggl. Toggl has become
one of the most popular choices for time tracking, not only in Norway, but internationally.
Since we love Tripletex and Toggl, we wanted to combine it.

Our intentions are to release this under Apache 2.0 License, but also support a hosted version,
where you easily can sign up, enter your credit card, and pay a small amount every month to
cover hosting and us building this integration. This repository will feature it as
soon as it's ready.

# Documentation

Apart from this README, you can find details and examples of using the SDK in
the following places:

- [SDK Documentation][sdk-doc]

## Contributing

We love contributions! 🙏 Bug reports and pull requests are welcome on [GitHub][github].

[banner]: ./.github/header.png
[build-badge]: https://github.com/bjerkio/trippl/workflows/build/badge.svg
[lgtm-badge]: https://img.shields.io/lgtm/alerts/g/bjerkio/trippl.svg?logo=lgtm&logoWidth=18
[lgtm-alerts]: https://lgtm.com/projects/g/bjerkio/trippl/alerts/
[codeclimate-badge]: https://api.codeclimate.com/v1/badges/58be7beec5935ef4531b/maintainability
[codeclimate]: https://codeclimate.com/github/bjerkio/trippl/maintainability
[sdk-doc]: https://pkg.go.dev/mod/github.com/trippl/tripletex-go
[github]: https://bjerkio/trippl
