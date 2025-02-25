// Code generated by fastssz. DO NOT EDIT.
// Hash: 3fdd5c5ff5c2e0e3e59c8a72be771c4d964b8c6d4716aac215372db3430a4eed
package bellatrix

import (
	"github.com/attestantio/go-eth2-client/spec/altair"
	"github.com/attestantio/go-eth2-client/spec/phase0"
	ssz "github.com/ferranbt/fastssz"
)

// MarshalSSZ ssz marshals the BeaconState object
func (b *BeaconState) MarshalSSZ() ([]byte, error) {
	return ssz.MarshalSSZ(b)
}

// MarshalSSZTo ssz marshals the BeaconState object to a target array
func (b *BeaconState) MarshalSSZTo(buf []byte) (dst []byte, err error) {
	dst = buf
	offset := int(2736633)

	// Field (0) 'GenesisTime'
	dst = ssz.MarshalUint64(dst, b.GenesisTime)

	// Field (1) 'GenesisValidatorsRoot'
	if size := len(b.GenesisValidatorsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("BeaconState.GenesisValidatorsRoot", size, 32)
		return
	}
	dst = append(dst, b.GenesisValidatorsRoot[:]...)

	// Field (2) 'Slot'
	dst = ssz.MarshalUint64(dst, uint64(b.Slot))

	// Field (3) 'Fork'
	if b.Fork == nil {
		b.Fork = new(phase0.Fork)
	}
	if dst, err = b.Fork.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (4) 'LatestBlockHeader'
	if b.LatestBlockHeader == nil {
		b.LatestBlockHeader = new(phase0.BeaconBlockHeader)
	}
	if dst, err = b.LatestBlockHeader.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (5) 'BlockRoots'
	if size := len(b.BlockRoots); size != 8192 {
		err = ssz.ErrVectorLengthFn("BeaconState.BlockRoots", size, 8192)
		return
	}
	for ii := 0; ii < 8192; ii++ {
		if size := len(b.BlockRoots[ii]); size != 32 {
			err = ssz.ErrBytesLengthFn("BeaconState.BlockRoots[ii]", size, 32)
			return
		}
		dst = append(dst, b.BlockRoots[ii][:]...)
	}

	// Field (6) 'StateRoots'
	if size := len(b.StateRoots); size != 8192 {
		err = ssz.ErrVectorLengthFn("BeaconState.StateRoots", size, 8192)
		return
	}
	for ii := 0; ii < 8192; ii++ {
		if size := len(b.StateRoots[ii]); size != 32 {
			err = ssz.ErrBytesLengthFn("BeaconState.StateRoots[ii]", size, 32)
			return
		}
		dst = append(dst, b.StateRoots[ii][:]...)
	}

	// Offset (7) 'HistoricalRoots'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.HistoricalRoots) * 32

	// Field (8) 'ETH1Data'
	if b.ETH1Data == nil {
		b.ETH1Data = new(phase0.ETH1Data)
	}
	if dst, err = b.ETH1Data.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (9) 'ETH1DataVotes'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.ETH1DataVotes) * 72

	// Field (10) 'ETH1DepositIndex'
	dst = ssz.MarshalUint64(dst, b.ETH1DepositIndex)

	// Offset (11) 'Validators'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Validators) * 121

	// Offset (12) 'Balances'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.Balances) * 8

	// Field (13) 'RANDAOMixes'
	if size := len(b.RANDAOMixes); size != 65536 {
		err = ssz.ErrVectorLengthFn("BeaconState.RANDAOMixes", size, 65536)
		return
	}
	for ii := 0; ii < 65536; ii++ {
		if size := len(b.RANDAOMixes[ii]); size != 32 {
			err = ssz.ErrBytesLengthFn("BeaconState.RANDAOMixes[ii]", size, 32)
			return
		}
		dst = append(dst, b.RANDAOMixes[ii][:]...)
	}

	// Field (14) 'Slashings'
	if size := len(b.Slashings); size != 8192 {
		err = ssz.ErrVectorLengthFn("BeaconState.Slashings", size, 8192)
		return
	}
	for ii := 0; ii < 8192; ii++ {
		dst = ssz.MarshalUint64(dst, uint64(b.Slashings[ii]))
	}

	// Offset (15) 'PreviousEpochParticipation'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.PreviousEpochParticipation) * 1

	// Offset (16) 'CurrentEpochParticipation'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.CurrentEpochParticipation) * 1

	// Field (17) 'JustificationBits'
	if size := len(b.JustificationBits); size != 1 {
		err = ssz.ErrBytesLengthFn("BeaconState.JustificationBits", size, 1)
		return
	}
	dst = append(dst, b.JustificationBits...)

	// Field (18) 'PreviousJustifiedCheckpoint'
	if b.PreviousJustifiedCheckpoint == nil {
		b.PreviousJustifiedCheckpoint = new(phase0.Checkpoint)
	}
	if dst, err = b.PreviousJustifiedCheckpoint.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (19) 'CurrentJustifiedCheckpoint'
	if b.CurrentJustifiedCheckpoint == nil {
		b.CurrentJustifiedCheckpoint = new(phase0.Checkpoint)
	}
	if dst, err = b.CurrentJustifiedCheckpoint.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (20) 'FinalizedCheckpoint'
	if b.FinalizedCheckpoint == nil {
		b.FinalizedCheckpoint = new(phase0.Checkpoint)
	}
	if dst, err = b.FinalizedCheckpoint.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (21) 'InactivityScores'
	dst = ssz.WriteOffset(dst, offset)
	offset += len(b.InactivityScores) * 8

	// Field (22) 'CurrentSyncCommittee'
	if b.CurrentSyncCommittee == nil {
		b.CurrentSyncCommittee = new(altair.SyncCommittee)
	}
	if dst, err = b.CurrentSyncCommittee.MarshalSSZTo(dst); err != nil {
		return
	}

	// Field (23) 'NextSyncCommittee'
	if b.NextSyncCommittee == nil {
		b.NextSyncCommittee = new(altair.SyncCommittee)
	}
	if dst, err = b.NextSyncCommittee.MarshalSSZTo(dst); err != nil {
		return
	}

	// Offset (24) 'LatestExecutionPayloadHeader'
	dst = ssz.WriteOffset(dst, offset)
	if b.LatestExecutionPayloadHeader == nil {
		b.LatestExecutionPayloadHeader = new(ExecutionPayloadHeader)
	}
	offset += b.LatestExecutionPayloadHeader.SizeSSZ()

	// Field (7) 'HistoricalRoots'
	if size := len(b.HistoricalRoots); size > 16777216 {
		err = ssz.ErrListTooBigFn("BeaconState.HistoricalRoots", size, 16777216)
		return
	}
	for ii := 0; ii < len(b.HistoricalRoots); ii++ {
		if size := len(b.HistoricalRoots[ii]); size != 32 {
			err = ssz.ErrBytesLengthFn("BeaconState.HistoricalRoots[ii]", size, 32)
			return
		}
		dst = append(dst, b.HistoricalRoots[ii][:]...)
	}

	// Field (9) 'ETH1DataVotes'
	if size := len(b.ETH1DataVotes); size > 2048 {
		err = ssz.ErrListTooBigFn("BeaconState.ETH1DataVotes", size, 2048)
		return
	}
	for ii := 0; ii < len(b.ETH1DataVotes); ii++ {
		if dst, err = b.ETH1DataVotes[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (11) 'Validators'
	if size := len(b.Validators); size > 1099511627776 {
		err = ssz.ErrListTooBigFn("BeaconState.Validators", size, 1099511627776)
		return
	}
	for ii := 0; ii < len(b.Validators); ii++ {
		if dst, err = b.Validators[ii].MarshalSSZTo(dst); err != nil {
			return
		}
	}

	// Field (12) 'Balances'
	if size := len(b.Balances); size > 1099511627776 {
		err = ssz.ErrListTooBigFn("BeaconState.Balances", size, 1099511627776)
		return
	}
	for ii := 0; ii < len(b.Balances); ii++ {
		dst = ssz.MarshalUint64(dst, uint64(b.Balances[ii]))
	}

	// Field (15) 'PreviousEpochParticipation'
	if size := len(b.PreviousEpochParticipation); size > 1099511627776 {
		err = ssz.ErrListTooBigFn("BeaconState.PreviousEpochParticipation", size, 1099511627776)
		return
	}
	for ii := 0; ii < len(b.PreviousEpochParticipation); ii++ {
		dst = ssz.MarshalUint8(dst, uint8(b.PreviousEpochParticipation[ii]))
	}

	// Field (16) 'CurrentEpochParticipation'
	if size := len(b.CurrentEpochParticipation); size > 1099511627776 {
		err = ssz.ErrListTooBigFn("BeaconState.CurrentEpochParticipation", size, 1099511627776)
		return
	}
	for ii := 0; ii < len(b.CurrentEpochParticipation); ii++ {
		dst = ssz.MarshalUint8(dst, uint8(b.CurrentEpochParticipation[ii]))
	}

	// Field (21) 'InactivityScores'
	if size := len(b.InactivityScores); size > 1099511627776 {
		err = ssz.ErrListTooBigFn("BeaconState.InactivityScores", size, 1099511627776)
		return
	}
	for ii := 0; ii < len(b.InactivityScores); ii++ {
		dst = ssz.MarshalUint64(dst, b.InactivityScores[ii])
	}

	// Field (24) 'LatestExecutionPayloadHeader'
	if dst, err = b.LatestExecutionPayloadHeader.MarshalSSZTo(dst); err != nil {
		return
	}

	return
}

// UnmarshalSSZ ssz unmarshals the BeaconState object
func (b *BeaconState) UnmarshalSSZ(buf []byte) error {
	var err error
	size := uint64(len(buf))
	if size < 2736633 {
		return ssz.ErrSize
	}

	tail := buf
	var o7, o9, o11, o12, o15, o16, o21, o24 uint64

	// Field (0) 'GenesisTime'
	b.GenesisTime = ssz.UnmarshallUint64(buf[0:8])

	// Field (1) 'GenesisValidatorsRoot'
	copy(b.GenesisValidatorsRoot[:], buf[8:40])

	// Field (2) 'Slot'
	b.Slot = phase0.Slot(ssz.UnmarshallUint64(buf[40:48]))

	// Field (3) 'Fork'
	if b.Fork == nil {
		b.Fork = new(phase0.Fork)
	}
	if err = b.Fork.UnmarshalSSZ(buf[48:64]); err != nil {
		return err
	}

	// Field (4) 'LatestBlockHeader'
	if b.LatestBlockHeader == nil {
		b.LatestBlockHeader = new(phase0.BeaconBlockHeader)
	}
	if err = b.LatestBlockHeader.UnmarshalSSZ(buf[64:176]); err != nil {
		return err
	}

	// Field (5) 'BlockRoots'
	b.BlockRoots = make([]phase0.Root, 8192)
	for ii := 0; ii < 8192; ii++ {
		copy(b.BlockRoots[ii][:], buf[176:262320][ii*32:(ii+1)*32])
	}

	// Field (6) 'StateRoots'
	b.StateRoots = make([]phase0.Root, 8192)
	for ii := 0; ii < 8192; ii++ {
		copy(b.StateRoots[ii][:], buf[262320:524464][ii*32:(ii+1)*32])
	}

	// Offset (7) 'HistoricalRoots'
	if o7 = ssz.ReadOffset(buf[524464:524468]); o7 > size {
		return ssz.ErrOffset
	}

	if o7 < 2736633 {
		return ssz.ErrInvalidVariableOffset
	}

	// Field (8) 'ETH1Data'
	if b.ETH1Data == nil {
		b.ETH1Data = new(phase0.ETH1Data)
	}
	if err = b.ETH1Data.UnmarshalSSZ(buf[524468:524540]); err != nil {
		return err
	}

	// Offset (9) 'ETH1DataVotes'
	if o9 = ssz.ReadOffset(buf[524540:524544]); o9 > size || o7 > o9 {
		return ssz.ErrOffset
	}

	// Field (10) 'ETH1DepositIndex'
	b.ETH1DepositIndex = ssz.UnmarshallUint64(buf[524544:524552])

	// Offset (11) 'Validators'
	if o11 = ssz.ReadOffset(buf[524552:524556]); o11 > size || o9 > o11 {
		return ssz.ErrOffset
	}

	// Offset (12) 'Balances'
	if o12 = ssz.ReadOffset(buf[524556:524560]); o12 > size || o11 > o12 {
		return ssz.ErrOffset
	}

	// Field (13) 'RANDAOMixes'
	b.RANDAOMixes = make([]phase0.Root, 65536)
	for ii := 0; ii < 65536; ii++ {
		copy(b.RANDAOMixes[ii][:], buf[524560:2621712][ii*32:(ii+1)*32])
	}

	// Field (14) 'Slashings'
	b.Slashings = make([]phase0.Gwei, 8192)
	for ii := 0; ii < 8192; ii++ {
		b.Slashings[ii] = phase0.Gwei(ssz.UnmarshallUint64(buf[2621712:2687248][ii*8 : (ii+1)*8]))
	}

	// Offset (15) 'PreviousEpochParticipation'
	if o15 = ssz.ReadOffset(buf[2687248:2687252]); o15 > size || o12 > o15 {
		return ssz.ErrOffset
	}

	// Offset (16) 'CurrentEpochParticipation'
	if o16 = ssz.ReadOffset(buf[2687252:2687256]); o16 > size || o15 > o16 {
		return ssz.ErrOffset
	}

	// Field (17) 'JustificationBits'
	if cap(b.JustificationBits) == 0 {
		b.JustificationBits = make([]byte, 0, len(buf[2687256:2687257]))
	}
	b.JustificationBits = append(b.JustificationBits, buf[2687256:2687257]...)

	// Field (18) 'PreviousJustifiedCheckpoint'
	if b.PreviousJustifiedCheckpoint == nil {
		b.PreviousJustifiedCheckpoint = new(phase0.Checkpoint)
	}
	if err = b.PreviousJustifiedCheckpoint.UnmarshalSSZ(buf[2687257:2687297]); err != nil {
		return err
	}

	// Field (19) 'CurrentJustifiedCheckpoint'
	if b.CurrentJustifiedCheckpoint == nil {
		b.CurrentJustifiedCheckpoint = new(phase0.Checkpoint)
	}
	if err = b.CurrentJustifiedCheckpoint.UnmarshalSSZ(buf[2687297:2687337]); err != nil {
		return err
	}

	// Field (20) 'FinalizedCheckpoint'
	if b.FinalizedCheckpoint == nil {
		b.FinalizedCheckpoint = new(phase0.Checkpoint)
	}
	if err = b.FinalizedCheckpoint.UnmarshalSSZ(buf[2687337:2687377]); err != nil {
		return err
	}

	// Offset (21) 'InactivityScores'
	if o21 = ssz.ReadOffset(buf[2687377:2687381]); o21 > size || o16 > o21 {
		return ssz.ErrOffset
	}

	// Field (22) 'CurrentSyncCommittee'
	if b.CurrentSyncCommittee == nil {
		b.CurrentSyncCommittee = new(altair.SyncCommittee)
	}
	if err = b.CurrentSyncCommittee.UnmarshalSSZ(buf[2687381:2712005]); err != nil {
		return err
	}

	// Field (23) 'NextSyncCommittee'
	if b.NextSyncCommittee == nil {
		b.NextSyncCommittee = new(altair.SyncCommittee)
	}
	if err = b.NextSyncCommittee.UnmarshalSSZ(buf[2712005:2736629]); err != nil {
		return err
	}

	// Offset (24) 'LatestExecutionPayloadHeader'
	if o24 = ssz.ReadOffset(buf[2736629:2736633]); o24 > size || o21 > o24 {
		return ssz.ErrOffset
	}

	// Field (7) 'HistoricalRoots'
	{
		buf = tail[o7:o9]
		num, err := ssz.DivideInt2(len(buf), 32, 16777216)
		if err != nil {
			return err
		}
		b.HistoricalRoots = make([]phase0.Root, num)
		for ii := 0; ii < num; ii++ {
			copy(b.HistoricalRoots[ii][:], buf[ii*32:(ii+1)*32])
		}
	}

	// Field (9) 'ETH1DataVotes'
	{
		buf = tail[o9:o11]
		num, err := ssz.DivideInt2(len(buf), 72, 2048)
		if err != nil {
			return err
		}
		b.ETH1DataVotes = make([]*phase0.ETH1Data, num)
		for ii := 0; ii < num; ii++ {
			if b.ETH1DataVotes[ii] == nil {
				b.ETH1DataVotes[ii] = new(phase0.ETH1Data)
			}
			if err = b.ETH1DataVotes[ii].UnmarshalSSZ(buf[ii*72 : (ii+1)*72]); err != nil {
				return err
			}
		}
	}

	// Field (11) 'Validators'
	{
		buf = tail[o11:o12]
		num, err := ssz.DivideInt2(len(buf), 121, 1099511627776)
		if err != nil {
			return err
		}
		b.Validators = make([]*phase0.Validator, num)
		for ii := 0; ii < num; ii++ {
			if b.Validators[ii] == nil {
				b.Validators[ii] = new(phase0.Validator)
			}
			if err = b.Validators[ii].UnmarshalSSZ(buf[ii*121 : (ii+1)*121]); err != nil {
				return err
			}
		}
	}

	// Field (12) 'Balances'
	{
		buf = tail[o12:o15]
		num, err := ssz.DivideInt2(len(buf), 8, 1099511627776)
		if err != nil {
			return err
		}
		b.Balances = make([]phase0.Gwei, num)
		for ii := 0; ii < num; ii++ {
			b.Balances[ii] = phase0.Gwei(ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8]))
		}
	}

	// Field (15) 'PreviousEpochParticipation'
	{
		buf = tail[o15:o16]
		num, err := ssz.DivideInt2(len(buf), 1, 1099511627776)
		if err != nil {
			return err
		}
		b.PreviousEpochParticipation = make([]altair.ParticipationFlags, num)
		for ii := 0; ii < num; ii++ {
			b.PreviousEpochParticipation[ii] = altair.ParticipationFlags(ssz.UnmarshallUint8(buf[ii*1 : (ii+1)*1]))
		}
	}

	// Field (16) 'CurrentEpochParticipation'
	{
		buf = tail[o16:o21]
		num, err := ssz.DivideInt2(len(buf), 1, 1099511627776)
		if err != nil {
			return err
		}
		b.CurrentEpochParticipation = make([]altair.ParticipationFlags, num)
		for ii := 0; ii < num; ii++ {
			b.CurrentEpochParticipation[ii] = altair.ParticipationFlags(ssz.UnmarshallUint8(buf[ii*1 : (ii+1)*1]))
		}
	}

	// Field (21) 'InactivityScores'
	{
		buf = tail[o21:o24]
		num, err := ssz.DivideInt2(len(buf), 8, 1099511627776)
		if err != nil {
			return err
		}
		b.InactivityScores = ssz.ExtendUint64(b.InactivityScores, num)
		for ii := 0; ii < num; ii++ {
			b.InactivityScores[ii] = ssz.UnmarshallUint64(buf[ii*8 : (ii+1)*8])
		}
	}

	// Field (24) 'LatestExecutionPayloadHeader'
	{
		buf = tail[o24:]
		if b.LatestExecutionPayloadHeader == nil {
			b.LatestExecutionPayloadHeader = new(ExecutionPayloadHeader)
		}
		if err = b.LatestExecutionPayloadHeader.UnmarshalSSZ(buf); err != nil {
			return err
		}
	}
	return err
}

