package perms

var ATOI = map[string]int64{
	"CREATE_INSTANT_INVITE":               CREATE_INSTANT_INVITE,
	"KICK_MEMBERS":                        KICK_MEMBERS,
	"BAN_MEMBERS":                         BAN_MEMBERS,
	"ADMINISTRATOR":                       ADMINISTRATOR,
	"MANAGE_CHANNELS":                     MANAGE_CHANNELS,
	"MANAGE_GUILD":                        MANAGE_GUILD,
	"ADD_REACTIONS":                       ADD_REACTIONS,
	"VIEW_AUDIT_LOG":                      VIEW_AUDIT_LOG,
	"PRIORITY_SPEAKER":                    PRIORITY_SPEAKER,
	"STREAM":                              STREAM,
	"VIEW_CHANNEL":                        VIEW_CHANNEL,
	"SEND_MESSAGES":                       SEND_MESSAGES,
	"SEND_TTS_MESSAGES":                   SEND_TTS_MESSAGES,
	"MANAGE_MESSAGES":                     MANAGE_MESSAGES,
	"EMBED_LINKS":                         EMBED_LINKS,
	"ATTACH_FILES":                        ATTACH_FILES,
	"READ_MESSAGE_HISTORY":                READ_MESSAGE_HISTORY,
	"MENTION_EVERYONE":                    MENTION_EVERYONE,
	"USE_EXTERNAL_EMOJIS":                 USE_EXTERNAL_EMOJIS,
	"VIEW_GUILD_INSIGHTS":                 VIEW_GUILD_INSIGHTS,
	"CONNECT":                             CONNECT,
	"SPEAK":                               SPEAK,
	"MUTE_MEMBERS":                        MUTE_MEMBERS,
	"DEAFEN_MEMBERS":                      DEAFEN_MEMBERS,
	"MOVE_MEMBERS":                        MOVE_MEMBERS,
	"USE_VAD":                             USE_VAD,
	"CHANGE_NICKNAME":                     CHANGE_NICKNAME,
	"MANAGE_NICKNAMES":                    MANAGE_NICKNAMES,
	"MANAGE_ROLES":                        MANAGE_ROLES,
	"MANAGE_WEBHOOKS":                     MANAGE_WEBHOOKS,
	"MANAGE_GUILD_EXPRESSIONS":            MANAGE_GUILD_EXPRESSIONS,
	"USE_APPLICATION_COMMANDS":            USE_APPLICATION_COMMANDS,
	"REQUEST_TO_SPEAK":                    REQUEST_TO_SPEAK,
	"MANAGE_EVENTS":                       MANAGE_EVENTS,
	"MANAGE_THREADS":                      MANAGE_THREADS,
	"CREATE_PUBLIC_THREADS":               CREATE_PUBLIC_THREADS,
	"CREATE_PRIVATE_THREADS":              CREATE_PRIVATE_THREADS,
	"USE_EXTERNAL_STICKERS":               USE_EXTERNAL_STICKERS,
	"SEND_MESSAGES_IN_THREADS":            SEND_MESSAGES_IN_THREADS,
	"USE_EMBEDDED_ACTIVITIES":             USE_EMBEDDED_ACTIVITIES,
	"MODERATE_MEMBERS":                    MODERATE_MEMBERS,
	"VIEW_CREATOR_MONETIZATION_ANALYTICS": VIEW_CREATOR_MONETIZATION_ANALYTICS,
	"USE_SOUNDBOARD":                      USE_SOUNDBOARD,
	"CREATE_GUILD_EXPRESSIONS":            CREATE_GUILD_EXPRESSIONS,
	"CREATE_EVENTS":                       CREATE_EVENTS,
	"USE_EXTERNAL_SOUNDS":                 USE_EXTERNAL_SOUNDS,
	"SEND_VOICE_MESSAGES":                 SEND_VOICE_MESSAGES,
}

