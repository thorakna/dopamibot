package stream

type Config struct {
	Prefix        string `json:"prefix"`
	ServiceUrl    string `json:"service_url"`
	BotToken      string `json:"bot_token"`
	OwnerId       string `json:"owner_id"`
	UseSharding   bool   `json:"use_sharding"`
	ShardId       int    `json:"shard_id"`
	ShardCount    int    `json:"shard_count"`
	DefaultStatus string `json:"default_status"`
}
