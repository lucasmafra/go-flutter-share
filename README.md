# open_file

This Go package has a noop host-side implementation of the Flutter [share](https://github.com/flutter/plugins/tree/master/packages/share) plugin.

## Usage

Import as:

```go
import share "github.com/lucasmafra/go-flutter-share"
```

Then add the following option to your go-flutter [application options](https://github.com/go-flutter-desktop/go-flutter/wiki/Plugin-info):

```go
flutter.AddPlugin(&share.SharePlugin{}),
```
