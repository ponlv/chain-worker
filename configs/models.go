package configs

import "github.com/ponlv/go-kit/mongodb"

type Schema struct {
	mongodb.DefaultModel `json:",inline" mapstructure:",inline"`

	Chain struct {
		Url string `mapstructure:"url" bson:"url"`
		WSS string `mapstructure:"wss" bson:"wss"`
	} `mapstructure:"chain" bson:"chain"`
}

func (Schema) CollectionName() string {
	return "configs"
}
