package dingTalk

import (
	"io"
	"net/http"
)

// Response Http response from oss
type Response struct {
	StatusCode int
	Headers    http.Header
	Body       io.ReadCloser
}

// 审批实例表单
type FormValues struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	ExtValue string `json:"ext_value"`
}

// access_token response
type AccessTokenRsp struct {
	ErrCode     int    `json:"errcode"`
	AccessToken string `json:"access_token"`
	ErrMsg      string `json:"errmsg"`
	ExpiresIn   int    `json:"expires_in"`
}

// create bpms_instance_task
type BpmsInstanceTaskRsp struct {
	ErrCode           int    `json:"errcode"`
	ProcessInstanceId string `json:"process_instance_id"`
	RequestId         string `json:"request_id"`
}

// get bpms_instance_task
type BpmsInstanceTaskDetailRsp struct {
	ErrCode         int                 `json:"errcode"`
	ProcessInstance ProcessInstanceItem `json:"process_instance"`
}

type ProcessInstanceItem struct {
	Title               string                 `json:"title"`              // 审批实例标题
	CreateTime          string                 `json:"create_time"`        // 开始时间
	FinishTime          string                 `json:"finish_time"`        // 结束时间
	OriginatorUserId    string                 `json:"originator_userid"`  // 发起人
	OriginatorDeptId    string                 `json:"originator_dept_id"` // 发起部门
	Status              string                 `json:"status"`
	CcUserIds           []string               `json:"cc_userids"`
	FormComponentValues []FormValues           `json:"form_component_values"`
	Result              string                 `json:"result"`            // 审批结果
	BusinessId          string                 `json:"business_id"`       // 审批实例业务编号
	OperationRecords    []OperationRecordsItem `json:"operation_records"` // 操作记录列表
	Tasks               []TaskItem             `json:"tasks"`
	OriginatorDeptName  string                 `json:"originator_dept_name"`
	// 审批实例业务动作，MODIFY表示该审批实例是基于原来的实例修改而来，REVOKE表示该审批实例对原来的实例进行撤销，NONE表示正常发起
	BizAction string `json:"biz_action"`
	// 审批附属实例列表，当已经通过的审批实例被修改或撤销，会生成一个新的实例，作为原有审批实例的附属。
	// 如果想知道当前已经通过的审批实例的状态，可以依次遍历它的附属列表，查询里面每个实例的biz_action
	AttachedProcessInstanceIds []string `json:"attached_process_instance_ids"`
}

// 操作记录
type OperationRecordsItem struct {
	Userid string `json:"userid"` // 操作人
	Date   string `json:"date"`   // 操作时间
	// 操作分类 EXECUTE_TASK_NORMAL（正常执行任务），EXECUTE_TASK_AGENT（代理人执行任务）
	// APPEND_TASK_BEFORE（前加签任务），APPEND_TASK_AFTER（后加签任务），REDIRECT_TASK（转交任务）
	// START_PROCESS_INSTANCE（发起流程实例），TERMINATE_PROCESS_INSTANCE（终止(撤销)流程实例）
	// FINISH_PROCESS_INSTANCE（结束流程实例），ADD_REMARK（添加评论）
	OperationType   string `json:"operation_type"`
	OperationResult string `json:"operation_result"` // 操作结果 AGREE（同意），REFUSE（拒绝）
	Remark          string `json:"remark"`           // 评论内容。审批操作附带评论时才返回该字段。
}

// 已审批任务列表
type TaskItem struct {
	UserId string `json:"userid"` // 任务处理人
	// 任务状态，分为 NEW（未启动），RUNNING（处理中），PAUSED（暂停），CANCELED（取消），COMPLETED（完成），TERMINATED（终止）
	TaskStatus string `json:"task_status"`
	// 结果，分为NONE（无），AGREE（同意），REFUSE（拒绝），REDIRECTED（转交）
	TaskResult string `json:"task_result"`
	CreateTime string `json:"create_time"` // 开始时间。yyyy-MM-dd HH:mm:ss格式
	FinishTime string `json:"finish_time"` // 结束时间。yyyy-MM-dd HH:mm:ss格式。当前任务结束时才会有这个字段返回。
	TaskId     string `json:"taskid"`      // 任务节点id
}

// 钉钉部门
// start ----------------------------------------------------
type QueryDept struct {
	Id         string `json:"id"`          // 部门id
	FetchChild bool   `json:"fetch_child"` // 是否递归部门的全部子部门，ISV微应用固定传递false
	Lang       string `json:"lang"`        // 通讯录语言（默认zh_CN，未来会支持en_US）
}

type DepartmentList struct {
	ErrCode        int          `json:"errcode"`
	ErrMsg         string       `json:"errmsg"`
	DepartmentItem []Department `json:"department"`
}