var ITOA = map[int64]string{
	CREATE_INSTANT_INVITE:               "CREATE_INSTANT_INVITE",
	KICK_MEMBERS:                        "KICK_MEMBERS",
	BAN_MEMBERS:                         "BAN_MEMBERS",
	ADMINISTRATOR:                       "ADMINISTRATOR",
	MANAGE_CHANNELS:                     "MANAGE_CHANNELS",
	MANAGE_GUILD:                        "MANAGE_GUILD",
	ADD_REACTIONS:                       "ADD_REACTIONS",
	VIEW_AUDIT_LOG:                      "VIEW_AUDIT_LOG",
	PRIORITY_SPEAKER:                    "PRIORITY_SPEAKER",
	STREAM:                              "STREAM",
	VIEW_CHANNEL:                        "VIEW_CHANNEL",
	SEND_MESSAGES:                       "SEND_MESSAGES",
	SEND_TTS_MESSAGES:                   "SEND_TTS_MESSAGES",
	MANAGE_MESSAGES:                     "MANAGE_MESSAGES",
	EMBED_LINKS:                         "EMBED_LINKS",
	ATTACH_FILES:                        "ATTACH_FILES",
	READ_MESSAGE_HISTORY:                "READ_MESSAGE_HISTORY",
	MENTION_EVERYONE:                    "MENTION_EVERYONE",
	USE_EXTERNAL_EMOJIS:                 "USE_EXTERNAL_EMOJIS",
	VIEW_GUILD_INSIGHTS:                 "VIEW_GUILD_INSIGHTS",
	CONNECT:                             "CONNECT",
	SPEAK:                               "SPEAK",
	MUTE_MEMBERS:                        "MUTE_MEMBERS",
	DEAFEN_MEMBERS:                      "DEAFEN_MEMBERS",
	MOVE_MEMBERS:                        "MOVE_MEMBERS",
	USE_VAD:                             "USE_VAD",
	CHANGE_NICKNAME:                     "CHANGE_NICKNAME",
	MANAGE_NICKNAMES:                    "MANAGE_NICKNAMES",
	MANAGE_ROLES:                        "MANAGE_ROLES",
	MANAGE_WEBHOOKS:                     "MANAGE_WEBHOOKS",
	MANAGE_GUILD_EXPRESSIONS:            "MANAGE_GUILD_EXPRESSIONS",
	USE_APPLICATION_COMMANDS:            "USE_APPLICATION_COMMANDS",
	REQUEST_TO_SPEAK:                    "REQUEST_TO_SPEAK",
	MANAGE_EVENTS:                       "MANAGE_EVENTS",
	MANAGE_THREADS:                      "MANAGE_THREADS",
	CREATE_PUBLIC_THREADS:               "CREATE_PUBLIC_THREADS",
	CREATE_PRIVATE_THREADS:              "CREATE_PRIVATE_THREADS",
	USE_EXTERNAL_STICKERS:               "USE_EXTERNAL_STICKERS",
	SEND_MESSAGES_IN_THREADS:            "SEND_MESSAGES_IN_THREADS",
	USE_EMBEDDED_ACTIVITIES:             "USE_EMBEDDED_ACTIVITIES",
	MODERATE_MEMBERS:                    "MODERATE_MEMBERS",
	VIEW_CREATOR_MONETIZATION_ANALYTICS: "VIEW_CREATOR_MONETIZATION_ANALYTICS",
	USE_SOUNDBOARD:                      "USE_SOUNDBOARD",
	CREATE_GUILD_EXPRESSIONS:            "CREATE_GUILD_EXPRESSIONS",
	CREATE_EVENTS:                       "CREATE_EVENTS",
	USE_EXTERNAL_SOUNDS:                 "USE_EXTERNAL_SOUNDS",
	SEND_VOICE_MESSAGES:                 "SEND_VOICE_MESSAGES",
}

const NONE = 0

const (
	T int64 = 1 << iota
	V
	S
)

const (
	PERM_WARN_NONE = iota
	PERM_WARN_2FA
	PERM_WARN_TIMEOUT
)

const PERM_UNTYPED = 0

var PERM_WARN_MESSAGES = map[int]string{
	PERM_WARN_NONE:    "",
	PERM_WARN_2FA:     "These permissions require the owner account to use two-factor authentication when used on a guild that has server-wide 2FA enabled.",
	PERM_WARN_TIMEOUT: "See Permissions for Timed Out Members to understand how permissions are temporarily modified for timed out users.",
}

type DictItem struct {
	Human        string
	Warn         int
	Bit          int
	Description  string
	ChannelTypes int64
}
