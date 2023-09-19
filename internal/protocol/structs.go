// generated via genstructs.py - All structure padding and member alignment verified
package protocol

const (
	SUCC                                                  = 0x1
	FAIL                                                  = 0x0
	SIZEOF_BYTE                                           = 0x1
	SIZEOF_DWORD                                          = 0x4
	SIZEOF_INT                                            = 0x4
	SIZEOF_FLOAT                                          = 0x4
	SIZEOF_SHORT                                          = 0x2
	SIZEOF_ULONG                                          = 0x4
	SIZEOF_UINT64                                         = 0x8
	SIZEOF_IP_STRING                                      = 0x10
	SIZEOF_CN_UID_STRING                                  = 0x32
	SIZEOF_ACCOUNT_STRING                                 = 0x21
	SIZEOF_PASSWORD_STRING                                = 0x21
	SIZEOF_AUTH_ID_STRING                                 = 0xff
	CN_MAX_COUNT_GROUP_MEMBER                             = 0x5
	CN_MAX_COUNT_PC_GROUP_MEMBER                          = 0x4
	CN_MAX_COUNT_NPC_GROUP_MEMBER                         = 0x5
	CHAT_MAX_STRING                                       = 0x80
	PC_START_LOCATION_RANDOM_RANGE                        = 0x2710
	SIZEOF_ANNOUNCE_STRING                                = 0x200
	SERVER_COUNT_SHARD_CLIENT                             = 0x19
	EXIT_CODE_DISCONNECT                                  = 0x0
	EXIT_CODE_REQ_BY_PC                                   = 0x1
	EXIT_CODE_REQ_BY_SVR                                  = 0x2
	EXIT_CODE_REQ_BY_GM                                   = 0x3
	EXIT_CODE_HACK                                        = 0x4
	EXIT_CODE_ERROR                                       = 0x5
	EXIT_CODE_LIVE_CHECK                                  = 0x6
	EXIT_CODE_REQ_BY_PC_DUPE_LOGIN                        = 0x7
	EXIT_CODE_SERVER_ERROR                                = 0x63
	SIZEOF_USER_ID                                        = 0x20
	SIZEOF_USER_PW                                        = 0x20
	SIZEOF_PC_SLOT                                        = 0x4
	SIZEOF_PC_NAME                                        = 0x10
	SIZEOF_PC_FIRST_NAME                                  = 0x9
	SIZEOF_PC_LAST_NAME                                   = 0x11
	SIZEOF_PC_NAME_FLAG                                   = 0x8
	GENDER_NONE                                           = 0x0
	GENDER_MALE                                           = 0x1
	GENDER_FEMALE                                         = 0x2
	MENTOR_CHANGE_BASE_COST                               = 0x64
	REPEAT_MISSION_RESET_TIME                             = 0x9
	SIZEOF_REPEAT_QUESTFLAG_NUMBER                        = 0x8
	FATIGUE_RESET_TIME                                    = 0x0
	PC_FATIGUE_KILL_UNIT                                  = 0x7
	PC_FATIGUE_1_LEVEL                                    = 0x2c9c
	PC_FATIGUE_2_LEVEL                                    = 0x1950
	PC_FATIGUE_MAX_LEVEL                                  = 0x2
	PC_FUSIONMATTER_MAX                                   = 0x3b9ac9ff
	PC_CANDY_MAX                                          = 0x3b9ac9ff
	PC_BATTERY_MAX                                        = 0x270f
	PC_LEVEL_MAX                                          = 0x24
	SIZEOF_PC_BULLET_SLOT                                 = 0x3
	PC_TICK_TIME                                          = 0x1388
	SIZEOF_EQUIP_SLOT                                     = 0x9
	EQUIP_SLOT_HAND                                       = 0x0
	EQUIP_SLOT_UPPERBODY                                  = 0x1
	EQUIP_SLOT_LOWERBODY                                  = 0x2
	EQUIP_SLOT_FOOT                                       = 0x3
	EQUIP_SLOT_HEAD                                       = 0x4
	EQUIP_SLOT_FACE                                       = 0x5
	EQUIP_SLOT_BACK                                       = 0x6
	EQUIP_SLOT_END                                        = 0x6
	EQUIP_SLOT_HAND_EX                                    = 0x7
	EQUIP_SLOT_VEHICLE                                    = 0x8
	WPN_EQUIP_TYPE_NONE                                   = 0x0
	WPN_EQUIP_TYPE_OH_BLADE                               = 0x1
	WPN_EQUIP_TYPE_OH_CLUB                                = 0x2
	WPN_EQUIP_TYPE_OH_PISTOL                              = 0x3
	WPN_EQUIP_TYPE_OH_RIPLE                               = 0x4
	WPN_EQUIP_TYPE_OH_THROW                               = 0x5
	WPN_EQUIP_TYPE_DH_BLADE                               = 0x6
	WPN_EQUIP_TYPE_DH_CLUB                                = 0x7
	WPN_EQUIP_TYPE_DH_DPISTOL                             = 0x8
	WPN_EQUIP_TYPE_DH_RIPLE                               = 0x9
	WPN_EQUIP_TYPE_DH_THROW                               = 0xa
	WPN_EQUIP_TYPE_DH_ROCKET                              = 0xb
	SIZEOF_INVEN_SLOT                                     = 0x32
	SIZEOF_QINVEN_SLOT                                    = 0x32
	SIZEOF_BANK_SLOT                                      = 0x77
	SIZEOF_RESTORE_SLOT                                   = 0x5
	SIZEOF_NANO_BANK_SLOT                                 = 0x25
	SIZEOF_QUEST_SLOT                                     = 0x400
	NANO_QUEST_INDEX                                      = 0x0
	SIZEOF_RQUEST_SLOT                                    = 0x9
	SIZEOF_QUESTFLAG_NUMBER                               = 0x20
	SIZEOF_EP_RECORD_SLOT                                 = 0x33
	SIZEOF_TRADE_SLOT                                     = 0xc
	SIZEOF_VENDOR_TABLE_SLOT                              = 0x14
	SIZEOF_VENDOR_RESTORE_SLOT                            = 0x5
	SIZEOF_QUEST_NPC_SLOT                                 = 0x3
	SIZEOF_QUEST_ITEM_SLOT                                = 0x3
	SIZEOF_MAX_ITEM_STACK                                 = 0x64
	SIZEOF_PC_SKILL_SLOT                                  = 0x21
	SIZEOF_QUICK_SLOT                                     = 0x8
	ENCHANT_WEAPON_MATERIAL_ID                            = 0x65
	ENCHANT_DEFENCE_MATERIAL_ID                           = 0x66
	SIZEOF_NANO_CARRY_SLOT                                = 0x3
	COUNTOF_NANO_PER_SET                                  = 0x3
	SIZEOF_NANO_SET                                       = 0xd
	SIZEOF_NANO_STYLE                                     = 0x3
	NANO_STYLE_NONE                                       = 0x1
	NANO_STYLE_CRYSTAL                                    = 0x0
	NANO_STYLE_ENERGY                                     = 0x1
	NANO_STYLE_FLUID                                      = 0x2
	SIZEOF_NANO_TYPE                                      = 0x4
	NANO_TYPE_POWER                                       = 0x0
	NANO_TYPE_ACCURACY                                    = 0x1
	NANO_TYPE_PROTECT                                     = 0x2
	NANO_TYPE_DODGE                                       = 0x3
	SIZEOF_NANO_TUNE_NEED_ITEM_SLOT                       = 0xa
	VALUE_ATTACK_MISS                                     = 0x1
	MSG_ONLINE                                            = 0x1
	MSG_BUSY                                              = 0x2
	MSG_OFFLINE                                           = 0x0
	SIZEOF_FREE_CHAT_STRING                               = 0x80
	SIZEOF_MENU_CHAT_STRING                               = 0x80
	SIZEOF_BUDDYLIST_SLOT                                 = 0x32
	SIZEOF_EMAIL_SUBJECT_STRING                           = 0x20
	SIZEOF_EMAIL_CONTENT_STRING                           = 0x200
	SIZEOF_EMAIL_PAGE_SIZE                                = 0x5
	SIZEOF_EMAIL_ITEM_CNT                                 = 0x4
	EMAIL_AND_MONEY_COST                                  = 0x32
	EMAIL_ITEM_COST                                       = 0x14
	BUDDYWARP_INTERVAL                                    = 0x3c
	EMAILSEND_TIME_DELAY                                  = 0x3c
	DB_ERROR_INVALID_DATA                                 = 0x1
	DB_ERROR_HACK_ATTEMPT                                 = 0x2
	DB_ERROR_ACCESS_FAIL                                  = 0x3
	DB_ERROR_PC_INSERT_FAIL                               = 0x4
	CALL_NPC_MAX_CNT                                      = 0x800
	CN_EP_RING_MAX_CNT                                    = 0x3e7
	HF_BIT_NONE                                           = 0x0
	HF_BIT_NORMAL                                         = 0x1
	HF_BIT_CRITICAL                                       = 0x2
	HF_BIT_STYLE_WIN                                      = 0x4
	HF_BIT_STYLE_TIE                                      = 0x8
	HF_BIT_STYLE_LOSE                                     = 0x10
	SKIN_COLOR_MAX                                        = 0xc
	HAIR_COLOR_MAX                                        = 0x12
	EYE_COLOR_MAX                                         = 0x5
	BODY_TYPE_MAX                                         = 0x3
	HEIGHT_TYPE_MAX                                       = 0x5
	CLASS_TYPE_MAX                                        = 0x4
	CN_EP_RACE_MODE_PRACTICE                              = 0x0
	CN_EP_RACE_MODE_RECORD                                = 0x1
	CN_EP_SECOM_NPC_TYPE_NUM                              = 0xd
	CN_EP_EECOM_NPC_TYPE_NUM                              = 0xe
	CN_EP_SIZE_SMALL                                      = 0x0
	CN_EP_SIZE_MIDDLE                                     = 0x1
	CN_EP_SIZE_BIG                                        = 0x2
	CN_EP_TICKET_ITEM_ID_SMALL                            = 0x73
	CN_EP_TICKET_ITEM_ID_MIDDLE                           = 0x74
	CN_EP_TICKET_ITEM_ID_BIG                              = 0x75
	CN_EP_TICKET_ITEM_ID_FREE                             = 0x76
	CN_EP_DISTANCE_ERROR_SAFE_RANGE                       = 0x4b0
	CN_ACCOUNT_LEVEL__MASTER                              = 0x1
	CN_ACCOUNT_LEVEL__POWER_DEVELOPER                     = 0xa
	CN_ACCOUNT_LEVEL__QA                                  = 0x14
	CN_ACCOUNT_LEVEL__GM                                  = 0x1e
	CN_ACCOUNT_LEVEL__CS                                  = 0x28
	CN_ACCOUNT_LEVEL__FREE_USER                           = 0x30
	CN_ACCOUNT_LEVEL__PAY_USER                            = 0x31
	CN_ACCOUNT_LEVEL__DEVELOPER                           = 0x32
	CN_ACCOUNT_LEVEL__CLOSEBETA_USER                      = 0x50
	CN_ACCOUNT_LEVEL__OPENBETA_USER                       = 0x55
	CN_ACCOUNT_LEVEL__USER                                = 0x63
	CN_SPECIAL_STATE_FLAG__PRINT_GM                       = 0x1
	CN_SPECIAL_STATE_FLAG__INVISIBLE                      = 0x2
	CN_SPECIAL_STATE_FLAG__INVULNERABLE                   = 0x4
	CN_SPECIAL_STATE_FLAG__FULL_UI                        = 0x10
	CN_SPECIAL_STATE_FLAG__COMBAT                         = 0x20
	CN_SPECIAL_STATE_FLAG__MUTE_FREECHAT                  = 0x40
	CN_GM_SET_VALUE_TYPE__HP                              = 0x1
	CN_GM_SET_VALUE_TYPE__WEAPON_BATTERY                  = 0x2
	CN_GM_SET_VALUE_TYPE__NANO_BATTERY                    = 0x3
	CN_GM_SET_VALUE_TYPE__FUSION_MATTER                   = 0x4
	CN_GM_SET_VALUE_TYPE__CANDY                           = 0x5
	CN_GM_SET_VALUE_TYPE__SPEED                           = 0x6
	CN_GM_SET_VALUE_TYPE__JUMP                            = 0x7
	CN_GM_SET_VALUE_TYPE__END                             = 0x8
	HEIGHT_CLIMBABLE                                      = 0x96
	CN_GROUP_WARP_CHECK_RANGE                             = 0x3e8
	WYVERN_LOCATION_FLAG_SIZE                             = 0x2
	CN_PC_EVENT_ID_GET_NANO_QUEST                         = 0x1
	CN_PC_EVENT_ID_DEFEAT_FUSE_AND_GET_NANO               = 0x2
	_dCN_STREETSTALL__ITEMLIST_COUNT_MAX                  = 0x5
	CSB_BIT_NONE                                          = 0x0
	CSB_BIT_UP_MOVE_SPEED                                 = 0x1
	CSB_BIT_UP_SWIM_SPEED                                 = 0x2
	CSB_BIT_UP_JUMP_HEIGHT                                = 0x4
	CSB_BIT_UP_STEALTH                                    = 0x8
	CSB_BIT_PHOENIX                                       = 0x10
	CSB_BIT_PROTECT_BATTERY                               = 0x20
	CSB_BIT_PROTECT_INFECTION                             = 0x40
	CSB_BIT_DN_MOVE_SPEED                                 = 0x80
	CSB_BIT_DN_ATTACK_SPEED                               = 0x100
	CSB_BIT_STUN                                          = 0x200
	CSB_BIT_MEZ                                           = 0x400
	CSB_BIT_KNOCKDOWN                                     = 0x800
	CSB_BIT_MINIMAP_ENEMY                                 = 0x1000
	CSB_BIT_MINIMAP_TRESURE                               = 0x2000
	CSB_BIT_REWARD_BLOB                                   = 0x4000
	CSB_BIT_REWARD_CASH                                   = 0x8000
	CSB_BIT_INFECTION                                     = 0x10000
	CSB_BIT_FREEDOM                                       = 0x20000
	CSB_BIT_BOUNDINGBALL                                  = 0x40000
	CSB_BIT_INVULNERABLE                                  = 0x80000
	CSB_BIT_STIMPAKSLOT1                                  = 0x100000
	CSB_BIT_STIMPAKSLOT2                                  = 0x200000
	CSB_BIT_STIMPAKSLOT3                                  = 0x400000
	CSB_BIT_HEAL                                          = 0x800000
	CSB_BIT_EXTRABANK                                     = 0x1000000
	TIME_BUFF_CONFIRM_KEY_MAX                             = 0x77359400
	READPACKET_SUCC                                       = 0x0
	READPACKET_FAIL                                       = 0x1
	READPACKET_RETURN                                     = 0x2
	BITMASK_FROM2TO                                       = 0xff000000
	BITMASK_FROM                                          = 0xf0000000
	BITMASK_TO                                            = 0xf000000
	BITMASK_SENDBLOCK                                     = 0x800000
	BITMASK_AUTHED                                        = 0x400000
	BITMASK_U_ID                                          = 0xfff
	CL2LS                                                 = 0x12000000
	CL2FE                                                 = 0x13000000
	LS2CL                                                 = 0x21000000
	LS2LS                                                 = 0x22000000
	LS2DBA                                                = 0x27000000
	FE2CL                                                 = 0x31000000
	FE2FE                                                 = 0x33000000
	FE2GS                                                 = 0x34000000
	FE2EP                                                 = 0x36000000
	FE2MSG                                                = 0x38000000
	GS2FE                                                 = 0x43000000
	GS2GS                                                 = 0x44000000
	GS2AI                                                 = 0x45000000
	GS2EP                                                 = 0x46000000
	GS2DBA                                                = 0x47000000
	GS2MSG                                                = 0x48000000
	GS2MGR                                                = 0x4a000000
	AI2GS                                                 = 0x54000000
	EP2FE                                                 = 0x63000000
	EP2GS                                                 = 0x64000000
	DBA2GS                                                = 0x74000000
	DBA2EP                                                = 0x75000000
	MSG2FE                                                = 0x83000000
	MSG2GS                                                = 0x84000000
	MSG2CMSG                                              = 0x89000000
	CMSG2MSG                                              = 0x98000000
	MGR2SPY                                               = 0xb3000000
	SPY2MGR                                               = 0xb4000000
	MGR2SA                                                = 0xb5000000
	SA2MGR                                                = 0xb6000000
	SA2SPY                                                = 0xb7000000
	SPY2SA                                                = 0xb8000000
	SPY2SVR                                               = 0xb9000000
	SVR2SPY                                               = 0xba000000
	SCH2SVR                                               = 0xc0000000
	SCH2LS                                                = 0xc2000000
	SCH2FE                                                = 0xc3000000
	SCH2GS                                                = 0xc4000000
	SCH2AI                                                = 0xc5000000
	SCH2EP                                                = 0xc6000000
	SCH2DBA                                               = 0xc7000000
	SCH2MSG                                               = 0xc8000000
	SCH2CMSG                                              = 0xc9000000
	CL2CDR                                                = 0x1f000000
	SENDBLOCK                                             = 0x800000
	AUTHED_X                                              = 0x0
	AUTHED_O                                              = 0x400000
	SEND_SVR_FE                                           = 0x1
	SEND_SVR_FE_ANY                                       = 0x2
	SEND_SVR_FE_ALL                                       = 0x3
	SEND_SVR_AI                                           = 0x4
	SEND_SVR_AI_ANY                                       = 0x5
	SEND_SVR_AI_ALL                                       = 0x6
	SEND_SVR_FE_AI_ALL                                    = 0x7
	SEND_SVR_DBA                                          = 0x8
	SEND_SVR_GS                                           = 0x9
	SEND_SVR_MSG                                          = 0xa
	SEND_SVR_MSG_ANY                                      = 0xb
	SEND_SVR_MSG_ALL                                      = 0xc
	SEND_UNICAST                                          = 0x1
	SEND_ANYCAST                                          = 0x2
	SEND_ANYCAST_NEW                                      = 0x3
	SEND_BROADCAST                                        = 0x4
	CN_PACKET_BUFFER_SIZE                                 = 0x1000
	P_CL2LS_REQ_LOGIN                                     = 0x12000001
	P_CL2LS_REQ_CHECK_CHAR_NAME                           = 0x12000002
	P_CL2LS_REQ_SAVE_CHAR_NAME                            = 0x12000003
	P_CL2LS_REQ_CHAR_CREATE                               = 0x12000004
	P_CL2LS_REQ_CHAR_SELECT                               = 0x12000005
	P_CL2LS_REQ_CHAR_DELETE                               = 0x12000006
	P_CL2LS_REQ_SHARD_SELECT                              = 0x12000007
	P_CL2LS_REQ_SHARD_LIST_INFO                           = 0x12000008
	P_CL2LS_CHECK_NAME_LIST                               = 0x12000009
	P_CL2LS_REQ_SAVE_CHAR_TUTOR                           = 0x1200000a
	P_CL2LS_REQ_PC_EXIT_DUPLICATE                         = 0x1200000b
	P_CL2LS_REP_LIVE_CHECK                                = 0x1200000c
	P_CL2LS_REQ_CHANGE_CHAR_NAME                          = 0x1200000d
	P_CL2LS_REQ_SERVER_SELECT                             = 0x1200000e
	P_CL2FE_REQ_PC_ENTER                                  = 0x13000001
	P_CL2FE_REQ_PC_EXIT                                   = 0x13000002
	P_CL2FE_REQ_PC_MOVE                                   = 0x13000003
	P_CL2FE_REQ_PC_STOP                                   = 0x13000004
	P_CL2FE_REQ_PC_JUMP                                   = 0x13000005
	P_CL2FE_REQ_PC_ATTACK_NPCs                            = 0x13000006
	P_CL2FE_REQ_SEND_FREECHAT_MESSAGE                     = 0x13000007
	P_CL2FE_REQ_SEND_MENUCHAT_MESSAGE                     = 0x13000008
	P_CL2FE_REQ_PC_REGEN                                  = 0x13000009
	P_CL2FE_REQ_ITEM_MOVE                                 = 0x1300000a
	P_CL2FE_REQ_PC_TASK_START                             = 0x1300000b
	P_CL2FE_REQ_PC_TASK_END                               = 0x1300000c
	P_CL2FE_REQ_NANO_EQUIP                                = 0x1300000d
	P_CL2FE_REQ_NANO_UNEQUIP                              = 0x1300000e
	P_CL2FE_REQ_NANO_ACTIVE                               = 0x1300000f
	P_CL2FE_REQ_NANO_TUNE                                 = 0x13000010
	P_CL2FE_REQ_NANO_SKILL_USE                            = 0x13000011
	P_CL2FE_REQ_PC_TASK_STOP                              = 0x13000012
	P_CL2FE_REQ_PC_TASK_CONTINUE                          = 0x13000013
	P_CL2FE_REQ_PC_GOTO                                   = 0x13000014
	P_CL2FE_REQ_CHARGE_NANO_STAMINA                       = 0x13000015
	P_CL2FE_REQ_PC_KILL_QUEST_NPCs                        = 0x13000016
	P_CL2FE_REQ_PC_VENDOR_ITEM_BUY                        = 0x13000017
	P_CL2FE_REQ_PC_VENDOR_ITEM_SELL                       = 0x13000018
	P_CL2FE_REQ_PC_ITEM_DELETE                            = 0x13000019
	P_CL2FE_REQ_PC_GIVE_ITEM                              = 0x1300001a
	P_CL2FE_REQ_PC_ROCKET_STYLE_READY                     = 0x1300001b
	P_CL2FE_REQ_PC_ROCKET_STYLE_FIRE                      = 0x1300001c
	P_CL2FE_REQ_PC_ROCKET_STYLE_HIT                       = 0x1300001d
	P_CL2FE_REQ_PC_GRENADE_STYLE_READY                    = 0x1300001e
	P_CL2FE_REQ_PC_GRENADE_STYLE_FIRE                     = 0x1300001f
	P_CL2FE_REQ_PC_GRENADE_STYLE_HIT                      = 0x13000020
	P_CL2FE_REQ_PC_NANO_CREATE                            = 0x13000021
	P_CL2FE_REQ_PC_TRADE_OFFER                            = 0x13000022
	P_CL2FE_REQ_PC_TRADE_OFFER_CANCEL                     = 0x13000023
	P_CL2FE_REQ_PC_TRADE_OFFER_ACCEPT                     = 0x13000024
	P_CL2FE_REQ_PC_TRADE_OFFER_REFUSAL                    = 0x13000025
	P_CL2FE_REQ_PC_TRADE_OFFER_ABORT                      = 0x13000026
	P_CL2FE_REQ_PC_TRADE_CONFIRM                          = 0x13000027
	P_CL2FE_REQ_PC_TRADE_CONFIRM_CANCEL                   = 0x13000028
	P_CL2FE_REQ_PC_TRADE_CONFIRM_ABORT                    = 0x13000029
	P_CL2FE_REQ_PC_TRADE_ITEM_REGISTER                    = 0x1300002a
	P_CL2FE_REQ_PC_TRADE_ITEM_UNREGISTER                  = 0x1300002b
	P_CL2FE_REQ_PC_TRADE_CASH_REGISTER                    = 0x1300002c
	P_CL2FE_REQ_PC_TRADE_EMOTES_CHAT                      = 0x1300002d
	P_CL2FE_REQ_PC_BANK_OPEN                              = 0x1300002e
	P_CL2FE_REQ_PC_BANK_CLOSE                             = 0x1300002f
	P_CL2FE_REQ_PC_VENDOR_START                           = 0x13000030
	P_CL2FE_REQ_PC_VENDOR_TABLE_UPDATE                    = 0x13000031
	P_CL2FE_REQ_PC_VENDOR_ITEM_RESTORE_BUY                = 0x13000032
	P_CL2FE_REQ_PC_COMBAT_BEGIN                           = 0x13000033
	P_CL2FE_REQ_PC_COMBAT_END                             = 0x13000034
	P_CL2FE_REQ_REQUEST_MAKE_BUDDY                        = 0x13000035
	P_CL2FE_REQ_ACCEPT_MAKE_BUDDY                         = 0x13000036
	P_CL2FE_REQ_SEND_BUDDY_FREECHAT_MESSAGE               = 0x13000037
	P_CL2FE_REQ_SEND_BUDDY_MENUCHAT_MESSAGE               = 0x13000038
	P_CL2FE_REQ_GET_BUDDY_STYLE                           = 0x13000039
	P_CL2FE_REQ_SET_BUDDY_BLOCK                           = 0x1300003a
	P_CL2FE_REQ_REMOVE_BUDDY                              = 0x1300003b
	P_CL2FE_REQ_GET_BUDDY_STATE                           = 0x1300003c
	P_CL2FE_REQ_PC_JUMPPAD                                = 0x1300003d
	P_CL2FE_REQ_PC_LAUNCHER                               = 0x1300003e
	P_CL2FE_REQ_PC_ZIPLINE                                = 0x1300003f
	P_CL2FE_REQ_PC_MOVEPLATFORM                           = 0x13000040
	P_CL2FE_REQ_PC_SLOPE                                  = 0x13000041
	P_CL2FE_REQ_PC_STATE_CHANGE                           = 0x13000042
	P_CL2FE_REQ_PC_MAP_WARP                               = 0x13000043
	P_CL2FE_REQ_PC_GIVE_NANO                              = 0x13000044
	P_CL2FE_REQ_NPC_SUMMON                                = 0x13000045
	P_CL2FE_REQ_NPC_UNSUMMON                              = 0x13000046
	P_CL2FE_REQ_ITEM_CHEST_OPEN                           = 0x13000047
	P_CL2FE_REQ_PC_GIVE_NANO_SKILL                        = 0x13000048
	P_CL2FE_DOT_DAMAGE_ONOFF                              = 0x13000049
	P_CL2FE_REQ_PC_VENDOR_BATTERY_BUY                     = 0x1300004a
	P_CL2FE_REQ_PC_WARP_USE_NPC                           = 0x1300004b
	P_CL2FE_REQ_PC_GROUP_INVITE                           = 0x1300004c
	P_CL2FE_REQ_PC_GROUP_INVITE_REFUSE                    = 0x1300004d
	P_CL2FE_REQ_PC_GROUP_JOIN                             = 0x1300004e
	P_CL2FE_REQ_PC_GROUP_LEAVE                            = 0x1300004f
	P_CL2FE_REQ_PC_AVATAR_EMOTES_CHAT                     = 0x13000050
	P_CL2FE_REQ_PC_BUDDY_WARP                             = 0x13000051
	P_CL2FE_REQ_GET_MEMBER_STYLE                          = 0x13000052
	P_CL2FE_REQ_GET_GROUP_STYLE                           = 0x13000053
	P_CL2FE_REQ_PC_CHANGE_MENTOR                          = 0x13000054
	P_CL2FE_REQ_GET_BUDDY_LOCATION                        = 0x13000055
	P_CL2FE_REQ_NPC_GROUP_SUMMON                          = 0x13000056
	P_CL2FE_REQ_PC_WARP_TO_PC                             = 0x13000057
	P_CL2FE_REQ_EP_RANK_GET_LIST                          = 0x13000058
	P_CL2FE_REQ_EP_RANK_GET_DETAIL                        = 0x13000059
	P_CL2FE_REQ_EP_RANK_GET_PC_INFO                       = 0x1300005a
	P_CL2FE_REQ_EP_RACE_START                             = 0x1300005b
	P_CL2FE_REQ_EP_RACE_END                               = 0x1300005c
	P_CL2FE_REQ_EP_RACE_CANCEL                            = 0x1300005d
	P_CL2FE_REQ_EP_GET_RING                               = 0x1300005e
	P_CL2FE_REQ_IM_CHANGE_SWITCH_STATUS                   = 0x1300005f
	P_CL2FE_REQ_SHINY_PICKUP                              = 0x13000060
	P_CL2FE_REQ_SHINY_SUMMON                              = 0x13000061
	P_CL2FE_REQ_PC_MOVETRANSPORTATION                     = 0x13000062
	P_CL2FE_REQ_SEND_ALL_GROUP_FREECHAT_MESSAGE           = 0x13000063
	P_CL2FE_REQ_SEND_ANY_GROUP_FREECHAT_MESSAGE           = 0x13000064
	P_CL2FE_REQ_BARKER                                    = 0x13000065
	P_CL2FE_REQ_SEND_ALL_GROUP_MENUCHAT_MESSAGE           = 0x13000066
	P_CL2FE_REQ_SEND_ANY_GROUP_MENUCHAT_MESSAGE           = 0x13000067
	P_CL2FE_REQ_REGIST_TRANSPORTATION_LOCATION            = 0x13000068
	P_CL2FE_REQ_PC_WARP_USE_TRANSPORTATION                = 0x13000069
	P_CL2FE_GM_REQ_PC_SPECIAL_STATE_SWITCH                = 0x1300006a
	P_CL2FE_GM_REQ_PC_SET_VALUE                           = 0x1300006b
	P_CL2FE_GM_REQ_KICK_PLAYER                            = 0x1300006c
	P_CL2FE_GM_REQ_TARGET_PC_TELEPORT                     = 0x1300006d
	P_CL2FE_GM_REQ_PC_LOCATION                            = 0x1300006e
	P_CL2FE_GM_REQ_PC_ANNOUNCE                            = 0x1300006f
	P_CL2FE_REQ_SET_PC_BLOCK                              = 0x13000070
	P_CL2FE_REQ_REGIST_RXCOM                              = 0x13000071
	P_CL2FE_GM_REQ_PC_MOTD_REGISTER                       = 0x13000072
	P_CL2FE_REQ_ITEM_USE                                  = 0x13000073
	P_CL2FE_REQ_WARP_USE_RECALL                           = 0x13000074
	P_CL2FE_REP_LIVE_CHECK                                = 0x13000075
	P_CL2FE_REQ_PC_MISSION_COMPLETE                       = 0x13000076
	P_CL2FE_REQ_PC_TASK_COMPLETE                          = 0x13000077
	P_CL2FE_REQ_NPC_INTERACTION                           = 0x13000078
	P_CL2FE_DOT_HEAL_ONOFF                                = 0x13000079
	P_CL2FE_REQ_PC_SPECIAL_STATE_SWITCH                   = 0x1300007a
	P_CL2FE_REQ_PC_EMAIL_UPDATE_CHECK                     = 0x1300007b
	P_CL2FE_REQ_PC_READ_EMAIL                             = 0x1300007c
	P_CL2FE_REQ_PC_RECV_EMAIL_PAGE_LIST                   = 0x1300007d
	P_CL2FE_REQ_PC_DELETE_EMAIL                           = 0x1300007e
	P_CL2FE_REQ_PC_SEND_EMAIL                             = 0x1300007f
	P_CL2FE_REQ_PC_RECV_EMAIL_ITEM                        = 0x13000080
	P_CL2FE_REQ_PC_RECV_EMAIL_CANDY                       = 0x13000081
	P_CL2FE_GM_REQ_TARGET_PC_SPECIAL_STATE_ONOFF          = 0x13000082
	P_CL2FE_REQ_PC_SET_CURRENT_MISSION_ID                 = 0x13000083
	P_CL2FE_REQ_NPC_GROUP_INVITE                          = 0x13000084
	P_CL2FE_REQ_NPC_GROUP_KICK                            = 0x13000085
	P_CL2FE_REQ_PC_FIRST_USE_FLAG_SET                     = 0x13000086
	P_CL2FE_REQ_PC_TRANSPORT_WARP                         = 0x13000087
	P_CL2FE_REQ_PC_TIME_TO_GO_WARP                        = 0x13000088
	P_CL2FE_REQ_PC_RECV_EMAIL_ITEM_ALL                    = 0x13000089
	P_CL2FE_REQ_CHANNEL_INFO                              = 0x1300008a
	P_CL2FE_REQ_PC_CHANNEL_NUM                            = 0x1300008b
	P_CL2FE_REQ_PC_WARP_CHANNEL                           = 0x1300008c
	P_CL2FE_REQ_PC_LOADING_COMPLETE                       = 0x1300008d
	P_CL2FE_REQ_PC_FIND_NAME_MAKE_BUDDY                   = 0x1300008e
	P_CL2FE_REQ_PC_FIND_NAME_ACCEPT_BUDDY                 = 0x1300008f
	P_CL2FE_REQ_PC_ATTACK_CHARs                           = 0x13000090
	P_CL2FE_PC_STREETSTALL_REQ_READY                      = 0x13000091
	P_CL2FE_PC_STREETSTALL_REQ_CANCEL                     = 0x13000092
	P_CL2FE_PC_STREETSTALL_REQ_REGIST_ITEM                = 0x13000093
	P_CL2FE_PC_STREETSTALL_REQ_UNREGIST_ITEM              = 0x13000094
	P_CL2FE_PC_STREETSTALL_REQ_SALE_START                 = 0x13000095
	P_CL2FE_PC_STREETSTALL_REQ_ITEM_LIST                  = 0x13000096
	P_CL2FE_PC_STREETSTALL_REQ_ITEM_BUY                   = 0x13000097
	P_CL2FE_REQ_PC_ITEM_COMBINATION                       = 0x13000098
	P_CL2FE_GM_REQ_SET_PC_SKILL                           = 0x13000099
	P_CL2FE_REQ_PC_SKILL_ADD                              = 0x1300009a
	P_CL2FE_REQ_PC_SKILL_DEL                              = 0x1300009b
	P_CL2FE_REQ_PC_SKILL_USE                              = 0x1300009c
	P_CL2FE_REQ_PC_ROPE                                   = 0x1300009d
	P_CL2FE_REQ_PC_BELT                                   = 0x1300009e
	P_CL2FE_REQ_PC_VEHICLE_ON                             = 0x1300009f
	P_CL2FE_REQ_PC_VEHICLE_OFF                            = 0x130000a0
	P_CL2FE_REQ_PC_REGIST_QUICK_SLOT                      = 0x130000a1
	P_CL2FE_REQ_PC_DISASSEMBLE_ITEM                       = 0x130000a2
	P_CL2FE_GM_REQ_REWARD_RATE                            = 0x130000a3
	P_CL2FE_REQ_PC_ITEM_ENCHANT                           = 0x130000a4
	P_FE2CL_ERROR                                         = 0x31000000
	P_FE2CL_REP_PC_ENTER_FAIL                             = 0x31000001
	P_FE2CL_REP_PC_ENTER_SUCC                             = 0x31000002
	P_FE2CL_PC_NEW                                        = 0x31000003
	P_FE2CL_REP_PC_EXIT_FAIL                              = 0x31000004
	P_FE2CL_REP_PC_EXIT_SUCC                              = 0x31000005
	P_FE2CL_PC_EXIT                                       = 0x31000006
	P_FE2CL_PC_AROUND                                     = 0x31000007
	P_FE2CL_PC_MOVE                                       = 0x31000008
	P_FE2CL_PC_STOP                                       = 0x31000009
	P_FE2CL_PC_JUMP                                       = 0x3100000a
	P_FE2CL_NPC_ENTER                                     = 0x3100000b
	P_FE2CL_NPC_EXIT                                      = 0x3100000c
	P_FE2CL_NPC_MOVE                                      = 0x3100000d
	P_FE2CL_NPC_NEW                                       = 0x3100000e
	P_FE2CL_NPC_AROUND                                    = 0x3100000f
	P_FE2CL_AROUND_DEL_PC                                 = 0x31000010
	P_FE2CL_AROUND_DEL_NPC                                = 0x31000011
	P_FE2CL_REP_SEND_FREECHAT_MESSAGE_SUCC                = 0x31000012
	P_FE2CL_REP_SEND_FREECHAT_MESSAGE_FAIL                = 0x31000013
	P_FE2CL_PC_ATTACK_NPCs_SUCC                           = 0x31000014
	P_FE2CL_PC_ATTACK_NPCs                                = 0x31000015
	P_FE2CL_NPC_ATTACK_PCs                                = 0x31000016
	P_FE2CL_REP_PC_REGEN_SUCC                             = 0x31000017
	P_FE2CL_REP_SEND_MENUCHAT_MESSAGE_SUCC                = 0x31000018
	P_FE2CL_REP_SEND_MENUCHAT_MESSAGE_FAIL                = 0x31000019
	P_FE2CL_PC_ITEM_MOVE_SUCC                             = 0x3100001a
	P_FE2CL_PC_EQUIP_CHANGE                               = 0x3100001b
	P_FE2CL_REP_PC_TASK_START_SUCC                        = 0x3100001c
	P_FE2CL_REP_PC_TASK_START_FAIL                        = 0x3100001d
	P_FE2CL_REP_PC_TASK_END_SUCC                          = 0x3100001e
	P_FE2CL_REP_PC_TASK_END_FAIL                          = 0x3100001f
	P_FE2CL_NPC_SKILL_READY                               = 0x31000020
	P_FE2CL_NPC_SKILL_FIRE                                = 0x31000021
	P_FE2CL_NPC_SKILL_HIT                                 = 0x31000022
	P_FE2CL_NPC_SKILL_CORRUPTION_READY                    = 0x31000023
	P_FE2CL_NPC_SKILL_CORRUPTION_HIT                      = 0x31000024
	P_FE2CL_NPC_SKILL_CANCEL                              = 0x31000025
	P_FE2CL_REP_NANO_EQUIP_SUCC                           = 0x31000026
	P_FE2CL_REP_NANO_UNEQUIP_SUCC                         = 0x31000027
	P_FE2CL_REP_NANO_ACTIVE_SUCC                          = 0x31000028
	P_FE2CL_REP_NANO_TUNE_SUCC                            = 0x31000029
	P_FE2CL_NANO_ACTIVE                                   = 0x3100002a
	P_FE2CL_NANO_SKILL_USE_SUCC                           = 0x3100002b
	P_FE2CL_NANO_SKILL_USE                                = 0x3100002c
	P_FE2CL_REP_PC_TASK_STOP_SUCC                         = 0x3100002d
	P_FE2CL_REP_PC_TASK_STOP_FAIL                         = 0x3100002e
	P_FE2CL_REP_PC_TASK_CONTINUE_SUCC                     = 0x3100002f
	P_FE2CL_REP_PC_TASK_CONTINUE_FAIL                     = 0x31000030
	P_FE2CL_REP_PC_GOTO_SUCC                              = 0x31000031
	P_FE2CL_REP_CHARGE_NANO_STAMINA                       = 0x31000032
	P_FE2CL_REP_PC_TICK                                   = 0x31000033
	P_FE2CL_REP_PC_KILL_QUEST_NPCs_SUCC                   = 0x31000034
	P_FE2CL_REP_PC_VENDOR_ITEM_BUY_SUCC                   = 0x31000035
	P_FE2CL_REP_PC_VENDOR_ITEM_BUY_FAIL                   = 0x31000036
	P_FE2CL_REP_PC_VENDOR_ITEM_SELL_SUCC                  = 0x31000037
	P_FE2CL_REP_PC_VENDOR_ITEM_SELL_FAIL                  = 0x31000038
	P_FE2CL_REP_PC_ITEM_DELETE_SUCC                       = 0x31000039
	P_FE2CL_PC_ROCKET_STYLE_READY                         = 0x3100003a
	P_FE2CL_REP_PC_ROCKET_STYLE_FIRE_SUCC                 = 0x3100003b
	P_FE2CL_PC_ROCKET_STYLE_FIRE                          = 0x3100003c
	P_FE2CL_PC_ROCKET_STYLE_HIT                           = 0x3100003d
	P_FE2CL_PC_GRENADE_STYLE_READY                        = 0x3100003e
	P_FE2CL_REP_PC_GRENADE_STYLE_FIRE_SUCC                = 0x3100003f
	P_FE2CL_PC_GRENADE_STYLE_FIRE                         = 0x31000040
	P_FE2CL_PC_GRENADE_STYLE_HIT                          = 0x31000041
	P_FE2CL_REP_PC_TRADE_OFFER                            = 0x31000042
	P_FE2CL_REP_PC_TRADE_OFFER_CANCEL                     = 0x31000043
	P_FE2CL_REP_PC_TRADE_OFFER_SUCC                       = 0x31000044
	P_FE2CL_REP_PC_TRADE_OFFER_REFUSAL                    = 0x31000045
	P_FE2CL_REP_PC_TRADE_OFFER_ABORT                      = 0x31000046
	P_FE2CL_REP_PC_TRADE_CONFIRM                          = 0x31000047
	P_FE2CL_REP_PC_TRADE_CONFIRM_CANCEL                   = 0x31000048
	P_FE2CL_REP_PC_TRADE_CONFIRM_ABORT                    = 0x31000049
	P_FE2CL_REP_PC_TRADE_CONFIRM_SUCC                     = 0x3100004a
	P_FE2CL_REP_PC_TRADE_CONFIRM_FAIL                     = 0x3100004b
	P_FE2CL_REP_PC_TRADE_ITEM_REGISTER_SUCC               = 0x3100004c
	P_FE2CL_REP_PC_TRADE_ITEM_REGISTER_FAIL               = 0x3100004d
	P_FE2CL_REP_PC_TRADE_ITEM_UNREGISTER_SUCC             = 0x3100004e
	P_FE2CL_REP_PC_TRADE_ITEM_UNREGISTER_FAIL             = 0x3100004f
	P_FE2CL_REP_PC_TRADE_CASH_REGISTER_SUCC               = 0x31000050
	P_FE2CL_REP_PC_TRADE_CASH_REGISTER_FAIL               = 0x31000051
	P_FE2CL_REP_PC_TRADE_EMOTES_CHAT                      = 0x31000052
	P_FE2CL_REP_PC_NANO_CREATE_SUCC                       = 0x31000053
	P_FE2CL_REP_PC_NANO_CREATE_FAIL                       = 0x31000054
	P_FE2CL_REP_NANO_TUNE_FAIL                            = 0x31000055
	P_FE2CL_REP_PC_BANK_OPEN_SUCC                         = 0x31000056
	P_FE2CL_REP_PC_BANK_OPEN_FAIL                         = 0x31000057
	P_FE2CL_REP_PC_BANK_CLOSE_SUCC                        = 0x31000058
	P_FE2CL_REP_PC_BANK_CLOSE_FAIL                        = 0x31000059
	P_FE2CL_REP_PC_VENDOR_START_SUCC                      = 0x3100005a
	P_FE2CL_REP_PC_VENDOR_START_FAIL                      = 0x3100005b
	P_FE2CL_REP_PC_VENDOR_TABLE_UPDATE_SUCC               = 0x3100005c
	P_FE2CL_REP_PC_VENDOR_TABLE_UPDATE_FAIL               = 0x3100005d
	P_FE2CL_REP_PC_VENDOR_ITEM_RESTORE_BUY_SUCC           = 0x3100005e
	P_FE2CL_REP_PC_VENDOR_ITEM_RESTORE_BUY_FAIL           = 0x3100005f
	P_FE2CL_CHAR_TIME_BUFF_TIME_OUT                       = 0x31000060
	P_FE2CL_REP_PC_GIVE_ITEM_SUCC                         = 0x31000061
	P_FE2CL_REP_PC_GIVE_ITEM_FAIL                         = 0x31000062
	P_FE2CL_REP_PC_BUDDYLIST_INFO_SUCC                    = 0x31000063
	P_FE2CL_REP_PC_BUDDYLIST_INFO_FAIL                    = 0x31000064
	P_FE2CL_REP_REQUEST_MAKE_BUDDY_SUCC                   = 0x83000065
	P_FE2CL_REP_REQUEST_MAKE_BUDDY_FAIL                   = 0x31000066
	P_FE2CL_REP_ACCEPT_MAKE_BUDDY_SUCC                    = 0x31000067
	P_FE2CL_REP_ACCEPT_MAKE_BUDDY_FAIL                    = 0x31000068
	P_FE2CL_REP_SEND_BUDDY_FREECHAT_MESSAGE_SUCC          = 0x31000069
	P_FE2CL_REP_SEND_BUDDY_FREECHAT_MESSAGE_FAIL          = 0x3100006a
	P_FE2CL_REP_SEND_BUDDY_MENUCHAT_MESSAGE_SUCC          = 0x3100006b
	P_FE2CL_REP_SEND_BUDDY_MENUCHAT_MESSAGE_FAIL          = 0x3100006c
	P_FE2CL_REP_GET_BUDDY_STYLE_SUCC                      = 0x3100006d
	P_FE2CL_REP_GET_BUDDY_STYLE_FAIL                      = 0x3100006e
	P_FE2CL_REP_GET_BUDDY_STATE_SUCC                      = 0x3100006f
	P_FE2CL_REP_GET_BUDDY_STATE_FAIL                      = 0x31000070
	P_FE2CL_REP_SET_BUDDY_BLOCK_SUCC                      = 0x31000071
	P_FE2CL_REP_SET_BUDDY_BLOCK_FAIL                      = 0x31000072
	P_FE2CL_REP_REMOVE_BUDDY_SUCC                         = 0x31000073
	P_FE2CL_REP_REMOVE_BUDDY_FAIL                         = 0x31000074
	P_FE2CL_PC_JUMPPAD                                    = 0x31000075
	P_FE2CL_PC_LAUNCHER                                   = 0x31000076
	P_FE2CL_PC_ZIPLINE                                    = 0x31000077
	P_FE2CL_PC_MOVEPLATFORM                               = 0x31000078
	P_FE2CL_PC_SLOPE                                      = 0x31000079
	P_FE2CL_PC_STATE_CHANGE                               = 0x3100007a
	P_FE2CL_REP_REQUEST_MAKE_BUDDY_SUCC_TO_ACCEPTER       = 0x3100007b
	P_FE2CL_REP_REWARD_ITEM                               = 0x3100007c
	P_FE2CL_REP_ITEM_CHEST_OPEN_SUCC                      = 0x3100007d
	P_FE2CL_REP_ITEM_CHEST_OPEN_FAIL                      = 0x3100007e
	P_FE2CL_CHAR_TIME_BUFF_TIME_TICK                      = 0x3100007f
	P_FE2CL_REP_PC_VENDOR_BATTERY_BUY_SUCC                = 0x31000080
	P_FE2CL_REP_PC_VENDOR_BATTERY_BUY_FAIL                = 0x31000081
	P_FE2CL_NPC_ROCKET_STYLE_FIRE                         = 0x31000082
	P_FE2CL_NPC_GRENADE_STYLE_FIRE                        = 0x31000083
	P_FE2CL_NPC_BULLET_STYLE_HIT                          = 0x31000084
	P_FE2CL_CHARACTER_ATTACK_CHARACTERs                   = 0x31000085
	P_FE2CL_PC_GROUP_INVITE                               = 0x31000086
	P_FE2CL_PC_GROUP_INVITE_FAIL                          = 0x31000087
	P_FE2CL_PC_GROUP_INVITE_REFUSE                        = 0x31000088
	P_FE2CL_PC_GROUP_JOIN                                 = 0x31000089
	P_FE2CL_PC_GROUP_JOIN_FAIL                            = 0x3100008a
	P_FE2CL_PC_GROUP_JOIN_SUCC                            = 0x3100008b
	P_FE2CL_PC_GROUP_LEAVE                                = 0x3100008c
	P_FE2CL_PC_GROUP_LEAVE_FAIL                           = 0x3100008d
	P_FE2CL_PC_GROUP_LEAVE_SUCC                           = 0x3100008e
	P_FE2CL_PC_GROUP_MEMBER_INFO                          = 0x3100008f
	P_FE2CL_REP_PC_WARP_USE_NPC_SUCC                      = 0x31000090
	P_FE2CL_REP_PC_WARP_USE_NPC_FAIL                      = 0x31000091
	P_FE2CL_REP_PC_AVATAR_EMOTES_CHAT                     = 0x31000092
	P_FE2CL_REP_PC_CHANGE_MENTOR_SUCC                     = 0x31000093
	P_FE2CL_REP_PC_CHANGE_MENTOR_FAIL                     = 0x31000094
	P_FE2CL_REP_GET_MEMBER_STYLE_FAIL                     = 0x31000095
	P_FE2CL_REP_GET_MEMBER_STYLE_SUCC                     = 0x31000096
	P_FE2CL_REP_GET_GROUP_STYLE_FAIL                      = 0x31000097
	P_FE2CL_REP_GET_GROUP_STYLE_SUCC                      = 0x31000098
	P_FE2CL_PC_REGEN                                      = 0x31000099
	P_FE2CL_INSTANCE_MAP_INFO                             = 0x3100009a
	P_FE2CL_TRANSPORTATION_ENTER                          = 0x3100009b
	P_FE2CL_TRANSPORTATION_EXIT                           = 0x3100009c
	P_FE2CL_TRANSPORTATION_MOVE                           = 0x3100009d
	P_FE2CL_TRANSPORTATION_NEW                            = 0x3100009e
	P_FE2CL_TRANSPORTATION_AROUND                         = 0x3100009f
	P_FE2CL_AROUND_DEL_TRANSPORTATION                     = 0x310000a0
	P_FE2CL_REP_EP_RANK_LIST                              = 0x310000a1
	P_FE2CL_REP_EP_RANK_DETAIL                            = 0x310000a2
	P_FE2CL_REP_EP_RANK_PC_INFO                           = 0x310000a3
	P_FE2CL_REP_EP_RACE_START_SUCC                        = 0x310000a4
	P_FE2CL_REP_EP_RACE_START_FAIL                        = 0x310000a5
	P_FE2CL_REP_EP_RACE_END_SUCC                          = 0x310000a6
	P_FE2CL_REP_EP_RACE_END_FAIL                          = 0x310000a7
	P_FE2CL_REP_EP_RACE_CANCEL_SUCC                       = 0x310000a8
	P_FE2CL_REP_EP_RACE_CANCEL_FAIL                       = 0x310000a9
	P_FE2CL_REP_EP_GET_RING_SUCC                          = 0x310000aa
	P_FE2CL_REP_EP_GET_RING_FAIL                          = 0x310000ab
	P_FE2CL_REP_IM_CHANGE_SWITCH_STATUS                   = 0x310000ac
	P_FE2CL_SHINY_ENTER                                   = 0x310000ad
	P_FE2CL_SHINY_EXIT                                    = 0x310000ae
	P_FE2CL_SHINY_NEW                                     = 0x310000af
	P_FE2CL_SHINY_AROUND                                  = 0x310000b0
	P_FE2CL_AROUND_DEL_SHINY                              = 0x310000b1
	P_FE2CL_REP_SHINY_PICKUP_FAIL                         = 0x310000b2
	P_FE2CL_REP_SHINY_PICKUP_SUCC                         = 0x310000b3
	P_FE2CL_PC_MOVETRANSPORTATION                         = 0x310000b4
	P_FE2CL_REP_SEND_ALL_GROUP_FREECHAT_MESSAGE_SUCC      = 0x310000b5
	P_FE2CL_REP_SEND_ALL_GROUP_FREECHAT_MESSAGE_FAIL      = 0x310000b6
	P_FE2CL_REP_SEND_ANY_GROUP_FREECHAT_MESSAGE_SUCC      = 0x310000b7
	P_FE2CL_REP_SEND_ANY_GROUP_FREECHAT_MESSAGE_FAIL      = 0x310000b8
	P_FE2CL_REP_BARKER                                    = 0x310000b9
	P_FE2CL_REP_SEND_ALL_GROUP_MENUCHAT_MESSAGE_SUCC      = 0x310000ba
	P_FE2CL_REP_SEND_ALL_GROUP_MENUCHAT_MESSAGE_FAIL      = 0x310000bb
	P_FE2CL_REP_SEND_ANY_GROUP_MENUCHAT_MESSAGE_SUCC      = 0x310000bc
	P_FE2CL_REP_SEND_ANY_GROUP_MENUCHAT_MESSAGE_FAIL      = 0x310000bd
	P_FE2CL_REP_PC_REGIST_TRANSPORTATION_LOCATION_FAIL    = 0x310000be
	P_FE2CL_REP_PC_REGIST_TRANSPORTATION_LOCATION_SUCC    = 0x310000bf
	P_FE2CL_REP_PC_WARP_USE_TRANSPORTATION_FAIL           = 0x310000c0
	P_FE2CL_REP_PC_WARP_USE_TRANSPORTATION_SUCC           = 0x310000c1
	P_FE2CL_ANNOUNCE_MSG                                  = 0x310000c2
	P_FE2CL_REP_PC_SPECIAL_STATE_SWITCH_SUCC              = 0x310000c3
	P_FE2CL_PC_SPECIAL_STATE_CHANGE                       = 0x310000c4
	P_FE2CL_GM_REP_PC_SET_VALUE                           = 0x310000c5
	P_FE2CL_GM_PC_CHANGE_VALUE                            = 0x310000c6
	P_FE2CL_GM_REP_PC_LOCATION                            = 0x310000c7
	P_FE2CL_GM_REP_PC_ANNOUNCE                            = 0x310000c8
	P_FE2CL_REP_PC_BUDDY_WARP_FAIL                        = 0x310000c9
	P_FE2CL_REP_PC_CHANGE_LEVEL                           = 0x310000ca
	P_FE2CL_REP_SET_PC_BLOCK_SUCC                         = 0x310000cb
	P_FE2CL_REP_SET_PC_BLOCK_FAIL                         = 0x310000cc
	P_FE2CL_REP_REGIST_RXCOM                              = 0x310000cd
	P_FE2CL_REP_REGIST_RXCOM_FAIL                         = 0x310000ce
	P_FE2CL_PC_INVEN_FULL_MSG                             = 0x310000cf
	P_FE2CL_REQ_LIVE_CHECK                                = 0x310000d0
	P_FE2CL_PC_MOTD_LOGIN                                 = 0x310000d1
	P_FE2CL_REP_PC_ITEM_USE_FAIL                          = 0x310000d2
	P_FE2CL_REP_PC_ITEM_USE_SUCC                          = 0x310000d3
	P_FE2CL_PC_ITEM_USE                                   = 0x310000d4
	P_FE2CL_REP_GET_BUDDY_LOCATION_SUCC                   = 0x310000d5
	P_FE2CL_REP_GET_BUDDY_LOCATION_FAIL                   = 0x310000d6
	P_FE2CL_REP_PC_RIDING_FAIL                            = 0x310000d7
	P_FE2CL_REP_PC_RIDING_SUCC                            = 0x310000d8
	P_FE2CL_PC_RIDING                                     = 0x310000d9
	P_FE2CL_PC_BROOMSTICK_MOVE                            = 0x310000da
	P_FE2CL_REP_PC_BUDDY_WARP_OTHER_SHARD_SUCC            = 0x310000db
	P_FE2CL_REP_WARP_USE_RECALL_FAIL                      = 0x310000dc
	P_FE2CL_REP_PC_EXIT_DUPLICATE                         = 0x310000dd
	P_FE2CL_REP_PC_MISSION_COMPLETE_SUCC                  = 0x310000de
	P_FE2CL_PC_BUFF_UPDATE                                = 0x310000df
	P_FE2CL_REP_PC_NEW_EMAIL                              = 0x310000e0
	P_FE2CL_REP_PC_READ_EMAIL_SUCC                        = 0x310000e1
	P_FE2CL_REP_PC_READ_EMAIL_FAIL                        = 0x310000e2
	P_FE2CL_REP_PC_RECV_EMAIL_PAGE_LIST_SUCC              = 0x310000e3
	P_FE2CL_REP_PC_RECV_EMAIL_PAGE_LIST_FAIL              = 0x310000e4
	P_FE2CL_REP_PC_DELETE_EMAIL_SUCC                      = 0x310000e5
	P_FE2CL_REP_PC_DELETE_EMAIL_FAIL                      = 0x310000e6
	P_FE2CL_REP_PC_SEND_EMAIL_SUCC                        = 0x310000e7
	P_FE2CL_REP_PC_SEND_EMAIL_FAIL                        = 0x310000e8
	P_FE2CL_REP_PC_RECV_EMAIL_ITEM_SUCC                   = 0x310000e9
	P_FE2CL_REP_PC_RECV_EMAIL_ITEM_FAIL                   = 0x310000ea
	P_FE2CL_REP_PC_RECV_EMAIL_CANDY_SUCC                  = 0x310000eb
	P_FE2CL_REP_PC_RECV_EMAIL_CANDY_FAIL                  = 0x310000ec
	P_FE2CL_PC_SUDDEN_DEAD                                = 0x310000ed
	P_FE2CL_REP_GM_REQ_TARGET_PC_SPECIAL_STATE_ONOFF_SUCC = 0x310000ee
	P_FE2CL_REP_PC_SET_CURRENT_MISSION_ID                 = 0x310000ef
	P_FE2CL_REP_NPC_GROUP_INVITE_FAIL                     = 0x310000f0
	P_FE2CL_REP_NPC_GROUP_INVITE_SUCC                     = 0x310000f1
	P_FE2CL_REP_NPC_GROUP_KICK_FAIL                       = 0x310000f2
	P_FE2CL_REP_NPC_GROUP_KICK_SUCC                       = 0x310000f3
	P_FE2CL_PC_EVENT                                      = 0x310000f4
	P_FE2CL_REP_PC_TRANSPORT_WARP_SUCC                    = 0x310000f5
	P_FE2CL_REP_PC_TRADE_EMOTES_CHAT_FAIL                 = 0x310000f6
	P_FE2CL_REP_PC_RECV_EMAIL_ITEM_ALL_SUCC               = 0x310000f7
	P_FE2CL_REP_PC_RECV_EMAIL_ITEM_ALL_FAIL               = 0x310000f8
	P_FE2CL_REP_PC_LOADING_COMPLETE_SUCC                  = 0x310000f9
	P_FE2CL_REP_CHANNEL_INFO                              = 0x310000fa
	P_FE2CL_REP_PC_CHANNEL_NUM                            = 0x310000fb
	P_FE2CL_REP_PC_WARP_CHANNEL_FAIL                      = 0x310000fc
	P_FE2CL_REP_PC_WARP_CHANNEL_SUCC                      = 0x310000fd
	P_FE2CL_REP_PC_FIND_NAME_MAKE_BUDDY_SUCC              = 0x310000fe
	P_FE2CL_REP_PC_FIND_NAME_MAKE_BUDDY_FAIL              = 0x310000ff
	P_FE2CL_REP_PC_FIND_NAME_ACCEPT_BUDDY_FAIL            = 0x31000100
	P_FE2CL_REP_PC_BUDDY_WARP_SAME_SHARD_SUCC             = 0x31000101
	P_FE2CL_PC_ATTACK_CHARs_SUCC                          = 0x31000102
	P_FE2CL_PC_ATTACK_CHARs                               = 0x31000103
	P_FE2CL_NPC_ATTACK_CHARs                              = 0x31000104
	P_FE2CL_REP_PC_CHANGE_LEVEL_SUCC                      = 0x31000105
	P_FE2CL_REP_PC_NANO_CREATE                            = 0x31000106
	P_FE2CL_PC_STREETSTALL_REP_READY_SUCC                 = 0x31000107
	P_FE2CL_PC_STREETSTALL_REP_READY_FAIL                 = 0x31000108
	P_FE2CL_PC_STREETSTALL_REP_CANCEL_SUCC                = 0x31000109
	P_FE2CL_PC_STREETSTALL_REP_CANCEL_FAIL                = 0x3100010a
	P_FE2CL_PC_STREETSTALL_REP_REGIST_ITEM_SUCC           = 0x3100010b
	P_FE2CL_PC_STREETSTALL_REP_REGIST_ITEM_FAIL           = 0x3100010c
	P_FE2CL_PC_STREETSTALL_REP_UNREGIST_ITEM_SUCC         = 0x3100010d
	P_FE2CL_PC_STREETSTALL_REP_UNREGIST_ITEM_FAIL         = 0x3100010e
	P_FE2CL_PC_STREETSTALL_REP_SALE_START_SUCC            = 0x3100010f
	P_FE2CL_PC_STREETSTALL_REP_SALE_START_FAIL            = 0x31000110
	P_FE2CL_PC_STREETSTALL_REP_ITEM_LIST                  = 0x31000111
	P_FE2CL_PC_STREETSTALL_REP_ITEM_LIST_FAIL             = 0x31000112
	P_FE2CL_PC_STREETSTALL_REP_ITEM_BUY_SUCC_BUYER        = 0x31000113
	P_FE2CL_PC_STREETSTALL_REP_ITEM_BUY_SUCC_SELLER       = 0x31000114
	P_FE2CL_PC_STREETSTALL_REP_ITEM_BUY_FAIL              = 0x31000115
	P_FE2CL_REP_PC_ITEM_COMBINATION_SUCC                  = 0x31000116
	P_FE2CL_REP_PC_ITEM_COMBINATION_FAIL                  = 0x31000117
	P_FE2CL_PC_CASH_BUFF_UPDATE                           = 0x31000118
	P_FE2CL_REP_PC_SKILL_ADD_SUCC                         = 0x31000119
	P_FE2CL_REP_PC_SKILL_ADD_FAIL                         = 0x3100011a
	P_FE2CL_REP_PC_SKILL_DEL_SUCC                         = 0x3100011b
	P_FE2CL_REP_PC_SKILL_DEL_FAIL                         = 0x3100011c
	P_FE2CL_REP_PC_SKILL_USE_SUCC                         = 0x3100011d
	P_FE2CL_REP_PC_SKILL_USE_FAIL                         = 0x3100011e
	P_FE2CL_PC_SKILL_USE                                  = 0x3100011f
	P_FE2CL_PC_ROPE                                       = 0x31000120
	P_FE2CL_PC_BELT                                       = 0x31000121
	P_FE2CL_PC_VEHICLE_ON_SUCC                            = 0x31000122
	P_FE2CL_PC_VEHICLE_ON_FAIL                            = 0x31000123
	P_FE2CL_PC_VEHICLE_OFF_SUCC                           = 0x31000124
	P_FE2CL_PC_VEHICLE_OFF_FAIL                           = 0x31000125
	P_FE2CL_PC_QUICK_SLOT_INFO                            = 0x31000126
	P_FE2CL_REP_PC_REGIST_QUICK_SLOT_FAIL                 = 0x31000127
	P_FE2CL_REP_PC_REGIST_QUICK_SLOT_SUCC                 = 0x31000128
	P_FE2CL_PC_DELETE_TIME_LIMIT_ITEM                     = 0x31000129
	P_FE2CL_REP_PC_DISASSEMBLE_ITEM_SUCC                  = 0x3100012a
	P_FE2CL_REP_PC_DISASSEMBLE_ITEM_FAIL                  = 0x3100012b
	P_FE2CL_GM_REP_REWARD_RATE_SUCC                       = 0x3100012c
	P_FE2CL_REP_PC_ITEM_ENCHANT_SUCC                      = 0x3100012d
	P_FE2CL_REP_PC_ITEM_ENCHANT_FAIL                      = 0x3100012e
	P_LS2CL_REP_LOGIN_SUCC                                = 0x21000001
	P_LS2CL_REP_LOGIN_FAIL                                = 0x21000002
	P_LS2CL_REP_CHAR_INFO                                 = 0x21000003
	P_LS2CL_REP_CHECK_CHAR_NAME_SUCC                      = 0x21000005
	P_LS2CL_REP_CHECK_CHAR_NAME_FAIL                      = 0x21000006
	P_LS2CL_REP_SAVE_CHAR_NAME_SUCC                       = 0x21000007
	P_LS2CL_REP_SAVE_CHAR_NAME_FAIL                       = 0x21000008
	P_LS2CL_REP_CHAR_CREATE_SUCC                          = 0x21000009
	P_LS2CL_REP_CHAR_CREATE_FAIL                          = 0x2100000a
	P_LS2CL_REP_CHAR_SELECT_SUCC                          = 0x2100000b
	P_LS2CL_REP_CHAR_SELECT_FAIL                          = 0x2100000c
	P_LS2CL_REP_CHAR_DELETE_SUCC                          = 0x2100000d
	P_LS2CL_REP_CHAR_DELETE_FAIL                          = 0x2100000e
	P_LS2CL_REP_SHARD_SELECT_SUCC                         = 0x2100000f
	P_LS2CL_REP_SHARD_SELECT_FAIL                         = 0x21000010
	P_LS2CL_REP_VERSION_CHECK_SUCC                        = 0x21000011
	P_LS2CL_REP_VERSION_CHECK_FAIL                        = 0x21000012
	P_LS2CL_REP_CHECK_NAME_LIST_SUCC                      = 0x21000013
	P_LS2CL_REP_CHECK_NAME_LIST_FAIL                      = 0x21000014
	P_LS2CL_REP_PC_EXIT_DUPLICATE                         = 0x21000015
	P_LS2CL_REQ_LIVE_CHECK                                = 0x21000016
	P_LS2CL_REP_CHANGE_CHAR_NAME_SUCC                     = 0x21000017
	P_LS2CL_REP_CHANGE_CHAR_NAME_FAIL                     = 0x21000018
	P_LS2CL_REP_SHARD_LIST_INFO_SUCC                      = 0x21000019
)

