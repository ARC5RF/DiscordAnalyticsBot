package perms

var Dict = map[int64]DictItem{
	CREATE_INSTANT_INVITE: {
		"CREATE_INSTANT_INVITE",
		PERM_WARN_NONE,
		0,
		"Allows creation of instant invites",
		T | V | S,
	},
	KICK_MEMBERS: {
		"KICK_MEMBERS",
		PERM_WARN_2FA,
		1,
		"Allows kicking members",
		PERM_UNTYPED,
	},
	BAN_MEMBERS: {
		"BAN_MEMBERS",
		PERM_WARN_2FA,
		2,
		"Allows banning members",
		PERM_UNTYPED,
	},
	ADMINISTRATOR: {
		"ADMINISTRATOR",
		PERM_WARN_2FA,
		3,
		"Allows all permissions and bypasses channel permission overwrites",
		PERM_UNTYPED,
	},
	MANAGE_CHANNELS: {
		"MANAGE_CHANNELS",
		PERM_WARN_2FA,
		4,
		"Allows management and editing of channels",
		T | V | S,
	},
	MANAGE_GUILD: {
		"MANAGE_GUILD",
		PERM_WARN_2FA,
		5,
		"Allows management and editing of the guild",
		PERM_UNTYPED,
	},
	ADD_REACTIONS: {
		"ADD_REACTIONS",
		PERM_WARN_NONE,
		6,
		"Allows for the addition of reactions to messages",
		T | V | S,
	},
	VIEW_AUDIT_LOG: {
		"VIEW_AUDIT_LOG",
		PERM_WARN_NONE,
		7,
		"Allows for viewing of audit logs",
		PERM_UNTYPED,
	},
	PRIORITY_SPEAKER: {
		"PRIORITY_SPEAKER",
		PERM_WARN_NONE,
		8,
		"Allows for using priority speaker in a voice channel",
		V,
	},
	STREAM: {
		"STREAM",
		PERM_WARN_NONE,
		9,
		"Allows the user to go live",
		V | S,
	},
	VIEW_CHANNEL: {
		"VIEW_CHANNEL",
		PERM_WARN_NONE,
		10,
		"Allows guild members to view a channel, which includes reading messages\n        in text channels and joining voice channels",
		T | V | S,
	},
	SEND_MESSAGES: {
		"SEND_MESSAGES",
		PERM_WARN_NONE,
		11,
		"Allows for sending messages in a channel and creating threads in a forum\n        (does not allow sending messages in threads)",
		T | V | S,
	},
	SEND_TTS_MESSAGES: {
		"SEND_TTS_MESSAGES",
		PERM_WARN_NONE,
		12,
		"Allows for sending of/ttsmessages",
		T | V | S,
	},
	MANAGE_MESSAGES: {
		"MANAGE_MESSAGES",
		PERM_WARN_2FA,
		13,
		"Allows for deletion of other users messages",
		T | V | S,
	},
	EMBED_LINKS: {
		"EMBED_LINKS",
		PERM_WARN_NONE,
		14,
		"Links sent by users with this permission will be auto-embedded",
		T | V | S,
	},
	ATTACH_FILES: {
		"ATTACH_FILES",
		PERM_WARN_NONE,
		15,
		"Allows for uploading images and files",
		T | V | S,
	},
	READ_MESSAGE_HISTORY: {
		"READ_MESSAGE_HISTORY",
		PERM_WARN_NONE,
		16,
		"Allows for reading of message history",
		T | V | S,
	},
	MENTION_EVERYONE: {
		"MENTION_EVERYONE",
		PERM_WARN_NONE,
		17,
		"Allows for using the@everyonetag to notify all users in a\n        channel, and the@heretag to notify all online users in a\n        channel",
		T | V | S,
	},
	USE_EXTERNAL_EMOJIS: {
		"USE_EXTERNAL_EMOJIS",
		PERM_WARN_NONE,
		18,
		"Allows the usage of custom emojis from other servers",
		T | V | S,
	},
	VIEW_GUILD_INSIGHTS: {
		"VIEW_GUILD_INSIGHTS",
		PERM_WARN_NONE,
		19,
		"Allows for viewing guild insights",
		PERM_UNTYPED,
	},
	CONNECT: {
		"CONNECT",
		PERM_WARN_NONE,
		20,
		"Allows for joining of a voice channel",
		V | S,
	},
	SPEAK: {
		"SPEAK",
		PERM_WARN_NONE,
		21,
		"Allows for speaking in a voice channel",
		V,
	},
	MUTE_MEMBERS: {
		"MUTE_MEMBERS",
		PERM_WARN_NONE,
		22,
		"Allows for muting members in a voice channel",
		V | S,
	},
	DEAFEN_MEMBERS: {
		"DEAFEN_MEMBERS",
		PERM_WARN_NONE,
		23,
		"Allows for deafening of members in a voice channel",
		V,
	},
	MOVE_MEMBERS: {
		"MOVE_MEMBERS",
		PERM_WARN_NONE,
		24,
		"Allows for moving of members between voice channels",
		V | S,
	},
	USE_VAD: {
		"USE_VAD",
		PERM_WARN_NONE,
		25,
		"Allows for using voice-activity-detection in a voice channel",
		V,
	},
	CHANGE_NICKNAME: {
		"CHANGE_NICKNAME",
		PERM_WARN_NONE,
		26,
		"Allows for modification of own nickname",
		PERM_UNTYPED,
	},
	MANAGE_NICKNAMES: {
		"MANAGE_NICKNAMES",
		PERM_WARN_NONE,
		27,
		"Allows for modification of other users nicknames",
		PERM_UNTYPED,
	},
	MANAGE_ROLES: {
		"MANAGE_ROLES",
		PERM_WARN_2FA,
		28,
		"Allows management and editing of roles",
		T | V | S,
	},
	MANAGE_WEBHOOKS: {
		"MANAGE_WEBHOOKS",
		PERM_WARN_2FA,
		29,
		"Allows management and editing of webhooks",
		T | V | S,
	},
	MANAGE_GUILD_EXPRESSIONS: {
		"MANAGE_GUILD_EXPRESSIONS",
		PERM_WARN_2FA,
		30,
		"Allows for editing and deleting emojis, stickers, and soundboard sounds\n        created by all users",
		PERM_UNTYPED,
	},
	USE_APPLICATION_COMMANDS: {
		"USE_APPLICATION_COMMANDS",
		PERM_WARN_NONE,
		31,
		"Allows members to use application commands, including slash commands and\n        context menu commands.",
		T | V | S,
	},
	REQUEST_TO_SPEAK: {
		"REQUEST_TO_SPEAK",
		PERM_WARN_NONE,
		32,
		"Allows for requesting to speak in stage channels. (This permission is under active development and may be changed or\n          removed.)",
		S,
	},
	MANAGE_EVENTS: {
		"MANAGE_EVENTS",
		PERM_WARN_NONE,
		33,
		"Allows for editing and deleting scheduled events created by all users",
		V | S,
	},
	MANAGE_THREADS: {
		"MANAGE_THREADS",
		PERM_WARN_2FA,
		34,
		"Allows for deleting and archiving threads, and viewing all private\n        threads",
		T,
	},
	CREATE_PUBLIC_THREADS: {
		"CREATE_PUBLIC_THREADS",
		PERM_WARN_NONE,
		35,
		"Allows for creating public and announcement threads",
		T,
	},
	CREATE_PRIVATE_THREADS: {
		"CREATE_PRIVATE_THREADS",
		PERM_WARN_NONE,
		36,
		"Allows for creating private threads",
		T,
	},
	USE_EXTERNAL_STICKERS: {
		"USE_EXTERNAL_STICKERS",
		PERM_WARN_NONE,
		37,
		"Allows the usage of custom stickers from other servers",
		T | V | S,
	},
	SEND_MESSAGES_IN_THREADS: {
		"SEND_MESSAGES_IN_THREADS",
		PERM_WARN_NONE,
		38,
		"Allows for sending messages in threads",
		T,
	},
	USE_EMBEDDED_ACTIVITIES: {
		"USE_EMBEDDED_ACTIVITIES",
		PERM_WARN_NONE,
		39,
		"Allows for using Activities (applications with theEMBEDDEDflag) in a voice channel",
		V,
	},
	MODERATE_MEMBERS: {
		"MODERATE_MEMBERS",
		PERM_WARN_TIMEOUT,
		40,
		"Allows for timing out users to prevent them from sending or reacting to\n        messages in chat and threads, and from speaking in voice and stage\n        channels",
		PERM_UNTYPED,
	},
	VIEW_CREATOR_MONETIZATION_ANALYTICS: {
		"VIEW_CREATOR_MONETIZATION_ANALYTICS",
		PERM_WARN_2FA,
		41,
		"Allows for viewing role subscription insights",
		PERM_UNTYPED,
	},
	USE_SOUNDBOARD: {
		"USE_SOUNDBOARD",
		PERM_WARN_NONE,
		42,
		"Allows for using soundboard in a voice channel",
		V,
	},
	CREATE_GUILD_EXPRESSIONS: {
		"CREATE_GUILD_EXPRESSIONS",
		PERM_WARN_NONE,
		43,
		"Allows for creating emojis, stickers, and soundboard sounds, and editing\n        and deleting those created by the current user",
		PERM_UNTYPED,
	},
	CREATE_EVENTS: {
		"CREATE_EVENTS",
		PERM_WARN_NONE,
		44,
		"Allows for creating scheduled events, and editing and deleting those\n        created by the current user",
		V | S,
	},
	USE_EXTERNAL_SOUNDS: {
		"USE_EXTERNAL_SOUNDS",
		PERM_WARN_NONE,
		45,
		"Allows the usage of custom soundboard sounds from other servers",
		V,
	},
	SEND_VOICE_MESSAGES: {
		"SEND_VOICE_MESSAGES",
		PERM_WARN_NONE,
		46,
		"Allows sending voice messages",
		T | V | S,
	},
}

