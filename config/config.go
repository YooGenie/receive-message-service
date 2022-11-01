package config


var Config = struct {
	ServiceName string
	HttpPort    string
	Environment string
	Database    struct {
		Driver           string
		User             string
		Connection       string
		ConnectionString string
	}
	Log struct {
		ShowSql     bool
		ShowHttpLog bool
		Path        string
		MaxSize     int
		MaxBackups  int
		MaxAge      int
		Compress    bool
		Topic       struct {
			Behaviorlog string
			SendLog     bool
		}
		Database struct {
			Driver string
			User   string
			Schema string
			Host   string
		}
	}
	Kakao struct {
		RestApiKey        string
		RedirectURL       string
		LogoutRedirectURL string
	}
	Domain struct {
		WebApp struct {
			LoginUrl string
			HomeUrl  string
			OrgUrl   string
		}
		Admin string
	}
	AwsS3 struct {
		SecretAccessKey string
		Bucket          string
		Region          string
		AccessKeyId     string
		HttpEndPoint    string
	}
	Service struct {
		Name string
	}
	KakaoBizmessage struct {
		Account struct {
			Id       string
			Password string
		}
		AuthUrl                   string
		SendingUrl                string
		KakaoChannelSenderKey     string
		MessageSendingDomain      string
		MessageSendingAdminDomain string
	}
	Bizmessage struct {
		Account struct {
			Id       string
			Password string
		}
		AuthUrl    string
		SendingUrl string
	}
	JwtSecret string
}{}
