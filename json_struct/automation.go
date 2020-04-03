package json_struct

type ActionDetail struct {
	ActionId string `json:"id"`
	ActionName string `json:"action_name"`
	ActionIcon string `json:"action_icon"`
	ActionPosition string `json:"action_position"`
	ActionType string `json:"action_type"`
	BehaviorInstanceList []*BehaviorDetail `json:"behavior_instance_list"`
}

type BehaviorDetail struct {
	Key string `json:"key"`
	Value string `json:"value"`
	Position string `json:"position"`
	Options []string `json:"options"`
}

type GetActionListRequest struct {
	UserId string `json:"user_id"`
	ActionType string `json:"action_type"`
}

type GetActionListResponse struct {
	ActionList []*ActionDetail `json:"action_list"`
	Base BaseResponse `json:"base"`
}

type GetActionDefinationRequest struct {
	UserId string `json:"user_id"`
	ActionId string `json:"action_id"`
}

type GetActionDefinationResponse struct {
	BehaviorDefinationList []*BehaviorDetail `json:"behavior_defination_list"`
	Base BaseResponse `json:"base"`
}

type GetFlowDefinationDetailRequest struct {
	FlowDefinationId string `json:"flow_defination_id"`
	UserId string `json:"user_id"`
}

type GetFlowDefinationDetailResponse struct {
	ActionList []*ActionDetail `json:"action_list"`
	Base BaseResponse `json:"base"`
}

type CreateFlowDefinationRequest struct {
	UserId string `json:"user_id"`
	FlowDefinationName string `json:"flow_defination_name"`
	ActionList []*ActionDetail `json:"action_list"`
}

type CreateFlowDefinationResponse struct {
	FlowDefinationId string `json:"flow_defination_id"`
	Base BaseResponse `json:"base"`
}

type UpdateFlowDefinationRequest struct {
	UserId int64 `json:"user_id"`
	FlowDefinationId string `json:"flow_defination_id"`
	ActionList []*ActionDetail `json:"action_list"`
}

type UpdateFlowDefinationResponse struct {
	Base BaseResponse `json:"base"`
}

type GetFlowDefinationListRequest struct {
	UserId string `json:"user_id"`
}

type GetFlowDefinationListResponse struct {
	ActionList []*ActionDetail `json:"action_list"`
	Base BaseResponse `json:"base"`
}