const (
	CREATE_INSTANT_INVITE = 1 << iota
	KICK_MEMBERS
	BAN_MEMBERS
	ADMINISTRATOR
	MANAGE_CHANNELS
	MANAGE_GUILD
	ADD_REACTIONS
	VIEW_AUDIT_LOG
	PRIORITY_SPEAKER
	STREAM
	VIEW_CHANNEL
	SEND_MESSAGES
	SEND_TTS_MESSAGES
	MANAGE_MESSAGES
	EMBED_LINKS
	ATTACH_FILES
	READ_MESSAGE_HISTORY
	MENTION_EVERYONE
	USE_EXTERNAL_EMOJIS
	VIEW_GUILD_INSIGHTS
	CONNECT
	SPEAK
	MUTE_MEMBERS
	DEAFEN_MEMBERS
	MOVE_MEMBERS
	USE_VAD
	CHANGE_NICKNAME
	MANAGE_NICKNAMES
	MANAGE_ROLES
	MANAGE_WEBHOOKS
	MANAGE_GUILD_EXPRESSIONS
	USE_APPLICATION_COMMANDS
	REQUEST_TO_SPEAK
	MANAGE_EVENTS
	MANAGE_THREADS
	CREATE_PUBLIC_THREADS
	CREATE_PRIVATE_THREADS
	USE_EXTERNAL_STICKERS
	SEND_MESSAGES_IN_THREADS
	USE_EMBEDDED_ACTIVITIES
	MODERATE_MEMBERS
	VIEW_CREATOR_MONETIZATION_ANALYTICS
	USE_SOUNDBOARD
	CREATE_GUILD_EXPRESSIONS
	CREATE_EVENTS
	USE_EXTERNAL_SOUNDS
	SEND_VOICE_MESSAGES
)
const ALL = CREATE_INSTANT_INVITE |
	KICK_MEMBERS |
	BAN_MEMBERS |
	ADMINISTRATOR |
	MANAGE_CHANNELS |
	MANAGE_GUILD |
	ADD_REACTIONS |
	VIEW_AUDIT_LOG |
	PRIORITY_SPEAKER |
	STREAM |
	VIEW_CHANNEL |
	SEND_MESSAGES |
	SEND_TTS_MESSAGES |
	MANAGE_MESSAGES |
	EMBED_LINKS |
	ATTACH_FILES |
	READ_MESSAGE_HISTORY |
	MENTION_EVERYONE |
	USE_EXTERNAL_EMOJIS |
	VIEW_GUILD_INSIGHTS |
	CONNECT |
	SPEAK |
	MUTE_MEMBERS |
	DEAFEN_MEMBERS |
	MOVE_MEMBERS |
	USE_VAD |
	CHANGE_NICKNAME |
	MANAGE_NICKNAMES |
	MANAGE_ROLES |
	MANAGE_WEBHOOKS |
	MANAGE_GUILD_EXPRESSIONS |
	USE_APPLICATION_COMMANDS |
	REQUEST_TO_SPEAK |
	MANAGE_EVENTS |
	MANAGE_THREADS |
	CREATE_PUBLIC_THREADS |
	CREATE_PRIVATE_THREADS |
	USE_EXTERNAL_STICKERS |
	SEND_MESSAGES_IN_THREADS |
	USE_EMBEDDED_ACTIVITIES |
	MODERATE_MEMBERS |
	VIEW_CREATOR_MONETIZATION_ANALYTICS |
	USE_SOUNDBOARD |
	CREATE_GUILD_EXPRESSIONS |
	CREATE_EVENTS |
	USE_EXTERNAL_SOUNDS |
	SEND_VOICE_MESSAGES
