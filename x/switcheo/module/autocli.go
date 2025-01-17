package switcheo

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"

	modulev1 "switcheo/api/switcheo/switcheo"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: modulev1.Query_ServiceDesc.ServiceName,
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Shows the parameters of the module",
				},
				{
					RpcMethod: "ItemAll",
					Use:       "list-item",
					Short:     "List all item",
				},
				{
					RpcMethod:      "Item",
					Use:            "show-item [id]",
					Short:          "Shows a item by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/query
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:              modulev1.Msg_ServiceDesc.ServiceName,
			EnhanceCustomCommand: true, // only required if you want to use the custom command
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "UpdateParams",
					Skip:      true, // skipped because authority gated
				},
				{
					RpcMethod:      "CreateItem",
					Use:            "create-item [name] [productType] [amount] [price] [discounted]",
					Short:          "Create item",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}, {ProtoField: "productType"}, {ProtoField: "amount"}, {ProtoField: "price"}, {ProtoField: "discounted"}},
				},
				{
					RpcMethod:      "UpdateItem",
					Use:            "update-item [id] [name] [productType] [amount] [price] [discounted]",
					Short:          "Update item",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}, {ProtoField: "name"}, {ProtoField: "productType"}, {ProtoField: "amount"}, {ProtoField: "price"}, {ProtoField: "discounted"}},
				},
				{
					RpcMethod:      "DeleteItem",
					Use:            "delete-item [id]",
					Short:          "Delete item",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				// this line is used by ignite scaffolding # autocli/tx
			},
		},
	}
}
