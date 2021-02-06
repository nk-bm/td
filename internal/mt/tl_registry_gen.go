// Code generated by gotdgen, DO NOT EDIT.

package mt

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// Types returns mapping from type ids to TL type names.
func TypesMap() map[uint32]string {
	return map[uint32]string{
		0x5162463:  "resPQ#5162463",
		0x83c95aec: "p_q_inner_data#83c95aec",
		0x79cb045d: "server_DH_params_fail#79cb045d",
		0xd0e8075c: "server_DH_params_ok#d0e8075c",
		0xb5890dba: "server_DH_inner_data#b5890dba",
		0x6643b654: "client_DH_inner_data#6643b654",
		0x3bcbf734: "dh_gen_ok#3bcbf734",
		0x46dc1fb9: "dh_gen_retry#46dc1fb9",
		0xa69dae02: "dh_gen_fail#a69dae02",
		0xf35c6d01: "rpc_result#f35c6d01",
		0x2144ca19: "rpc_error#2144ca19",
		0x5e2ad36e: "rpc_answer_unknown#5e2ad36e",
		0xcd78e586: "rpc_answer_dropped_running#cd78e586",
		0xa43ad8b7: "rpc_answer_dropped#a43ad8b7",
		0x949d9dc:  "future_salt#949d9dc",
		0xae500895: "future_salts#ae500895",
		0x347773c5: "pong#347773c5",
		0xe22045fc: "destroy_session_ok#e22045fc",
		0x62d350c9: "destroy_session_none#62d350c9",
		0x9ec20908: "new_session_created#9ec20908",
		0x5bb8e511: "message#5bb8e511",
		0xe06046b2: "msg_copy#e06046b2",
		0x3072cfa1: "gzip_packed#3072cfa1",
		0x62d6b459: "msgs_ack#62d6b459",
		0xa7eff811: "bad_msg_notification#a7eff811",
		0xedab447b: "bad_server_salt#edab447b",
		0x7d861a08: "msg_resend_req#7d861a08",
		0xda69fb52: "msgs_state_req#da69fb52",
		0x4deb57d:  "msgs_state_info#4deb57d",
		0x8cc0d131: "msgs_all_info#8cc0d131",
		0x276d3ec6: "msg_detailed_info#276d3ec6",
		0x809db6df: "msg_new_detailed_info#809db6df",
		0x60469778: "req_pq#60469778",
		0xbe7e8ef1: "req_pq_multi#be7e8ef1",
		0xd712e4be: "req_DH_params#d712e4be",
		0xf5045f1f: "set_client_DH_params#f5045f1f",
		0x58e4a740: "rpc_drop_answer#58e4a740",
		0xb921bd04: "get_future_salts#b921bd04",
		0x7abe77ec: "ping#7abe77ec",
		0xf3427b8c: "ping_delay_disconnect#f3427b8c",
		0xe7512126: "destroy_session#e7512126",
		0x9299359f: "http_wait#9299359f",
	}
}

// TypesConstructorMap maps type ids to constructors.
func TypesConstructorMap() map[uint32]func() bin.Object {
	return map[uint32]func() bin.Object{
		0x5162463:  func() bin.Object { return &ResPQ{} },
		0x83c95aec: func() bin.Object { return &PQInnerData{} },
		0x79cb045d: func() bin.Object { return &ServerDHParamsFail{} },
		0xd0e8075c: func() bin.Object { return &ServerDHParamsOk{} },
		0xb5890dba: func() bin.Object { return &ServerDHInnerData{} },
		0x6643b654: func() bin.Object { return &ClientDHInnerData{} },
		0x3bcbf734: func() bin.Object { return &DhGenOk{} },
		0x46dc1fb9: func() bin.Object { return &DhGenRetry{} },
		0xa69dae02: func() bin.Object { return &DhGenFail{} },
		0xf35c6d01: func() bin.Object { return &RPCResult{} },
		0x2144ca19: func() bin.Object { return &RPCError{} },
		0x5e2ad36e: func() bin.Object { return &RPCAnswerUnknown{} },
		0xcd78e586: func() bin.Object { return &RPCAnswerDroppedRunning{} },
		0xa43ad8b7: func() bin.Object { return &RPCAnswerDropped{} },
		0x949d9dc:  func() bin.Object { return &FutureSalt{} },
		0xae500895: func() bin.Object { return &FutureSalts{} },
		0x347773c5: func() bin.Object { return &Pong{} },
		0xe22045fc: func() bin.Object { return &DestroySessionOk{} },
		0x62d350c9: func() bin.Object { return &DestroySessionNone{} },
		0x9ec20908: func() bin.Object { return &NewSessionCreated{} },
		0x5bb8e511: func() bin.Object { return &Message{} },
		0xe06046b2: func() bin.Object { return &MsgCopy{} },
		0x3072cfa1: func() bin.Object { return &GzipPacked{} },
		0x62d6b459: func() bin.Object { return &MsgsAck{} },
		0xa7eff811: func() bin.Object { return &BadMsgNotification{} },
		0xedab447b: func() bin.Object { return &BadServerSalt{} },
		0x7d861a08: func() bin.Object { return &MsgResendReq{} },
		0xda69fb52: func() bin.Object { return &MsgsStateReq{} },
		0x4deb57d:  func() bin.Object { return &MsgsStateInfo{} },
		0x8cc0d131: func() bin.Object { return &MsgsAllInfo{} },
		0x276d3ec6: func() bin.Object { return &MsgDetailedInfo{} },
		0x809db6df: func() bin.Object { return &MsgNewDetailedInfo{} },
		0x60469778: func() bin.Object { return &ReqPqRequest{} },
		0xbe7e8ef1: func() bin.Object { return &ReqPqMultiRequest{} },
		0xd712e4be: func() bin.Object { return &ReqDHParamsRequest{} },
		0xf5045f1f: func() bin.Object { return &SetClientDHParamsRequest{} },
		0x58e4a740: func() bin.Object { return &RPCDropAnswerRequest{} },
		0xb921bd04: func() bin.Object { return &GetFutureSaltsRequest{} },
		0x7abe77ec: func() bin.Object { return &PingRequest{} },
		0xf3427b8c: func() bin.Object { return &PingDelayDisconnectRequest{} },
		0xe7512126: func() bin.Object { return &DestroySessionRequest{} },
		0x9299359f: func() bin.Object { return &HTTPWaitRequest{} },
	}
}