// SizeSSZ returns the ssz encoded size in bytes for the BeaconState object
func (b *BeaconState) SizeSSZ() (size int) {
	size = 2736633

	// Field (7) 'HistoricalRoots'
	size += len(b.HistoricalRoots) * 32

	// Field (9) 'ETH1DataVotes'
	size += len(b.ETH1DataVotes) * 72

	// Field (11) 'Validators'
	size += len(b.Validators) * 121

	// Field (12) 'Balances'
	size += len(b.Balances) * 8

	// Field (15) 'PreviousEpochParticipation'
	size += len(b.PreviousEpochParticipation) * 1

	// Field (16) 'CurrentEpochParticipation'
	size += len(b.CurrentEpochParticipation) * 1

	// Field (21) 'InactivityScores'
	size += len(b.InactivityScores) * 8

	// Field (24) 'LatestExecutionPayloadHeader'
	if b.LatestExecutionPayloadHeader == nil {
		b.LatestExecutionPayloadHeader = new(ExecutionPayloadHeader)
	}
	size += b.LatestExecutionPayloadHeader.SizeSSZ()

	return
}

// HashTreeRoot ssz hashes the BeaconState object
func (b *BeaconState) HashTreeRoot() ([32]byte, error) {
	return ssz.HashWithDefaultHasher(b)
}

