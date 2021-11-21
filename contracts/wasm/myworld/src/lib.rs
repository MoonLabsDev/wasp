// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]
#![allow(unused_imports)]

use myworld::*;
use wasmlib::*;
use wasmlib::host::*;

use crate::consts::*;
use crate::keys::*;
use crate::params::*;
use crate::results::*;
use crate::state::*;

mod consts;
mod contract;
mod keys;
mod params;
mod results;
mod state;
mod structs;
mod myworld;

#[no_mangle]
fn on_load() {
    let exports = ScExports::new();
    exports.add_func(FUNC_DEPOSIT_TREASURE,  func_deposit_treasure_thunk);
    exports.add_func(FUNC_INIT,              func_init_thunk);
    exports.add_func(FUNC_SET_OWNER,         func_set_owner_thunk);
    exports.add_view(VIEW_GET_ALL_TREASURES, view_get_all_treasures_thunk);
    exports.add_view(VIEW_GET_OWNER,         view_get_owner_thunk);

    unsafe {
        for i in 0..KEY_MAP_LEN {
            IDX_MAP[i] = get_key_id_from_string(KEY_MAP[i]);
        }
    }
}

pub struct DepositTreasureContext {
	params: ImmutableDepositTreasureParams,
	state: MutableMyWorldState,
}

fn func_deposit_treasure_thunk(ctx: &ScFuncContext) {
	ctx.log("myworld.funcDepositTreasure");
	let f = DepositTreasureContext {
		params: ImmutableDepositTreasureParams {
			id: OBJ_ID_PARAMS,
		},
		state: MutableMyWorldState {
			id: OBJ_ID_STATE,
		},
	};
	ctx.require(f.params.treasure().exists(), "missing mandatory treasure");
	func_deposit_treasure(ctx, &f);
	ctx.log("myworld.funcDepositTreasure ok");
}

pub struct InitContext {
	params: ImmutableInitParams,
	state: MutableMyWorldState,
}

fn func_init_thunk(ctx: &ScFuncContext) {
	ctx.log("myworld.funcInit");
	let f = InitContext {
		params: ImmutableInitParams {
			id: OBJ_ID_PARAMS,
		},
		state: MutableMyWorldState {
			id: OBJ_ID_STATE,
		},
	};
	func_init(ctx, &f);
	ctx.log("myworld.funcInit ok");
}

pub struct SetOwnerContext {
	params: ImmutableSetOwnerParams,
	state: MutableMyWorldState,
}

fn func_set_owner_thunk(ctx: &ScFuncContext) {
	ctx.log("myworld.funcSetOwner");

	// current owner of this smart contract
	let access = ctx.state().get_agent_id("owner");
	ctx.require(access.exists(), "access not set: owner");
	ctx.require(ctx.caller() == access.value(), "no permission");

	let f = SetOwnerContext {
		params: ImmutableSetOwnerParams {
			id: OBJ_ID_PARAMS,
		},
		state: MutableMyWorldState {
			id: OBJ_ID_STATE,
		},
	};
	ctx.require(f.params.owner().exists(), "missing mandatory owner");
	func_set_owner(ctx, &f);
	ctx.log("myworld.funcSetOwner ok");
}

pub struct GetAllTreasuresContext {
	results: MutableGetAllTreasuresResults,
	state: ImmutableMyWorldState,
}

fn view_get_all_treasures_thunk(ctx: &ScViewContext) {
	ctx.log("myworld.viewGetAllTreasures");
	let f = GetAllTreasuresContext {
		results: MutableGetAllTreasuresResults {
			id: OBJ_ID_RESULTS,
		},
		state: ImmutableMyWorldState {
			id: OBJ_ID_STATE,
		},
	};
	view_get_all_treasures(ctx, &f);
	ctx.log("myworld.viewGetAllTreasures ok");
}

pub struct GetOwnerContext {
	results: MutableGetOwnerResults,
	state: ImmutableMyWorldState,
}

fn view_get_owner_thunk(ctx: &ScViewContext) {
	ctx.log("myworld.viewGetOwner");
	let f = GetOwnerContext {
		results: MutableGetOwnerResults {
			id: OBJ_ID_RESULTS,
		},
		state: ImmutableMyWorldState {
			id: OBJ_ID_STATE,
		},
	};
	view_get_owner(ctx, &f);
	ctx.log("myworld.viewGetOwner ok");
}
