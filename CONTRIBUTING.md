# Contributing

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

You can use cobra to create a new feature, but with unit testing it may be easier to dupliate the files already in use:

```bash
    cobra add <command>
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
