package space

//Project struct
type Project struct {
	ID  string `json:"id"`
	Key struct {
		Key string `json:"key"`
	} `json:"key"`
	Name                     string      `json:"name"`
	Private                  bool        `json:"private"`
	Description              string      `json:"description"`
	Icon                     interface{} `json:"icon"`
	LatestRepositoryActivity interface{} `json:"latestRepositoryActivity"`
	CreatedAt                struct {
		Iso       string `json:"iso"`
		Timestamp int64  `json:"timestamp"`
	} `json:"createdAt"`
	Archived bool `json:"archived"`
}

//ProjectRepos struct
type ProjectRepos struct {
	Repos []struct {
		ID                        string      `json:"id"`
		Name                      string      `json:"name"`
		Description               string      `json:"description"`
		LatestActivity            interface{} `json:"latestActivity"`
		ProxyPushNotification     interface{} `json:"proxyPushNotification"`
		ProxyPushNotificationBody interface{} `json:"proxyPushNotificationBody"`
		State                     string      `json:"state"`
		InitProgress              interface{} `json:"initProgress"`
		ReadmeName                interface{} `json:"readmeName"`
		MonthlyActivity           interface{} `json:"monthlyActivity"`
		DefaultBranch             struct {
			Head string `json:"head"`
			Ref  string `json:"ref"`
		} `json:"defaultBranch"`
	} `json:"repos"`
}

//Repository struct
type Repository struct {
	ID                        string      `json:"id"`
	Name                      string      `json:"name"`
	Description               string      `json:"description"`
	LatestActivity            interface{} `json:"latestActivity"`
	ProxyPushNotification     interface{} `json:"proxyPushNotification"`
	ProxyPushNotificationBody interface{} `json:"proxyPushNotificationBody"`
	State                     string      `json:"state"`
	InitProgress              interface{} `json:"initProgress"`
	ReadmeName                interface{} `json:"readmeName"`
	MonthlyActivity           interface{} `json:"monthlyActivity"`
	DefaultBranch             struct {
		Head string `json:"head"`
		Ref  string `json:"ref"`
	} `json:"defaultBranch"`
}

//CreateRepositoryData struct
type CreateRepositoryData struct {
	Description   string `json:"description"`
	DefaultBranch string `json:"defaultBranch"`
	Initialize    bool   `json:"initialize"`
	DefaultSetup  bool   `json:"defaultSetup"`
}
