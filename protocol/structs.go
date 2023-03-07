// generated via genstructs.py - All structure padding and member alignment verified
package protocol

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
