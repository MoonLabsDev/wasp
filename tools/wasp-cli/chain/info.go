package chain

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/iotaledger/wasp/packages/kv/codec"
	"github.com/iotaledger/wasp/packages/kv/datatypes"
	"github.com/iotaledger/wasp/packages/vm/builtinvm/root"
	"github.com/iotaledger/wasp/tools/wasp-cli/config"
	"github.com/iotaledger/wasp/tools/wasp-cli/log"
)

func infoCmd(args []string) {
	chain, err := config.WaspClient().GetChainRecord(GetCurrentChainID())
	log.Check(err)

	log.Printf("Chain ID: %s\n", chain.ChainID)
	log.Printf("Chain Address: %s\n", address.Address(chain.ChainID))
	log.Printf("Chain Color: %s\n", chain.Color)
	log.Printf("Committee nodes: %+v\n", chain.CommitteeNodes)
	log.Printf("Active: %v\n", chain.Active)

	if chain.Active {
		info, err := SCClient(root.Interface.Hname()).CallView(root.FuncGetChainInfo, nil)
		log.Check(err)

		description, _, err := codec.DecodeString(info.MustGet(root.VarDescription))
		log.Check(err)
		log.Printf("Description: %s\n", description)

		contracts, err := root.DecodeContractRegistry(datatypes.NewMustMap(info, root.VarContractRegistry))
		log.Check(err)
		log.Printf("#Contracts: %d\n", len(contracts))

		ownerID, _, err := codec.DecodeAgentID(info.MustGet(root.VarChainOwnerID))
		log.Check(err)
		log.Printf("Owner: %s\n", ownerID)

		delegated, ok, err := codec.DecodeAgentID(info.MustGet(root.VarChainOwnerIDDelegated))
		log.Check(err)
		if ok {
			log.Printf("Delegated owner: %s\n", delegated)
		}

		feeColor, defaultOwnerFee, defaultValidatorFee, err := root.GetDefaultFeeInfo(info)
		log.Check(err)
		log.Printf("Default owner fee: %d %s\n", defaultOwnerFee, feeColor)
		log.Printf("Default validator fee: %d %s\n", defaultValidatorFee, feeColor)
	}
}