func PacketIDToString(id uint32) string {
	switch id {
	case P_CL2LS_REQ_LOGIN:
		return "P_CL2LS_REQ_LOGIN"
	case P_CL2LS_REQ_CHECK_CHAR_NAME:
		return "P_CL2LS_REQ_CHECK_CHAR_NAME"
	case P_CL2LS_REQ_SAVE_CHAR_NAME:
		return "P_CL2LS_REQ_SAVE_CHAR_NAME"
	case P_CL2LS_REQ_CHAR_CREATE:
		return "P_CL2LS_REQ_CHAR_CREATE"
	case P_CL2LS_REQ_CHAR_SELECT:
		return "P_CL2LS_REQ_CHAR_SELECT"
	case P_CL2LS_REQ_CHAR_DELETE:
		return "P_CL2LS_REQ_CHAR_DELETE"
	case P_CL2LS_REQ_SHARD_SELECT:
		return "P_CL2LS_REQ_SHARD_SELECT"
	case P_CL2LS_REQ_SHARD_LIST_INFO:
		return "P_CL2LS_REQ_SHARD_LIST_INFO"
	case P_CL2LS_CHECK_NAME_LIST:
		return "P_CL2LS_CHECK_NAME_LIST"
	case P_CL2LS_REQ_SAVE_CHAR_TUTOR:
		return "P_CL2LS_REQ_SAVE_CHAR_TUTOR"
	case P_CL2LS_REQ_PC_EXIT_DUPLICATE:
		return "P_CL2LS_REQ_PC_EXIT_DUPLICATE"
	case P_CL2LS_REP_LIVE_CHECK:
		return "P_CL2LS_REP_LIVE_CHECK"
	case P_CL2LS_REQ_CHANGE_CHAR_NAME:
		return "P_CL2LS_REQ_CHANGE_CHAR_NAME"
	case P_CL2LS_REQ_SERVER_SELECT:
		return "P_CL2LS_REQ_SERVER_SELECT"
	case P_CL2FE_REQ_PC_ENTER:
		return "P_CL2FE_REQ_PC_ENTER"
	case P_CL2FE_REQ_PC_EXIT:
		return "P_CL2FE_REQ_PC_EXIT"
	case P_CL2FE_REQ_PC_MOVE:
		return "P_CL2FE_REQ_PC_MOVE"
	case P_CL2FE_REQ_PC_STOP:
		return "P_CL2FE_REQ_PC_STOP"
	case P_CL2FE_REQ_PC_JUMP:
		return "P_CL2FE_REQ_PC_JUMP"
	case P_CL2FE_REQ_PC_ATTACK_NPCs:
		return "P_CL2FE_REQ_PC_ATTACK_NPCs"
	case P_CL2FE_REQ_SEND_FREECHAT_MESSAGE:
		return "P_CL2FE_REQ_SEND_FREECHAT_MESSAGE"
	case P_CL2FE_REQ_SEND_MENUCHAT_MESSAGE:
		return "P_CL2FE_REQ_SEND_MENUCHAT_MESSAGE"
	case P_CL2FE_REQ_PC_REGEN:
		return "P_CL2FE_REQ_PC_REGEN"
	case P_CL2FE_REQ_ITEM_MOVE:
		return "P_CL2FE_REQ_ITEM_MOVE"
	case P_CL2FE_REQ_PC_TASK_START:
		return "P_CL2FE_REQ_PC_TASK_START"
	case P_CL2FE_REQ_PC_TASK_END:
		return "P_CL2FE_REQ_PC_TASK_END"
	case P_CL2FE_REQ_NANO_EQUIP:
		return "P_CL2FE_REQ_NANO_EQUIP"
	case P_CL2FE_REQ_NANO_UNEQUIP:
		return "P_CL2FE_REQ_NANO_UNEQUIP"
	case P_CL2FE_REQ_NANO_ACTIVE:
		return "P_CL2FE_REQ_NANO_ACTIVE"
	case P_CL2FE_REQ_NANO_TUNE:
		return "P_CL2FE_REQ_NANO_TUNE"
	case P_CL2FE_REQ_NANO_SKILL_USE:
		return "P_CL2FE_REQ_NANO_SKILL_USE"
	case P_CL2FE_REQ_PC_TASK_STOP:
		return "P_CL2FE_REQ_PC_TASK_STOP"
	case P_CL2FE_REQ_PC_TASK_CONTINUE:
		return "P_CL2FE_REQ_PC_TASK_CONTINUE"
	case P_CL2FE_REQ_PC_GOTO:
		return "P_CL2FE_REQ_PC_GOTO"
	case P_CL2FE_REQ_CHARGE_NANO_STAMINA:
		return "P_CL2FE_REQ_CHARGE_NANO_STAMINA"
	case P_CL2FE_REQ_PC_KILL_QUEST_NPCs:
		return "P_CL2FE_REQ_PC_KILL_QUEST_NPCs"
	case P_CL2FE_REQ_PC_VENDOR_ITEM_BUY:
		return "P_CL2FE_REQ_PC_VENDOR_ITEM_BUY"
	case P_CL2FE_REQ_PC_VENDOR_ITEM_SELL:
		return "P_CL2FE_REQ_PC_VENDOR_ITEM_SELL"
	case P_CL2FE_REQ_PC_ITEM_DELETE:
		return "P_CL2FE_REQ_PC_ITEM_DELETE"
	case P_CL2FE_REQ_PC_GIVE_ITEM:
		return "P_CL2FE_REQ_PC_GIVE_ITEM"
	case P_CL2FE_REQ_PC_ROCKET_STYLE_READY:
		return "P_CL2FE_REQ_PC_ROCKET_STYLE_READY"
	case P_CL2FE_REQ_PC_ROCKET_STYLE_FIRE:
		return "P_CL2FE_REQ_PC_ROCKET_STYLE_FIRE"
	case P_CL2FE_REQ_PC_ROCKET_STYLE_HIT:
		return "P_CL2FE_REQ_PC_ROCKET_STYLE_HIT"
	case P_CL2FE_REQ_PC_GRENADE_STYLE_READY:
		return "P_CL2FE_REQ_PC_GRENADE_STYLE_READY"
	case P_CL2FE_REQ_PC_GRENADE_STYLE_FIRE:
		return "P_CL2FE_REQ_PC_GRENADE_STYLE_FIRE"
	case P_CL2FE_REQ_PC_GRENADE_STYLE_HIT:
		return "P_CL2FE_REQ_PC_GRENADE_STYLE_HIT"
	case P_CL2FE_REQ_PC_NANO_CREATE:
		return "P_CL2FE_REQ_PC_NANO_CREATE"
	case P_CL2FE_REQ_PC_TRADE_OFFER:
		return "P_CL2FE_REQ_PC_TRADE_OFFER"
	case P_CL2FE_REQ_PC_TRADE_OFFER_CANCEL:
		return "P_CL2FE_REQ_PC_TRADE_OFFER_CANCEL"
	case P_CL2FE_REQ_PC_TRADE_OFFER_ACCEPT:
		return "P_CL2FE_REQ_PC_TRADE_OFFER_ACCEPT"
	case P_CL2FE_REQ_PC_TRADE_OFFER_REFUSAL:
		return "P_CL2FE_REQ_PC_TRADE_OFFER_REFUSAL"
	case P_CL2FE_REQ_PC_TRADE_OFFER_ABORT:
		return "P_CL2FE_REQ_PC_TRADE_OFFER_ABORT"
	case P_CL2FE_REQ_PC_TRADE_CONFIRM:
		return "P_CL2FE_REQ_PC_TRADE_CONFIRM"
	case P_CL2FE_REQ_PC_TRADE_CONFIRM_CANCEL:
		return "P_CL2FE_REQ_PC_TRADE_CONFIRM_CANCEL"
	case P_CL2FE_REQ_PC_TRADE_CONFIRM_ABORT:
		return "P_CL2FE_REQ_PC_TRADE_CONFIRM_ABORT"
	case P_CL2FE_REQ_PC_TRADE_ITEM_REGISTER:
		return "P_CL2FE_REQ_PC_TRADE_ITEM_REGISTER"
	case P_CL2FE_REQ_PC_TRADE_ITEM_UNREGISTER:
		return "P_CL2FE_REQ_PC_TRADE_ITEM_UNREGISTER"
	case P_CL2FE_REQ_PC_TRADE_CASH_REGISTER:
		return "P_CL2FE_REQ_PC_TRADE_CASH_REGISTER"
	case P_CL2FE_REQ_PC_TRADE_EMOTES_CHAT:
		return "P_CL2FE_REQ_PC_TRADE_EMOTES_CHAT"
	case P_CL2FE_REQ_PC_BANK_OPEN:
		return "P_CL2FE_REQ_PC_BANK_OPEN"
	case P_CL2FE_REQ_PC_BANK_CLOSE:
		return "P_CL2FE_REQ_PC_BANK_CLOSE"
	case P_CL2FE_REQ_PC_VENDOR_START:
		return "P_CL2FE_REQ_PC_VENDOR_START"
	case P_CL2FE_REQ_PC_VENDOR_TABLE_UPDATE:
		return "P_CL2FE_REQ_PC_VENDOR_TABLE_UPDATE"
	case P_CL2FE_REQ_PC_VENDOR_ITEM_RESTORE_BUY:
		return "P_CL2FE_REQ_PC_VENDOR_ITEM_RESTORE_BUY"
	case P_CL2FE_REQ_PC_COMBAT_BEGIN:
		return "P_CL2FE_REQ_PC_COMBAT_BEGIN"
	case P_CL2FE_REQ_PC_COMBAT_END:
		return "P_CL2FE_REQ_PC_COMBAT_END"
	case P_CL2FE_REQ_REQUEST_MAKE_BUDDY:
		return "P_CL2FE_REQ_REQUEST_MAKE_BUDDY"
	case P_CL2FE_REQ_ACCEPT_MAKE_BUDDY:
		return "P_CL2FE_REQ_ACCEPT_MAKE_BUDDY"
	case P_CL2FE_REQ_SEND_BUDDY_FREECHAT_MESSAGE:
		return "P_CL2FE_REQ_SEND_BUDDY_FREECHAT_MESSAGE"
	case P_CL2FE_REQ_SEND_BUDDY_MENUCHAT_MESSAGE:
		return "P_CL2FE_REQ_SEND_BUDDY_MENUCHAT_MESSAGE"
	case P_CL2FE_REQ_GET_BUDDY_STYLE:
		return "P_CL2FE_REQ_GET_BUDDY_STYLE"
	case P_CL2FE_REQ_SET_BUDDY_BLOCK:
		return "P_CL2FE_REQ_SET_BUDDY_BLOCK"
	case P_CL2FE_REQ_REMOVE_BUDDY:
		return "P_CL2FE_REQ_REMOVE_BUDDY"
	case P_CL2FE_REQ_GET_BUDDY_STATE:
		return "P_CL2FE_REQ_GET_BUDDY_STATE"
	case P_CL2FE_REQ_PC_JUMPPAD:
		return "P_CL2FE_REQ_PC_JUMPPAD"
	case P_CL2FE_REQ_PC_LAUNCHER:
		return "P_CL2FE_REQ_PC_LAUNCHER"
	case P_CL2FE_REQ_PC_ZIPLINE:
		return "P_CL2FE_REQ_PC_ZIPLINE"
	case P_CL2FE_REQ_PC_MOVEPLATFORM:
		return "P_CL2FE_REQ_PC_MOVEPLATFORM"
	case P_CL2FE_REQ_PC_SLOPE:
		return "P_CL2FE_REQ_PC_SLOPE"
	case P_CL2FE_REQ_PC_STATE_CHANGE:
		return "P_CL2FE_REQ_PC_STATE_CHANGE"
	case P_CL2FE_REQ_PC_MAP_WARP:
		return "P_CL2FE_REQ_PC_MAP_WARP"
	case P_CL2FE_REQ_PC_GIVE_NANO:
		return "P_CL2FE_REQ_PC_GIVE_NANO"
	case P_CL2FE_REQ_NPC_SUMMON:
		return "P_CL2FE_REQ_NPC_SUMMON"
	case P_CL2FE_REQ_NPC_UNSUMMON:
		return "P_CL2FE_REQ_NPC_UNSUMMON"
	case P_CL2FE_REQ_ITEM_CHEST_OPEN:
		return "P_CL2FE_REQ_ITEM_CHEST_OPEN"
	case P_CL2FE_REQ_PC_GIVE_NANO_SKILL:
		return "P_CL2FE_REQ_PC_GIVE_NANO_SKILL"
	case P_CL2FE_DOT_DAMAGE_ONOFF:
		return "P_CL2FE_DOT_DAMAGE_ONOFF"
	case P_CL2FE_REQ_PC_VENDOR_BATTERY_BUY:
		return "P_CL2FE_REQ_PC_VENDOR_BATTERY_BUY"
	case P_CL2FE_REQ_PC_WARP_USE_NPC:
		return "P_CL2FE_REQ_PC_WARP_USE_NPC"
	case P_CL2FE_REQ_PC_GROUP_INVITE:
		return "P_CL2FE_REQ_PC_GROUP_INVITE"
	case P_CL2FE_REQ_PC_GROUP_INVITE_REFUSE:
		return "P_CL2FE_REQ_PC_GROUP_INVITE_REFUSE"
	case P_CL2FE_REQ_PC_GROUP_JOIN:
		return "P_CL2FE_REQ_PC_GROUP_JOIN"
	case P_CL2FE_REQ_PC_GROUP_LEAVE:
		return "P_CL2FE_REQ_PC_GROUP_LEAVE"
	case P_CL2FE_REQ_PC_AVATAR_EMOTES_CHAT:
		return "P_CL2FE_REQ_PC_AVATAR_EMOTES_CHAT"
	case P_CL2FE_REQ_PC_BUDDY_WARP:
		return "P_CL2FE_REQ_PC_BUDDY_WARP"
	case P_CL2FE_REQ_GET_MEMBER_STYLE:
		return "P_CL2FE_REQ_GET_MEMBER_STYLE"
	case P_CL2FE_REQ_GET_GROUP_STYLE:
		return "P_CL2FE_REQ_GET_GROUP_STYLE"
	case P_CL2FE_REQ_PC_CHANGE_MENTOR:
		return "P_CL2FE_REQ_PC_CHANGE_MENTOR"
	case P_CL2FE_REQ_GET_BUDDY_LOCATION:
		return "P_CL2FE_REQ_GET_BUDDY_LOCATION"
	case P_CL2FE_REQ_NPC_GROUP_SUMMON:
		return "P_CL2FE_REQ_NPC_GROUP_SUMMON"
	case P_CL2FE_REQ_PC_WARP_TO_PC:
		return "P_CL2FE_REQ_PC_WARP_TO_PC"
	case P_CL2FE_REQ_EP_RANK_GET_LIST:
		return "P_CL2FE_REQ_EP_RANK_GET_LIST"
	case P_CL2FE_REQ_EP_RANK_GET_DETAIL:
		return "P_CL2FE_REQ_EP_RANK_GET_DETAIL"
	case P_CL2FE_REQ_EP_RANK_GET_PC_INFO:
		return "P_CL2FE_REQ_EP_RANK_GET_PC_INFO"
	case P_CL2FE_REQ_EP_RACE_START:
		return "P_CL2FE_REQ_EP_RACE_START"
	case P_CL2FE_REQ_EP_RACE_END:
		return "P_CL2FE_REQ_EP_RACE_END"
	case P_CL2FE_REQ_EP_RACE_CANCEL:
		return "P_CL2FE_REQ_EP_RACE_CANCEL"
	case P_CL2FE_REQ_EP_GET_RING:
		return "P_CL2FE_REQ_EP_GET_RING"
	case P_CL2FE_REQ_IM_CHANGE_SWITCH_STATUS:
		return "P_CL2FE_REQ_IM_CHANGE_SWITCH_STATUS"
	case P_CL2FE_REQ_SHINY_PICKUP:
		return "P_CL2FE_REQ_SHINY_PICKUP"
	case P_CL2FE_REQ_SHINY_SUMMON:
		return "P_CL2FE_REQ_SHINY_SUMMON"
	case P_CL2FE_REQ_PC_MOVETRANSPORTATION:
		return "P_CL2FE_REQ_PC_MOVETRANSPORTATION"
	case P_CL2FE_REQ_SEND_ALL_GROUP_FREECHAT_MESSAGE:
		return "P_CL2FE_REQ_SEND_ALL_GROUP_FREECHAT_MESSAGE"
	case P_CL2FE_REQ_SEND_ANY_GROUP_FREECHAT_MESSAGE:
		return "P_CL2FE_REQ_SEND_ANY_GROUP_FREECHAT_MESSAGE"
	case P_CL2FE_REQ_BARKER:
		return "P_CL2FE_REQ_BARKER"
	case P_CL2FE_REQ_SEND_ALL_GROUP_MENUCHAT_MESSAGE:
		return "P_CL2FE_REQ_SEND_ALL_GROUP_MENUCHAT_MESSAGE"
	case P_CL2FE_REQ_SEND_ANY_GROUP_MENUCHAT_MESSAGE:
		return "P_CL2FE_REQ_SEND_ANY_GROUP_MENUCHAT_MESSAGE"
	case P_CL2FE_REQ_REGIST_TRANSPORTATION_LOCATION:
		return "P_CL2FE_REQ_REGIST_TRANSPORTATION_LOCATION"
	case P_CL2FE_REQ_PC_WARP_USE_TRANSPORTATION:
		return "P_CL2FE_REQ_PC_WARP_USE_TRANSPORTATION"
	case P_CL2FE_GM_REQ_PC_SPECIAL_STATE_SWITCH:
		return "P_CL2FE_GM_REQ_PC_SPECIAL_STATE_SWITCH"
	case P_CL2FE_GM_REQ_PC_SET_VALUE:
		return "P_CL2FE_GM_REQ_PC_SET_VALUE"
	case P_CL2FE_GM_REQ_KICK_PLAYER:
		return "P_CL2FE_GM_REQ_KICK_PLAYER"
	case P_CL2FE_GM_REQ_TARGET_PC_TELEPORT:
		return "P_CL2FE_GM_REQ_TARGET_PC_TELEPORT"
	case P_CL2FE_GM_REQ_PC_LOCATION:
		return "P_CL2FE_GM_REQ_PC_LOCATION"
	case P_CL2FE_GM_REQ_PC_ANNOUNCE:
		return "P_CL2FE_GM_REQ_PC_ANNOUNCE"
	case P_CL2FE_REQ_SET_PC_BLOCK:
		return "P_CL2FE_REQ_SET_PC_BLOCK"
	case P_CL2FE_REQ_REGIST_RXCOM:
		return "P_CL2FE_REQ_REGIST_RXCOM"
	case P_CL2FE_GM_REQ_PC_MOTD_REGISTER:
		return "P_CL2FE_GM_REQ_PC_MOTD_REGISTER"
	case P_CL2FE_REQ_ITEM_USE:
		return "P_CL2FE_REQ_ITEM_USE"
	case P_CL2FE_REQ_WARP_USE_RECALL:
		return "P_CL2FE_REQ_WARP_USE_RECALL"
	case P_CL2FE_REP_LIVE_CHECK:
		return "P_CL2FE_REP_LIVE_CHECK"
	case P_CL2FE_REQ_PC_MISSION_COMPLETE:
		return "P_CL2FE_REQ_PC_MISSION_COMPLETE"
	case P_CL2FE_REQ_PC_TASK_COMPLETE:
		return "P_CL2FE_REQ_PC_TASK_COMPLETE"
	case P_CL2FE_REQ_NPC_INTERACTION:
		return "P_CL2FE_REQ_NPC_INTERACTION"
	case P_CL2FE_DOT_HEAL_ONOFF:
		return "P_CL2FE_DOT_HEAL_ONOFF"
	case P_CL2FE_REQ_PC_SPECIAL_STATE_SWITCH:
		return "P_CL2FE_REQ_PC_SPECIAL_STATE_SWITCH"
	case P_CL2FE_REQ_PC_EMAIL_UPDATE_CHECK:
		return "P_CL2FE_REQ_PC_EMAIL_UPDATE_CHECK"
	case P_CL2FE_REQ_PC_READ_EMAIL:
		return "P_CL2FE_REQ_PC_READ_EMAIL"
	case P_CL2FE_REQ_PC_RECV_EMAIL_PAGE_LIST:
		return "P_CL2FE_REQ_PC_RECV_EMAIL_PAGE_LIST"
	case P_CL2FE_REQ_PC_DELETE_EMAIL:
		return "P_CL2FE_REQ_PC_DELETE_EMAIL"
	case P_CL2FE_REQ_PC_SEND_EMAIL:
		return "P_CL2FE_REQ_PC_SEND_EMAIL"
	case P_CL2FE_REQ_PC_RECV_EMAIL_ITEM:
		return "P_CL2FE_REQ_PC_RECV_EMAIL_ITEM"
	case P_CL2FE_REQ_PC_RECV_EMAIL_CANDY:
		return "P_CL2FE_REQ_PC_RECV_EMAIL_CANDY"
	case P_CL2FE_GM_REQ_TARGET_PC_SPECIAL_STATE_ONOFF:
		return "P_CL2FE_GM_REQ_TARGET_PC_SPECIAL_STATE_ONOFF"
	case P_CL2FE_REQ_PC_SET_CURRENT_MISSION_ID:
		return "P_CL2FE_REQ_PC_SET_CURRENT_MISSION_ID"
	case P_CL2FE_REQ_NPC_GROUP_INVITE:
		return "P_CL2FE_REQ_NPC_GROUP_INVITE"
	case P_CL2FE_REQ_NPC_GROUP_KICK:
		return "P_CL2FE_REQ_NPC_GROUP_KICK"
	case P_CL2FE_REQ_PC_FIRST_USE_FLAG_SET:
		return "P_CL2FE_REQ_PC_FIRST_USE_FLAG_SET"
	case P_CL2FE_REQ_PC_TRANSPORT_WARP:
		return "P_CL2FE_REQ_PC_TRANSPORT_WARP"
	case P_CL2FE_REQ_PC_TIME_TO_GO_WARP:
		return "P_CL2FE_REQ_PC_TIME_TO_GO_WARP"
	case P_CL2FE_REQ_PC_RECV_EMAIL_ITEM_ALL:
		return "P_CL2FE_REQ_PC_RECV_EMAIL_ITEM_ALL"
	case P_CL2FE_REQ_CHANNEL_INFO:
		return "P_CL2FE_REQ_CHANNEL_INFO"
	case P_CL2FE_REQ_PC_CHANNEL_NUM:
		return "P_CL2FE_REQ_PC_CHANNEL_NUM"
	case P_CL2FE_REQ_PC_WARP_CHANNEL:
		return "P_CL2FE_REQ_PC_WARP_CHANNEL"
	case P_CL2FE_REQ_PC_LOADING_COMPLETE:
		return "P_CL2FE_REQ_PC_LOADING_COMPLETE"
	case P_CL2FE_REQ_PC_FIND_NAME_MAKE_BUDDY:
		return "P_CL2FE_REQ_PC_FIND_NAME_MAKE_BUDDY"
	case P_CL2FE_REQ_PC_FIND_NAME_ACCEPT_BUDDY:
		return "P_CL2FE_REQ_PC_FIND_NAME_ACCEPT_BUDDY"
	case P_CL2FE_REQ_PC_ATTACK_CHARs:
		return "P_CL2FE_REQ_PC_ATTACK_CHARs"
	case P_CL2FE_PC_STREETSTALL_REQ_READY:
		return "P_CL2FE_PC_STREETSTALL_REQ_READY"
	case P_CL2FE_PC_STREETSTALL_REQ_CANCEL:
		return "P_CL2FE_PC_STREETSTALL_REQ_CANCEL"
	case P_CL2FE_PC_STREETSTALL_REQ_REGIST_ITEM:
		return "P_CL2FE_PC_STREETSTALL_REQ_REGIST_ITEM"
	case P_CL2FE_PC_STREETSTALL_REQ_UNREGIST_ITEM:
		return "P_CL2FE_PC_STREETSTALL_REQ_UNREGIST_ITEM"
	case P_CL2FE_PC_STREETSTALL_REQ_SALE_START:
		return "P_CL2FE_PC_STREETSTALL_REQ_SALE_START"
	case P_CL2FE_PC_STREETSTALL_REQ_ITEM_LIST:
		return "P_CL2FE_PC_STREETSTALL_REQ_ITEM_LIST"
	case P_CL2FE_PC_STREETSTALL_REQ_ITEM_BUY:
		return "P_CL2FE_PC_STREETSTALL_REQ_ITEM_BUY"
	case P_CL2FE_REQ_PC_ITEM_COMBINATION:
		return "P_CL2FE_REQ_PC_ITEM_COMBINATION"
	case P_CL2FE_GM_REQ_SET_PC_SKILL:
		return "P_CL2FE_GM_REQ_SET_PC_SKILL"
	case P_CL2FE_REQ_PC_SKILL_ADD:
		return "P_CL2FE_REQ_PC_SKILL_ADD"
	case P_CL2FE_REQ_PC_SKILL_DEL:
		return "P_CL2FE_REQ_PC_SKILL_DEL"
	case P_CL2FE_REQ_PC_SKILL_USE:
		return "P_CL2FE_REQ_PC_SKILL_USE"
	case P_CL2FE_REQ_PC_ROPE:
		return "P_CL2FE_REQ_PC_ROPE"
	case P_CL2FE_REQ_PC_BELT:
		return "P_CL2FE_REQ_PC_BELT"
	case P_CL2FE_REQ_PC_VEHICLE_ON:
		return "P_CL2FE_REQ_PC_VEHICLE_ON"
	case P_CL2FE_REQ_PC_VEHICLE_OFF:
		return "P_CL2FE_REQ_PC_VEHICLE_OFF"
	case P_CL2FE_REQ_PC_REGIST_QUICK_SLOT:
		return "P_CL2FE_REQ_PC_REGIST_QUICK_SLOT"
	case P_CL2FE_REQ_PC_DISASSEMBLE_ITEM:
		return "P_CL2FE_REQ_PC_DISASSEMBLE_ITEM"
	case P_CL2FE_GM_REQ_REWARD_RATE:
		return "P_CL2FE_GM_REQ_REWARD_RATE"
	case P_CL2FE_REQ_PC_ITEM_ENCHANT:
		return "P_CL2FE_REQ_PC_ITEM_ENCHANT"
	case P_FE2CL_ERROR:
		return "P_FE2CL_ERROR"
	case P_FE2CL_REP_PC_ENTER_FAIL:
		return "P_FE2CL_REP_PC_ENTER_FAIL"
	case P_FE2CL_REP_PC_ENTER_SUCC:
		return "P_FE2CL_REP_PC_ENTER_SUCC"
	case P_FE2CL_PC_NEW:
		return "P_FE2CL_PC_NEW"
	case P_FE2CL_REP_PC_EXIT_FAIL:
		return "P_FE2CL_REP_PC_EXIT_FAIL"
	case P_FE2CL_REP_PC_EXIT_SUCC:
		return "P_FE2CL_REP_PC_EXIT_SUCC"
	case P_FE2CL_PC_EXIT:
		return "P_FE2CL_PC_EXIT"
	case P_FE2CL_PC_AROUND:
		return "P_FE2CL_PC_AROUND"
	case P_FE2CL_PC_MOVE:
		return "P_FE2CL_PC_MOVE"
	case P_FE2CL_PC_STOP:
		return "P_FE2CL_PC_STOP"
	case P_FE2CL_PC_JUMP:
		return "P_FE2CL_PC_JUMP"
	case P_FE2CL_NPC_ENTER:
		return "P_FE2CL_NPC_ENTER"
	case P_FE2CL_NPC_EXIT:
		return "P_FE2CL_NPC_EXIT"
	case P_FE2CL_NPC_MOVE:
		return "P_FE2CL_NPC_MOVE"
	case P_FE2CL_NPC_NEW:
		return "P_FE2CL_NPC_NEW"
	case P_FE2CL_NPC_AROUND:
		return "P_FE2CL_NPC_AROUND"
	case P_FE2CL_AROUND_DEL_PC:
		return "P_FE2CL_AROUND_DEL_PC"
	case P_FE2CL_AROUND_DEL_NPC:
		return "P_FE2CL_AROUND_DEL_NPC"
	case P_FE2CL_REP_SEND_FREECHAT_MESSAGE_SUCC:
		return "P_FE2CL_REP_SEND_FREECHAT_MESSAGE_SUCC"
	case P_FE2CL_REP_SEND_FREECHAT_MESSAGE_FAIL:
		return "P_FE2CL_REP_SEND_FREECHAT_MESSAGE_FAIL"
	case P_FE2CL_PC_ATTACK_NPCs_SUCC:
		return "P_FE2CL_PC_ATTACK_NPCs_SUCC"
	case P_FE2CL_PC_ATTACK_NPCs:
		return "P_FE2CL_PC_ATTACK_NPCs"
	case P_FE2CL_NPC_ATTACK_PCs:
		return "P_FE2CL_NPC_ATTACK_PCs"
	case P_FE2CL_REP_PC_REGEN_SUCC:
		return "P_FE2CL_REP_PC_REGEN_SUCC"
	case P_FE2CL_REP_SEND_MENUCHAT_MESSAGE_SUCC:
		return "P_FE2CL_REP_SEND_MENUCHAT_MESSAGE_SUCC"
	case P_FE2CL_REP_SEND_MENUCHAT_MESSAGE_FAIL:
		return "P_FE2CL_REP_SEND_MENUCHAT_MESSAGE_FAIL"
	case P_FE2CL_PC_ITEM_MOVE_SUCC:
		return "P_FE2CL_PC_ITEM_MOVE_SUCC"
	case P_FE2CL_PC_EQUIP_CHANGE:
		return "P_FE2CL_PC_EQUIP_CHANGE"
	case P_FE2CL_REP_PC_TASK_START_SUCC:
		return "P_FE2CL_REP_PC_TASK_START_SUCC"
	case P_FE2CL_REP_PC_TASK_START_FAIL:
		return "P_FE2CL_REP_PC_TASK_START_FAIL"
	case P_FE2CL_REP_PC_TASK_END_SUCC:
		return "P_FE2CL_REP_PC_TASK_END_SUCC"
	case P_FE2CL_REP_PC_TASK_END_FAIL:
		return "P_FE2CL_REP_PC_TASK_END_FAIL"
	case P_FE2CL_NPC_SKILL_READY:
		return "P_FE2CL_NPC_SKILL_READY"
	case P_FE2CL_NPC_SKILL_FIRE:
		return "P_FE2CL_NPC_SKILL_FIRE"
	case P_FE2CL_NPC_SKILL_HIT:
		return "P_FE2CL_NPC_SKILL_HIT"
	case P_FE2CL_NPC_SKILL_CORRUPTION_READY:
		return "P_FE2CL_NPC_SKILL_CORRUPTION_READY"
	case P_FE2CL_NPC_SKILL_CORRUPTION_HIT:
		return "P_FE2CL_NPC_SKILL_CORRUPTION_HIT"
	case P_FE2CL_NPC_SKILL_CANCEL:
		return "P_FE2CL_NPC_SKILL_CANCEL"
	case P_FE2CL_REP_NANO_EQUIP_SUCC:
		return "P_FE2CL_REP_NANO_EQUIP_SUCC"
	case P_FE2CL_REP_NANO_UNEQUIP_SUCC:
		return "P_FE2CL_REP_NANO_UNEQUIP_SUCC"
	case P_FE2CL_REP_NANO_ACTIVE_SUCC:
		return "P_FE2CL_REP_NANO_ACTIVE_SUCC"
	case P_FE2CL_REP_NANO_TUNE_SUCC:
		return "P_FE2CL_REP_NANO_TUNE_SUCC"
	case P_FE2CL_NANO_ACTIVE:
		return "P_FE2CL_NANO_ACTIVE"
	case P_FE2CL_NANO_SKILL_USE_SUCC:
		return "P_FE2CL_NANO_SKILL_USE_SUCC"
	case P_FE2CL_NANO_SKILL_USE:
		return "P_FE2CL_NANO_SKILL_USE"
	case P_FE2CL_REP_PC_TASK_STOP_SUCC:
		return "P_FE2CL_REP_PC_TASK_STOP_SUCC"
	case P_FE2CL_REP_PC_TASK_STOP_FAIL:
		return "P_FE2CL_REP_PC_TASK_STOP_FAIL"
	case P_FE2CL_REP_PC_TASK_CONTINUE_SUCC:
		return "P_FE2CL_REP_PC_TASK_CONTINUE_SUCC"
	case P_FE2CL_REP_PC_TASK_CONTINUE_FAIL:
		return "P_FE2CL_REP_PC_TASK_CONTINUE_FAIL"
	case P_FE2CL_REP_PC_GOTO_SUCC:
		return "P_FE2CL_REP_PC_GOTO_SUCC"
	case P_FE2CL_REP_CHARGE_NANO_STAMINA:
		return "P_FE2CL_REP_CHARGE_NANO_STAMINA"
	case P_FE2CL_REP_PC_TICK:
		return "P_FE2CL_REP_PC_TICK"
	case P_FE2CL_REP_PC_KILL_QUEST_NPCs_SUCC:
		return "P_FE2CL_REP_PC_KILL_QUEST_NPCs_SUCC"
	case P_FE2CL_REP_PC_VENDOR_ITEM_BUY_SUCC:
		return "P_FE2CL_REP_PC_VENDOR_ITEM_BUY_SUCC"
	case P_FE2CL_REP_PC_VENDOR_ITEM_BUY_FAIL:
		return "P_FE2CL_REP_PC_VENDOR_ITEM_BUY_FAIL"
	case P_FE2CL_REP_PC_VENDOR_ITEM_SELL_SUCC:
		return "P_FE2CL_REP_PC_VENDOR_ITEM_SELL_SUCC"
	case P_FE2CL_REP_PC_VENDOR_ITEM_SELL_FAIL:
		return "P_FE2CL_REP_PC_VENDOR_ITEM_SELL_FAIL"
	case P_FE2CL_REP_PC_ITEM_DELETE_SUCC:
		return "P_FE2CL_REP_PC_ITEM_DELETE_SUCC"
	case P_FE2CL_PC_ROCKET_STYLE_READY:
		return "P_FE2CL_PC_ROCKET_STYLE_READY"
	case P_FE2CL_REP_PC_ROCKET_STYLE_FIRE_SUCC:
		return "P_FE2CL_REP_PC_ROCKET_STYLE_FIRE_SUCC"
	case P_FE2CL_PC_ROCKET_STYLE_FIRE:
		return "P_FE2CL_PC_ROCKET_STYLE_FIRE"
	case P_FE2CL_PC_ROCKET_STYLE_HIT:
		return "P_FE2CL_PC_ROCKET_STYLE_HIT"
	case P_FE2CL_PC_GRENADE_STYLE_READY:
		return "P_FE2CL_PC_GRENADE_STYLE_READY"
	case P_FE2CL_REP_PC_GRENADE_STYLE_FIRE_SUCC:
		return "P_FE2CL_REP_PC_GRENADE_STYLE_FIRE_SUCC"
	case P_FE2CL_PC_GRENADE_STYLE_FIRE:
		return "P_FE2CL_PC_GRENADE_STYLE_FIRE"
	case P_FE2CL_PC_GRENADE_STYLE_HIT:
		return "P_FE2CL_PC_GRENADE_STYLE_HIT"
	case P_FE2CL_REP_PC_TRADE_OFFER:
		return "P_FE2CL_REP_PC_TRADE_OFFER"
	case P_FE2CL_REP_PC_TRADE_OFFER_CANCEL:
		return "P_FE2CL_REP_PC_TRADE_OFFER_CANCEL"
	case P_FE2CL_REP_PC_TRADE_OFFER_SUCC:
		return "P_FE2CL_REP_PC_TRADE_OFFER_SUCC"
	case P_FE2CL_REP_PC_TRADE_OFFER_REFUSAL:
		return "P_FE2CL_REP_PC_TRADE_OFFER_REFUSAL"
	case P_FE2CL_REP_PC_TRADE_OFFER_ABORT:
		return "P_FE2CL_REP_PC_TRADE_OFFER_ABORT"
	case P_FE2CL_REP_PC_TRADE_CONFIRM:
		return "P_FE2CL_REP_PC_TRADE_CONFIRM"
	case P_FE2CL_REP_PC_TRADE_CONFIRM_CANCEL:
		return "P_FE2CL_REP_PC_TRADE_CONFIRM_CANCEL"
	case P_FE2CL_REP_PC_TRADE_CONFIRM_ABORT:
		return "P_FE2CL_REP_PC_TRADE_CONFIRM_ABORT"
	case P_FE2CL_REP_PC_TRADE_CONFIRM_SUCC:
		return "P_FE2CL_REP_PC_TRADE_CONFIRM_SUCC"
	case P_FE2CL_REP_PC_TRADE_CONFIRM_FAIL:
		return "P_FE2CL_REP_PC_TRADE_CONFIRM_FAIL"
	case P_FE2CL_REP_PC_TRADE_ITEM_REGISTER_SUCC:
		return "P_FE2CL_REP_PC_TRADE_ITEM_REGISTER_SUCC"
	case P_FE2CL_REP_PC_TRADE_ITEM_REGISTER_FAIL:
		return "P_FE2CL_REP_PC_TRADE_ITEM_REGISTER_FAIL"
	case P_FE2CL_REP_PC_TRADE_ITEM_UNREGISTER_SUCC:
		return "P_FE2CL_REP_PC_TRADE_ITEM_UNREGISTER_SUCC"
	case P_FE2CL_REP_PC_TRADE_ITEM_UNREGISTER_FAIL:
		return "P_FE2CL_REP_PC_TRADE_ITEM_UNREGISTER_FAIL"
	case P_FE2CL_REP_PC_TRADE_CASH_REGISTER_SUCC:
		return "P_FE2CL_REP_PC_TRADE_CASH_REGISTER_SUCC"
	case P_FE2CL_REP_PC_TRADE_CASH_REGISTER_FAIL:
		return "P_FE2CL_REP_PC_TRADE_CASH_REGISTER_FAIL"
	case P_FE2CL_REP_PC_TRADE_EMOTES_CHAT:
		return "P_FE2CL_REP_PC_TRADE_EMOTES_CHAT"
	case P_FE2CL_REP_PC_NANO_CREATE_SUCC:
		return "P_FE2CL_REP_PC_NANO_CREATE_SUCC"
	case P_FE2CL_REP_PC_NANO_CREATE_FAIL:
		return "P_FE2CL_REP_PC_NANO_CREATE_FAIL"
	case P_FE2CL_REP_NANO_TUNE_FAIL:
		return "P_FE2CL_REP_NANO_TUNE_FAIL"
	case P_FE2CL_REP_PC_BANK_OPEN_SUCC:
		return "P_FE2CL_REP_PC_BANK_OPEN_SUCC"
	case P_FE2CL_REP_PC_BANK_OPEN_FAIL:
		return "P_FE2CL_REP_PC_BANK_OPEN_FAIL"
	case P_FE2CL_REP_PC_BANK_CLOSE_SUCC:
		return "P_FE2CL_REP_PC_BANK_CLOSE_SUCC"
	case P_FE2CL_REP_PC_BANK_CLOSE_FAIL:
		return "P_FE2CL_REP_PC_BANK_CLOSE_FAIL"
	case P_FE2CL_REP_PC_VENDOR_START_SUCC:
		return "P_FE2CL_REP_PC_VENDOR_START_SUCC"
	case P_FE2CL_REP_PC_VENDOR_START_FAIL:
		return "P_FE2CL_REP_PC_VENDOR_START_FAIL"
	case P_FE2CL_REP_PC_VENDOR_TABLE_UPDATE_SUCC:
		return "P_FE2CL_REP_PC_VENDOR_TABLE_UPDATE_SUCC"
	case P_FE2CL_REP_PC_VENDOR_TABLE_UPDATE_FAIL:
		return "P_FE2CL_REP_PC_VENDOR_TABLE_UPDATE_FAIL"
	case P_FE2CL_REP_PC_VENDOR_ITEM_RESTORE_BUY_SUCC:
		return "P_FE2CL_REP_PC_VENDOR_ITEM_RESTORE_BUY_SUCC"
	case P_FE2CL_REP_PC_VENDOR_ITEM_RESTORE_BUY_FAIL:
		return "P_FE2CL_REP_PC_VENDOR_ITEM_RESTORE_BUY_FAIL"
	case P_FE2CL_CHAR_TIME_BUFF_TIME_OUT:
		return "P_FE2CL_CHAR_TIME_BUFF_TIME_OUT"
	case P_FE2CL_REP_PC_GIVE_ITEM_SUCC:
		return "P_FE2CL_REP_PC_GIVE_ITEM_SUCC"
	case P_FE2CL_REP_PC_GIVE_ITEM_FAIL:
		return "P_FE2CL_REP_PC_GIVE_ITEM_FAIL"
	case P_FE2CL_REP_PC_BUDDYLIST_INFO_SUCC:
		return "P_FE2CL_REP_PC_BUDDYLIST_INFO_SUCC"
	case P_FE2CL_REP_PC_BUDDYLIST_INFO_FAIL:
		return "P_FE2CL_REP_PC_BUDDYLIST_INFO_FAIL"
	case P_FE2CL_REP_REQUEST_MAKE_BUDDY_SUCC:
		return "P_FE2CL_REP_REQUEST_MAKE_BUDDY_SUCC"
	case P_FE2CL_REP_REQUEST_MAKE_BUDDY_FAIL:
		return "P_FE2CL_REP_REQUEST_MAKE_BUDDY_FAIL"
	case P_FE2CL_REP_ACCEPT_MAKE_BUDDY_SUCC:
		return "P_FE2CL_REP_ACCEPT_MAKE_BUDDY_SUCC"
	case P_FE2CL_REP_ACCEPT_MAKE_BUDDY_FAIL:
		return "P_FE2CL_REP_ACCEPT_MAKE_BUDDY_FAIL"
	case P_FE2CL_REP_SEND_BUDDY_FREECHAT_MESSAGE_SUCC:
		return "P_FE2CL_REP_SEND_BUDDY_FREECHAT_MESSAGE_SUCC"
	case P_FE2CL_REP_SEND_BUDDY_FREECHAT_MESSAGE_FAIL:
		return "P_FE2CL_REP_SEND_BUDDY_FREECHAT_MESSAGE_FAIL"
	case P_FE2CL_REP_SEND_BUDDY_MENUCHAT_MESSAGE_SUCC:
		return "P_FE2CL_REP_SEND_BUDDY_MENUCHAT_MESSAGE_SUCC"
	case P_FE2CL_REP_SEND_BUDDY_MENUCHAT_MESSAGE_FAIL:
		return "P_FE2CL_REP_SEND_BUDDY_MENUCHAT_MESSAGE_FAIL"
	case P_FE2CL_REP_GET_BUDDY_STYLE_SUCC:
		return "P_FE2CL_REP_GET_BUDDY_STYLE_SUCC"
	case P_FE2CL_REP_GET_BUDDY_STYLE_FAIL:
		return "P_FE2CL_REP_GET_BUDDY_STYLE_FAIL"
	case P_FE2CL_REP_GET_BUDDY_STATE_SUCC:
		return "P_FE2CL_REP_GET_BUDDY_STATE_SUCC"
	case P_FE2CL_REP_GET_BUDDY_STATE_FAIL:
		return "P_FE2CL_REP_GET_BUDDY_STATE_FAIL"
	case P_FE2CL_REP_SET_BUDDY_BLOCK_SUCC:
		return "P_FE2CL_REP_SET_BUDDY_BLOCK_SUCC"
	case P_FE2CL_REP_SET_BUDDY_BLOCK_FAIL:
		return "P_FE2CL_REP_SET_BUDDY_BLOCK_FAIL"
	case P_FE2CL_REP_REMOVE_BUDDY_SUCC:
		return "P_FE2CL_REP_REMOVE_BUDDY_SUCC"
	case P_FE2CL_REP_REMOVE_BUDDY_FAIL:
		return "P_FE2CL_REP_REMOVE_BUDDY_FAIL"
	case P_FE2CL_PC_JUMPPAD:
		return "P_FE2CL_PC_JUMPPAD"
	case P_FE2CL_PC_LAUNCHER:
		return "P_FE2CL_PC_LAUNCHER"
	case P_FE2CL_PC_ZIPLINE:
		return "P_FE2CL_PC_ZIPLINE"
	case P_FE2CL_PC_MOVEPLATFORM:
		return "P_FE2CL_PC_MOVEPLATFORM"
	case P_FE2CL_PC_SLOPE:
		return "P_FE2CL_PC_SLOPE"
	case P_FE2CL_PC_STATE_CHANGE:
		return "P_FE2CL_PC_STATE_CHANGE"
	case P_FE2CL_REP_REQUEST_MAKE_BUDDY_SUCC_TO_ACCEPTER:
		return "P_FE2CL_REP_REQUEST_MAKE_BUDDY_SUCC_TO_ACCEPTER"
	case P_FE2CL_REP_REWARD_ITEM:
		return "P_FE2CL_REP_REWARD_ITEM"
	case P_FE2CL_REP_ITEM_CHEST_OPEN_SUCC:
		return "P_FE2CL_REP_ITEM_CHEST_OPEN_SUCC"
	case P_FE2CL_REP_ITEM_CHEST_OPEN_FAIL:
		return "P_FE2CL_REP_ITEM_CHEST_OPEN_FAIL"
	case P_FE2CL_CHAR_TIME_BUFF_TIME_TICK:
		return "P_FE2CL_CHAR_TIME_BUFF_TIME_TICK"
	case P_FE2CL_REP_PC_VENDOR_BATTERY_BUY_SUCC:
		return "P_FE2CL_REP_PC_VENDOR_BATTERY_BUY_SUCC"
	case P_FE2CL_REP_PC_VENDOR_BATTERY_BUY_FAIL:
		return "P_FE2CL_REP_PC_VENDOR_BATTERY_BUY_FAIL"
	case P_FE2CL_NPC_ROCKET_STYLE_FIRE:
		return "P_FE2CL_NPC_ROCKET_STYLE_FIRE"
	case P_FE2CL_NPC_GRENADE_STYLE_FIRE:
		return "P_FE2CL_NPC_GRENADE_STYLE_FIRE"
	case P_FE2CL_NPC_BULLET_STYLE_HIT:
		return "P_FE2CL_NPC_BULLET_STYLE_HIT"
	case P_FE2CL_CHARACTER_ATTACK_CHARACTERs:
		return "P_FE2CL_CHARACTER_ATTACK_CHARACTERs"
	case P_FE2CL_PC_GROUP_INVITE:
		return "P_FE2CL_PC_GROUP_INVITE"
	case P_FE2CL_PC_GROUP_INVITE_FAIL:
		return "P_FE2CL_PC_GROUP_INVITE_FAIL"
	case P_FE2CL_PC_GROUP_INVITE_REFUSE:
		return "P_FE2CL_PC_GROUP_INVITE_REFUSE"
	case P_FE2CL_PC_GROUP_JOIN:
		return "P_FE2CL_PC_GROUP_JOIN"
	case P_FE2CL_PC_GROUP_JOIN_FAIL:
		return "P_FE2CL_PC_GROUP_JOIN_FAIL"
	case P_FE2CL_PC_GROUP_JOIN_SUCC:
		return "P_FE2CL_PC_GROUP_JOIN_SUCC"
	case P_FE2CL_PC_GROUP_LEAVE:
		return "P_FE2CL_PC_GROUP_LEAVE"
	case P_FE2CL_PC_GROUP_LEAVE_FAIL:
		return "P_FE2CL_PC_GROUP_LEAVE_FAIL"
	case P_FE2CL_PC_GROUP_LEAVE_SUCC:
		return "P_FE2CL_PC_GROUP_LEAVE_SUCC"
	case P_FE2CL_PC_GROUP_MEMBER_INFO:
		return "P_FE2CL_PC_GROUP_MEMBER_INFO"
	case P_FE2CL_REP_PC_WARP_USE_NPC_SUCC:
		return "P_FE2CL_REP_PC_WARP_USE_NPC_SUCC"
	case P_FE2CL_REP_PC_WARP_USE_NPC_FAIL:
		return "P_FE2CL_REP_PC_WARP_USE_NPC_FAIL"
	case P_FE2CL_REP_PC_AVATAR_EMOTES_CHAT:
		return "P_FE2CL_REP_PC_AVATAR_EMOTES_CHAT"
	case P_FE2CL_REP_PC_CHANGE_MENTOR_SUCC:
		return "P_FE2CL_REP_PC_CHANGE_MENTOR_SUCC"
	case P_FE2CL_REP_PC_CHANGE_MENTOR_FAIL:
		return "P_FE2CL_REP_PC_CHANGE_MENTOR_FAIL"
	case P_FE2CL_REP_GET_MEMBER_STYLE_FAIL:
		return "P_FE2CL_REP_GET_MEMBER_STYLE_FAIL"
	case P_FE2CL_REP_GET_MEMBER_STYLE_SUCC:
		return "P_FE2CL_REP_GET_MEMBER_STYLE_SUCC"
	case P_FE2CL_REP_GET_GROUP_STYLE_FAIL:
		return "P_FE2CL_REP_GET_GROUP_STYLE_FAIL"
	case P_FE2CL_REP_GET_GROUP_STYLE_SUCC:
		return "P_FE2CL_REP_GET_GROUP_STYLE_SUCC"
	case P_FE2CL_PC_REGEN:
		return "P_FE2CL_PC_REGEN"
	case P_FE2CL_INSTANCE_MAP_INFO:
		return "P_FE2CL_INSTANCE_MAP_INFO"
	case P_FE2CL_TRANSPORTATION_ENTER:
		return "P_FE2CL_TRANSPORTATION_ENTER"
	case P_FE2CL_TRANSPORTATION_EXIT:
		return "P_FE2CL_TRANSPORTATION_EXIT"
	case P_FE2CL_TRANSPORTATION_MOVE:
		return "P_FE2CL_TRANSPORTATION_MOVE"
	case P_FE2CL_TRANSPORTATION_NEW:
		return "P_FE2CL_TRANSPORTATION_NEW"
	case P_FE2CL_TRANSPORTATION_AROUND:
		return "P_FE2CL_TRANSPORTATION_AROUND"
	case P_FE2CL_AROUND_DEL_TRANSPORTATION:
		return "P_FE2CL_AROUND_DEL_TRANSPORTATION"
	case P_FE2CL_REP_EP_RANK_LIST:
		return "P_FE2CL_REP_EP_RANK_LIST"
	case P_FE2CL_REP_EP_RANK_DETAIL:
		return "P_FE2CL_REP_EP_RANK_DETAIL"
	case P_FE2CL_REP_EP_RANK_PC_INFO:
		return "P_FE2CL_REP_EP_RANK_PC_INFO"
	case P_FE2CL_REP_EP_RACE_START_SUCC:
		return "P_FE2CL_REP_EP_RACE_START_SUCC"
	case P_FE2CL_REP_EP_RACE_START_FAIL:
		return "P_FE2CL_REP_EP_RACE_START_FAIL"
	case P_FE2CL_REP_EP_RACE_END_SUCC:
		return "P_FE2CL_REP_EP_RACE_END_SUCC"
	case P_FE2CL_REP_EP_RACE_END_FAIL:
		return "P_FE2CL_REP_EP_RACE_END_FAIL"
	case P_FE2CL_REP_EP_RACE_CANCEL_SUCC:
		return "P_FE2CL_REP_EP_RACE_CANCEL_SUCC"
	case P_FE2CL_REP_EP_RACE_CANCEL_FAIL:
		return "P_FE2CL_REP_EP_RACE_CANCEL_FAIL"
	case P_FE2CL_REP_EP_GET_RING_SUCC:
		return "P_FE2CL_REP_EP_GET_RING_SUCC"
	case P_FE2CL_REP_EP_GET_RING_FAIL:
		return "P_FE2CL_REP_EP_GET_RING_FAIL"
	case P_FE2CL_REP_IM_CHANGE_SWITCH_STATUS:
		return "P_FE2CL_REP_IM_CHANGE_SWITCH_STATUS"
	case P_FE2CL_SHINY_ENTER:
		return "P_FE2CL_SHINY_ENTER"
	case P_FE2CL_SHINY_EXIT:
		return "P_FE2CL_SHINY_EXIT"
	case P_FE2CL_SHINY_NEW:
		return "P_FE2CL_SHINY_NEW"
	case P_FE2CL_SHINY_AROUND:
		return "P_FE2CL_SHINY_AROUND"
	case P_FE2CL_AROUND_DEL_SHINY:
		return "P_FE2CL_AROUND_DEL_SHINY"
	case P_FE2CL_REP_SHINY_PICKUP_FAIL:
		return "P_FE2CL_REP_SHINY_PICKUP_FAIL"
	case P_FE2CL_REP_SHINY_PICKUP_SUCC:
		return "P_FE2CL_REP_SHINY_PICKUP_SUCC"
	case P_FE2CL_PC_MOVETRANSPORTATION:
		return "P_FE2CL_PC_MOVETRANSPORTATION"
	case P_FE2CL_REP_SEND_ALL_GROUP_FREECHAT_MESSAGE_SUCC:
		return "P_FE2CL_REP_SEND_ALL_GROUP_FREECHAT_MESSAGE_SUCC"
	case P_FE2CL_REP_SEND_ALL_GROUP_FREECHAT_MESSAGE_FAIL:
		return "P_FE2CL_REP_SEND_ALL_GROUP_FREECHAT_MESSAGE_FAIL"
	case P_FE2CL_REP_SEND_ANY_GROUP_FREECHAT_MESSAGE_SUCC:
		return "P_FE2CL_REP_SEND_ANY_GROUP_FREECHAT_MESSAGE_SUCC"
	case P_FE2CL_REP_SEND_ANY_GROUP_FREECHAT_MESSAGE_FAIL:
		return "P_FE2CL_REP_SEND_ANY_GROUP_FREECHAT_MESSAGE_FAIL"
	case P_FE2CL_REP_BARKER:
		return "P_FE2CL_REP_BARKER"
	case P_FE2CL_REP_SEND_ALL_GROUP_MENUCHAT_MESSAGE_SUCC:
		return "P_FE2CL_REP_SEND_ALL_GROUP_MENUCHAT_MESSAGE_SUCC"
	case P_FE2CL_REP_SEND_ALL_GROUP_MENUCHAT_MESSAGE_FAIL:
		return "P_FE2CL_REP_SEND_ALL_GROUP_MENUCHAT_MESSAGE_FAIL"
	case P_FE2CL_REP_SEND_ANY_GROUP_MENUCHAT_MESSAGE_SUCC:
		return "P_FE2CL_REP_SEND_ANY_GROUP_MENUCHAT_MESSAGE_SUCC"
	case P_FE2CL_REP_SEND_ANY_GROUP_MENUCHAT_MESSAGE_FAIL:
		return "P_FE2CL_REP_SEND_ANY_GROUP_MENUCHAT_MESSAGE_FAIL"
	case P_FE2CL_REP_PC_REGIST_TRANSPORTATION_LOCATION_FAIL:
		return "P_FE2CL_REP_PC_REGIST_TRANSPORTATION_LOCATION_FAIL"
	case P_FE2CL_REP_PC_REGIST_TRANSPORTATION_LOCATION_SUCC:
		return "P_FE2CL_REP_PC_REGIST_TRANSPORTATION_LOCATION_SUCC"
	case P_FE2CL_REP_PC_WARP_USE_TRANSPORTATION_FAIL:
		return "P_FE2CL_REP_PC_WARP_USE_TRANSPORTATION_FAIL"
	case P_FE2CL_REP_PC_WARP_USE_TRANSPORTATION_SUCC:
		return "P_FE2CL_REP_PC_WARP_USE_TRANSPORTATION_SUCC"
	case P_FE2CL_ANNOUNCE_MSG:
		return "P_FE2CL_ANNOUNCE_MSG"
	case P_FE2CL_REP_PC_SPECIAL_STATE_SWITCH_SUCC:
		return "P_FE2CL_REP_PC_SPECIAL_STATE_SWITCH_SUCC"
	case P_FE2CL_PC_SPECIAL_STATE_CHANGE:
		return "P_FE2CL_PC_SPECIAL_STATE_CHANGE"
	case P_FE2CL_GM_REP_PC_SET_VALUE:
		return "P_FE2CL_GM_REP_PC_SET_VALUE"
	case P_FE2CL_GM_PC_CHANGE_VALUE:
		return "P_FE2CL_GM_PC_CHANGE_VALUE"
	case P_FE2CL_GM_REP_PC_LOCATION:
		return "P_FE2CL_GM_REP_PC_LOCATION"
	case P_FE2CL_GM_REP_PC_ANNOUNCE:
		return "P_FE2CL_GM_REP_PC_ANNOUNCE"
	case P_FE2CL_REP_PC_BUDDY_WARP_FAIL:
		return "P_FE2CL_REP_PC_BUDDY_WARP_FAIL"
	case P_FE2CL_REP_PC_CHANGE_LEVEL:
		return "P_FE2CL_REP_PC_CHANGE_LEVEL"
	case P_FE2CL_REP_SET_PC_BLOCK_SUCC:
		return "P_FE2CL_REP_SET_PC_BLOCK_SUCC"
	case P_FE2CL_REP_SET_PC_BLOCK_FAIL:
		return "P_FE2CL_REP_SET_PC_BLOCK_FAIL"
	case P_FE2CL_REP_REGIST_RXCOM:
		return "P_FE2CL_REP_REGIST_RXCOM"
	case P_FE2CL_REP_REGIST_RXCOM_FAIL:
		return "P_FE2CL_REP_REGIST_RXCOM_FAIL"
	case P_FE2CL_PC_INVEN_FULL_MSG:
		return "P_FE2CL_PC_INVEN_FULL_MSG"
	case P_FE2CL_REQ_LIVE_CHECK:
		return "P_FE2CL_REQ_LIVE_CHECK"
	case P_FE2CL_PC_MOTD_LOGIN:
		return "P_FE2CL_PC_MOTD_LOGIN"
	case P_FE2CL_REP_PC_ITEM_USE_FAIL:
		return "P_FE2CL_REP_PC_ITEM_USE_FAIL"
	case P_FE2CL_REP_PC_ITEM_USE_SUCC:
		return "P_FE2CL_REP_PC_ITEM_USE_SUCC"
	case P_FE2CL_PC_ITEM_USE:
		return "P_FE2CL_PC_ITEM_USE"
	case P_FE2CL_REP_GET_BUDDY_LOCATION_SUCC:
		return "P_FE2CL_REP_GET_BUDDY_LOCATION_SUCC"
	case P_FE2CL_REP_GET_BUDDY_LOCATION_FAIL:
		return "P_FE2CL_REP_GET_BUDDY_LOCATION_FAIL"
	case P_FE2CL_REP_PC_RIDING_FAIL:
		return "P_FE2CL_REP_PC_RIDING_FAIL"
	case P_FE2CL_REP_PC_RIDING_SUCC:
		return "P_FE2CL_REP_PC_RIDING_SUCC"
	case P_FE2CL_PC_RIDING:
		return "P_FE2CL_PC_RIDING"
	case P_FE2CL_PC_BROOMSTICK_MOVE:
		return "P_FE2CL_PC_BROOMSTICK_MOVE"
	case P_FE2CL_REP_PC_BUDDY_WARP_OTHER_SHARD_SUCC:
		return "P_FE2CL_REP_PC_BUDDY_WARP_OTHER_SHARD_SUCC"
	case P_FE2CL_REP_WARP_USE_RECALL_FAIL:
		return "P_FE2CL_REP_WARP_USE_RECALL_FAIL"
	case P_FE2CL_REP_PC_EXIT_DUPLICATE:
		return "P_FE2CL_REP_PC_EXIT_DUPLICATE"
	case P_FE2CL_REP_PC_MISSION_COMPLETE_SUCC:
		return "P_FE2CL_REP_PC_MISSION_COMPLETE_SUCC"
	case P_FE2CL_PC_BUFF_UPDATE:
		return "P_FE2CL_PC_BUFF_UPDATE"
	case P_FE2CL_REP_PC_NEW_EMAIL:
		return "P_FE2CL_REP_PC_NEW_EMAIL"
	case P_FE2CL_REP_PC_READ_EMAIL_SUCC:
		return "P_FE2CL_REP_PC_READ_EMAIL_SUCC"
	case P_FE2CL_REP_PC_READ_EMAIL_FAIL:
		return "P_FE2CL_REP_PC_READ_EMAIL_FAIL"
	case P_FE2CL_REP_PC_RECV_EMAIL_PAGE_LIST_SUCC:
		return "P_FE2CL_REP_PC_RECV_EMAIL_PAGE_LIST_SUCC"
	case P_FE2CL_REP_PC_RECV_EMAIL_PAGE_LIST_FAIL:
		return "P_FE2CL_REP_PC_RECV_EMAIL_PAGE_LIST_FAIL"
	case P_FE2CL_REP_PC_DELETE_EMAIL_SUCC:
		return "P_FE2CL_REP_PC_DELETE_EMAIL_SUCC"
	case P_FE2CL_REP_PC_DELETE_EMAIL_FAIL:
		return "P_FE2CL_REP_PC_DELETE_EMAIL_FAIL"
	case P_FE2CL_REP_PC_SEND_EMAIL_SUCC:
		return "P_FE2CL_REP_PC_SEND_EMAIL_SUCC"
	case P_FE2CL_REP_PC_SEND_EMAIL_FAIL:
		return "P_FE2CL_REP_PC_SEND_EMAIL_FAIL"
	case P_FE2CL_REP_PC_RECV_EMAIL_ITEM_SUCC:
		return "P_FE2CL_REP_PC_RECV_EMAIL_ITEM_SUCC"
	case P_FE2CL_REP_PC_RECV_EMAIL_ITEM_FAIL:
		return "P_FE2CL_REP_PC_RECV_EMAIL_ITEM_FAIL"
	case P_FE2CL_REP_PC_RECV_EMAIL_CANDY_SUCC:
		return "P_FE2CL_REP_PC_RECV_EMAIL_CANDY_SUCC"
	case P_FE2CL_REP_PC_RECV_EMAIL_CANDY_FAIL:
		return "P_FE2CL_REP_PC_RECV_EMAIL_CANDY_FAIL"
	case P_FE2CL_PC_SUDDEN_DEAD:
		return "P_FE2CL_PC_SUDDEN_DEAD"
	case P_FE2CL_REP_GM_REQ_TARGET_PC_SPECIAL_STATE_ONOFF_SUCC:
		return "P_FE2CL_REP_GM_REQ_TARGET_PC_SPECIAL_STATE_ONOFF_SUCC"
	case P_FE2CL_REP_PC_SET_CURRENT_MISSION_ID:
		return "P_FE2CL_REP_PC_SET_CURRENT_MISSION_ID"
	case P_FE2CL_REP_NPC_GROUP_INVITE_FAIL:
		return "P_FE2CL_REP_NPC_GROUP_INVITE_FAIL"
	case P_FE2CL_REP_NPC_GROUP_INVITE_SUCC:
		return "P_FE2CL_REP_NPC_GROUP_INVITE_SUCC"
	case P_FE2CL_REP_NPC_GROUP_KICK_FAIL:
		return "P_FE2CL_REP_NPC_GROUP_KICK_FAIL"
	case P_FE2CL_REP_NPC_GROUP_KICK_SUCC:
		return "P_FE2CL_REP_NPC_GROUP_KICK_SUCC"
	case P_FE2CL_PC_EVENT:
		return "P_FE2CL_PC_EVENT"
	case P_FE2CL_REP_PC_TRANSPORT_WARP_SUCC:
		return "P_FE2CL_REP_PC_TRANSPORT_WARP_SUCC"
	case P_FE2CL_REP_PC_TRADE_EMOTES_CHAT_FAIL:
		return "P_FE2CL_REP_PC_TRADE_EMOTES_CHAT_FAIL"
	case P_FE2CL_REP_PC_RECV_EMAIL_ITEM_ALL_SUCC:
		return "P_FE2CL_REP_PC_RECV_EMAIL_ITEM_ALL_SUCC"
	case P_FE2CL_REP_PC_RECV_EMAIL_ITEM_ALL_FAIL:
		return "P_FE2CL_REP_PC_RECV_EMAIL_ITEM_ALL_FAIL"
	case P_FE2CL_REP_PC_LOADING_COMPLETE_SUCC:
		return "P_FE2CL_REP_PC_LOADING_COMPLETE_SUCC"
	case P_FE2CL_REP_CHANNEL_INFO:
		return "P_FE2CL_REP_CHANNEL_INFO"
	case P_FE2CL_REP_PC_CHANNEL_NUM:
		return "P_FE2CL_REP_PC_CHANNEL_NUM"
	case P_FE2CL_REP_PC_WARP_CHANNEL_FAIL:
		return "P_FE2CL_REP_PC_WARP_CHANNEL_FAIL"
	case P_FE2CL_REP_PC_WARP_CHANNEL_SUCC:
		return "P_FE2CL_REP_PC_WARP_CHANNEL_SUCC"
	case P_FE2CL_REP_PC_FIND_NAME_MAKE_BUDDY_SUCC:
		return "P_FE2CL_REP_PC_FIND_NAME_MAKE_BUDDY_SUCC"
	case P_FE2CL_REP_PC_FIND_NAME_MAKE_BUDDY_FAIL:
		return "P_FE2CL_REP_PC_FIND_NAME_MAKE_BUDDY_FAIL"
	case P_FE2CL_REP_PC_FIND_NAME_ACCEPT_BUDDY_FAIL:
		return "P_FE2CL_REP_PC_FIND_NAME_ACCEPT_BUDDY_FAIL"
	case P_FE2CL_REP_PC_BUDDY_WARP_SAME_SHARD_SUCC:
		return "P_FE2CL_REP_PC_BUDDY_WARP_SAME_SHARD_SUCC"
	case P_FE2CL_PC_ATTACK_CHARs_SUCC:
		return "P_FE2CL_PC_ATTACK_CHARs_SUCC"
	case P_FE2CL_PC_ATTACK_CHARs:
		return "P_FE2CL_PC_ATTACK_CHARs"
	case P_FE2CL_NPC_ATTACK_CHARs:
		return "P_FE2CL_NPC_ATTACK_CHARs"
	case P_FE2CL_REP_PC_CHANGE_LEVEL_SUCC:
		return "P_FE2CL_REP_PC_CHANGE_LEVEL_SUCC"
	case P_FE2CL_REP_PC_NANO_CREATE:
		return "P_FE2CL_REP_PC_NANO_CREATE"
	case P_FE2CL_PC_STREETSTALL_REP_READY_SUCC:
		return "P_FE2CL_PC_STREETSTALL_REP_READY_SUCC"
	case P_FE2CL_PC_STREETSTALL_REP_READY_FAIL:
		return "P_FE2CL_PC_STREETSTALL_REP_READY_FAIL"
	case P_FE2CL_PC_STREETSTALL_REP_CANCEL_SUCC:
		return "P_FE2CL_PC_STREETSTALL_REP_CANCEL_SUCC"
	case P_FE2CL_PC_STREETSTALL_REP_CANCEL_FAIL:
		return "P_FE2CL_PC_STREETSTALL_REP_CANCEL_FAIL"
	case P_FE2CL_PC_STREETSTALL_REP_REGIST_ITEM_SUCC:
		return "P_FE2CL_PC_STREETSTALL_REP_REGIST_ITEM_SUCC"
	case P_FE2CL_PC_STREETSTALL_REP_REGIST_ITEM_FAIL:
		return "P_FE2CL_PC_STREETSTALL_REP_REGIST_ITEM_FAIL"
	case P_FE2CL_PC_STREETSTALL_REP_UNREGIST_ITEM_SUCC:
		return "P_FE2CL_PC_STREETSTALL_REP_UNREGIST_ITEM_SUCC"
	case P_FE2CL_PC_STREETSTALL_REP_UNREGIST_ITEM_FAIL:
		return "P_FE2CL_PC_STREETSTALL_REP_UNREGIST_ITEM_FAIL"
	case P_FE2CL_PC_STREETSTALL_REP_SALE_START_SUCC:
		return "P_FE2CL_PC_STREETSTALL_REP_SALE_START_SUCC"
	case P_FE2CL_PC_STREETSTALL_REP_SALE_START_FAIL:
		return "P_FE2CL_PC_STREETSTALL_REP_SALE_START_FAIL"
	case P_FE2CL_PC_STREETSTALL_REP_ITEM_LIST:
		return "P_FE2CL_PC_STREETSTALL_REP_ITEM_LIST"
	case P_FE2CL_PC_STREETSTALL_REP_ITEM_LIST_FAIL:
		return "P_FE2CL_PC_STREETSTALL_REP_ITEM_LIST_FAIL"
	case P_FE2CL_PC_STREETSTALL_REP_ITEM_BUY_SUCC_BUYER:
		return "P_FE2CL_PC_STREETSTALL_REP_ITEM_BUY_SUCC_BUYER"
	case P_FE2CL_PC_STREETSTALL_REP_ITEM_BUY_SUCC_SELLER:
		return "P_FE2CL_PC_STREETSTALL_REP_ITEM_BUY_SUCC_SELLER"
	case P_FE2CL_PC_STREETSTALL_REP_ITEM_BUY_FAIL:
		return "P_FE2CL_PC_STREETSTALL_REP_ITEM_BUY_FAIL"
	case P_FE2CL_REP_PC_ITEM_COMBINATION_SUCC:
		return "P_FE2CL_REP_PC_ITEM_COMBINATION_SUCC"
	case P_FE2CL_REP_PC_ITEM_COMBINATION_FAIL:
		return "P_FE2CL_REP_PC_ITEM_COMBINATION_FAIL"
	case P_FE2CL_PC_CASH_BUFF_UPDATE:
		return "P_FE2CL_PC_CASH_BUFF_UPDATE"
	case P_FE2CL_REP_PC_SKILL_ADD_SUCC:
		return "P_FE2CL_REP_PC_SKILL_ADD_SUCC"
	case P_FE2CL_REP_PC_SKILL_ADD_FAIL:
		return "P_FE2CL_REP_PC_SKILL_ADD_FAIL"
	case P_FE2CL_REP_PC_SKILL_DEL_SUCC:
		return "P_FE2CL_REP_PC_SKILL_DEL_SUCC"
	case P_FE2CL_REP_PC_SKILL_DEL_FAIL:
		return "P_FE2CL_REP_PC_SKILL_DEL_FAIL"
	case P_FE2CL_REP_PC_SKILL_USE_SUCC:
		return "P_FE2CL_REP_PC_SKILL_USE_SUCC"
	case P_FE2CL_REP_PC_SKILL_USE_FAIL:
		return "P_FE2CL_REP_PC_SKILL_USE_FAIL"
	case P_FE2CL_PC_SKILL_USE:
		return "P_FE2CL_PC_SKILL_USE"
	case P_FE2CL_PC_ROPE:
		return "P_FE2CL_PC_ROPE"
	case P_FE2CL_PC_BELT:
		return "P_FE2CL_PC_BELT"
	case P_FE2CL_PC_VEHICLE_ON_SUCC:
		return "P_FE2CL_PC_VEHICLE_ON_SUCC"
	case P_FE2CL_PC_VEHICLE_ON_FAIL:
		return "P_FE2CL_PC_VEHICLE_ON_FAIL"
	case P_FE2CL_PC_VEHICLE_OFF_SUCC:
		return "P_FE2CL_PC_VEHICLE_OFF_SUCC"
	case P_FE2CL_PC_VEHICLE_OFF_FAIL:
		return "P_FE2CL_PC_VEHICLE_OFF_FAIL"
	case P_FE2CL_PC_QUICK_SLOT_INFO:
		return "P_FE2CL_PC_QUICK_SLOT_INFO"
	case P_FE2CL_REP_PC_REGIST_QUICK_SLOT_FAIL:
		return "P_FE2CL_REP_PC_REGIST_QUICK_SLOT_FAIL"
	case P_FE2CL_REP_PC_REGIST_QUICK_SLOT_SUCC:
		return "P_FE2CL_REP_PC_REGIST_QUICK_SLOT_SUCC"
	case P_FE2CL_PC_DELETE_TIME_LIMIT_ITEM:
		return "P_FE2CL_PC_DELETE_TIME_LIMIT_ITEM"
	case P_FE2CL_REP_PC_DISASSEMBLE_ITEM_SUCC:
		return "P_FE2CL_REP_PC_DISASSEMBLE_ITEM_SUCC"
	case P_FE2CL_REP_PC_DISASSEMBLE_ITEM_FAIL:
		return "P_FE2CL_REP_PC_DISASSEMBLE_ITEM_FAIL"
	case P_FE2CL_GM_REP_REWARD_RATE_SUCC:
		return "P_FE2CL_GM_REP_REWARD_RATE_SUCC"
	case P_FE2CL_REP_PC_ITEM_ENCHANT_SUCC:
		return "P_FE2CL_REP_PC_ITEM_ENCHANT_SUCC"
	case P_FE2CL_REP_PC_ITEM_ENCHANT_FAIL:
		return "P_FE2CL_REP_PC_ITEM_ENCHANT_FAIL"
	case P_LS2CL_REP_LOGIN_SUCC:
		return "P_LS2CL_REP_LOGIN_SUCC"
	case P_LS2CL_REP_LOGIN_FAIL:
		return "P_LS2CL_REP_LOGIN_FAIL"
	case P_LS2CL_REP_CHAR_INFO:
		return "P_LS2CL_REP_CHAR_INFO"
	case P_LS2CL_REP_CHECK_CHAR_NAME_SUCC:
		return "P_LS2CL_REP_CHECK_CHAR_NAME_SUCC"
	case P_LS2CL_REP_CHECK_CHAR_NAME_FAIL:
		return "P_LS2CL_REP_CHECK_CHAR_NAME_FAIL"
	case P_LS2CL_REP_SAVE_CHAR_NAME_SUCC:
		return "P_LS2CL_REP_SAVE_CHAR_NAME_SUCC"
	case P_LS2CL_REP_SAVE_CHAR_NAME_FAIL:
		return "P_LS2CL_REP_SAVE_CHAR_NAME_FAIL"
	case P_LS2CL_REP_CHAR_CREATE_SUCC:
		return "P_LS2CL_REP_CHAR_CREATE_SUCC"
	case P_LS2CL_REP_CHAR_CREATE_FAIL:
		return "P_LS2CL_REP_CHAR_CREATE_FAIL"
	case P_LS2CL_REP_CHAR_SELECT_SUCC:
		return "P_LS2CL_REP_CHAR_SELECT_SUCC"
	case P_LS2CL_REP_CHAR_SELECT_FAIL:
		return "P_LS2CL_REP_CHAR_SELECT_FAIL"
	case P_LS2CL_REP_CHAR_DELETE_SUCC:
		return "P_LS2CL_REP_CHAR_DELETE_SUCC"
	case P_LS2CL_REP_CHAR_DELETE_FAIL:
		return "P_LS2CL_REP_CHAR_DELETE_FAIL"
	case P_LS2CL_REP_SHARD_SELECT_SUCC:
		return "P_LS2CL_REP_SHARD_SELECT_SUCC"
	case P_LS2CL_REP_SHARD_SELECT_FAIL:
		return "P_LS2CL_REP_SHARD_SELECT_FAIL"
	case P_LS2CL_REP_VERSION_CHECK_SUCC:
		return "P_LS2CL_REP_VERSION_CHECK_SUCC"
	case P_LS2CL_REP_VERSION_CHECK_FAIL:
		return "P_LS2CL_REP_VERSION_CHECK_FAIL"
	case P_LS2CL_REP_CHECK_NAME_LIST_SUCC:
		return "P_LS2CL_REP_CHECK_NAME_LIST_SUCC"
	case P_LS2CL_REP_CHECK_NAME_LIST_FAIL:
		return "P_LS2CL_REP_CHECK_NAME_LIST_FAIL"
	case P_LS2CL_REP_PC_EXIT_DUPLICATE:
		return "P_LS2CL_REP_PC_EXIT_DUPLICATE"
	case P_LS2CL_REQ_LIVE_CHECK:
		return "P_LS2CL_REQ_LIVE_CHECK"
	case P_LS2CL_REP_CHANGE_CHAR_NAME_SUCC:
		return "P_LS2CL_REP_CHANGE_CHAR_NAME_SUCC"
	case P_LS2CL_REP_CHANGE_CHAR_NAME_FAIL:
		return "P_LS2CL_REP_CHANGE_CHAR_NAME_FAIL"
	case P_LS2CL_REP_SHARD_LIST_INFO_SUCC:
		return "P_LS2CL_REP_SHARD_LIST_INFO_SUCC"
	}
	return "UNKNOWN"
}

