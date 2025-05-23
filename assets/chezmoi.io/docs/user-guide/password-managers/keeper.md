# Keeper

chezmoi includes support for [Keeper][keeper] using the
[Commander CLI][commander] to expose data as a template function.

Create a persistent login session as
[described in the Command CLI documentation][clidocs].

Passwords can be retrieved with the `keeperFindPassword` template function, for
example:

```text
examplePasswordFromPath = {{ keeperFindPassword "$PATH" }}
examplePasswordFromUid = {{ keeperFindPassword "$UID" }}
```

For retrieving more complex data, use the `keeper` template function with a UID
to retrieve structured data from [`keeper get`][get] or the `keeperDataFields`
template function which restructures the output of `keeper get` in to a more
convenient form, for example:

```text
keeperDataTitle = {{ (keeper "$UID").data.title }}
examplePassword = {{ index (keeperDataFields "$UID").password 0 }}
```

Extra arguments can be passed to the Keeper CLI command by setting the
`keeper.args` variable in chezmoi's config file, for example:

```toml title="~/.config/chezmoi/chezmoi.toml"
[keeper]
    args = ["--config", "/path/to/config.json"]
```

[keeper]: https://www.keepersecurity.com/
[commander]: https://docs.keeper.io/secrets-manager/commander-cli
[clidocs]: https://docs.keeper.io/secrets-manager/commander-cli/using-commander/logging-in#persistent-login-sessions
[get]: https://docs.keeper.io/secrets-manager/commander-cli/using-commander/command-reference/record-commands#get-command
