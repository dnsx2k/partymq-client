package client

type Binding struct {
	Exchange   string `json:"exchange"`
	RoutingKey string `json:"routingKey"`
}