type SP_CL2FE_REQ_PC_ENTER struct {
	SzID            string `size:"33" pad:"2"`
	ITempValue      int32
	IEnterSerialKey int64
	// SIZE: 80
}

type SP_CL2FE_REQ_PC_EXIT struct {
	IID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_MOVE struct {
	ICliTime  uint64
	IX        int32
	IY        int32
	IZ        int32
	FVX       float32
	FVY       float32
	FVZ       float32
	IAngle    int32
	CKeyValue uint8 `pad:"3"`
	ISpeed    int32
	// SIZE: 44
}

type SP_CL2FE_REQ_PC_STOP struct {
	ICliTime uint64
	IX       int32
	IY       int32
	IZ       int32
	// SIZE: 20
}

type SP_CL2FE_REQ_PC_JUMP struct {
	ICliTime  uint64
	IX        int32
	IY        int32
	IZ        int32
	IVX       int32
	IVY       int32
	IVZ       int32
	IAngle    int32
	CKeyValue uint8 `pad:"3"`
	ISpeed    int32
	// SIZE: 44
}

type SP_CL2FE_REQ_PC_ATTACK_NPCs struct {
	INPCCnt int32
	// SIZE: 4
}

type SP_CL2FE_REQ_SEND_FREECHAT_MESSAGE struct {
	SzFreeChat string `size:"128"`
	IEmoteCode int32
	// SIZE: 260
}

type SP_CL2FE_REQ_SEND_MENUCHAT_MESSAGE struct {
	SzFreeChat string `size:"128"`
	IEmoteCode int32
	// SIZE: 260
}

type SP_CL2FE_REQ_PC_REGEN struct {
	IRegenType int32
	EIL        int32
	IIndex     int32
	// SIZE: 12
}

type SP_CL2FE_REQ_ITEM_MOVE struct {
	EFrom        int32
	IFromSlotNum int32
	ETo          int32
	IToSlotNum   int32
	// SIZE: 16
}

type SP_CL2FE_REQ_PC_TASK_START struct {
	ITaskNum      int32
	INPC_ID       int32
	IEscortNPC_ID int32
	// SIZE: 12
}

type SP_CL2FE_REQ_PC_TASK_END struct {
	ITaskNum      int32
	INPC_ID       int32
	IBox1Choice   int8
	IBox2Choice   int8 `pad:"2"`
	IEscortNPC_ID int32
	// SIZE: 16
}

type SP_CL2FE_REQ_NANO_EQUIP struct {
	INanoID      int16
	INanoSlotNum int16
	// SIZE: 4
}

type SP_CL2FE_REQ_NANO_UNEQUIP struct {
	INanoSlotNum int16
	// SIZE: 2
}

type SP_CL2FE_REQ_NANO_ACTIVE struct {
	INanoSlotNum int16
	// SIZE: 2
}

type SP_CL2FE_REQ_NANO_TUNE struct {
	INanoID           int16
	ITuneID           int16
	AiNeedItemSlotNum [10]int32
	// SIZE: 44
}

type SP_CL2FE_REQ_NANO_SKILL_USE struct {
	IBulletID  int8 `pad:"3"`
	IArg1      int32
	IArg2      int32
	IArg3      int32
	ITargetCnt int32
	// SIZE: 20
}

type SP_CL2FE_REQ_PC_TASK_STOP struct {
	ITaskNum int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_TASK_CONTINUE struct {
	ITaskNum int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_GOTO struct {
	IToX int32
	IToY int32
	IToZ int32
	// SIZE: 12
}

type SP_CL2FE_REQ_CHARGE_NANO_STAMINA struct {
	IPC_ID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_KILL_QUEST_NPCs struct {
	INPCCnt int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_VENDOR_ITEM_SELL struct {
	IInvenSlotNum int32
	IItemCnt      int32
	// SIZE: 8
}

type SP_CL2FE_REQ_PC_ITEM_DELETE struct {
	EIL      int32
	ISlotNum int32
	// SIZE: 8
}

type SP_CL2FE_REQ_PC_ROCKET_STYLE_READY struct {
	ISkillID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_ROCKET_STYLE_FIRE struct {
	ISkillID int32
	IX       int32
	IY       int32
	IZ       int32
	IToX     int32
	IToY     int32
	IToZ     int32
	// SIZE: 28
}

type SP_CL2FE_REQ_PC_ROCKET_STYLE_HIT struct {
	IBulletID  int8 `pad:"3"`
	IX         int32
	IY         int32
	IZ         int32
	ITargetCnt int32
	// SIZE: 20
}

type SP_CL2FE_REQ_PC_GRENADE_STYLE_READY struct {
	ISkillID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_GRENADE_STYLE_FIRE struct {
	ISkillID int32
	IToX     int32
	IToY     int32
	IToZ     int32
	// SIZE: 16
}

type SP_CL2FE_REQ_PC_GRENADE_STYLE_HIT struct {
	IBulletID  int8 `pad:"3"`
	IX         int32
	IY         int32
	IZ         int32
	ITargetCnt int32
	// SIZE: 20
}

type SP_CL2FE_REQ_PC_NANO_CREATE struct {
	INanoID               int16 `pad:"2"`
	INeedQuestItemSlotNum int32
	// SIZE: 8
}

type SP_CL2FE_REQ_PC_TRADE_OFFER struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	// SIZE: 12
}

type SP_CL2FE_REQ_PC_TRADE_OFFER_CANCEL struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	// SIZE: 12
}

type SP_CL2FE_REQ_PC_TRADE_OFFER_ACCEPT struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	// SIZE: 12
}

type SP_CL2FE_REQ_PC_TRADE_OFFER_REFUSAL struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	// SIZE: 12
}

type SP_CL2FE_REQ_PC_TRADE_OFFER_ABORT struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	IErrorCode  int16 `pad:"2"`
	// SIZE: 16
}

type SP_CL2FE_REQ_PC_TRADE_CONFIRM struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	// SIZE: 12
}

type SP_CL2FE_REQ_PC_TRADE_CONFIRM_CANCEL struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	// SIZE: 12
}

