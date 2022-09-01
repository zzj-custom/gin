package kafka

import "time"

type Config struct {
	Host         string `toml:"host"`
	Port         int    `toml:"port"`
	Group        string `toml:"group"`
	Topic        string `toml:"topic"`
	Version      string `toml:"version"`
	ClientID     string `toml:"client_id"`
	MetadataFull bool   `toml:"metadata_full"`
	Assignor     string `toml:"assignor"`

	ProducerRequiredAcks    int  `toml:"producer_required_acks"`
	ProducerReturnSuccesses bool `toml:"producer_return_successes"`
	ProducerReturnErrors    bool `toml:"producer_return_errors"`

	ConsumerOffsetsAutoCommitEnable   bool          `toml:"consumer_offsets_auto_commit_enable"`
	ConsumerReturnErrors              bool          `toml:"return_errors"`
	ConsumerOffsetsAutoCommitInterval time.Duration `toml:"consumer_offsets_auto_commit_interval"`
	ConsumerOffsetsInitial            string        `toml:"consumer_offsets_initial"`
	ConsumerGroupRebalanceStrategy    string        `toml:"consumer_group_rebalance_strategy"`
}
