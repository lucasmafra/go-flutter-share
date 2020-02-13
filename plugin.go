package share

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	"github.com/go-flutter-desktop/go-flutter/plugin"
)

const channelName = "flutter_share"

// SharePlugin implements flutter.Plugin and handles method.
type SharePlugin struct{}

var _ flutter.Plugin = &SharePlugin{} // compile-time type check

// InitPlugin initializes the plugin.
func (p *SharePlugin) InitPlugin(messenger plugin.BinaryMessenger) error {
	channel := plugin.NewMethodChannel(messenger, channelName, plugin.StandardMethodCodec{})
	channel.HandleFunc("share", p.share)
	return nil
}

func (p *SharePlugin) share(arguments interface{}) (reply interface{}, err error) {		
	return nil, nil
}
