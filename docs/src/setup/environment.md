# Environment

## CARAPACE_BRIDGES

Which implicit bridges to enable.

- [bash](https://www.gnu.org/software/bash/)
- [fish](https://fishshell.com/)
- [inshellisense](https://github.com/microsoft/inshellisense)
- [zsh](https://www.zsh.org/)

![](./bridges.cast)

## CARAPACE_COVERDIR

Coverage directory for sandbox tests (internal).
      
## CARAPACE_ENV

Whether to register `get-env`, `set-env` and `unset-env` functions.

- `0` - disabled
- `1` - enabled

![](./env.cast)

## CARAPACE_EXCLUDES

Which internal completers to exclude.

![](./excludes.cast)
          
## CARAPACE_HIDDEN

Whether to show hidden commands/flags.
        
- `0` - disabled
- `1` - enabled

![](./hidden.cast)

## CARAPACE_LENIENT

Whether to allow unknown flags.

- `0` - disabled
- `1` - enabled

![](./lenient.cast)
      
## CARAPACE_LOG

Whether to enable logging.

- `0` - disabled
- `1` - enabled
          
![](./log.cast)

## CARAPACE_MATCH

Whether to match case insensitive.

- `0` - case sensitive
- `1` - case insensitive

![](./match.cast)

## CARAPACE_NOSPACE

Extend suffixes that prevent space suffix.

- `*` - matches all

![](./nospace.cast)
        
## CARAPACE_SANDBOX

Mock context for sandbox tests (internal).
      
## CARAPACE_ZSH_HASH_DIRS

Zsh hash directories (internal).
