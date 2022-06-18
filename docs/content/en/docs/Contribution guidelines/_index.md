---
title: "Contribution Guidelines"
linkTitle: "Contribution Guidelines"
weight: 10
description: >
  Pleaase read this before opening a PR!
---

## Contributing to Magento CLI

We love pull requests!

Fork, then clone the repo:

```bash
    git clone git@github.com:your-username/magento-cli.git
```
Set up your machine:

```bash   
    cd magento-cli
    make build
```

Make sure the existing tests pass:

```bash
    make test
```

While creating new commands, don't forget to create a `_test.go` file.

Make your changes, Add tests for your change, Make the tests pass:

```bash
    make test
```

Make your change. Add tests for your change. Make the tests pass:

```bash
    make test
```

Push to your fork and [submit a pull request][pr].

[pr]: https://github.com/superterran/magento-cli/compare/

At this point you're waiting on us. We like to at least comment on pull requests
within three business days (and, typically, one business day). We may suggest
some changes or improvements or alternatives.

Some things that will increase the chance that your pull request is accepted:

* Write tests.
* Write a [good commit message][commit].

[commit]: http://tbaggery.com/2008/04/19/a-note-about-git-commit-messages.html



## Contributing to Docs

We use [Hugo](https://gohugo.io/) to format and generate our website, the
[Docsy](https://github.com/google/docsy) theme for styling and site structure, 
and [Github Actions](https://www.github.com/) to manage the deployment of the site. 
Hugo is an open-source static site generator that provides us with templates, 
content organisation in a standard directory structure, and a website generation 
engine. You write the pages in Markdown (or HTML if you want), and Hugo wraps them up into a website.

All submissions, including submissions by project members, require review. We
use GitHub pull requests for this purpose. Consult
[GitHub Help](https://help.github.com/articles/about-pull-requests/) for more
information on using pull requests.

### Updating a single page

If you've just spotted something you'd like to change while using the docs, Docsy has a shortcut for you:

1. Click **Edit this page** in the top right hand corner of the page.
1. If you don't already have an up to date fork of the project repo, you are prompted to get one - click **Fork this repository and propose changes** or **Update your Fork** to get an up to date version of the project to edit. The appropriate page in your fork is displayed in edit mode.
1. Follow the rest of the [Quick start with Netlify](#quick-start-with-netlify) process above to make, preview, and propose your changes.

### Creating an issue

If you've found a problem in the docs, but you're not sure how to fix it yourself, please create an issue in the [Goldydocs repo](https://github.com/google/docsy-example/issues). You can also create an issue about a specific page by clicking the **Create Issue** button in the top right hand corner of the page.

### Useful resources

* [Docsy user guide](https://www.docsy.dev/docs/): All about Docsy, including how it manages navigation, look and feel, and multi-language support.
* [Hugo documentation](https://gohugo.io/documentation/): Comprehensive reference for Hugo.
* [Github Hello World!](https://guides.github.com/activities/hello-world/): A basic introduction to GitHub concepts and workflow.


