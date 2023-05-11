package cosmos

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	authSigning "github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/pkg/errors"
)

// Build performs WithMsgs and SignTx.
func (t *Tx) Build(msgs ...sdk.Msg) (client.TxBuilder, error) {
	txBuilder, err := t.WithMsgs(msgs...)
	if err != nil {
		return nil, err
	}

	err = t.SignTx(txBuilder)
	if err != nil {
		return nil, err
	}

	return txBuilder, nil
}

// WithMsgs creates an unsigned transaction builder from the given types.Msg's.
func (t *Tx) WithMsgs(msgs ...sdk.Msg) (client.TxBuilder, error) {
	return t.txf.BuildUnsignedTx(msgs...)
}

// SignTx performs the transaction signing.
func (t *Tx) SignTx(txBuilder client.TxBuilder) error {
	pubKey := t.params.MustGetPrivateKey().PubKey()

	err := t.prepareSignTx()
	if err != nil {
		return errors.Wrap(err, "prepareSignTx")
	}

	sigV2 := signing.SignatureV2{
		PubKey: pubKey,
		Data: &signing.SingleSignatureData{
			SignMode:  t.txf.SignMode(),
			Signature: nil,
		},
		Sequence: t.txf.Sequence(),
	}

	if err = txBuilder.SetSignatures(sigV2); err != nil {
		return errors.Wrap(err, "SetSignatures")
	}

	// Construct the SignatureV2 struct
	signerData := authSigning.SignerData{
		ChainID:       t.chainID,
		AccountNumber: t.txf.AccountNumber(),
		Sequence:      t.txf.Sequence(),
	}

	signWithPrivKey, err := tx.SignWithPrivKey(
		t.txf.SignMode(),
		signerData,
		txBuilder,
		t.params.MustGetPrivateKey().PrivKey,
		t.txConfig,
		t.txf.Sequence())

	if err != nil {
		return errors.Wrap(err, "SignWithPrivKey")
	}

	err = txBuilder.SetSignatures(signWithPrivKey)
	if err != nil {
		return errors.Wrap(err, "SetSignatures")
	}

	return nil
}

// PrintUnsignedTx displays the unsigned transaction from the given types.Msg's.
func (t *Tx) PrintUnsignedTx(msg sdk.Msg) (string, error) {
	unsignedTx, err := t.WithMsgs(msg)
	if err != nil {
		return "", errors.Wrap(err, "BuildUnsignedTx")
	}

	json, err := t.txConfig.TxJSONEncoder()(unsignedTx.GetTx())
	if err != nil {
		return "", errors.Wrap(err, "TxJSONEncoder")
	}

	return string(json), nil
}

// prepareSignTx performs the preparation for signing the transaction.
func (t *Tx) prepareSignTx() error {
	from := t.params.MustGetPrivateKey().AccAddress()

	if err := t.accountRetriever.AccountExists(from.String()); err != nil {
		return errors.Wrap(ErrAccountNotExisted, err.Error())
	}

	initNum, initSeq := t.txf.AccountNumber(), t.txf.Sequence()
	if initNum == 0 || initSeq == 0 {
		var accNum, accSeq uint64
		var err error

		cosmosAccount, err := t.accountRetriever.AccountInfo(from.String())
		if err != nil {
			return errors.Wrap(err, "CosmosAccount")
		}

		accNum = cosmosAccount.GetAccountNumber()
		accSeq = cosmosAccount.GetSequence()

		t.txf = t.txf.WithAccountNumber(accNum)
		t.txf = t.txf.WithSequence(accSeq)
	}

	return nil
}