// HashTreeRootWith ssz hashes the BeaconState object with a hasher
func (b *BeaconState) HashTreeRootWith(hh ssz.HashWalker) (err error) {
	indx := hh.Index()

	// Field (0) 'GenesisTime'
	hh.PutUint64(b.GenesisTime)

	// Field (1) 'GenesisValidatorsRoot'
	if size := len(b.GenesisValidatorsRoot); size != 32 {
		err = ssz.ErrBytesLengthFn("BeaconState.GenesisValidatorsRoot", size, 32)
		return
	}
	hh.PutBytes(b.GenesisValidatorsRoot[:])

	// Field (2) 'Slot'
	hh.PutUint64(uint64(b.Slot))

	// Field (3) 'Fork'
	if b.Fork == nil {
		b.Fork = new(phase0.Fork)
	}
	if err = b.Fork.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (4) 'LatestBlockHeader'
	if b.LatestBlockHeader == nil {
		b.LatestBlockHeader = new(phase0.BeaconBlockHeader)
	}
	if err = b.LatestBlockHeader.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (5) 'BlockRoots'
	{
		if size := len(b.BlockRoots); size != 8192 {
			err = ssz.ErrVectorLengthFn("BeaconState.BlockRoots", size, 8192)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.BlockRoots {
			hh.Append(i[:])
		}
		hh.Merkleize(subIndx)
	}

	// Field (6) 'StateRoots'
	{
		if size := len(b.StateRoots); size != 8192 {
			err = ssz.ErrVectorLengthFn("BeaconState.StateRoots", size, 8192)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.StateRoots {
			hh.Append(i[:])
		}
		hh.Merkleize(subIndx)
	}

	// Field (7) 'HistoricalRoots'
	{
		if size := len(b.HistoricalRoots); size > 16777216 {
			err = ssz.ErrListTooBigFn("BeaconState.HistoricalRoots", size, 16777216)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.HistoricalRoots {
			hh.Append(i[:])
		}
		numItems := uint64(len(b.HistoricalRoots))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(16777216, numItems, 32))
	}

	// Field (8) 'ETH1Data'
	if b.ETH1Data == nil {
		b.ETH1Data = new(phase0.ETH1Data)
	}
	if err = b.ETH1Data.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (9) 'ETH1DataVotes'
	{
		subIndx := hh.Index()
		num := uint64(len(b.ETH1DataVotes))
		if num > 2048 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.ETH1DataVotes {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 2048)
	}

	// Field (10) 'ETH1DepositIndex'
	hh.PutUint64(b.ETH1DepositIndex)

	// Field (11) 'Validators'
	{
		subIndx := hh.Index()
		num := uint64(len(b.Validators))
		if num > 1099511627776 {
			err = ssz.ErrIncorrectListSize
			return
		}
		for _, elem := range b.Validators {
			if err = elem.HashTreeRootWith(hh); err != nil {
				return
			}
		}
		hh.MerkleizeWithMixin(subIndx, num, 1099511627776)
	}

	// Field (12) 'Balances'
	{
		if size := len(b.Balances); size > 1099511627776 {
			err = ssz.ErrListTooBigFn("BeaconState.Balances", size, 1099511627776)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Balances {
			hh.AppendUint64(uint64(i))
		}
		hh.FillUpTo32()
		numItems := uint64(len(b.Balances))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(1099511627776, numItems, 8))
	}

	// Field (13) 'RANDAOMixes'
	{
		if size := len(b.RANDAOMixes); size != 65536 {
			err = ssz.ErrVectorLengthFn("BeaconState.RANDAOMixes", size, 65536)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.RANDAOMixes {
			hh.Append(i[:])
		}
		hh.Merkleize(subIndx)
	}

	// Field (14) 'Slashings'
	{
		if size := len(b.Slashings); size != 8192 {
			err = ssz.ErrVectorLengthFn("BeaconState.Slashings", size, 8192)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.Slashings {
			hh.AppendUint64(uint64(i))
		}
		hh.Merkleize(subIndx)
	}

	// Field (15) 'PreviousEpochParticipation'
	{
		if size := len(b.PreviousEpochParticipation); size > 1099511627776 {
			err = ssz.ErrListTooBigFn("BeaconState.PreviousEpochParticipation", size, 1099511627776)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.PreviousEpochParticipation {
			hh.AppendUint8(uint8(i))
		}
		hh.FillUpTo32()
		numItems := uint64(len(b.PreviousEpochParticipation))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(1099511627776, numItems, 1))
	}

	// Field (16) 'CurrentEpochParticipation'
	{
		if size := len(b.CurrentEpochParticipation); size > 1099511627776 {
			err = ssz.ErrListTooBigFn("BeaconState.CurrentEpochParticipation", size, 1099511627776)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.CurrentEpochParticipation {
			hh.AppendUint8(uint8(i))
		}
		hh.FillUpTo32()
		numItems := uint64(len(b.CurrentEpochParticipation))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(1099511627776, numItems, 1))
	}

	// Field (17) 'JustificationBits'
	if size := len(b.JustificationBits); size != 1 {
		err = ssz.ErrBytesLengthFn("BeaconState.JustificationBits", size, 1)
		return
	}
	hh.PutBytes(b.JustificationBits)

	// Field (18) 'PreviousJustifiedCheckpoint'
	if b.PreviousJustifiedCheckpoint == nil {
		b.PreviousJustifiedCheckpoint = new(phase0.Checkpoint)
	}
	if err = b.PreviousJustifiedCheckpoint.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (19) 'CurrentJustifiedCheckpoint'
	if b.CurrentJustifiedCheckpoint == nil {
		b.CurrentJustifiedCheckpoint = new(phase0.Checkpoint)
	}
	if err = b.CurrentJustifiedCheckpoint.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (20) 'FinalizedCheckpoint'
	if b.FinalizedCheckpoint == nil {
		b.FinalizedCheckpoint = new(phase0.Checkpoint)
	}
	if err = b.FinalizedCheckpoint.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (21) 'InactivityScores'
	{
		if size := len(b.InactivityScores); size > 1099511627776 {
			err = ssz.ErrListTooBigFn("BeaconState.InactivityScores", size, 1099511627776)
			return
		}
		subIndx := hh.Index()
		for _, i := range b.InactivityScores {
			hh.AppendUint64(i)
		}
		hh.FillUpTo32()
		numItems := uint64(len(b.InactivityScores))
		hh.MerkleizeWithMixin(subIndx, numItems, ssz.CalculateLimit(1099511627776, numItems, 8))
	}

	// Field (22) 'CurrentSyncCommittee'
	if b.CurrentSyncCommittee == nil {
		b.CurrentSyncCommittee = new(altair.SyncCommittee)
	}
	if err = b.CurrentSyncCommittee.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (23) 'NextSyncCommittee'
	if b.NextSyncCommittee == nil {
		b.NextSyncCommittee = new(altair.SyncCommittee)
	}
	if err = b.NextSyncCommittee.HashTreeRootWith(hh); err != nil {
		return
	}

	// Field (24) 'LatestExecutionPayloadHeader'
	if err = b.LatestExecutionPayloadHeader.HashTreeRootWith(hh); err != nil {
		return
	}

	hh.Merkleize(indx)
	return
}

// GetTree ssz hashes the BeaconState object
func (b *BeaconState) GetTree() (*ssz.Node, error) {
	return ssz.ProofTree(b)
}