type SP_CL2FE_REQ_PC_TRADE_CONFIRM_ABORT struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	// SIZE: 12
}

type SP_CL2FE_REQ_PC_TRADE_CASH_REGISTER struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	ICandy      int32
	// SIZE: 16
}

type SP_CL2FE_REQ_PC_TRADE_EMOTES_CHAT struct {
	IID_Request  int32
	IID_From     int32
	IID_To       int32
	SzFreeChat   string `size:"128"`
	IEmoteCode   int32
	IFreeChatUse int8 `pad:"3"`
	// SIZE: 276
}

type SP_CL2FE_REQ_PC_BANK_OPEN struct {
	IPC_ID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_BANK_CLOSE struct {
	IPC_ID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_VENDOR_START struct {
	INPC_ID   int32
	IVendorID int32
	// SIZE: 8
}

type SP_CL2FE_REQ_PC_VENDOR_TABLE_UPDATE struct {
	INPC_ID   int32
	IVendorID int32
	// SIZE: 8
}

type SP_CL2FE_REQ_PC_COMBAT_BEGIN struct {
	IPC_ID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_COMBAT_END struct {
	IPC_ID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_REQUEST_MAKE_BUDDY struct {
	IBuddyID    int32
	IBuddyPCUID int64
	// SIZE: 12
}

type SP_CL2FE_REQ_ACCEPT_MAKE_BUDDY struct {
	IAcceptFlag int8 `pad:"3"`
	IBuddyID    int32
	IBuddyPCUID int64
	// SIZE: 16
}

type SP_CL2FE_REQ_SEND_BUDDY_FREECHAT_MESSAGE struct {
	SzFreeChat  string `size:"128"`
	IEmoteCode  int32
	IBuddyPCUID int64
	IBuddySlot  int8 `pad:"3"`
	// SIZE: 272
}

type SP_CL2FE_REQ_SEND_BUDDY_MENUCHAT_MESSAGE struct {
	SzFreeChat  string `size:"128"`
	IEmoteCode  int32
	IBuddyPCUID int64
	IBuddySlot  int8 `pad:"3"`
	// SIZE: 272
}

type SP_CL2FE_REQ_GET_BUDDY_STYLE struct {
	IBuddyPCUID int64
	IBuddySlot  int8 `pad:"3"`
	// SIZE: 12
}

type SP_CL2FE_REQ_SET_BUDDY_BLOCK struct {
	IBuddyPCUID int64
	IBuddySlot  int8 `pad:"3"`
	// SIZE: 12
}

type SP_CL2FE_REQ_REMOVE_BUDDY struct {
	IBuddyPCUID int64
	IBuddySlot  int8 `pad:"3"`
	// SIZE: 12
}

type SP_CL2FE_REQ_PC_JUMPPAD struct {
	ICliTime  uint64
	IX        int32
	IY        int32
	IZ        int32
	IVX       int32
	IVY       int32
	IVZ       int32
	IAngle    int32
	CKeyValue uint8 `pad:"3"`
	// SIZE: 40
}

type SP_CL2FE_REQ_PC_LAUNCHER struct {
	ICliTime uint64
	IX       int32
	IY       int32
	IZ       int32
	IVX      int32
	IVY      int32
	IVZ      int32
	IAngle   int32
	ISpeed   int32
	// SIZE: 40
}

type SP_CL2FE_REQ_PC_ZIPLINE struct {
	ICliTime     uint64
	IStX         int32
	IStY         int32
	IStZ         int32
	FMovDistance float32
	FMaxDistance float32
	FDummy       float32
	IX           int32
	IY           int32
	IZ           int32
	FVX          float32
	FVY          float32
	FVZ          float32
	BDown        int32
	IRollMax     int32
	IRoll        uint8 `pad:"3"`
	IAngle       int32
	ISpeed       int32
	// SIZE: 76
}

type SP_CL2FE_REQ_PC_MOVEPLATFORM struct {
	ICliTime    uint64
	ILcX        int32
	ILcY        int32
	ILcZ        int32
	IX          int32
	IY          int32
	IZ          int32
	FVX         float32
	FVY         float32
	FVZ         float32
	BDown       int32
	IPlatformID int32
	IAngle      int32
	CKeyValue   uint8 `pad:"3"`
	ISpeed      int32
	// SIZE: 64
}

type SP_CL2FE_REQ_PC_SLOPE struct {
	ICliTime  uint64
	IX        int32
	IY        int32
	IZ        int32
	IAngle    int32
	ISpeed    int32
	CKeyValue uint8 `pad:"3"`
	FVX       float32
	FVY       float32
	FVZ       float32
	ISlopeID  int32
	// SIZE: 48
}

type SP_CL2FE_REQ_PC_STATE_CHANGE struct {
	IState int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_MAP_WARP struct {
	IMapNum int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_GIVE_NANO struct {
	INanoID int16
	// SIZE: 2
}

type SP_CL2FE_REQ_NPC_SUMMON struct {
	INPCType int32
	INPCCnt  int16 `pad:"2"`
	// SIZE: 8
}

type SP_CL2FE_REQ_NPC_UNSUMMON struct {
	INPC_ID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_GIVE_NANO_SKILL struct {
	INanoID      int16
	INanoSkillID int16
	// SIZE: 4
}

type SP_CL2FE_DOT_DAMAGE_ONOFF struct {
	IFlag int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_WARP_USE_NPC struct {
	INPC_ID    int32
	IWarpID    int32
	EIL1       int32
	IItemSlot1 int32
	EIL2       int32
	IItemSlot2 int32
	// SIZE: 24
}

type SP_CL2FE_REQ_PC_GROUP_INVITE struct {
	IID_To int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_GROUP_INVITE_REFUSE struct {
	IID_From int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_GROUP_JOIN struct {
	IID_From int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_AVATAR_EMOTES_CHAT struct {
	IID_From   int32
	IEmoteCode int32
	// SIZE: 8
}

type SP_CL2FE_REQ_PC_BUDDY_WARP struct {
	IBuddyPCUID int64
	ISlotNum    int8 `pad:"3"`
	// SIZE: 12
}

type SP_CL2FE_REQ_GET_MEMBER_STYLE struct {
	IMemberID  int32
	IMemberUID int64
	// SIZE: 12
}

type SP_CL2FE_REQ_PC_CHANGE_MENTOR struct {
	IMentor int16
	// SIZE: 2
}

type SP_CL2FE_REQ_GET_BUDDY_LOCATION struct {
	IBuddyPCUID int64
	ISlotNum    int8 `pad:"3"`
	// SIZE: 12
}

type SP_CL2FE_REQ_NPC_GROUP_SUMMON struct {
	INPCGroupType int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_WARP_TO_PC struct {
	IPC_ID int32
	IPCUID int32
	// SIZE: 8
}

type SP_CL2FE_REQ_EP_RANK_GET_LIST struct {
	IRankListPageNum int32
	// SIZE: 4
}

type SP_CL2FE_REQ_EP_RANK_GET_DETAIL struct {
	IEP_ID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_EP_RANK_GET_PC_INFO struct {
	IEP_ID      int32
	SzFirstName string `size:"9"`
	SzLastName  string `size:"17"`
	// SIZE: 56
}

type SP_CL2FE_REQ_EP_RACE_START struct {
	IStartEcomID         int32
	IEPRaceMode          int32
	IEPTicketItemSlotNum int32
	// SIZE: 12
}

type SP_CL2FE_REQ_EP_RACE_END struct {
	IEndEcomID           int32
	IEPTicketItemSlotNum int32
	// SIZE: 8
}

type SP_CL2FE_REQ_EP_RACE_CANCEL struct {
	IStartEcomID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_EP_GET_RING struct {
	IRingLID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_IM_CHANGE_SWITCH_STATUS struct {
	ISwitchLID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_SHINY_PICKUP struct {
	IShinyID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_SHINY_SUMMON struct {
	IShinyType int32
	IX         int32
	IY         int32
	IZ         int32
	// SIZE: 16
}

type SP_CL2FE_REQ_PC_MOVETRANSPORTATION struct {
	ICliTime  uint64
	ILcX      int32
	ILcY      int32
	ILcZ      int32
	IX        int32
	IY        int32
	IZ        int32
	FVX       float32
	FVY       float32
	FVZ       float32
	IT_ID     int32
	IAngle    int32
	CKeyValue uint8 `pad:"3"`
	ISpeed    int32
	// SIZE: 60
}

type SP_CL2FE_REQ_SEND_ALL_GROUP_FREECHAT_MESSAGE struct {
	SzFreeChat string `size:"128"`
	IEmoteCode int32
	// SIZE: 260
}

type SP_CL2FE_REQ_SEND_ANY_GROUP_FREECHAT_MESSAGE struct {
	SzFreeChat  string `size:"128"`
	IEmoteCode  int32
	IGroupPC_ID int32
	// SIZE: 264
}

type SP_CL2FE_REQ_BARKER struct {
	IMissionTaskID int32
	INPC_ID        int32
	// SIZE: 8
}

type SP_CL2FE_REQ_SEND_ALL_GROUP_MENUCHAT_MESSAGE struct {
	SzFreeChat string `size:"128"`
	IEmoteCode int32
	// SIZE: 260
}

type SP_CL2FE_REQ_SEND_ANY_GROUP_MENUCHAT_MESSAGE struct {
	SzFreeChat  string `size:"128"`
	IEmoteCode  int32
	IGroupPC_ID int32
	// SIZE: 264
}

type SP_CL2FE_REQ_REGIST_TRANSPORTATION_LOCATION struct {
	ETT         int32
	INPC_ID     int32
	ILocationID int32
	// SIZE: 12
}

type SP_CL2FE_REQ_PC_WARP_USE_TRANSPORTATION struct {
	INPC_ID          int32
	ITransporationID int32
	EIL              int32
	ISlotNum         int32
	// SIZE: 16
}

type SP_CL2FE_GM_REQ_PC_SPECIAL_STATE_SWITCH struct {
	IPC_ID            int32
	ISpecialStateFlag int8 `pad:"3"`
	// SIZE: 8
}

type SP_CL2FE_GM_REQ_PC_SET_VALUE struct {
	IPC_ID        int32
	ISetValueType int32
	ISetValue     int32
	// SIZE: 12
}

type SP_CL2FE_GM_REQ_KICK_PLAYER struct {
	IPC_ID               int32
	ETargetSearchBy      int32
	ITargetPC_ID         int32
	SzTargetPC_FirstName string `size:"10"`
	SzTargetPC_LastName  string `size:"18"`
	ITargetPC_UID        int64
	// SIZE: 76
}

type SP_CL2FE_GM_REQ_TARGET_PC_TELEPORT struct {
	IPC_ID               int32
	ETargetPCSearchBy    int32
	ITargetPC_ID         int32
	SzTargetPC_FirstName string `size:"10"`
	SzTargetPC_LastName  string `size:"18"`
	ITargetPC_UID        int64
	ETeleportType        int32
	IToMapType           int32
	IToMap               int32
	IToX                 int32
	IToY                 int32
	IToZ                 int32
	EGoalPCSearchBy      int32
	IGoalPC_ID           int32
	SzGoalPC_FirstName   string `size:"10"`
	SzGoalPC_LastName    string `size:"18"`
	IGoalPC_UID          int64
	// SIZE: 172
}

type SP_CL2FE_GM_REQ_PC_LOCATION struct {
	ETargetSearchBy      int32
	ITargetPC_ID         int32
	SzTargetPC_FirstName string `size:"10"`
	SzTargetPC_LastName  string `size:"18"`
	ITargetPC_UID        int64
	// SIZE: 72
}

type SP_CL2FE_GM_REQ_PC_ANNOUNCE struct {
	IAreaType     int8
	IAnnounceType int8 `pad:"2"`
	IDuringTime   int32
	SzAnnounceMsg string `size:"512"`
	// SIZE: 1032
}

type SP_CL2FE_REQ_SET_PC_BLOCK struct {
	IBlock_ID    int32
	IBlock_PCUID int64
	// SIZE: 12
}

type SP_CL2FE_REQ_REGIST_RXCOM struct {
	INPCID int32
	// SIZE: 4
}

type SP_CL2FE_GM_REQ_PC_MOTD_REGISTER struct {
	IType       int8   `pad:"1"`
	SzSystemMsg string `size:"512"`
	// SIZE: 1026
}

type SP_CL2FE_REQ_ITEM_USE struct {
	EIL       int32
	ISlotNum  int32
	INanoSlot int16 `pad:"2"`
	// SIZE: 12
}

type SP_CL2FE_REQ_WARP_USE_RECALL struct {
	IGroupMemberID int32
	// SIZE: 4
}

type SP_CL2FE_REP_LIVE_CHECK struct {
	ITempValue int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_MISSION_COMPLETE struct {
	IMissionNum int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_TASK_COMPLETE struct {
	ITaskNum int32
	// SIZE: 4
}

type SP_CL2FE_REQ_NPC_INTERACTION struct {
	INPC_ID int32
	BFlag   int32
	// SIZE: 8
}

type SP_CL2FE_DOT_HEAL_ONOFF struct {
	IFlag int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_SPECIAL_STATE_SWITCH struct {
	IPC_ID            int32
	ISpecialStateFlag int8 `pad:"3"`
	// SIZE: 8
}

type SP_CL2FE_REQ_PC_READ_EMAIL struct {
	IEmailIndex int64
	// SIZE: 8
}

type SP_CL2FE_REQ_PC_RECV_EMAIL_PAGE_LIST struct {
	IPageNum int8
	// SIZE: 1
}

type SP_CL2FE_REQ_PC_DELETE_EMAIL struct {
	IEmailIndexArray [5]int64
	// SIZE: 40
}

type SP_CL2FE_REQ_PC_RECV_EMAIL_ITEM struct {
	IEmailIndex    int64
	ISlotNum       int32
	IEmailItemSlot int32
	// SIZE: 16
}

type SP_CL2FE_REQ_PC_RECV_EMAIL_CANDY struct {
	IEmailIndex int64
	// SIZE: 8
}

type SP_CL2FE_GM_REQ_TARGET_PC_SPECIAL_STATE_ONOFF struct {
	ETargetSearchBy      int32
	ITargetPC_ID         int32
	SzTargetPC_FirstName string `size:"10"`
	SzTargetPC_LastName  string `size:"18"`
	ITargetPC_UID        int64
	IONOFF               int32
	ISpecialStateFlag    int8 `pad:"3"`
	// SIZE: 80
}

type SP_CL2FE_REQ_PC_SET_CURRENT_MISSION_ID struct {
	ICurrentMissionID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_NPC_GROUP_INVITE struct {
	INPC_ID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_NPC_GROUP_KICK struct {
	INPC_ID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_FIRST_USE_FLAG_SET struct {
	IFlagCode int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_TRANSPORT_WARP struct {
	ITransport_ID int32
	ILcX          int32
	ILcY          int32
	ILcZ          int32
	// SIZE: 16
}

type SP_CL2FE_REQ_PC_TIME_TO_GO_WARP struct {
	INPC_ID    int32
	IWarpID    int32
	EIL1       int32
	IItemSlot1 int32
	EIL2       int32
	IItemSlot2 int32
	IPC_Level  int32
	IPayFlag   int32
	// SIZE: 32
}

type SP_CL2FE_REQ_PC_RECV_EMAIL_ITEM_ALL struct {
	IEmailIndex int64
	// SIZE: 8
}

type SP_CL2FE_REQ_PC_WARP_CHANNEL struct {
	IChannelNum int32
	IWarpType   int8 `pad:"3"`
	// SIZE: 8
}

type SP_CL2FE_REQ_PC_LOADING_COMPLETE struct {
	IPC_ID int32
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_FIND_NAME_MAKE_BUDDY struct {
	SzFirstName string `size:"9"`
	SzLastName  string `size:"17"`
	// SIZE: 52
}

type SP_CL2FE_REQ_PC_FIND_NAME_ACCEPT_BUDDY struct {
	IAcceptFlag int32
	IBuddyPCUID int64
	SzFirstName string `size:"9"`
	SzLastName  string `size:"17"`
	// SIZE: 64
}

type SP_CL2FE_REQ_PC_ATTACK_CHARs struct {
	ITargetCnt int32
	// SIZE: 4
}

type SP_CL2FE_PC_STREETSTALL_REQ_READY struct {
	IStreetStallItemInvenSlotNum int32
	// SIZE: 4
}

type SP_CL2FE_PC_STREETSTALL_REQ_CANCEL struct {
	IPC_ID int32
	// SIZE: 4
}

type SP_CL2FE_PC_STREETSTALL_REQ_UNREGIST_ITEM struct {
	IItemListNum int32
	// SIZE: 4
}

type SP_CL2FE_PC_STREETSTALL_REQ_SALE_START struct {
	IStreetStallItemInvenSlotNum int32
	// SIZE: 4
}

type SP_CL2FE_PC_STREETSTALL_REQ_ITEM_LIST struct {
	IStreetStallPC_ID int32
	// SIZE: 4
}

type SP_CL2FE_PC_STREETSTALL_REQ_ITEM_BUY struct {
	IStreetStallPC_ID  int32
	IItemListNum       int32
	IEmptyInvenSlotNum int32
	// SIZE: 12
}

type SP_CL2FE_REQ_PC_ITEM_COMBINATION struct {
	ICostumeItemSlot int32
	IStatItemSlot    int32
	ICashItemSlot1   int32
	ICashItemSlot2   int32
	// SIZE: 16
}

type SP_CL2FE_GM_REQ_SET_PC_SKILL struct {
	ISkillSlotNum int32
	ISkillID      int32
	// SIZE: 8
}

type SP_CL2FE_REQ_PC_SKILL_ADD struct {
	ISkillSlotNum          int32
	ISkillID               int32
	ISkillItemInvenSlotNum int32
	IPreSkillSlotNum       int32
	IPreSkillID            int32
	// SIZE: 20
}

type SP_CL2FE_REQ_PC_SKILL_DEL struct {
	ISkillSlotNum int32
	ISkillID      int32
	// SIZE: 8
}

type SP_CL2FE_REQ_PC_SKILL_USE struct {
	ISkillSlotNum    int32
	ISkillID         int32
	IMoveFlag        int32
	IFromX           int32
	IFromY           int32
	IFromZ           int32
	IToX             int32
	IToY             int32
	IToZ             int32
	IMainTargetType  int32
	IMainTargetID    int32
	ITargetLocationX int32
	ITargetLocationY int32
	ITargetLocationZ int32
	ITargetCount     int32
	// SIZE: 60
}

type SP_CL2FE_REQ_PC_ROPE struct {
	ICliTime  uint64
	IX        int32
	IY        int32
	IZ        int32
	FVX       float32
	FVY       float32
	FVZ       float32
	IRopeID   int32
	IAngle    int32
	CKeyValue uint8 `pad:"3"`
	ISpeed    int32
	// SIZE: 48
}

type SP_CL2FE_REQ_PC_BELT struct {
	ICliTime  uint64
	ILcX      int32
	ILcY      int32
	ILcZ      int32
	IX        int32
	IY        int32
	IZ        int32
	FVX       float32
	FVY       float32
	FVZ       float32
	BDown     int32
	IBeltID   int32
	IAngle    int32
	CKeyValue uint8 `pad:"3"`
	ISpeed    int32
	// SIZE: 64
}

type SP_CL2FE_REQ_PC_REGIST_QUICK_SLOT struct {
	ISlotNum  int32
	IItemType int16
	IItemID   int16
	// SIZE: 8
}

type SP_CL2FE_REQ_PC_DISASSEMBLE_ITEM struct {
	IItemSlot int32
	// SIZE: 4
}

type SP_CL2FE_GM_REQ_REWARD_RATE struct {
	IGetSet          int32
	IRewardType      int32
	IRewardRateIndex int32
	ISetRateValue    int32
	// SIZE: 16
}

type SP_CL2FE_REQ_PC_ITEM_ENCHANT struct {
	IEnchantItemSlot         int32
	IWeaponMaterialItemSlot  int32
	IDefenceMaterialItemSlot int32
	ICashItemSlot1           int32
	ICashItemSlot2           int32
	// SIZE: 20
}

type SP_CL2LS_REQ_LOGIN struct {
	SzID            string `size:"33"`
	SzPassword      string `size:"33"`
	IClientVerA     int32
	IClientVerB     int32
	IClientVerC     int32
	ILoginType      int32
	SzCookie_TEGid  [64]byte
	SzCookie_authid [255]byte `pad:"1"`
	// SIZE: 468
}

type SP_CL2LS_REQ_CHECK_CHAR_NAME struct {
	IFNCode     int32
	ILNCode     int32
	IMNCode     int32
	SzFirstName string `size:"9"`
	SzLastName  string `size:"17"`
	// SIZE: 64
}

type SP_CL2LS_REQ_SAVE_CHAR_NAME struct {
	ISlotNum    int8
	IGender     int8 `pad:"2"`
	IFNCode     int32
	ILNCode     int32
	IMNCode     int32
	SzFirstName string `size:"9"`
	SzLastName  string `size:"17"`
	// SIZE: 68
}

type SP_CL2LS_REQ_CHAR_SELECT struct {
	IPC_UID int64
	// SIZE: 8
}

type SP_CL2LS_REQ_CHAR_DELETE struct {
	IPC_UID int64
	// SIZE: 8
}

type SP_CL2LS_REQ_SHARD_SELECT struct {
	ShardNum int8
	// SIZE: 1
}

type SP_CL2LS_CHECK_NAME_LIST struct {
	IFNCode int32
	IMNCode int32
	ILNCode int32
	// SIZE: 12
}

type SP_CL2LS_REQ_SAVE_CHAR_TUTOR struct {
	IPC_UID       int64
	ITutorialFlag int8 `pad:"3"`
	// SIZE: 12
}

type SP_CL2LS_REQ_PC_EXIT_DUPLICATE struct {
	SzID       string `size:"33"`
	SzPassword string `size:"33"`
	// SIZE: 132
}

type SP_CL2LS_REP_LIVE_CHECK struct {
	ITempValue int32
	// SIZE: 4
}

type SP_CL2LS_REQ_CHANGE_CHAR_NAME struct {
	IPCUID      int64
	ISlotNum    int8
	IGender     int8 `pad:"2"`
	IFNCode     int32
	ILNCode     int32
	IMNCode     int32
	SzFirstName string `size:"9"`
	SzLastName  string `size:"17"`
	// SIZE: 76
}

type SP_CL2LS_REQ_SERVER_SELECT struct {
	ServerNum int8
	// SIZE: 1
}

type SPacket struct {
	DwType uint32
	SzData [4096]byte
	// SIZE: 4100
}

type SPacket_Full struct {
	DwSize uint32
	DwType uint32
	SzData [4096]byte
	// SIZE: 4104
}

type SPacket2x struct {
	DwType uint32
	SzData [8192]byte
	// SIZE: 8196
}

type SPacket2x_Full struct {
	DwSize uint32
	DwType uint32
	SzData [8192]byte
	// SIZE: 8200
}

type SP_FE2CL_ERROR struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_ENTER_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_EXIT_FAIL struct {
	IID        int32
	IErrorCode int32
	// SIZE: 8
}

type SP_FE2CL_REP_PC_EXIT_SUCC struct {
	IID       int32
	IExitCode int32
	// SIZE: 8
}

type SP_FE2CL_PC_EXIT struct {
	IID       int32
	IExitType int32
	// SIZE: 8
}

type SP_FE2CL_PC_AROUND struct {
	IPCCnt int32
	// SIZE: 4
}

type SP_FE2CL_PC_MOVE struct {
	ICliTime  uint64
	IX        int32
	IY        int32
	IZ        int32
	FVX       float32
	FVY       float32
	FVZ       float32
	IAngle    int32
	CKeyValue uint8 `pad:"3"`
	ISpeed    int32
	IID       int32
	ISvrTime  uint64
	// SIZE: 56
}

type SP_FE2CL_PC_STOP struct {
	ICliTime uint64
	IX       int32
	IY       int32
	IZ       int32
	IID      int32
	ISvrTime uint64
	// SIZE: 32
}

type SP_FE2CL_PC_JUMP struct {
	ICliTime  uint64
	IX        int32
	IY        int32
	IZ        int32
	IVX       int32
	IVY       int32
	IVZ       int32
	IAngle    int32
	CKeyValue uint8 `pad:"3"`
	ISpeed    int32
	IID       int32
	ISvrTime  uint64
	// SIZE: 56
}

type SP_FE2CL_NPC_EXIT struct {
	INPC_ID int32
	// SIZE: 4
}

type SP_FE2CL_NPC_MOVE struct {
	INPC_ID    int32
	IToX       int32
	IToY       int32
	IToZ       int32
	ISpeed     int32
	IMoveStyle int16 `pad:"2"`
	// SIZE: 24
}

type SP_FE2CL_NPC_AROUND struct {
	INPCCnt int32
	// SIZE: 4
}

type SP_FE2CL_AROUND_DEL_PC struct {
	IPCCnt int32
	// SIZE: 4
}

type SP_FE2CL_AROUND_DEL_NPC struct {
	INPCCnt int32
	// SIZE: 4
}

type SP_FE2CL_REP_SEND_FREECHAT_MESSAGE_SUCC struct {
	IPC_ID     int32
	SzFreeChat string `size:"128"`
	IEmoteCode int32
	// SIZE: 264
}

type SP_FE2CL_REP_SEND_FREECHAT_MESSAGE_FAIL struct {
	IErrorCode int32
	SzFreeChat string `size:"128"`
	IEmoteCode int32
	// SIZE: 264
}

type SP_FE2CL_PC_ATTACK_NPCs_SUCC struct {
	IBatteryW int32
	INPCCnt   int32
	// SIZE: 8
}

type SP_FE2CL_PC_ATTACK_NPCs struct {
	IPC_ID  int32
	INPCCnt int32
	// SIZE: 8
}

type SP_FE2CL_NPC_ATTACK_PCs struct {
	INPC_ID int32
	IPCCnt  int32
	// SIZE: 8
}

type SP_FE2CL_REP_SEND_MENUCHAT_MESSAGE_SUCC struct {
	IPC_ID     int32
	SzFreeChat string `size:"128"`
	IEmoteCode int32
	// SIZE: 264
}

type SP_FE2CL_REP_SEND_MENUCHAT_MESSAGE_FAIL struct {
	IErrorCode int32
	SzFreeChat string `size:"128"`
	IEmoteCode int32
	// SIZE: 264
}

type SP_FE2CL_REP_PC_TASK_START_SUCC struct {
	ITaskNum    int32
	IRemainTime int32
	// SIZE: 8
}

type SP_FE2CL_REP_PC_TASK_START_FAIL struct {
	ITaskNum   int32
	IErrorCode int32
	// SIZE: 8
}

type SP_FE2CL_REP_PC_TASK_END_SUCC struct {
	ITaskNum int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_TASK_END_FAIL struct {
	ITaskNum   int32
	IErrorCode int32
	// SIZE: 8
}

type SP_FE2CL_NPC_SKILL_READY struct {
	INPC_ID  int32
	ISkillID int16 `pad:"2"`
	IValue1  int32
	IValue2  int32
	IValue3  int32
	// SIZE: 20
}

type SP_FE2CL_NPC_SKILL_FIRE struct {
	INPC_ID  int32
	ISkillID int16 `pad:"2"`
	IVX      int32
	IVY      int32
	IVZ      int32
	// SIZE: 20
}

type SP_FE2CL_NPC_SKILL_HIT struct {
	INPC_ID    int32
	ISkillID   int16 `pad:"2"`
	IValue1    int32
	IValue2    int32
	IValue3    int32
	EST        int32
	ITargetCnt int32
	// SIZE: 28
}

type SP_FE2CL_NPC_SKILL_CORRUPTION_READY struct {
	INPC_ID  int32
	ISkillID int16
	IStyle   int16
	IValue1  int32
	IValue2  int32
	IValue3  int32
	// SIZE: 20
}

type SP_FE2CL_NPC_SKILL_CORRUPTION_HIT struct {
	INPC_ID    int32
	ISkillID   int16
	IStyle     int16
	IValue1    int32
	IValue2    int32
	IValue3    int32
	ITargetCnt int32
	// SIZE: 24
}

type SP_FE2CL_NPC_SKILL_CANCEL struct {
	INPC_ID int32
	// SIZE: 4
}

type SP_FE2CL_REP_NANO_EQUIP_SUCC struct {
	INanoID       int16
	INanoSlotNum  int16
	BNanoDeactive int32
	// SIZE: 8
}

type SP_FE2CL_REP_NANO_UNEQUIP_SUCC struct {
	INanoSlotNum  int16 `pad:"2"`
	BNanoDeactive int32
	// SIZE: 8
}

type SP_FE2CL_REP_NANO_ACTIVE_SUCC struct {
	IActiveNanoSlotNum int16 `pad:"2"`
	ECSTB___Add        int32
	// SIZE: 8
}

type SP_FE2CL_NANO_SKILL_USE_SUCC struct {
	IPC_ID        int32
	IBulletID     int8 `pad:"1"`
	ISkillID      int16
	IArg1         int32
	IArg2         int32
	IArg3         int32
	BNanoDeactive int32
	INanoID       int16
	INanoStamina  int16
	EST           int32
	ITargetCnt    int32
	// SIZE: 36
}

type SP_FE2CL_NANO_SKILL_USE struct {
	IPC_ID        int32
	IBulletID     int8 `pad:"1"`
	ISkillID      int16
	IArg1         int32
	IArg2         int32
	IArg3         int32
	BNanoDeactive int32
	INanoID       int16
	INanoStamina  int16
	EST           int32
	ITargetCnt    int32
	// SIZE: 36
}

type SP_FE2CL_REP_PC_TASK_STOP_SUCC struct {
	ITaskNum int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_TASK_STOP_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_TASK_CONTINUE_SUCC struct {
	ITaskNum int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_TASK_CONTINUE_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_GOTO_SUCC struct {
	IX int32
	IY int32
	IZ int32
	// SIZE: 12
}

type SP_FE2CL_REP_CHARGE_NANO_STAMINA struct {
	IBatteryN    int32
	INanoID      int16
	INanoStamina int16
	// SIZE: 8
}

type SP_FE2CL_REP_PC_KILL_QUEST_NPCs_SUCC struct {
	INPCID int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_VENDOR_ITEM_BUY_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_VENDOR_ITEM_SELL_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_ITEM_DELETE_SUCC struct {
	EIL      int32
	ISlotNum int32
	// SIZE: 8
}

type SP_FE2CL_PC_ROCKET_STYLE_READY struct {
	IPC_ID   int32
	ISkillID int32
	// SIZE: 8
}

type SP_FE2CL_PC_GRENADE_STYLE_READY struct {
	IPC_ID   int32
	ISkillID int32
	// SIZE: 8
}

type SP_FE2CL_REP_PC_TRADE_OFFER struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	// SIZE: 12
}

type SP_FE2CL_REP_PC_TRADE_OFFER_CANCEL struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	// SIZE: 12
}

type SP_FE2CL_REP_PC_TRADE_OFFER_SUCC struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	// SIZE: 12
}

type SP_FE2CL_REP_PC_TRADE_OFFER_REFUSAL struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	// SIZE: 12
}

type SP_FE2CL_REP_PC_TRADE_OFFER_ABORT struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	IErrorCode  int16 `pad:"2"`
	// SIZE: 16
}

type SP_FE2CL_REP_PC_TRADE_CONFIRM struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	// SIZE: 12
}

type SP_FE2CL_REP_PC_TRADE_CONFIRM_CANCEL struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	// SIZE: 12
}

type SP_FE2CL_REP_PC_TRADE_CONFIRM_ABORT struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	// SIZE: 12
}

type SP_FE2CL_REP_PC_TRADE_CONFIRM_FAIL struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	IErrorCode  int32
	// SIZE: 16
}

type SP_FE2CL_REP_PC_TRADE_ITEM_REGISTER_FAIL struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	IErrorCode  int32
	// SIZE: 16
}

type SP_FE2CL_REP_PC_TRADE_ITEM_UNREGISTER_FAIL struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	IErrorCode  int32
	// SIZE: 16
}

type SP_FE2CL_REP_PC_TRADE_CASH_REGISTER_SUCC struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	ICandy      int32
	// SIZE: 16
}

type SP_FE2CL_REP_PC_TRADE_CASH_REGISTER_FAIL struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	IErrorCode  int32
	// SIZE: 16
}

type SP_FE2CL_REP_PC_TRADE_EMOTES_CHAT struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	SzFreeChat  string `size:"128"`
	IEmoteCode  int32
	// SIZE: 272
}

type SP_FE2CL_REP_PC_NANO_CREATE_FAIL struct {
	IPC_ID     int32
	IErrorCode int32
	// SIZE: 8
}

type SP_FE2CL_REP_NANO_TUNE_FAIL struct {
	IPC_ID     int32
	IErrorCode int32
	// SIZE: 8
}

type SP_FE2CL_REP_PC_BANK_OPEN_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_BANK_CLOSE_SUCC struct {
	IPC_ID int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_BANK_CLOSE_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_VENDOR_START_SUCC struct {
	INPC_ID   int32
	IVendorID int32
	// SIZE: 8
}

type SP_FE2CL_REP_PC_VENDOR_START_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_VENDOR_TABLE_UPDATE_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_VENDOR_ITEM_RESTORE_BUY_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_CHAR_TIME_BUFF_TIME_OUT struct {
	ECT               int32
	IID               int32
	IConditionBitFlag int32
	// SIZE: 12
}

type SP_FE2CL_REP_PC_GIVE_ITEM_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_BUDDYLIST_INFO_SUCC struct {
	IID       int32
	IPCUID    int64
	IListNum  int8
	IBuddyCnt int8 `pad:"2"`
	// SIZE: 16
}

type SP_FE2CL_REP_PC_BUDDYLIST_INFO_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_REQUEST_MAKE_BUDDY_SUCC struct {
	IRequestID  int32
	IBuddyID    int32
	IBuddyPCUID int64
	// SIZE: 16
}

type SP_FE2CL_REP_REQUEST_MAKE_BUDDY_FAIL struct {
	IBuddyID    int32
	IBuddyPCUID int64
	IErrorCode  int32
	// SIZE: 16
}

type SP_FE2CL_REP_ACCEPT_MAKE_BUDDY_FAIL struct {
	IBuddyID    int32
	IBuddyPCUID int64
	IErrorCode  int32
	// SIZE: 16
}

type SP_FE2CL_REP_SEND_BUDDY_FREECHAT_MESSAGE_SUCC struct {
	IFromPCUID int64
	IToPCUID   int64
	SzFreeChat string `size:"128"`
	IEmoteCode int32
	// SIZE: 276
}

type SP_FE2CL_REP_SEND_BUDDY_FREECHAT_MESSAGE_FAIL struct {
	IErrorCode int32
	IToPCUID   int64
	SzFreeChat string `size:"128"`
	IEmoteCode int32
	// SIZE: 272
}

type SP_FE2CL_REP_SEND_BUDDY_MENUCHAT_MESSAGE_SUCC struct {
	IFromPCUID int64
	IToPCUID   int64
	SzFreeChat string `size:"128"`
	IEmoteCode int32
	// SIZE: 276
}

type SP_FE2CL_REP_SEND_BUDDY_MENUCHAT_MESSAGE_FAIL struct {
	IErrorCode int32
	IToPCUID   int64
	SzFreeChat string `size:"128"`
	IEmoteCode int32
	// SIZE: 272
}

type SP_FE2CL_REP_GET_BUDDY_STYLE_FAIL struct {
	IErrorCode  int32
	IBuddyPCUID int64
	// SIZE: 12
}

type SP_FE2CL_REP_GET_BUDDY_STATE_SUCC struct {
	ABuddyID    [50]int32
	ABuddyState [50]byte `pad:"2"`
	// SIZE: 252
}

type SP_FE2CL_REP_GET_BUDDY_STATE_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_SET_BUDDY_BLOCK_SUCC struct {
	IBuddyPCUID int64
	IBuddySlot  int8 `pad:"3"`
	// SIZE: 12
}

type SP_FE2CL_REP_SET_BUDDY_BLOCK_FAIL struct {
	IBuddyPCUID int64
	IErrorCode  int32
	// SIZE: 12
}

type SP_FE2CL_REP_REMOVE_BUDDY_SUCC struct {
	IBuddyPCUID int64
	IBuddySlot  int8 `pad:"3"`
	// SIZE: 12
}

type SP_FE2CL_REP_REMOVE_BUDDY_FAIL struct {
	IBuddyPCUID int64
	IErrorCode  int32
	// SIZE: 12
}

type SP_FE2CL_PC_JUMPPAD struct {
	ICliTime  uint64
	IX        int32
	IY        int32
	IZ        int32
	IVX       int32
	IVY       int32
	IVZ       int32
	IAngle    int32
	CKeyValue uint8 `pad:"3"`
	IPC_ID    int32
	ISvrTime  uint64
	// SIZE: 52
}

type SP_FE2CL_PC_LAUNCHER struct {
	ICliTime uint64
	IX       int32
	IY       int32
	IZ       int32
	IVX      int32
	IVY      int32
	IVZ      int32
	IAngle   int32
	ISpeed   int32
	IPC_ID   int32
	ISvrTime uint64
	// SIZE: 52
}

type SP_FE2CL_PC_ZIPLINE struct {
	ICliTime     uint64
	IStX         int32
	IStY         int32
	IStZ         int32
	FMovDistance float32
	FMaxDistance float32
	FDummy       float32
	IX           int32
	IY           int32
	IZ           int32
	FVX          float32
	FVY          float32
	FVZ          float32
	BDown        int32
	IRollMax     int32
	IRoll        uint8 `pad:"3"`
	IAngle       int32
	ISpeed       int32
	IPC_ID       int32
	ISvrTime     uint64
	// SIZE: 88
}

type SP_FE2CL_PC_MOVEPLATFORM struct {
	ICliTime    uint64
	ILcX        int32
	ILcY        int32
	ILcZ        int32
	IX          int32
	IY          int32
	IZ          int32
	FVX         float32
	FVY         float32
	FVZ         float32
	BDown       int32
	IPlatformID int32
	IAngle      int32
	CKeyValue   uint8 `pad:"3"`
	ISpeed      int32
	IPC_ID      int32
	ISvrTime    uint64
	// SIZE: 76
}

type SP_FE2CL_PC_SLOPE struct {
	ICliTime  uint64
	IX        int32
	IY        int32
	IZ        int32
	IAngle    int32
	ISpeed    int32
	CKeyValue uint8 `pad:"3"`
	IPC_ID    int32
	ISvrTime  uint64
	FVX       float32
	FVY       float32
	FVZ       float32
	ISlopeID  int32
	// SIZE: 60
}

type SP_FE2CL_PC_STATE_CHANGE struct {
	IPC_ID int32
	IState int8 `pad:"3"`
	// SIZE: 8
}

type SP_FE2CL_REP_REQUEST_MAKE_BUDDY_SUCC_TO_ACCEPTER struct {
	IRequestID  int32
	IBuddyID    int32
	SzFirstName string `size:"9"`
	SzLastName  string `size:"17"`
	// SIZE: 60
}

type SP_FE2CL_REP_REWARD_ITEM struct {
	M_iCandy        int32
	M_iFusionMatter int32
	M_iBatteryN     int32
	M_iBatteryW     int32
	IItemCnt        int8 `pad:"3"`
	IFatigue        int32
	IFatigue_Level  int32
	INPC_TypeID     int32
	ITaskID         int32
	// SIZE: 36
}

type SP_FE2CL_REP_ITEM_CHEST_OPEN_SUCC struct {
	ISlotNum int32
	// SIZE: 4
}

type SP_FE2CL_REP_ITEM_CHEST_OPEN_FAIL struct {
	ISlotNum   int32
	IErrorCode int32
	// SIZE: 8
}

type SP_FE2CL_CHAR_TIME_BUFF_TIME_TICK struct {
	ECT    int32
	IID    int32
	ITB_ID int16 `pad:"2"`
	// SIZE: 12
}

type SP_FE2CL_REP_PC_VENDOR_BATTERY_BUY_SUCC struct {
	ICandy    int32
	IBatteryW int32
	IBatteryN int32
	// SIZE: 12
}

type SP_FE2CL_REP_PC_VENDOR_BATTERY_BUY_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_CHARACTER_ATTACK_CHARACTERs struct {
	ECT          int32
	ICharacterID int32
	ITargetCnt   int32
	// SIZE: 12
}

type SP_FE2CL_PC_GROUP_INVITE struct {
	IHostID int32
	// SIZE: 4
}

type SP_FE2CL_PC_GROUP_INVITE_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_PC_GROUP_INVITE_REFUSE struct {
	IID_To int32
	// SIZE: 4
}

type SP_FE2CL_PC_GROUP_JOIN struct {
	IID_NewMember int32
	IMemberPCCnt  int32
	IMemberNPCCnt int32
	// SIZE: 12
}

type SP_FE2CL_PC_GROUP_JOIN_FAIL struct {
	IID        int32
	IErrorCode int32
	// SIZE: 8
}

type SP_FE2CL_PC_GROUP_JOIN_SUCC struct {
	IID_NewMember int32
	IMemberPCCnt  int32
	IMemberNPCCnt int32
	// SIZE: 12
}

type SP_FE2CL_PC_GROUP_LEAVE struct {
	IID_LeaveMember int32
	IMemberPCCnt    int32
	IMemberNPCCnt   int32
	// SIZE: 12
}

type SP_FE2CL_PC_GROUP_LEAVE_FAIL struct {
	IID        int32
	IErrorCode int32
	// SIZE: 8
}

type SP_FE2CL_PC_GROUP_MEMBER_INFO struct {
	IID           int32
	IMemberPCCnt  int32
	IMemberNPCCnt int32
	// SIZE: 12
}

type SP_FE2CL_REP_PC_WARP_USE_NPC_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_AVATAR_EMOTES_CHAT struct {
	IID_From   int32
	IEmoteCode int32
	// SIZE: 8
}

type SP_FE2CL_REP_PC_CHANGE_MENTOR_SUCC struct {
	IMentor       int16
	IMentorCnt    int16
	IFusionMatter int32
	// SIZE: 8
}

type SP_FE2CL_REP_PC_CHANGE_MENTOR_FAIL struct {
	IMentor    int16 `pad:"2"`
	IErrorCode int32
	// SIZE: 8
}

type SP_FE2CL_REP_GET_MEMBER_STYLE_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_GET_GROUP_STYLE_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_GET_GROUP_STYLE_SUCC struct {
	IMemberCnt int32
	// SIZE: 4
}

type SP_FE2CL_INSTANCE_MAP_INFO struct {
	IInstanceMapNum        int32
	ICreateTick            uint64
	IMapCoordX_Min         int32
	IMapCoordX_Max         int32
	IMapCoordY_Min         int32
	IMapCoordY_Max         int32
	IMapCoordZ_Min         int32
	IMapCoordZ_Max         int32
	IEP_ID                 int32
	IEPTopRecord_Score     int32
	IEPTopRecord_Rank      int32
	IEPTopRecord_Time      int32
	IEPTopRecord_RingCount int32
	IEPSwitch_StatusON_Cnt int32
	// SIZE: 60
}

type SP_FE2CL_TRANSPORTATION_EXIT struct {
	ETT   int32
	IT_ID int32
	// SIZE: 8
}

type SP_FE2CL_TRANSPORTATION_MOVE struct {
	ETT        int32
	IT_ID      int32
	IToX       int32
	IToY       int32
	IToZ       int32
	ISpeed     int32
	IMoveStyle int16 `pad:"2"`
	// SIZE: 28
}

type SP_FE2CL_TRANSPORTATION_AROUND struct {
	ICnt int32
	// SIZE: 4
}

type SP_FE2CL_AROUND_DEL_TRANSPORTATION struct {
	ETT  int32
	ICnt int32
	// SIZE: 8
}

type SP_FE2CL_REP_EP_RACE_START_SUCC struct {
	IStartTick uint64
	ILimitTime int32
	// SIZE: 12
}

type SP_FE2CL_REP_EP_RACE_START_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_EP_RACE_END_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_EP_RACE_CANCEL_SUCC struct {
	ITemp int32
	// SIZE: 4
}

type SP_FE2CL_REP_EP_RACE_CANCEL_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_EP_GET_RING_SUCC struct {
	IRingLID       int32
	IRingCount_Get int32
	// SIZE: 8
}

type SP_FE2CL_REP_EP_GET_RING_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_IM_CHANGE_SWITCH_STATUS struct {
	IMapNum       int32
	ISwitchLID    int32
	ISwitchGID    int32
	ISwitchStatus int32
	// SIZE: 16
}

type SP_FE2CL_SHINY_EXIT struct {
	IShinyID int32
	// SIZE: 4
}

type SP_FE2CL_SHINY_AROUND struct {
	IShinyCnt int32
	// SIZE: 4
}

type SP_FE2CL_AROUND_DEL_SHINY struct {
	IShinyCnt int32
	// SIZE: 4
}

type SP_FE2CL_REP_SHINY_PICKUP_SUCC struct {
	ISkillID int16 `pad:"2"`
	ECSTB    int32
	// SIZE: 8
}

type SP_FE2CL_PC_MOVETRANSPORTATION struct {
	ICliTime  uint64
	ILcX      int32
	ILcY      int32
	ILcZ      int32
	IX        int32
	IY        int32
	IZ        int32
	FVX       float32
	FVY       float32
	FVZ       float32
	IT_ID     int32
	IAngle    int32
	CKeyValue uint8 `pad:"3"`
	ISpeed    int32
	IPC_ID    int32
	ISvrTime  uint64
	// SIZE: 72
}

type SP_FE2CL_REP_SEND_ALL_GROUP_FREECHAT_MESSAGE_SUCC struct {
	ISendPCID  int32
	SzFreeChat string `size:"128"`
	IEmoteCode int32
	// SIZE: 264
}

type SP_FE2CL_REP_SEND_ALL_GROUP_FREECHAT_MESSAGE_FAIL struct {
	ISendPCID  int32
	SzFreeChat string `size:"128"`
	IEmoteCode int32
	IErrorCode int32
	// SIZE: 268
}

type SP_FE2CL_REP_SEND_ANY_GROUP_FREECHAT_MESSAGE_SUCC struct {
	ISendPCID   int32
	IGroupPC_ID int32
	SzFreeChat  string `size:"128"`
	IEmoteCode  int32
	// SIZE: 268
}

type SP_FE2CL_REP_SEND_ANY_GROUP_FREECHAT_MESSAGE_FAIL struct {
	ISendPCID   int32
	IGroupPC_ID int32
	SzFreeChat  string `size:"128"`
	IEmoteCode  int32
	IErrorCode  int32
	// SIZE: 272
}

type SP_FE2CL_REP_BARKER struct {
	INPC_ID          int32
	IMissionStringID int32
	// SIZE: 8
}

type SP_FE2CL_REP_SEND_ALL_GROUP_MENUCHAT_MESSAGE_SUCC struct {
	ISendPCID  int32
	SzFreeChat string `size:"128"`
	IEmoteCode int32
	// SIZE: 264
}

type SP_FE2CL_REP_SEND_ALL_GROUP_MENUCHAT_MESSAGE_FAIL struct {
	ISendPCID  int32
	SzFreeChat string `size:"128"`
	IEmoteCode int32
	IErrorCode int32
	// SIZE: 268
}

type SP_FE2CL_REP_SEND_ANY_GROUP_MENUCHAT_MESSAGE_SUCC struct {
	ISendPCID   int32
	IGroupPC_ID int32
	SzFreeChat  string `size:"128"`
	IEmoteCode  int32
	// SIZE: 268
}

type SP_FE2CL_REP_SEND_ANY_GROUP_MENUCHAT_MESSAGE_FAIL struct {
	ISendPCID   int32
	IGroupPC_ID int32
	SzFreeChat  string `size:"128"`
	IEmoteCode  int32
	IErrorCode  int32
	// SIZE: 272
}

type SP_FE2CL_REP_PC_REGIST_TRANSPORTATION_LOCATION_FAIL struct {
	ETT         int32
	ILocationID int32
	IErrorCode  int32
	// SIZE: 12
}

type SP_FE2CL_REP_PC_REGIST_TRANSPORTATION_LOCATION_SUCC struct {
	ETT                 int32
	ILocationID         int32
	IWarpLocationFlag   int32
	AWyvernLocationFlag [2]int64
	// SIZE: 28
}

type SP_FE2CL_REP_PC_WARP_USE_TRANSPORTATION_FAIL struct {
	ITransportationID int32
	IErrorCode        int32
	// SIZE: 8
}

type SP_FE2CL_REP_PC_WARP_USE_TRANSPORTATION_SUCC struct {
	ETT    int32
	IX     int32
	IY     int32
	IZ     int32
	ICandy int32
	// SIZE: 20
}

type SP_FE2CL_ANNOUNCE_MSG struct {
	IAnnounceType int8 `pad:"3"`
	IDuringTime   int32
	SzAnnounceMsg string `size:"512"`
	// SIZE: 1032
}

type SP_FE2CL_REP_PC_SPECIAL_STATE_SWITCH_SUCC struct {
	IPC_ID               int32
	IReqSpecialStateFlag int8
	ISpecialState        int8 `pad:"2"`
	// SIZE: 8
}

type SP_FE2CL_PC_SPECIAL_STATE_CHANGE struct {
	IPC_ID               int32
	IReqSpecialStateFlag int8
	ISpecialState        int8 `pad:"2"`
	// SIZE: 8
}

type SP_FE2CL_GM_REP_PC_SET_VALUE struct {
	IPC_ID        int32
	ISetValueType int32
	ISetValue     int32
	// SIZE: 12
}

type SP_FE2CL_GM_PC_CHANGE_VALUE struct {
	IPC_ID        int32
	ISetValueType int32
	ISetValue     int32
	// SIZE: 12
}

type SP_FE2CL_GM_REP_PC_LOCATION struct {
	ITargetPC_UID        int64
	ITargetPC_ID         int32
	IShardID             int32
	IMapType             int32
	IMapID               int32
	IMapNum              int32
	IX                   int32
	IY                   int32
	IZ                   int32
	SzTargetPC_FirstName string `size:"10"`
	SzTargetPC_LastName  string `size:"18"`
	// SIZE: 96
}

type SP_FE2CL_GM_REP_PC_ANNOUNCE struct {
	IAnnounceType int8 `pad:"3"`
	IDuringTime   int32
	SzAnnounceMsg string `size:"512"`
	// SIZE: 1032
}

type SP_FE2CL_REP_PC_BUDDY_WARP_FAIL struct {
	IBuddyPCUID int64
	IErrorCode  int32
	// SIZE: 12
}

type SP_FE2CL_REP_PC_CHANGE_LEVEL struct {
	IPC_ID    int32
	IPC_Level int16 `pad:"2"`
	// SIZE: 8
}

type SP_FE2CL_REP_SET_PC_BLOCK_SUCC struct {
	IBlock_ID    int32
	IBlock_PCUID int64
	IBuddySlot   int8 `pad:"3"`
	// SIZE: 16
}

type SP_FE2CL_REP_SET_PC_BLOCK_FAIL struct {
	IBlock_ID    int32
	IBlock_PCUID int64
	IErrorCode   int32
	// SIZE: 16
}

type SP_FE2CL_REP_REGIST_RXCOM struct {
	IMapNum int32
	IX      int32
	IY      int32
	IZ      int32
	// SIZE: 16
}

type SP_FE2CL_REP_REGIST_RXCOM_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_PC_INVEN_FULL_MSG struct {
	IType      int8 `pad:"3"`
	IErrorCode int32
	// SIZE: 8
}

type SP_FE2CL_REQ_LIVE_CHECK struct {
	ITempValue int32
	// SIZE: 4
}

type SP_FE2CL_PC_MOTD_LOGIN struct {
	IType       int8   `pad:"1"`
	SzSystemMsg string `size:"512"`
	// SIZE: 1026
}

type SP_FE2CL_REP_PC_ITEM_USE_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_PC_ITEM_USE struct {
	IPC_ID     int32
	ISkillID   int16 `pad:"2"`
	EST        int32
	ITargetCnt int32
	// SIZE: 16
}

type SP_FE2CL_REP_GET_BUDDY_LOCATION_SUCC struct {
	IBuddyPCUID int64
	ISlotNum    int8 `pad:"3"`
	IX          int32
	IY          int32
	IZ          int32
	IShardNum   int8 `pad:"3"`
	// SIZE: 28
}

type SP_FE2CL_REP_GET_BUDDY_LOCATION_FAIL struct {
	IBuddyPCUID int64
	ISlotNum    int8 `pad:"3"`
	IErrorCode  int32
	// SIZE: 16
}

type SP_FE2CL_REP_PC_RIDING_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_RIDING_SUCC struct {
	IPC_ID int32
	ERT    int32
	// SIZE: 8
}

type SP_FE2CL_PC_RIDING struct {
	IPC_ID int32
	ERT    int32
	// SIZE: 8
}

type SP_FE2CL_PC_BROOMSTICK_MOVE struct {
	IPC_ID int32
	IToX   int32
	IToY   int32
	IToZ   int32
	ISpeed int32
	// SIZE: 20
}

type SP_FE2CL_REP_PC_BUDDY_WARP_OTHER_SHARD_SUCC struct {
	IBuddyPCUID int64
	IShardNum   int8 `pad:"3"`
	IChannelNum int32
	// SIZE: 16
}

type SP_FE2CL_REP_WARP_USE_RECALL_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_EXIT_DUPLICATE struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_MISSION_COMPLETE_SUCC struct {
	IMissionNum int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_NEW_EMAIL struct {
	INewEmailCnt int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_READ_EMAIL_FAIL struct {
	IEmailIndex int64
	IErrorCode  int32
	// SIZE: 12
}

type SP_FE2CL_REP_PC_RECV_EMAIL_PAGE_LIST_FAIL struct {
	IPageNum   int8 `pad:"3"`
	IErrorCode int32
	// SIZE: 8
}

type SP_FE2CL_REP_PC_DELETE_EMAIL_SUCC struct {
	IEmailIndexArray [5]int64
	// SIZE: 40
}

type SP_FE2CL_REP_PC_DELETE_EMAIL_FAIL struct {
	IEmailIndexArray [5]int64
	IErrorCode       int32
	// SIZE: 44
}

type SP_FE2CL_REP_PC_SEND_EMAIL_FAIL struct {
	ITo_PCUID  int64
	IErrorCode int32
	// SIZE: 12
}

type SP_FE2CL_REP_PC_RECV_EMAIL_ITEM_SUCC struct {
	IEmailIndex    int64
	ISlotNum       int32
	IEmailItemSlot int32
	// SIZE: 16
}

type SP_FE2CL_REP_PC_RECV_EMAIL_ITEM_FAIL struct {
	IEmailIndex    int64
	ISlotNum       int32
	IEmailItemSlot int32
	IErrorCode     int32
	// SIZE: 20
}

type SP_FE2CL_REP_PC_RECV_EMAIL_CANDY_SUCC struct {
	IEmailIndex int64
	ICandy      int32
	// SIZE: 12
}

type SP_FE2CL_REP_PC_RECV_EMAIL_CANDY_FAIL struct {
	IEmailIndex int64
	IErrorCode  int32
	// SIZE: 12
}

type SP_FE2CL_PC_SUDDEN_DEAD struct {
	IPC_ID            int32
	ISuddenDeadReason int32
	IDamage           int32
	IHP               int32
	// SIZE: 16
}

type SP_FE2CL_REP_GM_REQ_TARGET_PC_SPECIAL_STATE_ONOFF_SUCC struct {
	ITargetPC_ID         int32
	SzTargetPC_FirstName string `size:"10"`
	SzTargetPC_LastName  string `size:"18"`
	IReqSpecialStateFlag int8
	ISpecialState        int8 `pad:"2"`
	// SIZE: 64
}

type SP_FE2CL_REP_PC_SET_CURRENT_MISSION_ID struct {
	ICurrentMissionID int32
	// SIZE: 4
}

type SP_FE2CL_REP_NPC_GROUP_INVITE_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_NPC_GROUP_INVITE_SUCC struct {
	IPC_ID        int32
	INPC_ID       int32
	IMemberPCCnt  int32
	IMemberNPCCnt int32
	// SIZE: 16
}

type SP_FE2CL_REP_NPC_GROUP_KICK_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_NPC_GROUP_KICK_SUCC struct {
	IPC_ID        int32
	INPC_ID       int32
	IMemberPCCnt  int32
	IMemberNPCCnt int32
	// SIZE: 16
}

type SP_FE2CL_PC_EVENT struct {
	IPC_ID       int32
	IEventID     int32
	IEventValue1 int32
	IEventValue2 int32
	IEventValue3 int32
	// SIZE: 20
}

type SP_FE2CL_REP_PC_TRADE_EMOTES_CHAT_FAIL struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	SzFreeChat  string `size:"128"`
	IEmoteCode  int32
	IErrorCode  int32
	// SIZE: 276
}

type SP_FE2CL_REP_PC_RECV_EMAIL_ITEM_ALL_SUCC struct {
	IEmailIndex int64
	// SIZE: 8
}

type SP_FE2CL_REP_PC_RECV_EMAIL_ITEM_ALL_FAIL struct {
	IEmailIndex int64
	IErrorCode  int32
	// SIZE: 12
}

type SP_FE2CL_REP_PC_LOADING_COMPLETE_SUCC struct {
	IPC_ID int32
	// SIZE: 4
}

type SChannelInfo struct {
	IChannelNum     int32
	ICurrentUserCnt int32
	// SIZE: 8
}

type SP_FE2CL_REP_CHANNEL_INFO struct {
	ICurrChannelNum int32
	IChannelCnt     int32
	// SIZE: 8
}

type SP_FE2CL_REP_PC_CHANNEL_NUM struct {
	IChannelNum int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_WARP_CHANNEL_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_FIND_NAME_MAKE_BUDDY_SUCC struct {
	SzFirstName    string `size:"9"`
	SzLastName     string `size:"17"`
	IPCUID         int64
	INameCheckFlag int8 `pad:"3"`
	// SIZE: 64
}

type SP_FE2CL_REP_PC_FIND_NAME_MAKE_BUDDY_FAIL struct {
	SzFirstName string `size:"9"`
	SzLastName  string `size:"17"`
	IErrorCode  int32
	// SIZE: 56
}

type SP_FE2CL_REP_PC_FIND_NAME_ACCEPT_BUDDY_FAIL struct {
	SzFirstName    string `size:"9"`
	SzLastName     string `size:"17"`
	IPCUID         int64
	INameCheckFlag int8 `pad:"3"`
	IErrorCode     int32
	// SIZE: 68
}

type SP_FE2CL_PC_ATTACK_CHARs_SUCC struct {
	IBatteryW  int32
	ITargetCnt int32
	// SIZE: 8
}

type SP_FE2CL_PC_ATTACK_CHARs struct {
	IPC_ID     int32
	ITargetCnt int32
	// SIZE: 8
}

type SP_FE2CL_NPC_ATTACK_CHARs struct {
	INPC_ID    int32
	ITargetCnt int32
	// SIZE: 8
}

type SP_FE2CL_REP_PC_CHANGE_LEVEL_SUCC struct {
	ILevel        int32
	IFusionMatter int32
	// SIZE: 8
}

type SP_FE2CL_REP_PC_NANO_CREATE struct {
	IPC_ID  int32
	INanoID int16 `pad:"2"`
	// SIZE: 8
}

type SP_FE2CL_PC_STREETSTALL_REP_READY_SUCC struct {
	IStreetStallItemInvenSlotNum int32
	IItemListCountMax            int32
	FTaxPercentage               float32
	IPCCharState                 int8 `pad:"3"`
	// SIZE: 16
}

type SP_FE2CL_PC_STREETSTALL_REP_READY_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_PC_STREETSTALL_REP_CANCEL_SUCC struct {
	IPCCharState int8
	// SIZE: 1
}

type SP_FE2CL_PC_STREETSTALL_REP_CANCEL_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_PC_STREETSTALL_REP_REGIST_ITEM_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_PC_STREETSTALL_REP_UNREGIST_ITEM_SUCC struct {
	IItemListNum int32
	// SIZE: 4
}

type SP_FE2CL_PC_STREETSTALL_REP_UNREGIST_ITEM_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_PC_STREETSTALL_REP_SALE_START_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_PC_STREETSTALL_REP_ITEM_LIST struct {
	IStreetStallPC_ID int32
	IItemListCount    int32
	// SIZE: 8
}

type SP_FE2CL_PC_STREETSTALL_REP_ITEM_LIST_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_PC_STREETSTALL_REP_ITEM_BUY_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_ITEM_COMBINATION_FAIL struct {
	IErrorCode       int32
	ICostumeItemSlot int32
	IStatItemSlot    int32
	ICashItemSlot1   int32
	ICashItemSlot2   int32
	// SIZE: 20
}

type SP_FE2CL_REP_PC_SKILL_ADD_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_SKILL_DEL_SUCC struct {
	ISkillSlotNum int32
	ISkillID      int32
	// SIZE: 8
}

type SP_FE2CL_REP_PC_SKILL_DEL_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_SKILL_USE_SUCC struct {
	IPC_ID           int32
	ISkillSlotNum    int32
	ISkillID         int32
	IX               int32
	IY               int32
	IZ               int32
	IAngle           int32
	IBlockMove       int32
	EST              int32
	ITargetID        int32
	ITargetType      int32
	ITargetLocationX int32
	ITargetLocationY int32
	ITargetLocationZ int32
	ITargetCnt       int32
	// SIZE: 60
}

type SP_FE2CL_REP_PC_SKILL_USE_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_PC_SKILL_USE struct {
	IPC_ID           int32
	ISkillSlotNum    int32
	ISkillID         int32
	IX               int32
	IY               int32
	IZ               int32
	IAngle           int32
	IBlockMove       int32
	EST              int32
	ITargetID        int32
	ITargetType      int32
	ITargetLocationX int32
	ITargetLocationY int32
	ITargetLocationZ int32
	ITargetCnt       int32
	// SIZE: 60
}

type SP_FE2CL_PC_ROPE struct {
	ICliTime  uint64
	IX        int32
	IY        int32
	IZ        int32
	FVX       float32
	FVY       float32
	FVZ       float32
	IRopeID   int32
	IAngle    int32
	CKeyValue uint8 `pad:"3"`
	ISpeed    int32
	IPC_ID    int32
	ISvrTime  uint64
	// SIZE: 60
}

type SP_FE2CL_PC_BELT struct {
	ICliTime  uint64
	ILcX      int32
	ILcY      int32
	ILcZ      int32
	IX        int32
	IY        int32
	IZ        int32
	FVX       float32
	FVY       float32
	FVZ       float32
	BDown     int32
	IBeltID   int32
	IAngle    int32
	CKeyValue uint8 `pad:"3"`
	ISpeed    int32
	IPC_ID    int32
	ISvrTime  uint64
	// SIZE: 76
}

type SP_FE2CL_PC_VEHICLE_ON_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_PC_VEHICLE_OFF_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_REGIST_QUICK_SLOT_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_REGIST_QUICK_SLOT_SUCC struct {
	ISlotNum  int32
	IItemType int16
	IItemID   int16
	// SIZE: 8
}

type SP_FE2CL_PC_DELETE_TIME_LIMIT_ITEM struct {
	IItemListCount int32
	// SIZE: 4
}

type SP_FE2CL_REP_PC_DISASSEMBLE_ITEM_FAIL struct {
	IErrorCode int32
	IItemSlot  int32
	// SIZE: 8
}

type SP_FE2CL_GM_REP_REWARD_RATE_SUCC struct {
	AfRewardRate_Taros        [5]float32
	AfRewardRate_FusionMatter [5]float32
	// SIZE: 40
}

type SP_FE2CL_REP_PC_ITEM_ENCHANT_FAIL struct {
	IErrorCode               int32
	IEnchantItemSlot         int32
	IWeaponMaterialItemSlot  int32
	IDefenceMaterialItemSlot int32
	ICashItemSlot1           int32
	ICashItemSlot2           int32
	// SIZE: 24
}

type SP_LS2CL_REP_LOGIN_SUCC struct {
	ICharCount       int8
	ISlotNum         int8
	IPaymentFlag     int8
	ITempForPacking4 int8
	UiSvrTime        uint64
	SzID             string `size:"33" pad:"2"`
	IOpenBetaFlag    int32
	// SIZE: 84
}

type SP_LS2CL_REP_LOGIN_FAIL struct {
	IErrorCode int32
	SzID       string `size:"33" pad:"2"`
	// SIZE: 72
}

type SP_LS2CL_REP_CHECK_CHAR_NAME_SUCC struct {
	SzFirstName string `size:"9"`
	SzLastName  string `size:"17"`
	// SIZE: 52
}

type SP_LS2CL_REP_CHECK_CHAR_NAME_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_LS2CL_REP_SAVE_CHAR_NAME_SUCC struct {
	IPC_UID     int64
	ISlotNum    int8
	IGender     int8
	SzFirstName string `size:"9"`
	SzLastName  string `size:"17" pad:"2"`
	// SIZE: 64
}

type SP_LS2CL_REP_SAVE_CHAR_NAME_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_LS2CL_REP_CHAR_CREATE_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_LS2CL_REP_CHAR_SELECT_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_LS2CL_REP_CHAR_DELETE_SUCC struct {
	ISlotNum int8
	// SIZE: 1
}

type SP_LS2CL_REP_CHAR_DELETE_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_LS2CL_REP_SHARD_SELECT_SUCC struct {
	G_FE_ServerIP   [16]byte
	G_FE_ServerPort int32
	IEnterSerialKey int64
	// SIZE: 28
}

type SP_LS2CL_REP_SHARD_SELECT_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_LS2CL_REP_VERSION_CHECK_FAIL struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_LS2CL_REP_CHECK_NAME_LIST_SUCC struct {
	IFNCode       int32
	IMNCode       int32
	ILNCode       int32
	ANameCodeFlag [8]int64
	// SIZE: 76
}

type SP_LS2CL_REP_CHECK_NAME_LIST_FAIL struct {
	IFNCode    int32
	IMNCode    int32
	ILNCode    int32
	IErrorCode int32
	// SIZE: 16
}

type SP_LS2CL_REP_PC_EXIT_DUPLICATE struct {
	IErrorCode int32
	// SIZE: 4
}

type SP_LS2CL_REQ_LIVE_CHECK struct {
	ITempValue int32
	// SIZE: 4
}

type SP_LS2CL_REP_CHANGE_CHAR_NAME_SUCC struct {
	IPC_UID     int64
	ISlotNum    int8   `pad:"1"`
	SzFirstName string `size:"9"`
	SzLastName  string `size:"17" pad:"2"`
	// SIZE: 64
}

type SP_LS2CL_REP_CHANGE_CHAR_NAME_FAIL struct {
	IPC_UID    int64
	ISlotNum   int8 `pad:"3"`
	IErrorCode int32
	// SIZE: 16
}

type SP_LS2CL_REP_SHARD_LIST_INFO_SUCC struct {
	AShardConnectFlag [26]byte
	// SIZE: 26
}

type SPCStyle struct {
	IPC_UID     int64
	INameCheck  int8   `pad:"1"`
	SzFirstName string `size:"9"`
	SzLastName  string `size:"17"`
	IGender     int8
	IFaceStyle  int8
	IHairStyle  int8
	IHairColor  int8
	ISkinColor  int8
	IEyeColor   int8
	IHeight     int8
	IBody       int8 `pad:"2"`
	IClass      int32
	// SIZE: 76
}

type SPCStyle2 struct {
	IAppearanceFlag int8
	ITutorialFlag   int8
	IPayzoneFlag    int8
	// SIZE: 3
}

type SRunningQuest struct {
	M_aCurrTaskID      int32
	M_aKillNPCID       [3]int32
	M_aKillNPCCount    [3]int32
	M_aNeededItemID    [3]int32
	M_aNeededItemCount [3]int32
	// SIZE: 52
}

type SOnItem struct {
	IEquipHandID int16
	IEquipUBID   int16
	IEquipLBID   int16
	IEquipFootID int16
	IEquipHeadID int16
	IEquipFaceID int16
	IEquipBackID int16
	// SIZE: 14
}

type SOnItem_Index struct {
	IEquipUBID_index   int16
	IEquipLBID_index   int16
	IEquipFootID_index int16
	IFaceStyle         int16
	IHairStyle         int16
	// SIZE: 10
}

type SItemBase struct {
	IType      int16
	IID        int16
	IOpt       int32
	ITimeLimit int32
	// SIZE: 12
}

type SItemTrade struct {
	IType     int16
	IID       int16
	IOpt      int32
	IInvenNum int32
	ISlotNum  int32
	// SIZE: 16
}

type SItemVendor struct {
	IVendorID int32
	FBuyCost  float32
	Item      SItemBase
	ISortNum  int32
	// SIZE: 24
}

type SItemReward struct {
	SItem    SItemBase
	EIL      int32
	ISlotNum int32
	// SIZE: 20
}

type STimeLimitItemDeleteInfo2CL struct {
	EIL      int32
	ISlotNum int32
	// SIZE: 8
}

type SNanoTuneNeedItemInfo2CL struct {
	ISlotNum int32
	ItemBase SItemBase
	// SIZE: 16
}

type SEmailItemInfoFromCL struct {
	ISlotNum  int32
	ItemInven SItemBase
	// SIZE: 16
}

type SEPRecord struct {
	UiScore int16
	UiRank  int8
	UiRing  int8
	UiTime  int16
	// SIZE: 6
}

type SBuddyBaseInfo struct {
	IID            int32
	IPCUID         int64
	BBlocked       int8
	BFreeChat      int8
	IPCState       int8   `pad:"1"`
	SzFirstName    string `size:"9"`
	SzLastName     string `size:"17"`
	IGender        int8
	INameCheckFlag int8 `pad:"2"`
	// SIZE: 72
}

type SBuddyStyleInfo struct {
	SBuddyStyle SPCStyle
	AEquip      [9]SItemBase
	// SIZE: 184
}

type SSYSTEMTIME struct {
	WYear         int32
	WMonth        int32
	WDayOfWeek    int32
	WDay          int32
	WHour         int32
	WMinute       int32
	WSecond       int32
	WMilliseconds int32
	// SIZE: 32
}

type SEmailInfo struct {
	IEmailIndex    int64
	IFromPCUID     int64
	SzFirstName    string `size:"9"`
	SzLastName     string `size:"17"`
	SzSubject      string `size:"32"`
	IReadFlag      int32
	SendTime       SSYSTEMTIME
	DeleteTime     SSYSTEMTIME
	IItemCandyFlag int32
	// SIZE: 204
}

type SNano struct {
	IID      int16
	ISkillID int16
	IStamina int16
	// SIZE: 6
}

type SNanoBank struct {
	ISkillID int16
	IStamina int16
	// SIZE: 4
}

type STimeBuff struct {
	ITimeLimit    uint64
	ITimeDuration uint64
	ITimeRepeat   int32
	IValue        int32
	IConfirmNum   int32
	// SIZE: 28
}

type STimeBuff_Svr struct {
	ITimeLimit    uint64
	ITimeDuration uint64
	ITimeRepeat   int32
	IValue        int32
	IConfirmNum   int32
	ITimeFlow     int16 `pad:"2"`
	// SIZE: 32
}

type SPCLoadData2CL struct {
	IUserLevel          int16 `pad:"2"`
	PCStyle             SPCStyle
	PCStyle2            SPCStyle2 `pad:"1"`
	ILevel              int16
	IMentor             int16
	IMentorCount        int16 `pad:"2"`
	IHP                 int32
	IBatteryW           int32
	IBatteryN           int32
	ICandy              int32
	IFusionMatter       int32
	ISpecialState       int8 `pad:"3"`
	IMapNum             int32
	IX                  int32
	IY                  int32
	IZ                  int32
	IAngle              int32
	AEquip              [9]SItemBase
	AInven              [50]SItemBase
	AQInven             [50]SItemBase
	ANanoBank           [37]SNano
	ANanoSlots          [3]int16
	IActiveNanoSlotNum  int16 `pad:"2"`
	IConditionBitFlag   int32
	ECSTB___Add         int32
	TimeBuff            STimeBuff
	AQuestFlag          [32]int64
	ARepeatQuestFlag    [8]int64
	ARunningQuest       [9]SRunningQuest
	ICurrentMissionID   int32
	IWarpLocationFlag   int32
	AWyvernLocationFlag [2]int64
	IBuddyWarpTime      int32
	IFatigue            int32
	IFatigue_Level      int32
	IFatigueRate        int32
	IFirstUseFlag1      int64
	IFirstUseFlag2      int64
	AiPCSkill           [33]int32
	// SIZE: 2688
}

type SPCAppearanceData struct {
	IID               int32
	PCStyle           SPCStyle
	IConditionBitFlag int32
	IPCState          int8
	ISpecialState     int8
	ILv               int16
	IHP               int32
	IMapNum           int32
	IX                int32
	IY                int32
	IZ                int32
	IAngle            int32
	ItemEquip         [9]SItemBase
	Nano              SNano `pad:"2"`
	ERT               int32
	// SIZE: 232
}

type SNPCAppearanceData struct {
	INPC_ID           int32
	INPCType          int32
	IHP               int32
	IConditionBitFlag int32
	IX                int32
	IY                int32
	IZ                int32
	IAngle            int32
	IBarkerType       int32
	// SIZE: 36
}

type SBulletAppearanceData struct {
	IBullet_ID int32
	IX         int32
	IY         int32
	IZ         int32
	IAngle     int32
	// SIZE: 20
}

type STransportationLoadData struct {
	IAISvrID int32
	ETT      int32
	IT_Type  int32
	IMapType int32
	IMapNum  int32
	IX       int32
	IY       int32
	IZ       int32
	// SIZE: 32
}

type STransportationAppearanceData struct {
	ETT     int32
	IT_ID   int32
	IT_Type int32
	IX      int32
	IY      int32
	IZ      int32
	// SIZE: 24
}

type SShinyAppearanceData struct {
	IShiny_ID  int32
	IShinyType int32
	IMapNum    int32
	IX         int32
	IY         int32
	IZ         int32
	// SIZE: 24
}

type SAttackResult struct {
	ECT        int32
	IID        int32
	BProtected int32
	IDamage    int32
	IHP        int32
	IHitFlag   int8 `pad:"3"`
	// SIZE: 24
}

type SCAttackResult struct {
	ECT                int32
	IID                int32
	BProtected         int32
	IDamage            int32
	IHP                int32
	IHitFlag           int8 `pad:"1"`
	IActiveNanoSlotNum int16
	BNanoDeactive      int32
	INanoID            int16
	INanoStamina       int16
	IConditionBitFlag  int32
	ECSTB___Del        int32
	// SIZE: 40
}

type SSkillResult_Damage struct {
	ECT        int32
	IID        int32
	BProtected int32
	IDamage    int32
	IHP        int32
	// SIZE: 20
}

type SSkillResult_DotDamage struct {
	ECT               int32
	IID               int32
	BProtected        int32
	IDamage           int32
	IHP               int32
	IStamina          int16 `pad:"2"`
	BNanoDeactive     int32
	IConditionBitFlag int32
	// SIZE: 32
}

type SSkillResult_Heal_HP struct {
	ECT     int32
	IID     int32
	IHealHP int32
	IHP     int32
	// SIZE: 16
}

type SSkillResult_Heal_Stamina struct {
	ECT              int32
	IID              int32
	IHealNanoStamina int16
	Nano             SNano
	// SIZE: 16
}

type SSkillResult_Stamina_Self struct {
	ECT              int32
	IID              int32
	IReduceHP        int32
	IHP              int32
	IHealNanoStamina int16
	Nano             SNano
	// SIZE: 24
}

type SSkillResult_Damage_N_Debuff struct {
	ECT               int32
	IID               int32
	BProtected        int32
	IDamage           int32
	IHP               int32
	IStamina          int16 `pad:"2"`
	BNanoDeactive     int32
	IConditionBitFlag int32
	// SIZE: 32
}

type SSkillResult_Buff struct {
	ECT               int32
	IID               int32
	BProtected        int32
	IConditionBitFlag int32
	// SIZE: 16
}

type SSkillResult_BatteryDrain struct {
	ECT               int32
	IID               int32
	BProtected        int32
	IDrainW           int32
	IBatteryW         int32
	IDrainN           int32
	IBatteryN         int32
	IStamina          int16 `pad:"2"`
	BNanoDeactive     int32
	IConditionBitFlag int32
	// SIZE: 40
}

type SSkillResult_Damage_N_Move struct {
	ECT        int32
	IID        int32
	BProtected int32
	IDamage    int32
	IHP        int32
	IMoveX     int32
	IMoveY     int32
	IMoveZ     int32
	IBlockMove int32
	// SIZE: 36
}

type SSkillResult_Move struct {
	ECT     int32
	IID     int32
	IMapNum int32
	IMoveX  int32
	IMoveY  int32
	IMoveZ  int32
	// SIZE: 24
}

type SSkillResult_Resurrect struct {
	ECT      int32
	IID      int32
	IRegenHP int32
	// SIZE: 12
}

type SPC_HP struct {
	IPC_ID int32
	IHP    int32
	// SIZE: 8
}

type SPC_BATTERYs struct {
	IPC_ID    int32
	IBatteryW int32
	IBatteryN int32
	// SIZE: 12
}

type SPC_NanoSlots struct {
	ANanoSlots         [3]int32
	IActiveNanoSlotNum int16 `pad:"2"`
	// SIZE: 16
}

type SPC_Nano struct {
	IPC_ID             int32
	Nano               SNano
	IActiveNanoSlotNum int16
	// SIZE: 12
}

type SPCRegenData struct {
	IHP                int32
	IMapNum            int32
	IX                 int32
	IY                 int32
	IZ                 int32
	IActiveNanoSlotNum int16
	Nanos              [3]SNano
	// SIZE: 40
}

type SPCRegenDataForOtherPC struct {
	IPC_ID            int32
	IHP               int32
	IX                int32
	IY                int32
	IZ                int32
	IAngle            int32
	IConditionBitFlag int32
	IPCState          int8
	ISpecialState     int8
	Nano              SNano
	// SIZE: 36
}

type SPCBullet struct {
	EAT      int32
	IID      int32
	BCharged int32
	// SIZE: 12
}

type SNPCBullet struct {
	EAT      int32
	IID      int32
	BCharged int32
	EST      int32
	// SIZE: 16
}

type SNPCLocationData struct {
	INPC_Type int32
	IX        int32
	IY        int32
	IZ        int32
	IAngle    int32
	IRoute    int32
	// SIZE: 24
}

type SGroupNPCLocationData struct {
	IGroupType   int32
	IX           int32
	IY           int32
	IZ           int32
	IAngle       int32
	IRoute       int32
	AGroupNPCIDs [5]int32
	// SIZE: 44
}

type SPCGroupMemberInfo struct {
	IPC_ID        int32
	IPCUID        uint64
	INameCheck    int8   `pad:"1"`
	SzFirstName   string `size:"9"`
	SzLastName    string `size:"17"`
	ISpecialState int8   `pad:"1"`
	ILv           int16  `pad:"2"`
	IHP           int32
	IMaxHP        int32
	IMapType      int32
	IMapNum       int32
	IX            int32
	IY            int32
	IZ            int32
	BNano         int32
	Nano          SNano `pad:"2"`
	// SIZE: 112
}

type SNPCGroupMemberInfo struct {
	INPC_ID   int32
	INPC_Type int32
	IHP       int32
	IMapType  int32
	IMapNum   int32
	IX        int32
	IY        int32
	IZ        int32
	// SIZE: 32
}

type SEPElement struct {
	ILID       int32
	IGID       int32
	IType      int32
	ITargetGID int32
	IX         int32
	IY         int32
	IZ         int32
	IEnable    int32
	IONOFF     int32
	// SIZE: 36
}

type SCNStreetStall_ItemInfo_for_Client struct {
	IListNum int32
	Item     SItemBase
	IPrice   int32
	// SIZE: 20
}

type SQuickSlot struct {
	IType int16
	IID   int16
	// SIZE: 4
}

type SP_CL2FE_REQ_PC_VENDOR_ITEM_BUY struct {
	INPC_ID       int32
	IVendorID     int32
	IListID       int8 `pad:"3"`
	Item          SItemBase
	IInvenSlotNum int32
	// SIZE: 28
}

type SP_CL2FE_REQ_PC_GIVE_ITEM struct {
	EIL       int32
	ISlotNum  int32
	Item      SItemBase
	ITimeLeft int32
	// SIZE: 24
}

type SP_CL2FE_REQ_PC_TRADE_ITEM_REGISTER struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	Item        SItemTrade
	// SIZE: 28
}

type SP_CL2FE_REQ_PC_TRADE_ITEM_UNREGISTER struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	Item        SItemTrade
	// SIZE: 28
}

type SP_CL2FE_REQ_PC_VENDOR_ITEM_RESTORE_BUY struct {
	INPC_ID       int32
	IVendorID     int32
	IListID       int8 `pad:"3"`
	Item          SItemBase
	IInvenSlotNum int32
	// SIZE: 28
}

type SP_CL2FE_REQ_ITEM_CHEST_OPEN struct {
	EIL       int32
	ISlotNum  int32
	ChestItem SItemBase
	// SIZE: 20
}

type SP_CL2FE_REQ_PC_VENDOR_BATTERY_BUY struct {
	INPC_ID   int32
	IVendorID int32
	IListID   int8 `pad:"3"`
	Item      SItemBase
	// SIZE: 24
}

type SP_CL2FE_REQ_PC_SEND_EMAIL struct {
	ITo_PCUID int64
	SzSubject string `size:"32"`
	SzContent string `size:"512"`
	AItem     [4]SEmailItemInfoFromCL
	ICash     int32
	// SIZE: 1164
}

type SP_CL2FE_PC_STREETSTALL_REQ_REGIST_ITEM struct {
	IItemListNum      int32
	IItemInvenSlotNum int32
	Item              SItemBase
	IPrice            int32
	// SIZE: 24
}

type SP_CL2LS_REQ_CHAR_CREATE struct {
	PCStyle        SPCStyle
	SOn_Item       SOnItem
	SOn_Item_Index SOnItem_Index
	// SIZE: 100
}

type SP_FE2CL_REP_PC_ENTER_SUCC struct {
	IID           int32
	PCLoadData2CL SPCLoadData2CL
	UiSvrTime     uint64
	// SIZE: 2700
}

type SP_FE2CL_PC_NEW struct {
	PCAppearanceData SPCAppearanceData
	// SIZE: 232
}

type SP_FE2CL_NPC_ENTER struct {
	NPCAppearanceData SNPCAppearanceData
	// SIZE: 36
}

type SP_FE2CL_NPC_NEW struct {
	NPCAppearanceData SNPCAppearanceData
	// SIZE: 36
}

type SP_FE2CL_REP_PC_REGEN_SUCC struct {
	PCRegenData   SPCRegenData
	BMoveLocation int32
	IFusionMatter int32
	// SIZE: 48
}

type SP_FE2CL_PC_ITEM_MOVE_SUCC struct {
	EFrom        int32
	IFromSlotNum int32
	FromSlotItem SItemBase
	ETo          int32
	IToSlotNum   int32
	ToSlotItem   SItemBase
	// SIZE: 40
}

type SP_FE2CL_PC_EQUIP_CHANGE struct {
	IPC_ID        int32
	IEquipSlotNum int32
	EquipSlotItem SItemBase
	// SIZE: 20
}

type SP_FE2CL_REP_NANO_TUNE_SUCC struct {
	INanoID          int16
	ISkillID         int16
	IPC_FusionMatter int32
	AiItemSlotNum    [10]int32
	AItem            [10]SItemBase
	// SIZE: 168
}

type SP_FE2CL_NANO_ACTIVE struct {
	IPC_ID            int32
	Nano              SNano `pad:"2"`
	IConditionBitFlag int32
	ECSTB___Add       int32
	// SIZE: 20
}

type SP_FE2CL_REP_PC_TICK struct {
	IHP               int32
	ANano             [3]SNano `pad:"2"`
	IBatteryN         int32
	BResetMissionFlag int32
	// SIZE: 32
}

type SP_FE2CL_REP_PC_VENDOR_ITEM_BUY_SUCC struct {
	ICandy        int32
	IInvenSlotNum int32
	Item          SItemBase
	// SIZE: 20
}

type SP_FE2CL_REP_PC_VENDOR_ITEM_SELL_SUCC struct {
	ICandy        int32
	IInvenSlotNum int32
	Item          SItemBase
	ItemStay      SItemBase
	// SIZE: 32
}

type SP_FE2CL_REP_PC_ROCKET_STYLE_FIRE_SUCC struct {
	ISkillID      int32
	IX            int32
	IY            int32
	IZ            int32
	IToX          int32
	IToY          int32
	IToZ          int32
	IBulletID     int8 `pad:"3"`
	Bullet        SPCBullet
	IBatteryW     int32
	BNanoDeactive int32
	INanoID       int16
	INanoStamina  int16
	// SIZE: 56
}

type SP_FE2CL_PC_ROCKET_STYLE_FIRE struct {
	IPC_ID        int32
	IX            int32
	IY            int32
	IZ            int32
	IToX          int32
	IToY          int32
	IToZ          int32
	IBulletID     int8 `pad:"3"`
	Bullet        SPCBullet
	BNanoDeactive int32
	// SIZE: 48
}

type SP_FE2CL_PC_ROCKET_STYLE_HIT struct {
	IPC_ID     int32
	IBulletID  int8 `pad:"3"`
	Bullet     SPCBullet
	ITargetCnt int32
	// SIZE: 24
}

type SP_FE2CL_REP_PC_GRENADE_STYLE_FIRE_SUCC struct {
	ISkillID      int32
	IToX          int32
	IToY          int32
	IToZ          int32
	IBulletID     int8 `pad:"3"`
	Bullet        SPCBullet
	IBatteryW     int32
	BNanoDeactive int32
	INanoID       int16
	INanoStamina  int16
	// SIZE: 44
}

type SP_FE2CL_PC_GRENADE_STYLE_FIRE struct {
	IPC_ID        int32
	IToX          int32
	IToY          int32
	IToZ          int32
	IBulletID     int8 `pad:"3"`
	Bullet        SPCBullet
	BNanoDeactive int32
	// SIZE: 36
}

type SP_FE2CL_PC_GRENADE_STYLE_HIT struct {
	IPC_ID     int32
	IBulletID  int8 `pad:"3"`
	Bullet     SPCBullet
	ITargetCnt int32
	// SIZE: 24
}

type SP_FE2CL_REP_PC_TRADE_CONFIRM_SUCC struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	Item        [12]SItemTrade
	ICandy      int32
	ItemStay    [12]SItemTrade
	// SIZE: 400
}

type SP_FE2CL_REP_PC_TRADE_ITEM_REGISTER_SUCC struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	TradeItem   SItemTrade
	InvenItem   SItemTrade
	// SIZE: 44
}

type SP_FE2CL_REP_PC_TRADE_ITEM_UNREGISTER_SUCC struct {
	IID_Request int32
	IID_From    int32
	IID_To      int32
	TradeItem   SItemTrade
	InvenItem   SItemTrade
	// SIZE: 44
}

type SP_FE2CL_REP_PC_NANO_CREATE_SUCC struct {
	IPC_FusionMatter  int32
	IQuestItemSlotNum int32
	QuestItem         SItemBase
	Nano              SNano
	IPC_Level         int16
	// SIZE: 28
}

type SP_FE2CL_REP_PC_BANK_OPEN_SUCC struct {
	ABank      [119]SItemBase
	IExtraBank int32
	// SIZE: 1432
}

type SP_FE2CL_REP_PC_VENDOR_TABLE_UPDATE_SUCC struct {
	Item [20]SItemVendor
	// SIZE: 480
}

type SP_FE2CL_REP_PC_VENDOR_ITEM_RESTORE_BUY_SUCC struct {
	ICandy        int32
	IInvenSlotNum int32
	Item          SItemBase
	// SIZE: 20
}

type SP_FE2CL_REP_PC_GIVE_ITEM_SUCC struct {
	EIL      int32
	ISlotNum int32
	Item     SItemBase
	// SIZE: 20
}

type SP_FE2CL_REP_ACCEPT_MAKE_BUDDY_SUCC struct {
	IBuddySlot int8 `pad:"3"`
	BuddyInfo  SBuddyBaseInfo
	// SIZE: 76
}

type SP_FE2CL_REP_GET_BUDDY_STYLE_SUCC struct {
	IBuddyPCUID int64
	IBuddySlot  int8 `pad:"3"`
	SBuddyStyle SBuddyStyleInfo
	// SIZE: 196
}

type SP_FE2CL_NPC_ROCKET_STYLE_FIRE struct {
	INPC_ID   int32
	IX        int32
	IY        int32
	IZ        int32
	IToX      int32
	IToY      int32
	IToZ      int32
	IBulletID int8 `pad:"3"`
	Bullet    SNPCBullet
	// SIZE: 48
}

type SP_FE2CL_NPC_GRENADE_STYLE_FIRE struct {
	INPC_ID   int32
	IToX      int32
	IToY      int32
	IToZ      int32
	IBulletID int8 `pad:"3"`
	Bullet    SNPCBullet
	// SIZE: 36
}

type SP_FE2CL_NPC_BULLET_STYLE_HIT struct {
	INPC_ID    int32
	IBulletID  int8 `pad:"3"`
	Bullet     SNPCBullet
	ITargetCnt int32
	// SIZE: 28
}

type SP_FE2CL_REP_PC_WARP_USE_NPC_SUCC struct {
	IX           int32
	IY           int32
	IZ           int32
	EIL          int32
	IItemSlotNum int32
	Item         SItemBase
	ICandy       int32
	// SIZE: 36
}

type SP_FE2CL_REP_GET_MEMBER_STYLE_SUCC struct {
	IMemberID      int32
	IMemberUID     int64
	BuddyStyleInfo SBuddyStyleInfo
	// SIZE: 196
}

type SP_FE2CL_PC_REGEN struct {
	PCRegenDataForOtherPC SPCRegenDataForOtherPC
	// SIZE: 36
}

type SP_FE2CL_TRANSPORTATION_ENTER struct {
	AppearanceData STransportationAppearanceData
	// SIZE: 24
}

type SP_FE2CL_TRANSPORTATION_NEW struct {
	AppearanceData STransportationAppearanceData
	// SIZE: 24
}

type SP_FE2CL_REP_EP_RACE_END_SUCC struct {
	IEPRaceMode     int32
	IEPRaceTime     int32
	IEPRingCnt      int32
	IEPScore        int32
	IEPRank         int32
	IEPRewardFM     int32
	IEPTopScore     int32
	IEPTopRank      int32
	IEPTopTime      int32
	IEPTopRingCount int32
	IFusionMatter   int32
	RewardItem      SItemReward
	IFatigue        int32
	IFatigue_Level  int32
	// SIZE: 72
}

type SP_FE2CL_SHINY_ENTER struct {
	ShinyAppearanceData SShinyAppearanceData
	// SIZE: 24
}

type SP_FE2CL_SHINY_NEW struct {
	ShinyAppearanceData SShinyAppearanceData
	// SIZE: 24
}

type SP_FE2CL_REP_PC_ITEM_USE_SUCC struct {
	IPC_ID     int32
	EIL        int32
	ISlotNum   int32
	RemainItem SItemBase
	ISkillID   int16 `pad:"2"`
	EST        int32
	ITargetCnt int32
	// SIZE: 36
}

type SP_FE2CL_PC_BUFF_UPDATE struct {
	ECSTB             int32
	ETBU              int32
	ETBT              int32
	TimeBuff          STimeBuff
	IConditionBitFlag int32
	// SIZE: 44
}

type SP_FE2CL_REP_PC_READ_EMAIL_SUCC struct {
	IEmailIndex int64
	SzContent   string `size:"512"`
	AItem       [4]SItemBase
	ICash       int32
	// SIZE: 1084
}

type SP_FE2CL_REP_PC_RECV_EMAIL_PAGE_LIST_SUCC struct {
	IPageNum   int8 `pad:"3"`
	AEmailInfo [5]SEmailInfo
	// SIZE: 1024
}

type SP_FE2CL_REP_PC_SEND_EMAIL_SUCC struct {
	ITo_PCUID int64
	ICandy    int32
	AItem     [4]SEmailItemInfoFromCL
	// SIZE: 76
}

type SP_FE2CL_REP_PC_TRANSPORT_WARP_SUCC struct {
	TransportationAppearanceData STransportationAppearanceData
	ILcX                         int32
	ILcY                         int32
	ILcZ                         int32
	// SIZE: 36
}

type SP_FE2CL_PC_STREETSTALL_REP_REGIST_ITEM_SUCC struct {
	IItemListNum      int32
	IItemInvenSlotNum int32
	Item              SItemBase
	IPrice            int32
	// SIZE: 24
}

type SP_FE2CL_PC_STREETSTALL_REP_SALE_START_SUCC struct {
	IStreetStallItemInvenSlotNum int32
	OpenItem                     SItemBase
	EPCCharState                 int32
	// SIZE: 20
}

type SP_FE2CL_PC_STREETSTALL_REP_ITEM_BUY_SUCC_BUYER struct {
	IStreetStallPC_ID    int32
	IPC_Candy            int32
	IPC_ItemInvenSlotNum int32
	PC_Item              SItemBase
	IItemListNum         int32
	// SIZE: 28
}

type SP_FE2CL_PC_STREETSTALL_REP_ITEM_BUY_SUCC_SELLER struct {
	IBuyerPC_ID                     int32
	IStreetStallPC_Candy            int32
	IStreetStallPC_ItemInvenSlotNum int32
	StreetStallPC_Item              SItemBase
	IItemListNum                    int32
	// SIZE: 28
}

type SP_FE2CL_REP_PC_ITEM_COMBINATION_SUCC struct {
	INewItemSlot   int32
	SNewItem       SItemBase
	IStatItemSlot  int32
	ICashItemSlot1 int32
	ICashItemSlot2 int32
	ICandy         int32
	ISuccessFlag   int32
	// SIZE: 36
}

type SP_FE2CL_PC_CASH_BUFF_UPDATE struct {
	ECSTB             int32
	ETBU              int32
	TimeBuff          STimeBuff
	IConditionBitFlag int32
	// SIZE: 40
}

type SP_FE2CL_REP_PC_SKILL_ADD_SUCC struct {
	ISkillSlotNum          int32
	ISkillID               int32
	ISkillItemInvenSlotNum int32
	SkillItem              SItemBase
	// SIZE: 24
}

type SP_FE2CL_PC_QUICK_SLOT_INFO struct {
	AQuickSlot [8]SQuickSlot
	// SIZE: 32
}

type SP_FE2CL_REP_PC_DISASSEMBLE_ITEM_SUCC struct {
	INewItemSlot int32
	SNewItem     SItemBase
	// SIZE: 16
}

type SP_FE2CL_REP_PC_ITEM_ENCHANT_SUCC struct {
	IEnchantItemSlot         int32
	SEnchantItem             SItemBase
	IWeaponMaterialItemSlot  int32
	SWeaponMaterialItem      SItemBase
	IDefenceMaterialItemSlot int32
	SDefenceMaterialItem     SItemBase
	ICashItemSlot1           int32
	ICashItemSlot2           int32
	ICandy                   int32
	ISuccessFlag             int32
	// SIZE: 64
}

type SP_LS2CL_REP_CHAR_INFO struct {
	ISlot      int8 `pad:"1"`
	ILevel     int16
	SPC_Style  SPCStyle
	SPC_Style2 SPCStyle2 `pad:"1"`
	IX         int32
	IY         int32
	IZ         int32
	AEquip     [9]SItemBase
	// SIZE: 204
}

type SP_LS2CL_REP_CHAR_CREATE_SUCC struct {
	ILevel     int16 `pad:"2"`
	SPC_Style  SPCStyle
	SPC_Style2 SPCStyle2 `pad:"1"`
	SOn_Item   SOnItem   `pad:"2"`
	// SIZE: 100
}
