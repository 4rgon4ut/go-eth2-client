// Copyright © 2021 Attestant Limited.
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package multi

import (
	"context"

	consensusclient "github.com/attestantio/go-eth2-client"
	"github.com/attestantio/go-eth2-client/spec"
)

// SignedBeaconBlock fetches a signed beacon block given a block ID.
// N.B if a signed beacon block for the block ID is not available this will return nil without an error.
func (s *Service) SignedBeaconBlock(ctx context.Context,
	blockID string,
) (
	*spec.VersionedSignedBeaconBlock,
	error,
) {
	res, err := s.doCall(ctx, func(ctx context.Context, client consensusclient.Service) (interface{}, error) {
		block, err := client.(consensusclient.SignedBeaconBlockProvider).SignedBeaconBlock(ctx, blockID)
		if err != nil {
			return nil, err
		}
		return block, nil
	}, nil)
	if err != nil {
		return nil, err
	}
	if res == nil {
		return nil, nil
	}
	return res.(*spec.VersionedSignedBeaconBlock), nil
}