type Department struct {
	ErrCode               int    `json:"errcode"`
	ErrMsg                string `json:"errmsg"`
	Id                    int    `json:"id"`              // 部门id
	Name                  string `json:"name"`            // 部门名称
	ParentId              int    `json:"parentid"`        // 父部门id，根部门为1
	Order                 int    `json:"order"`           // 当前部门在父部门下的所有子部门中的排序值
	CreateDeptGroup       bool   `json:"createDeptGroup"` // 是否同步创建一个关联此部门的企业群，true表示是，false表示不是
	AutoAddUser           bool   `json:"autoAddUser"`     // 当部门群已经创建后，是否有新人加入部门会自动加入该群，true表示是，false表示不是
	DeptHiding            bool   `json:"deptHiding"`      // 是否隐藏部门，true表示隐藏，false表示显示
	DeptPermits           string `json:"deptPermits"`     // 可以查看指定隐藏部门的其他部门列表，如果部门隐藏，则此值生效，取值为其他的部门id组成的的字符串，使用“|”符号进行分割
	UserPermits           string `json:"userPermits"`     // 可以查看指定隐藏部门的其他人员列表，如果部门隐藏，则此值生效，取值为其他的人员userid组成的的字符串，使用“|”符号进行分割
	OuterDept             bool   `json:"outerDept"`       // 是否本部门的员工仅可见员工自己，为true时，本部门员工默认只能看到员工自己
	OuterPermitDepts      string `json:"outerPermitDepts"`
	OuterPermitUsers      string `json:"outerPermitUsers"`
	OrgDeptOwner          string `json:"orgDeptOwner"`          // 企业群群主
	DeptManagerUserIdList string `json:"deptManagerUseridList"` // 部门的主管列表，取值为由主管的userid组成的字符串，不同的userid使用“|”符号进行分割
	SourceIdentifier      string `json:"sourceIdentifier"`      // 部门标识字段，开发者可用该字段来唯一标识一个部门，并与钉钉外部通讯录里的部门做映射
}

// ending ---------------------------------------------------

// 钉钉用户
// start ----------------------------------------------------
type QueryUser struct {
	UserId string `json:"userid"` // 员工id
	Lang   string `json:"lang"`   // 通讯录语言（默认zh_CN，未来会支持en_US）
}

type QueryUserList struct {
	DepartmentId uint64 `json:"department_id"` // 获取的部门id，1表示根部门
	Lang         string `json:"lang"`          // 通讯录语言（默认zh_CN，未来会支持en_US）
	OffSet       uint64 `json:"offset"`        // 支持分页查询，与size参数同时设置时才生效，此参数代表偏移量,偏移量从0开始
	Size         uint64 `json:"size"`          // 支持分页查询，与offset参数同时设置时才生效，此参数代表分页大小，最大100
	// 支持分页查询，部门成员的排序规则，默认 是按自定义排序；
	// entry_asc：代表按照进入部门的时间升序，
	// entry_desc：代表按照进入部门的时间降序，
	// modify_asc：代表按照部门信息修改时间升序，
	// modify_desc：代表按照部门信息修改时间降序，
	// custom：代表用户定义(未定义时按照拼音)排序
	Order string `json:"order"`
}

type UserList struct {
	ErrCode  int    `json:"errcode"`
	ErrMsg   string `json:"errmsg"`
	UserItem []User `json:"userlist"`
}

type User struct {
	ErrCode   int    `json:"errcode"`
	ErrMsg    string `json:"errmsg"`
	UserId    string `json:"userid"`    // 员工在当前企业内的唯一标识，也称staffId
	UnionId   string `json:"unionid"`   // 员工在当前开发者企业账号范围内的唯一标识，系统生成，固定值，不会改变
	Email     string `json:"email"`     // 邮箱。长度为0~64个字符。企业内必须唯一，不可重复
	Mobile    string `json:"mobile"`    // 手机号码，企业内必须唯一，不可重复。如果是国际号码，请使用+xx-xxxxxx的格式
	OrgEmail  string `json:"orgEmail"`  // 员工的企业邮箱，员工的企业邮箱已开通，才能增加此字段， 否则会报错
	Tel       string `json:"tel"`       // 分机号，长度为0~50个字符，企业内必须唯一，不可重复
	WorkPlace string `json:"workPlace"` // 办公地点，长度为0~50个字符
	Remark    string `json:"remark"`    // 备注，长度为0~1000个字符
	// 表示人员在此部门中的排序，列表是按order的倒序排列输出的，即从大到小排列输出的
	// （钉钉管理后台里面调整了顺序的话order才有值）
	Order           int    `json:"order"`
	Name            string `json:"name"`            // 员工名字
	Active          bool   `json:"active"`          // 是否已经激活，true表示已激活，false表示未激活
	OrderInDepts    string `json:"orderInDepts"`    // 在对应的部门中的排序，Map结构的json字符串，key是部门的Id，value是人员在这个部门的排序值
	IsAdmin         bool   `json:"isAdmin"`         // 是否为企业的管理员，true表示是，false表示不是
	IsBoss          bool   `json:"isBoss"`          // 是否为企业的老板，true表示是，false表示不是
	IsLeaderInDepts string `json:"isLeaderInDepts"` // 在对应的部门中是否为主管：Map结构的json字符串，key是部门的Id，value是人员在这个部门中是否为主管，true表示是，false表示不是
	IsHide          bool   `json:"isHide"`          // 是否号码隐藏，true表示隐藏，false表示不隐藏
	Department      []int  `json:"department"`      // 成员所属部门id列表
	Position        string `json:"position"`        // 职位信息
	Avatar          string `json:"avatar"`          // 头像url
	HiredDate       int    `json:"hiredDate"`       // 入职时间。Unix时间戳 （在OA后台通讯录中的员工基础信息中维护过入职时间才会返回)
	JobNumber       string `json:"jobnumber"`       // 员工工号
	IsSenior        bool   `json:"isSenior"`        // 部门的主管列表，取值为由主管的userid组成的字符串，不同的userid使用“|”符号进行分割
	Roles           []Role `json:"roles"`           // 部门标识字段，开发者可用该字段来唯一标识一个部门，并与钉钉外部通讯录里的部门做映射
}

type Role struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	GroupName string `json:"groupName"`
	Type      int    `json:"type"`
}

// ending ----------------------------------------------------
