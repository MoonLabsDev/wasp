// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

#![allow(dead_code)]

use wasmlib::*;

pub const SC_NAME        : &str = "myworld";
pub const SC_DESCRIPTION : &str = "MyWorld description";
pub const HSC_NAME       : ScHname = ScHname(0x5efde160);

pub const PARAM_OWNER    : &str = "owner";
pub const PARAM_TREASURE : &str = "treasure";

pub const RESULT_OWNER     : &str = "owner";
pub const RESULT_TREASURES : &str = "treasures";

pub const STATE_OWNER     : &str = "owner";
pub const STATE_TREASURES : &str = "treasures";

pub const FUNC_DEPOSIT_TREASURE  : &str = "depositTreasure";
pub const FUNC_INIT              : &str = "init";
pub const FUNC_SET_OWNER         : &str = "setOwner";
pub const VIEW_GET_ALL_TREASURES : &str = "getAllTreasures";
pub const VIEW_GET_OWNER         : &str = "getOwner";

pub const HFUNC_DEPOSIT_TREASURE  : ScHname = ScHname(0x186cdf0f);
pub const HFUNC_INIT              : ScHname = ScHname(0x1f44d644);
pub const HFUNC_SET_OWNER         : ScHname = ScHname(0x2a15fe7b);
pub const HVIEW_GET_ALL_TREASURES : ScHname = ScHname(0x193d3714);
pub const HVIEW_GET_OWNER         : ScHname = ScHname(0x137107a6);
