package room

import ()

type Room struct {
	roomId  int       //4位数字的房间号
	players []*Player //当前房间内所有玩家

	setupKey string //当前房间配置
}

//---------- 新建房间 ----------

// `n`
func NewRoom(creator *Player) int {
	//新建房间对象, 并设置creator为当前`上帝`moderator
	//在全局map中, 保存下<roomId -> *Room>
	//房间新建后 会等待玩家们陆续进入房间
	//当所有想参与的玩家都进入后, 由`上帝`回复1开始选择配置
}

//---------- `玩家`进入房间 ----------

// `????`
func (r *Room) Enter(p *Player) {
	//玩家进入房间:
	// players[] 中追加 p对象, 向p发送确认响应
	// 同时向moderator发送下当前房间状态信息
}

//---------- `上帝`锁定玩家人数 ----------

// `r`  {可以多次发送ready}
func (r *Room) Ready(p *Player) {
	//根据当前玩家的数量, 向上帝发送配置选项
	// {只有p==moderator}才能执行该操作
}

//---------- `上帝`选择配置 ----------

// `c???`  {配置采用全局编号}?
func (r *Room) Config(p *Player, cfg string) {
	//解析cfg字符串, 完成房间角色配置{随机算法,但避免同一身份连续出现}, 并启动游戏
	// {只有p==moderator}才能执行该操作
}

//
//---------- 分配角色 ----------
//---------- 开始游戏 ----------

// 第一版 先做成智能首杀(回复1满意,回复0不满意  超过半数不满意则重新选择)  女巫,预言家还是用人工操作   1~使用  0~使用
// 为了减少干扰, 平民也需要操作下   回复自己的数字 #??

//---------- 狼人行动 ----------  ?? 怎么判定 狼人活动结束??
//---------- 女巫行动 ----------
//---------- 预言家行动 ----------

// -- 怎么能让每个人都操作手机, 减少场外干扰 ??
//

//---------- 查看第一晚结果(已完成警长竞选) ----------
// `v`
func (r *Room) ViewFirstNightResult(p *Player) {
	//查看第一晚结果, 必须是竞选警长完成后
	// {只有p==moderator}才能执行该操作
	// 群发
}

//---------- `上帝`转让给下家 ----------

// `t??`
func (r *Room) Transfer(p *Player, playerId int) {
	//将此房间的`上帝`moderator 转让给 #playerId 号玩家
	// {只有p==moderator}才能执行该操作
}

//---------- `玩家`退出房间 ----------

// `q`
func (r *Room) Quit(p *Player) {
	//结束房间, 将结果发送给各个玩家
	// {只有p==moderator}才能执行该操作
}